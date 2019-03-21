package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//fileContents:=[]byte
	fileName := os.Args[1] // os.Args is slice of string that returns slice of string with temp file of the running this program as well as argumenst passed into program when it excute
	//with no argument passed into the program it will give the error
	fp, err := os.Open(fileName)
	if err != nil { //if there is error
		fmt.Println("Error!", err) //print out the error statement
		os.Exit(1)                 //exit the program
	}
	io.Copy(os.Stdout, fp) //copy the content of the file(fp) to the terminal(os.Stdout)

}
