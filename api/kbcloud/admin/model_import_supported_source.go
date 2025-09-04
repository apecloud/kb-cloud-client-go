// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportSupportedSource Supported data source information
type ImportSupportedSource struct {
	// Data source type
	Type string `json:"type"`
	// Data source category
	Category ImportSourceCategory `json:"category"`
	// Display name
	Name string `json:"name"`
	// Description
	Description *string `json:"description,omitempty"`
	// Usage guide
	Guide *string `json:"guide,omitempty"`
	// Whether supported
	Supported bool `json:"supported"`
	// List of supported capabilities
	Capabilities []ImportCapabilityType `json:"capabilities,omitempty"`
	// List of usage requirements
	Requirements []string `json:"requirements,omitempty"`
	// Connection field configuration
	ConnectionFields []ImportConnectionField `json:"connectionFields,omitempty"`
	// List of supported versions
	SupportedVersions []string `json:"supportedVersions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewImportSupportedSource instantiates a new ImportSupportedSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImportSupportedSource(typeVar string, category ImportSourceCategory, name string, supported bool) *ImportSupportedSource {
	this := ImportSupportedSource{}
	this.Type = typeVar
	this.Category = category
	this.Name = name
	this.Supported = supported
	return &this
}

// NewImportSupportedSourceWithDefaults instantiates a new ImportSupportedSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewImportSupportedSourceWithDefaults() *ImportSupportedSource {
	this := ImportSupportedSource{}
	var supported bool = true
	this.Supported = supported
	return &this
}

// GetType returns the Type field value.
func (o *ImportSupportedSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ImportSupportedSource) SetType(v string) {
	o.Type = v
}

// GetCategory returns the Category field value.
func (o *ImportSupportedSource) GetCategory() ImportSourceCategory {
	if o == nil {
		var ret ImportSourceCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetCategoryOk() (*ImportSourceCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *ImportSupportedSource) SetCategory(v ImportSourceCategory) {
	o.Category = v
}

// GetName returns the Name field value.
func (o *ImportSupportedSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ImportSupportedSource) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ImportSupportedSource) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ImportSupportedSource) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ImportSupportedSource) SetDescription(v string) {
	o.Description = &v
}

// GetGuide returns the Guide field value if set, zero value otherwise.
func (o *ImportSupportedSource) GetGuide() string {
	if o == nil || o.Guide == nil {
		var ret string
		return ret
	}
	return *o.Guide
}

// GetGuideOk returns a tuple with the Guide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetGuideOk() (*string, bool) {
	if o == nil || o.Guide == nil {
		return nil, false
	}
	return o.Guide, true
}

// HasGuide returns a boolean if a field has been set.
func (o *ImportSupportedSource) HasGuide() bool {
	return o != nil && o.Guide != nil
}

// SetGuide gets a reference to the given string and assigns it to the Guide field.
func (o *ImportSupportedSource) SetGuide(v string) {
	o.Guide = &v
}

// GetSupported returns the Supported field value.
func (o *ImportSupportedSource) GetSupported() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetSupportedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supported, true
}

// SetSupported sets field value.
func (o *ImportSupportedSource) SetSupported(v bool) {
	o.Supported = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *ImportSupportedSource) GetCapabilities() []ImportCapabilityType {
	if o == nil || o.Capabilities == nil {
		var ret []ImportCapabilityType
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetCapabilitiesOk() (*[]ImportCapabilityType, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return &o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *ImportSupportedSource) HasCapabilities() bool {
	return o != nil && o.Capabilities != nil
}

// SetCapabilities gets a reference to the given []ImportCapabilityType and assigns it to the Capabilities field.
func (o *ImportSupportedSource) SetCapabilities(v []ImportCapabilityType) {
	o.Capabilities = v
}

// GetRequirements returns the Requirements field value if set, zero value otherwise.
func (o *ImportSupportedSource) GetRequirements() []string {
	if o == nil || o.Requirements == nil {
		var ret []string
		return ret
	}
	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetRequirementsOk() (*[]string, bool) {
	if o == nil || o.Requirements == nil {
		return nil, false
	}
	return &o.Requirements, true
}

// HasRequirements returns a boolean if a field has been set.
func (o *ImportSupportedSource) HasRequirements() bool {
	return o != nil && o.Requirements != nil
}

// SetRequirements gets a reference to the given []string and assigns it to the Requirements field.
func (o *ImportSupportedSource) SetRequirements(v []string) {
	o.Requirements = v
}

// GetConnectionFields returns the ConnectionFields field value if set, zero value otherwise.
func (o *ImportSupportedSource) GetConnectionFields() []ImportConnectionField {
	if o == nil || o.ConnectionFields == nil {
		var ret []ImportConnectionField
		return ret
	}
	return o.ConnectionFields
}

// GetConnectionFieldsOk returns a tuple with the ConnectionFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetConnectionFieldsOk() (*[]ImportConnectionField, bool) {
	if o == nil || o.ConnectionFields == nil {
		return nil, false
	}
	return &o.ConnectionFields, true
}

// HasConnectionFields returns a boolean if a field has been set.
func (o *ImportSupportedSource) HasConnectionFields() bool {
	return o != nil && o.ConnectionFields != nil
}

// SetConnectionFields gets a reference to the given []ImportConnectionField and assigns it to the ConnectionFields field.
func (o *ImportSupportedSource) SetConnectionFields(v []ImportConnectionField) {
	o.ConnectionFields = v
}

// GetSupportedVersions returns the SupportedVersions field value if set, zero value otherwise.
func (o *ImportSupportedSource) GetSupportedVersions() []string {
	if o == nil || o.SupportedVersions == nil {
		var ret []string
		return ret
	}
	return o.SupportedVersions
}

// GetSupportedVersionsOk returns a tuple with the SupportedVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetSupportedVersionsOk() (*[]string, bool) {
	if o == nil || o.SupportedVersions == nil {
		return nil, false
	}
	return &o.SupportedVersions, true
}

// HasSupportedVersions returns a boolean if a field has been set.
func (o *ImportSupportedSource) HasSupportedVersions() bool {
	return o != nil && o.SupportedVersions != nil
}

// SetSupportedVersions gets a reference to the given []string and assigns it to the SupportedVersions field.
func (o *ImportSupportedSource) SetSupportedVersions(v []string) {
	o.SupportedVersions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ImportSupportedSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["category"] = o.Category
	toSerialize["name"] = o.Name
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Guide != nil {
		toSerialize["guide"] = o.Guide
	}
	toSerialize["supported"] = o.Supported
	if o.Capabilities != nil {
		toSerialize["capabilities"] = o.Capabilities
	}
	if o.Requirements != nil {
		toSerialize["requirements"] = o.Requirements
	}
	if o.ConnectionFields != nil {
		toSerialize["connectionFields"] = o.ConnectionFields
	}
	if o.SupportedVersions != nil {
		toSerialize["supportedVersions"] = o.SupportedVersions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ImportSupportedSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type              *string                 `json:"type"`
		Category          *ImportSourceCategory   `json:"category"`
		Name              *string                 `json:"name"`
		Description       *string                 `json:"description,omitempty"`
		Guide             *string                 `json:"guide,omitempty"`
		Supported         *bool                   `json:"supported"`
		Capabilities      []ImportCapabilityType  `json:"capabilities,omitempty"`
		Requirements      []string                `json:"requirements,omitempty"`
		ConnectionFields  []ImportConnectionField `json:"connectionFields,omitempty"`
		SupportedVersions []string                `json:"supportedVersions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Supported == nil {
		return fmt.Errorf("required field supported missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "category", "name", "description", "guide", "supported", "capabilities", "requirements", "connectionFields", "supportedVersions"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Type = *all.Type
	if !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = *all.Category
	}
	o.Name = *all.Name
	o.Description = all.Description
	o.Guide = all.Guide
	o.Supported = *all.Supported
	o.Capabilities = all.Capabilities
	o.Requirements = all.Requirements
	o.ConnectionFields = all.ConnectionFields
	o.SupportedVersions = all.SupportedVersions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
