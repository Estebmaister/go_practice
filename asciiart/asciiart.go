package asciiart

import (
	"fmt"
	t "time"
)

// ASCIIArt makes and illustration of a dog and a goopher and print it in the console
func ASCIIArt() {
	fmt.Println(t.Now())
	printDoggy()
	fmt.Println()
	printGopher()
}

func printDoggy() {
	fmt.Println("  __      _")
	fmt.Println("o'')}____//")
	fmt.Println(" `_/      )")
	fmt.Println(" (_(_/-(_/ ")
}

func printGopher() {
	fmt.Println("    `.-::::::-.`    ")
	fmt.Println(".:-::::::::::::::-:.")
	fmt.Println("`_:::    ::    :::_`")
	fmt.Println(" .:( ^   :: ^   ):. ")
	fmt.Println(" `:::   (..)   :::. ")
	fmt.Println(" `:::::::UU:::::::` ")
	fmt.Println(" .::::::::::::::::. ")
	fmt.Println(" O::::::::::::::::O ")
	fmt.Println(" -::::::::::::::::- ")
	fmt.Println(" `::::::::::::::::` ")
	fmt.Println("  .::::::::::::::.  ")
	fmt.Println("    oO:::::::Oo     ")
}