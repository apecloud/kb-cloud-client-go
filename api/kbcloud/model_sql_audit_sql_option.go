// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type SqlAuditSQLOption struct {
	Query SqlAuditSQLQueryOption `json:"query"`
	// SQL statements to enable audit log. Multiple statements will be joined and executed.
	Enable []string `json:"enable"`
	// SQL statements to disable audit log. Multiple statements will be joined and executed.
	Disable []string `json:"disable"`
	// Specific user account name to execute SQL statements. If not provided, uses the default datasource account.
	User *string `json:"user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSqlAuditSQLOption instantiates a new SqlAuditSQLOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSqlAuditSQLOption(query SqlAuditSQLQueryOption, enable []string, disable []string) *SqlAuditSQLOption {
	this := SqlAuditSQLOption{}
	this.Query = query
	this.Enable = enable
	this.Disable = disable
	return &this
}

// NewSqlAuditSQLOptionWithDefaults instantiates a new SqlAuditSQLOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSqlAuditSQLOptionWithDefaults() *SqlAuditSQLOption {
	this := SqlAuditSQLOption{}
	return &this
}

// GetQuery returns the Query field value.
func (o *SqlAuditSQLOption) GetQuery() SqlAuditSQLQueryOption {
	if o == nil {
		var ret SqlAuditSQLQueryOption
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SqlAuditSQLOption) GetQueryOk() (*SqlAuditSQLQueryOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *SqlAuditSQLOption) SetQuery(v SqlAuditSQLQueryOption) {
	o.Query = v
}

// GetEnable returns the Enable field value.
func (o *SqlAuditSQLOption) GetEnable() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value
// and a boolean to check if the value has been set.
func (o *SqlAuditSQLOption) GetEnableOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enable, true
}

// SetEnable sets field value.
func (o *SqlAuditSQLOption) SetEnable(v []string) {
	o.Enable = v
}

// GetDisable returns the Disable field value.
func (o *SqlAuditSQLOption) GetDisable() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Disable
}

// GetDisableOk returns a tuple with the Disable field value
// and a boolean to check if the value has been set.
func (o *SqlAuditSQLOption) GetDisableOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disable, true
}

// SetDisable sets field value.
func (o *SqlAuditSQLOption) SetDisable(v []string) {
	o.Disable = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *SqlAuditSQLOption) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlAuditSQLOption) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *SqlAuditSQLOption) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *SqlAuditSQLOption) SetUser(v string) {
	o.User = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SqlAuditSQLOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["query"] = o.Query
	toSerialize["enable"] = o.Enable
	toSerialize["disable"] = o.Disable
	if o.User != nil {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SqlAuditSQLOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query   *SqlAuditSQLQueryOption `json:"query"`
		Enable  *[]string               `json:"enable"`
		Disable *[]string               `json:"disable"`
		User    *string                 `json:"user,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.Enable == nil {
		return fmt.Errorf("required field enable missing")
	}
	if all.Disable == nil {
		return fmt.Errorf("required field disable missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"query", "enable", "disable", "user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = *all.Query
	o.Enable = *all.Enable
	o.Disable = *all.Disable
	o.User = all.User

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
