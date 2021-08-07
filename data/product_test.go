package data

import (
	"testing"
)

func TestProductValidator(t *testing.T) {
	p := &Product{
		Name:        "Test",
		Description: "Test",
		Price:       1,
		SKU:         "abc-abc-def",
	}
	if err := p.Validate(); err != nil {
		t.Fatal(err)
	}
}
