package main
import (
	"fmt"
)
func scanUserInput() {
	var appleNum int
	fmt.Println("Enter the Number of apples:")
	num, err := fmt.Scan(&appleNum)
	fmt.Println("Number of items:",num, " Error:",err)
	//_,_ = num, err
	fmt.Printf("Entered No. of Apples:%d\n",appleNum)
}

