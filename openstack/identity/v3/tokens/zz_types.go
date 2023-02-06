package tokens

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/mdbooth/openstack-go/meta"
)

type RequestAuthIdentityMethods string

const (
	RequestAuthIdentityMethodsApplicationCredential = "application_credential"
)

type Request struct {
	Auth RequestAuth `json:"auth"`
}

func (r *Request) Path() string {
	return "/v3/auth/tokens"
}

func (r *Request) Method() string {
	return "POST"
}

func (r *Request) AuthenticateUser(ctx context.Context, requester meta.Requester) (*Response, error) {
	var errs []error
	jsonBytes, err := requester.Request(ctx, r)
	if err != nil {
		errs = append(errs, err)
	}

	var response *Response
	if jsonBytes != nil {
		response = &Response{}
		err = json.Unmarshal(jsonBytes, response)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return response, errors.Join(errs...)
}

type RequestUser struct {
	ID string `json:"id"`
}

type RequestAuth struct {
	Identity RequestAuthIdentity `json:"identity"`
}

type RequestAuthIdentity struct {
	Methods               []RequestAuthIdentityMethods              `json:"methods"`
	ApplicationCredential *RequestAuthIdentityApplicationCredential `json:"application_credential,omitempty"`
}

// ApplicationCredential is oneOf...
// N.B. `byID` and `byName` must be obtained from custom metadata

type RequestAuthIdentityApplicationCredential struct {
	ID     interface{} `json:"id,omitempty"`
	Name   interface{} `json:"name,omitempty"`
	Secret interface{} `json:"secret"`
	User   interface{} `json:"user,omitempty"`
}

type RequestAuthIdentityApplicationCredentialByID struct {
	ID     string `json:"id,omitempty"`
	Secret string `json:"secret"`
}

func (r RequestAuthIdentityApplicationCredentialByID) AsRequestAuthIdentityApplicationCredential() *RequestAuthIdentityApplicationCredential {
	return &RequestAuthIdentityApplicationCredential{
		ID:     r.ID,
		Secret: r.Secret,
	}
}

type RequestAuthIdentityApplicationCredentialByName struct {
	Name   string       `json:"name,omitempty"`
	Secret string       `json:"secret"`
	User   *RequestUser `json:"user,omitempty"`
}

func (r RequestAuthIdentityApplicationCredentialByName) AsRequestAuthIdentityApplicationCredential() *RequestAuthIdentityApplicationCredential {
	return &RequestAuthIdentityApplicationCredential{
		Name:   r.Name,
		Secret: r.Secret,
		User:   r.User,
	}
}

type Response struct {
}
