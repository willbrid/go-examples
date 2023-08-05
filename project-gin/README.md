# Installation du framework Gin

Initialisation d'un projet **hello-world**

```
mkdir ~/hello-world
cd ~/hello-world
go mod init hello-world
```

Pour installer le package Gin pour notre projet **hello-world**, nous exécutons la commande :

```
go get -u github.com/gin-gonic/gin
```

Pour importer Gin dans notre code :

```
import "github.com/gin-gonic/gin"
```

(Facultatif) nous pouvons importer **net/http**. Ceci est requis par exemple si nous utilisons des constantes telles que **http.StatusOK**

```
import "net/http"
```