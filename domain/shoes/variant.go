package shoes

import (
	"fmt"
	"strconv"
)

type Variant struct {
	Size  float64
	Price string
}

func NewVariant(
	size string,
	price string,
) (*Variant, error) {
	sizeF, err := strconv.ParseFloat(size, 64)
	if err != nil {
		return nil, err
	}
	return &Variant{
		Size:  sizeF,
		Price: price,
	}, nil
}

func NewVariants(
	sizes []string,
	prices []string,
) ([]*Variant, error) {
	if len(sizes) == 0 {
		return nil, fmt.Errorf("not found sizes")
	}
	if len(sizes) != len(prices) {
		return nil, fmt.Errorf("mismatch: len(sizes)=%d, len(prices)=%d\n", len(sizes), len(prices))
	}

	var variants []*Variant
	for i := 0; i < len(sizes); i++ {
		variant, err2 := NewVariant(
			sizes[i],
			prices[i],
		)
		if err2 != nil {
			return nil, err2
		}
		variants = append(variants, variant)
	}
	return variants, nil
}
