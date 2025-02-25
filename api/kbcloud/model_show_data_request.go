// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ShowDataRequest struct {
	// the database of the table or view
	Database *string `json:"database,omitempty"`
	// the target table or view name
	Table *string `json:"table,omitempty"`
	// return limited number of data
	Limit common.NullableInt32 `json:"limit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewShowDataRequest instantiates a new ShowDataRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewShowDataRequest() *ShowDataRequest {
	this := ShowDataRequest{}
	var database string = ""
	this.Database = &database
	return &this
}

// NewShowDataRequestWithDefaults instantiates a new ShowDataRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewShowDataRequestWithDefaults() *ShowDataRequest {
	this := ShowDataRequest{}
	var database string = ""
	this.Database = &database
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ShowDataRequest) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShowDataRequest) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ShowDataRequest) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *ShowDataRequest) SetDatabase(v string) {
	o.Database = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *ShowDataRequest) GetTable() string {
	if o == nil || o.Table == nil {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShowDataRequest) GetTableOk() (*string, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *ShowDataRequest) HasTable() bool {
	return o != nil && o.Table != nil
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *ShowDataRequest) SetTable(v string) {
	o.Table = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShowDataRequest) GetLimit() int32 {
	if o == nil || o.Limit.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ShowDataRequest) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *ShowDataRequest) HasLimit() bool {
	return o != nil && o.Limit.IsSet()
}

// SetLimit gets a reference to the given common.NullableInt32 and assigns it to the Limit field.
func (o *ShowDataRequest) SetLimit(v int32) {
	o.Limit.Set(&v)
}

// SetLimitNil sets the value for Limit to be an explicit nil.
func (o *ShowDataRequest) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil.
func (o *ShowDataRequest) UnsetLimit() {
	o.Limit.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ShowDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.Table != nil {
		toSerialize["table"] = o.Table
	}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ShowDataRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database *string              `json:"database,omitempty"`
		Table    *string              `json:"table,omitempty"`
		Limit    common.NullableInt32 `json:"limit,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "table", "limit"})
	} else {
		return err
	}
	o.Database = all.Database
	o.Table = all.Table
	o.Limit = all.Limit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
