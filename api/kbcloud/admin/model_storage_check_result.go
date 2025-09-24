// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StorageCheckResult storageCheck is the result when checking storage connectivity
type StorageCheckResult struct {
	// the result that the storage connectivity
	Connected *bool `json:"connected,omitempty"`
	// the error msg
	Error *string `json:"error,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageCheckResult instantiates a new StorageCheckResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageCheckResult() *StorageCheckResult {
	this := StorageCheckResult{}
	return &this
}

// NewStorageCheckResultWithDefaults instantiates a new StorageCheckResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageCheckResultWithDefaults() *StorageCheckResult {
	this := StorageCheckResult{}
	return &this
}

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *StorageCheckResult) GetConnected() bool {
	if o == nil || o.Connected == nil {
		var ret bool
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageCheckResult) GetConnectedOk() (*bool, bool) {
	if o == nil || o.Connected == nil {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *StorageCheckResult) HasConnected() bool {
	return o != nil && o.Connected != nil
}

// SetConnected gets a reference to the given bool and assigns it to the Connected field.
func (o *StorageCheckResult) SetConnected(v bool) {
	o.Connected = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *StorageCheckResult) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageCheckResult) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *StorageCheckResult) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *StorageCheckResult) SetError(v string) {
	o.Error = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageCheckResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Connected != nil {
		toSerialize["connected"] = o.Connected
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageCheckResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Connected *bool   `json:"connected,omitempty"`
		Error     *string `json:"error,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"connected", "error"})
	} else {
		return err
	}
	o.Connected = all.Connected
	o.Error = all.Error

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
