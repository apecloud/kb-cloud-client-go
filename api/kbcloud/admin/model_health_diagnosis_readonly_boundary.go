// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HealthDiagnosisReadonlyBoundary struct {
	// Whether the module only performs read-only collection.
	ReadOnly bool `json:"readOnly"`
	// Source used for the read-only collection.
	Source string `json:"source"`
	// Dangerous action types intentionally not provided by this module.
	DangerousActions []string `json:"dangerousActions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHealthDiagnosisReadonlyBoundary instantiates a new HealthDiagnosisReadonlyBoundary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHealthDiagnosisReadonlyBoundary(readOnly bool, source string, dangerousActions []string) *HealthDiagnosisReadonlyBoundary {
	this := HealthDiagnosisReadonlyBoundary{}
	this.ReadOnly = readOnly
	this.Source = source
	this.DangerousActions = dangerousActions
	return &this
}

// NewHealthDiagnosisReadonlyBoundaryWithDefaults instantiates a new HealthDiagnosisReadonlyBoundary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHealthDiagnosisReadonlyBoundaryWithDefaults() *HealthDiagnosisReadonlyBoundary {
	this := HealthDiagnosisReadonlyBoundary{}
	return &this
}

// GetReadOnly returns the ReadOnly field value.
func (o *HealthDiagnosisReadonlyBoundary) GetReadOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisReadonlyBoundary) GetReadOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadOnly, true
}

// SetReadOnly sets field value.
func (o *HealthDiagnosisReadonlyBoundary) SetReadOnly(v bool) {
	o.ReadOnly = v
}

// GetSource returns the Source field value.
func (o *HealthDiagnosisReadonlyBoundary) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisReadonlyBoundary) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *HealthDiagnosisReadonlyBoundary) SetSource(v string) {
	o.Source = v
}

// GetDangerousActions returns the DangerousActions field value.
func (o *HealthDiagnosisReadonlyBoundary) GetDangerousActions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DangerousActions
}

// GetDangerousActionsOk returns a tuple with the DangerousActions field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisReadonlyBoundary) GetDangerousActionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DangerousActions, true
}

// SetDangerousActions sets field value.
func (o *HealthDiagnosisReadonlyBoundary) SetDangerousActions(v []string) {
	o.DangerousActions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HealthDiagnosisReadonlyBoundary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["readOnly"] = o.ReadOnly
	toSerialize["source"] = o.Source
	toSerialize["dangerousActions"] = o.DangerousActions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HealthDiagnosisReadonlyBoundary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ReadOnly         *bool     `json:"readOnly"`
		Source           *string   `json:"source"`
		DangerousActions *[]string `json:"dangerousActions"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ReadOnly == nil {
		return fmt.Errorf("required field readOnly missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.DangerousActions == nil {
		return fmt.Errorf("required field dangerousActions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"readOnly", "source", "dangerousActions"})
	} else {
		return err
	}
	o.ReadOnly = *all.ReadOnly
	o.Source = *all.Source
	o.DangerousActions = *all.DangerousActions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
