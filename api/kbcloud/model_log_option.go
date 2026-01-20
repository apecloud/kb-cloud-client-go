// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type LogOption struct {
	Component                    string `json:"component"`
	Error                        bool   `json:"error"`
	Slow                         bool   `json:"slow"`
	Audit                        bool   `json:"audit"`
	Running                      bool   `json:"running"`
	DisableDisasterRecoveryAudit *bool  `json:"disableDisasterRecoveryAudit,omitempty"`
	// SQL audit switch configuration. Either sql or parameter must be provided.
	SqlAuditSwitch *SqlAuditSwitchOption `json:"sqlAuditSwitch,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogOption instantiates a new LogOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogOption(component string, error bool, slow bool, audit bool, running bool) *LogOption {
	this := LogOption{}
	this.Component = component
	this.Error = error
	this.Slow = slow
	this.Audit = audit
	this.Running = running
	return &this
}

// NewLogOptionWithDefaults instantiates a new LogOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogOptionWithDefaults() *LogOption {
	this := LogOption{}
	return &this
}

// GetComponent returns the Component field value.
func (o *LogOption) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *LogOption) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *LogOption) SetComponent(v string) {
	o.Component = v
}

// GetError returns the Error field value.
func (o *LogOption) GetError() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *LogOption) GetErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value.
func (o *LogOption) SetError(v bool) {
	o.Error = v
}

// GetSlow returns the Slow field value.
func (o *LogOption) GetSlow() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Slow
}

// GetSlowOk returns a tuple with the Slow field value
// and a boolean to check if the value has been set.
func (o *LogOption) GetSlowOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slow, true
}

// SetSlow sets field value.
func (o *LogOption) SetSlow(v bool) {
	o.Slow = v
}

// GetAudit returns the Audit field value.
func (o *LogOption) GetAudit() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Audit
}

// GetAuditOk returns a tuple with the Audit field value
// and a boolean to check if the value has been set.
func (o *LogOption) GetAuditOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audit, true
}

// SetAudit sets field value.
func (o *LogOption) SetAudit(v bool) {
	o.Audit = v
}

// GetRunning returns the Running field value.
func (o *LogOption) GetRunning() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Running
}

// GetRunningOk returns a tuple with the Running field value
// and a boolean to check if the value has been set.
func (o *LogOption) GetRunningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Running, true
}

// SetRunning sets field value.
func (o *LogOption) SetRunning(v bool) {
	o.Running = v
}

// GetDisableDisasterRecoveryAudit returns the DisableDisasterRecoveryAudit field value if set, zero value otherwise.
func (o *LogOption) GetDisableDisasterRecoveryAudit() bool {
	if o == nil || o.DisableDisasterRecoveryAudit == nil {
		var ret bool
		return ret
	}
	return *o.DisableDisasterRecoveryAudit
}

// GetDisableDisasterRecoveryAuditOk returns a tuple with the DisableDisasterRecoveryAudit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogOption) GetDisableDisasterRecoveryAuditOk() (*bool, bool) {
	if o == nil || o.DisableDisasterRecoveryAudit == nil {
		return nil, false
	}
	return o.DisableDisasterRecoveryAudit, true
}

// HasDisableDisasterRecoveryAudit returns a boolean if a field has been set.
func (o *LogOption) HasDisableDisasterRecoveryAudit() bool {
	return o != nil && o.DisableDisasterRecoveryAudit != nil
}

// SetDisableDisasterRecoveryAudit gets a reference to the given bool and assigns it to the DisableDisasterRecoveryAudit field.
func (o *LogOption) SetDisableDisasterRecoveryAudit(v bool) {
	o.DisableDisasterRecoveryAudit = &v
}

// GetSqlAuditSwitch returns the SqlAuditSwitch field value if set, zero value otherwise.
func (o *LogOption) GetSqlAuditSwitch() SqlAuditSwitchOption {
	if o == nil || o.SqlAuditSwitch == nil {
		var ret SqlAuditSwitchOption
		return ret
	}
	return *o.SqlAuditSwitch
}

// GetSqlAuditSwitchOk returns a tuple with the SqlAuditSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogOption) GetSqlAuditSwitchOk() (*SqlAuditSwitchOption, bool) {
	if o == nil || o.SqlAuditSwitch == nil {
		return nil, false
	}
	return o.SqlAuditSwitch, true
}

// HasSqlAuditSwitch returns a boolean if a field has been set.
func (o *LogOption) HasSqlAuditSwitch() bool {
	return o != nil && o.SqlAuditSwitch != nil
}

// SetSqlAuditSwitch gets a reference to the given SqlAuditSwitchOption and assigns it to the SqlAuditSwitch field.
func (o *LogOption) SetSqlAuditSwitch(v SqlAuditSwitchOption) {
	o.SqlAuditSwitch = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["error"] = o.Error
	toSerialize["slow"] = o.Slow
	toSerialize["audit"] = o.Audit
	toSerialize["running"] = o.Running
	if o.DisableDisasterRecoveryAudit != nil {
		toSerialize["disableDisasterRecoveryAudit"] = o.DisableDisasterRecoveryAudit
	}
	if o.SqlAuditSwitch != nil {
		toSerialize["sqlAuditSwitch"] = o.SqlAuditSwitch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component                    *string               `json:"component"`
		Error                        *bool                 `json:"error"`
		Slow                         *bool                 `json:"slow"`
		Audit                        *bool                 `json:"audit"`
		Running                      *bool                 `json:"running"`
		DisableDisasterRecoveryAudit *bool                 `json:"disableDisasterRecoveryAudit,omitempty"`
		SqlAuditSwitch               *SqlAuditSwitchOption `json:"sqlAuditSwitch,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Error == nil {
		return fmt.Errorf("required field error missing")
	}
	if all.Slow == nil {
		return fmt.Errorf("required field slow missing")
	}
	if all.Audit == nil {
		return fmt.Errorf("required field audit missing")
	}
	if all.Running == nil {
		return fmt.Errorf("required field running missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "error", "slow", "audit", "running", "disableDisasterRecoveryAudit", "sqlAuditSwitch"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Component = *all.Component
	o.Error = *all.Error
	o.Slow = *all.Slow
	o.Audit = *all.Audit
	o.Running = *all.Running
	o.DisableDisasterRecoveryAudit = all.DisableDisasterRecoveryAudit
	if all.SqlAuditSwitch != nil && all.SqlAuditSwitch.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SqlAuditSwitch = all.SqlAuditSwitch

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
