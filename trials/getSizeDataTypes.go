package main
import(
	"fmt"
	"unsafe"
	"math"
	"strings"
)

func main() {
	var x int
	y := int64(123)
	z := int32(123)
	z1 := 123
	f1 := float32(1234.0)
	//Scientific notation
	f2 := 1.2345e-2

	fmt.Printf("***I AM INSIDE getSizeDataTypes.go FILE***\n")
	fmt.Printf("x: %T :%d bytes\n",x,unsafe.Sizeof(x))
	fmt.Printf("y: %T :%d bytes\n",y,unsafe.Sizeof(y))
	fmt.Printf("z: %T :%d bytes\n",z,unsafe.Sizeof(z))
	fmt.Printf("z1: %T :%d bytes\n",z1, unsafe.Sizeof(z1))
	fmt.Printf("f1: %T :%d bytes\n",f1, unsafe.Sizeof(f1))
	fmt.Printf("f2: %T :%d bytes\n",f2, unsafe.Sizeof(f2))
	fmt.Printf("f2: %f\n",f2)
	fmt.Printf("ASCII value of A=%d\n",'A')

	/* Max of types*/
	fmt.Printf("max float32 = %f\n", math.MaxFloat32)
	fmt.Printf("max int = %d\n", math.MaxInt)
	f3 := 9223372036854775807.0
	fmt.Printf("Max Int value assigned to f3:%f\n",f3)
	/*

	*Invoking the function defined in the file: hello.go
	*For this to work, please run using the command:
	*go build getSizeDataTypes.go hello.go

	*/
	fmt.Printf("Calling Function Defined in another file in the same package\n")

	hellomain()

	//STRINGS
	var r1,r2 rune
	r1 = 1
	r2 = 0
	fmt.Println("Rune:",r1)
	fmt.Println(strings.Compare(string(r1),string(r2)))

	//Try Scan
	//scanUserInput()

	//Test Arrays
	learnArrays()
}

func learnArrays() {
	var unsetArray [5]int 
	for index,item := range unsetArray {
		fmt.Printf("Index:%d is having:%d\n",index,item)
	}

	var initializedArray [5]int = [5]int{1,2,3,4,5}
	for index,item := range initializedArray {
		fmt.Printf("Index:%d is having:%d\n",index,item)
	}

}