// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ObjectPrivilegePrerequisiteCheck struct {
	Target    ObjectPrivilegeTarget             `json:"target"`
	Privilege string                            `json:"privilege"`
	Status    ObjectPrivilegePrerequisiteStatus `json:"status"`
	// Optional explanation. Checks are not a complete proof of executable access. Missing checks do not imply granted.
	Reason *string `json:"reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObjectPrivilegePrerequisiteCheck instantiates a new ObjectPrivilegePrerequisiteCheck object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObjectPrivilegePrerequisiteCheck(target ObjectPrivilegeTarget, privilege string, status ObjectPrivilegePrerequisiteStatus) *ObjectPrivilegePrerequisiteCheck {
	this := ObjectPrivilegePrerequisiteCheck{}
	this.Target = target
	this.Privilege = privilege
	this.Status = status
	return &this
}

// NewObjectPrivilegePrerequisiteCheckWithDefaults instantiates a new ObjectPrivilegePrerequisiteCheck object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObjectPrivilegePrerequisiteCheckWithDefaults() *ObjectPrivilegePrerequisiteCheck {
	this := ObjectPrivilegePrerequisiteCheck{}
	return &this
}

// GetTarget returns the Target field value.
func (o *ObjectPrivilegePrerequisiteCheck) GetTarget() ObjectPrivilegeTarget {
	if o == nil {
		var ret ObjectPrivilegeTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegePrerequisiteCheck) GetTargetOk() (*ObjectPrivilegeTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *ObjectPrivilegePrerequisiteCheck) SetTarget(v ObjectPrivilegeTarget) {
	o.Target = v
}

// GetPrivilege returns the Privilege field value.
func (o *ObjectPrivilegePrerequisiteCheck) GetPrivilege() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Privilege
}

// GetPrivilegeOk returns a tuple with the Privilege field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegePrerequisiteCheck) GetPrivilegeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privilege, true
}

// SetPrivilege sets field value.
func (o *ObjectPrivilegePrerequisiteCheck) SetPrivilege(v string) {
	o.Privilege = v
}

// GetStatus returns the Status field value.
func (o *ObjectPrivilegePrerequisiteCheck) GetStatus() ObjectPrivilegePrerequisiteStatus {
	if o == nil {
		var ret ObjectPrivilegePrerequisiteStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegePrerequisiteCheck) GetStatusOk() (*ObjectPrivilegePrerequisiteStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ObjectPrivilegePrerequisiteCheck) SetStatus(v ObjectPrivilegePrerequisiteStatus) {
	o.Status = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ObjectPrivilegePrerequisiteCheck) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegePrerequisiteCheck) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ObjectPrivilegePrerequisiteCheck) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ObjectPrivilegePrerequisiteCheck) SetReason(v string) {
	o.Reason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObjectPrivilegePrerequisiteCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["target"] = o.Target
	toSerialize["privilege"] = o.Privilege
	toSerialize["status"] = o.Status
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObjectPrivilegePrerequisiteCheck) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Target    *ObjectPrivilegeTarget             `json:"target"`
		Privilege *string                            `json:"privilege"`
		Status    *ObjectPrivilegePrerequisiteStatus `json:"status"`
		Reason    *string                            `json:"reason,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Privilege == nil {
		return fmt.Errorf("required field privilege missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"target", "privilege", "status", "reason"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = *all.Target
	o.Privilege = *all.Privilege
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Reason = all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
