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

	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	} else {

	}

	staff, err = reader.ReadJSON(path, -1)
	if err != nil {
		fmt.Print(err)
	}

	printer.PrintStaff(staff)
}
