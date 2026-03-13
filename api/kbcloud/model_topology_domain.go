// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TopologyDomain Specifies a topology domain and the set of domain values for scheduling constraints.
type TopologyDomain struct {
	// Topology level at which pod placement constraints are evaluated.
	Type TopologyType `json:"type"`
	// Domain values within the topology type. Use ["*"] to include all available domains.
	// For example, when type is Zone and values is ["*"], the rule applies across all zones.
	//
	Values []string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTopologyDomain instantiates a new TopologyDomain object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopologyDomain(typeVar TopologyType, values []string) *TopologyDomain {
	this := TopologyDomain{}
	this.Type = typeVar
	this.Values = values
	return &this
}

// NewTopologyDomainWithDefaults instantiates a new TopologyDomain object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopologyDomainWithDefaults() *TopologyDomain {
	this := TopologyDomain{}
	var typeVar TopologyType = TopologyTypeNode
	this.Type = typeVar
	return &this
}

// GetType returns the Type field value.
func (o *TopologyDomain) GetType() TopologyType {
	if o == nil {
		var ret TopologyType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TopologyDomain) GetTypeOk() (*TopologyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TopologyDomain) SetType(v TopologyType) {
	o.Type = v
}

// GetValues returns the Values field value.
func (o *TopologyDomain) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *TopologyDomain) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *TopologyDomain) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopologyDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TopologyDomain) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type   *TopologyType `json:"type"`
		Values *[]string     `json:"values"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
