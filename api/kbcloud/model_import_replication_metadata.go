// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportReplicationMetadata Replication metadata tree definition for import object selection.
type ImportReplicationMetadata struct {
	// Metadata node type identifier (e.g. database, table).
	MetadataType string `json:"metadataType"`
	// Child metadata nodes.
	Children []ImportReplicationMetadata `json:"children,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewImportReplicationMetadata instantiates a new ImportReplicationMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImportReplicationMetadata(metadataType string) *ImportReplicationMetadata {
	this := ImportReplicationMetadata{}
	this.MetadataType = metadataType
	return &this
}

// NewImportReplicationMetadataWithDefaults instantiates a new ImportReplicationMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewImportReplicationMetadataWithDefaults() *ImportReplicationMetadata {
	this := ImportReplicationMetadata{}
	return &this
}

// GetMetadataType returns the MetadataType field value.
func (o *ImportReplicationMetadata) GetMetadataType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MetadataType
}

// GetMetadataTypeOk returns a tuple with the MetadataType field value
// and a boolean to check if the value has been set.
func (o *ImportReplicationMetadata) GetMetadataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataType, true
}

// SetMetadataType sets field value.
func (o *ImportReplicationMetadata) SetMetadataType(v string) {
	o.MetadataType = v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *ImportReplicationMetadata) GetChildren() []ImportReplicationMetadata {
	if o == nil || o.Children == nil {
		var ret []ImportReplicationMetadata
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportReplicationMetadata) GetChildrenOk() (*[]ImportReplicationMetadata, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return &o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *ImportReplicationMetadata) HasChildren() bool {
	return o != nil && o.Children != nil
}

// SetChildren gets a reference to the given []ImportReplicationMetadata and assigns it to the Children field.
func (o *ImportReplicationMetadata) SetChildren(v []ImportReplicationMetadata) {
	o.Children = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ImportReplicationMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["metadataType"] = o.MetadataType
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ImportReplicationMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MetadataType *string                     `json:"metadataType"`
		Children     []ImportReplicationMetadata `json:"children,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.MetadataType == nil {
		return fmt.Errorf("required field metadataType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"metadataType", "children"})
	} else {
		return err
	}
	o.MetadataType = *all.MetadataType
	o.Children = all.Children

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
