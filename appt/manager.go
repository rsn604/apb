package appt

// ----------------------------------------------------------------------
type Manager interface {
	Connect(databaseName string, connectString string) error
	Define() error
	AddAppt(u Appointment) (int, error)
	UpdateAppt(u Appointment, id int) error
	DeleteAppt(id int) error
	GetAppts() []Appointment
	GetApptByID(id int) Appointment
	//PrintAppts()
	Close()
}

func GetManager(name string) Manager {
	if name == "BOLT" {
		return new(BoltManager)
	} else {
		//return new(ListDBManager)
	}
	return nil
}
