# Go-Examples

### Installation et mise à jour du binaire Go sous ubuntu 20.04

- Téléchargement du fichier archive Go pour la version **1.22.7**

```
cd ~
wget https://go.dev/dl/go1.22.7.linux-amd64.tar.gz
```

- décompression du fichier archive Go 

Supprimons toute installation Go précédente en supprimant le dossier **/usr/local/go** (s'il existe), puis extrayons l'archive que nous venons de télécharger dans **/usr/local**, en créant une nouvelle arborescence **Go** dans **/usr/local/go** :

```
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.7.linux-amd64.tar.gz
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

- Installer **Delve** le débogueur standard pour les applications Go

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

--- vérifions la configuration en affichant la version de la commande **dlv**

```
dlv version
```

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

### Référence
 
- [Go Documentation](https://go.dev/doc/) 
- Livre **The Complete Guide to Programming Reliable and efficient Software Using Golang d'Adam Freeman**
- [Mockery](https://vektra.github.io/mockery/latest/)
- [Testify](https://github.com/stretchr/testify)
- [Zerolog](https://github.com/rs/zerolog)