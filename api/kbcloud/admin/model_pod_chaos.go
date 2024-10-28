// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PodChaos struct {
	// specify the mode of chaos
	Mode PodChaosMode `json:"mode"`
	// the id of cluster to perform chaos
	ClusterId string  `json:"clusterID"`
	Value     *string `json:"value,omitempty"`
	// specify the action to be performed
	Action PodChaosAction `json:"action"`
	// specify the duration of the chaos action
	Duration *string `json:"duration,omitempty"`
	// when action is pod-kill, specify the duration to wait before killing the pod
	GracePeriod *int32 `json:"gracePeriod,omitempty"`
	// the name of the pod to perform the chaos
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPodChaos instantiates a new PodChaos object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPodChaos(mode PodChaosMode, clusterId string, action PodChaosAction) *PodChaos {
	this := PodChaos{}
	this.Mode = mode
	this.ClusterId = clusterId
	this.Action = action
	return &this
}

// NewPodChaosWithDefaults instantiates a new PodChaos object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPodChaosWithDefaults() *PodChaos {
	this := PodChaos{}
	return &this
}

// GetMode returns the Mode field value.
func (o *PodChaos) GetMode() PodChaosMode {
	if o == nil {
		var ret PodChaosMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *PodChaos) GetModeOk() (*PodChaosMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *PodChaos) SetMode(v PodChaosMode) {
	o.Mode = v
}

// GetClusterId returns the ClusterId field value.
func (o *PodChaos) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *PodChaos) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value.
func (o *PodChaos) SetClusterId(v string) {
	o.ClusterId = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PodChaos) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodChaos) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PodChaos) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PodChaos) SetValue(v string) {
	o.Value = &v
}

// GetAction returns the Action field value.
func (o *PodChaos) GetAction() PodChaosAction {
	if o == nil {
		var ret PodChaosAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *PodChaos) GetActionOk() (*PodChaosAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *PodChaos) SetAction(v PodChaosAction) {
	o.Action = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *PodChaos) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodChaos) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *PodChaos) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *PodChaos) SetDuration(v string) {
	o.Duration = &v
}

// GetGracePeriod returns the GracePeriod field value if set, zero value otherwise.
func (o *PodChaos) GetGracePeriod() int32 {
	if o == nil || o.GracePeriod == nil {
		var ret int32
		return ret
	}
	return *o.GracePeriod
}

// GetGracePeriodOk returns a tuple with the GracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodChaos) GetGracePeriodOk() (*int32, bool) {
	if o == nil || o.GracePeriod == nil {
		return nil, false
	}
	return o.GracePeriod, true
}

// HasGracePeriod returns a boolean if a field has been set.
func (o *PodChaos) HasGracePeriod() bool {
	return o != nil && o.GracePeriod != nil
}

// SetGracePeriod gets a reference to the given int32 and assigns it to the GracePeriod field.
func (o *PodChaos) SetGracePeriod(v int32) {
	o.GracePeriod = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PodChaos) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodChaos) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PodChaos) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PodChaos) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PodChaos) MarshalJSON() ([]byte, error) {
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
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.GracePeriod != nil {
		toSerialize["gracePeriod"] = o.GracePeriod
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PodChaos) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mode        *PodChaosMode   `json:"mode"`
		ClusterId   *string         `json:"clusterID"`
		Value       *string         `json:"value,omitempty"`
		Action      *PodChaosAction `json:"action"`
		Duration    *string         `json:"duration,omitempty"`
		GracePeriod *int32          `json:"gracePeriod,omitempty"`
		Name        *string         `json:"name,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"mode", "clusterID", "value", "action", "duration", "gracePeriod", "name"})
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
	o.Duration = all.Duration
	o.GracePeriod = all.GracePeriod
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
