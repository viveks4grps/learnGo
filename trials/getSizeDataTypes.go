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
	/*syntax error: unexpected :=, expected =
	//var xInt int := 4*/
	xInt := 4
	_ = xInt

	f1 := float32(1234.0)
	//Scientific notation
	f2 := 1.2345e-2
	const aConstant = 10

	fmt.Printf("***I AM INSIDE getSizeDataTypes.go FILE***\n")
	fmt.Printf("x: %T :%d bytes\n",x,unsafe.Sizeof(x))
	fmt.Printf("y: %T :%d bytes\n",y,unsafe.Sizeof(y))
	fmt.Printf("z: %T :%d bytes\n",z,unsafe.Sizeof(z))
	fmt.Printf("z1: %T :%d bytes\n",z1, unsafe.Sizeof(z1))
	fmt.Printf("f1: %T :%d bytes\n",f1, unsafe.Sizeof(f1))
	fmt.Printf("f2: %T :%d bytes\n",f2, unsafe.Sizeof(f2))
	fmt.Printf("f2: %f\n",f2)
	fmt.Printf("ASCII value of A=%d\n",'A')
	fmt.Printf("Constant Value=%d\n",aConstant)

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

	//Test For loops
	learnForLoops()

	//Test Arrays
	learnArrays()

	//Test Strings
	exploreStrings()

	//Test Slices
	learnSlices()

	//Test HashMaps
	learnMaps()

	//Test Struct (Structure)
	learnStruct()

	//Test multi-return Value Function
	a,b,c := callFunc2(1,"STR2")
	fmt.Printf("***Type of a:%T, Type of b:%T, Type of c:%T\n",a,b,c)
}

func learnForLoops() {
	for i:=0; i<10; i++ {
		//Do something
		//fmt.Println("For loop1")
	}

	i := 0
	for i<10 {
		//Do something
		i++
		//fmt.Println("For loop2")
	}

	for {
		//fmt.Println("For loop3")
		break
	}

}

func learnArrays() {
	var unsetArray [5]int 
	//Easy way of traversal using range expression
	fmt.Println("Using a range Expression based for loop for array traversal:")
	for index,item := range unsetArray {
		fmt.Printf("Index:%d is having:%d\n",index,item)
	}
	//Without accessing the item inside the range expression
	fmt.Println("Using a range Expression WITHOUT item:")
	for index := range unsetArray {
		fmt.Printf("Index:%d is having:%d\n",index,unsetArray[index])
	}

	//Another way of traversal using the normal for loop:
	fmt.Println("Using a normal for loop for array traversal:")
	for index:=0; index < len(unsetArray); index++ {
		fmt.Printf("Index:%d is having:%d\n",index,unsetArray[index])
	}

	/*
	ARRAY LITERAL using type + curly brackets
	//syntax error: unexpected {, expected type
	var arrayLiteralWithInferredType [2]int = [2]{3,4}
	*/

	var initializedArray [6]int = [6]int{1,2,3,4,5,6}
	fmt.Println("Using an array LITERAL, during Initialization:")
	for index,item := range initializedArray {
		fmt.Printf("Index:%d is having:%d\n",index,item)
	}

	//How about using the KEY-WORD "..."
	/* An ellipsis is a series of three dots (...) used to indicate the 
	deliberate omission of a word or words 
	from a sentence without preventing the meaning from being understood. 
	It can also imply a pause or trail off in thought.
	*/
	var arrWithEllipsis [2]int = [...]int{3,4}
	fmt.Println("Array with ellipsis:")
	for index,item := range arrWithEllipsis {
		fmt.Println("index:",index,"item:",item)
	}
	fmt.Println("Arrays are mutable")
	arrWithEllipsis[1] = 5
	for index,item := range arrWithEllipsis {
		fmt.Println("index:",index,"item:",item)
	}

}

func exploreStrings() {

	simpleString := "abcde"
	fmt.Println("simpleString:",simpleString)
	for i := 0; i < len(simpleString); i++ {
        fmt.Printf("%x ", simpleString[i])
    }
    fmt.Println("");

	s := "ceлёдка" //"селёдка"
	fmt.Println("String s:",s)
	for i:=0; i < len(s); i++ {
		fmt.Printf("%d: %x ",i,s[i])
	}
	fmt.Println("");
	nonAscii := "л"
	fmt.Println("Len of nonAscii:",len(nonAscii))

	//Rune is a Unicode character
	runeArray := []rune(s)
	fmt.Println("RuneArray:");
	for i:=0; i < len(runeArray); i++ {
		fmt.Printf("%d: %c ",i,runeArray[i])
	}
	fmt.Println("");

	//Extracting runes from String using range function
	const nihongo = "日本語"
    for index, runeValue := range nihongo {
        fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
    }

    for index,runeChar := range s {
		fmt.Printf("Position:%d  Character: %c\n",index,runeChar)
	}

	//Strings are immutable
	var immutableStr string = "STRING IS IMMUTABLE in GoLang"
	fmt.Printf("Before mutation:%s\n", immutableStr)
	//cannot assign to immutableStr[3] (neither addressable nor a map index expression)
	/*
	immutableStr[3] = "c"
	fmt.Printf("After mutation:%s", immutableStr)
	*/
	

}

func learnSlices() {
	sli := []int {1, 2, 3}
	for _,v := range sli {
		fmt.Printf("Value in Slice:%d\n",v)
	}
	fmt.Printf("Type of sli:%T Sizeof sli:%d\n",sli,len(sli))
	sli = append(sli,100)
	fmt.Printf("After append, Sizeof sli:%d\n",len(sli))
	arr := [3]int {1,2,3}
	fmt.Printf("Type of arr:%T\n",arr)
}

func learnMaps() {
	//map is the implementation of HashMap in GoLang
	//Exception 1: key is on the right of map, instead of being on the left, 
	//             unlike in Array
	var idMap map[string]int
	idMap = make(map[string]int) // This is initialised as empty Map
	fmt.Println("Lenght of idMap:",len(idMap))

	//Add a new key to the map
	idMap["jane"] = 456
	fmt.Println("After Addition, lenght of idMap:",len(idMap))
	fmt.Println("Value of 'jane'->", idMap["jane"])

	//Remove a key value pair from the map
	delete(idMap, "jane")
	fmt.Println("After Removal, lenght of idMap:",len(idMap))
	fmt.Println("Value of 'jane' after delete->", idMap["jane"])
	
	//Using Map Literal
	mapLiteral := map[string]int {
		"joe": 123}
	fmt.Printf("Value of Key:'joe' in mapLiteral->%d\n",mapLiteral["joe"])

	//Checking the presence of a key using 2 value assignment
	id,present := mapLiteral["joe"]
	fmt.Println("Value of Key:'joe' in mapLiteral->",id, "Is it present?:",present)
	_,present2 := mapLiteral["IAMNOTAKEY"]
	fmt.Println("Key:'IAMNOTAKEY' in mapLiteral:Is it present?:",present2)

	//Iteration through a Map
	fmt.Println("Iteration through a Map:")
	for key,val := range mapLiteral {
		fmt.Println(key,val)
	}
	//Only Values to be accessed?
	fmt.Println("Values of a Map:")
	for _,val := range mapLiteral {
		fmt.Println(val)
	}
	//Only Keys to be accessed?
	fmt.Println("Keys of a Map:")
	for key,_ := range mapLiteral {
		fmt.Println(key)
	}


}


type Person struct  {
	name string
	addr string
	phone string 
}
func learnStruct() {
	// Empty structure
	var p1 Person 
	_ = p1

	// Creation using new function. Empty structure
	p2 := new(Person)  
	_ = p2

	//struct literal
	p3 := Person{name: "joe", addr: "a st.", 
	phone: "123" }
	fmt.Println("p3:", p3)
}

func callFunc2(x int, strPar string) (string, int, bool) {
	return strPar+" STR3", x+1, true
}