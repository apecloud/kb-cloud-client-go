// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceTemplateCreate Create assignment for one instance template. name and replicas are required. classCode, storageClass, availabilityZone, env, and annotations are optional overlays; omitted fields inherit the component. Node group is cluster-level and is inherited by every component and template. Volume sizes are defined once on the component and copied onto each template.
type InstanceTemplateCreate struct {
	// Instance template name declared on the engine option.
	Name string `json:"name"`
	// Replica count for this instance template.
	Replicas int32 `json:"replicas"`
	// Class code for this template. Omit to inherit the component classCode.
	ClassCode *string `json:"classCode,omitempty"`
	// StorageClass for this template. Omit to inherit the component storageClass.
	StorageClass *string `json:"storageClass,omitempty"`
	// Availability zone for this template. Omit to inherit component scheduling (no default zone is invented).
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	// Env vars merged into this template (same-name keys overwritten).
	Env []InstanceTemplateCreateEnvItem `json:"env,omitempty"`
	// Annotations merged into this template (same-name keys overwritten).
	Annotations map[string]string `json:"annotations,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstanceTemplateCreate instantiates a new InstanceTemplateCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstanceTemplateCreate(name string, replicas int32) *InstanceTemplateCreate {
	this := InstanceTemplateCreate{}
	this.Name = name
	this.Replicas = replicas
	return &this
}

// NewInstanceTemplateCreateWithDefaults instantiates a new InstanceTemplateCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceTemplateCreateWithDefaults() *InstanceTemplateCreate {
	this := InstanceTemplateCreate{}
	return &this
}

// GetName returns the Name field value.
func (o *InstanceTemplateCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InstanceTemplateCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *InstanceTemplateCreate) SetName(v string) {
	o.Name = v
}

// GetReplicas returns the Replicas field value.
func (o *InstanceTemplateCreate) GetReplicas() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *InstanceTemplateCreate) GetReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replicas, true
}

// SetReplicas sets field value.
func (o *InstanceTemplateCreate) SetReplicas(v int32) {
	o.Replicas = v
}

// GetClassCode returns the ClassCode field value if set, zero value otherwise.
func (o *InstanceTemplateCreate) GetClassCode() string {
	if o == nil || o.ClassCode == nil {
		var ret string
		return ret
	}
	return *o.ClassCode
}

// GetClassCodeOk returns a tuple with the ClassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTemplateCreate) GetClassCodeOk() (*string, bool) {
	if o == nil || o.ClassCode == nil {
		return nil, false
	}
	return o.ClassCode, true
}

// HasClassCode returns a boolean if a field has been set.
func (o *InstanceTemplateCreate) HasClassCode() bool {
	return o != nil && o.ClassCode != nil
}

// SetClassCode gets a reference to the given string and assigns it to the ClassCode field.
func (o *InstanceTemplateCreate) SetClassCode(v string) {
	o.ClassCode = &v
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise.
func (o *InstanceTemplateCreate) GetStorageClass() string {
	if o == nil || o.StorageClass == nil {
		var ret string
		return ret
	}
	return *o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTemplateCreate) GetStorageClassOk() (*string, bool) {
	if o == nil || o.StorageClass == nil {
		return nil, false
	}
	return o.StorageClass, true
}

// HasStorageClass returns a boolean if a field has been set.
func (o *InstanceTemplateCreate) HasStorageClass() bool {
	return o != nil && o.StorageClass != nil
}

// SetStorageClass gets a reference to the given string and assigns it to the StorageClass field.
func (o *InstanceTemplateCreate) SetStorageClass(v string) {
	o.StorageClass = &v
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *InstanceTemplateCreate) GetAvailabilityZone() string {
	if o == nil || o.AvailabilityZone == nil {
		var ret string
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTemplateCreate) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil || o.AvailabilityZone == nil {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *InstanceTemplateCreate) HasAvailabilityZone() bool {
	return o != nil && o.AvailabilityZone != nil
}

// SetAvailabilityZone gets a reference to the given string and assigns it to the AvailabilityZone field.
func (o *InstanceTemplateCreate) SetAvailabilityZone(v string) {
	o.AvailabilityZone = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *InstanceTemplateCreate) GetEnv() []InstanceTemplateCreateEnvItem {
	if o == nil || o.Env == nil {
		var ret []InstanceTemplateCreateEnvItem
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTemplateCreate) GetEnvOk() (*[]InstanceTemplateCreateEnvItem, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return &o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *InstanceTemplateCreate) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given []InstanceTemplateCreateEnvItem and assigns it to the Env field.
func (o *InstanceTemplateCreate) SetEnv(v []InstanceTemplateCreateEnvItem) {
	o.Env = v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *InstanceTemplateCreate) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTemplateCreate) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *InstanceTemplateCreate) HasAnnotations() bool {
	return o != nil && o.Annotations != nil
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *InstanceTemplateCreate) SetAnnotations(v map[string]string) {
	o.Annotations = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InstanceTemplateCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["replicas"] = o.Replicas
	if o.ClassCode != nil {
		toSerialize["classCode"] = o.ClassCode
	}
	if o.StorageClass != nil {
		toSerialize["storageClass"] = o.StorageClass
	}
	if o.AvailabilityZone != nil {
		toSerialize["availabilityZone"] = o.AvailabilityZone
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceTemplateCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name             *string                         `json:"name"`
		Replicas         *int32                          `json:"replicas"`
		ClassCode        *string                         `json:"classCode,omitempty"`
		StorageClass     *string                         `json:"storageClass,omitempty"`
		AvailabilityZone *string                         `json:"availabilityZone,omitempty"`
		Env              []InstanceTemplateCreateEnvItem `json:"env,omitempty"`
		Annotations      map[string]string               `json:"annotations,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Replicas == nil {
		return fmt.Errorf("required field replicas missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "replicas", "classCode", "storageClass", "availabilityZone", "env", "annotations"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Replicas = *all.Replicas
	o.ClassCode = all.ClassCode
	o.StorageClass = all.StorageClass
	o.AvailabilityZone = all.AvailabilityZone
	o.Env = all.Env
	o.Annotations = all.Annotations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
