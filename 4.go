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
	maxi:=arr[0]
	for i:=0; i<n; i++{
		if arr[i]>maxi{
			maxi=arr[i]
		}
	}
	fmt.Println("Максимальный элемент равен: ",maxi)
	mini:=arr[0]
	for i:=0; i<n; i++{
		if arr[i]<mini{
			mini=arr[i]
		}
	}
	fmt.Println("Минимальный элемент равен: ",mini)
	raznica := maxi-mini
	fmt.Print("Разница между максимальным и минимальным элементом равна: ", raznica)

}
