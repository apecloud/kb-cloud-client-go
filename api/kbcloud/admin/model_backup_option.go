// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BackupOption If present, it must be set defaultMethod and fullMethod
type BackupOption struct {
	DefaultMethod    string               `json:"defaultMethod"`
	FullMethod       []BackupMethodOption `json:"fullMethod"`
	ContinuousMethod []BackupMethodOption `json:"continuousMethod,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupOption instantiates a new BackupOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupOption(defaultMethod string, fullMethod []BackupMethodOption) *BackupOption {
	this := BackupOption{}
	this.DefaultMethod = defaultMethod
	this.FullMethod = fullMethod
	return &this
}

// NewBackupOptionWithDefaults instantiates a new BackupOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupOptionWithDefaults() *BackupOption {
	this := BackupOption{}
	return &this
}

// GetDefaultMethod returns the DefaultMethod field value.
func (o *BackupOption) GetDefaultMethod() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DefaultMethod
}

// GetDefaultMethodOk returns a tuple with the DefaultMethod field value
// and a boolean to check if the value has been set.
func (o *BackupOption) GetDefaultMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultMethod, true
}

// SetDefaultMethod sets field value.
func (o *BackupOption) SetDefaultMethod(v string) {
	o.DefaultMethod = v
}

// GetFullMethod returns the FullMethod field value.
func (o *BackupOption) GetFullMethod() []BackupMethodOption {
	if o == nil {
		var ret []BackupMethodOption
		return ret
	}
	return o.FullMethod
}

// GetFullMethodOk returns a tuple with the FullMethod field value
// and a boolean to check if the value has been set.
func (o *BackupOption) GetFullMethodOk() (*[]BackupMethodOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullMethod, true
}

// SetFullMethod sets field value.
func (o *BackupOption) SetFullMethod(v []BackupMethodOption) {
	o.FullMethod = v
}

// GetContinuousMethod returns the ContinuousMethod field value if set, zero value otherwise.
func (o *BackupOption) GetContinuousMethod() []BackupMethodOption {
	if o == nil || o.ContinuousMethod == nil {
		var ret []BackupMethodOption
		return ret
	}
	return o.ContinuousMethod
}

// GetContinuousMethodOk returns a tuple with the ContinuousMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOption) GetContinuousMethodOk() (*[]BackupMethodOption, bool) {
	if o == nil || o.ContinuousMethod == nil {
		return nil, false
	}
	return &o.ContinuousMethod, true
}

// HasContinuousMethod returns a boolean if a field has been set.
func (o *BackupOption) HasContinuousMethod() bool {
	return o != nil && o.ContinuousMethod != nil
}

// SetContinuousMethod gets a reference to the given []BackupMethodOption and assigns it to the ContinuousMethod field.
func (o *BackupOption) SetContinuousMethod(v []BackupMethodOption) {
	o.ContinuousMethod = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["defaultMethod"] = o.DefaultMethod
	toSerialize["fullMethod"] = o.FullMethod
	if o.ContinuousMethod != nil {
		toSerialize["continuousMethod"] = o.ContinuousMethod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefaultMethod    *string               `json:"defaultMethod"`
		FullMethod       *[]BackupMethodOption `json:"fullMethod"`
		ContinuousMethod []BackupMethodOption  `json:"continuousMethod,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DefaultMethod == nil {
		return fmt.Errorf("required field defaultMethod missing")
	}
	if all.FullMethod == nil {
		return fmt.Errorf("required field fullMethod missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"defaultMethod", "fullMethod", "continuousMethod"})
	} else {
		return err
	}
	o.DefaultMethod = *all.DefaultMethod
	o.FullMethod = *all.FullMethod
	o.ContinuousMethod = all.ContinuousMethod

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
