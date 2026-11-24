// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DataEndpointCreate Connection configuration for an independently managed data endpoint.
type DataEndpointCreate struct {
	Name           common.NullableString     `json:"name,omitempty"`
	EngineName     *string                   `json:"engineName,omitempty"`
	DefinitionName *string                   `json:"definitionName,omitempty"`
	EndpointType   *DataEndpointTypeCreation `json:"endpointType,omitempty"`
	Custom         *CustomEndpointCreate     `json:"custom,omitempty"`
	Kubeblocks     *KubeblocksEndpoint       `json:"kubeblocks,omitempty"`
	Description    common.NullableString     `json:"description,omitempty"`
	ExtraCfgs      map[string]string         `json:"extraCfgs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataEndpointCreate instantiates a new DataEndpointCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataEndpointCreate() *DataEndpointCreate {
	this := DataEndpointCreate{}
	return &this
}

// NewDataEndpointCreateWithDefaults instantiates a new DataEndpointCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataEndpointCreateWithDefaults() *DataEndpointCreate {
	this := DataEndpointCreate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataEndpointCreate) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataEndpointCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DataEndpointCreate) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given common.NullableString and assigns it to the Name field.
func (o *DataEndpointCreate) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *DataEndpointCreate) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *DataEndpointCreate) UnsetName() {
	o.Name.Unset()
}

// GetEngineName returns the EngineName field value if set, zero value otherwise.
func (o *DataEndpointCreate) GetEngineName() string {
	if o == nil || o.EngineName == nil {
		var ret string
		return ret
	}
	return *o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataEndpointCreate) GetEngineNameOk() (*string, bool) {
	if o == nil || o.EngineName == nil {
		return nil, false
	}
	return o.EngineName, true
}

// HasEngineName returns a boolean if a field has been set.
func (o *DataEndpointCreate) HasEngineName() bool {
	return o != nil && o.EngineName != nil
}

// SetEngineName gets a reference to the given string and assigns it to the EngineName field.
func (o *DataEndpointCreate) SetEngineName(v string) {
	o.EngineName = &v
}

// GetDefinitionName returns the DefinitionName field value if set, zero value otherwise.
func (o *DataEndpointCreate) GetDefinitionName() string {
	if o == nil || o.DefinitionName == nil {
		var ret string
		return ret
	}
	return *o.DefinitionName
}

// GetDefinitionNameOk returns a tuple with the DefinitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataEndpointCreate) GetDefinitionNameOk() (*string, bool) {
	if o == nil || o.DefinitionName == nil {
		return nil, false
	}
	return o.DefinitionName, true
}

// HasDefinitionName returns a boolean if a field has been set.
func (o *DataEndpointCreate) HasDefinitionName() bool {
	return o != nil && o.DefinitionName != nil
}

// SetDefinitionName gets a reference to the given string and assigns it to the DefinitionName field.
func (o *DataEndpointCreate) SetDefinitionName(v string) {
	o.DefinitionName = &v
}

// GetEndpointType returns the EndpointType field value if set, zero value otherwise.
func (o *DataEndpointCreate) GetEndpointType() DataEndpointTypeCreation {
	if o == nil || o.EndpointType == nil {
		var ret DataEndpointTypeCreation
		return ret
	}
	return *o.EndpointType
}

// GetEndpointTypeOk returns a tuple with the EndpointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataEndpointCreate) GetEndpointTypeOk() (*DataEndpointTypeCreation, bool) {
	if o == nil || o.EndpointType == nil {
		return nil, false
	}
	return o.EndpointType, true
}

// HasEndpointType returns a boolean if a field has been set.
func (o *DataEndpointCreate) HasEndpointType() bool {
	return o != nil && o.EndpointType != nil
}

// SetEndpointType gets a reference to the given DataEndpointTypeCreation and assigns it to the EndpointType field.
func (o *DataEndpointCreate) SetEndpointType(v DataEndpointTypeCreation) {
	o.EndpointType = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *DataEndpointCreate) GetCustom() CustomEndpointCreate {
	if o == nil || o.Custom == nil {
		var ret CustomEndpointCreate
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataEndpointCreate) GetCustomOk() (*CustomEndpointCreate, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *DataEndpointCreate) HasCustom() bool {
	return o != nil && o.Custom != nil
}

// SetCustom gets a reference to the given CustomEndpointCreate and assigns it to the Custom field.
func (o *DataEndpointCreate) SetCustom(v CustomEndpointCreate) {
	o.Custom = &v
}

// GetKubeblocks returns the Kubeblocks field value if set, zero value otherwise.
func (o *DataEndpointCreate) GetKubeblocks() KubeblocksEndpoint {
	if o == nil || o.Kubeblocks == nil {
		var ret KubeblocksEndpoint
		return ret
	}
	return *o.Kubeblocks
}

// GetKubeblocksOk returns a tuple with the Kubeblocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataEndpointCreate) GetKubeblocksOk() (*KubeblocksEndpoint, bool) {
	if o == nil || o.Kubeblocks == nil {
		return nil, false
	}
	return o.Kubeblocks, true
}

// HasKubeblocks returns a boolean if a field has been set.
func (o *DataEndpointCreate) HasKubeblocks() bool {
	return o != nil && o.Kubeblocks != nil
}

// SetKubeblocks gets a reference to the given KubeblocksEndpoint and assigns it to the Kubeblocks field.
func (o *DataEndpointCreate) SetKubeblocks(v KubeblocksEndpoint) {
	o.Kubeblocks = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataEndpointCreate) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataEndpointCreate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DataEndpointCreate) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given common.NullableString and assigns it to the Description field.
func (o *DataEndpointCreate) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *DataEndpointCreate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *DataEndpointCreate) UnsetDescription() {
	o.Description.Unset()
}

// GetExtraCfgs returns the ExtraCfgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataEndpointCreate) GetExtraCfgs() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ExtraCfgs
}

// GetExtraCfgsOk returns a tuple with the ExtraCfgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataEndpointCreate) GetExtraCfgsOk() (*map[string]string, bool) {
	if o == nil || o.ExtraCfgs == nil {
		return nil, false
	}
	return &o.ExtraCfgs, true
}

// HasExtraCfgs returns a boolean if a field has been set.
func (o *DataEndpointCreate) HasExtraCfgs() bool {
	return o != nil && o.ExtraCfgs != nil
}

// SetExtraCfgs gets a reference to the given map[string]string and assigns it to the ExtraCfgs field.
func (o *DataEndpointCreate) SetExtraCfgs(v map[string]string) {
	o.ExtraCfgs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataEndpointCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.EngineName != nil {
		toSerialize["engineName"] = o.EngineName
	}
	if o.DefinitionName != nil {
		toSerialize["definitionName"] = o.DefinitionName
	}
	if o.EndpointType != nil {
		toSerialize["endpointType"] = o.EndpointType
	}
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}
	if o.Kubeblocks != nil {
		toSerialize["kubeblocks"] = o.Kubeblocks
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ExtraCfgs != nil {
		toSerialize["extraCfgs"] = o.ExtraCfgs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataEndpointCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           common.NullableString     `json:"name,omitempty"`
		EngineName     *string                   `json:"engineName,omitempty"`
		DefinitionName *string                   `json:"definitionName,omitempty"`
		EndpointType   *DataEndpointTypeCreation `json:"endpointType,omitempty"`
		Custom         *CustomEndpointCreate     `json:"custom,omitempty"`
		Kubeblocks     *KubeblocksEndpoint       `json:"kubeblocks,omitempty"`
		Description    common.NullableString     `json:"description,omitempty"`
		ExtraCfgs      map[string]string         `json:"extraCfgs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "engineName", "definitionName", "endpointType", "custom", "kubeblocks", "description", "extraCfgs"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	o.EngineName = all.EngineName
	o.DefinitionName = all.DefinitionName
	if all.EndpointType != nil && !all.EndpointType.IsValid() {
		hasInvalidField = true
	} else {
		o.EndpointType = all.EndpointType
	}
	if all.Custom != nil && all.Custom.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Custom = all.Custom
	if all.Kubeblocks != nil && all.Kubeblocks.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Kubeblocks = all.Kubeblocks
	o.Description = all.Description
	o.ExtraCfgs = all.ExtraCfgs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
