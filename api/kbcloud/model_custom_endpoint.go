// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type CustomEndpoint struct {
	ConnectionString *string `json:"connectionString,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomEndpoint instantiates a new CustomEndpoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomEndpoint() *CustomEndpoint {
	this := CustomEndpoint{}
	return &this
}

// NewCustomEndpointWithDefaults instantiates a new CustomEndpoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomEndpointWithDefaults() *CustomEndpoint {
	this := CustomEndpoint{}
	return &this
}

// GetConnectionString returns the ConnectionString field value if set, zero value otherwise.
func (o *CustomEndpoint) GetConnectionString() string {
	if o == nil || o.ConnectionString == nil {
		var ret string
		return ret
	}
	return *o.ConnectionString
}

// GetConnectionStringOk returns a tuple with the ConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEndpoint) GetConnectionStringOk() (*string, bool) {
	if o == nil || o.ConnectionString == nil {
		return nil, false
	}
	return o.ConnectionString, true
}

// HasConnectionString returns a boolean if a field has been set.
func (o *CustomEndpoint) HasConnectionString() bool {
	return o != nil && o.ConnectionString != nil
}

// SetConnectionString gets a reference to the given string and assigns it to the ConnectionString field.
func (o *CustomEndpoint) SetConnectionString(v string) {
	o.ConnectionString = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ConnectionString != nil {
		toSerialize["connectionString"] = o.ConnectionString
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConnectionString *string `json:"connectionString,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"connectionString"})
	} else {
		return err
	}
	o.ConnectionString = all.ConnectionString

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
