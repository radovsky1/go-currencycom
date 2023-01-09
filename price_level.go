package go_currencycom

type PriceLevel struct {
	Price    float64
	Quantity float64
}

type PriceLevelList []PriceLevel

func (p PriceLevelList) Len() int {
	return len(p)
}

func (p PriceLevelList) Less(i, j int) bool {
	return p[i].Price < p[j].Price
}

func (p PriceLevelList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
