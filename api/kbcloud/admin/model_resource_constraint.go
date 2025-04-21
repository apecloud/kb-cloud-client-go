// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ResourceConstraint struct {
	// resource constraint id
	Id *string `json:"id,omitempty"`
	// engine name
	Engine *string `json:"engine,omitempty"`
	// engine mode
	Mode *string `json:"mode,omitempty"`
	// engine component
	Component *string         `json:"component,omitempty"`
	Replicas  *IntegerOption  `json:"replicas,omitempty"`
	Shards    *IntegerOption  `json:"shards,omitempty"`
	Volumes   []StorageOption `json:"volumes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewResourceConstraint instantiates a new ResourceConstraint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewResourceConstraint() *ResourceConstraint {
	this := ResourceConstraint{}
	return &this
}

// NewResourceConstraintWithDefaults instantiates a new ResourceConstraint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewResourceConstraintWithDefaults() *ResourceConstraint {
	this := ResourceConstraint{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceConstraint) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceConstraint) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceConstraint) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceConstraint) SetId(v string) {
	o.Id = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *ResourceConstraint) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceConstraint) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *ResourceConstraint) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *ResourceConstraint) SetEngine(v string) {
	o.Engine = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ResourceConstraint) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceConstraint) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ResourceConstraint) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ResourceConstraint) SetMode(v string) {
	o.Mode = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ResourceConstraint) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceConstraint) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ResourceConstraint) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ResourceConstraint) SetComponent(v string) {
	o.Component = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *ResourceConstraint) GetReplicas() IntegerOption {
	if o == nil || o.Replicas == nil {
		var ret IntegerOption
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceConstraint) GetReplicasOk() (*IntegerOption, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *ResourceConstraint) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given IntegerOption and assigns it to the Replicas field.
func (o *ResourceConstraint) SetReplicas(v IntegerOption) {
	o.Replicas = &v
}

// GetShards returns the Shards field value if set, zero value otherwise.
func (o *ResourceConstraint) GetShards() IntegerOption {
	if o == nil || o.Shards == nil {
		var ret IntegerOption
		return ret
	}
	return *o.Shards
}

// GetShardsOk returns a tuple with the Shards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceConstraint) GetShardsOk() (*IntegerOption, bool) {
	if o == nil || o.Shards == nil {
		return nil, false
	}
	return o.Shards, true
}

// HasShards returns a boolean if a field has been set.
func (o *ResourceConstraint) HasShards() bool {
	return o != nil && o.Shards != nil
}

// SetShards gets a reference to the given IntegerOption and assigns it to the Shards field.
func (o *ResourceConstraint) SetShards(v IntegerOption) {
	o.Shards = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *ResourceConstraint) GetVolumes() []StorageOption {
	if o == nil || o.Volumes == nil {
		var ret []StorageOption
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceConstraint) GetVolumesOk() (*[]StorageOption, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return &o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *ResourceConstraint) HasVolumes() bool {
	return o != nil && o.Volumes != nil
}

// SetVolumes gets a reference to the given []StorageOption and assigns it to the Volumes field.
func (o *ResourceConstraint) SetVolumes(v []StorageOption) {
	o.Volumes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ResourceConstraint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Shards != nil {
		toSerialize["shards"] = o.Shards
	}
	if o.Volumes != nil {
		toSerialize["volumes"] = o.Volumes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ResourceConstraint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id        *string         `json:"id,omitempty"`
		Engine    *string         `json:"engine,omitempty"`
		Mode      *string         `json:"mode,omitempty"`
		Component *string         `json:"component,omitempty"`
		Replicas  *IntegerOption  `json:"replicas,omitempty"`
		Shards    *IntegerOption  `json:"shards,omitempty"`
		Volumes   []StorageOption `json:"volumes,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "engine", "mode", "component", "replicas", "shards", "volumes"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Engine = all.Engine
	o.Mode = all.Mode
	o.Component = all.Component
	if all.Replicas != nil && all.Replicas.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Replicas = all.Replicas
	if all.Shards != nil && all.Shards.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Shards = all.Shards
	o.Volumes = all.Volumes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
