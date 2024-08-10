package main

import (
	"fmt"

	"github.com/DZaychuk/home/hw02_fix_app/printer"
	"github.com/DZaychuk/home/hw02_fix_app/reader"
	"github.com/DZaychuk/home/hw02_fix_app/types"
)

func main() {
	var path string = "data.json"

	fmt.Printf("Enter data file path: ")
	_, err := fmt.Scanln(&path)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if len(path) == 0 {
		path = "data.json"
	}

	var staff []types.Employee

	staff, err = reader.ReadJSON(path, -1)
	if err != nil {
		fmt.Println("error reading input:", err)
		return
	}

	printer.PrintStaff(staff)
}
