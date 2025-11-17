# Go-Examples

### Installation et mise à jour du binaire Go sous ubuntu 24.04

- Téléchargement du fichier archive Go pour la version **1.24.10**

```
cd $HOME && wget https://go.dev/dl/go1.24.10.linux-amd64.tar.gz
```

- décompression du fichier archive Go 

Supprimons toute installation Go précédente en supprimant le dossier **/usr/local/go** (s'il existe), puis extrayons l'archive que nous venons de télécharger dans **/usr/local**, en créant une nouvelle arborescence **Go** dans **/usr/local/go** :

```
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.24.10.linux-amd64.tar.gz
```

Ne décompressons pas l'archive dans une arborescence **/usr/local/go** existante. Ceci est connu pour produire des installations Go cassées.

- Ajoutons **/usr/local/go/bin** à la variable d'environnement **PATH**.

Nous pouvons le faire en ajoutant la ligne suivante à notre **$HOME/.profile** ou **/etc/profile** (pour une installation à l'échelle du système) :

```
export PATH=$PATH:/usr/local/go/bin
```

**Remarque** : Les modifications apportées à un fichier de profil peuvent ne s'appliquer qu'à la prochaine connexion à notre ordinateur. Pour appliquer les modifications immédiatement, exécutons simplement les commandes shell directement ou exécutons-les à partir du profil à l'aide d'une commande telle que 

```
source $HOME/.profile
```

- Vérifions que nous avons installé Go en ouvrant une invite de commande et en saisissant la commande suivante

```
go version
```

Nous confirmons que la commande affiche la version installée de Go.

### Utilisation des outils Go

```
mkdir ~/tools && cd ~/tools
```

```
vi main.go
```

```
package main

import "fmt"

func main() {
	PrintHello()

	for i := 0; i < 5; i++ {
		PrintNumber(i)
	}
}

func PrintHello() {
	fmt.Println("Hello, Willbrid")
}

func PrintNumber(number int) {
	fmt.Println(number)
}
```

- Compiler le code source Go et produire un exécutable (**main**).

```
go build main.go
```

```
./main
```

- Supprimer la sortie du processus de compilation

```
go clean main.go
```

La commande "**go clean**" supprime la sortie produite par la commande "**go build**", y compris l'exécutable et tous les fichiers temporaires créés pendant la construction.

- Générer la documentation 

```
go doc
```

La commande "**go doc**" génère de la documentation à partir du code source.

- Compiler et exécuter en une seule étape

```
go run main.go
```

Sans la création d'un module Go, nous sommes obligés de préciser le nom du fichier de code dont nous souhaiterons compiler et exécuter.

- Définir et exécuter un module Go

```
mkdir tools && cd tools
```

```
go mod init tools
```

```
go run .
```

Avec la création d'un module Go, au lieu de spécifier un fichier de code particulier, le projet peut être compilé et exécuté à l'aide d'un point, indiquant le projet dans le répertoire courant.

- Installer **Delve** le débogueur standard pour les applications Go

Le **débogueur standard** pour les applications Go s'appelle **Delve**. Il s'agit d'un outil tiers, mais il est bien pris en charge et recommandé par l'équipe de développement de Go.

```
go install github.com/go-delve/delve/cmd/dlv@latest
```

La commande "**go install**" télécharge des packages et est généralement utilisée pour installer des packages d'outils.

--- ajoutons **$HOME/go/bin** à la variable d'environnement **PATH**.

```
vi $HOME/.profile
```

```
export PATH=$PATH:$HOME/go/bin
```

```
source $HOME/.profile
```

--- Vérifions la configuration en affichant la version de la commande **dlv**

```
dlv version
```

--- Supposons un code Go dans le fichier **main.go**

```
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go")
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

On peut démarrer le debogeur sur ce fichier **main.go**

```
dlv debug main.go
```

On peut créer un point d'arrêt, ce qui se fait en spécifiant un emplacement dans le code

```
break bp1 main.main:3
```

Le nom du point d'arrêt est **bp1** et l'emplacement spécifie la troisième ligne de la fonction **main** du package **main**.

On peut créer une condition pour le point d’arrêt afin que l’exécution soit interrompue uniquement lorsqu’une expression spécifiée est évaluée à vrai.

```
condition bp1 i == 2
```

Les arguments de la commande **condition** spécifient un point d'arrêt et une expression. Cette commande indique au débogueur que le point d'arrêt nommé **bp1** doit interrompre l'exécution uniquement lorsque l'expression **i == 2** est vraie. Pour démarrer l'exécution, l'on entre la commande

```
continue
```

On peut afficher la valeur courante de la variable **i**

```
print i
```

Le débogueur fournit un ensemble complet de commandes pour inspecter et modifier l'état de l'application dont ci-dessous quelques unes :

----- **print <expr>** : Cette commande évalue une expression et affiche le résultat. Elle peut être utilisée pour afficher une valeur (**print i**) ou effectuer un test plus complexe (**print i > 0**) <br>
----- **set <variable> = <value>** : Cette commande modifie la valeur de la variable spécifiée <br>
----- **locals** : Cette commande imprime la valeur de toutes les variables locales <br>
----- **whatis <expr>** : Cette commande imprime le type de l'expression spécifiée (**whatis i**)

Quelques commandes de débogage utiles pour contrôler l'exécution :

----- **continue** : cette commande reprend l'exécution de l'application <br>
----- **next** : cette commande passe à l'instruction suivante <br>
----- **step** : cette commande passe à l'instruction en cours <br>
----- **stepout** : cette commande quitte l'instruction en cours <br>
----- **restart** : cette commande redémarre le processus. Utilisez la commande continue pour commencer l'exécution <br>
----- **exit** : cette commande quitte le débogueur

- linter de code Go

Un **linter** est un outil qui vérifie les fichiers de code à l'aide d'un ensemble de règles décrivant les problèmes susceptibles de créer de la confusion, de produire des résultats inattendus ou de réduire la lisibilité du code.

Exemple de package linter pour Go : **revive** [https://github.com/mgechev/revive](https://github.com/mgechev/revive)

```
go install github.com/mgechev/revive@latest
```

Cet utilitaire est installé dans le repertoire **go/bin** dans notre dossier personnel et pour exécuter cette commande, nous saisissons

```
revive
```

Le package **revive** peut être configuré à l'aide de commentaires dans les fichiers de code. Par exemple désactiver une ou plusieurs règles pour des sections de code. Nous pouvons aussi utiliser un fichier de configuration au format TOML par exemple un ficier **revive.toml** .

La liste des règles supportées par le linter **revive** : [https://github.com/mgechev/revive#available-rules](https://github.com/mgechev/revive#available-rules).

Syntaxe
```
revive:<enable|disable>:<rule>
```

Exemples
```
package main

import "fmt"

func main() {
    PrintHello()
    for i := 0; i < 5; i++ {
        PrintNumber(i)
    }
}

// revive:disable:exported

func PrintHello() {
    fmt.Println("Hello, Go")
}

// revive:enable:exported

// PrintNumber writes a number using the fmt.Println function
func PrintNumber(number int) {
    fmt.Println(number)
}
```

--- [https://github.com/mgechev/revive#recommended-configuration](https://github.com/mgechev/revive#recommended-configuration)

Exécuter la commande **revive** avec le fichier de configuration **revive.toml**

```
revive -config revive.toml
```

En Go, les commentaires de fonction doivent contenir une phrase commençant par le nom de la fonction et doivent fournir un aperçu concis de l’objectif de la fonction, comme décrit par [https://golang.org/doc/effective_go.html#commentary](https://golang.org/doc/effective_go.html#commentary).

- Analyser un code Go

La commande "**go vet**" identifie les déclarations susceptibles d'être des erreurs.

```
mkdir testGOVETCmd && cd testGOVETCmd
```

```
go mod init testGOVETCmd
```

```
vi main.go
```

```
package main

import "fmt"

func main() {
    PrintHello()
    for i := 0; i < 5; i++ {
        i = i
        PrintNumber(i)
    }
}

func PrintHello() {
    fmt.Println("Hello, Go")
}

func PrintNumber(number int) {
    fmt.Println(number)
}
```

```
go vet main.go
```

Les avertissements produits par la commande **go vet** précisent l'emplacement dans le code où un problème a été détecté et fournissent une description du problème. La commande **go vet** applique plusieurs analyseurs au code, et nous pouvons voir la liste des analyseurs sur [https://golang.org/cmd/vet](https://golang.org/cmd/vet).

Pour déterminer quel analyseur est responsable d'un avertissement, nous utilisons l'argument **-json** qui génère une sortie au format **JSON** permettant de regrouper les avertissements par analyseur.

```
go vet -json main.go
```

Dans notre cas l'analyseur **assign**. Une fois l'analyseur connu, il peut être activé ou désactivé.

```
go vet -assign
```

```
go vet -assign=false
```

L'analyse avec **Go Vet** est activée par défaut dans l'extension **Go** de vscode. Nous pouvons la désactiver via : **Settings > Extensions > Go > Manage > Settings > Vet on save**.

- Utiliser la commande **go fmt**

La commande **go fmt** reformate automatiquement le code Go selon le style officiel (style spécifié par l'équipe de développement Go) : indentation par tabulations, alignement des commentaires et suppression des points-virgules inutiles. Le format est fixe et non personnalisable.

```
go fmt
```

- Afficher l'aide de la fonctionnalité **build**

```
go help build
```

La commande "**go help**" affiche des informations d'aide pour d'autres fonctionnalités de Go.

- Analyser les dépendances et supprimer les dépendances inutilisées dans un module

```
go mod tidy
```

- Utiliser la commande **go get**

--- ajouter ou mettre à jour une dépendance (par exemple **mux**) dans un module

```
go get github.com/gorilla/mux
```

La commande "**go get**" télécharge et installe des packages externes.

--- supprimer une dépendance dans un module

```
go get github.com/gorilla/mux@none
```

- Mettre à jour toutes les dépendances d'un projet

```
go get -u ./...
```

- Exécuter un test unitaire

```
go test ./...
```

La commande "**go test**" permet d'exécuter des tests (unitaires, fonctionnels,...).

### Go test packages

- **mockery**

Mockery est un projet qui crée des implémentations fictives d'interfaces Golang. Les simulations générées dans ce projet sont basées sur la suite de packages de test github.com/stretchr/testify.

```
go install github.com/vektra/mockery/v2@v2.46.3
```

- **testify**

Ensemble de packages Go Code (golang) qui fournissent de nombreux outils pour vérifier que notre code se comportera comme nous le souhaitons.

```
go get github.com/stretchr/testify
```

Générer des mocks pour nos interfaces en utilisant la commande suivante :

```
$HOME/go/bin/mockery --dir $HOME/go-examples/33-tdd/calculator-project --output $HOME/go-examples/33-tdd/calculator-project/mocks --all
```

La commande **mockery** prend en charge une variété d'indicateurs. Voici quelques-uns des indicateurs courants que nous pourrions utiliser :

--- **--dir** : spécifie le répertoire dans lequel rechercher les interfaces à simuler. <br>
--- **--all** spécifie de rechercher dans tous les sous-répertoires et de générer des mocks. <br>
--- **--name** spécifie le nom ou l'expression régulière à faire correspondre lors de la recherche d'interfaces pour générer des mocks. <br>
--- **--output** spécifie le répertoire dans lequel placer les mocks générés. Par défaut, il est configuré sur **/mocks**.

### Références
 
- [Go Documentation](https://go.dev/doc/) 
- [Mockery](https://vektra.github.io/mockery/latest/)
- [Testify](https://github.com/stretchr/testify)
- [Zerolog](https://github.com/rs/zerolog)
- [Book The Complete Guide to Programming Reliable and efficient Software Using Golang of Adam Freeman]()