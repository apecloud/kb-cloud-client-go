// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsRequestName OpsRequestName is the name of a KubeBlocks OpsRequest

type OpsRequestName struct {
	// NODESCRIPTION OpsRequestName
	OpsRequestName string `json:"opsRequestName"`
	// NODESCRIPTION DependentOpsName
	DependentOpsName *string `json:"dependentOpsName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsRequestName instantiates a new OpsRequestName object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsRequestName(opsRequestName string) *OpsRequestName {
	this := OpsRequestName{}
	this.OpsRequestName = opsRequestName
	return &this
}

// NewOpsRequestNameWithDefaults instantiates a new OpsRequestName object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsRequestNameWithDefaults() *OpsRequestName {
	this := OpsRequestName{}
	return &this
}

// GetOpsRequestName returns the OpsRequestName field value.
func (o *OpsRequestName) GetOpsRequestName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OpsRequestName
}

// GetOpsRequestNameOk returns a tuple with the OpsRequestName field value
// and a boolean to check if the value has been set.
func (o *OpsRequestName) GetOpsRequestNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpsRequestName, true
}

// SetOpsRequestName sets field value.
func (o *OpsRequestName) SetOpsRequestName(v string) {
	o.OpsRequestName = v
}

// GetDependentOpsName returns the DependentOpsName field value if set, zero value otherwise.
func (o *OpsRequestName) GetDependentOpsName() string {
	if o == nil || o.DependentOpsName == nil {
		var ret string
		return ret
	}
	return *o.DependentOpsName
}

// GetDependentOpsNameOk returns a tuple with the DependentOpsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsRequestName) GetDependentOpsNameOk() (*string, bool) {
	if o == nil || o.DependentOpsName == nil {
		return nil, false
	}
	return o.DependentOpsName, true
}

// HasDependentOpsName returns a boolean if a field has been set.
func (o *OpsRequestName) HasDependentOpsName() bool {
	return o != nil && o.DependentOpsName != nil
}

// SetDependentOpsName gets a reference to the given string and assigns it to the DependentOpsName field.
func (o *OpsRequestName) SetDependentOpsName(v string) {
	o.DependentOpsName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsRequestName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["opsRequestName"] = o.OpsRequestName
	if o.DependentOpsName != nil {
		toSerialize["dependentOpsName"] = o.DependentOpsName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsRequestName) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpsRequestName   *string `json:"opsRequestName"`
		DependentOpsName *string `json:"dependentOpsName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.OpsRequestName == nil {
		return fmt.Errorf("required field opsRequestName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"opsRequestName", "dependentOpsName"})
	} else {
		return err
	}
	o.OpsRequestName = *all.OpsRequestName
	o.DependentOpsName = all.DependentOpsName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
