package main

import (
	"fmt"
	"github.com/NathanFrasier/sorters"
	"math/rand"
	"time"
)

func isSorted(data []int) bool{
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false;
		}
	}
	return true;
}

func main() {
	const size = 1000000;
	//fmt.Println("Setup Started")
	//create array that holds original values. Once this array is populated it is never modified
	var original[size] int	//may need to be bigger
	rand.Seed(0)
	for i := 0; i < size; i++ {
		//original[i] =  rand.Int()
		original[i] = size - i
	}
	//fmt.Println("Setup Complete")
	for i, j := 10, 10;i <= size; {
		//create copy of first i elements into slice
		data := make([]int, i);
		copy(data, original[:])
		var start, end int64;
		//get start time
		start = time.Now().UnixNano()
		//run test on copied slice
		sorters.Mrginssort(data)
		//get end time
		end = time.Now().UnixNano()
		
		fmt.Print(i)
		fmt.Print(",")
		fmt.Println(((end - start)*int64(time.Nanosecond)))
		//verify the list is sorted
		if !isSorted(data) {
			fmt.Println("it is not sorted!")
		}
		//logarithmic increment
		if i/j == 10 {
			j = j * 10;
		}
		i = i + j;
	}
}