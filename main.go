package main
import "fmt"
import "unicode/utf8"

func main() {
	var intNum int = 32;
	var floatNum float64 = 1.23;
	var myString string = "Working..."
	var myRune rune = 'a'; // single quote gives you a rune datatype
	var myBoolean bool = true

	// inferred data type
	var myVar = "text"; // automatically a string

	newVar := "testing"; // shorthand

	var var1, var2 = 1,2 // init multiple variables

	fmt.Println("Hello World", intNum)
	fmt.Println("Hello World", floatNum)
	fmt.Println("Hello World", myString)
	fmt.Println("Hello World", myRune)
	fmt.Println("Hello World", myBoolean)
	fmt.Println("Hello World", myVar)
	fmt.Println("Hello World", newVar)
	fmt.Println("Hello World", var1)
	fmt.Println("Hello World", var2)
	fmt.Println("Hello World", len("Y")) // length in bytes
	fmt.Println("Hello World", utf8.RuneCountInString("Y")) // length
}