package main

import "fmt"
import (
	"io/ioutil"
	"os" // More controlled access to File Operations as compared with ioutil package
)


func main() {
	fmt.Println("Hello, Welcome to File Operations")
	//A VERY in-efficient way of reading the FULL FILE into memory
	//A MORE Efficient way is to use "os" package 
	dat, err :=
		ioutil.ReadFile("test.txt")
	fmt.Printf("dat Type:%T err:%T\n",dat,err)
	fmt.Printf("Length of Byte Array read:%d\n",len(dat))
	if err == nil {
		
		//Approach #1
		str := string(dat)
		fmt.Println(str)

		//Approach #2
		//Using the rune approach:
		for index,runeChar := range str {
			fmt.Printf("Index:%d is having:%c\n",index,runeChar)
		}

		//Approach #3
		//Print it directly from the []uint8
		fmt.Printf("File Read:%s", dat)
	}

	fmt.Println("File Writting:")
	dat = []byte("Write this string into a File\n")
	ioutil.WriteFile("outfile.txt", dat, 0777)

	readFileEfficiently()

	writeFileEfficiently()
}

/*
 	* Using os package read the File only as much as you NEED
 	* os.Open opens a file, and returns a 
 		File descriptor(A File struct, used to access the file)
 	* os.Close closes a file once you are done
 	* os.Read : Reads from a file into a []byte
 		You can control the amount of bytes read
 	* os.Write : Writes a []byte into a file
*/
func readFileEfficiently() {
	//Step 1: Open a File, and track the file descriptor
	fDescriptor,err1 := os.Open("test.txt")
	if err1 != nil {
			fmt.Println("Error while Opening the file")
			return
	}
	byteArray := make([]byte, 5)
	for {
		nb, err2 := fDescriptor.Read(byteArray)
		if err2 != nil {
			fmt.Println("Error while Reading the file. Error:", err2)
			break
		}
		if nb<cap(byteArray) {
			fmt.Println("Reaching End of File")
			fmt.Println("Writing the last", cap(byteArray) ,"bytes")
			fmt.Printf("No. of bytes read:%d type:%T\n", nb,byteArray)
			for idx:=0; idx<nb; idx++ {
				fmt.Printf(" Index:%d -> %d ",idx,byteArray[idx])
			} 
			//break
		}
	}
	fDescriptor.Close()
}

func writeFileEfficiently() {
	f, err := os.Create("outfileTwo.txt")
	if err == nil {
		bSlice := []byte{1,2,3,10}
		nb, err := f.Write(bSlice)
		fmt.Println("No of bytes written:", nb,"Error:", err)
		nb2, err2 := f.WriteString("Hi")
		_,_ = nb2, err2
		f.Close()
	}
	
}