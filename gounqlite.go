// Package gounqlite provides access to the UnQLite C API, version 1.1.6.
package gounqlite

/*
#cgo LDFLAGS: -lunqlite
#include <unqlite.h>
*/
import "C"

func Threadsafe() bool {
	return int(C.unqlite_lib_is_threadsafe()) == 1
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
