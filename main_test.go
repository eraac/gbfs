package gbfs

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	server *httptest.Server
	serverHTTPOptions []HTTPOption
)

func TestMain(m *testing.M) {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("files")))
	mux.HandleFunc("/en/500.json", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	server = httptest.NewServer(mux)
	defer server.Close()

	serverHTTPOptions = []HTTPOption{
		HTTPOptionBaseURL(server.URL),
		HTTPOptionLanguage(lang),
		HTTPOptionForceURL(FeedKeyAutoDiscovery, fmt.Sprintf("%s/%s.json", server.URL, FeedKeyAutoDiscovery)),
	}

	code := m.Run()

	os.Exit(code)
}
