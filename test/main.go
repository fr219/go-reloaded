package main

import (
	"fmt"
	"os"
	"piscine"
)

func main() {
	// where the user will enter 2 files as a string 

	// checking filesArgs (handling filesArgs)
	// - where user might enter wrong format or more than one file 
	// check the files that the user has entered 
	if len(os.Args) < 3 {
		fmt.Print("Error: (Usage) <inputfile> <outputfile>")
		return
	}

	// read file names  
	inputFile := os.Args[1]
	outPutFile := os.Args[2]

	// open and read the input file 
	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(outPutFile)
	if err != nil {
		fmt.Println("Error creating ")
	}
	defer outFile.Close()

	//	Proccess the input file 
	data := make([]byte, 1024) //  It will temporarily hold data read from the input file.
	modifiedLine := ""
	for {
		n, err := inFile.Read(data)
        if n > 0 {
            line := string(data[:n])
            modifiedLine = piscine.Modifiy(line)
            _, err := outFile.WriteString(modifiedLine)
            if err != nil {
                fmt.Println("Error writing to output file:", err)
                return
            }
		}
		if err != nil {
            if err.Error() != "EOF" {
                fmt.Println("Error reading input file:", err)
            }
            break
        }
	}
	fmt.Println(modifiedLine)


}
