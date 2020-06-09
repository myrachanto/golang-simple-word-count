package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"log"
)

var ()

func main(){
	if len(os.Args) == 1 {
		log.Println("No files to process")
		return
	}
	result := make(map[string]int)
	start := time.Now()
	for _, fn := range os.Args[1:]{
		processFile(result, fn)
	}
	defer fmt.Printf("Process took:%v\n", time.Since(start))
	printResult(result)
}

func processFile(result map[string]int, fn string){
	var w string
	r, err := os.Open(fn)
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Close()
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	for sc.Scan(){
		w = strings.ToLower(sc.Text())
		result[w] = result[w] + 1
	}
}
func printResult(result map[string]int){
	fmt.Printf("%-10s%s\n", "count", "word")
	fmt.Printf("%-10s%s\n", "------", "-----")
	
	for w,c := range result {
		fmt.Printf("%-10v%s\n", c, w)
	}
}