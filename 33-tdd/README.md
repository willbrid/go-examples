# Quelques notes

### testify/mock

**testify/mock** est un framework opensource qui se compose d'un package de **mock** et d'un outil de génération de code, **mockery**. <br>
**testify** fournit également un framework d'assertion très puissant et populaire.

Pour utiliser ce framework, nous devons installer ses deux composants principaux en exécutant ces commandes :

```
go get github.com/stretchr/testify

go install github.com/vektra/mockery/v2@latest
```

```
vi $HOME/.bashrc
```

```
...
# Go binary packages
export PATH="$HOME/go/bin:$PATH"
```

Cet outil permet de générer facilement du code **mock** standard, de sorte que nous n'avons pas besoin de le créer et de le maintenir à la main. La génération de code **mock** dans **testify** ne nécessite aucune annotation particulière. Des codes **mock** peuvent être générés pour les interfaces et les fonctions, ce qui les rend adaptées à la fois à la substitution de fonctions et à la substitution d'interfaces.

```
mockery --dir "calculator-project/input" --output "calculator-project/mocks" --all
```

- L'indicateur de chaîne **--dir** spécifie le répertoire dans lequel rechercher les interfaces à générer du code mock.
- L'indicateur **--all** spécifie de rechercher dans tous les sous-répertoires et de générer des codes mock.
- L'indicateur de chaîne **--name** spécifie le nom ou l'expression régulière à faire correspondre lors de la recherche d'interfaces pour générer des codes mock.
- L'indicateur de chaîne **--output** spécifie le répertoire dans lequel placer les code mock générés. Par défaut, il est configuré sur **/mocks**.