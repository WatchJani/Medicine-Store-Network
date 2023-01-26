package os

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadFile(structure interface{}, path string) {
	file, err := os.Open(path)
	ErrorHandler(err)

	data, err := ioutil.ReadAll(file)
	ErrorHandler(err)

	err = json.Unmarshal(data, &structure)
}