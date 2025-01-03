// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type CdcSqlExecutor struct {
	Sql          []string              `json:"sql,omitempty"`
	Result       []string              `json:"result,omitempty"`
	AuthDatabase common.NullableString `json:"authDatabase,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCdcSqlExecutor instantiates a new CdcSqlExecutor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCdcSqlExecutor() *CdcSqlExecutor {
	this := CdcSqlExecutor{}
	return &this
}

// NewCdcSqlExecutorWithDefaults instantiates a new CdcSqlExecutor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCdcSqlExecutorWithDefaults() *CdcSqlExecutor {
	this := CdcSqlExecutor{}
	return &this
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *CdcSqlExecutor) GetSql() []string {
	if o == nil || o.Sql == nil {
		var ret []string
		return ret
	}
	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcSqlExecutor) GetSqlOk() (*[]string, bool) {
	if o == nil || o.Sql == nil {
		return nil, false
	}
	return &o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *CdcSqlExecutor) HasSql() bool {
	return o != nil && o.Sql != nil
}

// SetSql gets a reference to the given []string and assigns it to the Sql field.
func (o *CdcSqlExecutor) SetSql(v []string) {
	o.Sql = v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CdcSqlExecutor) GetResult() []string {
	if o == nil || o.Result == nil {
		var ret []string
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcSqlExecutor) GetResultOk() (*[]string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return &o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CdcSqlExecutor) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given []string and assigns it to the Result field.
func (o *CdcSqlExecutor) SetResult(v []string) {
	o.Result = v
}

// GetAuthDatabase returns the AuthDatabase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CdcSqlExecutor) GetAuthDatabase() string {
	if o == nil || o.AuthDatabase.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthDatabase.Get()
}

// GetAuthDatabaseOk returns a tuple with the AuthDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CdcSqlExecutor) GetAuthDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthDatabase.Get(), o.AuthDatabase.IsSet()
}

// HasAuthDatabase returns a boolean if a field has been set.
func (o *CdcSqlExecutor) HasAuthDatabase() bool {
	return o != nil && o.AuthDatabase.IsSet()
}

// SetAuthDatabase gets a reference to the given common.NullableString and assigns it to the AuthDatabase field.
func (o *CdcSqlExecutor) SetAuthDatabase(v string) {
	o.AuthDatabase.Set(&v)
}

// SetAuthDatabaseNil sets the value for AuthDatabase to be an explicit nil.
func (o *CdcSqlExecutor) SetAuthDatabaseNil() {
	o.AuthDatabase.Set(nil)
}

// UnsetAuthDatabase ensures that no value is present for AuthDatabase, not even an explicit nil.
func (o *CdcSqlExecutor) UnsetAuthDatabase() {
	o.AuthDatabase.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o CdcSqlExecutor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Sql != nil {
		toSerialize["sql"] = o.Sql
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.AuthDatabase.IsSet() {
		toSerialize["authDatabase"] = o.AuthDatabase.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CdcSqlExecutor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Sql          []string              `json:"sql,omitempty"`
		Result       []string              `json:"result,omitempty"`
		AuthDatabase common.NullableString `json:"authDatabase,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sql", "result", "authDatabase"})
	} else {
		return err
	}
	o.Sql = all.Sql
	o.Result = all.Result
	o.AuthDatabase = all.AuthDatabase

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
