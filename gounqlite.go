// Package gounqlite provides access to the UnQLite C API, version 1.1.6.
package gounqlite

/*
#cgo LDFLAGS: -lunqlite
#include <unqlite.h>
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

type Errno int

func (e Errno) Error() string {
	s := errText[e]
	if s == "" {
		return fmt.Sprintf("errno %d", int(e))
	}
	return s
}

var (
	ErrLock     error = Errno(-76) // /* Locking protocol error */
	ErrReadOnly error = Errno(-75) // /* Read only Key/Value storage engine */
	ErrOpen     error = Errno(-74) // /* Unable to open the database file */
	ErrFull     error = Errno(-73) // /* Full database (unlikely) */
	ErrVM       error = Errno(-71) // /* Virtual machine error */
	ErrCompile  error = Errno(-70) // /* Compilation error */
	Done              = Errno(-28) // Not an error. /* Operation done */
	ErrCorrupt  error = Errno(-24) // /* Corrupt pointer */
	ErrNoOp     error = Errno(-20) // /* No such method */
	ErrPerm     error = Errno(-19) // /* Permission error */
	ErrEOF      error = Errno(-18) // /* End Of Input */
	ErrNotImpl  error = Errno(-17) // /* Method not implemented by the underlying Key/Value storage engine */
	ErrBusy     error = Errno(-14) // /* The database file is locked */
	ErrUnknown  error = Errno(-13) // /* Unknown configuration option */
	ErrExists   error = Errno(-11) // /* Record exists */
	ErrAbort    error = Errno(-10) // /* Another thread have released this instance */
	ErrInvalid  error = Errno(-9)  // /* Invalid parameter */
	ErrLimit    error = Errno(-7)  // /* Database limit reached */
	ErrNotFound error = Errno(-6)  // /* No such record */
	ErrLocked   error = Errno(-4)  // /* Forbidden Operation */
	ErrEmpty    error = Errno(-3)  // /* Empty record */
	ErrIO       error = Errno(-2)  // /* IO error */
	ErrNoMem    error = Errno(-1)  // /* Out of memory */
)

var errText = map[Errno]string{
	-76: "Locking protocol error",
	-75: "Read only Key/Value storage engine",
	-74: "Unable to open the database file",
	-73: "Full database",
	-71: "Virtual machine error",
	-70: "Compilation error",
	-28: "Operation done", // Not an error.
	-24: "Corrupt pointer",
	-20: "No such method",
	-19: "Permission error",
	-18: "End Of Input",
	-17: "Method not implemented by the underlying Key/Value storage engine",
	-14: "The database file is locked",
	-13: "Unknown configuration option",
	-11: "Record exists",
	-10: "Another thread have released this instance",
	-9:  "Invalid parameter",
	-7:  "Database limit reached",
	-6:  "No such record",
	-4:  "Forbidden Operation",
	-3:  "Empty record",
	-2:  "IO error",
	-1:  "Out of memory",
}

type Conn struct {
	db *C.unqlite
}

func Open(filename string) (*Conn, error) {
	// TODO(ceh): default to thread-safe operation only?
	var db *C.unqlite
	name := C.CString(filename)
	defer C.free(unsafe.Pointer(name))
	rv := C.unqlite_open(&db, name, C.UNQLITE_OPEN_CREATE)
	if rv != C.UNQLITE_OK {
		return nil, errors.New(Errno(rv).Error())
	}
	if db == nil {
		return nil, errors.New("unqlite unable to allocate memory to hold the database")
	}
	return &Conn{db}, nil
}

func (c *Conn) Close() error {
	if c == nil || c.db == nil {
		return errors.New("nil unqlite database")
	}
	rv := C.unqlite_close(c.db)
	if rv != C.UNQLITE_OK {
		return errors.New(Errno(rv).Error())
	}
	c.db = nil
	return nil
}

func Threadsafe() bool {
	return C.unqlite_lib_is_threadsafe() == 1
}

func Version() string {
	p := C.unqlite_lib_version()
	return C.GoString(p)
}

func Signature() string {
	p := C.unqlite_lib_signature()
	return C.GoString(p)
}

func Ident() string {
	p := C.unqlite_lib_ident()
	return C.GoString(p)
}

func Copyright() string {
	p := C.unqlite_lib_copyright()
	return C.GoString(p)
}
