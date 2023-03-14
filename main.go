// file ini digunakan sebagai file utama yg dijalankan golang

// standart minimum suatu project
// 1. package main
// 2. func mmain
package main

import "fmt"

func main()  {
	// fmt --> build in package yg biasa digunakan
	//untuk memunculkan output ke terminal

	//syntax
	// fmt.Println("Hello World sayang")

	//variable
	
	// var1 := 100 // var1 merupakan variable yang menampung value 100
	// var2 := "this is var2" //var 2 merupakan variable yang menampung value bertipe string
	// var3 := uint64(100)

	// var4 := var1 + int(var3) // type cast adalah mengubah tipe dari variable
	// fmt.Printf("var1:%d var2%s var3%d var4%s",var1,var2,var3,var4)

	//challange 1
	var i int = 21
	var j bool = true;

	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%% \n")
	fmt.Printf("%t \n\n", j)
	fmt.Printf("%b \n", i)
	fmt.Printf("%c \n", '\u042F')
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)
	fmt.Printf("%x \n", 15)
	fmt.Printf("%X \n", 15)
	fmt.Printf("%U \n\n", 'Я')
	
	var k float64 = 123.456;
	fmt.Printf("%f \n", k)
	fmt.Printf("%E \n", k)
}