package main
import "fmt"
import "unicode/utf8"

func main() {
	var intNum int = 32;
	var floatNum float64 = 1.23;
	var myString string = "Working..."



	fmt.Println("Hello World", intNum)
	fmt.Println("Hello World", floatNum)
	fmt.Println("Hello World", myString)
	fmt.Println("Hello World", len("Y")) // length in bytes
	fmt.Println("Hello World", utf8.RuneCountInString("Y")) // length
}