package classifier

import (
	"bytes"
	"encoding/csv"

	"github.com/sjwhitworth/golearn/base"
)

type MenuItem struct {
	Name     string
	Category string
}

type MenuItemList []MenuItem

// ToDense to desnse
func (list MenuItemList) ToDense() (instances *base.DenseInstances, err error) {
	var buf = bytes.Buffer{}
	var writer = csv.NewWriter(&buf)

	var headers = []string{"Item Name", "Category"}
	writer.Write(headers)

	for _, item := range list {
		writer.Write([]string{item.Name, item.Category})
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return nil, err
	}

	var reader = bytes.NewReader(buf.Bytes())

	return base.ParseCSVToInstancesFromReader(reader, true)

}
