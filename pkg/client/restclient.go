package client

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mcolomerc/cc-tools/pkg/config"
	"net/http"
)

type RestClient struct {
	Client http.Client
	Bearer string
}

func NewRestClient(url string, credentials config.Credentials) *RestClient {
	var client *http.Client
	tls := credentials.Certificates != config.Certificates{}
	if tls {
		client = &http.Client{
			Transport: getTransport(credentials.Certificates),
		}
	} else {
		client = &http.Client{}
	}

	user := credentials.Key + ":" + credentials.Secret
	bearer := b64.StdEncoding.EncodeToString([]byte(user))

	return &RestClient{
		Client: *client,
		Bearer: bearer,
	}
}

// POST request
func (kClient *RestClient) Post(requestURL string, requestBody []byte) ([]interface{}, error) {

	req, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	resp, err := kClient.buildArrayRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (kClient *RestClient) Get(requestURL string) (interface{}, error) {
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := kClient.buildRequest(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Get Request
// Expect results --> data:[]
func (kClient *RestClient) GetList(requestURL string) ([]interface{}, error) {
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := kClient.buildArrayRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (kClient *RestClient) buildArrayRequest(req *http.Request) ([]interface{}, error) {
	result, err := kClient.build(req)
	if err != nil {
		return nil, err
	}
	switch v := result.(type) {
	case map[string]interface{}:
		if v["data"] != nil {
			return v["data"].([]interface{}), nil
		}
	default:
		return result.([]interface{}), nil
	}

	return nil, errors.New("No data result")
}

func (kClient *RestClient) buildRequest(req *http.Request) (interface{}, error) {
	return kClient.build(req)
}

// Build request - Client Do
func (kClient *RestClient) build(req *http.Request) (interface{}, error) {
	req.Header.Set("Authorization", "Basic "+kClient.Bearer)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	res, err := kClient.Client.Do(req)
	if err != nil {
		log.Printf("Rest client: error making http request: %s\n", err)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		errorString := fmt.Sprintf("Rest client:: %d - %s : %v \n", res.StatusCode, req.Method, req.URL)
		return nil, errors.New(errorString)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Rest client:: could not read response body: %s\n", err)
		return nil, err
	}

	var result interface{}
	json.Unmarshal([]byte(resBody), &result)
	return result, nil
}

// Get Transport from certificates
func getTransport(certificates config.Certificates) *http.Transport {
	certFile := certificates.CertFile
	keyFile := certificates.KeyFile
	caFile := certificates.CAFile

	// Load client cert
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Printf("Error loading cert files")
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Printf("Error reading the CA cert")
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	return &http.Transport{TLSClientConfig: tlsConfig}

}
