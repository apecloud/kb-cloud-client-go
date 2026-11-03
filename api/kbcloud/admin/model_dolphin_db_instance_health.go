// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DolphinDBInstanceHealth struct {
	Pod           *string               `json:"pod,omitempty"`
	Container     *string               `json:"container,omitempty"`
	NodeName      *string               `json:"nodeName,omitempty"`
	Mode          *int32                `json:"mode,omitempty"`
	Process       *DolphinDBHealthCheck `json:"process,omitempty"`
	PortReachable *DolphinDBHealthCheck `json:"portReachable,omitempty"`
	Registered    *DolphinDBHealthCheck `json:"registered,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBInstanceHealth instantiates a new DolphinDBInstanceHealth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBInstanceHealth() *DolphinDBInstanceHealth {
	this := DolphinDBInstanceHealth{}
	return &this
}

// NewDolphinDBInstanceHealthWithDefaults instantiates a new DolphinDBInstanceHealth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBInstanceHealthWithDefaults() *DolphinDBInstanceHealth {
	this := DolphinDBInstanceHealth{}
	return &this
}

// GetPod returns the Pod field value if set, zero value otherwise.
func (o *DolphinDBInstanceHealth) GetPod() string {
	if o == nil || o.Pod == nil {
		var ret string
		return ret
	}
	return *o.Pod
}

// GetPodOk returns a tuple with the Pod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBInstanceHealth) GetPodOk() (*string, bool) {
	if o == nil || o.Pod == nil {
		return nil, false
	}
	return o.Pod, true
}

// HasPod returns a boolean if a field has been set.
func (o *DolphinDBInstanceHealth) HasPod() bool {
	return o != nil && o.Pod != nil
}

// SetPod gets a reference to the given string and assigns it to the Pod field.
func (o *DolphinDBInstanceHealth) SetPod(v string) {
	o.Pod = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *DolphinDBInstanceHealth) GetContainer() string {
	if o == nil || o.Container == nil {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBInstanceHealth) GetContainerOk() (*string, bool) {
	if o == nil || o.Container == nil {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *DolphinDBInstanceHealth) HasContainer() bool {
	return o != nil && o.Container != nil
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *DolphinDBInstanceHealth) SetContainer(v string) {
	o.Container = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *DolphinDBInstanceHealth) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBInstanceHealth) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *DolphinDBInstanceHealth) HasNodeName() bool {
	return o != nil && o.NodeName != nil
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *DolphinDBInstanceHealth) SetNodeName(v string) {
	o.NodeName = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *DolphinDBInstanceHealth) GetMode() int32 {
	if o == nil || o.Mode == nil {
		var ret int32
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBInstanceHealth) GetModeOk() (*int32, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *DolphinDBInstanceHealth) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given int32 and assigns it to the Mode field.
func (o *DolphinDBInstanceHealth) SetMode(v int32) {
	o.Mode = &v
}

// GetProcess returns the Process field value if set, zero value otherwise.
func (o *DolphinDBInstanceHealth) GetProcess() DolphinDBHealthCheck {
	if o == nil || o.Process == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return *o.Process
}

// GetProcessOk returns a tuple with the Process field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBInstanceHealth) GetProcessOk() (*DolphinDBHealthCheck, bool) {
	if o == nil || o.Process == nil {
		return nil, false
	}
	return o.Process, true
}

// HasProcess returns a boolean if a field has been set.
func (o *DolphinDBInstanceHealth) HasProcess() bool {
	return o != nil && o.Process != nil
}

// SetProcess gets a reference to the given DolphinDBHealthCheck and assigns it to the Process field.
func (o *DolphinDBInstanceHealth) SetProcess(v DolphinDBHealthCheck) {
	o.Process = &v
}

// GetPortReachable returns the PortReachable field value if set, zero value otherwise.
func (o *DolphinDBInstanceHealth) GetPortReachable() DolphinDBHealthCheck {
	if o == nil || o.PortReachable == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return *o.PortReachable
}

// GetPortReachableOk returns a tuple with the PortReachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBInstanceHealth) GetPortReachableOk() (*DolphinDBHealthCheck, bool) {
	if o == nil || o.PortReachable == nil {
		return nil, false
	}
	return o.PortReachable, true
}

// HasPortReachable returns a boolean if a field has been set.
func (o *DolphinDBInstanceHealth) HasPortReachable() bool {
	return o != nil && o.PortReachable != nil
}

// SetPortReachable gets a reference to the given DolphinDBHealthCheck and assigns it to the PortReachable field.
func (o *DolphinDBInstanceHealth) SetPortReachable(v DolphinDBHealthCheck) {
	o.PortReachable = &v
}

// GetRegistered returns the Registered field value if set, zero value otherwise.
func (o *DolphinDBInstanceHealth) GetRegistered() DolphinDBHealthCheck {
	if o == nil || o.Registered == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return *o.Registered
}

// GetRegisteredOk returns a tuple with the Registered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBInstanceHealth) GetRegisteredOk() (*DolphinDBHealthCheck, bool) {
	if o == nil || o.Registered == nil {
		return nil, false
	}
	return o.Registered, true
}

// HasRegistered returns a boolean if a field has been set.
func (o *DolphinDBInstanceHealth) HasRegistered() bool {
	return o != nil && o.Registered != nil
}

// SetRegistered gets a reference to the given DolphinDBHealthCheck and assigns it to the Registered field.
func (o *DolphinDBInstanceHealth) SetRegistered(v DolphinDBHealthCheck) {
	o.Registered = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBInstanceHealth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Pod != nil {
		toSerialize["pod"] = o.Pod
	}
	if o.Container != nil {
		toSerialize["container"] = o.Container
	}
	if o.NodeName != nil {
		toSerialize["nodeName"] = o.NodeName
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Process != nil {
		toSerialize["process"] = o.Process
	}
	if o.PortReachable != nil {
		toSerialize["portReachable"] = o.PortReachable
	}
	if o.Registered != nil {
		toSerialize["registered"] = o.Registered
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBInstanceHealth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pod           *string               `json:"pod,omitempty"`
		Container     *string               `json:"container,omitempty"`
		NodeName      *string               `json:"nodeName,omitempty"`
		Mode          *int32                `json:"mode,omitempty"`
		Process       *DolphinDBHealthCheck `json:"process,omitempty"`
		PortReachable *DolphinDBHealthCheck `json:"portReachable,omitempty"`
		Registered    *DolphinDBHealthCheck `json:"registered,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"pod", "container", "nodeName", "mode", "process", "portReachable", "registered"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Pod = all.Pod
	o.Container = all.Container
	o.NodeName = all.NodeName
	o.Mode = all.Mode
	if all.Process != nil && all.Process.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Process = all.Process
	if all.PortReachable != nil && all.PortReachable.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PortReachable = all.PortReachable
	if all.Registered != nil && all.Registered.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Registered = all.Registered

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
