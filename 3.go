package main

import "fmt"

func main(){
	var n int
	fmt.Print("n= ")
	fmt.Scan(&n)
	var arr [10]int
	for i:=0; i<n; i++{
		var numbers int
		fmt.Print("numbers= ")
		fmt.Scan(&numbers)
		arr[i]=numbers
	}
	sumi:=0;
	for i:=0; i<n; i++{
		sumi= sumi+arr[i]
	}
	avg:=sumi/n;
	fmt.Println(avg)
	x:=0
	for i:=0; i<n; i++{
		if avg>arr[i]{
			x++;
		}
	}
	fmt.Print("Количество элементов, которые меньше среднего арифметического равна: ", x)
}
