// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2022-03-31

package client

import (
	"encoding/json"
	"time"
)

// KeysetPaginationResponse struct for KeysetPaginationResponse.
type KeysetPaginationResponse struct {
	Next                 *string    `json:"next,omitempty"`
	Last                 *string    `json:"last,omitempty"`
	Limit                *int32     `json:"limit,omitempty"`
	Time                 *time.Time `json:"time,omitempty"`
	Order                *SortOrder `json:"order,omitempty"`
	AdditionalProperties map[string]interface{}
}

type keysetPaginationResponse KeysetPaginationResponse

// NewKeysetPaginationResponse instantiates a new KeysetPaginationResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeysetPaginationResponse() *KeysetPaginationResponse {
	p := KeysetPaginationResponse{}
	var order SortOrder = SORTORDER_ASC
	p.Order = &order
	return &p
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *KeysetPaginationResponse) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *KeysetPaginationResponse) SetNext(v string) {
	o.Next = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *KeysetPaginationResponse) GetLast() string {
	if o == nil || o.Last == nil {
		var ret string
		return ret
	}
	return *o.Last
}

// SetLast gets a reference to the given string and assigns it to the Last field.
func (o *KeysetPaginationResponse) SetLast(v string) {
	o.Last = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *KeysetPaginationResponse) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *KeysetPaginationResponse) SetLimit(v int32) {
	o.Limit = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *KeysetPaginationResponse) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *KeysetPaginationResponse) SetTime(v time.Time) {
	o.Time = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *KeysetPaginationResponse) GetOrder() SortOrder {
	if o == nil || o.Order == nil {
		var ret SortOrder
		return ret
	}
	return *o.Order
}

// SetOrder gets a reference to the given SortOrder and assigns it to the Order field.
func (o *KeysetPaginationResponse) SetOrder(v SortOrder) {
	o.Order = &v
}

func (o KeysetPaginationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Last != nil {
		toSerialize["last"] = o.Last
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KeysetPaginationResponse) UnmarshalJSON(bytes []byte) (err error) {
	varKeysetPaginationResponse := keysetPaginationResponse{}

	if err = json.Unmarshal(bytes, &varKeysetPaginationResponse); err == nil {
		*o = KeysetPaginationResponse(varKeysetPaginationResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "next")
		delete(additionalProperties, "last")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "time")
		delete(additionalProperties, "order")
		o.AdditionalProperties = additionalProperties
	}

	return err
}