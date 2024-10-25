// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// NODESCRIPTION BackupRepoCheck
type BackupRepoCheck struct {
	// whether backup repo pass the check
	Success *bool `json:"success,omitempty"`
	// the info when failed to pass the check
	Message interface{} `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupRepoCheck instantiates a new BackupRepoCheck object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupRepoCheck() *BackupRepoCheck {
	this := BackupRepoCheck{}
	return &this
}

// NewBackupRepoCheckWithDefaults instantiates a new BackupRepoCheck object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupRepoCheckWithDefaults() *BackupRepoCheck {
	this := BackupRepoCheck{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *BackupRepoCheck) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoCheck) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *BackupRepoCheck) HasSuccess() bool {
	return o != nil && o.Success != nil
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *BackupRepoCheck) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BackupRepoCheck) GetMessage() interface{} {
	if o == nil || o.Message == nil {
		var ret interface{}
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoCheck) GetMessageOk() (*interface{}, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return &o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BackupRepoCheck) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given interface{} and assigns it to the Message field.
func (o *BackupRepoCheck) SetMessage(v interface{}) {
	o.Message = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupRepoCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupRepoCheck) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Success *bool       `json:"success,omitempty"`
		Message interface{} `json:"message,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"success", "message"})
	} else {
		return err
	}
	o.Success = all.Success
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
