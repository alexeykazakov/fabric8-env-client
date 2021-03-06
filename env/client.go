// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "env": Client
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-env/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-env-client
// --pkg=env
// --version=v1.3.0

package env

import (
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
)

// Client is the env service client.
type Client struct {
	*goaclient.Client
	JWTSigner goaclient.Signer
	Encoder   *goa.HTTPEncoder
	Decoder   *goa.HTTPDecoder
}

// New instantiates the client.
func New(c goaclient.Doer) *Client {
	client := &Client{
		Client:  goaclient.New(c),
		Encoder: goa.NewHTTPEncoder(),
		Decoder: goa.NewHTTPDecoder(),
	}

	// Setup encoders and decoders
	client.Encoder.Register(goa.NewJSONEncoder, "application/json")
	client.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	client.Encoder.Register(goa.NewJSONEncoder, "*/*")
	client.Decoder.Register(goa.NewJSONDecoder, "*/*")

	return client
}

// SetJWTSigner sets the request signer for the jwt security scheme.
func (c *Client) SetJWTSigner(signer goaclient.Signer) {
	c.JWTSigner = signer
}
