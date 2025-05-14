package batch

import (
	"context"
	"time"
)

func RunBatchJob(ctx context.Context) {
	time.Sleep(20 * time.Millisecond)
}

/**
En Go, la concurrence est souvent utilisée pour exécuter des tâches longues, comme des appels à des services distants ou des traitements
par lots, sans bloquer le reste du programme. Pour gérer ces tâches, notamment pour pouvoir les annuler si elles deviennent inutiles ou
trop longues, on utilise le package context. Celui-ci permet de créer un contexte que l'on passe à la fonction concernée. L'appelant peut
ensuite annuler ce contexte pour indiquer à la tâche qu'elle doit s'arrêter.
**/
