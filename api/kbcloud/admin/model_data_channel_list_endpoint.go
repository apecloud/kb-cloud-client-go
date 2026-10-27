// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataChannelListEndpoint struct {
	Id           common.NullableString        `json:"id,omitempty"`
	Name         common.NullableString        `json:"name,omitempty"`
	EngineName   *string                      `json:"engineName,omitempty"`
	EndpointType *DataReplicationEndpointType `json:"endpointType,omitempty"`
	Custom       *CustomEndpoint              `json:"custom,omitempty"`
	Kubeblocks   *KubeblocksEndpoint          `json:"kubeblocks,omitempty"`
	// Reference to an existing data endpoint in the same organization.
	Definition  *DefinitionEndpoint   `json:"definition,omitempty"`
	Description common.NullableString `json:"description,omitempty"`
	ExtraCfgs   map[string]string     `json:"extraCfgs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataChannelListEndpoint instantiates a new DataChannelListEndpoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataChannelListEndpoint() *DataChannelListEndpoint {
	this := DataChannelListEndpoint{}
	return &this
}

// NewDataChannelListEndpointWithDefaults instantiates a new DataChannelListEndpoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataChannelListEndpointWithDefaults() *DataChannelListEndpoint {
	this := DataChannelListEndpoint{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataChannelListEndpoint) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataChannelListEndpoint) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DataChannelListEndpoint) HasId() bool {
	return o != nil && o.Id.IsSet()
}

// SetId gets a reference to the given common.NullableString and assigns it to the Id field.
func (o *DataChannelListEndpoint) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil.
func (o *DataChannelListEndpoint) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil.
func (o *DataChannelListEndpoint) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataChannelListEndpoint) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataChannelListEndpoint) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DataChannelListEndpoint) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given common.NullableString and assigns it to the Name field.
func (o *DataChannelListEndpoint) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *DataChannelListEndpoint) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *DataChannelListEndpoint) UnsetName() {
	o.Name.Unset()
}

// GetEngineName returns the EngineName field value if set, zero value otherwise.
func (o *DataChannelListEndpoint) GetEngineName() string {
	if o == nil || o.EngineName == nil {
		var ret string
		return ret
	}
	return *o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelListEndpoint) GetEngineNameOk() (*string, bool) {
	if o == nil || o.EngineName == nil {
		return nil, false
	}
	return o.EngineName, true
}

// HasEngineName returns a boolean if a field has been set.
func (o *DataChannelListEndpoint) HasEngineName() bool {
	return o != nil && o.EngineName != nil
}

// SetEngineName gets a reference to the given string and assigns it to the EngineName field.
func (o *DataChannelListEndpoint) SetEngineName(v string) {
	o.EngineName = &v
}

// GetEndpointType returns the EndpointType field value if set, zero value otherwise.
func (o *DataChannelListEndpoint) GetEndpointType() DataReplicationEndpointType {
	if o == nil || o.EndpointType == nil {
		var ret DataReplicationEndpointType
		return ret
	}
	return *o.EndpointType
}

// GetEndpointTypeOk returns a tuple with the EndpointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelListEndpoint) GetEndpointTypeOk() (*DataReplicationEndpointType, bool) {
	if o == nil || o.EndpointType == nil {
		return nil, false
	}
	return o.EndpointType, true
}

// HasEndpointType returns a boolean if a field has been set.
func (o *DataChannelListEndpoint) HasEndpointType() bool {
	return o != nil && o.EndpointType != nil
}

// SetEndpointType gets a reference to the given DataReplicationEndpointType and assigns it to the EndpointType field.
func (o *DataChannelListEndpoint) SetEndpointType(v DataReplicationEndpointType) {
	o.EndpointType = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *DataChannelListEndpoint) GetCustom() CustomEndpoint {
	if o == nil || o.Custom == nil {
		var ret CustomEndpoint
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelListEndpoint) GetCustomOk() (*CustomEndpoint, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *DataChannelListEndpoint) HasCustom() bool {
	return o != nil && o.Custom != nil
}

// SetCustom gets a reference to the given CustomEndpoint and assigns it to the Custom field.
func (o *DataChannelListEndpoint) SetCustom(v CustomEndpoint) {
	o.Custom = &v
}

// GetKubeblocks returns the Kubeblocks field value if set, zero value otherwise.
func (o *DataChannelListEndpoint) GetKubeblocks() KubeblocksEndpoint {
	if o == nil || o.Kubeblocks == nil {
		var ret KubeblocksEndpoint
		return ret
	}
	return *o.Kubeblocks
}

// GetKubeblocksOk returns a tuple with the Kubeblocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelListEndpoint) GetKubeblocksOk() (*KubeblocksEndpoint, bool) {
	if o == nil || o.Kubeblocks == nil {
		return nil, false
	}
	return o.Kubeblocks, true
}

// HasKubeblocks returns a boolean if a field has been set.
func (o *DataChannelListEndpoint) HasKubeblocks() bool {
	return o != nil && o.Kubeblocks != nil
}

// SetKubeblocks gets a reference to the given KubeblocksEndpoint and assigns it to the Kubeblocks field.
func (o *DataChannelListEndpoint) SetKubeblocks(v KubeblocksEndpoint) {
	o.Kubeblocks = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *DataChannelListEndpoint) GetDefinition() DefinitionEndpoint {
	if o == nil || o.Definition == nil {
		var ret DefinitionEndpoint
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelListEndpoint) GetDefinitionOk() (*DefinitionEndpoint, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *DataChannelListEndpoint) HasDefinition() bool {
	return o != nil && o.Definition != nil
}

// SetDefinition gets a reference to the given DefinitionEndpoint and assigns it to the Definition field.
func (o *DataChannelListEndpoint) SetDefinition(v DefinitionEndpoint) {
	o.Definition = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataChannelListEndpoint) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataChannelListEndpoint) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DataChannelListEndpoint) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given common.NullableString and assigns it to the Description field.
func (o *DataChannelListEndpoint) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *DataChannelListEndpoint) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *DataChannelListEndpoint) UnsetDescription() {
	o.Description.Unset()
}

// GetExtraCfgs returns the ExtraCfgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataChannelListEndpoint) GetExtraCfgs() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ExtraCfgs
}

// GetExtraCfgsOk returns a tuple with the ExtraCfgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataChannelListEndpoint) GetExtraCfgsOk() (*map[string]string, bool) {
	if o == nil || o.ExtraCfgs == nil {
		return nil, false
	}
	return &o.ExtraCfgs, true
}

// HasExtraCfgs returns a boolean if a field has been set.
func (o *DataChannelListEndpoint) HasExtraCfgs() bool {
	return o != nil && o.ExtraCfgs != nil
}

// SetExtraCfgs gets a reference to the given map[string]string and assigns it to the ExtraCfgs field.
func (o *DataChannelListEndpoint) SetExtraCfgs(v map[string]string) {
	o.ExtraCfgs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataChannelListEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.EngineName != nil {
		toSerialize["engineName"] = o.EngineName
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
	if o.Definition != nil {
		toSerialize["definition"] = o.Definition
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
func (o *DataChannelListEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id           common.NullableString        `json:"id,omitempty"`
		Name         common.NullableString        `json:"name,omitempty"`
		EngineName   *string                      `json:"engineName,omitempty"`
		EndpointType *DataReplicationEndpointType `json:"endpointType,omitempty"`
		Custom       *CustomEndpoint              `json:"custom,omitempty"`
		Kubeblocks   *KubeblocksEndpoint          `json:"kubeblocks,omitempty"`
		Definition   *DefinitionEndpoint          `json:"definition,omitempty"`
		Description  common.NullableString        `json:"description,omitempty"`
		ExtraCfgs    map[string]string            `json:"extraCfgs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "engineName", "endpointType", "custom", "kubeblocks", "definition", "description", "extraCfgs"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Name = all.Name
	o.EngineName = all.EngineName
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
	if all.Definition != nil && all.Definition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Definition = all.Definition
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
