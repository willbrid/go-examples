package main

import (
	"os"

	"timer"
)

func main() {
	t := timer.NewTimerFromArgs(os.Args)
	t.Sleep()
}

/**
Une solution à la question épineuse du test des outils CLI consiste à découpler l'interface de ligne de commande du reste du programme et
à les tester séparément. Une conception découplée est une meilleure conception.
**/
