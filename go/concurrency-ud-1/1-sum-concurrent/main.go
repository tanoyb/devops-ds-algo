package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func SumSequential(nums []int64) int64 {
	var sum int64
	for _, n := range nums {
		sum += n
	}
	return sum
}

func SumConcurrent(nums []int64) int64 {
	//Utilize all cores in machine
	numOfCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numOfCores)
	fmt.Println("====number of cores==", numOfCores)
	var sum int64
	max := len(nums)
	sizeOfParts := max / numOfCores
	var wg sync.WaitGroup

	for i := 0; i < numOfCores; i++ {
		//devide the input into parts
		start := i * sizeOfParts
		end := start + sizeOfParts
		part := nums[start:end]

		if i == numOfCores-1 {
			fmt.Println("taking the rest of the array")
			part = nums[start:]
		}
		fmt.Println("length of ful array = ", len(nums),
			" length of part i=", i, " is ", len(part))
		//run compuation in each part in separate goroutine
		wg.Add(1)
		go func(nums1 []int64) {
			defer wg.Done()
			var partSum int64
			//calculate the sum for each part
			for _, n := range nums1 {
				partSum += n
			}
			//add sum of each part to cumulative sum
			atomic.AddInt64(&sum, partSum)
		}(part)

	}
	wg.Wait()

	return sum
}

func main() {
	nums := make([]int64, 300000000)
	for i := 0; i < 300000000; i++ {
		nums[i] = int64(i)
	}
	t := time.Now()
	sum := SumSequential(nums)
	fmt.Println("sequencial add, sum =", sum, " \n time taken = ", time.Since(t))

	t1 := time.Now()
	sum = SumConcurrent(nums)
	fmt.Println("sequencial concurrent, sum =", sum, " \n time taken = ", time.Since(t1))

}
