# Les tests en Go

Go recherche le code de test dans les fichiers dont le nom se termine par **_test.go**. Aucun autre fichier ne sera pris en compte lors de l'exécution des tests. Il est pratique, mais pas obligatoire, de placer les fichiers sources contenant le code de test dans le même dossier que le package qu'ils testent.

Exemple :

```
service/
   service.go
   service_test.go
```

L’avantage d’utiliser le package de test de Go, par opposition à la création de notre propre package, est que le code de test est ignoré lorsque nous construisons le système pour la publication. Il n’y a donc pas lieu de s’inquiéter de gonfler le binaire exécutable avec du code de test qui ne profite pas aux utilisateurs.

```
// service_test.go
package service_test

import "testing"
...
```

Les tests Go sont des fonctions ordinaires, définies avec le mot-clé **func**, mais ce qui les rend spéciaux, c'est que leurs noms commencent par le mot **Test**.

```
func TestNewReturnsValidService(t *testing.T)
```

Le corps de la fonction de test peut être complètement vide et si nous ne disons pas explicitement à un test d’échouer, il sera considéré comme réussi.

Les fonctions de test prennent ce paramètre **t \*testing.T** et si nous essayons d'écrire un test sans lui, nous obtenons une erreur.

Pour exécuter un ensemble de tests :

```
go test
```

Méthodologie de test : **red, green, refactor**

#### Référence -> LIVRE : The power of GO - Tests [bitfieldconsulting](https://bitfieldconsulting.com/)