package main

import "fmt"

func main(){
	var n int
	fmt.Print("n= ")
	fmt.Scan(&n)
	var arr [10]int
	for i:=0; i<n;i++{
		var numbers int
		fmt.Print("numbers= ")
		fmt.Scan(&numbers)
		arr[i]=numbers
	}
	fmt.Print("Количество четных элементов массива: ")
	x:=0
	for i:=0; i<n; i++{
		if arr[i]%2==0{
			x++;
		}
	}
	fmt.Println(x)
	fmt.Print("Количество нечетных элементов массива: ")
	y:=0
	for i:=0; i<n; i++{
		if arr[i]%2==1{
			y++;
		}
	}
	fmt.Print(y)



}

