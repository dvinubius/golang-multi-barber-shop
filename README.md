# Multi Barber Shop in Go

The classic CS problem by Dijkstra. 

Version with multiple barbers 🧔🏻 🧔🏻 🧔🏻

Tweak the params to see varying degrees of efficiency (customers served / customers arrived)

## Requirements 

Customers come to the barber shop during business hours.
If waiting room (fixed capacity) is not full, a customer enters the waiting room
If waiting room is full, customer goes somewhere else.
There can be multiple barbers in the shop
A barber is asleep as long as there are no customers

## Simulation

First the shop opens, then customers begin to come.
The shop has a set time for business hours. After that time passes, no more customers attempt to visit.

Solved via **channels**, **wait groups** and a **mutex**

- Shop is a go routine
- Every barber is a go routine
- Arrival of customers is go routine
- Waiting room is a buffered channel

### Asymmetric (unfair) resource allocation

While making full use of the automatic scheduling provided by Go (code simplicity), there is **no specific strategy for dispatching** incoming customers to existing barbers. So the shop may have unfair work distribution among barbers. 

e.g. If there are so few customers that a single barber can handle them, all the other barbers may be asleep all the time.

> Pretty terminal output 🤩