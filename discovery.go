package discovery

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// kv stands for KEY VALUE
const (
	kvStorage = "http://localhost:8500"
	kvPath    = "/v1/kv"
)

type client struct {
	key    string
	data   []byte
	errors []error
}

type consulKVResponse struct {
	LockIndex   int
	Key         string
	Flags       int
	Value       string
	CreateIndex int
	ModifyIndex int
}

/*
New - discovery client.
key - key ins KV storage. Example: "/v1/queue1"
v — pointer to data interface (as in json.Unmarshal)
*/
func New(key string, v interface{}) error {
	if v == nil {
		return errors.New("Data container is nil")
	}
	kv := client{key: key}
	return kv.fetch(v)
}

func (c *client) fetch(v interface{}) (err error) {
	data, err := c.get()
	if err != nil {
		return
	}
	var kvResponse []consulKVResponse
	if err = json.Unmarshal(data, &kvResponse); err != nil {
		return
	}
	var encStr []byte
	if encStr, err = base64.StdEncoding.DecodeString(kvResponse[0].Value); err != nil {
		return
	}
	return json.Unmarshal(encStr, v)
}

func (c *client) get() (body []byte, err error) {
	res, err := http.Get(c.prepareURL())
	if err != nil {
		return
	}
	body, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	return
}

func (c *client) prepareURL() string {
	url := kvStorage
	env := os.Getenv("AMAIZ_CONSUL_URL")
	if env != "" {
		fmt.Println("Using custom Consul url: ", env)
		url = env
	}
	return url + kvPath + c.key
}
