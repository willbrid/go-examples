# Fabrique

La **Fabrique** est un patron de conception de création qui permet de créer des produits sans avoir à préciser leurs classes concrètes.

La **fabrique** définit une méthode qui doit être utilisée pour créer des objets à la place de l’appel au constructeur (opérateur **new**). Les sous-classes peuvent redéfinir cette méthode pour modifier la classe des objets qui seront créés.

**Exemple :** Prenons le cas d'un projet de e-commerce avec une fonctionnalité permettant aux clients d'acheter des produits de sport dont un produit peut être : soit une paire de chaussures ou soit un maillot.