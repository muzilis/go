package main

import "fmt"

func main(){
	a :=20  
	var ip *int
	ip = &a
	fmt.Printf("a 变量的地址是: %x\n", &a  )
	 /* 指针变量的存储地址 */
	 fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

	 /* 使用指针访问值 */ 
	 fmt.Printf("*ip 变量的值: %d\n", *ip )
}