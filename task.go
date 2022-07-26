package main

import (
    "fmt"
)

func main() {
// variable ************************* DEFINED WITH ALfabet 
    var i, j, k , q = 1, 2, 3 , 6// variable with value 

    var (
        name       = "Deepa Kumar "
        occupation = " Teacher  ")
		

    fmt.Println(i, j , k, q )
    fmt.Printf("%s is a %s\n" ,name, occupation)
//******************CONSTANT WITH numbers or value
	var age int = 35// variable 
    const  WIDTH = 100// CONSTANT are written in uppercase 

    age = 39
    age = 38
	age = 67// it takes the last value 

    // WIDTH = 101

    fmt.Println(" constant and variable :" , age, WIDTH)
	//****************** boolean datatype ************************************************//
	var b1 bool = true // typed declaration with initial value
	var b2 = true // untyped declaration with initial value
	var b3 bool // typed declaration without initial value
	b4 := true // untyped declaration with initial value
  // though yo do not difine the datatype  compiler take the data type 
	fmt.Println(b1) // Returns true
	fmt.Println(b2) // Returns true
	fmt.Println(b3) // Returns false
	fmt.Println(b4) // Returns true
	//************** integer*****
	var x int = 60// positive 
	var y int = -6666// negative
	fmt.Printf("The value and type is displayed: %T , value: %v", x, x)//%T Type of datatype 
	fmt.Printf("The type and value is displayed : %T, value: %v", y, y)//%v Value if the data type 
//********************UNSigned integer with unit **************************7//
	var x uint = 500// with unit only positive value are stored 
	var y uint = 4500
	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)
//*********float ****
var x float32 = 123.78
  var y float32 = 3.4e+38
  fmt.Printf("Type: %T, value: %v\n", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
  //******************
  var txt1 string = "Hello!"
  var txt2 string
  txt3 := "STRING WORD "

  fmt.Printf("Type: %T, value: %\n", txt1, txt1)
  fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
  fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
  //





}