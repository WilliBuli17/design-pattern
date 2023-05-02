package flyweight

import (
	"design-patterns/structural/proxy/refactoring_guru"
	"fmt"
	"net/http"
	"testing"
)

func TestNginX(t *testing.T) {
	nginxServer := refactoring_guru.NewNginx()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"

	httpCode, body := nginxServer.HandleRequest(appStatusURL, http.MethodGet)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, http.MethodGet)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, http.MethodGet)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createuserURL, http.MethodPost)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createuserURL, http.MethodGet)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
}
