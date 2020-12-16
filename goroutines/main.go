package main

import (
	"fmt"
	"time"
	"unsafe"
)

func printSleep(s string) {
	for index, stringVal := range s {
		fmt.Printf("%#U at index %d\n", stringVal, index)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	const t time.Duration = 14
	go printSleep("HELLO GOPHERS")
	time.Sleep(t * time.Millisecond)
	fmt.Println("sleep complete")
}

type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters
	// lock protects all field in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking
	lock mutex
}
