package util

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

func DoGet(url string, param map[string]string, resp interface{}) error {
	apiUrl, err := parseUrl(url, param)
	if apiUrl == "" {
		return err
	}

	res, err := http.Get(apiUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(resp)
}

func DoPost(url string, param interface{}, resp interface{}) error {
	var form []byte
	form, _ = json.Marshal(param)
	res, err := http.Post(url, "", bytes.NewReader(form))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(resp)
}

func DoPostXMLWithTLS(uri string, obj interface{}, resp interface{}, ca, key string) error {
	xmlData, err := xml.Marshal(obj)
	if err != nil {
		return err
	}

	client, err := httpWithTLS(ca, key)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(xmlData)
	res, err := client.Post(uri, "application/xml;charset=utf-8", body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("http code error : uri=%v , statusCode=%v", uri, res.StatusCode)
	}

	return xml.NewDecoder(res.Body).Decode(resp)
}

func httpWithTLS(rootCa, key string) (*http.Client, error) {
	var client *http.Client
	cert, err := tls.LoadX509KeyPair(rootCa, key)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	tr := &http.Transport{
		TLSClientConfig:    config,
		DisableCompression: true,
	}
	client = &http.Client{Transport: tr}
	return client, nil
}

func parseUrl(api string, params map[string]string) (string, error) {
	url, err := url.Parse(api)
	if err != nil {
		return "", err
	}

	query := url.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	url.RawQuery = query.Encode()
	return url.String(), nil
}
