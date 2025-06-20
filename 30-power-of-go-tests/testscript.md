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

### Fournir des entrées aux programmes à l'aide de stdin

Puisqu’on peut créer des fichiers arbitraires dans le répertoire de travail d’un script, on peut s’en servir pour simuler une saisie utilisateur ou fournir des données à un programme, comme si elles étaient entrées via un shell ou en mode interactif.

On peut aller encore plus loin avec **stdin**. On ne se limite pas à fournir une entrée à partir d'un fichier ; on peut aussi utiliser la sortie d'un exécutable précédent. Cela peut être utile lorsqu'un programme doit générer une sortie qui sera redirigée vers un autre, par exemple :

```
exec echo hello
stdin stdout
exec cat
stdout 'hello'
```

### Opérations sur les fichiers

Tout comme dans un script shell traditionnel, nous pouvons copier un fichier vers un autre en utilisant **cp**.

```
cp a.txt b.txt
```

Cependant, le premier argument de **cp** peut également être **stdout** ou **stderr**, indiquant que nous voulons copier la sortie d'un exécutable précédent dans un fichier.

```
exec echo hello
cp stdout tmp.txt
```

Nous pouvons également utiliser **mv** pour déplacer un fichier (c'est-à-dire le renommer).

```
mv a.txt b.txt
```

Nous pouvons également créer un répertoire en utilisant **mkdir**, puis y copier plusieurs fichiers avec **cp**.

```
mkdir data
cp a.txt b.txt c.txt data
```

L'instruction **cd** modifiera le répertoire actuel pour les programmes ultérieurs exécutés par **exec**.

```
cd data
```

Pour supprimer un fichier ou un répertoire, nous utilisons l'instruction **rm**. Utilisée avec un répertoire, **rm** agit de manière récursive : elle supprime tous les fichiers et sous-répertoires qu'il contient avant de supprimer le répertoire lui-même.

```
rm
```

Pour créer un lien symbolique d'un fichier ou d'un répertoire vers un autre, nous pouvons utiliser l'instruction **symlink**. Pour cela nous utilisons le caractère obligatoire **->** qui indique la direction du lien symbolique.

```
mkdir target
symlink source -> target
```

### Différences avec les scripts shell

Les scripts de test ressemblent aux scripts shell, mais sans structures de contrôle comme les boucles ou fonctions. Un échec d’assertion arrête simplement le script. Cela les rend simples et lisibles. Contrairement aux scripts shell, les commandes **exec** sont exécutées directement, sans passer par un shell, donc des fonctionnalités comme le globbing (*) ne sont pas disponibles.

Si nous devons développer une expression glob, nous pouvons simplement demander au shell de le faire pour nous.

```
exec sh -c 'ls .*'
```

Si nous souhaitons utiliser le caractère pipe (**|**) (sans utiliser la technique **stdin** et **stdout**) pour envoyer la sortie d'une commande à une autre, nous pouvons simplement demander au shell de le faire pour nous.

```
exec sh -c 'echo hello | wc -l'
```

### Commentaires et phases

Dans un script de test, un **#** démarre un commentaire. Tout ce qui suit sur la ligne est ignoré, ce qui permet d’ajouter des explications après les commandes.

```
exec echo hello # this comment will not appear in output
```

Les commentaires ne servent pas uniquement à documenter. Ils délimitent aussi des phases du script. Lors d’un échec, **testscript** n’affiche que le journal de la phase en cours ; les phases précédentes sont résumées par leurs commentaires, indiquant qu’elles ont réussi.

```
# run an existing command: this will succeed
exec echo hello

# try to run a command that doesn't exist
exec bogus
```

### Conditions

Dans les scripts de test, il peut être nécessaire de conditionner l'exécution d'une action selon l’environnement (ex. : la présence d’un programme). Pour cela, on peut précéder une ligne de script par une condition qui vérifie si l’action doit être exécutée ou non.

```
[exec:sh] exec echo yay, we have a shell
```

Les crochets contiennent une condition, qui peut être vraie ou fausse. Si la condition est vraie, le reste de la ligne est exécuté. Sinon, elle est ignorée.
Dans cet exemple, la condition **[exec:sh]** est vraie si un programme nommé sh est présent dans notre $PATH et que nous avons l'autorisation d'exécuter. Dans le cas contraire, cette ligne du script sera ignorée.

Si nous devons vérifier la présence d'un programme sur un chemin spécifique, nous pouvons indiquer ce chemin dans la condition.

```
[exec:/bin/sh] exec echo yay, we have /bin/sh
```

Parmi les conditions intégrées les plus utiles dans les scripts de test figurent celles permettant de vérifier la **version de Go**, le **système d’exploitation** et l’**architecture du processeur** de la machine d’exécution.

```
# 'go1.x' is true if this is Go 1.x or higher
[go1.99] exec echo 'We have at least Go 1.99'

# Any known value of GOOS is also a valid condition
[darwin] exec echo 'We''re on macOS'

# As is any known value of GOARCH
[!arm64] exec echo 'This is a non-arm64 machine'
```

Comme dans l'exemple ci-dessus, nous pouvons également nier une condition en la préfixant avec **!** .

Nous pouvons utiliser l'instruction **skip** pour ignorer le test dans certaines circonstances.

```
# Skip this test unless we have Go 1.99 or later
[!go1.99] skip

# Skip this test on Linux
[linux] skip
```

La condition **Unix** est vraie lorsque le système d'exploitation cible est l'un de ceux que Go considère comme « de type **Unix** », notamment **Linux**, **macOS**, **FreeBSD** et autres.

```
[unix] exec echo 'It''s a UNIX system! I know this!'
```

En fait, les seuls systèmes d’exploitation pris en charge par Go pour lesquels la condition **Unix** n’est pas vraie sont **js**, **nacl**, **plan9**, **Windows** et **zos**.

### Définition des variables d'environnement avec env

Pour définir une variable spécifique pour la durée du script, nous utilisons une instruction **env**. <br>
Par exemple, pour tester qu'un programme **myprog** échoue et affiche un message d'erreur lorsque la variable d'environnement dont il a besoin n'a pas de valeur, nous pourrions écrire :

```
env AUTH_TOKEN=
! exec myprog
stderr 'AUTH_TOKEN must be set'
```

Dans les scripts, nous pouvons faire référence à la valeur d'une variable d'environnement en utilisant le symbole **$** suivi du nom de la variable, comme dans un script shell.

```
exec echo $PATH
```

Chaque script démarre avec un environnement vide, à l'exception de **$PATH** et de quelques autres variables prédéfinies. Par exemple, **HOME** est défini sur **/no-home**, car de nombreux programmes s'attendent à pouvoir trouver le répertoire personnel de l'utilisateur dans **$HOME**.

Une variable utile pour les tests multiplateformes est **$exe**. Sa valeur sous Windows est **.exe**, mais sur toutes les autres plateformes, c'est une chaîne vide.

```
exec prog$exe
```

### Transmission de valeurs aux scripts via des variables d'environnement

Pour fournir des valeurs dynamiques à un script (ex. : une adresse générée au moment du test), on peut utiliser des variables d’environnement. Cela se fait via la fonction **Setup** passée à **testscript.Run** dans les paramètres. Cette méthode permet d’injecter des données calculées au moment de l’exécution dans les scripts de test.

La fonction **Setup** s'exécute juste avant le lancement du script et reçoit un objet **Env** représentant l’environnement du script. Elle peut utiliser **env.Setenv** pour définir des variables d’environnement personnalisées (ex. : **SERVER_ADDR**). Le script peut ensuite y accéder via **$SERVER_ADDR**, comme avec toute variable d’environnement classique.

### Exécution de programmes en arrière-plan avec &

Les scripts de test permettent, tout comme les shells classiques, d’exécuter des programmes en arrière-plan. Cela signifie que le programme démarre comme un processus séparé, et que le script continue immédiatement son exécution, sans attendre la fin de ce programme.

```
exec sleep 10 &
```

Dans un script testscript, ajouter une esperluette (**&**) à la fin d’une commande lance le programme en arrière-plan, sans bloquer le reste du script. Son exécution est mise en mémoire tampon, ce qui permet d’y accéder plus tard. Pour s'assurer qu'un script n'expire pas avant la fin de ces programmes (programes en arrière plan), on utilise l’instruction **wait**, qui attend explicitement leur terminaison.

L’instruction **wait** suspend l’exécution du script jusqu’à ce que tous les programmes lancés en arrière-plan soient terminés. Leurs sorties (stdout, stderr) peuvent ensuite être testées comme s’ils avaient été exécutés au premier plan.

```
exec sleep 10 &
wait
```

### L'exécuteur de scripts de test autonome

Le langage de scripts de test est tellement pratique qu’il serait utile de l’utiliser en dehors des tests Go, par exemple dans des pipelines CI ou des projets non-Go. Bonne nouvelle : c’est possible grâce à un outil autonome permettant d’exécuter des scripts de test directement depuis la ligne de commande, sans écrire de code Go. Pour l'installer on exécute la commande :

```
go install github.com/rogpeppe/go-internal/cmd/testscript@latest
```

Pour l'utiliser, il suffit de donner le chemin vers un script, ou plusieurs scripts.

```
testscript testdata/script/*
```

Cela exécutera chaque script tour à tour et affichera PASS s'il réussit (ainsi que les commentaires décrivant chaque phase réussie). Sinon, le code affichera FAIL, accompagné du même message d'erreur que lors de l'exécution du script dans un test Go. <br>
Si un script échoue, le code de sortie de testscript sera 1, ce qui est utile pour détecter les échecs dans les automatisations.

Pour enregistrer l'activité du script, nous pouvons utiliser l'option **-v**, qui affiche un message détaillé, que le script réussisse ou échoue.

```
testscript -v echo.txtar
```

Pour transmettre des variables d'environnement aux scripts, nous pouvons les spécifier à l'aide de l'option **-e**. Nous répétons l'option **-e** pour chaque paire variable-valeur.

```
testscript -e VAR1=hello -e VAR2=goodbye script.txtar
```

Tout comme lors de l'exécution de scripts à partir de tests Go, chaque script dispose de son propre répertoire de travail, qui est ensuite nettoyé. Pour conserver ce répertoire et son contenu, par exemple pour le dépannage, nous utilisons l'option **-work**.

```
testscript -work script.txtar
```

Il est possible d'utiliser testscript comme un interpréteur de script, grâce à la ligne **shebang (#!)**, sur les systèmes Unix comme Linux ou macOS. Cela permet d'exécuter directement un fichier **.txtar** (exemple ci-dessus fichier **hello.txtar**) comme s’il s’agissait d’un script autonome.

```
vim hello.txtar
```

```
#!/usr/bin/env testscript
exec echo hello
stdout 'hello'
```

Ainsi, si nous modifions les permissions du fichier pour le rendre exécutable, nous pouvons exécuter ce script directement depuis la ligne de commande.

```
chmod +x hello.txtar

./hello.txtar
```

### Scripts de test en tant que reproductions de problèmes

Le format **txtar** est particulièrement utile pour soumettre des rapports de bugs, car il permet de représenter à la fois un script de test et plusieurs fichiers/dossiers dans un seul bloc de texte facilement copiable. Il est souvent utilisé dans l’écosystème Go pour illustrer des problèmes de manière concise dans les demandes d’incidents.

```
# I was promised 'Go 2', are we there yet?
exec go version
stdout 'go version go2.\d+'
```

Les mainteneurs peuvent utiliser le script fourni pour reproduire facilement un bug, tester une correction proposée et même l’intégrer aux tests automatisés du projet. Le format testscript simplifie ainsi la soumission, la reproduction et la discussion des cas de test, ce qui encourage l’adoption des tests automatisés, même dans des projets où ils sont peu utilisés.

### Scripts de test comme tests

Au lieu d’exécuter des scripts de test avec **go test**, on peut faire l’inverse : écrire un script de test qui lance **go test**. <br>
Avec testscript, on peut écrire un seul script qui vérifie automatiquement si chaque test donne bien le résultat attendu (succès ou échec).

```
! exec go test
stdout 'want 4, got 5'
```