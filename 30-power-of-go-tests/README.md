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

- En appelant **t.Parallel()**, tout test peut se déclarer apte à être exécuté simultanément avec d'autres tests.
- Un bon nom de test est une phrase décrivant le comportement attendu (**Running returns true when service is running**).
- La valeur exacte de la mauvaise réponse d'un test peut être un indice utile pour savoir pourquoi elle est fausse.

```
want := "hello"
got := Greeting()
if want != got {
    t.Errorf("want %q, got %q", want, got)
}
```

- **t.Error** (ou son équivalent de formatage **t.Errorf**) marque le test comme ayant échoué, mais ne l’arrête pas. Il est souvent préférable d'abandonner le test dès que quelque chose a échoué. 

```
f, err := os.Open("testdata/input")
if err != nil {
    t.Fatal(err)
}
```

Utiliser **t.Error** lorsqu’il vaut la peine de poursuivre le test; utiliser **t.Fatal** lorsque ce n’est pas le cas.

- Par commodité, seuls les messages de logs des tests ayant échoué sont imprimés. Nous pouvons donc enregistrer autant de messages que nous le souhaitons avec **t.Log** à chaque test, et nous ne serons pas submergé de messages indésirables lorsqu'un autre test échoue.

```
got := StageOne()
t.Log("StageOne result", got)
got = StageTwo(got)
t.Log("StageTwo result", got)
```

- Pour tester tous les packages à partir du répertoire actuel, on fait :

```
got test ./...
```

- L'indicateur **-v** (pour « **verbose** ») imprime les noms de chaque test au fur et à mesure de son exécution, ses résultats et tous les messages de journal, quel que soit l'état d'échec :

```
go test -v
```

```
go test -v .
```

- Si un package particulier n’a pas changé depuis le dernier test, nous verrons ses résultats mis en cache, marqués en conséquence.
- Il peut arriver que nous souhaitions forcer l'exécution d'un test pour une raison quelconque, même si son résultat est mis en cache. Dans ce cas, nous pouvons utiliser l'indicateur **-count** pour remplacer le cache :

```
go test -count=1 .
```

- Pour exécuter un seul test, utiliser l'indicateur **-run** avec le nom du test

```
go test -run TestRunningIsTrueWhenServiceIsRunning
```

- Nous pouvez également l'utiliser (option **-run**) pour exécuter un groupe de tests associés, dont les noms correspondent à l'expression que nous fournissons. Supposons que nous avons les fonctions de test : **TestDatabaseOperation1**, **TestDatabaseOperation2** et **TestDatabaseOperation3**

```
go test -run TestDatabase
```

- **t.Helper** marque la fonction comme une aide au test, ce qui signifie que tout échec sera signalé sur la ligne appropriée du test, et non dans cette fonction.

```
func TestUserCanLogin(t *testing.T) {
    createTestUser(t, "Jo Schmo", "dummy password")
    ... // check that User can log in
}

func createTestUser(t *testing.T, user, pass string) {
    t.Helper()
    ... // create user ...
    if err != nil {
        t.Fatal(err)
    }
}
```

L’appel de **t.Helper** rend effectivement la fonction d’assistance invisible aux tests en échec. Cela est logique, car nous sommes généralement plus intéressés par le fait de savoir quel test a rencontré un problème.

- Si le test doit créer ou écrire des données dans des fichiers, nous pouvons utiliser **t.TempDir** pour créer un répertoire temporaire pour eux.

```
f, err := os.Create(t.TempDir()+"/result.txt")
```

Ce répertoire est propre au test, donc aucun autre test n'interférera avec lui. Une fois le test terminé, ce répertoire et tout son contenu seront également automatiquement nettoyés

- Lorsque nous devons nettoyer quelque chose nous-mêmes à la fin d'un test, nous pouvons utiliser **t.Cleanup** pour enregistrer une fonction appropriée. La fonction enregistrée par **t.Cleanup** sera appelée une fois le test terminé.

```
func createTestDB(t *testing.T) *DB {
    db := ... // create db
    t.Cleanup(func() {
        db.Close()
    })
    return db
}

res := someResource()
t.Cleanup(func() {
    res.GracefulShutdown()
    res.Close()
})
```

- Une façon utile de réfléchir à la valeur d’un test donné est de se demander « **Quelles implémentations incorrectes réussiraient encore ce test ?** ».
- Dans nos projets, il faut toujours détecter et améliorer les **test faibles**.
- **cmp.Equal** et **cmp.Diff** : deux fonctions de comparaison importantes du package [go-cmp](https://github.com/google/go-cmp/cmp).

```
import "github.com/google/go-cmp/cmp"

func TestNewThing_ReturnsThingWithGivenXYZValues(t *testing.T) {
    t.Parallel()
    x, y, z := 1, 2, 3
    want := &thing.Thing{
        X: x,
        Y: y,
        Z: z,
    }
    got, err := thing.NewThing(x, y, z)
    if err != nil {
        t.Fatal(err)
    }
    if !cmp.Equal(want, got) {
        t.Error(cmp.Diff(want, got))
    }
}
```

```
func NewThing(x, y, z int) (*Thing, error) {
    return &Thing{
        X: x,
        Y: y,
        Z: z,
    }, nil
}
```

<br>

#### Référence -> LIVRE : The power of GO - Tests [bitfieldconsulting](https://bitfieldconsulting.com/)