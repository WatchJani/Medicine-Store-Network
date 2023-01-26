package os

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const OUTPUT string = "res"

func WriteFile(data any) {
	jp, err := json.Marshal(data)
	ErrorHandler(err)

	_, err = os.Stat(OUTPUT)

	if os.IsNotExist(err) {
		err = os.Mkdir(OUTPUT, 0755)
		ErrorHandler(err)
		fmt.Println("Directory created!")
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/res.json", OUTPUT), jp, 0644)
	ErrorHandler(err)
}
