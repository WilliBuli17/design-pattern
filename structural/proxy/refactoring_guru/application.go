package refactoring_guru

import "net/http"

type Application struct {
}

func (a *Application) HandleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == http.MethodGet {
		return http.StatusOK, http.StatusText(http.StatusOK)
	}
	if url == "/create/user" && method == http.MethodPost {
		return http.StatusCreated, http.StatusText(http.StatusCreated)
	}

	return http.StatusNotFound, http.StatusText(http.StatusNotFound)
}
