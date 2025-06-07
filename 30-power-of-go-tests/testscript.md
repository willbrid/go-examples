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

### Comparaison de la sortie avec les fichiers à l'aide de cmp

Par exemple, supposons que nous souhaitions comparer la sortie du programme non pas à une chaîne ou à une expression régulière, mais au contenu d'un fichier (fichier de référence).
Nous pouvons fournir un fichier de référence dans le fichier de script lui-même, en délimitant son contenu par une ligne de marquage spéciale commençant et se terminant par un double tiret (--).

```
exec echo "hello world"
cmp stdout golden.txt

-- golden.txt --
hello world
```

La ligne de marquage contenant **golden.txt** ouvre une entrée de fichier : tout ce qui suit sera écrit dans **golden.txt** et placé dans le répertoire de travail du script avant son exécution.
L'assertion **cmp** permet de comparer deux fichiers pour vérifier leur concordance. Si les correspondances sont exactes, le test réussit. Dans le cas contraire, l'échec sera accompagné d'un différentiel indiquant les parties non concordantes.

Alternativement, nous pouvons utiliser **!** pour annuler la comparaison, auquel cas les fichiers ne doivent pas correspondre, et le test échouera s'ils le font :

```
exec echo hello
! cmp stdout golden.txt

-- golden.txt --
goodbye world
```

Le premier argument de **cmp** peut être le nom d'un fichier, mais on peut aussi utiliser le nom spécial **stdout**, qui correspond à la sortie standard de l'exécutable précédent. De même, **stderr** désigne la sortie d'erreur standard.

Si le programme produit une sortie différente selon la valeur d'une variable d'environnement, on peut utiliser l'assertion **cmpenv**. Cela fonctionne comme **cmp**, mais en interpolant les variables d'environnement dans le fichier source.

```
exec echo Running with home directory $HOME
cmpenv stdout golden.txt

-- golden.txt --
Running with home directory $HOME
```

Lors de l'exécution de ce script, la variable **$HOME** de la commande **echo** sera étendue à la valeur réelle de la variable d'environnement **HOME**, quelle qu'elle soit. Cependant, comme nous utilisons **cmpenv**, nous étendons également la variable **$HOME** du fichier source à la même valeur.

### Plus de correspondances : exists, grep et -count

Certains programmes créent des fichiers directement, sans afficher de sortie sur le terminal. Si nous souhaitons simplement vérifier l'existence d'un fichier donné suite à l'exécution du programme, sans nous soucier de son contenu, nous pouvons utiliser l'assertion **exists**.

Par exemple, supposons un programme **myprog** qui écrit sa sortie dans un fichier spécifié par l'option **-o**. Nous pouvons vérifier l'existence de ce fichier après l'exécution du programme en utilisant exists.

```
exec myprog -o results.txt
exists results.txt
```

Si nous souhaitons comparer le contenu exact du fichier de résultats, nous pouvons utiliser **cmp** pour le comparer à un fichier.

```
exec myprog -o results.txt
cmp results.txt golden.txt

-- golden.txt --
hello world
```

Si les deux fichiers correspondent parfaitement, l'assertion réussit. Sinon, elle échoue et génère un différentiel indiquant la non-correspondance. Si le fichier de résultats n'existe pas, c'est également un échec.
En revanche, si nous ne devons pas rechercher la correspondance du fichier entier, mais seulement d'une partie, nous pouvons utiliser l'assertion **grep** pour rechercher une expression régulière.

```
exec myprog -o results.txt
grep '^hello' results.txt

-- golden.txt --
hello world
```

Une assertion **grep** réussit si le fichier correspond à l'expression donnée au moins une fois, quel que soit le nombre de correspondances. En revanche, s'il est important d'avoir un nombre précis de correspondances, nous pouvons utiliser l'option **-count** pour spécifier ce nombre.

```
grep -count=1 'beep' result.txt

-- result.txt --
beep beep
```

Dans cet exemple, nous avons spécifié que le modèle `beep` ne doit correspondre qu'une seule fois dans le fichier cible.

Par défaut le répertoire de travail du script est automatiquement supprimé après le test, nous ne pouvons pas consulter son contenu. Pour conserver ce répertoire en cas de dépannage, nous pouvons utiliser l'option **-testwork** de la commande **go test** : cela préservera le répertoire de travail du script et affichera également son environnement, y compris la variable **WORK** qui indique où trouver ce répertoire.

### Le format txtar : construction de fichiers de données de test

Nous pouvons créer des fichiers dans le répertoire de travail du script, en utilisant cette syntaxe spéciale pour indiquer une ou plusieurs entrées de fichier.

```
-- filename1.txt --
... file1 contents ...

-- filename2.txt --
... file2 contents ...

-- filename3.txt --
... file3 contents ...
```

La ligne commençant par **--**, appelée ligne de **marqueur de fichier**, indique à **testscript** que tout ce qui suit cette ligne (jusqu'au marqueur de fichier suivant) doit être traité comme le contenu du fichier.
Une ligne de marqueur de fichier doit commencer par deux tirets et un espace et se terminer par un espace et deux tirets. La partie entre ces marqueurs spécifie le nom du fichier, qui sera débarrassé de tout espace.

Chaque ligne de marqueur indique le début d'un nouveau fichier, suivi de zéro ou plusieurs lignes de contenu, et se terminant à la ligne de marqueur de fichier suivante, le cas échéant. Tous ces fichiers seront créés dans le répertoire de travail avant le démarrage du script.

Si nous devons créer des dossiers, voire des arborescences complètes de fichiers et de dossiers, nous pouvons le faire en utilisant des chemins séparés par des barres obliques dans les noms de fichiers.

```
-- folder1/filename1.txt --
... file1 contents ...

-- folder2/subfolder2/filename2.txt --
... file2 contents ...

-- folder3/filename3.txt --
... file3 contents ...
```

Le format txtar (abréviation de **text archive**) permet de définir facilement un ensemble de fichiers et dossiers directement dans un fichier texte, souvent utilisé dans les tests avec **testscript**. Cela évite d’écrire du code Go ou de copier des fichiers depuis **testdata**. Un fichier **.txtar** peut contenir plusieurs fichiers texte et être utilisé aussi en dehors des tests, comme format simple d'archivage. Pour l’exploiter dans un programme Go, on peut importer le package **txtar**.