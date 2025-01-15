// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NodePoolNode Node info for environment
type NodePoolNode struct {
	Name             string `json:"name"`
	AvailabilityZone string `json:"availabilityZone"`
	MarkControlPlane *bool  `json:"markControlPlane,omitempty"`
	// Mark the node as a data plane, which will always be true when markControlPlane is not specified
	MarkDataPlane *bool `json:"markDataPlane,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNodePoolNode instantiates a new NodePoolNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNodePoolNode(name string, availabilityZone string) *NodePoolNode {
	this := NodePoolNode{}
	this.Name = name
	this.AvailabilityZone = availabilityZone
	var markControlPlane bool = false
	this.MarkControlPlane = &markControlPlane
	var markDataPlane bool = true
	this.MarkDataPlane = &markDataPlane
	return &this
}

// NewNodePoolNodeWithDefaults instantiates a new NodePoolNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNodePoolNodeWithDefaults() *NodePoolNode {
	this := NodePoolNode{}
	var markControlPlane bool = false
	this.MarkControlPlane = &markControlPlane
	var markDataPlane bool = true
	this.MarkDataPlane = &markDataPlane
	return &this
}

// GetName returns the Name field value.
func (o *NodePoolNode) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NodePoolNode) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *NodePoolNode) SetName(v string) {
	o.Name = v
}

// GetAvailabilityZone returns the AvailabilityZone field value.
func (o *NodePoolNode) GetAvailabilityZone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
func (o *NodePoolNode) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value.
func (o *NodePoolNode) SetAvailabilityZone(v string) {
	o.AvailabilityZone = v
}

// GetMarkControlPlane returns the MarkControlPlane field value if set, zero value otherwise.
func (o *NodePoolNode) GetMarkControlPlane() bool {
	if o == nil || o.MarkControlPlane == nil {
		var ret bool
		return ret
	}
	return *o.MarkControlPlane
}

// GetMarkControlPlaneOk returns a tuple with the MarkControlPlane field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePoolNode) GetMarkControlPlaneOk() (*bool, bool) {
	if o == nil || o.MarkControlPlane == nil {
		return nil, false
	}
	return o.MarkControlPlane, true
}

// HasMarkControlPlane returns a boolean if a field has been set.
func (o *NodePoolNode) HasMarkControlPlane() bool {
	return o != nil && o.MarkControlPlane != nil
}

// SetMarkControlPlane gets a reference to the given bool and assigns it to the MarkControlPlane field.
func (o *NodePoolNode) SetMarkControlPlane(v bool) {
	o.MarkControlPlane = &v
}

// GetMarkDataPlane returns the MarkDataPlane field value if set, zero value otherwise.
func (o *NodePoolNode) GetMarkDataPlane() bool {
	if o == nil || o.MarkDataPlane == nil {
		var ret bool
		return ret
	}
	return *o.MarkDataPlane
}

// GetMarkDataPlaneOk returns a tuple with the MarkDataPlane field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePoolNode) GetMarkDataPlaneOk() (*bool, bool) {
	if o == nil || o.MarkDataPlane == nil {
		return nil, false
	}
	return o.MarkDataPlane, true
}

// HasMarkDataPlane returns a boolean if a field has been set.
func (o *NodePoolNode) HasMarkDataPlane() bool {
	return o != nil && o.MarkDataPlane != nil
}

// SetMarkDataPlane gets a reference to the given bool and assigns it to the MarkDataPlane field.
func (o *NodePoolNode) SetMarkDataPlane(v bool) {
	o.MarkDataPlane = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NodePoolNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["availabilityZone"] = o.AvailabilityZone
	if o.MarkControlPlane != nil {
		toSerialize["markControlPlane"] = o.MarkControlPlane
	}
	if o.MarkDataPlane != nil {
		toSerialize["markDataPlane"] = o.MarkDataPlane
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NodePoolNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name             *string `json:"name"`
		AvailabilityZone *string `json:"availabilityZone"`
		MarkControlPlane *bool   `json:"markControlPlane,omitempty"`
		MarkDataPlane    *bool   `json:"markDataPlane,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.AvailabilityZone == nil {
		return fmt.Errorf("required field availabilityZone missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "availabilityZone", "markControlPlane", "markDataPlane"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.AvailabilityZone = *all.AvailabilityZone
	o.MarkControlPlane = all.MarkControlPlane
	o.MarkDataPlane = all.MarkDataPlane

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
