// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoDeleteRequest struct {
	// document selector
	Filter interface{} `json:"filter,omitempty"`
	// reserved for future bulk delete; S1 rejects true
	Many *bool `json:"many,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoDeleteRequest instantiates a new MongoDeleteRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoDeleteRequest() *MongoDeleteRequest {
	this := MongoDeleteRequest{}
	return &this
}

// NewMongoDeleteRequestWithDefaults instantiates a new MongoDeleteRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoDeleteRequestWithDefaults() *MongoDeleteRequest {
	this := MongoDeleteRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MongoDeleteRequest) GetFilter() interface{} {
	if o == nil || o.Filter == nil {
		var ret interface{}
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDeleteRequest) GetFilterOk() (*interface{}, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return &o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MongoDeleteRequest) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given interface{} and assigns it to the Filter field.
func (o *MongoDeleteRequest) SetFilter(v interface{}) {
	o.Filter = v
}

// GetMany returns the Many field value if set, zero value otherwise.
func (o *MongoDeleteRequest) GetMany() bool {
	if o == nil || o.Many == nil {
		var ret bool
		return ret
	}
	return *o.Many
}

// GetManyOk returns a tuple with the Many field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDeleteRequest) GetManyOk() (*bool, bool) {
	if o == nil || o.Many == nil {
		return nil, false
	}
	return o.Many, true
}

// HasMany returns a boolean if a field has been set.
func (o *MongoDeleteRequest) HasMany() bool {
	return o != nil && o.Many != nil
}

// SetMany gets a reference to the given bool and assigns it to the Many field.
func (o *MongoDeleteRequest) SetMany(v bool) {
	o.Many = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Many != nil {
		toSerialize["many"] = o.Many
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoDeleteRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter interface{} `json:"filter,omitempty"`
		Many   *bool       `json:"many,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"filter", "many"})
	} else {
		return err
	}
	o.Filter = all.Filter
	o.Many = all.Many

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
