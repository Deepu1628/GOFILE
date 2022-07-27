package main

import "fmt"

func main() {
	var monate = [12]string{"jan", "Feb", "Mar", "Apr", "May", "jun", "July", "AUG ", "Sep", "oct", "NOV", "Dec"}
	q1 := monate[0:3]
	fmt.Println(q1)

	fmt.Println("Length: ", len(q1), " \n Capacity:", cap(q1))
	q2 := monate[3:6]
	fmt.Println(q2)
	fmt.Println("Length: ", len(q2), " \n Capacity:", cap(q2))
	q3 := monate[6:9]
	fmt.Println(q3)
	fmt.Println("Length: ", len(q3), " \n Capacity:", cap(q3))
	q4 := monate[9:12]
	fmt.Println(q4)
	fmt.Println("Length: ", len(q4), " \n Capacity:", cap(q4))
}
