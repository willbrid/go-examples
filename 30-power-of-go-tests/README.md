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

- Cela n’a pas de sens de comparer les valeurs d’erreur dans Go en utilisant l’opérateur **!=** : elles ne seront jamais égales.

- Il y a quatre points clés concernant les tests des erreurs :

1. Toujours vérifier les erreurs dans les tests, qu'elles soient attendues ou non. <br>
2. Leur nature importe généralement peu, tant qu'elles ne sont pas nulles. <br>
3. Lorsque leur nature est importante, utilisons **errors.Is** pour les comparer à une erreur sentinelle. <br>
4. Lorsque nous devons combiner une erreur sentinelle avec des informations dynamiques, utilisons **fmt.Errorf** et le verbe %w pour créer une erreur encapsulée.

- **Les entrées de tests**

Un bon testeur adopte une approche conflictuelle lors de la conception des entrées de test. Les utilisateurs ne chercheront peut-être pas délibérément à faire planter notre logiciel, mais cela peut parfois sembler le cas. Si nous ne parvenons pas à trouver des entrées improbables et bizarres qui font planter le système, nos utilisateurs le feront certainement.

Lorsque nous concevons des entrées de test, essayons de penser à des choses qui ne peuvent théoriquement pas se produire, ou que seule une personne folle ferait, et testons-les.

Nous pouvons et devons attaquer le système avec tous ces types d'entrées contradictoires, mais nous pouvons faire encore plus. Comme nous connaissons précisément l'implémentation du système, nous avons une bonne idée des bugs qui pourraient s'y cacher.

Chaque fois qu'une instruction if dans le corps du test sélectionne une logique de test différente en fonction d'un élément du cas, nous devrions probablement refactoriser en plusieurs tests, un pour chaque type de cas. <br>
Inversement, nous pouvons dire que chaque fois que nous testons exactement la même chose sur un ensemble d'entrées différentes, nous avons un bon candidat pour un test de table.

- Externalisation des données de test vers des variables ou des fonctions

Lorsque plusieurs tests doivent utiliser les mêmes entrées de test, utilisons des fonctions pour construire en toute sécurité tout type d'entrée de test contenant des map et des slice, y compris les structures dont les champs sont des map ou des slice.

- **Chargement des données de test à partir de fichiers**

Lorsque les données de test peuvent devenir très volumineuses, il est conseillé de les placer dans un fichier pour éviter d'encombrer le code source. Par convention, nous plaçons ces fichiers dans un dossier nommé **testdata**, car les outils Go ignorent tout dossier portant ce nom lors de la recherche de packages.

En particulier, évitons d'ouvrir un fichier uniquement pour créer un **io.Reader**, si c'est ce que la fonction doit utiliser, utilisons plutôt **strings.NewReader** pour créer un reader à partir de données statiques :

```
...
input := strings.NewReader("hello world")
got, err := parse.ParseReader(input)
...
```

De même, ne créons pas de fichier uniquement pour obtenir une valeur **io.Writer**. Si la fonction testée utilise un script d'écriture, et que son contenu est indifférent, nous n'avons même pas besoin d'un **strings.Reader**. Nous pouvons simplement utiliser **io.Discard** prédéclaré, qui est un script d'écriture vers nulle part.

```
...
tps.WriteReportTo(io.Discard)
...
```

D'un autre côté, si le test doit examiner le **writer** par la suite, pour voir ce que la fonction lui a réellement écrit, alors nous pouvons utiliser un tampon :

```
...
buf := new(bytes.Buffer)
tps.WriteReportTo(buf)
want := "PC Load Letter"
got := buf.String()
...
```

Si nous devons créer ou ouvrir un fichier parce que la fonction testée attend un fichier **\*os.File**, nous avons deux cas : <br>
--- Si la fonction souhaite simplement lire des octets, nous pouvons lui passer un **io.Reader**. Si elle doit écrire des octets (par exemple, pour afficher du texte), utilisons un **io.Writer**. C'est pratique pour les utilisateurs, car de nombreux éléments implémentent **io.Reader/Writer**. Et c'est pratique pour les tests : un **\*bytes.Buffer** implémente les deux interfaces. <br>
--- En revanche, si la fonction traite réellement des fichiers, et pas seulement des flux d'octets, nous ne pouvons pas utiliser ces interfaces. Par exemple, si elle doit appeler des méthodes sur son argument spécifiques à **\*os.File**, comme **Name** ou **Stat**, nous devons lui passer un fichier.

- **L'abstraction du système de fichiers**

Le package **fs.FS** en Go fournit une abstraction pour le système de fichiers, permettant de parcourir une arborescence de fichiers sans se soucier de leur emplacement réel (sur disque ou en mémoire). Cette interface est utile pour éviter des accès disque coûteux, notamment lors des tests.

L'implémentation fstest.MapFS propose un système de fichiers en mémoire, ce qui accélère considérablement les tests en éliminant les opérations d'E/S disque. Cela permet aux fonctions manipulant des fichiers de fonctionner de manière identique, qu'elles accèdent à un vrai disque ou à une structure simulée en mémoire.

```
...
fsys := fstest.MapFS{
  "file.go": {},
  "subfolder/subfolder.go": {},
  "subfolder2/another.go":{},
  "subfolder2/file.go":{},
}
results := find.GoFiles(fsys)
...
```

- **Utilisation de t.TempDir pour la sortie de test avec nettoyage**

La fonction **t.TempDir()** en Go permet de créer un répertoire temporaire unique pour chaque test, qui est automatiquement supprimé à la fin du test. Cela évite l'encombrement du système de fichiers et garantit un nettoyage efficace.

Il est recommandé de nommer les fichiers créés avec **t.Name()**, qui retourne le nom du test ou du sous-test en cours. Cela facilite l'identification de l'origine des fichiers en cas d'erreur ou de persistance inattendue.

```
...
image.Generate(t.TempDir()+"/"+t.Name()+".png")
...
```

### Qu'est-ce que la couverture des tests ?

La couverture des tests mesure la proportion du code exécutée lors des tests. Par exemple, une couverture de 80 % signifie que 80 % des instructions du package ont été testées. <br>
Lorsqu'une partie du code n'est pas couverte, il est utile de se demander si cette partie a un impact sur le comportement utilisateur :

- Si oui, il faut écrire des tests pour la couvrir
- Si non, le code peut être supprimé

Dans tous les 2 cas ci-dessus, cela améliore la qualité globale du code.

- Obtenir un simple pourcentage de couverture de tests

```
go test -cover
```

- Génèrer un rapport de couverture du code pendant les tests dans un fichier **coverage.out**

```
go test -coverprofile=coverage.out
```

- Génèrer une visualisation HTML de la couverture de test se trouvant dans le fichier **coverage.out**

```
go tool cover -html=coverage.out
```

Cela ouvrira le navigateur web par défaut et affichera une page HTML : <br>
--- le code couvert est surligné en vert <br>
--- le code non couvert est surligné en rouge <br>
--- certaines lignes, comme les importations, noms de fonctions ou commentaires, ne sont pas pertinentes pour la couverture. Elles sont grisés et peuvent être ignorés.

Tous les changements qui augmentent la couverture des tests ne constituent pas nécessairement une amélioration. En d’autres termes, il est possible d’avoir un test qui exécute un morceau de code, mais qui ne nous dit toujours rien d’utile sur son exactitude.

- Utiliser le **bugging** pour identifier les tests faibles
Un test faible est un test qui ne teste pas autant qu'il le devrait. Une façon de détecter de tels tests pourrait être d'ajouter délibérément des bugs au système et de voir si un test les détecte.

- **Détection de code inutile ou inaccessible**
si un bug délibérément introduit dans le code ne fait échouer aucun test, cela ne signifie pas forcément que les tests sont faibles. Il est possible que le code modifié ne soit pas nécessaire, car son bon ou mauvais fonctionnement n'affecte pas le comportement observable du système. Dans ce cas, ce code est inutile ou inaccessible, et peut être supprimé pour améliorer la qualité du programme. <br>
Le bugging automatisé est appelé **test de mutation**.

- **test de mutation automatisé**
C'est processus qui insère automatiquement des bugs dans le code pour tester l’efficacité de la suite de tests. Il analyse le code, modifie certaines instructions (par exemple, remplacer == par !=), puis exécute les tests : <br>
--- si un test échoue : le code est bien testé. <br>
--- si aucun test n’échoue : cela peut révéler un test manquant ou insuffisant, ou du code inutile.

À la fin, l’outil fournit une liste de modifications qui n’ont provoqué aucun échec, que le développeur peut utiliser pour améliorer les tests ou nettoyer le code. <br>
L'outil **go-mutesting** est un framework de test de mutaiton sur un code source Go, que nous pouvons installer en exécutant la commande :

```
go install github.com/avito-tech/go-mutesting/cmd/go-mutesting@latest
```

Il est recommandé d'effectuer une sauvegarde de l'intégralité de notre projet, y compris le répertoire **.git**, avant d'exécuter **go-mutesting**. <br>
Nous exécutons la commande au niveau du repertoire de notre projet Go.

```
go-mutesting .
```

si l'on obtient un score de **1,0** avec un testeur de mutation, cela signifie que toutes les mutations ont été détectées par les tests, ce qui est idéal. Si l’on veut examiner les mutations produites, il est possible de conserver les fichiers temporaires générés par go-mutesting en utilisant l'option **--do-not-remove-tmp-folder**, afin de consulter les versions modifiées du code car par défaut ils sont supprimés.

L’utilisation d’outils comme **go-mutesting** (ou le nouveau **gremlins**(**https://github.com/go‐gremlins/gremlins**)) peut générer beaucoup de résultats sur de gros projets, dont certains seront pertinents, révélant des bugs réels, des tests faibles ou du code inutile. Il n’est pas nécessaire de faire ces tests souvent, mais plutôt comme un bilan de santé périodique, par exemple tous les deux mois ou après des changements importants. Il est aussi utile de suivre le score de mutation dans le temps : une baisse peut indiquer un affaiblissement des tests. Ces outils permettent d’améliorer la qualité des tests et de mieux cibler les comportements critiques.

### Tester l'intestable

- Il est préférable de lancer un projet rapidement, même incomplet, en mettant en place un squelette fonctionnel minimal.
- Transformer rapidement le squelette en un package testable pour construire progressivement le système. Le squelette doit être couvert par des tests pour en assurer la fiabilité. Ne jamais trop s’éloigner d’un état exécutable. Avancer de façon incrémentale pour garder un système fonctionnel à chaque étape.
- Commencer par un test très simple, même trivial, pour amorcer le développement guidé par les tests. Ne pas chercher à tester immédiatement un comportement complexe. Créer une version très simplifiée du problème pour pouvoir écrire un premier test réalisable.
- Décomposer un problème complexe en petites unités testables. Approche progressive et modulaire pour traiter les grandes fonctionnalités.

Les tests internes, placés dans des fichiers **_internal_test.go**, permettent de vérifier le bon fonctionnement des composants internes, mais ils sont plus fragiles car ils dépendent des détails d'implémentation. Une modification interne peut les casser même si le comportement global reste inchangé.

- **Concurrence**

Pour tester une fonction concurrente comme une tentative de connexion avec timeout, on utilise un select avec un canal de temporisation (timeout.C). Si le timeout est atteint (réception sur le canal), le test échoue. Sinon, on remet le timer en veille et on réessaie. Ce mécanisme permet de contrôler le comportement des goroutines dans le temps et de tester la robustesse de la logique concurrente.

Une approche de gérer aussi une fonction concurrente consiste à laisser la concurrence à la fonction appelante.

- **Sécurité de la concurrence**

Pour tester la sécurité concurrente d’un code, on peut utiliser un test de détection de fuite qui exécute plusieurs goroutines simultanément, dans l’espoir de provoquer une course de données. Toutefois, ce type de test reste incertain car l’ordre d’exécution est imprévisible. Pour fiabiliser cela, Go fournit le détecteur de course (**-race**), un outil puissant capable d’identifier de nombreuses conditions de course durant l’exécution des tests.

```
go test -race
```

<br>

### Référence 

- LIVRE : The power of GO - Tests [bitfieldconsulting](https://bitfieldconsulting.com/)
- Practical Go : [Practical Go](https://dave.cheney.net/practical-go/presentations/qcon-china.html)