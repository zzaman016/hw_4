package main

import "fmt"

func main() {
	var n int
	fmt.Print("n= ")
	fmt.Scan(&n)
	var arr [5]int
	for i:=0; i<n; i++{
		var num int
		fmt.Print("num= ")
		fmt.Scan(&num)
		arr[i]=num
	}
	fmt.Println(arr)
	var leftPart[2]int
	j:=0
	for i:=0; i<n/2; i++{
		leftPart[j]=arr[i]
		j++
	}
	fmt.Println(leftPart)
	var rightPart[3]int
	k:=0
	for i:=2; i<n; i++{
		rightPart[k]=arr[i]
		k++
	}
	fmt.Print(rightPart)
}

