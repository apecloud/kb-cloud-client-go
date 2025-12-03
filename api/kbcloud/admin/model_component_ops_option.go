// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ComponentOpsOption struct {
	// List of modes that this component supports. if not set, support all modes.
	Modes     []string `json:"modes,omitempty"`
	Component string   `json:"component"`
	DisableHa *bool    `json:"disableHA,omitempty"`
	// parameter for rebuild instance ops
	InPlace *bool `json:"inPlace,omitempty"`
	// indicate whether backup is required when Inplace is true
	NeedBackupWhenInPlace *bool `json:"needBackupWhenInPlace,omitempty"`
	// indicate whether backup is required for this ops
	BackupRequired *bool `json:"backupRequired,omitempty"`
	// indicate the backup method when inplace is true
	BackupMethod *ComponentOpsOptionBackupMethod    `json:"backupMethod,omitempty"`
	RestoreEnv   []ComponentOpsOptionRestoreEnvItem `json:"restoreEnv,omitempty"`
	// indicate whether rollback is supported for this ops
	RollbackSupported  *bool                                 `json:"rollbackSupported,omitempty"`
	DependentCustomOps *ComponentOpsOptionDependentCustomOps `json:"dependentCustomOps,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentOpsOption instantiates a new ComponentOpsOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentOpsOption(component string) *ComponentOpsOption {
	this := ComponentOpsOption{}
	this.Component = component
	var disableHa bool = false
	this.DisableHa = &disableHa
	var inPlace bool = false
	this.InPlace = &inPlace
	var needBackupWhenInPlace bool = false
	this.NeedBackupWhenInPlace = &needBackupWhenInPlace
	var backupRequired bool = false
	this.BackupRequired = &backupRequired
	var rollbackSupported bool = false
	this.RollbackSupported = &rollbackSupported
	return &this
}

// NewComponentOpsOptionWithDefaults instantiates a new ComponentOpsOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentOpsOptionWithDefaults() *ComponentOpsOption {
	this := ComponentOpsOption{}
	var disableHa bool = false
	this.DisableHa = &disableHa
	var inPlace bool = false
	this.InPlace = &inPlace
	var needBackupWhenInPlace bool = false
	this.NeedBackupWhenInPlace = &needBackupWhenInPlace
	var backupRequired bool = false
	this.BackupRequired = &backupRequired
	var rollbackSupported bool = false
	this.RollbackSupported = &rollbackSupported
	return &this
}

// GetModes returns the Modes field value if set, zero value otherwise.
func (o *ComponentOpsOption) GetModes() []string {
	if o == nil || o.Modes == nil {
		var ret []string
		return ret
	}
	return o.Modes
}

// GetModesOk returns a tuple with the Modes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOption) GetModesOk() (*[]string, bool) {
	if o == nil || o.Modes == nil {
		return nil, false
	}
	return &o.Modes, true
}

// HasModes returns a boolean if a field has been set.
func (o *ComponentOpsOption) HasModes() bool {
	return o != nil && o.Modes != nil
}

// SetModes gets a reference to the given []string and assigns it to the Modes field.
func (o *ComponentOpsOption) SetModes(v []string) {
	o.Modes = v
}

// GetComponent returns the Component field value.
func (o *ComponentOpsOption) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ComponentOpsOption) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *ComponentOpsOption) SetComponent(v string) {
	o.Component = v
}

// GetDisableHa returns the DisableHa field value if set, zero value otherwise.
func (o *ComponentOpsOption) GetDisableHa() bool {
	if o == nil || o.DisableHa == nil {
		var ret bool
		return ret
	}
	return *o.DisableHa
}

// GetDisableHaOk returns a tuple with the DisableHa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOption) GetDisableHaOk() (*bool, bool) {
	if o == nil || o.DisableHa == nil {
		return nil, false
	}
	return o.DisableHa, true
}

// HasDisableHa returns a boolean if a field has been set.
func (o *ComponentOpsOption) HasDisableHa() bool {
	return o != nil && o.DisableHa != nil
}

// SetDisableHa gets a reference to the given bool and assigns it to the DisableHa field.
func (o *ComponentOpsOption) SetDisableHa(v bool) {
	o.DisableHa = &v
}

// GetInPlace returns the InPlace field value if set, zero value otherwise.
func (o *ComponentOpsOption) GetInPlace() bool {
	if o == nil || o.InPlace == nil {
		var ret bool
		return ret
	}
	return *o.InPlace
}

// GetInPlaceOk returns a tuple with the InPlace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOption) GetInPlaceOk() (*bool, bool) {
	if o == nil || o.InPlace == nil {
		return nil, false
	}
	return o.InPlace, true
}

// HasInPlace returns a boolean if a field has been set.
func (o *ComponentOpsOption) HasInPlace() bool {
	return o != nil && o.InPlace != nil
}

// SetInPlace gets a reference to the given bool and assigns it to the InPlace field.
func (o *ComponentOpsOption) SetInPlace(v bool) {
	o.InPlace = &v
}

// GetNeedBackupWhenInPlace returns the NeedBackupWhenInPlace field value if set, zero value otherwise.
func (o *ComponentOpsOption) GetNeedBackupWhenInPlace() bool {
	if o == nil || o.NeedBackupWhenInPlace == nil {
		var ret bool
		return ret
	}
	return *o.NeedBackupWhenInPlace
}

// GetNeedBackupWhenInPlaceOk returns a tuple with the NeedBackupWhenInPlace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOption) GetNeedBackupWhenInPlaceOk() (*bool, bool) {
	if o == nil || o.NeedBackupWhenInPlace == nil {
		return nil, false
	}
	return o.NeedBackupWhenInPlace, true
}

// HasNeedBackupWhenInPlace returns a boolean if a field has been set.
func (o *ComponentOpsOption) HasNeedBackupWhenInPlace() bool {
	return o != nil && o.NeedBackupWhenInPlace != nil
}

// SetNeedBackupWhenInPlace gets a reference to the given bool and assigns it to the NeedBackupWhenInPlace field.
func (o *ComponentOpsOption) SetNeedBackupWhenInPlace(v bool) {
	o.NeedBackupWhenInPlace = &v
}

// GetBackupRequired returns the BackupRequired field value if set, zero value otherwise.
func (o *ComponentOpsOption) GetBackupRequired() bool {
	if o == nil || o.BackupRequired == nil {
		var ret bool
		return ret
	}
	return *o.BackupRequired
}

// GetBackupRequiredOk returns a tuple with the BackupRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOption) GetBackupRequiredOk() (*bool, bool) {
	if o == nil || o.BackupRequired == nil {
		return nil, false
	}
	return o.BackupRequired, true
}

// HasBackupRequired returns a boolean if a field has been set.
func (o *ComponentOpsOption) HasBackupRequired() bool {
	return o != nil && o.BackupRequired != nil
}

// SetBackupRequired gets a reference to the given bool and assigns it to the BackupRequired field.
func (o *ComponentOpsOption) SetBackupRequired(v bool) {
	o.BackupRequired = &v
}

// GetBackupMethod returns the BackupMethod field value if set, zero value otherwise.
func (o *ComponentOpsOption) GetBackupMethod() ComponentOpsOptionBackupMethod {
	if o == nil || o.BackupMethod == nil {
		var ret ComponentOpsOptionBackupMethod
		return ret
	}
	return *o.BackupMethod
}

// GetBackupMethodOk returns a tuple with the BackupMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOption) GetBackupMethodOk() (*ComponentOpsOptionBackupMethod, bool) {
	if o == nil || o.BackupMethod == nil {
		return nil, false
	}
	return o.BackupMethod, true
}

// HasBackupMethod returns a boolean if a field has been set.
func (o *ComponentOpsOption) HasBackupMethod() bool {
	return o != nil && o.BackupMethod != nil
}

// SetBackupMethod gets a reference to the given ComponentOpsOptionBackupMethod and assigns it to the BackupMethod field.
func (o *ComponentOpsOption) SetBackupMethod(v ComponentOpsOptionBackupMethod) {
	o.BackupMethod = &v
}

// GetRestoreEnv returns the RestoreEnv field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentOpsOption) GetRestoreEnv() []ComponentOpsOptionRestoreEnvItem {
	if o == nil {
		var ret []ComponentOpsOptionRestoreEnvItem
		return ret
	}
	return o.RestoreEnv
}

// GetRestoreEnvOk returns a tuple with the RestoreEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ComponentOpsOption) GetRestoreEnvOk() (*[]ComponentOpsOptionRestoreEnvItem, bool) {
	if o == nil || o.RestoreEnv == nil {
		return nil, false
	}
	return &o.RestoreEnv, true
}

// HasRestoreEnv returns a boolean if a field has been set.
func (o *ComponentOpsOption) HasRestoreEnv() bool {
	return o != nil && o.RestoreEnv != nil
}

// SetRestoreEnv gets a reference to the given []ComponentOpsOptionRestoreEnvItem and assigns it to the RestoreEnv field.
func (o *ComponentOpsOption) SetRestoreEnv(v []ComponentOpsOptionRestoreEnvItem) {
	o.RestoreEnv = v
}

// GetRollbackSupported returns the RollbackSupported field value if set, zero value otherwise.
func (o *ComponentOpsOption) GetRollbackSupported() bool {
	if o == nil || o.RollbackSupported == nil {
		var ret bool
		return ret
	}
	return *o.RollbackSupported
}

// GetRollbackSupportedOk returns a tuple with the RollbackSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOption) GetRollbackSupportedOk() (*bool, bool) {
	if o == nil || o.RollbackSupported == nil {
		return nil, false
	}
	return o.RollbackSupported, true
}

// HasRollbackSupported returns a boolean if a field has been set.
func (o *ComponentOpsOption) HasRollbackSupported() bool {
	return o != nil && o.RollbackSupported != nil
}

// SetRollbackSupported gets a reference to the given bool and assigns it to the RollbackSupported field.
func (o *ComponentOpsOption) SetRollbackSupported(v bool) {
	o.RollbackSupported = &v
}

// GetDependentCustomOps returns the DependentCustomOps field value if set, zero value otherwise.
func (o *ComponentOpsOption) GetDependentCustomOps() ComponentOpsOptionDependentCustomOps {
	if o == nil || o.DependentCustomOps == nil {
		var ret ComponentOpsOptionDependentCustomOps
		return ret
	}
	return *o.DependentCustomOps
}

// GetDependentCustomOpsOk returns a tuple with the DependentCustomOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOption) GetDependentCustomOpsOk() (*ComponentOpsOptionDependentCustomOps, bool) {
	if o == nil || o.DependentCustomOps == nil {
		return nil, false
	}
	return o.DependentCustomOps, true
}

// HasDependentCustomOps returns a boolean if a field has been set.
func (o *ComponentOpsOption) HasDependentCustomOps() bool {
	return o != nil && o.DependentCustomOps != nil
}

// SetDependentCustomOps gets a reference to the given ComponentOpsOptionDependentCustomOps and assigns it to the DependentCustomOps field.
func (o *ComponentOpsOption) SetDependentCustomOps(v ComponentOpsOptionDependentCustomOps) {
	o.DependentCustomOps = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentOpsOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Modes != nil {
		toSerialize["modes"] = o.Modes
	}
	toSerialize["component"] = o.Component
	if o.DisableHa != nil {
		toSerialize["disableHA"] = o.DisableHa
	}
	if o.InPlace != nil {
		toSerialize["inPlace"] = o.InPlace
	}
	if o.NeedBackupWhenInPlace != nil {
		toSerialize["needBackupWhenInPlace"] = o.NeedBackupWhenInPlace
	}
	if o.BackupRequired != nil {
		toSerialize["backupRequired"] = o.BackupRequired
	}
	if o.BackupMethod != nil {
		toSerialize["backupMethod"] = o.BackupMethod
	}
	if o.RestoreEnv != nil {
		toSerialize["restoreEnv"] = o.RestoreEnv
	}
	if o.RollbackSupported != nil {
		toSerialize["rollbackSupported"] = o.RollbackSupported
	}
	if o.DependentCustomOps != nil {
		toSerialize["dependentCustomOps"] = o.DependentCustomOps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentOpsOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Modes                 []string                              `json:"modes,omitempty"`
		Component             *string                               `json:"component"`
		DisableHa             *bool                                 `json:"disableHA,omitempty"`
		InPlace               *bool                                 `json:"inPlace,omitempty"`
		NeedBackupWhenInPlace *bool                                 `json:"needBackupWhenInPlace,omitempty"`
		BackupRequired        *bool                                 `json:"backupRequired,omitempty"`
		BackupMethod          *ComponentOpsOptionBackupMethod       `json:"backupMethod,omitempty"`
		RestoreEnv            []ComponentOpsOptionRestoreEnvItem    `json:"restoreEnv,omitempty"`
		RollbackSupported     *bool                                 `json:"rollbackSupported,omitempty"`
		DependentCustomOps    *ComponentOpsOptionDependentCustomOps `json:"dependentCustomOps,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"modes", "component", "disableHA", "inPlace", "needBackupWhenInPlace", "backupRequired", "backupMethod", "restoreEnv", "rollbackSupported", "dependentCustomOps"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Modes = all.Modes
	o.Component = *all.Component
	o.DisableHa = all.DisableHa
	o.InPlace = all.InPlace
	o.NeedBackupWhenInPlace = all.NeedBackupWhenInPlace
	o.BackupRequired = all.BackupRequired
	if all.BackupMethod != nil && all.BackupMethod.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BackupMethod = all.BackupMethod
	o.RestoreEnv = all.RestoreEnv
	o.RollbackSupported = all.RollbackSupported
	if all.DependentCustomOps != nil && all.DependentCustomOps.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DependentCustomOps = all.DependentCustomOps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
