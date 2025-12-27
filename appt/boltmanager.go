package appt

import (
	"encoding/binary"
	"encoding/json"
	//"fmt"
	"github.com/boltdb/bolt"
)

// -------------------------------------------------
type BoltManager struct {
	Db *bolt.DB
}

const (
	APPTDB = "ApptDB"
)

// -------------------------------------------------
/*
func (self *BoltManager) PrintAppts() {
	appts := self.GetAppts()
	for _, ap := range appts {
		fmt.Printf("id:%d repeat:%d desc:%s sdate:%s stime:%s etime:%s location:%s ap.ConDays:%d frequency:%d DayNumber:%d DayPosition:%d DayOfWeek:%d Month:%d DateOfYear:%s RepeatSdate:%s RepeatEdate:%s\n", ap.Id, ap.RepeatType, ap.Description, ap.Sdate, ap.Stime, ap.Etime, ap.Location, ap.ConDays, ap.Frequency, ap.DayNumber, ap.DayPosition, ap.DayOfWeek, ap.Month, ap.DateOfYear, ap.RepeatSdate, ap.RepeatEdate)
	}
}
*/

// -------------------------------------------------
func (self *BoltManager) GetDb() *bolt.DB {
	return self.Db
}

func (self *BoltManager) Connect(databaseName string, connectString string) error {
	var err error
	self.Db, err = bolt.Open(connectString, 0600, nil)
	if err != nil {
		return err
	}
	return err
}

func (self *BoltManager) Close() {
	self.GetDb().Close()
}

func (self *BoltManager) Define() error {
	err := self.GetDb().Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(APPTDB))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (self *BoltManager) itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func (self *BoltManager) AddAppt(ap Appointment) (int, error) {
	err := self.GetDb().Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(APPTDB))
		seq, _ := b.NextSequence()
		ap.Id = int(seq)
		buf, err := json.Marshal(ap)
		if err != nil {
			return err
		}
		b.Put(self.itob(ap.Id), buf)
		return nil
	})
	return ap.Id, err
}

func (self *BoltManager) UpdateAppt(ap Appointment, id int) error {
	err := self.GetDb().Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(APPTDB))
		buf, err := json.Marshal(ap)
		if err != nil {
			return err
		}
		b.Put(self.itob(id), buf)
		return nil
	})
	//return nil, err
	return err

}

func (self *BoltManager) DeleteAppt(id int) error {
	err := self.GetDb().Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(APPTDB))
		b.Delete(self.itob(id))
		return nil
	})
	return err
}

func (self *BoltManager) GetAppts() []Appointment {
	var appts []Appointment
	_ = self.GetDb().View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(APPTDB))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			ap := Appointment{}
			if err := json.Unmarshal(v, &ap); err != nil {
				return err
			}
			appts = append(appts, ap)
		}
		return nil
	})
	return appts
}
func (self *BoltManager) GetApptByID(id int) Appointment {
	//var ap Appointment
	ap := Appointment{}
	_ = self.GetDb().View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(APPTDB))
		v := b.Get([]byte(self.itob(id)))
		if err := json.Unmarshal(v, &ap); err != nil {
			return err
		}
		return nil
	})
	return ap
}
