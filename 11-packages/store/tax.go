package store

const defaultTaxRate float64 = 0.2
const minThreshold = 10

type TaxRate struct {
	rate, threshold float64
}

func newTaxRate(rate, threshold float64) *TaxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}
	if threshold < minThreshold {
		threshold = minThreshold
	}

	return &TaxRate{rate, threshold}
}

func (taxRate *TaxRate) calcTax(product *Product) float64 {
	if product.price > taxRate.threshold {
		return product.price + (product.price * taxRate.rate)
	}

	return product.price
}
