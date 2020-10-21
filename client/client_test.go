package client

import "testing"

func TestApiClientSuccess(t *testing.T) {
	c := ApiClient{}
	t.Log("?", c)

	if false {
		t.Fatal("hello")
	}
}