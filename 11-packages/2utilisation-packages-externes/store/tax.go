package store

/**
Les packages peuvent contenir plusieurs fichiers de code et, pour simplifier le développement, les règles de contrôle d'accès et
les préfixes de package ne s'appliquent pas lors de l'accès aux fonctionnalités définies dans le même package.

Toutes les fonctionnalités définies dans le fichier `tax.go` ne sont pas exportées, ce qui signifie qu'elles ne peuvent être utilisées que
dans le package `store`. Notons que la méthode calcTax peut accéder au champ `price` du type `Product` sans avoir à faire référence au
type comme `store.Product`, car il se trouve dans le même package.

Chaque fichier de code peut contenir une fonction d'initialisation qui n'est exécutée que lorsque tous les packages ont été chargés et
que toutes les autres initialisations (comme la définition des constantes et des variables) ont été effectuées. L'utilisation la plus
courante des fonctions d'initialisation est d'effectuer des calculs complexes ou nécessitant une duplication.
La fonction d'initialisation est appelée `init`, et elle est définie sans paramètres ni résultat. La fonction `init` est
appelée automatiquement et permet de préparer le paquetage à l'utilisation.
La fonction `init` n'est pas une fonction Go classique et ne peut pas être appelée directement. De plus, contrairement aux fonctions classiques,
un même fichier peut définir plusieurs fonctions `init`, qui seront toutes exécutées.
Chaque fichier de code peut avoir sa propre fonction d'initialisation. Avec le compilateur Go standard, ces fonctions sont exécutées par
ordre alphabétique des noms de fichiers : la fonction du fichier a.go sera exécutée avant celle du fichier b.go, et ainsi de suite.
Cependant, cet ordre ne fait pas partie des spécifications du langage Go et ne doit pas être pris en compte. Nos fonctions d'initialisation
doivent être autonomes et ne pas dépendre de l'appel préalable d'autres fonctions d'initialisation.
**/

const defaultTaxRate float64 = 0.2
const minThreshold float64 = 10

var categoryMaxPrices map[string]float64 = map[string]float64{
	"Watersports": 250,
	"Soccer":      150,
	"Chess":       50,
}

func init() {
	for category, price := range categoryMaxPrices {
		categoryMaxPrices[category] = price + price*defaultTaxRate
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

func (t *TaxRate) calcTax(product *Product) float64 {
	if product.price > t.threshold {
		return product.price + product.price*t.rate
	}

	return product.price
}

func (t *TaxRate) calcTaxWithCategory(product *Product) (price float64) {
	if product.price > t.threshold {
		price = product.price + product.price*t.rate
	} else {
		price = product.price
	}
	if max, ok := categoryMaxPrices[product.Category]; ok && price > max {
		price = max
	}

	return
}
