package main

import (
_	"fmt"
_	"io/ioutil"
	"os"
	"log"
	"bufio"
)

func main() {

	var anchorLeft = "<a href=\""
	var anchorRight = "</a>"

	/*
	input, err := ioutil.ReadFile("clipboard.txt")
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
*/


	file, err := os.Open("clipboard.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	f, _ := os.OpenFile("clipboard.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	  defer f.Close()
	  dw := bufio.NewWriter(f)
	  
	  

	scanner := bufio.NewScanner(file)
        for scanner.Scan() {
		if scanner.Text() != "" {
	            var link = anchorLeft + scanner.Text()+"\">"+scanner.Text() + anchorRight
		    dw.WriteString(link + "\n")

		}
	}
	dw.WriteString("<link rel=\"stylesheet\" href=\"clipboard.css\"/>")
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	  dw.Flush()
}

