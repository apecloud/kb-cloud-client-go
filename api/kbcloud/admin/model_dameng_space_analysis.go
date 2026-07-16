// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DamengSpaceAnalysis struct {
	// Tablespace size and usage information from dba_free_space and dba_data_files.
	Tablespaces []DamengTablespaceItem `json:"tablespaces"`
	// Schema-level size aggregation from DBA_SEGMENTS grouped by OWNER.
	Schemas []DamengSchemaItem `json:"schemas,omitempty"`
	// Top-N tables ordered by total size from DBA_SEGMENTS (SEGMENT_TYPE='TABLE').
	Tables []DamengTableSizeItem `json:"tables,omitempty"`
	// Top-N indexes ordered by size from DBA_INDEXES joined with DBA_SEGMENTS.
	Indexes []DamengIndexItem `json:"indexes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDamengSpaceAnalysis instantiates a new DamengSpaceAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDamengSpaceAnalysis(tablespaces []DamengTablespaceItem) *DamengSpaceAnalysis {
	this := DamengSpaceAnalysis{}
	this.Tablespaces = tablespaces
	return &this
}

// NewDamengSpaceAnalysisWithDefaults instantiates a new DamengSpaceAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDamengSpaceAnalysisWithDefaults() *DamengSpaceAnalysis {
	this := DamengSpaceAnalysis{}
	return &this
}

// GetTablespaces returns the Tablespaces field value.
func (o *DamengSpaceAnalysis) GetTablespaces() []DamengTablespaceItem {
	if o == nil {
		var ret []DamengTablespaceItem
		return ret
	}
	return o.Tablespaces
}

// GetTablespacesOk returns a tuple with the Tablespaces field value
// and a boolean to check if the value has been set.
func (o *DamengSpaceAnalysis) GetTablespacesOk() (*[]DamengTablespaceItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tablespaces, true
}

// SetTablespaces sets field value.
func (o *DamengSpaceAnalysis) SetTablespaces(v []DamengTablespaceItem) {
	o.Tablespaces = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *DamengSpaceAnalysis) GetSchemas() []DamengSchemaItem {
	if o == nil || o.Schemas == nil {
		var ret []DamengSchemaItem
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSpaceAnalysis) GetSchemasOk() (*[]DamengSchemaItem, bool) {
	if o == nil || o.Schemas == nil {
		return nil, false
	}
	return &o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *DamengSpaceAnalysis) HasSchemas() bool {
	return o != nil && o.Schemas != nil
}

// SetSchemas gets a reference to the given []DamengSchemaItem and assigns it to the Schemas field.
func (o *DamengSpaceAnalysis) SetSchemas(v []DamengSchemaItem) {
	o.Schemas = v
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *DamengSpaceAnalysis) GetTables() []DamengTableSizeItem {
	if o == nil || o.Tables == nil {
		var ret []DamengTableSizeItem
		return ret
	}
	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSpaceAnalysis) GetTablesOk() (*[]DamengTableSizeItem, bool) {
	if o == nil || o.Tables == nil {
		return nil, false
	}
	return &o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *DamengSpaceAnalysis) HasTables() bool {
	return o != nil && o.Tables != nil
}

// SetTables gets a reference to the given []DamengTableSizeItem and assigns it to the Tables field.
func (o *DamengSpaceAnalysis) SetTables(v []DamengTableSizeItem) {
	o.Tables = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *DamengSpaceAnalysis) GetIndexes() []DamengIndexItem {
	if o == nil || o.Indexes == nil {
		var ret []DamengIndexItem
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSpaceAnalysis) GetIndexesOk() (*[]DamengIndexItem, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *DamengSpaceAnalysis) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []DamengIndexItem and assigns it to the Indexes field.
func (o *DamengSpaceAnalysis) SetIndexes(v []DamengIndexItem) {
	o.Indexes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DamengSpaceAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["tablespaces"] = o.Tablespaces
	if o.Schemas != nil {
		toSerialize["schemas"] = o.Schemas
	}
	if o.Tables != nil {
		toSerialize["tables"] = o.Tables
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DamengSpaceAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Tablespaces *[]DamengTablespaceItem `json:"tablespaces"`
		Schemas     []DamengSchemaItem      `json:"schemas,omitempty"`
		Tables      []DamengTableSizeItem   `json:"tables,omitempty"`
		Indexes     []DamengIndexItem       `json:"indexes,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Tablespaces == nil {
		return fmt.Errorf("required field tablespaces missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"tablespaces", "schemas", "tables", "indexes"})
	} else {
		return err
	}
	o.Tablespaces = *all.Tablespaces
	o.Schemas = all.Schemas
	o.Tables = all.Tables
	o.Indexes = all.Indexes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
