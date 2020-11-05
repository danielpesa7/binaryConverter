package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadText(){
	fmt.Println("Write what you want to convert to binary: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	toConvert := string(scanner.Text())
	sb := []byte(toConvert)
	fmt.Printf("%08b \n", sb)
}

func ReadBinary() {
	fmt.Println("Write what binary do you want to convert to text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	toConvert := scanner.Text()
	toConvertA := strings.Split(toConvert, " ")
	var byteA []byte
	for _, v := range toConvertA{
		convertedS, err:= strconv.ParseInt(v, 2, 8)
		if err != nil{
			fmt.Println(err)
		}else {
			byteA = append(byteA, byte(convertedS))
		}
	}
	s := string(byteA)
	fmt.Println(s)
}
