// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2022-03-31

package client

import (
	"encoding/json"
)

// UpdateClusterSpecification struct for UpdateClusterSpecification.
type UpdateClusterSpecification struct {
	Dedicated            *DedicatedClusterUpdateSpecification  `json:"dedicated,omitempty"`
	Serverless           *ServerlessClusterUpdateSpecification `json:"serverless,omitempty"`
	AdditionalProperties map[string]interface{}
}

type updateClusterSpecification UpdateClusterSpecification

// NewUpdateClusterSpecification instantiates a new UpdateClusterSpecification object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateClusterSpecification() *UpdateClusterSpecification {
	p := UpdateClusterSpecification{}
	return &p
}

// GetDedicated returns the Dedicated field value if set, zero value otherwise.
func (o *UpdateClusterSpecification) GetDedicated() DedicatedClusterUpdateSpecification {
	if o == nil || o.Dedicated == nil {
		var ret DedicatedClusterUpdateSpecification
		return ret
	}
	return *o.Dedicated
}

// SetDedicated gets a reference to the given DedicatedClusterUpdateSpecification and assigns it to the Dedicated field.
func (o *UpdateClusterSpecification) SetDedicated(v DedicatedClusterUpdateSpecification) {
	o.Dedicated = &v
}

// GetServerless returns the Serverless field value if set, zero value otherwise.
func (o *UpdateClusterSpecification) GetServerless() ServerlessClusterUpdateSpecification {
	if o == nil || o.Serverless == nil {
		var ret ServerlessClusterUpdateSpecification
		return ret
	}
	return *o.Serverless
}

// SetServerless gets a reference to the given ServerlessClusterUpdateSpecification and assigns it to the Serverless field.
func (o *UpdateClusterSpecification) SetServerless(v ServerlessClusterUpdateSpecification) {
	o.Serverless = &v
}

func (o UpdateClusterSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dedicated != nil {
		toSerialize["dedicated"] = o.Dedicated
	}
	if o.Serverless != nil {
		toSerialize["serverless"] = o.Serverless
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateClusterSpecification) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateClusterSpecification := updateClusterSpecification{}

	if err = json.Unmarshal(bytes, &varUpdateClusterSpecification); err == nil {
		*o = UpdateClusterSpecification(varUpdateClusterSpecification)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dedicated")
		delete(additionalProperties, "serverless")
		o.AdditionalProperties = additionalProperties
	}

	return err
}