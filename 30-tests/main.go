package main

import (
	"fmt"
	"log"
	"sort"
)

/**
Les tests unitaires sont définis dans des fichiers dont le nom se termine par _test.go

La bibliothèque standard Go prend en charge l'écriture de tests unitaires via le package testing. Les tests unitaires sont exprimés sous forme de
fonctions dont le nom commence par Test, suivi d'un terme commençant par une lettre majuscule, comme TestSum.
(La lettre majuscule est importante car les outils de test ne reconnaîtront pas un nom de fonction tel que Testsum comme test unitaire.)

Les fonctions de test unitaire reçoivent un pointeur vers une classe T (testing.T), qui définit les méthodes de gestion des tests et de rapport des résultats
des tests. Les tests Go ne reposent pas sur des assertions et sont écrits à l'aide d'instructions de code standard. Tout ce dont les outils de test
se soucient, c'est de savoir si le test échoue, ce qui est signalé à l'aide des méthodes :
- Log(...vals) : cette méthode écrit les valeurs spécifiées dans le journal des erreurs de test.
- Logf(template, ...vals) : cette méthode utilise le modèle et les valeurs spécifiés pour écrire un message dans le journal des erreurs de test.
- Fail() : l'appel de cette méthode marque le test comme ayant échoué mais poursuit l'exécution du test.
- FailNow() : l'appel de cette méthode marque le test comme ayant échoué et arrête l'exécution du test.
- Failed() ; cette méthode renvoie true si le test a échoué.
- Error(...errs) : l'appel de cette méthode équivaut à appeler la méthode Log, suivie de la méthode Fail.
- Errorf(template, ...vals) : l'appel de cette méthode équivaut à appeler la méthode Logf, suivie de la méthode Fail.
- Fatal(...vals) : l'appel de cette méthode équivaut à appeler la méthode Log, suivie de la méthode FailNow.
- Fatalf(template, ...vals) : l'appel de cette méthode équivaut à appeler la méthode Logf, suivie de la méthode FailNow.

Un fichier de test peut contenir plusieurs tests.

La commande <go test> ne rapporte aucun détail par défaut, mais plus d'informations peuvent être générées en exécutant la commande : <go test -v>

La commande go test peut être utilisée pour exécuter des tests sélectionnés par leur nom : <go test -v -run "um">
Les tests sont sélectionnés avec une expression régulière. Le seul test dont le nom correspond à l'expression est <TestSum>.

La classe T (testing.T) fournit également un ensemble de méthodes pour gérer l'exécution des tests :
- Run(name, func) : l'appel de cette méthode exécute la fonction spécifiée en tant que sous-test. La méthode se bloque pendant que le test est exécuté
                    dans sa propre goroutine et renvoie un booléen qui indique si le test a réussi.
					La méthode Run est utilisée pour exécuter un sous-test, ce qui est un moyen pratique d'exécuter une série de tests liés à partir
					d'une seule fonction.
- SkipNow() : l'appel de cette méthode arrête l'exécution du test et le marque comme ignoré.
- Skip(...args) : cette méthode équivaut à appeler la méthode Log, suivie de la méthode SkipNow.
- Skipf(template, ...args) : cette méthode équivaut à appeler la méthode Logf, suivie de la méthode SkipNow.
- Skipped() : cette méthode renvoie true si le test a été ignoré.

Les tests peuvent être ignorés, ce qui peut être utile lorsqu'un échec d'un test signifie qu'il est inutile d'effectuer des tests connexes.


# Code d'analyse comparative
Les fonctions dont le nom commence par Benchmark, suivi d'un terme commençant par une lettre majuscule sont des benchmarks, dont l'exécution est temporisée.
Les fonctions de benchmark reçoivent un pointeur vers la classe testing.B, qui définit le champ :
- N : ce champ int spécifie le nombre de fois que la fonction de benchmark doit exécuter le code à mesurer.

La valeur de N est utilisée dans une boucle for au sein de la fonction benchmark pour répéter le code dont les performances sont mesurées.
Les outils benchmark peuvent appeler la fonction benchmark à plusieurs reprises, en utilisant différentes valeurs de N, pour établir une mesure stable.

Pour effectuer le benchmark, nous exécutons la commande : go test -bench . -run notest

La période qui suit l'argument -bench entraîne l'exécution de tous les tests de performance découverts par l'outil de test go. Le point peut être
remplacé par une expression régulière pour sélectionner des repères spécifiques. Par défaut, les tests unitaires sont également effectués, mais nous avons
utilisé l'argument -run pour spécifier une valeur qui ne correspondra à aucun des noms de fonction de test dans le projet, avec pour résultat
que seuls les tests de performance seront effectués.

Le nom de la fonction de référence est suivi du nombre de processeurs ou de cœurs.
Le champ suivant indique la valeur de N qui a été transmise à la fonction de référence pour générer ces résultats.
La valeur finale rend compte de la durée, en nanosecondes, nécessaire pour effectuer chaque itération de la boucle de référence.

Pour chaque itération de la boucle for, la fonction BenchmarkSort doit générer des données aléatoires, et le temps nécessaire pour produire ces données
est inclus dans les résultats du benchmark. La classe testing.B définit les méthodes utilisées pour contrôler la minuterie utilisée pour l'analyse comparative.
- StopTimer() : cette méthode arrête le chronomètre.
- StartTimer() : cette méthode démarre le minuteur.
- ResetTimer() : cette méthode réinitialise la minuterie.

La méthode ResetTimer est utile lorsqu'un benchmark nécessite une configuration initiale, et les autres méthodes sont utiles lorsqu'il y a
une surcharge associée à chaque activité de benchmark.

Une fonction de benchmark peut exécuter des sous-benchmarks, tout comme une fonction de test peut exécuter des sous-tests.
La méthode utilisée pour exécuter un sous-benchmark :
- Run(name, func) : l'appel de cette méthode exécute la fonction spécifiée en tant que sous-benchmark. La méthode se bloque pendant l'exécution du benchmark.


Le package log fournit une API de journalisation simple qui crée des entrées de log et les envoie à un io.Writer, permettant à une application de générer
des données de journalisation sans avoir besoin de savoir où ces données seront stockées.
Les fonctions les plus utiles définies par le package log :
- Output() : cette fonction renvoie le Writer auquel les messages de log seront passés. Par défaut, les messages de log sont écrits sur la sortie standard.
- SetOutput(writer) : cette fonction utilise le Writer spécifié pour la journalisation.
- Flags() : cette fonction renvoie les flags utilisés pour formater les messages de journalisation.
- SetFlags(flags) : cette fonction utilise les indicateurs spécifiés pour formater les messages de journalisation.
- Prefix() : cette fonction renvoie le préfixe appliqué aux messages de journalisation. Il n'y a pas de préfixe par défaut.
- SetPrefix(prefix) : cette fonction utilise la chaîne spécifiée comme préfixe pour les messages de journalisation.
- Output(depth, message) : cette fonction écrit le message spécifié dans le Writer renvoyé par la fonction Output, avec la profondeur d'appel spécifiée,
                           qui est par défaut de 2. La profondeur d'appel est utilisée pour contrôler la sélection du fichier de code et
						   n'est généralement pas modifiée .
- Print(...vals) : cette fonction crée un message de journal en appelant fmt.Sprint et en transmettant le résultat à la fonction Output.
- Printf(template, ...vals) : cette fonction crée un message de journal en appelant fmt.Sprintf et en transmettant le résultat à la fonction Output.
- Fatal(...vals) : cette fonction crée un message de journal en appelant fmt.Sprint, transmet le résultat à la fonction Output, puis termine l'application.
- Fatalf(template, ...vals) : cette fonction crée un message de journal en appelant fmt.Sprintf, transmet le résultat à la fonction Output,
                              puis termine l'application.
- Panic(...vals) : cette fonction crée un message de journal en appelant fmt.Sprint, puis transmet le résultat à la fonction Output, puis à la fonction panic.
- Panicf(template, ...vals) : cette fonction crée un message de journal en appelant fmt.Sprintf et transmet le résultat à la fonction Output
                              puis à la fonction panic.

Le format des messages de log est contrôlé avec la fonction SetFlags, pour laquelle le package log définit les constantes :
- Ldate : la sélection de cet indicateur inclut la date dans la sortie du journal.
- Ltime : la sélection de cet indicateur inclut l'heure dans la sortie du journal.
- Lmicrosecondes : la sélection de cet indicateur inclut les microsecondes dans le temps.
- Llongfile : la sélection de cet indicateur inclut le nom du fichier de code, y compris les répertoires, et le numéro de ligne qui a consigné le message.
- Lshortfile : la sélection de cet indicateur inclut le nom du fichier de code, à l'exclusion des répertoires, et le numéro de ligne qui a consigné le message.
- LUTC : la sélection de cet indicateur utilise UTC pour les dates et les heures, au lieu du fuseau horaire local.
- Lmsgprefix : la sélection de cet indicateur déplace le préfixe de sa position par défaut, qui est au début du message de journal,
               juste avant la chaîne transmise à la fonction Output.
- LstdFlags : cette constante représente le format par défaut, qui sélectionne Ldate et Ltime.

Le package log peut être utilisé pour configurer différentes options de journalisation afin que différentes parties de l'application puissent écrire
des messages de log vers différentes destinations ou utiliser différentes options de formatage. La fonction ci-après est utilisée pour créer
une destination de journalisation personnalisée :
- New(writer, prefix, flags) : cette fonction renvoie un Logger qui écrira des messages au writer spécifié, configuré avec le préfixe et les flags spécifiés.
**/

/*
*
La classe Logger est créée avec un nouveau préfixe et l'ajout de l'indicateur Lmsgprefix, en utilisant le Writer obtenu à partir de la fonction Output.
Le résultat est que les messages de log sont toujours écrits vers la même destination, mais avec un préfixe supplémentaire qui désigne
les messages de la fonction sortAndTotal.
*
*/
func sortAndTotal(vals []int) (sorted []int, total int) {
	var logger = log.New(log.Writer(), "sortAndTotal: ", log.Flags()|log.Lmsgprefix)
	logger.Printf("Invoked with %v values", len(vals))
	sorted = make([]int, len(vals))
	copy(sorted, vals)
	sort.Ints(sorted)
	logger.Printf("Sorted data : %v", sorted)
	for _, val := range sorted {
		total += val
		// total++ // Ici on fixe l'erreur qui a conduit à un échec de test
	}
	logger.Printf("Total : %v", total)
	return
}

func main() {
	nums := []int{100, 20, 1, 7, 84}
	sorted, total := sortAndTotal(nums)
	fmt.Println("Sorted Data:", sorted)
	fmt.Println("Total:", total)

	log.Print("Sorted Data:", sorted)
	log.Print("Total:", total)
}

/*
*
La fonction d'initialisation utilise la fonction SetFlags pour sélectionner les indicateurs Lshortfile et Ltime, qui incluront le nom du fichier et
l'heure dans la sortie de journalisation. Dans la fonction main, les messages de log sont créés à l'aide de la fonction log.Print.
*
*/
func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}
