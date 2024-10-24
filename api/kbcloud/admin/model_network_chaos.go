// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type NetworkChaos struct {
	// Specifies the mode of network chaos to be applied. This determines how the chaos is distributed among the target pods.
	Mode NetworkChaosMode `json:"mode"`
	// the id of cluster to perform chaos
	ClusterId string `json:"clusterID"`
	// NODESCRIPTION Value
	Value *string `json:"value,omitempty"`
	// Specifies the type of network chaos action to be performed. Options include introducing delay, packet loss, duplication, corruption, network partition, or bandwidth limitation.
	Action NetworkChaosAction `json:"action"`
	// Specifies the direction of network chaos effects. 'to' affects outgoing traffic, 'from' affects incoming traffic, and 'both' affects traffic in both directions.
	Direction *NetworkChaosDirection `json:"direction,omitempty"`
	// the name of the pod to perform the chaos
	Name *string `json:"name,omitempty"`
	// specify the delay in the chaos action
	Delay *NetworkChaosDelay `json:"delay,omitempty"`
	// specify the loss in the chaos action
	Loss *NetworkChaosLoss `json:"loss,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNetworkChaos instantiates a new NetworkChaos object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNetworkChaos(mode NetworkChaosMode, clusterId string, action NetworkChaosAction) *NetworkChaos {
	this := NetworkChaos{}
	this.Mode = mode
	this.ClusterId = clusterId
	this.Action = action
	return &this
}

// NewNetworkChaosWithDefaults instantiates a new NetworkChaos object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNetworkChaosWithDefaults() *NetworkChaos {
	this := NetworkChaos{}
	return &this
}

// GetMode returns the Mode field value.
func (o *NetworkChaos) GetMode() NetworkChaosMode {
	if o == nil {
		var ret NetworkChaosMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *NetworkChaos) GetModeOk() (*NetworkChaosMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *NetworkChaos) SetMode(v NetworkChaosMode) {
	o.Mode = v
}

// GetClusterId returns the ClusterId field value.
func (o *NetworkChaos) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *NetworkChaos) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value.
func (o *NetworkChaos) SetClusterId(v string) {
	o.ClusterId = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *NetworkChaos) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkChaos) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *NetworkChaos) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *NetworkChaos) SetValue(v string) {
	o.Value = &v
}

// GetAction returns the Action field value.
func (o *NetworkChaos) GetAction() NetworkChaosAction {
	if o == nil {
		var ret NetworkChaosAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *NetworkChaos) GetActionOk() (*NetworkChaosAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *NetworkChaos) SetAction(v NetworkChaosAction) {
	o.Action = v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *NetworkChaos) GetDirection() NetworkChaosDirection {
	if o == nil || o.Direction == nil {
		var ret NetworkChaosDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkChaos) GetDirectionOk() (*NetworkChaosDirection, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *NetworkChaos) HasDirection() bool {
	return o != nil && o.Direction != nil
}

// SetDirection gets a reference to the given NetworkChaosDirection and assigns it to the Direction field.
func (o *NetworkChaos) SetDirection(v NetworkChaosDirection) {
	o.Direction = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkChaos) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkChaos) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkChaos) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkChaos) SetName(v string) {
	o.Name = &v
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *NetworkChaos) GetDelay() NetworkChaosDelay {
	if o == nil || o.Delay == nil {
		var ret NetworkChaosDelay
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkChaos) GetDelayOk() (*NetworkChaosDelay, bool) {
	if o == nil || o.Delay == nil {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *NetworkChaos) HasDelay() bool {
	return o != nil && o.Delay != nil
}

// SetDelay gets a reference to the given NetworkChaosDelay and assigns it to the Delay field.
func (o *NetworkChaos) SetDelay(v NetworkChaosDelay) {
	o.Delay = &v
}

// GetLoss returns the Loss field value if set, zero value otherwise.
func (o *NetworkChaos) GetLoss() NetworkChaosLoss {
	if o == nil || o.Loss == nil {
		var ret NetworkChaosLoss
		return ret
	}
	return *o.Loss
}

// GetLossOk returns a tuple with the Loss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkChaos) GetLossOk() (*NetworkChaosLoss, bool) {
	if o == nil || o.Loss == nil {
		return nil, false
	}
	return o.Loss, true
}

// HasLoss returns a boolean if a field has been set.
func (o *NetworkChaos) HasLoss() bool {
	return o != nil && o.Loss != nil
}

// SetLoss gets a reference to the given NetworkChaosLoss and assigns it to the Loss field.
func (o *NetworkChaos) SetLoss(v NetworkChaosLoss) {
	o.Loss = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NetworkChaos) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["mode"] = o.Mode
	toSerialize["clusterID"] = o.ClusterId
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	toSerialize["action"] = o.Action
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Delay != nil {
		toSerialize["delay"] = o.Delay
	}
	if o.Loss != nil {
		toSerialize["loss"] = o.Loss
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NetworkChaos) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mode      *NetworkChaosMode      `json:"mode"`
		ClusterId *string                `json:"clusterID"`
		Value     *string                `json:"value,omitempty"`
		Action    *NetworkChaosAction    `json:"action"`
		Direction *NetworkChaosDirection `json:"direction,omitempty"`
		Name      *string                `json:"name,omitempty"`
		Delay     *NetworkChaosDelay     `json:"delay,omitempty"`
		Loss      *NetworkChaosLoss      `json:"loss,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.ClusterId == nil {
		return fmt.Errorf("required field clusterID missing")
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"mode", "clusterID", "value", "action", "direction", "name", "delay", "loss"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
	o.ClusterId = *all.ClusterId
	o.Value = all.Value
	if !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = *all.Action
	}
	if all.Direction != nil && !all.Direction.IsValid() {
		hasInvalidField = true
	} else {
		o.Direction = all.Direction
	}
	o.Name = all.Name
	if all.Delay != nil && all.Delay.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Delay = all.Delay
	if all.Loss != nil && all.Loss.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Loss = all.Loss

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
