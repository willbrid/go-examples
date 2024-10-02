# Monteur

Le Monteur est un patron de conception de création qui permet de construire des objets complexes étape par étape.

Le monteur n’est pas comme les autres patrons de création : les produits n’ont pas besoin d’avoir une interface commune. Il est ainsi possible de créer différents produits en utilisant le même procédé de fabrication.

Le monteur est également utilisé lorsqu’un produit est complexe et requiert plusieurs étapes pour être fabriqué. Dans ce cas, il est plus simple de mettre en place plusieurs méthodes de construction qu’un seul constructeur monstrueux. Ce processus de fabrication en plusieurs étapes peut poser problème, car on va potentiellement pouvoir passer un produit incomplet et instable au client. Le monteur pallie à ce problème en gardant le produit privé jusqu’à ce qu’il soit complètement fabriqué.

**Exemple :** Prenons l'exemple de deux types de maisons : un **igloo** et une maison traditionnelle (**NormalHouse**). Ces maisons partagent certaines caractéristiques comme le type de fenêtre, le type de porte, le nombre de sols. Mais elles diffèrent dans la façon dont elles sont construites.