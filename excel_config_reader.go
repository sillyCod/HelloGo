package main

import (
    "fmt"
    "excelize"
)

func init()

func main() {
	data := readExcel("so")
	fmt.Printf("%T\n", data)
	for _, slice := range data {
		for _, v := range slice {
			fmt.Println(v)
		}
	}
}


func readExcelData(path string) interface{} {
	return
}


func readExcel(file string) [][]string {
	return [][]string{{"hello"}, {"world"}}
}