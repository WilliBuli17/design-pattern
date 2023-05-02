package refactoring_guru

import "net/http"

type Nginx struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func NewNginx() *Nginx {
	return &Nginx{
		application:       &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}

	n.rateLimiter[url] += 1
	return true
}

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)
	}

	return n.application.HandleRequest(url, method)
}
