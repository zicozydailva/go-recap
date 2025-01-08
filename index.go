package main
import "fmt";

func main() {
	printInfo := "Working"
	printMe(printInfo);

	var numerator int = 11;
	var demonimator int = 2;
	var result, remainder int = intDivision(numerator, demonimator);
	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
}

func printMe(printVal string) {
	fmt.Println(printVal)
}

func intDivision(numerator int, demonimator int) (int, int) {
	var result int = numerator / demonimator;
	var remainder int = numerator % demonimator;

	return result, remainder;
}