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

type KafkaRestClient struct {
	Client http.Client
	Bearer string
}

func New(conf config.Config) *KafkaRestClient {
	var client *http.Client
	tls := conf.Credentials.Certificates != config.Certificates{}
	if tls {
		client = &http.Client{
			Transport: getTransport(conf.Credentials.Certificates),
		}
	} else {
		client = &http.Client{}
	}

	user := conf.Credentials.Key + ":" + conf.Credentials.Secret
	bearer := b64.StdEncoding.EncodeToString([]byte(user))

	return &KafkaRestClient{
		Client: *client,
		Bearer: bearer,
	}
}

// POST request
func (kClient *KafkaRestClient) Post(requestURL string, requestBody []byte) ([]interface{}, error) {

	req, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Printf("Kafka Rest client: could not create request: %s\n", err)
		return nil, err
	}
	resp, err := kClient.buildArrayRequest(req)
	if err != nil {
		log.Printf("Kafka Rest client: POST %s response error: %s\n", requestURL, err)
		return nil, err
	}

	return resp, nil
}

func (kClient *KafkaRestClient) Get(requestURL string) (interface{}, error) {
	log.Printf("Building request %s", requestURL)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return nil, err
	}
	resp, err := kClient.buildRequest(req)
	if err != nil {
		log.Printf("Kafka Rest client: GET %s response error: %s\n", requestURL, err)
		return nil, err
	}
	return resp, nil
}

// Get Request
// Expect results --> data:[]
func (kClient *KafkaRestClient) GetList(requestURL string) ([]interface{}, error) {
	log.Printf("Building request %s", requestURL)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return nil, err
	}
	resp, err := kClient.buildArrayRequest(req)
	if err != nil {
		log.Printf("Kafka Rest client: GET %s response error: %s\n", requestURL, err)
		return nil, err
	}

	return resp, nil
}

func (kClient *KafkaRestClient) buildArrayRequest(req *http.Request) ([]interface{}, error) {
	result, err := kClient.build(req)
	if err != nil {
		fmt.Printf("client: could not get result: %s\n", err)
		return nil, err
	}

	if result["data"] != nil {
		return result["data"].([]interface{}), nil
	}
	return nil, errors.New("No data result")
}

func (kClient *KafkaRestClient) buildRequest(req *http.Request) (interface{}, error) {
	return kClient.build(req)
}

// Build request - Client Do
func (kClient *KafkaRestClient) build(req *http.Request) (map[string]any, error) {
	req.Header.Set("Authorization", "Basic "+kClient.Bearer)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	res, err := kClient.Client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return nil, err
	}
	if res.StatusCode == http.StatusNotFound {
		fmt.Printf("client: 404 - Not found %v \n", err)
		return nil, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return nil, err
	}

	var result map[string]any
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
		fmt.Printf("error loading cert files")
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		fmt.Printf("Error reading the CA cert")
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
