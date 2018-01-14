package main

import "fmt"

func main(){

	//self_pow:=1
	tot:=0

	for i:=1; i<=10;i++{
		self_pow:=1
		for n:=1;n<=i;n++{

			self_pow*=i

		}
		tot+=self_pow
	}

	fmt.Println(tot)

}
