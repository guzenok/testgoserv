package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMyHandler(t *testing.T) {
	os.Setenv("INSTANCE_NAME", "app")
	handler := &MyHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()
	for _, i := range []int{1, 2} {
		resp, err := http.Get(server.URL)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != 200 {
			t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
		}
		expected := fmt.Sprintf("%s get %d-nd response", "app", i)
		actual, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		result := string(actual)
		if expected != result {
			t.Errorf("Expected the message '%s', get '%s'\n", expected, result)
		}
	}
}
