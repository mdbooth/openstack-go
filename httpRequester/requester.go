package httpRequester

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/mdbooth/openstack-go/meta"
)

type Requester struct {
	endpoint string
	client   http.Client
}

func NewRequester(endpoint string) *Requester {
	return &Requester{
		endpoint: endpoint,
		client:   http.Client{},
	}
}

func (r *Requester) Request(ctx context.Context, req meta.Request) ([]byte, error) {
	requestURL, err := url.JoinPath(r.endpoint, req.Path())
	if err != nil {
		return nil, err
	}

	requestBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequestWithContext(ctx, req.Method(), requestURL, bytes.NewReader(requestBytes))
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	return body, nil
}
