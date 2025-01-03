// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ComponentPodsPodsItem struct {
	// Pod name
	Name *string `json:"name,omitempty"`
	// Pod namespace
	Namespace *string `json:"namespace,omitempty"`
	// Node name where pod is running
	NodeName *string `json:"node_name,omitempty"`
	// Pod status
	Status *string `json:"status,omitempty"`
	// Pod phase
	Phase *string `json:"phase,omitempty"`
	// Pod IP address
	Ip *string `json:"ip,omitempty"`
	// Pod creation time
	CreationTimestamp *time.Time `json:"creation_timestamp,omitempty"`
	// Resource usage (only for kb-cluster type)
	Resources       *ComponentPodsPodsItemResources            `json:"resources,omitempty"`
	OwnerReferences []ComponentPodsPodsItemOwnerReferencesItem `json:"owner_references,omitempty"`
	Containers      []ComponentPodsPodsItemContainersItem      `json:"containers,omitempty"`
	Conditions      []ComponentPodsPodsItemConditionsItem      `json:"conditions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentPodsPodsItem instantiates a new ComponentPodsPodsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentPodsPodsItem() *ComponentPodsPodsItem {
	this := ComponentPodsPodsItem{}
	return &this
}

// NewComponentPodsPodsItemWithDefaults instantiates a new ComponentPodsPodsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentPodsPodsItemWithDefaults() *ComponentPodsPodsItem {
	this := ComponentPodsPodsItem{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComponentPodsPodsItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPodsPodsItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComponentPodsPodsItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComponentPodsPodsItem) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ComponentPodsPodsItem) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPodsPodsItem) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ComponentPodsPodsItem) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ComponentPodsPodsItem) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *ComponentPodsPodsItem) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPodsPodsItem) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *ComponentPodsPodsItem) HasNodeName() bool {
	return o != nil && o.NodeName != nil
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *ComponentPodsPodsItem) SetNodeName(v string) {
	o.NodeName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ComponentPodsPodsItem) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPodsPodsItem) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ComponentPodsPodsItem) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ComponentPodsPodsItem) SetStatus(v string) {
	o.Status = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *ComponentPodsPodsItem) GetPhase() string {
	if o == nil || o.Phase == nil {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPodsPodsItem) GetPhaseOk() (*string, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *ComponentPodsPodsItem) HasPhase() bool {
	return o != nil && o.Phase != nil
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *ComponentPodsPodsItem) SetPhase(v string) {
	o.Phase = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *ComponentPodsPodsItem) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPodsPodsItem) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *ComponentPodsPodsItem) HasIp() bool {
	return o != nil && o.Ip != nil
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *ComponentPodsPodsItem) SetIp(v string) {
	o.Ip = &v
}

// GetCreationTimestamp returns the CreationTimestamp field value if set, zero value otherwise.
func (o *ComponentPodsPodsItem) GetCreationTimestamp() time.Time {
	if o == nil || o.CreationTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPodsPodsItem) GetCreationTimestampOk() (*time.Time, bool) {
	if o == nil || o.CreationTimestamp == nil {
		return nil, false
	}
	return o.CreationTimestamp, true
}

// HasCreationTimestamp returns a boolean if a field has been set.
func (o *ComponentPodsPodsItem) HasCreationTimestamp() bool {
	return o != nil && o.CreationTimestamp != nil
}

// SetCreationTimestamp gets a reference to the given time.Time and assigns it to the CreationTimestamp field.
func (o *ComponentPodsPodsItem) SetCreationTimestamp(v time.Time) {
	o.CreationTimestamp = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ComponentPodsPodsItem) GetResources() ComponentPodsPodsItemResources {
	if o == nil || o.Resources == nil {
		var ret ComponentPodsPodsItemResources
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPodsPodsItem) GetResourcesOk() (*ComponentPodsPodsItemResources, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ComponentPodsPodsItem) HasResources() bool {
	return o != nil && o.Resources != nil
}

// SetResources gets a reference to the given ComponentPodsPodsItemResources and assigns it to the Resources field.
func (o *ComponentPodsPodsItem) SetResources(v ComponentPodsPodsItemResources) {
	o.Resources = &v
}

// GetOwnerReferences returns the OwnerReferences field value if set, zero value otherwise.
func (o *ComponentPodsPodsItem) GetOwnerReferences() []ComponentPodsPodsItemOwnerReferencesItem {
	if o == nil || o.OwnerReferences == nil {
		var ret []ComponentPodsPodsItemOwnerReferencesItem
		return ret
	}
	return o.OwnerReferences
}

// GetOwnerReferencesOk returns a tuple with the OwnerReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPodsPodsItem) GetOwnerReferencesOk() (*[]ComponentPodsPodsItemOwnerReferencesItem, bool) {
	if o == nil || o.OwnerReferences == nil {
		return nil, false
	}
	return &o.OwnerReferences, true
}

// HasOwnerReferences returns a boolean if a field has been set.
func (o *ComponentPodsPodsItem) HasOwnerReferences() bool {
	return o != nil && o.OwnerReferences != nil
}

// SetOwnerReferences gets a reference to the given []ComponentPodsPodsItemOwnerReferencesItem and assigns it to the OwnerReferences field.
func (o *ComponentPodsPodsItem) SetOwnerReferences(v []ComponentPodsPodsItemOwnerReferencesItem) {
	o.OwnerReferences = v
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *ComponentPodsPodsItem) GetContainers() []ComponentPodsPodsItemContainersItem {
	if o == nil || o.Containers == nil {
		var ret []ComponentPodsPodsItemContainersItem
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPodsPodsItem) GetContainersOk() (*[]ComponentPodsPodsItemContainersItem, bool) {
	if o == nil || o.Containers == nil {
		return nil, false
	}
	return &o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *ComponentPodsPodsItem) HasContainers() bool {
	return o != nil && o.Containers != nil
}

// SetContainers gets a reference to the given []ComponentPodsPodsItemContainersItem and assigns it to the Containers field.
func (o *ComponentPodsPodsItem) SetContainers(v []ComponentPodsPodsItemContainersItem) {
	o.Containers = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ComponentPodsPodsItem) GetConditions() []ComponentPodsPodsItemConditionsItem {
	if o == nil || o.Conditions == nil {
		var ret []ComponentPodsPodsItemConditionsItem
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPodsPodsItem) GetConditionsOk() (*[]ComponentPodsPodsItemConditionsItem, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ComponentPodsPodsItem) HasConditions() bool {
	return o != nil && o.Conditions != nil
}

// SetConditions gets a reference to the given []ComponentPodsPodsItemConditionsItem and assigns it to the Conditions field.
func (o *ComponentPodsPodsItem) SetConditions(v []ComponentPodsPodsItemConditionsItem) {
	o.Conditions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentPodsPodsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.NodeName != nil {
		toSerialize["node_name"] = o.NodeName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Phase != nil {
		toSerialize["phase"] = o.Phase
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.CreationTimestamp != nil {
		if o.CreationTimestamp.Nanosecond() == 0 {
			toSerialize["creation_timestamp"] = o.CreationTimestamp.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["creation_timestamp"] = o.CreationTimestamp.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.OwnerReferences != nil {
		toSerialize["owner_references"] = o.OwnerReferences
	}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentPodsPodsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name              *string                                    `json:"name,omitempty"`
		Namespace         *string                                    `json:"namespace,omitempty"`
		NodeName          *string                                    `json:"node_name,omitempty"`
		Status            *string                                    `json:"status,omitempty"`
		Phase             *string                                    `json:"phase,omitempty"`
		Ip                *string                                    `json:"ip,omitempty"`
		CreationTimestamp *time.Time                                 `json:"creation_timestamp,omitempty"`
		Resources         *ComponentPodsPodsItemResources            `json:"resources,omitempty"`
		OwnerReferences   []ComponentPodsPodsItemOwnerReferencesItem `json:"owner_references,omitempty"`
		Containers        []ComponentPodsPodsItemContainersItem      `json:"containers,omitempty"`
		Conditions        []ComponentPodsPodsItemConditionsItem      `json:"conditions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "namespace", "node_name", "status", "phase", "ip", "creation_timestamp", "resources", "owner_references", "containers", "conditions"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	o.Namespace = all.Namespace
	o.NodeName = all.NodeName
	o.Status = all.Status
	o.Phase = all.Phase
	o.Ip = all.Ip
	o.CreationTimestamp = all.CreationTimestamp
	if all.Resources != nil && all.Resources.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Resources = all.Resources
	o.OwnerReferences = all.OwnerReferences
	o.Containers = all.Containers
	o.Conditions = all.Conditions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
