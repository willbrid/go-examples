# Langage testscript

Les scripts de test utilisés avec testscript ne sont ni des scripts shell ni du code Go, mais sont écrits dans un langage spécifique (DSL) conçu pour tester des programmes en ligne de commande. Ce DSL n’est pas un langage généraliste, mais un outil simple et élégant pour écrire des tests automatisés.

### Négation d'une assertion avec le préfixe !

Nous pouvons également affirmer que la sortie du programme ne doit pas correspondre, en préfixant la ligne stdout par le caractère **!**.

```
exec echo 'hello world'
! stdout 'hello world\n'
```

Nous pouvons également utiliser l'opérateur ! pour nier d'autres assertions, par exemple **exec**.
Un **exec** seul affirme que le programme donné réussit, c'est-à-dire qu'il renvoie un état de sortie **nul** (**zéro**). L'effet de l'opérateur **! exec** est d'affirmer que le programme échoue, c'est-à-dire que son état de sortie n'est pas **nul** (**différent de zéro**).

Pourquoi vouloir affirmer que l'exécution d'une commande donnée échoue ? C'est le comportement attendu d'un outil en ligne de commande lorsque l'utilisateur spécifie un indicateur ou un argument non valide, par exemple. Dans ce cas, il est courant d'afficher un message d'erreur et de quitter l'application avec un code de sortie différent de zéro. L'assertion correspondante pour l'erreur standard correspondante est nommée **stderr**.

```
! exec cat doesntexist
stderr 'cat: doesntexist: No such file or directory'
```

 Dans ce cas, pour éviter que **cat** affiche quoi que ce soit sur sa sortie standard en cas d'erreur, on peut utiliser **! stdout** :

 ```
! exec cat doesntexist
stderr 'cat: doesntexist: No such file or directory'
! stdout .
 ```

### Passer des arguments aux programmes

Nous pouvons passer des arguments au programme exécuté par **exec**, en les plaçant après le nom du programme.

```
exec echo success

exec cat 'data file.txt'
```

Puisque l'apostrophe a pour effet particulier de regrouper les arguments séparés par des espaces, pour écrire une apostrophe littérale lorsque c'est nécessaire nous écrivons deux apostrophes consécutives.

```
exec echo 'Here''s how to escape single quotes'
```

Les guillemets n'ont aucun effet de guillemet dans les scripts de test. Ils sont simplement traités comme des guillemets littéraux :

```
exec echo "This will print with literal double quotes"
```

### Tester les outils CLI avec testscript

Si TestScript se limitait à exécuter des programmes existants avec certains arguments, à vérifier leur réussite (ou leur échec) et à produire certains résultats (ou non), ce serait encore très utile.
Mais nous ne sommes pas limités à l'exécution de programmes existants. Si nous voulons tester notre propre binaire, par exemple, nous n'avons pas besoin de le compiler et de l'installer au préalable pour l'exécuter dans un script. TestScript peut nous épargner cette tâche.

Supposons que nous écrivions un programme nommé hello, par exemple, dont le travail consiste simplement à imprimer un message « hello world » sur le terminal.

```
func main() {
    fmt.Println("hello world")
}
```

La fonction main ne peut pas être appelée directement depuis un test Go ; elle est invoquée automatiquement lors de l'exécution du binaire compilé. Nous ne pouvons pas la tester directement. Nous allons donc déléguer ses tâches à une autre fonction que nous pouvons appeler depuis un test:hello.Main, par exemple.

```
func main() {
    os.Exit(hello.Main())
}
```

Maintenant que nous avons délégué toutes les fonctionnalités du programme à cette fonction hello.Main, nous pouvons maintenant créer un binaire via une fonction **TestMain**.

```
func TestMain(m *testing.M) {
    os.Exit(testscript.RunMain(m, map[string]func() int{
        "hello": hello.Main,
    }))
}
```

Cette fonction **TestMain** en Go sert à préparer l’environnement avant d’exécuter les tests. Ici, elle utilise **testscript.RunMain** pour configurer un binaire nommé **hello** lié à la fonction **hello.Main**. Ce binaire est placé dans un répertoire temporaire ajouté au **$PATH**, ce qui permet de l’exécuter via exec dans les scripts de test, sans compilation manuelle. Une fois les tests terminés, le binaire et son répertoire temporaire seront automatiquement supprimés.

La possibilité de définir et d'exécuter des programmes personnalisés de cette manière est essentielle pour tester les outils en ligne de commande avec TestScript. Nous pouvons invoquer le programme avec les arguments, les variables d'environnement et les fichiers nécessaires pour tester un comportement donné. Ainsi, nous pouvons tester des comportements même complexes avec un minimum de code.

[https://github.com/rogpeppe/go‐internal](https://github.com/rogpeppe/go‐internal)

### Vérification de la couverture des tests des scripts

Une fonctionnalité particulièrement intéressante de **testscript** est qu'il peut même nous fournir des informations de couverture lors du test de notre binaire. Ce serait difficile à réaliser si nous construisions et exécutions le binaire nous-mêmes, mais **testscript** le fait fonctionner pour nous :

```
go test -coverprofile=cover.out
```