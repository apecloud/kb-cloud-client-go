// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentObjectRef struct {
	Kind      *string `json:"kind,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Name      *string `json:"name,omitempty"`
	Component *string `json:"component,omitempty"`
	Container *string `json:"container,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentObjectRef instantiates a new AiAgentObjectRef object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentObjectRef() *AiAgentObjectRef {
	this := AiAgentObjectRef{}
	return &this
}

// NewAiAgentObjectRefWithDefaults instantiates a new AiAgentObjectRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentObjectRefWithDefaults() *AiAgentObjectRef {
	this := AiAgentObjectRef{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *AiAgentObjectRef) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentObjectRef) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *AiAgentObjectRef) HasKind() bool {
	return o != nil && o.Kind != nil
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *AiAgentObjectRef) SetKind(v string) {
	o.Kind = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *AiAgentObjectRef) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentObjectRef) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *AiAgentObjectRef) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *AiAgentObjectRef) SetNamespace(v string) {
	o.Namespace = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AiAgentObjectRef) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentObjectRef) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AiAgentObjectRef) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AiAgentObjectRef) SetName(v string) {
	o.Name = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AiAgentObjectRef) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentObjectRef) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AiAgentObjectRef) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AiAgentObjectRef) SetComponent(v string) {
	o.Component = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *AiAgentObjectRef) GetContainer() string {
	if o == nil || o.Container == nil {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentObjectRef) GetContainerOk() (*string, bool) {
	if o == nil || o.Container == nil {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *AiAgentObjectRef) HasContainer() bool {
	return o != nil && o.Container != nil
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *AiAgentObjectRef) SetContainer(v string) {
	o.Container = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentObjectRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.Container != nil {
		toSerialize["container"] = o.Container
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentObjectRef) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Kind      *string `json:"kind,omitempty"`
		Namespace *string `json:"namespace,omitempty"`
		Name      *string `json:"name,omitempty"`
		Component *string `json:"component,omitempty"`
		Container *string `json:"container,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"kind", "namespace", "name", "component", "container"})
	} else {
		return err
	}
	o.Kind = all.Kind
	o.Namespace = all.Namespace
	o.Name = all.Name
	o.Component = all.Component
	o.Container = all.Container

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
