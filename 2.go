package main

import "fmt"

func main(){
	var n int
	fmt.Print("n= ")
	fmt.Scan(&n)
	var arr [10]int
	for i:=0; i<n; i++{
		var number int
		fmt.Print("number= ")
		fmt.Scan(&number)
		arr[i]=number
	}
	sumi:=0;
	for i:=0; i<n; i++{
		if arr[i]<0{
			continue
		}
		sumi = sumi + arr[i]
	}
	fmt.Print(sumi)
}
