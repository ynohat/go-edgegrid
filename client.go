package open

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/satori/go.uuid"
	"gopkg.in/ini.v1"
)

// EdgeGrid Api interface to Akamai
type EdgeGrid struct {
	host         string
	clientToken  string
	clientSecret string
	accessToken  string
	maxBody      string
}

// New returns a new instance of Edgecast
func New(host, clientToken, clientSecret, accessToken, maxBody string) *EdgeGrid {

	e := new(EdgeGrid)

	e.host = host
	e.clientToken = clientToken
	e.clientSecret = clientSecret
	e.accessToken = accessToken
	e.maxBody = maxBody

	return e
}

// NewFromIni returns a new instance of Edgecast
func NewFromIni(iniFile string) *EdgeGrid {

	cfg, err := ini.Load(iniFile)
	if err != nil {
		panic(err)
	}

	section, err := cfg.GetSection("default")
	if err != nil {
		panic(err)
	}

	host, err := section.GetKey("host")
	if err != nil {
		panic("Host name not found in config file")
	}

	clientToken, err := section.GetKey("client_token")
	if err != nil {
		panic("Client Token not found in config file")
	}

	clientSecret, err := section.GetKey("client_secret")
	if err != nil {
		panic("Client Secret not found in config file")
	}

	accessToken, err := section.GetKey("access_token")
	if err != nil {
		panic("Acces Token not found in config file")
	}

	maxBody, err := section.GetKey("max_body")
	if err != nil {
		panic("Max Body not found in config file")
	}

	return New(host.String(), clientToken.String(), clientSecret.String(), accessToken.String(), maxBody.String())
}

// Send a api request to Akamai
func (e *EdgeGrid) Send(method, path, body string) (*http.Response, error) {
	client := &http.Client{}
	method = strings.ToUpper(method)

	if method == "POST" && (string(body) == "{}" || string(body) == "") {
		return nil, errors.New("found empty data set")
	}

	r, _ := http.NewRequest(method, e.host+path, bytes.NewBufferString(body))
	r.Header.Add("Content-Type", "application/json")
	//r.Header.Add("Content-Length", strconv.Itoa(len(body)))
	r.Header.Add("Authorization", e.signedRequest(method, path, r.Header, body))
	return client.Do(r)
}

func (e *EdgeGrid) signedRequest(method, path string, headers http.Header, body string) string {

	t := time.Now().UTC()

	joinedPairs := map[string]string{
		"client_token": e.clientToken,
		"access_token": e.accessToken,
		"timestamp":    t.Format("20060102T15:04:05-0700"),
		"nonce":        uuid.NewV4().String(),
	}

	authHeader := "EG1-HMAC-SHA256 " + formatMap(sorting(joinedPairs), "=", ";")
	authHeader += "signature=" + ComputeHmac256(dataToSign(method, e.host+path, authHeader, headers, body), ComputeHmac256(joinedPairs["timestamp"], e.clientSecret))

	return authHeader
}

func dataToSign(method, requestURL string, authHeader string, headers http.Header, body string) string {

	parsedURL, err := url.Parse(requestURL)
	if err != nil {
		return ""
	}

	parsedHeaders := formatHeaders(headers)
	parsedHeadersString := formatMap(parsedHeaders, ":", "\t")

	dataToSign := []string{
		method,
		parsedURL.Scheme,
		parsedURL.Host,
		parsedURL.Path,
		strings.TrimSpace(parsedHeadersString[:len(parsedHeadersString)-1]),
		Compute256(body),
		authHeader,
	}

	var returnString string
	for _, k := range dataToSign {
		returnString += k + "\\t"
	}

	log.Println(returnString)

	return returnString
}

func formatHeaders(headers http.Header) map[string]string {
	parsedHeaders := map[string]string{}

	for i, v := range headers {
		parsedHeaders[strings.TrimSpace(strings.ToLower(i))] = strings.TrimSpace(v[0])
	}

	return parsedHeaders
}

func formatMap(m map[string]string, sep1, sep2 string) string {
	var r string
	for k, v := range m {
		r += k + sep1 + v + sep2
	}
	return r
}

func sorting(m map[string]string) map[string]string {

	placeholder := make(map[string]string, len(m))

	placeholder["client_token"] = m["client_token"]
	placeholder["access_token"] = m["access_token"]
	placeholder["timestamp"] = m["timestamp"]
	placeholder["nonce"] = m["nonce"]

	return placeholder
}
