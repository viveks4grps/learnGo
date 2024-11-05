package main

import "fmt"
import (
	"io/ioutil"
	//"os"
)


func main() {
	fmt.Println("Hello, Welcome to File Operations")
	//A VERY in-efficient way of reading the FULL FILE into memory
	dat, err :=
		ioutil.ReadFile("test.txt")
	fmt.Printf("dat Type:%T err:%T\n",dat,err)
	if err == nil {
		
		//Approach #1
		str := string(dat)
		fmt.Println(str)

		//Approach #2
		//Using the rune approach:
		/*for index,runeChar := range str {
			fmt.Printf("Index:%d is having:%c\n",index,runeChar)
		}*/

		//Approach #3
		//Print it directly from the []uint8
		fmt.Printf("File Read:%s", dat)
	}

	fmt.Println("File Writting:")
	dat = []byte("Write this string into a File\n")
	ioutil.WriteFile("outfile.txt", dat, 0777)
	
}
