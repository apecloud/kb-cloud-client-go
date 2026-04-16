// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BlueGreenDeploymentPrecheckItem blueGreenDeploymentPrecheckItem is a precheck item in the blue-green deployment precheck result.
type BlueGreenDeploymentPrecheckItem struct {
	// The name of the checker.
	CheckerName *string `json:"checkerName,omitempty"`
	// The description of the checker.
	CheckerDesc *string `json:"checkerDesc,omitempty"`
	// The status of a step in a blue-green deployment precheck.
	Status *BlueGreenDeploymentPrecheckStatus `json:"status,omitempty"`
	// The message from the checker.
	Message common.NullableString `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBlueGreenDeploymentPrecheckItem instantiates a new BlueGreenDeploymentPrecheckItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBlueGreenDeploymentPrecheckItem() *BlueGreenDeploymentPrecheckItem {
	this := BlueGreenDeploymentPrecheckItem{}
	return &this
}

// NewBlueGreenDeploymentPrecheckItemWithDefaults instantiates a new BlueGreenDeploymentPrecheckItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBlueGreenDeploymentPrecheckItemWithDefaults() *BlueGreenDeploymentPrecheckItem {
	this := BlueGreenDeploymentPrecheckItem{}
	return &this
}

// GetCheckerName returns the CheckerName field value if set, zero value otherwise.
func (o *BlueGreenDeploymentPrecheckItem) GetCheckerName() string {
	if o == nil || o.CheckerName == nil {
		var ret string
		return ret
	}
	return *o.CheckerName
}

// GetCheckerNameOk returns a tuple with the CheckerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentPrecheckItem) GetCheckerNameOk() (*string, bool) {
	if o == nil || o.CheckerName == nil {
		return nil, false
	}
	return o.CheckerName, true
}

// HasCheckerName returns a boolean if a field has been set.
func (o *BlueGreenDeploymentPrecheckItem) HasCheckerName() bool {
	return o != nil && o.CheckerName != nil
}

// SetCheckerName gets a reference to the given string and assigns it to the CheckerName field.
func (o *BlueGreenDeploymentPrecheckItem) SetCheckerName(v string) {
	o.CheckerName = &v
}

// GetCheckerDesc returns the CheckerDesc field value if set, zero value otherwise.
func (o *BlueGreenDeploymentPrecheckItem) GetCheckerDesc() string {
	if o == nil || o.CheckerDesc == nil {
		var ret string
		return ret
	}
	return *o.CheckerDesc
}

// GetCheckerDescOk returns a tuple with the CheckerDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentPrecheckItem) GetCheckerDescOk() (*string, bool) {
	if o == nil || o.CheckerDesc == nil {
		return nil, false
	}
	return o.CheckerDesc, true
}

// HasCheckerDesc returns a boolean if a field has been set.
func (o *BlueGreenDeploymentPrecheckItem) HasCheckerDesc() bool {
	return o != nil && o.CheckerDesc != nil
}

// SetCheckerDesc gets a reference to the given string and assigns it to the CheckerDesc field.
func (o *BlueGreenDeploymentPrecheckItem) SetCheckerDesc(v string) {
	o.CheckerDesc = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BlueGreenDeploymentPrecheckItem) GetStatus() BlueGreenDeploymentPrecheckStatus {
	if o == nil || o.Status == nil {
		var ret BlueGreenDeploymentPrecheckStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentPrecheckItem) GetStatusOk() (*BlueGreenDeploymentPrecheckStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BlueGreenDeploymentPrecheckItem) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given BlueGreenDeploymentPrecheckStatus and assigns it to the Status field.
func (o *BlueGreenDeploymentPrecheckItem) SetStatus(v BlueGreenDeploymentPrecheckStatus) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlueGreenDeploymentPrecheckItem) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BlueGreenDeploymentPrecheckItem) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *BlueGreenDeploymentPrecheckItem) HasMessage() bool {
	return o != nil && o.Message.IsSet()
}

// SetMessage gets a reference to the given common.NullableString and assigns it to the Message field.
func (o *BlueGreenDeploymentPrecheckItem) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil.
func (o *BlueGreenDeploymentPrecheckItem) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil.
func (o *BlueGreenDeploymentPrecheckItem) UnsetMessage() {
	o.Message.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o BlueGreenDeploymentPrecheckItem) MarshalJSON() ([]byte, error) {
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
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BlueGreenDeploymentPrecheckItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CheckerName *string                            `json:"checkerName,omitempty"`
		CheckerDesc *string                            `json:"checkerDesc,omitempty"`
		Status      *BlueGreenDeploymentPrecheckStatus `json:"status,omitempty"`
		Message     common.NullableString              `json:"message,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"checkerName", "checkerDesc", "status", "message"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CheckerName = all.CheckerName
	o.CheckerDesc = all.CheckerDesc
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
