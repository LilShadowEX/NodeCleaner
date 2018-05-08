package main

import (
	"fmt"
	"bufio"
	"encoding/csv"
	"os"
	"io"
)

func main() {
	f, _ := os.Open("c:\\test\\test.csv")

	r := csv.NewReader(bufio.NewReader(f)) 
	for {
		record, err := r.Read()

		if err == io.EOF {
			// valdria la pena un Log
			break
		}
		fmt.Println(record)
		fmt.Println(len(record))
		for value := range record	{
			fmt.Printf(" %v\n", record[value])
		}
	}
}