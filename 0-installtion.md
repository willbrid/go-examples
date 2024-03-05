# Installation et mise à jour du binaire Go sous ubuntu 20.04

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