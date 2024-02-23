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

	var err error
	var staff []types.Employee

	if len(path) == " " {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path, -1)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
