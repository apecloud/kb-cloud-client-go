// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsCustom Create a custom KubeBlocks OpsRequest. Specify exactly one of compName and compNames. Switchover requires a single component.
type OpsCustom struct {
	// component name
	CompName *string `json:"compName,omitempty"`
	// Component names for a single OpsRequest. Mutually exclusive with compName. All components share the same params.
	CompNames []string `json:"compNames,omitempty"`
	// ops definition name.
	OpsType string `json:"opsType"`
	// ops definition name.
	DependentOnOps []string `json:"dependentOnOps,omitempty"`
	// custom ops parameters
	Params   []OpsParameter `json:"params,omitempty"`
	Schedule *TaskSchedule  `json:"schedule,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsCustom instantiates a new OpsCustom object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsCustom(opsType string) *OpsCustom {
	this := OpsCustom{}
	this.OpsType = opsType
	return &this
}

// NewOpsCustomWithDefaults instantiates a new OpsCustom object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsCustomWithDefaults() *OpsCustom {
	this := OpsCustom{}
	return &this
}

// GetCompName returns the CompName field value if set, zero value otherwise.
func (o *OpsCustom) GetCompName() string {
	if o == nil || o.CompName == nil {
		var ret string
		return ret
	}
	return *o.CompName
}

// GetCompNameOk returns a tuple with the CompName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsCustom) GetCompNameOk() (*string, bool) {
	if o == nil || o.CompName == nil {
		return nil, false
	}
	return o.CompName, true
}

// HasCompName returns a boolean if a field has been set.
func (o *OpsCustom) HasCompName() bool {
	return o != nil && o.CompName != nil
}

// SetCompName gets a reference to the given string and assigns it to the CompName field.
func (o *OpsCustom) SetCompName(v string) {
	o.CompName = &v
}

// GetCompNames returns the CompNames field value if set, zero value otherwise.
func (o *OpsCustom) GetCompNames() []string {
	if o == nil || o.CompNames == nil {
		var ret []string
		return ret
	}
	return o.CompNames
}

// GetCompNamesOk returns a tuple with the CompNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsCustom) GetCompNamesOk() (*[]string, bool) {
	if o == nil || o.CompNames == nil {
		return nil, false
	}
	return &o.CompNames, true
}

// HasCompNames returns a boolean if a field has been set.
func (o *OpsCustom) HasCompNames() bool {
	return o != nil && o.CompNames != nil
}

// SetCompNames gets a reference to the given []string and assigns it to the CompNames field.
func (o *OpsCustom) SetCompNames(v []string) {
	o.CompNames = v
}

// GetOpsType returns the OpsType field value.
func (o *OpsCustom) GetOpsType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OpsType
}

// GetOpsTypeOk returns a tuple with the OpsType field value
// and a boolean to check if the value has been set.
func (o *OpsCustom) GetOpsTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpsType, true
}

// SetOpsType sets field value.
func (o *OpsCustom) SetOpsType(v string) {
	o.OpsType = v
}

// GetDependentOnOps returns the DependentOnOps field value if set, zero value otherwise.
func (o *OpsCustom) GetDependentOnOps() []string {
	if o == nil || o.DependentOnOps == nil {
		var ret []string
		return ret
	}
	return o.DependentOnOps
}

// GetDependentOnOpsOk returns a tuple with the DependentOnOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsCustom) GetDependentOnOpsOk() (*[]string, bool) {
	if o == nil || o.DependentOnOps == nil {
		return nil, false
	}
	return &o.DependentOnOps, true
}

// HasDependentOnOps returns a boolean if a field has been set.
func (o *OpsCustom) HasDependentOnOps() bool {
	return o != nil && o.DependentOnOps != nil
}

// SetDependentOnOps gets a reference to the given []string and assigns it to the DependentOnOps field.
func (o *OpsCustom) SetDependentOnOps(v []string) {
	o.DependentOnOps = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *OpsCustom) GetParams() []OpsParameter {
	if o == nil || o.Params == nil {
		var ret []OpsParameter
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsCustom) GetParamsOk() (*[]OpsParameter, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return &o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *OpsCustom) HasParams() bool {
	return o != nil && o.Params != nil
}

// SetParams gets a reference to the given []OpsParameter and assigns it to the Params field.
func (o *OpsCustom) SetParams(v []OpsParameter) {
	o.Params = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *OpsCustom) GetSchedule() TaskSchedule {
	if o == nil || o.Schedule == nil {
		var ret TaskSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsCustom) GetScheduleOk() (*TaskSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *OpsCustom) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given TaskSchedule and assigns it to the Schedule field.
func (o *OpsCustom) SetSchedule(v TaskSchedule) {
	o.Schedule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsCustom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CompName != nil {
		toSerialize["compName"] = o.CompName
	}
	if o.CompNames != nil {
		toSerialize["compNames"] = o.CompNames
	}
	toSerialize["opsType"] = o.OpsType
	if o.DependentOnOps != nil {
		toSerialize["dependentOnOps"] = o.DependentOnOps
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsCustom) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompName       *string        `json:"compName,omitempty"`
		CompNames      []string       `json:"compNames,omitempty"`
		OpsType        *string        `json:"opsType"`
		DependentOnOps []string       `json:"dependentOnOps,omitempty"`
		Params         []OpsParameter `json:"params,omitempty"`
		Schedule       *TaskSchedule  `json:"schedule,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.OpsType == nil {
		return fmt.Errorf("required field opsType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"compName", "compNames", "opsType", "dependentOnOps", "params", "schedule"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CompName = all.CompName
	o.CompNames = all.CompNames
	o.OpsType = *all.OpsType
	o.DependentOnOps = all.DependentOnOps
	o.Params = all.Params
	if all.Schedule != nil && all.Schedule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Schedule = all.Schedule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
