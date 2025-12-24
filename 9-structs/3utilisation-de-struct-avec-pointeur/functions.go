package main

/**
Accéder aux champs d’une structure via un pointeur peut s’avérer peu pratique, d’autant plus que les structures sont fréquemment utilisées
comme paramètres et valeurs de retour des fonctions.
L’utilisation de pointeurs reste néanmoins indispensable afin d’éviter des copies inutiles de structures et de garantir que les
modifications effectuées dans une fonction s’appliquent bien aux valeurs passées en argument.
Pour simplifier l’écriture et améliorer la lisibilité du code, Go permet d’accéder directement aux champs d’une structure référencée
par un pointeur, sans nécessiter l’utilisation explicite de l’opérateur déréférencement (*).
Cette fonctionnalité ne modifie pas le type de données du paramètre de la fonction, qui reste *Product, et s'applique uniquement
lors de l'accès aux champs.
**/

func calcTax(product *Product) {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
}

func calcTaxWithResult(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}

	return product
}

/**
Pour garantir la duplication du fournisseur, la fonction `copyProduct` l'assigne à une variable distincte, puis crée un pointeur vers
cette variable. Cette méthode est peu élégante, mais elle permet de forcer la copie de la structure. Cependant, cette technique est
spécifique à un seul type de structure et doit être répétée pour chaque champ de structure imbriqué.
**/

func copyArticle(article *Article) Article {
	a := *article
	s := *article.Supplier
	a.Supplier = &s

	return a
}
