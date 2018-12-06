// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "env": Application Media Types
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
	"net/http"
)

// Holds a single environment (default view)
//
// Identifier: application/vnd.environment+json; view=default
type EnvironmentSingle struct {
	Data *Environment `form:"data" json:"data" xml:"data"`
	// An array of mixed types
	Included []interface{} `form:"included,omitempty" json:"included,omitempty" xml:"included,omitempty"`
}

// Validate validates the EnvironmentSingle media type instance.
func (mt *EnvironmentSingle) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	if mt.Data != nil {
		if err2 := mt.Data.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeEnvironmentSingle decodes the EnvironmentSingle instance encoded in resp body.
func (c *Client) DecodeEnvironmentSingle(resp *http.Response) (*EnvironmentSingle, error) {
	var decoded EnvironmentSingle
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Holds the list of environments (default view)
//
// Identifier: application/vnd.environmentslist+json; view=default
type EnvironmentsList struct {
	Data []*Environment `form:"data" json:"data" xml:"data"`
	// An array of mixed types
	Included []interface{}        `form:"included,omitempty" json:"included,omitempty" xml:"included,omitempty"`
	Links    *PagingLinks         `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Meta     *EnvironmentListMeta `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
}

// Validate validates the EnvironmentsList media type instance.
func (mt *EnvironmentsList) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeEnvironmentsList decodes the EnvironmentsList instance encoded in resp body.
func (c *Client) DecodeEnvironmentsList(resp *http.Response) (*EnvironmentsList, error) {
	var decoded EnvironmentsList
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// JSONAPIErrors media type (default view)
//
// Identifier: application/vnd.jsonapierrors+json; view=default
type JSONAPIErrors struct {
	Errors []*JSONAPIError `form:"errors" json:"errors" xml:"errors"`
}

// Validate validates the JSONAPIErrors media type instance.
func (mt *JSONAPIErrors) Validate() (err error) {
	if mt.Errors == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "errors"))
	}
	for _, e := range mt.Errors {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeJSONAPIErrors decodes the JSONAPIErrors instance encoded in resp body.
func (c *Client) DecodeJSONAPIErrors(resp *http.Response) (*JSONAPIErrors, error) {
	var decoded JSONAPIErrors
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// The status of the current running instance (default view)
//
// Identifier: application/vnd.status+json; view=default
type Status struct {
	// The time when built
	BuildTime string `form:"buildTime" json:"buildTime" xml:"buildTime"`
	// Commit SHA this build is based on
	Commit string `form:"commit" json:"commit" xml:"commit"`
	// The status of Database connection. 'OK' or an error message is displayed.
	DatabaseStatus string `form:"databaseStatus" json:"databaseStatus" xml:"databaseStatus"`
	// The time when started
	StartTime string `form:"startTime" json:"startTime" xml:"startTime"`
}

// Validate validates the Status media type instance.
func (mt *Status) Validate() (err error) {
	if mt.Commit == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "commit"))
	}
	if mt.BuildTime == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "buildTime"))
	}
	if mt.StartTime == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "startTime"))
	}
	if mt.DatabaseStatus == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "databaseStatus"))
	}
	return
}

// DecodeStatus decodes the Status instance encoded in resp body.
func (c *Client) DecodeStatus(resp *http.Response) (*Status, error) {
	var decoded Status
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}