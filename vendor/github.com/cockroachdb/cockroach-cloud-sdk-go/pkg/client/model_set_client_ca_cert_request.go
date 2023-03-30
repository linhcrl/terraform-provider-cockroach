// Copyright 2023 The Cockroach Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2022-09-20

package client

import (
	"encoding/json"
)

// SetClientCACertRequest struct for SetClientCACertRequest.
type SetClientCACertRequest struct {
	X509PemCert string `json:"x509_pem_cert"`
}

// NewSetClientCACertRequest instantiates a new SetClientCACertRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetClientCACertRequest(x509PemCert string) *SetClientCACertRequest {
	p := SetClientCACertRequest{}
	p.X509PemCert = x509PemCert
	return &p
}

// NewSetClientCACertRequestWithDefaults instantiates a new SetClientCACertRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetClientCACertRequestWithDefaults() *SetClientCACertRequest {
	p := SetClientCACertRequest{}
	return &p
}

// GetX509PemCert returns the X509PemCert field value.
func (o *SetClientCACertRequest) GetX509PemCert() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.X509PemCert
}

// SetX509PemCert sets field value.
func (o *SetClientCACertRequest) SetX509PemCert(v string) {
	o.X509PemCert = v
}

func (o SetClientCACertRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["x509_pem_cert"] = o.X509PemCert
	}
	return json.Marshal(toSerialize)
}