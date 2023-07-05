# Package utmp

The utmp package provides functionality to read and manipulate utmp entries in Go. It is a wrapper for the C function `getuntent()`. The package includes structs and methods to work with the utmp file, which contains information about user login sessions and terminated processes.

It is recommended to refer to the utmp manual **(man utmp)** for a deeper understanding of the utmp entries and their fields.

## Types
### type `Utmp`

Type Utmp is a struct that represents a utmp entry. It contains various fields to store information such as the entry type, process ID, device, user ID, username, hostname, exit status, session ID, timestamp, login time, and IPv6 address.

```go
type Utmp struct {
	Type      int
	Pid       int
	Device    string
	Id        string
	Username  string
	HostName  string
	Exit      exit_status
	Session   int
	Seconds   int64
	MicroSecs int64
	Login     time.Time
	AddrV6    [4]int
}
```

### type `Users`
Type Users holds all the historical data since boot time which includes logged out users and terminated processes. This data is stored in `AllUsers`. To filter out the historical data to show only the currently logged in users, `LiveUsers` has been defined.

```go
type Users struct {
    AllUsers []Utmp  // All users including logged out users
    LiveUsers []Utmp // Utmp Type 7 users, see man Utmo
}
```

### type `exit_status`
Type `exit_status` is a struct that represents the exit status of a terminated process. It has two fields: e_termination and e_exit, which store the termination and exit values, respectively.

```go
type exit_status struct {
	e_termination int16
	e_exit        int16
}
```

## Functions
### ReadUtmp

ReadUtmp reads the utmp file and returns type `Users`.

```go
func ReadUtmp() Users {
	...
}
```