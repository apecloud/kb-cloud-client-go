// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

type NodeScaleInRequest struct {
	// List of node names to be removed
	Nodes []string `json:"nodes"`
	// Whether to drain nodes before removal
	Drain    *bool            `json:"drain,omitempty"`
	Strategy *ScaleInStrategy `json:"strategy,omitempty"`
	// Force remove nodes even if they have pods not managed by ReplicationController, ReplicaSet, Job, DaemonSet or StatefulSet
	Force *bool `json:"force,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNodeScaleInRequest instantiates a new NodeScaleInRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNodeScaleInRequest(nodes []string) *NodeScaleInRequest {
	this := NodeScaleInRequest{}
	this.Nodes = nodes
	var drain bool = true
	this.Drain = &drain
	var force bool = false
	this.Force = &force
	return &this
}

// NewNodeScaleInRequestWithDefaults instantiates a new NodeScaleInRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNodeScaleInRequestWithDefaults() *NodeScaleInRequest {
	this := NodeScaleInRequest{}
	var drain bool = true
	this.Drain = &drain
	var force bool = false
	this.Force = &force
	return &this
}

// GetNodes returns the Nodes field value.
func (o *NodeScaleInRequest) GetNodes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *NodeScaleInRequest) GetNodesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value.
func (o *NodeScaleInRequest) SetNodes(v []string) {
	o.Nodes = v
}

// GetDrain returns the Drain field value if set, zero value otherwise.
func (o *NodeScaleInRequest) GetDrain() bool {
	if o == nil || o.Drain == nil {
		var ret bool
		return ret
	}
	return *o.Drain
}

// GetDrainOk returns a tuple with the Drain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeScaleInRequest) GetDrainOk() (*bool, bool) {
	if o == nil || o.Drain == nil {
		return nil, false
	}
	return o.Drain, true
}

// HasDrain returns a boolean if a field has been set.
func (o *NodeScaleInRequest) HasDrain() bool {
	return o != nil && o.Drain != nil
}

// SetDrain gets a reference to the given bool and assigns it to the Drain field.
func (o *NodeScaleInRequest) SetDrain(v bool) {
	o.Drain = &v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *NodeScaleInRequest) GetStrategy() ScaleInStrategy {
	if o == nil || o.Strategy == nil {
		var ret ScaleInStrategy
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeScaleInRequest) GetStrategyOk() (*ScaleInStrategy, bool) {
	if o == nil || o.Strategy == nil {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *NodeScaleInRequest) HasStrategy() bool {
	return o != nil && o.Strategy != nil
}

// SetStrategy gets a reference to the given ScaleInStrategy and assigns it to the Strategy field.
func (o *NodeScaleInRequest) SetStrategy(v ScaleInStrategy) {
	o.Strategy = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *NodeScaleInRequest) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeScaleInRequest) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *NodeScaleInRequest) HasForce() bool {
	return o != nil && o.Force != nil
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *NodeScaleInRequest) SetForce(v bool) {
	o.Force = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NodeScaleInRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["nodes"] = o.Nodes
	if o.Drain != nil {
		toSerialize["drain"] = o.Drain
	}
	if o.Strategy != nil {
		toSerialize["strategy"] = o.Strategy
	}
	if o.Force != nil {
		toSerialize["force"] = o.Force
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NodeScaleInRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Nodes    *[]string        `json:"nodes"`
		Drain    *bool            `json:"drain,omitempty"`
		Strategy *ScaleInStrategy `json:"strategy,omitempty"`
		Force    *bool            `json:"force,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Nodes == nil {
		return fmt.Errorf("required field nodes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nodes", "drain", "strategy", "force"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Nodes = *all.Nodes
	o.Drain = all.Drain
	if all.Strategy != nil && all.Strategy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Strategy = all.Strategy
	o.Force = all.Force

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
