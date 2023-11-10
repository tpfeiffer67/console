package theme

import (
	"encoding/json"
	"io/ioutil"
)

func SaveToFile(filename string) {
	file, _ := json.MarshalIndent(CurrentTheme, "", " ")
	_ = ioutil.WriteFile(filename, file, 0644)
}
