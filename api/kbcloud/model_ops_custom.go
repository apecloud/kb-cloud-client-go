// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsCustom OpsCustom is the payload to create a custom KubeBlocks OpsRequest
type OpsCustom struct {
	// component name
	CompName string `json:"compName"`
	// ops definition name.
	OpsType string `json:"opsType"`
	// ops definition name.
	DependentOnOps []string `json:"dependentOnOps,omitempty"`
	// custom ops parameters
	Params []interface{} `json:"params,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsCustom instantiates a new OpsCustom object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsCustom(compName string, opsType string) *OpsCustom {
	this := OpsCustom{}
	this.CompName = compName
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

// GetCompName returns the CompName field value.
func (o *OpsCustom) GetCompName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CompName
}

// GetCompNameOk returns a tuple with the CompName field value
// and a boolean to check if the value has been set.
func (o *OpsCustom) GetCompNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompName, true
}

// SetCompName sets field value.
func (o *OpsCustom) SetCompName(v string) {
	o.CompName = v
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
func (o *OpsCustom) GetParams() []interface{} {
	if o == nil || o.Params == nil {
		var ret []interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsCustom) GetParamsOk() (*[]interface{}, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return &o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *OpsCustom) HasParams() bool {
	return o != nil && o.Params != nil
}

// SetParams gets a reference to the given []interface{} and assigns it to the Params field.
func (o *OpsCustom) SetParams(v []interface{}) {
	o.Params = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsCustom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["compName"] = o.CompName
	toSerialize["opsType"] = o.OpsType
	if o.DependentOnOps != nil {
		toSerialize["dependentOnOps"] = o.DependentOnOps
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsCustom) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompName       *string       `json:"compName"`
		OpsType        *string       `json:"opsType"`
		DependentOnOps []string      `json:"dependentOnOps,omitempty"`
		Params         []interface{} `json:"params,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CompName == nil {
		return fmt.Errorf("required field compName missing")
	}
	if all.OpsType == nil {
		return fmt.Errorf("required field opsType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"compName", "opsType", "dependentOnOps", "params"})
	} else {
		return err
	}
	o.CompName = *all.CompName
	o.OpsType = *all.OpsType
	o.DependentOnOps = all.DependentOnOps
	o.Params = all.Params

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
