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

// EnableLogExportRequest struct for EnableLogExportRequest.
type EnableLogExportRequest struct {
	// auth_principal is either the AWS Role ARN that identifies a role that the cluster account can assume to write to CloudWatch or the GCP Project ID that the cluster service account has permissions to write to for cloud logging.
	AuthPrincipal string `json:"auth_principal"`
	// groups is a collection of log group configurations that allows the customer to define collections of CRDB log channels that are aggregated separately at the target sink.
	Groups *[]LogExportGroup `json:"groups,omitempty"`
	// log_name is an identifier for the logs in the customer's log sink.
	LogName string `json:"log_name"`
	// redact allows the customer to set a default redaction policy for logs before they are exported to the target sink. If a group config omits a redact flag and this one is set to `true`, then that group will receive redacted logs.
	Redact *bool `json:"redact,omitempty"`
	// region allows the customer to override the destination region for all logs for a cluster.
	Region *string       `json:"region,omitempty"`
	Type   LogExportType `json:"type"`
}

// NewEnableLogExportRequest instantiates a new EnableLogExportRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableLogExportRequest(authPrincipal string, logName string, type_ LogExportType) *EnableLogExportRequest {
	p := EnableLogExportRequest{}
	p.AuthPrincipal = authPrincipal
	p.LogName = logName
	p.Type = type_
	return &p
}

// NewEnableLogExportRequestWithDefaults instantiates a new EnableLogExportRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableLogExportRequestWithDefaults() *EnableLogExportRequest {
	p := EnableLogExportRequest{}
	return &p
}

// GetAuthPrincipal returns the AuthPrincipal field value.
func (o *EnableLogExportRequest) GetAuthPrincipal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthPrincipal
}

// SetAuthPrincipal sets field value.
func (o *EnableLogExportRequest) SetAuthPrincipal(v string) {
	o.AuthPrincipal = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *EnableLogExportRequest) GetGroups() []LogExportGroup {
	if o == nil || o.Groups == nil {
		var ret []LogExportGroup
		return ret
	}
	return *o.Groups
}

// SetGroups gets a reference to the given []LogExportGroup and assigns it to the Groups field.
func (o *EnableLogExportRequest) SetGroups(v []LogExportGroup) {
	o.Groups = &v
}

// GetLogName returns the LogName field value.
func (o *EnableLogExportRequest) GetLogName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogName
}

// SetLogName sets field value.
func (o *EnableLogExportRequest) SetLogName(v string) {
	o.LogName = v
}

// GetRedact returns the Redact field value if set, zero value otherwise.
func (o *EnableLogExportRequest) GetRedact() bool {
	if o == nil || o.Redact == nil {
		var ret bool
		return ret
	}
	return *o.Redact
}

// SetRedact gets a reference to the given bool and assigns it to the Redact field.
func (o *EnableLogExportRequest) SetRedact(v bool) {
	o.Redact = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *EnableLogExportRequest) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *EnableLogExportRequest) SetRegion(v string) {
	o.Region = &v
}

// GetType returns the Type field value.
func (o *EnableLogExportRequest) GetType() LogExportType {
	if o == nil {
		var ret LogExportType
		return ret
	}

	return o.Type
}

// SetType sets field value.
func (o *EnableLogExportRequest) SetType(v LogExportType) {
	o.Type = v
}

func (o EnableLogExportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["auth_principal"] = o.AuthPrincipal
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if true {
		toSerialize["log_name"] = o.LogName
	}
	if o.Redact != nil {
		toSerialize["redact"] = o.Redact
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}