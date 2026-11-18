// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type PlatformComponentNode struct {
	Name *string `json:"name,omitempty"`
	// Internal node IP, falling back to external IP
	Ip *string `json:"ip,omitempty"`
	// Kubernetes node role, control-plane or worker
	Role *string `json:"role,omitempty"`
	// Ready when the NodeReady condition is True, otherwise NotReady
	Status *string `json:"status,omitempty"`
	// Names of platform components with Pods assigned to this node
	Components []string `json:"components,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPlatformComponentNode instantiates a new PlatformComponentNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPlatformComponentNode() *PlatformComponentNode {
	this := PlatformComponentNode{}
	return &this
}

// NewPlatformComponentNodeWithDefaults instantiates a new PlatformComponentNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPlatformComponentNodeWithDefaults() *PlatformComponentNode {
	this := PlatformComponentNode{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PlatformComponentNode) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentNode) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PlatformComponentNode) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlatformComponentNode) SetName(v string) {
	o.Name = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *PlatformComponentNode) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentNode) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *PlatformComponentNode) HasIp() bool {
	return o != nil && o.Ip != nil
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *PlatformComponentNode) SetIp(v string) {
	o.Ip = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *PlatformComponentNode) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentNode) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *PlatformComponentNode) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *PlatformComponentNode) SetRole(v string) {
	o.Role = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PlatformComponentNode) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentNode) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PlatformComponentNode) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PlatformComponentNode) SetStatus(v string) {
	o.Status = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *PlatformComponentNode) GetComponents() []string {
	if o == nil || o.Components == nil {
		var ret []string
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentNode) GetComponentsOk() (*[]string, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *PlatformComponentNode) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []string and assigns it to the Components field.
func (o *PlatformComponentNode) SetComponents(v []string) {
	o.Components = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PlatformComponentNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PlatformComponentNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string  `json:"name,omitempty"`
		Ip         *string  `json:"ip,omitempty"`
		Role       *string  `json:"role,omitempty"`
		Status     *string  `json:"status,omitempty"`
		Components []string `json:"components,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "ip", "role", "status", "components"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Ip = all.Ip
	o.Role = all.Role
	o.Status = all.Status
	o.Components = all.Components

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
