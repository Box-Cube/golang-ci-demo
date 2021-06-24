package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"
)


func TestStartServer(t *testing.T) {
	go startServer()

	r, err := http.Get("http://localhost:80")
	if err != nil {
		t.Fatalf("failed to request http://localhost:80, err: %s", err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		t.Fatalf("failed to reader response body: %v", err)
	}

	t.Logf("now body is %v", string(body))

	if r.StatusCode == http.StatusOK {
		t.Log("server is running...")
		if match,_ := regexp.MatchString("(.*)Golang(.*)", string(body)); match {
			t.Logf("Server response is corrcet...")
		} else {
			t.Fatalf("Server response is not corrcet...")
		}
	} else {
		t.Fatalf("http statusCode not 200")
	}
}