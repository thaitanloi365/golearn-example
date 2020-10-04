package classifier

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMenuClassifier(t *testing.T) {
	var sampleData = []MenuItem{
		{
			Name: "L-Cup Soup",
		},
		{
			Name: "L-Bowl Soup",
		},
		{
			Name: "L-Soup/Salad",
		},
		{
			Name: "Jack Daniels Rye",
		},
		{
			Name: "SazRac Rye",
		},
		{
			Name: "RT Bourbon",
		},
		{
			Name: "Corsair Sm Bourb",
		},
	}

	var classifier = New("./menu.csv")
	result, err := classifier.Predict(sampleData)

	assert.NoError(t, err)

	fmt.Println(result)
}
