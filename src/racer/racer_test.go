package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	// 在标准库中有一个 net/http/httptest 包，它可以让你轻易建立一个 HTTP 模拟服务器（mock HTTP server）
	// httptest.NewServer 接受一个我们传入的 匿名函数 http.HandlerFunc。
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	// 在某个函数调用前加上 defer 前缀会在 包含它的函数结束时 调用它。
	defer slowServer.Close()
	defer fastServer.Close()
	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
