package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

var wrCapacity = 4
var shopOpenDuration = 25
var newCustInterval = 1
var custServiceDuration = 4
var numBarbers = 2

var waiting = make(chan int, wrCapacity) // waiting room
var wrClosed = false

var servedCust int
var unservedCust int
var numWaiting int

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	wg.Add(1)
	go shop()
	time.Sleep(3 * time.Second)
	go customers()
	wg.Wait()

	fmt.Println()
	color.Cyan(fmt.Sprintf("Customers SERVED: %v\n", servedCust))
	color.Red(fmt.Sprintf("Customers PASSED: %v\n", unservedCust))
	fmt.Println()
}

func shop() {
	// open shop
	color.Cyan("\n= = = Shop OPEN! = = =\n\n")

	// setup lifetime
	go func() {
		time.Sleep(time.Duration(shopOpenDuration) * time.Second)
		color.Cyan(fmt.Sprintln("\n= = = Waiting Room closed! = = ="))
		wrClosed = true
	}()

	var innerWg sync.WaitGroup
	innerWg.Add(numBarbers)
	// create barbers and init barbers
	for i := 0; i < numBarbers; i++ {
		go barber(i+1, &innerWg)
	}
	innerWg.Wait()
	wg.Done()
	color.Cyan(fmt.Sprintln("\n= = = Shop CLOSED! = = ="))
}

func barber(barberNo int, bwg *sync.WaitGroup) {
	barberAsleep := true

	// no specific strategy: letting go scheduler assign clients to barbers (regardless of sleeping status or accumulated workload)
	for nextCust := range waiting {
		pad := strings.Repeat("\t", barberNo*3)

		if barberAsleep {
			barberAsleep = false
			color.Blue(fmt.Sprintf("%v🔔 Barber #%v wakes up...\n", pad, barberNo))
			time.Sleep(50 * time.Millisecond) // wake up time
		}
		time.Sleep(100 * time.Millisecond) // invite customer into the barber chair
		color.Cyan(fmt.Sprintf("%v🪒 Serving 👨 #%v...\n", pad, nextCust))
		mu.Lock()
		numWaiting--
		printCustomers()
		mu.Unlock()
		time.Sleep(time.Duration(custServiceDuration) * time.Second)
		color.Cyan(fmt.Sprintf("%v✅ Finished 👨 #%v...\n", pad, nextCust))
		if numWaiting == 0 {
			barberAsleep = true
			color.Blue(fmt.Sprintf("%v💤 Barber #%v snoozing...\n", pad, barberNo))
		}
	}
	bwg.Done()
}

func customers() {
	custNo := 0

	for {
		custNo++
		if wrClosed {
			close(waiting)
			break
		}
		select {
		case waiting <- custNo:
			servedCust++
			fmt.Printf(" 👨 #%v ENTER\n", custNo)
			mu.Lock()
			numWaiting++
			printCustomers()
			mu.Unlock()
		default:
			unservedCust++
			color.Red(fmt.Sprintf("\t\t\t\t👨 #%v PASS\n", custNo))
		}
		time.Sleep(time.Duration(newCustInterval) * time.Second)
		if custNo == wrCapacity*numBarbers+2 {
			newCustInterval = newCustInterval * 5
		}
	}
}

func printCustomers() {
	color.Green(fmt.Sprintf(" WAITING ROOM  * * *  -%v- \n", strings.Repeat(" |👨| ", numWaiting)))
}
