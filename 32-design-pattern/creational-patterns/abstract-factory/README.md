# Fabrique abstraite

La **Fabrique abstraite** est un patron de conception de création qui permet de créer des familles de produits complètes sans avoir à préciser leurs classes concrètes.

La **fabrique abstraite** définit une interface pour la création de chaque produit, mais délègue la véritable création des produits aux classes concrètes de la fabrique. Chaque type de fabrique correspond à une certaine variété de produits.

Le code client appelle les méthodes de création d’un objet **Fabrique** plutôt que de créer directement les produits à l’aide d’un constructeur (opérateur **new**). Comme chaque fabrique possède sa propre variante de produit, tous ses produits seront compatibles.

Le code client manipule les fabriques et les produits uniquement via leurs interfaces abstraites, ce qui lui permet de travailler avec n’importe quelle variante de produit créée par un objet Fabrique. L'on crée juste une nouvelle classe concrète **Fabrique** et la passe au code client.