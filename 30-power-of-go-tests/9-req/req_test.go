package req_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"req"
)

func TestRequest_ReturnsErrRateLimitWhenRatelimited(t *testing.T) {
	t.Parallel()

	ts := newRateLimitingServer()
	defer ts.Close()

	err := req.Request(ts.URL)
	if !errors.Is(err, req.ErrRateLimit) {
		t.Errorf("wrong error: %v", err)
	}
}

func newRateLimitingServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusTooManyRequests)
		}))
}

// Prenons le scénario où nous requetons vers un serveur qui nous renvoie http 429 à partir d'un seuil de nombre de requêtes.
// Une valeur sentinelle conviendrait bien ici car c'est utile pour les utilisateurs de notre package, ainsi qu'à des fins de test.
// Et puisque nous disons qu’il s’agit d’une partie importante de notre comportement face à l’utilisateur que la fonction renvoie
// cette valeur d’erreur spécifique quand elle le devrait, elle doit être testée.
// La fonction errors.Is(err error, target error) permet de comparer les valeurs d'interface err et target.
