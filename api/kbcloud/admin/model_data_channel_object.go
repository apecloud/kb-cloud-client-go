// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataChannelObject struct {
	IncludeObjects []string                    `json:"includeObjects,omitempty"`
	ExcludeObjects common.NullableList[string] `json:"excludeObjects,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataChannelObject instantiates a new DataChannelObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataChannelObject() *DataChannelObject {
	this := DataChannelObject{}
	return &this
}

// NewDataChannelObjectWithDefaults instantiates a new DataChannelObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataChannelObjectWithDefaults() *DataChannelObject {
	this := DataChannelObject{}
	return &this
}

// GetIncludeObjects returns the IncludeObjects field value if set, zero value otherwise.
func (o *DataChannelObject) GetIncludeObjects() []string {
	if o == nil || o.IncludeObjects == nil {
		var ret []string
		return ret
	}
	return o.IncludeObjects
}

// GetIncludeObjectsOk returns a tuple with the IncludeObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelObject) GetIncludeObjectsOk() (*[]string, bool) {
	if o == nil || o.IncludeObjects == nil {
		return nil, false
	}
	return &o.IncludeObjects, true
}

// HasIncludeObjects returns a boolean if a field has been set.
func (o *DataChannelObject) HasIncludeObjects() bool {
	return o != nil && o.IncludeObjects != nil
}

// SetIncludeObjects gets a reference to the given []string and assigns it to the IncludeObjects field.
func (o *DataChannelObject) SetIncludeObjects(v []string) {
	o.IncludeObjects = v
}

// GetExcludeObjects returns the ExcludeObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataChannelObject) GetExcludeObjects() []string {
	if o == nil || o.ExcludeObjects.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ExcludeObjects.Get()
}

// GetExcludeObjectsOk returns a tuple with the ExcludeObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataChannelObject) GetExcludeObjectsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExcludeObjects.Get(), o.ExcludeObjects.IsSet()
}

// HasExcludeObjects returns a boolean if a field has been set.
func (o *DataChannelObject) HasExcludeObjects() bool {
	return o != nil && o.ExcludeObjects.IsSet()
}

// SetExcludeObjects gets a reference to the given common.NullableList[string] and assigns it to the ExcludeObjects field.
func (o *DataChannelObject) SetExcludeObjects(v []string) {
	o.ExcludeObjects.Set(&v)
}

// SetExcludeObjectsNil sets the value for ExcludeObjects to be an explicit nil.
func (o *DataChannelObject) SetExcludeObjectsNil() {
	o.ExcludeObjects.Set(nil)
}

// UnsetExcludeObjects ensures that no value is present for ExcludeObjects, not even an explicit nil.
func (o *DataChannelObject) UnsetExcludeObjects() {
	o.ExcludeObjects.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DataChannelObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.IncludeObjects != nil {
		toSerialize["includeObjects"] = o.IncludeObjects
	}
	if o.ExcludeObjects.IsSet() {
		toSerialize["excludeObjects"] = o.ExcludeObjects.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataChannelObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeObjects []string                    `json:"includeObjects,omitempty"`
		ExcludeObjects common.NullableList[string] `json:"excludeObjects,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"includeObjects", "excludeObjects"})
	} else {
		return err
	}
	o.IncludeObjects = all.IncludeObjects
	o.ExcludeObjects = all.ExcludeObjects

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
