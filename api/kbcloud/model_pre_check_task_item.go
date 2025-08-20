// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type PreCheckTaskItem struct {
	CheckerName *string               `json:"checkerName,omitempty"`
	CheckerDesc *string               `json:"checkerDesc,omitempty"`
	Message     common.NullableString `json:"message,omitempty"`
	Status      *PreCheckStatus       `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPreCheckTaskItem instantiates a new PreCheckTaskItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPreCheckTaskItem() *PreCheckTaskItem {
	this := PreCheckTaskItem{}
	return &this
}

// NewPreCheckTaskItemWithDefaults instantiates a new PreCheckTaskItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPreCheckTaskItemWithDefaults() *PreCheckTaskItem {
	this := PreCheckTaskItem{}
	return &this
}

// GetCheckerName returns the CheckerName field value if set, zero value otherwise.
func (o *PreCheckTaskItem) GetCheckerName() string {
	if o == nil || o.CheckerName == nil {
		var ret string
		return ret
	}
	return *o.CheckerName
}

// GetCheckerNameOk returns a tuple with the CheckerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreCheckTaskItem) GetCheckerNameOk() (*string, bool) {
	if o == nil || o.CheckerName == nil {
		return nil, false
	}
	return o.CheckerName, true
}

// HasCheckerName returns a boolean if a field has been set.
func (o *PreCheckTaskItem) HasCheckerName() bool {
	return o != nil && o.CheckerName != nil
}

// SetCheckerName gets a reference to the given string and assigns it to the CheckerName field.
func (o *PreCheckTaskItem) SetCheckerName(v string) {
	o.CheckerName = &v
}

// GetCheckerDesc returns the CheckerDesc field value if set, zero value otherwise.
func (o *PreCheckTaskItem) GetCheckerDesc() string {
	if o == nil || o.CheckerDesc == nil {
		var ret string
		return ret
	}
	return *o.CheckerDesc
}

// GetCheckerDescOk returns a tuple with the CheckerDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreCheckTaskItem) GetCheckerDescOk() (*string, bool) {
	if o == nil || o.CheckerDesc == nil {
		return nil, false
	}
	return o.CheckerDesc, true
}

// HasCheckerDesc returns a boolean if a field has been set.
func (o *PreCheckTaskItem) HasCheckerDesc() bool {
	return o != nil && o.CheckerDesc != nil
}

// SetCheckerDesc gets a reference to the given string and assigns it to the CheckerDesc field.
func (o *PreCheckTaskItem) SetCheckerDesc(v string) {
	o.CheckerDesc = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PreCheckTaskItem) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PreCheckTaskItem) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *PreCheckTaskItem) HasMessage() bool {
	return o != nil && o.Message.IsSet()
}

// SetMessage gets a reference to the given common.NullableString and assigns it to the Message field.
func (o *PreCheckTaskItem) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil.
func (o *PreCheckTaskItem) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil.
func (o *PreCheckTaskItem) UnsetMessage() {
	o.Message.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PreCheckTaskItem) GetStatus() PreCheckStatus {
	if o == nil || o.Status == nil {
		var ret PreCheckStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreCheckTaskItem) GetStatusOk() (*PreCheckStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PreCheckTaskItem) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given PreCheckStatus and assigns it to the Status field.
func (o *PreCheckTaskItem) SetStatus(v PreCheckStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PreCheckTaskItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CheckerName != nil {
		toSerialize["checkerName"] = o.CheckerName
	}
	if o.CheckerDesc != nil {
		toSerialize["checkerDesc"] = o.CheckerDesc
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PreCheckTaskItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CheckerName *string               `json:"checkerName,omitempty"`
		CheckerDesc *string               `json:"checkerDesc,omitempty"`
		Message     common.NullableString `json:"message,omitempty"`
		Status      *PreCheckStatus       `json:"status,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"checkerName", "checkerDesc", "message", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CheckerName = all.CheckerName
	o.CheckerDesc = all.CheckerDesc
	o.Message = all.Message
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
