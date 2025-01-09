# Go Map Access: Panic on nil map

This repository demonstrates a common error in Go when working with maps: accessing a key in a nil map. Attempting to access an uninitialized map results in a runtime panic.  This example shows how to avoid the panic by checking if the map is nil before accessing it.

## Bug
The `bug.go` file contains code that panics when trying to access an element of an uninitialized map.

## Solution
The `bugSolution.go` file demonstrates how to safely access map elements by checking for nil maps.