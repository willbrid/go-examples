package batch_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"batch"
)

func TestRunBatchJob_RunsJobWithinReasonableTime(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	go func() {
		batch.RunBatchJob(ctx)
		cancel()
	}()
	<-ctx.Done()
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		t.Fatal("timed out")
	}
}

/**
Un contexte avec un timeout de 10 ms est créé et son annulation est différée pour libérer les ressources.
La tâche RunBatchJob est lancée dans une goroutine avec ce contexte. Le test principal attend ensuite la fin de la tâche en bloquant
sur ctx.Done(), un canal fermé soit à la fin de la tâche, soit à l’expiration du timeout.

Lorsque le canal est fermé, on vérifie ctx.Err() pour savoir pourquoi :

- Si la tâche s’est terminée à temps, l’erreur sera context.Canceled, ce qui est un succès.
- Si le délai est dépassé, l’erreur sera context.DeadlineExceeded, ce qui échoue le test avec t.Fatal.

Ce mécanisme permet de s'assurer qu'une tâche longue ne bloque pas indéfiniment et que le test reste maîtrisé.
**/
