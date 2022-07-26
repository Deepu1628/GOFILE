package main

import (
    "fmt"
	"strings"
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
	var c uint = 500// with unit only positive value are stored 
	var d uint = 4500
	fmt.Printf("Type: %T, value: %v", c, d)
	fmt.Printf("Type: %T, value: %v", c, d)
//*********float ****
var a float32 = 123.78
  var b float32 = 3.4e+38
  fmt.Printf("Type: %T, value: %v\n", a, b)
  fmt.Printf("Type: %T, value: %v", a, b)
  //******************
  
  //arithmetic operators are there in Go 

  cat := 10
  dog:= 2

  fmt.Println(cat +dog)
// take two strings and concatenate them using Join() function, with single space as a separator.
  var str1 = "Hello GO"
    var str2 = "You are new to me "
    var output = strings.Join([]string{str1, str2}, " ")
    fmt.Println(output)
//four strings in an a string array and joined them with the separator: comma ,.
var w = []string{"dog", "cat", "pig", "monkey"}
var sep = ","
var answer = strings.Join(w, sep)
fmt.Println(answer)
}