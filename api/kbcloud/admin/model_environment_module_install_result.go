// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentModuleInstallResult Environment module installation check results and optional submitted task
type EnvironmentModuleInstallResult struct {
	// Checks that passed during this synchronous evaluation. These items are informational and do not block task submission.
	Pass []EnvironmentModuleInstallCheck `json:"pass"`
	// Checks that failed during this synchronous evaluation. When non-empty, no asynchronous installation task is submitted.
	Fail []EnvironmentModuleInstallCheck `json:"fail"`
	// ID of the submitted asynchronous installation task. It is absent for dry runs or failed checks.
	TaskId *string `json:"taskId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentModuleInstallResult instantiates a new EnvironmentModuleInstallResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentModuleInstallResult(pass []EnvironmentModuleInstallCheck, fail []EnvironmentModuleInstallCheck) *EnvironmentModuleInstallResult {
	this := EnvironmentModuleInstallResult{}
	this.Pass = pass
	this.Fail = fail
	return &this
}

// NewEnvironmentModuleInstallResultWithDefaults instantiates a new EnvironmentModuleInstallResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentModuleInstallResultWithDefaults() *EnvironmentModuleInstallResult {
	this := EnvironmentModuleInstallResult{}
	return &this
}

// GetPass returns the Pass field value.
func (o *EnvironmentModuleInstallResult) GetPass() []EnvironmentModuleInstallCheck {
	if o == nil {
		var ret []EnvironmentModuleInstallCheck
		return ret
	}
	return o.Pass
}

// GetPassOk returns a tuple with the Pass field value
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleInstallResult) GetPassOk() (*[]EnvironmentModuleInstallCheck, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pass, true
}

// SetPass sets field value.
func (o *EnvironmentModuleInstallResult) SetPass(v []EnvironmentModuleInstallCheck) {
	o.Pass = v
}

// GetFail returns the Fail field value.
func (o *EnvironmentModuleInstallResult) GetFail() []EnvironmentModuleInstallCheck {
	if o == nil {
		var ret []EnvironmentModuleInstallCheck
		return ret
	}
	return o.Fail
}

// GetFailOk returns a tuple with the Fail field value
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleInstallResult) GetFailOk() (*[]EnvironmentModuleInstallCheck, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fail, true
}

// SetFail sets field value.
func (o *EnvironmentModuleInstallResult) SetFail(v []EnvironmentModuleInstallCheck) {
	o.Fail = v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *EnvironmentModuleInstallResult) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleInstallResult) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *EnvironmentModuleInstallResult) HasTaskId() bool {
	return o != nil && o.TaskId != nil
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *EnvironmentModuleInstallResult) SetTaskId(v string) {
	o.TaskId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentModuleInstallResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["pass"] = o.Pass
	toSerialize["fail"] = o.Fail
	if o.TaskId != nil {
		toSerialize["taskId"] = o.TaskId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentModuleInstallResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pass   *[]EnvironmentModuleInstallCheck `json:"pass"`
		Fail   *[]EnvironmentModuleInstallCheck `json:"fail"`
		TaskId *string                          `json:"taskId,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Pass == nil {
		return fmt.Errorf("required field pass missing")
	}
	if all.Fail == nil {
		return fmt.Errorf("required field fail missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"pass", "fail", "taskId"})
	} else {
		return err
	}
	o.Pass = *all.Pass
	o.Fail = *all.Fail
	o.TaskId = all.TaskId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
