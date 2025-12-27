package store

type Describable interface {
	GetName() string
	GetCategory() string
}

/**
Go permet de composer des interfaces à partir d'autres interfaces.
Une interface peut en englober une autre, de sorte que les types doivent implémenter toutes les méthodes définies par les
interfaces englobante et sous-jacente. Les interfaces sont plus simples que les structures et ne comportent ni champs ni méthodes à promouvoir.
La composition d'interfaces produit une union des méthodes définies par les types englobant et sous-jacent. Dans cet exemple,
l'implémentation de l'interface `DescribableItem` requiert les méthodes `GetName`, `GetCategory` et `Price`.
Les méthodes `GetName` et `GetCategory`, définies directement par l'interface `DescribableItem`, sont combinées avec la méthode `Price`
définie par l'interface `ItemForSale`.
**/

type DescribableItem interface {
	GetName() string
	GetCategory() string
	ItemForSale
}
