package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func init() {
	conf = &Config{
		Addr:  ":8891",
		Dir:   "/tmp",
		Debug: true,
	}

	go main()
	<-time.After(100 * time.Millisecond)
}

func TestServerGetExistingFile(t *testing.T) {
	ioutil.WriteFile("/tmp/test.txt", []byte("hello!"), 0644)
	res, err := http.Get("http://localhost:8891/test.txt")
	if err != nil {
		t.Errorf("Expected perform request correctly, got error: %v", err)
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	if string(body) != "hello!" {
		t.Errorf("Expected to get correct body, got: '%s'", body)
	}
}

func TestServerGetMissingFile(t *testing.T) {
	res, err := http.Get("http://localhost:8891/not-exists-for-sure")
	if err != nil {
		t.Errorf("Expected perform request correctly, got error: %v", err)
		return
	}
	if res.StatusCode != 404 {
		t.Errorf("Expected to get 404 status, got: %d", res.Status)
	}
}

func TestServerPost(t *testing.T) {
	ioutil.WriteFile("/tmp/test.txt", []byte("hello!"), 0644)
	res, err := http.PostForm("http://localhost:8891/test.txt", url.Values{})
	if err != nil {
		t.Errorf("Expected perform request correctly, got error: %v", err)
		return
	}
	if res.StatusCode != 405 {
		t.Errorf("Expected to get 405 status, got: %d", res.Status)
	}
}
