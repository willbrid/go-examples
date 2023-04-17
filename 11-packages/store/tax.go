package store

const defaultTaxRate float64 = 0.2
const minThreshold = 10

var categoryMaxPrices map[string]float64 = map[string]float64{
	"Watersports": 250,
	"Soccer":      150,
	"Chess":       50,
}

/**
Chaque fichier de code peut contenir une fonction d'initialisation qui n'est exécutée que lorsque tous
les packages ont été chargés et toutes les autres initialisations.
La fonction d'initialisation s'appelle init et est définie sans paramètres ni résultat.
La fonction init est appelée automatiquement et offre la possibilité de préparer le package à utiliser.
La fonction init n'est pas une fonction Go standard et ne peut pas être appelée directement. Et, contrairement
aux fonctions régulières, un seul fichier peut définir plusieurs fonctions init, qui seront toutes exécutées.
**/
func init() {
	for category, price := range categoryMaxPrices {
		categoryMaxPrices[category] = price + (price + defaultTaxRate)
	}
}

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
	var price float64

	if product.price > taxRate.threshold {
		price = product.price + (product.price * taxRate.rate)
	} else {
		price = product.price
	}

	if max, ok := categoryMaxPrices[product.Category]; ok && price > max {
		price = max
	}

	return price
}
