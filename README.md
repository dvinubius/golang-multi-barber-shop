# Multi Barber Shop in Go

The classic CS problem by Dijkstra. 

Version with multiple barbers 🧔🏻 🧔🏻 🧔🏻

## Requirements 

Customers come to the barber shop during business hours.
If waiting room (fixed capacity) is not full, a customer enters the waiting room
If waiting room is full, customer goes somewhere else.
There can be multiple barbers in the shop
A barber is asleep as long as there are no customers

## Simulation

First the shop opens, then customers begin to come.
The shop has a set time for business hours. After that time passes, no more customers attempt to visit.

Solved via **channels** and a **mutex**.

Pretty terminal output 🤩