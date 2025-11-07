// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportSupportedSource Supported data source definition for an import capability.
type ImportSupportedSource struct {
	// Import source type
	Type ImportSourceType `json:"type"`
	// Data source category
	Category *ImportSourceCategory `json:"category,omitempty"`
	// Display name of the supported source.
	Name string `json:"name"`
	// Whether the source is currently supported.
	Supported   *bool                 `json:"supported,omitempty"`
	Description *LocalizedDescription `json:"description,omitempty"`
	Guide       *LocalizedDescription `json:"guide,omitempty"`
	// Supported import capabilities.
	Capabilities []ImportCapabilityType `json:"capabilities,omitempty"`
	// Prerequisite steps for enabling the source.
	Requirements []LocalizedDescription `json:"requirements,omitempty"`
	// Required connection parameters for the source.
	ConnectionFields []ImportConnectionField `json:"connectionFields,omitempty"`
	// Compatible source version patterns.
	SupportedVersions []string `json:"supportedVersions,omitempty"`
	// Replication metadata tree definition for import object selection.
	ReplicationMetadata *ImportReplicationMetadata `json:"replicationMetadata,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewImportSupportedSource instantiates a new ImportSupportedSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImportSupportedSource(typeVar ImportSourceType, name string) *ImportSupportedSource {
	this := ImportSupportedSource{}
	this.Type = typeVar
	this.Name = name
	return &this
}

// NewImportSupportedSourceWithDefaults instantiates a new ImportSupportedSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewImportSupportedSourceWithDefaults() *ImportSupportedSource {
	this := ImportSupportedSource{}
	return &this
}

// GetType returns the Type field value.
func (o *ImportSupportedSource) GetType() ImportSourceType {
	if o == nil {
		var ret ImportSourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetTypeOk() (*ImportSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ImportSupportedSource) SetType(v ImportSourceType) {
	o.Type = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ImportSupportedSource) GetCategory() ImportSourceCategory {
	if o == nil || o.Category == nil {
		var ret ImportSourceCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetCategoryOk() (*ImportSourceCategory, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ImportSupportedSource) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given ImportSourceCategory and assigns it to the Category field.
func (o *ImportSupportedSource) SetCategory(v ImportSourceCategory) {
	o.Category = &v
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

// GetSupported returns the Supported field value if set, zero value otherwise.
func (o *ImportSupportedSource) GetSupported() bool {
	if o == nil || o.Supported == nil {
		var ret bool
		return ret
	}
	return *o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetSupportedOk() (*bool, bool) {
	if o == nil || o.Supported == nil {
		return nil, false
	}
	return o.Supported, true
}

// HasSupported returns a boolean if a field has been set.
func (o *ImportSupportedSource) HasSupported() bool {
	return o != nil && o.Supported != nil
}

// SetSupported gets a reference to the given bool and assigns it to the Supported field.
func (o *ImportSupportedSource) SetSupported(v bool) {
	o.Supported = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ImportSupportedSource) GetDescription() LocalizedDescription {
	if o == nil || o.Description == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ImportSupportedSource) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given LocalizedDescription and assigns it to the Description field.
func (o *ImportSupportedSource) SetDescription(v LocalizedDescription) {
	o.Description = &v
}

// GetGuide returns the Guide field value if set, zero value otherwise.
func (o *ImportSupportedSource) GetGuide() LocalizedDescription {
	if o == nil || o.Guide == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Guide
}

// GetGuideOk returns a tuple with the Guide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetGuideOk() (*LocalizedDescription, bool) {
	if o == nil || o.Guide == nil {
		return nil, false
	}
	return o.Guide, true
}

// HasGuide returns a boolean if a field has been set.
func (o *ImportSupportedSource) HasGuide() bool {
	return o != nil && o.Guide != nil
}

// SetGuide gets a reference to the given LocalizedDescription and assigns it to the Guide field.
func (o *ImportSupportedSource) SetGuide(v LocalizedDescription) {
	o.Guide = &v
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
func (o *ImportSupportedSource) GetRequirements() []LocalizedDescription {
	if o == nil || o.Requirements == nil {
		var ret []LocalizedDescription
		return ret
	}
	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetRequirementsOk() (*[]LocalizedDescription, bool) {
	if o == nil || o.Requirements == nil {
		return nil, false
	}
	return &o.Requirements, true
}

// HasRequirements returns a boolean if a field has been set.
func (o *ImportSupportedSource) HasRequirements() bool {
	return o != nil && o.Requirements != nil
}

// SetRequirements gets a reference to the given []LocalizedDescription and assigns it to the Requirements field.
func (o *ImportSupportedSource) SetRequirements(v []LocalizedDescription) {
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

// GetReplicationMetadata returns the ReplicationMetadata field value if set, zero value otherwise.
func (o *ImportSupportedSource) GetReplicationMetadata() ImportReplicationMetadata {
	if o == nil || o.ReplicationMetadata == nil {
		var ret ImportReplicationMetadata
		return ret
	}
	return *o.ReplicationMetadata
}

// GetReplicationMetadataOk returns a tuple with the ReplicationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSupportedSource) GetReplicationMetadataOk() (*ImportReplicationMetadata, bool) {
	if o == nil || o.ReplicationMetadata == nil {
		return nil, false
	}
	return o.ReplicationMetadata, true
}

// HasReplicationMetadata returns a boolean if a field has been set.
func (o *ImportSupportedSource) HasReplicationMetadata() bool {
	return o != nil && o.ReplicationMetadata != nil
}

// SetReplicationMetadata gets a reference to the given ImportReplicationMetadata and assigns it to the ReplicationMetadata field.
func (o *ImportSupportedSource) SetReplicationMetadata(v ImportReplicationMetadata) {
	o.ReplicationMetadata = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ImportSupportedSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	toSerialize["name"] = o.Name
	if o.Supported != nil {
		toSerialize["supported"] = o.Supported
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Guide != nil {
		toSerialize["guide"] = o.Guide
	}
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
	if o.ReplicationMetadata != nil {
		toSerialize["replicationMetadata"] = o.ReplicationMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ImportSupportedSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type                *ImportSourceType          `json:"type"`
		Category            *ImportSourceCategory      `json:"category,omitempty"`
		Name                *string                    `json:"name"`
		Supported           *bool                      `json:"supported,omitempty"`
		Description         *LocalizedDescription      `json:"description,omitempty"`
		Guide               *LocalizedDescription      `json:"guide,omitempty"`
		Capabilities        []ImportCapabilityType     `json:"capabilities,omitempty"`
		Requirements        []LocalizedDescription     `json:"requirements,omitempty"`
		ConnectionFields    []ImportConnectionField    `json:"connectionFields,omitempty"`
		SupportedVersions   []string                   `json:"supportedVersions,omitempty"`
		ReplicationMetadata *ImportReplicationMetadata `json:"replicationMetadata,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "category", "name", "supported", "description", "guide", "capabilities", "requirements", "connectionFields", "supportedVersions", "replicationMetadata"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if all.Category != nil && !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = all.Category
	}
	o.Name = *all.Name
	o.Supported = all.Supported
	if all.Description != nil && all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = all.Description
	if all.Guide != nil && all.Guide.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Guide = all.Guide
	o.Capabilities = all.Capabilities
	o.Requirements = all.Requirements
	o.ConnectionFields = all.ConnectionFields
	o.SupportedVersions = all.SupportedVersions
	if all.ReplicationMetadata != nil && all.ReplicationMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReplicationMetadata = all.ReplicationMetadata

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
