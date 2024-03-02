package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Employee struct {
	UserID       int    `json:"user_id"`
	Age          int    `json:"age"`
	Name         string `json:"name"`
	DepartmentID int    `json:"department_id"`
}

func ReadJSON(filePath string) ([]Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)

	}

	var data []Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}
	res := data

	return res, nil
}
