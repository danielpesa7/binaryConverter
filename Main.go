package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main(){
	fmt.Println("Welcome to the best binary converter in the World!")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Type String-Binary to convert text to binary code.")
	fmt.Println("Type Binary-String to convert binary code to text.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	option := scanner.Text()
	switch option {
	case "String-Binary":
		ReadText()
	case "Binary-String":
		ReadBinary()
	default:
		fmt.Printf("%s no es una opción valida\n", option)
	}


}