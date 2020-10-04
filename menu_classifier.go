package classifier

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/ensemble"
)

var instance *Classifier

type Classifier struct {
	classifier classifier
}

type classifier interface {
	Predict(with base.FixedDataGrid) (base.FixedDataGrid, error)
}

func New(datasetFile string) *Classifier {
	trainData, err := base.ParseCSVToInstances(datasetFile, true)
	if err != nil {
		panic(err)
	}

	var rf = ensemble.NewRandomForest(20, 1)
	err = rf.Fit(trainData)
	if err != nil {
		panic(err)
	}

	return &Classifier{
		classifier: rf,
	}
}

func (c *Classifier) Predict(items []MenuItem) ([]MenuItem, error) {
	var itemList MenuItemList = items

	input, err := itemList.ToDense()
	if err != nil {
		return nil, err
	}

	resp, err := c.classifier.Predict(input)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	return itemList, nil
}
