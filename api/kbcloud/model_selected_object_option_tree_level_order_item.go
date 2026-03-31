// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type SelectedObjectOptionTreeLevelOrderItem struct {
	// The object type, e.g., "database", "schema", "table"
	Type *string `json:"type,omitempty"`
	// The column name of the object type which is defined in the sql, e.g., "database_name", "schema_name", "table_name"
	ColumnName *string `json:"columnName,omitempty"`
	// The object type which is defined in the sql, mutually exclusive with `type`
	TypeFromColumn *string `json:"typeFromColumn,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSelectedObjectOptionTreeLevelOrderItem instantiates a new SelectedObjectOptionTreeLevelOrderItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSelectedObjectOptionTreeLevelOrderItem() *SelectedObjectOptionTreeLevelOrderItem {
	this := SelectedObjectOptionTreeLevelOrderItem{}
	return &this
}

// NewSelectedObjectOptionTreeLevelOrderItemWithDefaults instantiates a new SelectedObjectOptionTreeLevelOrderItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSelectedObjectOptionTreeLevelOrderItemWithDefaults() *SelectedObjectOptionTreeLevelOrderItem {
	this := SelectedObjectOptionTreeLevelOrderItem{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SelectedObjectOptionTreeLevelOrderItem) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedObjectOptionTreeLevelOrderItem) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SelectedObjectOptionTreeLevelOrderItem) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SelectedObjectOptionTreeLevelOrderItem) SetType(v string) {
	o.Type = &v
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *SelectedObjectOptionTreeLevelOrderItem) GetColumnName() string {
	if o == nil || o.ColumnName == nil {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedObjectOptionTreeLevelOrderItem) GetColumnNameOk() (*string, bool) {
	if o == nil || o.ColumnName == nil {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *SelectedObjectOptionTreeLevelOrderItem) HasColumnName() bool {
	return o != nil && o.ColumnName != nil
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *SelectedObjectOptionTreeLevelOrderItem) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetTypeFromColumn returns the TypeFromColumn field value if set, zero value otherwise.
func (o *SelectedObjectOptionTreeLevelOrderItem) GetTypeFromColumn() string {
	if o == nil || o.TypeFromColumn == nil {
		var ret string
		return ret
	}
	return *o.TypeFromColumn
}

// GetTypeFromColumnOk returns a tuple with the TypeFromColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedObjectOptionTreeLevelOrderItem) GetTypeFromColumnOk() (*string, bool) {
	if o == nil || o.TypeFromColumn == nil {
		return nil, false
	}
	return o.TypeFromColumn, true
}

// HasTypeFromColumn returns a boolean if a field has been set.
func (o *SelectedObjectOptionTreeLevelOrderItem) HasTypeFromColumn() bool {
	return o != nil && o.TypeFromColumn != nil
}

// SetTypeFromColumn gets a reference to the given string and assigns it to the TypeFromColumn field.
func (o *SelectedObjectOptionTreeLevelOrderItem) SetTypeFromColumn(v string) {
	o.TypeFromColumn = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SelectedObjectOptionTreeLevelOrderItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ColumnName != nil {
		toSerialize["columnName"] = o.ColumnName
	}
	if o.TypeFromColumn != nil {
		toSerialize["typeFromColumn"] = o.TypeFromColumn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SelectedObjectOptionTreeLevelOrderItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type           *string `json:"type,omitempty"`
		ColumnName     *string `json:"columnName,omitempty"`
		TypeFromColumn *string `json:"typeFromColumn,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "columnName", "typeFromColumn"})
	} else {
		return err
	}
	o.Type = all.Type
	o.ColumnName = all.ColumnName
	o.TypeFromColumn = all.TypeFromColumn

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
