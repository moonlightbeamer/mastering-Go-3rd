package main
import (
	"fmt"
  "encoding/json"
  "bytes"
)
type m struct {
  Name string `json:"name"`
  Tel  int    `json:"tel"`
}
func main() {
  buf := new(bytes.Buffer)
  encoder := json.NewEncoder(buf)
  encoder.SetIndent("", "\t")
  encoder.Encode(m{"asda", 121})
  fmt.Println(buf)
}