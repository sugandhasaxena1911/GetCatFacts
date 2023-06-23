package controllers

import (
	"net/http"
	"testing"
)

func TestGetCatFacts(t *testing.T) {
	statusCode := makeGetRequest("https://cat-fact.herokuapp.com/facts", "")
	if statusCode != 200 {
		t.Error("Expected 200. Got ", statusCode)
	}
}

func BenchmarkGetCatFacts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetCatFacts()
	}

}

func makeGetRequest(url string, body interface{}) int {
	r, _ := http.Get(url)
	return r.StatusCode

}
