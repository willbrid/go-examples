package script_test

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func Test(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

/**
testscript définit un petit « langage » pour écrire des scripts de test, que nous stockons sous forme de fichiers séparés dans
notre projet, avec notre code source Go.

La fonction Test est une fonction de test Go classique, mais l'appel à testscript.Run est notable. Cette fonction exécute plusieurs scripts de
test comme des sous-tests parallèles, en utilisant t pour signaler les échecs. On lui passe aussi une structure Params pour la configuration,
notamment le dossier testdata/script où sont stockés les scripts de test.

L'instruction exec dans un script testscript sert à la fois à exécuter un programme et à vérifier qu'il réussit (code de sortie nul).
Si le programme échoue, le script s'arrête immédiatement, comme un t.Fatal en Go. Chaque script s'exécute dans un répertoire temporaire vide,
ce qui garantit des tests reproductibles et permet de détecter facilement les erreurs liées à l'absence de fichiers ou de répertoires
nécessaires à l'exécution du programme.

stdout est aussi une assertion, concernant cette fois le contenu de la sortie standard de l'instruction exec.
stdout affirme que la sortie correspondra à une expression régulière, dans ce cas la chaîne hello world suivie d'un saut de ligne.
Le programme peut afficher plus que cela, mais ce script affirme qu'il produit au moins hello world\n parmi ses sorties.
Dans le cas contraire, l'assertion stdout échouera, et ce sous-test échouera également.


Exemple de résultat d'un cas d'échec du script

--- FAIL: Test (0.00s)
    --- FAIL: Test/hello (0.00s)
        testscript.go:584: > exec echo 'hello world'
            [stdout]
            hello world
            > stdout 'ghello world\n'
            FAIL: testdata/script/hello.txtar:2: no match for `ghello world\n` found in stdout

Le test parent défaillant est Test, et le nom du sous-test défaillant représenté par le script est hello. Ce nom est dérivé du nom de
fichier du script, hello.txtar, mais sans l'extension .txtar
--- FAIL: Test/hello (0.00s)

Chaque ligne du script est affichée dans la sortie du test au fur et à mesure de son exécution, préfixée par un caractère >
The standard output of the program run by exec is shown next in the test output:
[stdout]
hello world

Et voici la ligne qui a réellement déclenché l’échec du test :
> stdout 'ghello world\n'
FAIL: testdata/script/hello.txtar:2: no match for `ghello world\n` found in stdout

L'assertion stdout échoue, car la sortie du programme ne correspond pas à l'expression donnée, ce qui entraîne l'échec du test.
En plus d'afficher la ligne défaillante, le message indique également le nom du fichier du script et le numéro de la ligne où le trouver.
testdata/script/hello.txtar:2

La fonction de test invoque testscript.Run, qui exécute tous les scripts présents dans le répertoire testdata/script comme des sous-tests
parallèles. On peut organiser les scripts dans un seul répertoire ou les répartir dans plusieurs selon les besoins du projet, bien que
la plupart des projets utilisent un seul dossier.
Il est aussi possible d’exécuter un seul script spécifique en utilisant l’option -run, comme pour tout autre test Go.
go test -run Test/hello

Il est judicieux de limiter la taille de chaque script et de le concentrer sur un ou deux comportements connexes. Par exemple,
vous pourriez avoir un script pour vérifier les différents types de comportements « entrée invalide », et un autre pour
la partie de réussite. Un script trop chargé est difficile à lire et à comprendre, tout comme un test trop chargé.
**/
