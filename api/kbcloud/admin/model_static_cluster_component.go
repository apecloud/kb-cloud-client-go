// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type StaticClusterComponent struct {
	// Exact component name for the selected engine. Each component may appear only once.
	Component StaticClusterComponentName `json:"component"`
	// Number of replicas. Omit to keep the installer default.
	Replicas common.NullableInt64 `json:"replicas,omitempty"`
	// Resource class for this exact engine, mode and component. Omit to keep the installer default.
	ClassCode *string `json:"classCode,omitempty"`
	// Disk capacity in Gi. Only vm-storage, vlstorage and minio accept this field.
	Storage common.NullableFloat64 `json:"storage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStaticClusterComponent instantiates a new StaticClusterComponent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStaticClusterComponent(component StaticClusterComponentName) *StaticClusterComponent {
	this := StaticClusterComponent{}
	this.Component = component
	return &this
}

// NewStaticClusterComponentWithDefaults instantiates a new StaticClusterComponent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStaticClusterComponentWithDefaults() *StaticClusterComponent {
	this := StaticClusterComponent{}
	return &this
}

// GetComponent returns the Component field value.
func (o *StaticClusterComponent) GetComponent() StaticClusterComponentName {
	if o == nil {
		var ret StaticClusterComponentName
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *StaticClusterComponent) GetComponentOk() (*StaticClusterComponentName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *StaticClusterComponent) SetComponent(v StaticClusterComponentName) {
	o.Component = v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StaticClusterComponent) GetReplicas() int64 {
	if o == nil || o.Replicas.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Replicas.Get()
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StaticClusterComponent) GetReplicasOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Replicas.Get(), o.Replicas.IsSet()
}

// HasReplicas returns a boolean if a field has been set.
func (o *StaticClusterComponent) HasReplicas() bool {
	return o != nil && o.Replicas.IsSet()
}

// SetReplicas gets a reference to the given common.NullableInt64 and assigns it to the Replicas field.
func (o *StaticClusterComponent) SetReplicas(v int64) {
	o.Replicas.Set(&v)
}

// SetReplicasNil sets the value for Replicas to be an explicit nil.
func (o *StaticClusterComponent) SetReplicasNil() {
	o.Replicas.Set(nil)
}

// UnsetReplicas ensures that no value is present for Replicas, not even an explicit nil.
func (o *StaticClusterComponent) UnsetReplicas() {
	o.Replicas.Unset()
}

// GetClassCode returns the ClassCode field value if set, zero value otherwise.
func (o *StaticClusterComponent) GetClassCode() string {
	if o == nil || o.ClassCode == nil {
		var ret string
		return ret
	}
	return *o.ClassCode
}

// GetClassCodeOk returns a tuple with the ClassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticClusterComponent) GetClassCodeOk() (*string, bool) {
	if o == nil || o.ClassCode == nil {
		return nil, false
	}
	return o.ClassCode, true
}

// HasClassCode returns a boolean if a field has been set.
func (o *StaticClusterComponent) HasClassCode() bool {
	return o != nil && o.ClassCode != nil
}

// SetClassCode gets a reference to the given string and assigns it to the ClassCode field.
func (o *StaticClusterComponent) SetClassCode(v string) {
	o.ClassCode = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StaticClusterComponent) GetStorage() float64 {
	if o == nil || o.Storage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StaticClusterComponent) GetStorageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *StaticClusterComponent) HasStorage() bool {
	return o != nil && o.Storage.IsSet()
}

// SetStorage gets a reference to the given common.NullableFloat64 and assigns it to the Storage field.
func (o *StaticClusterComponent) SetStorage(v float64) {
	o.Storage.Set(&v)
}

// SetStorageNil sets the value for Storage to be an explicit nil.
func (o *StaticClusterComponent) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil.
func (o *StaticClusterComponent) UnsetStorage() {
	o.Storage.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o StaticClusterComponent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	if o.Replicas.IsSet() {
		toSerialize["replicas"] = o.Replicas.Get()
	}
	if o.ClassCode != nil {
		toSerialize["classCode"] = o.ClassCode
	}
	if o.Storage.IsSet() {
		toSerialize["storage"] = o.Storage.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StaticClusterComponent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component *StaticClusterComponentName `json:"component"`
		Replicas  common.NullableInt64        `json:"replicas,omitempty"`
		ClassCode *string                     `json:"classCode,omitempty"`
		Storage   common.NullableFloat64      `json:"storage,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "replicas", "classCode", "storage"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Component.IsValid() {
		hasInvalidField = true
	} else {
		o.Component = *all.Component
	}
	o.Replicas = all.Replicas
	o.ClassCode = all.ClassCode
	o.Storage = all.Storage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
