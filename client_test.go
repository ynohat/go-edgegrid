package edgegrid

import (
	"io/ioutil"
	"testing"
)

func TestApiRequest(t *testing.T) {

	api := NewFromIni(".edgerc")
	resp, err := api.Send("POST", "/ccu/v3/invalidate/url/production", "{ \"objects\" : [ \"http://example.com\" ]}")
	if err != nil {
		t.Error(err)
	}

	if resp.Status != "201" {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		t.Log(string(contents))
	}
}
