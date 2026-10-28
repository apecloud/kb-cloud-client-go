// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBPermissionSnapshot struct {
	// user or group
	SubjectType string                     `json:"subjectType"`
	Name        string                     `json:"name"`
	Direct      []DolphinDBPermissionState `json:"direct"`
	Effective   []DolphinDBPermissionState `json:"effective"`
	// True for users: effective is read from native getUserAccess(finalAccess=true). Groups expose direct settings only.
	EffectiveKnown bool                        `json:"effectiveKnown"`
	Inherited      []DolphinDBGroupPermissions `json:"inherited"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBPermissionSnapshot instantiates a new DolphinDBPermissionSnapshot object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBPermissionSnapshot(subjectType string, name string, direct []DolphinDBPermissionState, effective []DolphinDBPermissionState, effectiveKnown bool, inherited []DolphinDBGroupPermissions) *DolphinDBPermissionSnapshot {
	this := DolphinDBPermissionSnapshot{}
	this.SubjectType = subjectType
	this.Name = name
	this.Direct = direct
	this.Effective = effective
	this.EffectiveKnown = effectiveKnown
	this.Inherited = inherited
	return &this
}

// NewDolphinDBPermissionSnapshotWithDefaults instantiates a new DolphinDBPermissionSnapshot object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBPermissionSnapshotWithDefaults() *DolphinDBPermissionSnapshot {
	this := DolphinDBPermissionSnapshot{}
	return &this
}

// GetSubjectType returns the SubjectType field value.
func (o *DolphinDBPermissionSnapshot) GetSubjectType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SubjectType
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value
// and a boolean to check if the value has been set.
func (o *DolphinDBPermissionSnapshot) GetSubjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectType, true
}

// SetSubjectType sets field value.
func (o *DolphinDBPermissionSnapshot) SetSubjectType(v string) {
	o.SubjectType = v
}

// GetName returns the Name field value.
func (o *DolphinDBPermissionSnapshot) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DolphinDBPermissionSnapshot) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DolphinDBPermissionSnapshot) SetName(v string) {
	o.Name = v
}

// GetDirect returns the Direct field value.
func (o *DolphinDBPermissionSnapshot) GetDirect() []DolphinDBPermissionState {
	if o == nil {
		var ret []DolphinDBPermissionState
		return ret
	}
	return o.Direct
}

// GetDirectOk returns a tuple with the Direct field value
// and a boolean to check if the value has been set.
func (o *DolphinDBPermissionSnapshot) GetDirectOk() (*[]DolphinDBPermissionState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direct, true
}

// SetDirect sets field value.
func (o *DolphinDBPermissionSnapshot) SetDirect(v []DolphinDBPermissionState) {
	o.Direct = v
}

// GetEffective returns the Effective field value.
func (o *DolphinDBPermissionSnapshot) GetEffective() []DolphinDBPermissionState {
	if o == nil {
		var ret []DolphinDBPermissionState
		return ret
	}
	return o.Effective
}

// GetEffectiveOk returns a tuple with the Effective field value
// and a boolean to check if the value has been set.
func (o *DolphinDBPermissionSnapshot) GetEffectiveOk() (*[]DolphinDBPermissionState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Effective, true
}

// SetEffective sets field value.
func (o *DolphinDBPermissionSnapshot) SetEffective(v []DolphinDBPermissionState) {
	o.Effective = v
}

// GetEffectiveKnown returns the EffectiveKnown field value.
func (o *DolphinDBPermissionSnapshot) GetEffectiveKnown() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.EffectiveKnown
}

// GetEffectiveKnownOk returns a tuple with the EffectiveKnown field value
// and a boolean to check if the value has been set.
func (o *DolphinDBPermissionSnapshot) GetEffectiveKnownOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveKnown, true
}

// SetEffectiveKnown sets field value.
func (o *DolphinDBPermissionSnapshot) SetEffectiveKnown(v bool) {
	o.EffectiveKnown = v
}

// GetInherited returns the Inherited field value.
func (o *DolphinDBPermissionSnapshot) GetInherited() []DolphinDBGroupPermissions {
	if o == nil {
		var ret []DolphinDBGroupPermissions
		return ret
	}
	return o.Inherited
}

// GetInheritedOk returns a tuple with the Inherited field value
// and a boolean to check if the value has been set.
func (o *DolphinDBPermissionSnapshot) GetInheritedOk() (*[]DolphinDBGroupPermissions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inherited, true
}

// SetInherited sets field value.
func (o *DolphinDBPermissionSnapshot) SetInherited(v []DolphinDBGroupPermissions) {
	o.Inherited = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBPermissionSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["subjectType"] = o.SubjectType
	toSerialize["name"] = o.Name
	toSerialize["direct"] = o.Direct
	toSerialize["effective"] = o.Effective
	toSerialize["effectiveKnown"] = o.EffectiveKnown
	toSerialize["inherited"] = o.Inherited

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBPermissionSnapshot) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SubjectType    *string                      `json:"subjectType"`
		Name           *string                      `json:"name"`
		Direct         *[]DolphinDBPermissionState  `json:"direct"`
		Effective      *[]DolphinDBPermissionState  `json:"effective"`
		EffectiveKnown *bool                        `json:"effectiveKnown"`
		Inherited      *[]DolphinDBGroupPermissions `json:"inherited"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SubjectType == nil {
		return fmt.Errorf("required field subjectType missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Direct == nil {
		return fmt.Errorf("required field direct missing")
	}
	if all.Effective == nil {
		return fmt.Errorf("required field effective missing")
	}
	if all.EffectiveKnown == nil {
		return fmt.Errorf("required field effectiveKnown missing")
	}
	if all.Inherited == nil {
		return fmt.Errorf("required field inherited missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"subjectType", "name", "direct", "effective", "effectiveKnown", "inherited"})
	} else {
		return err
	}
	o.SubjectType = *all.SubjectType
	o.Name = *all.Name
	o.Direct = *all.Direct
	o.Effective = *all.Effective
	o.EffectiveKnown = *all.EffectiveKnown
	o.Inherited = *all.Inherited

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
