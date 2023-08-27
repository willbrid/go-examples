# Utilisation du framework platform dans le module sportsstore

Au préalable, les sous-dossiers platform et sportsstore doivent se trouver dans le même répertoire.
<br>
Pour utiliser le framework platform comme dépendance dans le module sportsstore, nous faisons :

```
mkdir sportsstore
cd sportsstore
go mod init sportsstore
go mod edit -require="platform@v1.0.0"
go mod edit -replace="platform@v1.0.0"="../platform"
go get -d "platform@v1.0.0"
```

La directive **require** déclare une dépendance sur le module **platform**. Dans les projets réels, cela peut être spécifié comme URL de notre référentiel de contrôle de version, comme une URL GitHub. <br> 
La directive **replace** fournit un chemin local où le module **platform** peut être trouvé. Lorsque les outils Go résolvent une dépendance sur un package dans le module **platform**, ils le feront en utilisant le dossier platform, qui se trouve au même niveau que le dossier **sportsstore**.
Le projet **platform** présente des dépendances vis-à-vis de packages tiers, qui doivent être résolues avant de pouvoir être utilisées. Cela a été fait par la commande **go get**, qui a produit la directive **require**, qui déclare les dépendances indirectes sur les packages utilisés pour implémenter les sessions utilisés par le projet **platform**.
