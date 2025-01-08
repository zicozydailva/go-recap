package main
import "fmt";

func main() {
	printInfo := "Working"
	printMe(printInfo);

	var numerator int = 11;
	var demonimator int = 2;
	var result int = intDivision(numerator, demonimator);
	fmt.Println(result);
}

func printMe(printVal string) {
	fmt.Println(printVal)
}

func intDivision(numerator int, demonimator int) int {
	var result int = numerator / demonimator;

	return result;
}