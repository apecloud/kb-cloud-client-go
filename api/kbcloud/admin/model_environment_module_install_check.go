// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentModuleInstallCheck A localized result item produced by an environment module installation check
type EnvironmentModuleInstallCheck struct {
	// Identifier included with this check result. The localized message is returned directly and is not resolved from this value.
	Code    string               `json:"code"`
	Message LocalizedDescription `json:"message"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentModuleInstallCheck instantiates a new EnvironmentModuleInstallCheck object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentModuleInstallCheck(code string, message LocalizedDescription) *EnvironmentModuleInstallCheck {
	this := EnvironmentModuleInstallCheck{}
	this.Code = code
	this.Message = message
	return &this
}

// NewEnvironmentModuleInstallCheckWithDefaults instantiates a new EnvironmentModuleInstallCheck object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentModuleInstallCheckWithDefaults() *EnvironmentModuleInstallCheck {
	this := EnvironmentModuleInstallCheck{}
	return &this
}

// GetCode returns the Code field value.
func (o *EnvironmentModuleInstallCheck) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleInstallCheck) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *EnvironmentModuleInstallCheck) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value.
func (o *EnvironmentModuleInstallCheck) GetMessage() LocalizedDescription {
	if o == nil {
		var ret LocalizedDescription
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleInstallCheck) GetMessageOk() (*LocalizedDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *EnvironmentModuleInstallCheck) SetMessage(v LocalizedDescription) {
	o.Message = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentModuleInstallCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentModuleInstallCheck) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code    *string               `json:"code"`
		Message *LocalizedDescription `json:"message"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Code == nil {
		return fmt.Errorf("required field code missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"code", "message"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Code = *all.Code
	if all.Message.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Message = *all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
