package main

import (
	"fmt"
	"github.com/rjeczalik/notify"
	"log"
	"time"
)

func main() {
	coalesce := time.After(5 * time.Second)

	// evs := []notify.Event{ notify.InCreate, notify.InDelete, notify.InCloseWrite, notify.InMovedFrom}
	evs := []notify.Event{notify.Create, notify.Remove, notify.Write, notify.Rename}

	c := make(chan notify.EventInfo, 1)
	if err := notify.Watch(".", c, evs...); err != nil {
		log.Fatal(err)
	}
	defer notify.Stop(c)

	// Block until an event is received.
	ei := <-c
	log.Println("Got event:", ei)

	// Deadline
	select {
	case t := <-coalesce:
		fmt.Println("Time:", t)
	}
}

func build() {

}
