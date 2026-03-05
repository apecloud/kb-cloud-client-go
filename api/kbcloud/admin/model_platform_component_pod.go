// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// PlatformComponentPod Pod summary for platform component detail
type PlatformComponentPod struct {
	// Pod name
	Name *string `json:"name,omitempty"`
	// Whether the pod is ready
	Ready *bool `json:"ready,omitempty"`
	// Pod status
	Status *string `json:"status,omitempty"`
	// Pod age (e.g. 5m, 1h)
	Age *string `json:"age,omitempty"`
	// Pod IP address
	Ip *string `json:"ip,omitempty"`
	// Node name where pod is running
	Node *string `json:"node,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPlatformComponentPod instantiates a new PlatformComponentPod object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPlatformComponentPod() *PlatformComponentPod {
	this := PlatformComponentPod{}
	return &this
}

// NewPlatformComponentPodWithDefaults instantiates a new PlatformComponentPod object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPlatformComponentPodWithDefaults() *PlatformComponentPod {
	this := PlatformComponentPod{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PlatformComponentPod) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentPod) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PlatformComponentPod) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlatformComponentPod) SetName(v string) {
	o.Name = &v
}

// GetReady returns the Ready field value if set, zero value otherwise.
func (o *PlatformComponentPod) GetReady() bool {
	if o == nil || o.Ready == nil {
		var ret bool
		return ret
	}
	return *o.Ready
}

// GetReadyOk returns a tuple with the Ready field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentPod) GetReadyOk() (*bool, bool) {
	if o == nil || o.Ready == nil {
		return nil, false
	}
	return o.Ready, true
}

// HasReady returns a boolean if a field has been set.
func (o *PlatformComponentPod) HasReady() bool {
	return o != nil && o.Ready != nil
}

// SetReady gets a reference to the given bool and assigns it to the Ready field.
func (o *PlatformComponentPod) SetReady(v bool) {
	o.Ready = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PlatformComponentPod) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentPod) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PlatformComponentPod) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PlatformComponentPod) SetStatus(v string) {
	o.Status = &v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *PlatformComponentPod) GetAge() string {
	if o == nil || o.Age == nil {
		var ret string
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentPod) GetAgeOk() (*string, bool) {
	if o == nil || o.Age == nil {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *PlatformComponentPod) HasAge() bool {
	return o != nil && o.Age != nil
}

// SetAge gets a reference to the given string and assigns it to the Age field.
func (o *PlatformComponentPod) SetAge(v string) {
	o.Age = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *PlatformComponentPod) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentPod) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *PlatformComponentPod) HasIp() bool {
	return o != nil && o.Ip != nil
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *PlatformComponentPod) SetIp(v string) {
	o.Ip = &v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *PlatformComponentPod) GetNode() string {
	if o == nil || o.Node == nil {
		var ret string
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentPod) GetNodeOk() (*string, bool) {
	if o == nil || o.Node == nil {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *PlatformComponentPod) HasNode() bool {
	return o != nil && o.Node != nil
}

// SetNode gets a reference to the given string and assigns it to the Node field.
func (o *PlatformComponentPod) SetNode(v string) {
	o.Node = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PlatformComponentPod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Ready != nil {
		toSerialize["ready"] = o.Ready
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Age != nil {
		toSerialize["age"] = o.Age
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Node != nil {
		toSerialize["node"] = o.Node
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PlatformComponentPod) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name   *string `json:"name,omitempty"`
		Ready  *bool   `json:"ready,omitempty"`
		Status *string `json:"status,omitempty"`
		Age    *string `json:"age,omitempty"`
		Ip     *string `json:"ip,omitempty"`
		Node   *string `json:"node,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "ready", "status", "age", "ip", "node"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Ready = all.Ready
	o.Status = all.Status
	o.Age = all.Age
	o.Ip = all.Ip
	o.Node = all.Node

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
