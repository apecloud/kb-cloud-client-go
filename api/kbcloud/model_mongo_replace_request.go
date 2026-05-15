// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoReplaceRequest struct {
	// document selector
	Filter interface{} `json:"filter,omitempty"`
	// replacement document
	Replacement interface{} `json:"replacement,omitempty"`
	// create document when no match exists
	Upsert *bool `json:"upsert,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoReplaceRequest instantiates a new MongoReplaceRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoReplaceRequest() *MongoReplaceRequest {
	this := MongoReplaceRequest{}
	return &this
}

// NewMongoReplaceRequestWithDefaults instantiates a new MongoReplaceRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoReplaceRequestWithDefaults() *MongoReplaceRequest {
	this := MongoReplaceRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MongoReplaceRequest) GetFilter() interface{} {
	if o == nil || o.Filter == nil {
		var ret interface{}
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoReplaceRequest) GetFilterOk() (*interface{}, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return &o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MongoReplaceRequest) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given interface{} and assigns it to the Filter field.
func (o *MongoReplaceRequest) SetFilter(v interface{}) {
	o.Filter = v
}

// GetReplacement returns the Replacement field value if set, zero value otherwise.
func (o *MongoReplaceRequest) GetReplacement() interface{} {
	if o == nil || o.Replacement == nil {
		var ret interface{}
		return ret
	}
	return o.Replacement
}

// GetReplacementOk returns a tuple with the Replacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoReplaceRequest) GetReplacementOk() (*interface{}, bool) {
	if o == nil || o.Replacement == nil {
		return nil, false
	}
	return &o.Replacement, true
}

// HasReplacement returns a boolean if a field has been set.
func (o *MongoReplaceRequest) HasReplacement() bool {
	return o != nil && o.Replacement != nil
}

// SetReplacement gets a reference to the given interface{} and assigns it to the Replacement field.
func (o *MongoReplaceRequest) SetReplacement(v interface{}) {
	o.Replacement = v
}

// GetUpsert returns the Upsert field value if set, zero value otherwise.
func (o *MongoReplaceRequest) GetUpsert() bool {
	if o == nil || o.Upsert == nil {
		var ret bool
		return ret
	}
	return *o.Upsert
}

// GetUpsertOk returns a tuple with the Upsert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoReplaceRequest) GetUpsertOk() (*bool, bool) {
	if o == nil || o.Upsert == nil {
		return nil, false
	}
	return o.Upsert, true
}

// HasUpsert returns a boolean if a field has been set.
func (o *MongoReplaceRequest) HasUpsert() bool {
	return o != nil && o.Upsert != nil
}

// SetUpsert gets a reference to the given bool and assigns it to the Upsert field.
func (o *MongoReplaceRequest) SetUpsert(v bool) {
	o.Upsert = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoReplaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Replacement != nil {
		toSerialize["replacement"] = o.Replacement
	}
	if o.Upsert != nil {
		toSerialize["upsert"] = o.Upsert
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoReplaceRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter      interface{} `json:"filter,omitempty"`
		Replacement interface{} `json:"replacement,omitempty"`
		Upsert      *bool       `json:"upsert,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"filter", "replacement", "upsert"})
	} else {
		return err
	}
	o.Filter = all.Filter
	o.Replacement = all.Replacement
	o.Upsert = all.Upsert

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
