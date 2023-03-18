package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Michael",
		Price: 5.00,
		SKU:   "zxc-asd-qwe",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
