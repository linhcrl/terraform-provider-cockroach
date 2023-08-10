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
// API version: 2023-04-10

package client

import (
	"fmt"
)

// SetAWSEndpointConnectionStatusType  - AVAILABLE: accept/verify the connection on the service side.  - REJECTED: reject the connection on the service side.
type SetAWSEndpointConnectionStatusType string

// List of SetAWSEndpointConnectionStatus.Type.
const (
	SETAWSENDPOINTCONNECTIONSTATUSTYPE_AVAILABLE SetAWSEndpointConnectionStatusType = "AVAILABLE"
	SETAWSENDPOINTCONNECTIONSTATUSTYPE_REJECTED  SetAWSEndpointConnectionStatusType = "REJECTED"
)

// All allowed values of SetAWSEndpointConnectionStatusType enum.
var AllowedSetAWSEndpointConnectionStatusTypeEnumValues = []SetAWSEndpointConnectionStatusType{
	"AVAILABLE",
	"REJECTED",
}

// NewSetAWSEndpointConnectionStatusTypeFromValue returns a pointer to a valid SetAWSEndpointConnectionStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSetAWSEndpointConnectionStatusTypeFromValue(v string) (*SetAWSEndpointConnectionStatusType, error) {
	ev := SetAWSEndpointConnectionStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SetAWSEndpointConnectionStatusType: valid values are %v", v, AllowedSetAWSEndpointConnectionStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SetAWSEndpointConnectionStatusType) IsValid() bool {
	for _, existing := range AllowedSetAWSEndpointConnectionStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SetAWSEndpointConnectionStatus.Type value.
func (v SetAWSEndpointConnectionStatusType) Ptr() *SetAWSEndpointConnectionStatusType {
	return &v
}
