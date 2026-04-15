// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OfflineInstance Offline instance is the information of KubeBlocks cluster instances
type OfflineInstance struct {
	// Cluster name
	Cluster string `json:"cluster"`
	// Component name
	ComponentName *string `json:"componentName,omitempty"`
	// component type, refer to componentDef and support NamePrefix
	Component string `json:"component"`
	// Instance name
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOfflineInstance instantiates a new OfflineInstance object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOfflineInstance(cluster string, component string, name string) *OfflineInstance {
	this := OfflineInstance{}
	this.Cluster = cluster
	this.Component = component
	this.Name = name
	return &this
}

// NewOfflineInstanceWithDefaults instantiates a new OfflineInstance object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOfflineInstanceWithDefaults() *OfflineInstance {
	this := OfflineInstance{}
	return &this
}

// GetCluster returns the Cluster field value.
func (o *OfflineInstance) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *OfflineInstance) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value.
func (o *OfflineInstance) SetCluster(v string) {
	o.Cluster = v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *OfflineInstance) GetComponentName() string {
	if o == nil || o.ComponentName == nil {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineInstance) GetComponentNameOk() (*string, bool) {
	if o == nil || o.ComponentName == nil {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *OfflineInstance) HasComponentName() bool {
	return o != nil && o.ComponentName != nil
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *OfflineInstance) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetComponent returns the Component field value.
func (o *OfflineInstance) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *OfflineInstance) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *OfflineInstance) SetComponent(v string) {
	o.Component = v
}

// GetName returns the Name field value.
func (o *OfflineInstance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OfflineInstance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OfflineInstance) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OfflineInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["cluster"] = o.Cluster
	if o.ComponentName != nil {
		toSerialize["componentName"] = o.ComponentName
	}
	toSerialize["component"] = o.Component
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OfflineInstance) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cluster       *string `json:"cluster"`
		ComponentName *string `json:"componentName,omitempty"`
		Component     *string `json:"component"`
		Name          *string `json:"name"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Cluster == nil {
		return fmt.Errorf("required field cluster missing")
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cluster", "componentName", "component", "name"})
	} else {
		return err
	}
	o.Cluster = *all.Cluster
	o.ComponentName = all.ComponentName
	o.Component = *all.Component
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
