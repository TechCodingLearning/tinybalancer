package demo

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type HTTPProxy struct {
	proxy *httputil.ReverseProxy
}

func NewHTTPProxy(target string) (*HTTPProxy, error) {
	u, err := url.Parse(target)
	if err != nil {
		return nil, err
	}
	return &HTTPProxy{httputil.NewSingleHostReverseProxy(u)}, nil
}

func (h *HTTPProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.proxy.ServeHTTP(w, r)
}

func main() {
	proxy, err := NewHTTPProxy("http:127.0.0.1:8088")
	if err != nil {
		log.Fatalln(err)
	}

	http.Handle("/", proxy)
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalln(err)
	}
}