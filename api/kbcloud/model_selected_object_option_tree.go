// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type SelectedObjectOptionTree struct {
	// if set to true, the tree will switch when database is switched. This is used for database-level backup, and the root level of the tree is database. If not set or set to false, the tree will not switch when database is switched, and the root level of the tree is instance.
	SwitchWithDatabase *bool `json:"switchWithDatabase,omitempty"`
	// The SQL query to retrieve the object tree
	Sql string `json:"sql"`
	// The order of object types in the hierarchy
	LevelOrder []SelectedObjectOptionTreeLevelOrderItem `json:"levelOrder"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSelectedObjectOptionTree instantiates a new SelectedObjectOptionTree object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSelectedObjectOptionTree(sql string, levelOrder []SelectedObjectOptionTreeLevelOrderItem) *SelectedObjectOptionTree {
	this := SelectedObjectOptionTree{}
	this.Sql = sql
	this.LevelOrder = levelOrder
	return &this
}

// NewSelectedObjectOptionTreeWithDefaults instantiates a new SelectedObjectOptionTree object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSelectedObjectOptionTreeWithDefaults() *SelectedObjectOptionTree {
	this := SelectedObjectOptionTree{}
	return &this
}

// GetSwitchWithDatabase returns the SwitchWithDatabase field value if set, zero value otherwise.
func (o *SelectedObjectOptionTree) GetSwitchWithDatabase() bool {
	if o == nil || o.SwitchWithDatabase == nil {
		var ret bool
		return ret
	}
	return *o.SwitchWithDatabase
}

// GetSwitchWithDatabaseOk returns a tuple with the SwitchWithDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedObjectOptionTree) GetSwitchWithDatabaseOk() (*bool, bool) {
	if o == nil || o.SwitchWithDatabase == nil {
		return nil, false
	}
	return o.SwitchWithDatabase, true
}

// HasSwitchWithDatabase returns a boolean if a field has been set.
func (o *SelectedObjectOptionTree) HasSwitchWithDatabase() bool {
	return o != nil && o.SwitchWithDatabase != nil
}

// SetSwitchWithDatabase gets a reference to the given bool and assigns it to the SwitchWithDatabase field.
func (o *SelectedObjectOptionTree) SetSwitchWithDatabase(v bool) {
	o.SwitchWithDatabase = &v
}

// GetSql returns the Sql field value.
func (o *SelectedObjectOptionTree) GetSql() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value
// and a boolean to check if the value has been set.
func (o *SelectedObjectOptionTree) GetSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sql, true
}

// SetSql sets field value.
func (o *SelectedObjectOptionTree) SetSql(v string) {
	o.Sql = v
}

// GetLevelOrder returns the LevelOrder field value.
func (o *SelectedObjectOptionTree) GetLevelOrder() []SelectedObjectOptionTreeLevelOrderItem {
	if o == nil {
		var ret []SelectedObjectOptionTreeLevelOrderItem
		return ret
	}
	return o.LevelOrder
}

// GetLevelOrderOk returns a tuple with the LevelOrder field value
// and a boolean to check if the value has been set.
func (o *SelectedObjectOptionTree) GetLevelOrderOk() (*[]SelectedObjectOptionTreeLevelOrderItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LevelOrder, true
}

// SetLevelOrder sets field value.
func (o *SelectedObjectOptionTree) SetLevelOrder(v []SelectedObjectOptionTreeLevelOrderItem) {
	o.LevelOrder = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SelectedObjectOptionTree) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.SwitchWithDatabase != nil {
		toSerialize["switchWithDatabase"] = o.SwitchWithDatabase
	}
	toSerialize["sql"] = o.Sql
	toSerialize["levelOrder"] = o.LevelOrder

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SelectedObjectOptionTree) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SwitchWithDatabase *bool                                     `json:"switchWithDatabase,omitempty"`
		Sql                *string                                   `json:"sql"`
		LevelOrder         *[]SelectedObjectOptionTreeLevelOrderItem `json:"levelOrder"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Sql == nil {
		return fmt.Errorf("required field sql missing")
	}
	if all.LevelOrder == nil {
		return fmt.Errorf("required field levelOrder missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"switchWithDatabase", "sql", "levelOrder"})
	} else {
		return err
	}
	o.SwitchWithDatabase = all.SwitchWithDatabase
	o.Sql = *all.Sql
	o.LevelOrder = *all.LevelOrder

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
