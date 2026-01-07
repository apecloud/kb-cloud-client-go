// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type SqlAuditSwitchOption struct {
	Type SqlAuditSwitchOptionType `json:"type"`
	// SQL query to check audit log status (required when type is sql)
	Sql *string `json:"sql,omitempty"`
	// Parameter configuration for querying and controlling (required when type is parameter). When type is parameter, both query and control (reconfigure) are inferred from this configuration.
	Parameter *SqlAuditParameterOption `json:"parameter,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSqlAuditSwitchOption instantiates a new SqlAuditSwitchOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSqlAuditSwitchOption(typeVar SqlAuditSwitchOptionType) *SqlAuditSwitchOption {
	this := SqlAuditSwitchOption{}
	this.Type = typeVar
	return &this
}

// NewSqlAuditSwitchOptionWithDefaults instantiates a new SqlAuditSwitchOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSqlAuditSwitchOptionWithDefaults() *SqlAuditSwitchOption {
	this := SqlAuditSwitchOption{}
	return &this
}

// GetType returns the Type field value.
func (o *SqlAuditSwitchOption) GetType() SqlAuditSwitchOptionType {
	if o == nil {
		var ret SqlAuditSwitchOptionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SqlAuditSwitchOption) GetTypeOk() (*SqlAuditSwitchOptionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SqlAuditSwitchOption) SetType(v SqlAuditSwitchOptionType) {
	o.Type = v
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *SqlAuditSwitchOption) GetSql() string {
	if o == nil || o.Sql == nil {
		var ret string
		return ret
	}
	return *o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlAuditSwitchOption) GetSqlOk() (*string, bool) {
	if o == nil || o.Sql == nil {
		return nil, false
	}
	return o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *SqlAuditSwitchOption) HasSql() bool {
	return o != nil && o.Sql != nil
}

// SetSql gets a reference to the given string and assigns it to the Sql field.
func (o *SqlAuditSwitchOption) SetSql(v string) {
	o.Sql = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *SqlAuditSwitchOption) GetParameter() SqlAuditParameterOption {
	if o == nil || o.Parameter == nil {
		var ret SqlAuditParameterOption
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlAuditSwitchOption) GetParameterOk() (*SqlAuditParameterOption, bool) {
	if o == nil || o.Parameter == nil {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *SqlAuditSwitchOption) HasParameter() bool {
	return o != nil && o.Parameter != nil
}

// SetParameter gets a reference to the given SqlAuditParameterOption and assigns it to the Parameter field.
func (o *SqlAuditSwitchOption) SetParameter(v SqlAuditParameterOption) {
	o.Parameter = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SqlAuditSwitchOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	if o.Sql != nil {
		toSerialize["sql"] = o.Sql
	}
	if o.Parameter != nil {
		toSerialize["parameter"] = o.Parameter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SqlAuditSwitchOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type      *SqlAuditSwitchOptionType `json:"type"`
		Sql       *string                   `json:"sql,omitempty"`
		Parameter *SqlAuditParameterOption  `json:"parameter,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "sql", "parameter"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Sql = all.Sql
	if all.Parameter != nil && all.Parameter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Parameter = all.Parameter

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
