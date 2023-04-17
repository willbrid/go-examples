#### Utilisation des packages externes

Les projets peuvent être étendus à l'aide de packages développés par des tiers. Les packages sont téléchargés et installés à l'aide de la commande **go get**.

```
go get github.com/fatih/color@v1.10.0
```

L'argument de la commande **go get** est le chemin du module qui contient le package que nous souhaitons utiliser. <br>
Le nom est suivi du caractère **@** puis du numéro de version du package, qui est précédé de la lettre **v**.
<br>
La commande **go get** est sophistiquée et sait que le chemin spécifié est une URL GitHub. La version spécifiée du module est téléchargée <br>
et les packages qu'elle contient sont compilés et installés afin de pouvoir être utilisés dans le projet. <br>
(Les packages sont distribués sous forme de code source, ce qui permet de les compiler pour la plate-forme sur laquelle nous travaillons.)

<br>

Lien pour trouver les packages

- https://pkg.go.dev
- https://github.com/golang/go/wiki/Projects

L'instruction **require** note la dépendance au module **github.com/fatih/color** et aux autres modules dont il a besoin. Le commentaire **indirect** à la fin des instructions est ajouté automatiquement car les packages ne sont pas utilisés par le code dans le projet. Un fichier nommé **go.sum** est créé lors de l'obtention du module et contient les sommes de contrôle utilisées pour valider les packages.

<br>

Nous pouvons constater que notre projet a des dépendances sur différentes versions d'un module, en particulier dans les projets complexes qui ont beaucoup de dépendances. Dans ces situations, Go résout cette dépendance en utilisant la version la plus récente spécifiée par ces dépendances. Ainsi, par exemple, s'il existe des dépendances sur les versions 1.1 et 1.5 d'un module, Go utilisera la version 1.5 lors de la construction du projet. Go utilisera uniquement la version la plus récente spécifiée par une dépendance, même si une version plus récente est disponible. Si la dépendance la plus récente pour un module spécifie la version 1.5, par exemple, Go n'utilisera pas la version 1.6, même si elle est disponible.
<br>
L'effet de cette approche est que notre projet peut ne pas être compilé à l'aide de la version de module que nous avons sélectionnée avec la commande **go get** si un module dépend d'une version ultérieure. De même, un module peut ne pas être compilé avec les versions qu'il attend pour ses dépendances si un autre module — ou le fichier **go.mod** — spécifie une version plus récente.

<br>

Pour désinstaller un package installé, on peut commencer par enlever toutes les instructions qui utilisent ce package puis exécuter la commande 

```
go mod tidy
```