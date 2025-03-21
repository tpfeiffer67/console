package theme

import (
	"encoding/json"
	"os"
)

func SaveToFile(filename string) {
	file, _ := json.MarshalIndent(CurrentTheme, "", " ")
	_ = os.WriteFile(filename, file, 0644)
}
