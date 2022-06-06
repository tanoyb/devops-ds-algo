package main

import "fmt"

var arr = []int{237, 146, 259, 348, 152, 163, 235, 48, 36, 62, 6545}

type Node struct {
	data int
	next *Node
}

var bin = [10]*Node{} //automatically assigned with nil
//first take the last digit in the number, by doing mod 10,
//then again take the second last digit by doing mod 100
//and so on
//for a number 876(876/1=876), if we mod by 10, 876%10 we get 6 the last digit
// if we do 876/10 = 87 and 87%10 = 7, we get the second last number
// 876/100=8 , 8%10 = 8, we get the third last digit from the number

func radixSort() {
	//fmt.Println(bin)
	modValue := 10
	devideValue := 1 //1, 10, 100, 1000 .....goes not on each loop
	maxNum := findMax()
	pass := 1
	//56789 : get last digit -> (56789 / 1 ) % 10 = 9
	//get 2nd last digit  (56789 / 10 ) = 5678 % 10 = 8
	//get 3rd last digit (56789 / 100 ) = 567 % 10 = 7
	//get 4th last digit (56789 / 1000 ) = 56 % 10 = 6
	//get 5th last digit (56789 / 10000 ) = 5 % 10 = 5
	//now devidevalue will be 100000, which is greater than the number, exit the loop and all the elements should be sorted
	for devideValue < maxNum { //here we are calculating the number of pass
		//maximum number of pass depends on the max digits, 100 for 123, 1000 1234 and 10000 for 12345
		//so when the devideBy will be greter than max Num we will stop the loop
		//and meanwhile on the loop , the numbers should be sorted.
		arrInsertIndex := 0
		fmt.Println("\n=====PASS NO ", pass)
		pass++
		fmt.Println("devideValue=", devideValue, " maxNum=", maxNum)
		//insert into bin by last digits
		for i := 0; i < len(arr); i++ {
			//get the bin index
			binIndex := (arr[i] / devideValue) % modValue
			fmt.Println("bin index for ", arr[i], " is ", binIndex, " and devidevalue is ", devideValue)
			if bin[binIndex] == nil {
				//create a node and put it in the bin
				node := &Node{data: arr[i]}
				bin[binIndex] = node
			} else {
				headNode := bin[binIndex]
				for headNode.next != nil {
					headNode = headNode.next
				}
				tempnode := &Node{data: arr[i]}
				headNode.next = tempnode

			}

		}

		//move the elements from bin to arr

		printBin()

		for i := 0; i < len(bin); i++ {
			headNode := bin[i]
			for headNode != nil {
				arr[arrInsertIndex] = headNode.data
				arrInsertIndex++
				headNode = headNode.next
			}
			bin[i] = nil
		}

		fmt.Println("length after ", len(arr), " after each loop  arr = ", arr)
		//multiply devide value by 10 after each pass
		devideValue = devideValue * 10
	}

}

func findMax() int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func printBin() {
	fmt.Println(" ========================")
	for i := 0; i < len(bin); i++ {
		if bin[i] != nil {
			node := bin[i]
			for node != nil {
				fmt.Println("for bin ", i, " the values are: ")
				fmt.Println("\t ", node.data)
				node = node.next
			}
		}
	}
	fmt.Println(" ========================")
}

func main() {
	fmt.Println("length before ", len(arr), " after each loop  arr = ", arr)
	radixSort()

}
