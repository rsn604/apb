# About apb

**Japanese** is [here](README_JP.md) 

As an example of using the Golang-based TUI framework <a href="https://github.com/rsn604/taps" target="_blank">taps</a>, I created the appointment management application apb.

It's a typical planner, managing data by date and time. It can easily register one-time appointments, as well as recurring events by day, week, month, or year.

The screen layout and transitions are based on the Appointment Book, a built-in application for the legendary HP100LX and HP200LX.

--
## [1] Compiling

Compile with the following command:
```
go build -o apb main.go
```
Note that binaries for Linux and Windows are stored in "bin" directory.

---
## [2] Usage

```
apb <DB Name>
```
Start with apb <DB Name>. If the **DB Name** does not exist, a new one will be created.

--
### (1) Apptlist
Displays a list of today's schedules. The current month's calendar is displayed in the upper right corner, and the next schedule and **TodoList** are displayed below it.

![apptlist](/images/01apptlist.png)

Note that the following main screens have "**HELP**" buttons, so you can select and view them using the **F1** key or the mouse. In the following explanation, only function keys will be used; for other options, please refer to "HELP."

![apptlist](/images/01apptlist_help.png)

---
### (2) Data Registration
Selecting an appropriate time from the list screen will take you to the "**Data Registration**" screen.

![Detail](/images/02detail.png)

Enter the required information and press **F2**. When complete, "**D.MG Record added.**" should appear at the bottom.

![Detail](/images/02detail_add.png)

---
### (3) Repeat Settings
Once the record has been added, you can now set the "**Repeat Settings**." The **F8** key should appear in the guidance section at the bottom of the screen.
Pressing this key will display the screen below, where you can set the "Repeat Settings" for the appointment.

![Repeat](/images/03repeat_menu.png)

#### 1. Daily
![Repeat](/images/03repeat_daily.png)

As you can see from the screen, the **Duration** setting is set to 5 years by default, so please set it to an appropriate period.

#### 2. Weekly
![Repeat](/images/03repeat_weekly.png)

In the "Weekly Settings," you set the day of the week.

#### 3. Monthly
![Repeat](/images/03repeat_monthly.png)

In the "Monthly Settings," you set the date or day of the week.

#### 4. Yearly
![Repeat](/images/03repeat_yearly.png)

In the "Monthly Settings," you set the date or day of the week and month.

#### 5. Custom
![Repeat](/images/03repeat_custom.png)

In the "Custom Settings," more options are available.

--
### (4) View

#### 1. Weekly
Press **F8** from list view to switch to weekly view.

![View](/images/04wlist.png)

#### 2. Monthly
Press **F7** from list view to switch to monthly view.

![View](/images/04mlist.png)

---
### (5) Changing the Date

#### 1. Godate
Press **F8** from list view to display the monthly calendar.

![Calendar](/images/05godate.png)

Select a date from this screen or enter a date directly in the **Goto** field to switch to that date.

#### 2. Calendar (6 Months)
Press **F7** from list view to display the 6-month calendar.

![Calendar](/images/05calendar.png)

Select a date from this screen to switch to that date.

---
### (6) Delete
Pressing the **DEL** key on a list item will delete that item.
If "Repeat Setting" is enabled, select from the three options shown below. For a single schedule, the top two options will be unavailable.

![Delete](/images/06delete.png)

---
### (7) Todo
Pressing **F10** on a list item will display the **TodoList**. To add a new item, select the last blank item.

![Todo](/images/07todolist.png)

The registration screen will appear, allowing you to enter the required information. As with **Appointment**, "Repeat Setting" is also available.

![Todo](/images/07todo.png)

---
### (8) Utility
The **bin** directory contains binaries for **Linux** and **Windows**.

#### 1. Unload, Load
You can **Unload** and **Load** the appointment database.

```
apbunload <DB Name>
```
Displays data in **JSON** format. Please redirect as appropriate.

```
apbload <load file> <DB Name>
```
Loads data in JSON format. Note that if you specify an existing database, identical data will be registered twice. This is intended for registering new databases only.

#### 2. adbdump, hpload
This program extracts data from the HP100LX and 200LX and converts it to **apb**. However, please note that since I do not have the actual device, thorough testing has not been performed. Please use this information for reference only.

To extract data from the HP database, I used a tool called **ADBIO**, which is registered on the <a href="http://mizj.com/" target="_blank">S.U.P.E.R Site</a>.

>ADBIO is freeware and may be distributed/copied/used/modified/eaten/microwaved, etc. freely.</a>

As a result, I made some modifications to it so that it could be compiled with **gcc**.

First, save the HP data to a CSV file. The file name **appt.adb** may vary depending on your environment, but it represents the data within the HP database.

```
./adbdump2 -c -i atnr -d "yyyy/mm/dd" -q1 appt.adb >t1.txt
```
If the data contains double-byte characters, convert them to UTF-8 using a program such as **nkf**.

Next, load the data into **apb**.

```
./hpload t1.txt HPL.boltdb
```
Next, check the contents as shown below.

```
apb HPL.boltdb
```
