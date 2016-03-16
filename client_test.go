package edgegrid

import (
	"io/ioutil"
	"testing"

	"strings"
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

func TestSignedRequest(t *testing.T) {
	e := new(EdgeGrid)
	e.clientToken = "test"
	e.clientSecret = "testSecret"
	e.accessToken = "testAccess"
	signedRequest := e.signedRequest("POST", "/somwhere", nil, "")

	if strings.Index(signedRequest, "EG1-HMAC-SHA256") != 0 {
		t.Error("Bad signed request: missing hash declaration: ", signedRequest)
	}

	if !strings.Contains(signedRequest, "client_token="+e.clientToken+";") {
		t.Error("Bad signed request: missing client token: ", signedRequest)
	}
}
