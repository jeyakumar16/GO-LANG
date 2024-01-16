//Simple Go Program to compare the compilation time between preallocation of capacity of slices and slices without preallocation
//using the time package

//time.Now() return the current monotonic clock readings
//time.Duration represents the time elapsed in nanoseconds
//time.Since(t) is equivalent to time.Now().Sub(t) --- provides the elapsed time since t

package main

import (
	"fmt"
	"time"
)

func main() {
	var n int = 100000
	var testSlice []int = []int{}
	var testSlice2 = make([]int, 0, n)
	fmt.Printf("Total time taken without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time taken with preallocation: %v\n", timeLoop(testSlice2, n))

}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
