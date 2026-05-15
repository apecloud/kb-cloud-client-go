// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoUpdateRequest struct {
	// document selector
	Filter interface{} `json:"filter,omitempty"`
	// update document
	Update interface{} `json:"update,omitempty"`
	// create document when no match exists
	Upsert *bool `json:"upsert,omitempty"`
	// reserved for future bulk update; S1 rejects true
	Many *bool `json:"many,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoUpdateRequest instantiates a new MongoUpdateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoUpdateRequest() *MongoUpdateRequest {
	this := MongoUpdateRequest{}
	return &this
}

// NewMongoUpdateRequestWithDefaults instantiates a new MongoUpdateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoUpdateRequestWithDefaults() *MongoUpdateRequest {
	this := MongoUpdateRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MongoUpdateRequest) GetFilter() interface{} {
	if o == nil || o.Filter == nil {
		var ret interface{}
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoUpdateRequest) GetFilterOk() (*interface{}, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return &o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MongoUpdateRequest) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given interface{} and assigns it to the Filter field.
func (o *MongoUpdateRequest) SetFilter(v interface{}) {
	o.Filter = v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *MongoUpdateRequest) GetUpdate() interface{} {
	if o == nil || o.Update == nil {
		var ret interface{}
		return ret
	}
	return o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoUpdateRequest) GetUpdateOk() (*interface{}, bool) {
	if o == nil || o.Update == nil {
		return nil, false
	}
	return &o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *MongoUpdateRequest) HasUpdate() bool {
	return o != nil && o.Update != nil
}

// SetUpdate gets a reference to the given interface{} and assigns it to the Update field.
func (o *MongoUpdateRequest) SetUpdate(v interface{}) {
	o.Update = v
}

// GetUpsert returns the Upsert field value if set, zero value otherwise.
func (o *MongoUpdateRequest) GetUpsert() bool {
	if o == nil || o.Upsert == nil {
		var ret bool
		return ret
	}
	return *o.Upsert
}

// GetUpsertOk returns a tuple with the Upsert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoUpdateRequest) GetUpsertOk() (*bool, bool) {
	if o == nil || o.Upsert == nil {
		return nil, false
	}
	return o.Upsert, true
}

// HasUpsert returns a boolean if a field has been set.
func (o *MongoUpdateRequest) HasUpsert() bool {
	return o != nil && o.Upsert != nil
}

// SetUpsert gets a reference to the given bool and assigns it to the Upsert field.
func (o *MongoUpdateRequest) SetUpsert(v bool) {
	o.Upsert = &v
}

// GetMany returns the Many field value if set, zero value otherwise.
func (o *MongoUpdateRequest) GetMany() bool {
	if o == nil || o.Many == nil {
		var ret bool
		return ret
	}
	return *o.Many
}

// GetManyOk returns a tuple with the Many field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoUpdateRequest) GetManyOk() (*bool, bool) {
	if o == nil || o.Many == nil {
		return nil, false
	}
	return o.Many, true
}

// HasMany returns a boolean if a field has been set.
func (o *MongoUpdateRequest) HasMany() bool {
	return o != nil && o.Many != nil
}

// SetMany gets a reference to the given bool and assigns it to the Many field.
func (o *MongoUpdateRequest) SetMany(v bool) {
	o.Many = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Update != nil {
		toSerialize["update"] = o.Update
	}
	if o.Upsert != nil {
		toSerialize["upsert"] = o.Upsert
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
func (o *MongoUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter interface{} `json:"filter,omitempty"`
		Update interface{} `json:"update,omitempty"`
		Upsert *bool       `json:"upsert,omitempty"`
		Many   *bool       `json:"many,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"filter", "update", "upsert", "many"})
	} else {
		return err
	}
	o.Filter = all.Filter
	o.Update = all.Update
	o.Upsert = all.Upsert
	o.Many = all.Many

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
