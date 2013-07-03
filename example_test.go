package gounqlite_test

import (
	"fmt"
	"gounqlite"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ExampleOpen() {
	f, err := ioutil.TempFile("./", "unqlite_")
	check(err)
	defer os.Remove(f.Name())
	// Open file database.
	db, err := gounqlite.Open(f.Name())
	check(err)
	db.Close()

	// Open in memory database.
	db, err = gounqlite.Open(":mem:")
	check(err)
	db.Close()
}

func ExampleClose() {
	db, err := gounqlite.Open(":mem:")
	check(err)

	err = db.Close()
	check(err)

	err = db.Close()
	fmt.Println(err)
	// Output: nil unqlite database
}

func ExampleConn_Store() {
	db, err := gounqlite.Open(":mem:")
	check(err)
	defer db.Close()

	db.Store([]byte("key"), []byte("value"))
}

func ExampleConn_Append() {
	db, err := gounqlite.Open(":mem:")
	check(err)
	defer db.Close()

	db.Append([]byte("key"), []byte{'a'})
	v, err := db.Fetch([]byte("key"))
	check(err)
	fmt.Println(v)

	db.Append([]byte("key"), []byte{'b'})
	v, err = db.Fetch([]byte("key"))
	check(err)
	fmt.Println(v)

	// Output: [97]
	// [97 98]
}

func ExampleConn_Fetch() {
	db, err := gounqlite.Open(":mem:")
	check(err)
	defer db.Close()

	v, err := db.Fetch([]byte("key"))
	fmt.Println(err)

	err = db.Store([]byte("key"), []byte{'a'})
	v, err = db.Fetch([]byte("key"))
	check(err)
	fmt.Println(v)

	// Output: No such record
	// [97]
}

func ExampleConn_Delete() {
	db, err := gounqlite.Open(":mem:")
	check(err)
	defer db.Close()

	err = db.Store([]byte("key"), []byte{'a'})
	v, err := db.Fetch([]byte("key"))
	check(err)
	fmt.Println(v)

	err = db.Delete([]byte("key"))
	check(err)

	v, err = db.Fetch([]byte("key"))
	fmt.Println(err)

	// Output: [97]
	// No such record
}
