# Go-Examples

### Installation et mise à jour du binaire Go sous ubuntu 20.04

- Téléchargement du fichier archive Go pour la version **1.21.7**

```
cd ~
wget https://go.dev/dl/go1.21.7.linux-amd64.tar.gz
```

- décompression du fichier archive Go 

Supprimons toute installation Go précédente en supprimant le dossier **/usr/local/go** (s'il existe), puis extrayons l'archive que nous venons de télécharger dans **/usr/local**, en créant une nouvelle arborescence **Go** dans **/usr/local/go** :

```
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.7.linux-amd64.tar.gz
```

(Nous devons exécuter la commande en tant que root ou via sudo).
<br>
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

La commande **go vet** identifie les déclarations susceptibles d'être des erreurs.

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


- Analyser les dépendances et supprimer les dépendances inutilisées dans un module

```
go mod tidy
```

### Référence
 
- [Go Documentation](https://go.dev/doc/) 
- Livre **The Complete Guide to Programming Reliable and efficient Software Using Golang d'Adam Freeman**