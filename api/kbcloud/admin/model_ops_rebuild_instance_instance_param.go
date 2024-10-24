// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsRebuildInstanceInstanceParam instance params for RebuildInstance ops
// NODESCRIPTION OpsRebuildInstanceInstanceParam
//
// Deprecated: This model is deprecated.
type OpsRebuildInstanceInstanceParam struct {
	// Pod name of the instance
	Name string `json:"name"`
	// The instance will rebuild on the specified node. If not set, it will rebuild on a random node.
	TargetNodeName *string `json:"targetNodeName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsRebuildInstanceInstanceParam instantiates a new OpsRebuildInstanceInstanceParam object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsRebuildInstanceInstanceParam(name string) *OpsRebuildInstanceInstanceParam {
	this := OpsRebuildInstanceInstanceParam{}
	this.Name = name
	return &this
}

// NewOpsRebuildInstanceInstanceParamWithDefaults instantiates a new OpsRebuildInstanceInstanceParam object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsRebuildInstanceInstanceParamWithDefaults() *OpsRebuildInstanceInstanceParam {
	this := OpsRebuildInstanceInstanceParam{}
	return &this
}

// GetName returns the Name field value.
func (o *OpsRebuildInstanceInstanceParam) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpsRebuildInstanceInstanceParam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OpsRebuildInstanceInstanceParam) SetName(v string) {
	o.Name = v
}

// GetTargetNodeName returns the TargetNodeName field value if set, zero value otherwise.
func (o *OpsRebuildInstanceInstanceParam) GetTargetNodeName() string {
	if o == nil || o.TargetNodeName == nil {
		var ret string
		return ret
	}
	return *o.TargetNodeName
}

// GetTargetNodeNameOk returns a tuple with the TargetNodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsRebuildInstanceInstanceParam) GetTargetNodeNameOk() (*string, bool) {
	if o == nil || o.TargetNodeName == nil {
		return nil, false
	}
	return o.TargetNodeName, true
}

// HasTargetNodeName returns a boolean if a field has been set.
func (o *OpsRebuildInstanceInstanceParam) HasTargetNodeName() bool {
	return o != nil && o.TargetNodeName != nil
}

// SetTargetNodeName gets a reference to the given string and assigns it to the TargetNodeName field.
func (o *OpsRebuildInstanceInstanceParam) SetTargetNodeName(v string) {
	o.TargetNodeName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsRebuildInstanceInstanceParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.TargetNodeName != nil {
		toSerialize["targetNodeName"] = o.TargetNodeName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsRebuildInstanceInstanceParam) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *string `json:"name"`
		TargetNodeName *string `json:"targetNodeName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "targetNodeName"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.TargetNodeName = all.TargetNodeName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
