package writer

import (
	"encoding/json"
	"os"
)

func jsonout() {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})
}
