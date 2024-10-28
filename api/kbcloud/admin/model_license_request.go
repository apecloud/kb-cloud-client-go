// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// NODESCRIPTION LicenseRequest
type LicenseRequest struct {
	// License key
	Key *string `json:"key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLicenseRequest instantiates a new LicenseRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLicenseRequest() *LicenseRequest {
	this := LicenseRequest{}
	return &this
}

// NewLicenseRequestWithDefaults instantiates a new LicenseRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLicenseRequestWithDefaults() *LicenseRequest {
	this := LicenseRequest{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *LicenseRequest) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseRequest) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *LicenseRequest) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *LicenseRequest) SetKey(v string) {
	o.Key = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LicenseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LicenseRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key *string `json:"key,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"key"})
	} else {
		return err
	}
	o.Key = all.Key

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
