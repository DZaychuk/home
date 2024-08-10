package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/DZaychuk/home/hw02_fix_app/types"
)

type Employee struct {
	UserID       int    `json:"user_id"`
	Age          int    `json:"age"`
	Name         string `json:"name"`
	DepartmentID int    `json:"department_id"`
}

func ReadJSON(filePath string, limit int) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return []types.Employee{}, fmt.Errorf("error: %v", err)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		return []types.Employee{}, fmt.Errorf("error: %v", err)

	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return []types.Employee{}, fmt.Errorf("error: %v", err)
	}
	res := data

	return res, nil
}
