package reader

import (
	"encoding/json"
	"fmt"
	"go/types"
	"io"
	"os"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)

	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}
	res := data

	return res, nil
}
