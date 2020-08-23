package httpclient

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/sony/gobreaker"
)

// Breaker contém um map de circuit breakers que funcionam como wrappers para executar ações custosas e com possibilidade de falha, como uma chamada http
type Breaker struct {
	breakers map[string]*gobreaker.CircuitBreaker
	client   *http.Client
}

// NewBreaker retorna uma instancia de Breaker
func NewBreaker() *Breaker {
	return &Breaker{
		breakers: make(map[string]*gobreaker.CircuitBreaker),
		client:   http.DefaultClient,
	}
}

// HTTPRequest usa um circuit breaker como wrapper para executar uma chamada http
func (b *Breaker) HTTPRequest(method string, URL string, path string, body io.Reader) ([]byte, error) {
	var breaker *gobreaker.CircuitBreaker
	if val, ok := b.breakers[URL]; ok {
		breaker = val
	} else {
		var st gobreaker.Settings
		st.Name = URL
		st.ReadyToTrip = func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.6
		}
		breaker = gobreaker.NewCircuitBreaker(st)
		b.breakers[URL] = breaker
	}
	resBody, err := breaker.Execute(func() (interface{}, error) {
		req, err := http.NewRequest(method, fmt.Sprintf("%s%s", URL, path), body)
		if err != nil {
			return nil, err
		}
		res, err := b.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	})
	if err != nil {
		return nil, err
	}
	return resBody.([]byte), nil
}
