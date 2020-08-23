package httpclient

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sony/gobreaker"
)

func TestBreaker(t *testing.T) {
	Convey("Dado um http client breaker", t, func() {
		Convey("Caso seja passada uma URL errada, deve retornar erro", func() {
			breaker := NewBreaker()
			_, err := breaker.HTTPRequest("Errado", "url errada", "/", nil)
			So(err, ShouldNotBeNil)
		})
		Convey("Caso seja usado com os parâmetros corretos, deve retornar o payload recebido do serviço alvo", func() {
			payload := []byte("payload")
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
				rw.Write(payload)
			}))
			breaker := &Breaker{
				breakers: make(map[string]*gobreaker.CircuitBreaker),
				client:   server.Client(),
			}
			defer server.Close()
			res, err := breaker.HTTPRequest("GET", server.URL, "/", nil)
			So(err, ShouldBeNil)
			So(res, ShouldResemble, payload)
		})
	})
}
