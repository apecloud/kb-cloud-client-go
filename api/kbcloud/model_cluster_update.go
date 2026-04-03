// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterUpdate ClusterUpdate is the payload to update a KubeBlocks cluster
type ClusterUpdate struct {
	// Components is the list of components
	Components []ComponentItemCreate `json:"components,omitempty"`
	// The termination policy of cluster.
	TerminationPolicy *ClusterTerminationPolicy `json:"terminationPolicy,omitempty"`
	// Display name of cluster.
	DisplayName common.NullableString `json:"displayName,omitempty"`
	// the maintenance window for a cluster
	MaintainceWindow *ClusterMaintainceWindow `json:"maintainceWindow,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterUpdate instantiates a new ClusterUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterUpdate() *ClusterUpdate {
	this := ClusterUpdate{}
	var terminationPolicy ClusterTerminationPolicy = ClusterTerminationPolicyDelete
	this.TerminationPolicy = &terminationPolicy
	return &this
}

// NewClusterUpdateWithDefaults instantiates a new ClusterUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterUpdateWithDefaults() *ClusterUpdate {
	this := ClusterUpdate{}
	var terminationPolicy ClusterTerminationPolicy = ClusterTerminationPolicyDelete
	this.TerminationPolicy = &terminationPolicy
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *ClusterUpdate) GetComponents() []ComponentItemCreate {
	if o == nil || o.Components == nil {
		var ret []ComponentItemCreate
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetComponentsOk() (*[]ComponentItemCreate, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *ClusterUpdate) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []ComponentItemCreate and assigns it to the Components field.
func (o *ClusterUpdate) SetComponents(v []ComponentItemCreate) {
	o.Components = v
}

// GetTerminationPolicy returns the TerminationPolicy field value if set, zero value otherwise.
func (o *ClusterUpdate) GetTerminationPolicy() ClusterTerminationPolicy {
	if o == nil || o.TerminationPolicy == nil {
		var ret ClusterTerminationPolicy
		return ret
	}
	return *o.TerminationPolicy
}

// GetTerminationPolicyOk returns a tuple with the TerminationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetTerminationPolicyOk() (*ClusterTerminationPolicy, bool) {
	if o == nil || o.TerminationPolicy == nil {
		return nil, false
	}
	return o.TerminationPolicy, true
}

// HasTerminationPolicy returns a boolean if a field has been set.
func (o *ClusterUpdate) HasTerminationPolicy() bool {
	return o != nil && o.TerminationPolicy != nil
}

// SetTerminationPolicy gets a reference to the given ClusterTerminationPolicy and assigns it to the TerminationPolicy field.
func (o *ClusterUpdate) SetTerminationPolicy(v ClusterTerminationPolicy) {
	o.TerminationPolicy = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterUpdate) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterUpdate) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ClusterUpdate) HasDisplayName() bool {
	return o != nil && o.DisplayName.IsSet()
}

// SetDisplayName gets a reference to the given common.NullableString and assigns it to the DisplayName field.
func (o *ClusterUpdate) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// SetDisplayNameNil sets the value for DisplayName to be an explicit nil.
func (o *ClusterUpdate) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil.
func (o *ClusterUpdate) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetMaintainceWindow returns the MaintainceWindow field value if set, zero value otherwise.
func (o *ClusterUpdate) GetMaintainceWindow() ClusterMaintainceWindow {
	if o == nil || o.MaintainceWindow == nil {
		var ret ClusterMaintainceWindow
		return ret
	}
	return *o.MaintainceWindow
}

// GetMaintainceWindowOk returns a tuple with the MaintainceWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetMaintainceWindowOk() (*ClusterMaintainceWindow, bool) {
	if o == nil || o.MaintainceWindow == nil {
		return nil, false
	}
	return o.MaintainceWindow, true
}

// HasMaintainceWindow returns a boolean if a field has been set.
func (o *ClusterUpdate) HasMaintainceWindow() bool {
	return o != nil && o.MaintainceWindow != nil
}

// SetMaintainceWindow gets a reference to the given ClusterMaintainceWindow and assigns it to the MaintainceWindow field.
func (o *ClusterUpdate) SetMaintainceWindow(v ClusterMaintainceWindow) {
	o.MaintainceWindow = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.TerminationPolicy != nil {
		toSerialize["terminationPolicy"] = o.TerminationPolicy
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.MaintainceWindow != nil {
		toSerialize["maintainceWindow"] = o.MaintainceWindow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Components        []ComponentItemCreate     `json:"components,omitempty"`
		TerminationPolicy *ClusterTerminationPolicy `json:"terminationPolicy,omitempty"`
		DisplayName       common.NullableString     `json:"displayName,omitempty"`
		MaintainceWindow  *ClusterMaintainceWindow  `json:"maintainceWindow,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"components", "terminationPolicy", "displayName", "maintainceWindow"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Components = all.Components
	if all.TerminationPolicy != nil && !all.TerminationPolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.TerminationPolicy = all.TerminationPolicy
	}
	o.DisplayName = all.DisplayName
	if all.MaintainceWindow != nil && all.MaintainceWindow.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MaintainceWindow = all.MaintainceWindow

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
