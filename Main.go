package main

import (
	"fmt"
	"bufio"
	"encoding/csv"
	"os"
	"io"
	"strings"
)
type Wordseeker struct {
	word string
	recurrence int64
}

func main() {
	path := os.Args[1]
	f, _ := os.Open(path)

	r := csv.NewReader(bufio.NewReader(f))
	// positive := Wordseeker{}
	// negative := Wordseeker{}
	var wordtotal []string
	var temppos []string
	var tempneg []string
	for {
		record, err := r.Read()

		if err == io.EOF {
			// valdria la pena un Log
			break
		}

		//fmt.Println(record)
		if "n" == record[4]{
			fmt.Println("esto fue un no")
			fmt.Println(record[1])
			wordtotal = strings.Fields(record[1])
			tempneg = append(tempneg,wordtotal[])
			wordtotal = strings.Fields(record[2])
			tempneg = append(tempneg,wordtotal[])
			fmt.Println(tempneg)
			fmt.Println("length : ",len(tempneg))
			}

		if "y" == record[4]{
			fmt.Println("esto fue un si")
			fmt.Println(record[1])
			wordtotal = strings.Fields(record[1])
			fmt.Println(wordtotal)
			fmt.Println("length : ",len(wordtotal))

			//words := strings.fields(record[2])
			//aggregatedwords = append(aggregatedwords,words)
			//words := strings.fields(record[3])
			//aggregatedwords = append(aggregatedwords,words)
		}
	}
	fmt.Println("agregated: ")
	//fmt.Println(aggregatedwords)
}
