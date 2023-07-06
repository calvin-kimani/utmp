package utmp

//#include <utmp.h>
import "C"
import "time"
import "unsafe"

type exit_status struct {
    e_termination int16
    e_exit        int16
}

// man utmp
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
    Login     time.Time // added to make it easier getting the time
    AddrV6    [4]int // issue?
}

type Users struct {
    AllUsers []Utmp  // All users including logged out users
    LiveUsers []Utmp // Utmp Type 7 users, see man Utmo
}

// Read the utmp file
func ReadUtmp() Users {
    var users Users

    for {
        utmp := C.getutent() // read the current line("user") from the utmp file
        if utmp == nil {
            break
        }

        // // Converting ut_addr_v6 to [4]int32
        var addrV6 [4]C.int
        copy(addrV6[:], utmp.ut_addr_v6[:])
        addrV6Int32 := *(*[4]int)(unsafe.Pointer(&addrV6)) // weary...

        user := Utmp{
            Type:       int(utmp.ut_type),
            Pid:        int(utmp.ut_pid),
            Device:     C.GoString(&utmp.ut_line[0]),
            Id:         C.GoString(&utmp.ut_id[0]),
            Username:   C.GoString(&utmp.ut_user[0]),
            HostName:   C.GoString(&utmp.ut_host[0]),
            Exit:       exit_status{int16(utmp.ut_exit.e_termination), int16(utmp.ut_exit.e_exit)},
            Session:    int(utmp.ut_session),
            Seconds:    int64(utmp.ut_tv.tv_sec),
            MicroSecs:  int64(utmp.ut_tv.tv_usec),
            Login:      time.Unix(int64(utmp.ut_tv.tv_sec), 0),
            AddrV6:    addrV6Int32, // not quite good
        }

        users.AllUsers = append(users.AllUsers, user)
    }

    // get logged in (live) users, Type 7 -> man utmp
    for _, user := range users.AllUsers{
        if user.Type == 7{
            users.LiveUsers = append(users.LiveUsers, user)
        }
    }

    return users // Return Users
}