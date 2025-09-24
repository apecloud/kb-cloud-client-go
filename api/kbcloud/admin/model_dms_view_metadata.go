// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsViewMetadata struct {
	// The name of the view
	ViewName *string `json:"viewName,omitempty"`
	// The database in which the view will be created
	Database *string `json:"database,omitempty"`
	// Whether to replace the existing view if it exists
	Replace common.NullableBool `json:"replace,omitempty"`
	// The definer of the view
	Definer common.NullableString `json:"definer,omitempty"`
	// The SQL security context of the view (e.g., DEFINER or INVOKER)
	SqlSecurity common.NullableString `json:"sqlSecurity,omitempty"`
	// The WITH CHECK OPTION clause (e.g., CASCADED or LOCAL)
	CheckOption common.NullableString `json:"checkOption,omitempty"`
	// The algorithm used for the view (e.g., UNDEFINED, MERGE, TEMPTABLE)
	Algorithm common.NullableString `json:"algorithm,omitempty"`
	// The SQL statement defining the view
	Definition common.NullableString `json:"definition,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsViewMetadata instantiates a new DmsViewMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsViewMetadata() *DmsViewMetadata {
	this := DmsViewMetadata{}
	return &this
}

// NewDmsViewMetadataWithDefaults instantiates a new DmsViewMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsViewMetadataWithDefaults() *DmsViewMetadata {
	this := DmsViewMetadata{}
	return &this
}

// GetViewName returns the ViewName field value if set, zero value otherwise.
func (o *DmsViewMetadata) GetViewName() string {
	if o == nil || o.ViewName == nil {
		var ret string
		return ret
	}
	return *o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsViewMetadata) GetViewNameOk() (*string, bool) {
	if o == nil || o.ViewName == nil {
		return nil, false
	}
	return o.ViewName, true
}

// HasViewName returns a boolean if a field has been set.
func (o *DmsViewMetadata) HasViewName() bool {
	return o != nil && o.ViewName != nil
}

// SetViewName gets a reference to the given string and assigns it to the ViewName field.
func (o *DmsViewMetadata) SetViewName(v string) {
	o.ViewName = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *DmsViewMetadata) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsViewMetadata) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *DmsViewMetadata) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *DmsViewMetadata) SetDatabase(v string) {
	o.Database = &v
}

// GetReplace returns the Replace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsViewMetadata) GetReplace() bool {
	if o == nil || o.Replace.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Replace.Get()
}

// GetReplaceOk returns a tuple with the Replace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsViewMetadata) GetReplaceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Replace.Get(), o.Replace.IsSet()
}

// HasReplace returns a boolean if a field has been set.
func (o *DmsViewMetadata) HasReplace() bool {
	return o != nil && o.Replace.IsSet()
}

// SetReplace gets a reference to the given common.NullableBool and assigns it to the Replace field.
func (o *DmsViewMetadata) SetReplace(v bool) {
	o.Replace.Set(&v)
}

// SetReplaceNil sets the value for Replace to be an explicit nil.
func (o *DmsViewMetadata) SetReplaceNil() {
	o.Replace.Set(nil)
}

// UnsetReplace ensures that no value is present for Replace, not even an explicit nil.
func (o *DmsViewMetadata) UnsetReplace() {
	o.Replace.Unset()
}

// GetDefiner returns the Definer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsViewMetadata) GetDefiner() string {
	if o == nil || o.Definer.Get() == nil {
		var ret string
		return ret
	}
	return *o.Definer.Get()
}

// GetDefinerOk returns a tuple with the Definer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsViewMetadata) GetDefinerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Definer.Get(), o.Definer.IsSet()
}

// HasDefiner returns a boolean if a field has been set.
func (o *DmsViewMetadata) HasDefiner() bool {
	return o != nil && o.Definer.IsSet()
}

// SetDefiner gets a reference to the given common.NullableString and assigns it to the Definer field.
func (o *DmsViewMetadata) SetDefiner(v string) {
	o.Definer.Set(&v)
}

// SetDefinerNil sets the value for Definer to be an explicit nil.
func (o *DmsViewMetadata) SetDefinerNil() {
	o.Definer.Set(nil)
}

// UnsetDefiner ensures that no value is present for Definer, not even an explicit nil.
func (o *DmsViewMetadata) UnsetDefiner() {
	o.Definer.Unset()
}

// GetSqlSecurity returns the SqlSecurity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsViewMetadata) GetSqlSecurity() string {
	if o == nil || o.SqlSecurity.Get() == nil {
		var ret string
		return ret
	}
	return *o.SqlSecurity.Get()
}

// GetSqlSecurityOk returns a tuple with the SqlSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsViewMetadata) GetSqlSecurityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SqlSecurity.Get(), o.SqlSecurity.IsSet()
}

// HasSqlSecurity returns a boolean if a field has been set.
func (o *DmsViewMetadata) HasSqlSecurity() bool {
	return o != nil && o.SqlSecurity.IsSet()
}

// SetSqlSecurity gets a reference to the given common.NullableString and assigns it to the SqlSecurity field.
func (o *DmsViewMetadata) SetSqlSecurity(v string) {
	o.SqlSecurity.Set(&v)
}

// SetSqlSecurityNil sets the value for SqlSecurity to be an explicit nil.
func (o *DmsViewMetadata) SetSqlSecurityNil() {
	o.SqlSecurity.Set(nil)
}

// UnsetSqlSecurity ensures that no value is present for SqlSecurity, not even an explicit nil.
func (o *DmsViewMetadata) UnsetSqlSecurity() {
	o.SqlSecurity.Unset()
}

// GetCheckOption returns the CheckOption field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsViewMetadata) GetCheckOption() string {
	if o == nil || o.CheckOption.Get() == nil {
		var ret string
		return ret
	}
	return *o.CheckOption.Get()
}

// GetCheckOptionOk returns a tuple with the CheckOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsViewMetadata) GetCheckOptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckOption.Get(), o.CheckOption.IsSet()
}

// HasCheckOption returns a boolean if a field has been set.
func (o *DmsViewMetadata) HasCheckOption() bool {
	return o != nil && o.CheckOption.IsSet()
}

// SetCheckOption gets a reference to the given common.NullableString and assigns it to the CheckOption field.
func (o *DmsViewMetadata) SetCheckOption(v string) {
	o.CheckOption.Set(&v)
}

// SetCheckOptionNil sets the value for CheckOption to be an explicit nil.
func (o *DmsViewMetadata) SetCheckOptionNil() {
	o.CheckOption.Set(nil)
}

// UnsetCheckOption ensures that no value is present for CheckOption, not even an explicit nil.
func (o *DmsViewMetadata) UnsetCheckOption() {
	o.CheckOption.Unset()
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsViewMetadata) GetAlgorithm() string {
	if o == nil || o.Algorithm.Get() == nil {
		var ret string
		return ret
	}
	return *o.Algorithm.Get()
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsViewMetadata) GetAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Algorithm.Get(), o.Algorithm.IsSet()
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *DmsViewMetadata) HasAlgorithm() bool {
	return o != nil && o.Algorithm.IsSet()
}

// SetAlgorithm gets a reference to the given common.NullableString and assigns it to the Algorithm field.
func (o *DmsViewMetadata) SetAlgorithm(v string) {
	o.Algorithm.Set(&v)
}

// SetAlgorithmNil sets the value for Algorithm to be an explicit nil.
func (o *DmsViewMetadata) SetAlgorithmNil() {
	o.Algorithm.Set(nil)
}

// UnsetAlgorithm ensures that no value is present for Algorithm, not even an explicit nil.
func (o *DmsViewMetadata) UnsetAlgorithm() {
	o.Algorithm.Unset()
}

// GetDefinition returns the Definition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsViewMetadata) GetDefinition() string {
	if o == nil || o.Definition.Get() == nil {
		var ret string
		return ret
	}
	return *o.Definition.Get()
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsViewMetadata) GetDefinitionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Definition.Get(), o.Definition.IsSet()
}

// HasDefinition returns a boolean if a field has been set.
func (o *DmsViewMetadata) HasDefinition() bool {
	return o != nil && o.Definition.IsSet()
}

// SetDefinition gets a reference to the given common.NullableString and assigns it to the Definition field.
func (o *DmsViewMetadata) SetDefinition(v string) {
	o.Definition.Set(&v)
}

// SetDefinitionNil sets the value for Definition to be an explicit nil.
func (o *DmsViewMetadata) SetDefinitionNil() {
	o.Definition.Set(nil)
}

// UnsetDefinition ensures that no value is present for Definition, not even an explicit nil.
func (o *DmsViewMetadata) UnsetDefinition() {
	o.Definition.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsViewMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ViewName != nil {
		toSerialize["viewName"] = o.ViewName
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.Replace.IsSet() {
		toSerialize["replace"] = o.Replace.Get()
	}
	if o.Definer.IsSet() {
		toSerialize["definer"] = o.Definer.Get()
	}
	if o.SqlSecurity.IsSet() {
		toSerialize["sqlSecurity"] = o.SqlSecurity.Get()
	}
	if o.CheckOption.IsSet() {
		toSerialize["checkOption"] = o.CheckOption.Get()
	}
	if o.Algorithm.IsSet() {
		toSerialize["algorithm"] = o.Algorithm.Get()
	}
	if o.Definition.IsSet() {
		toSerialize["definition"] = o.Definition.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsViewMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ViewName    *string               `json:"viewName,omitempty"`
		Database    *string               `json:"database,omitempty"`
		Replace     common.NullableBool   `json:"replace,omitempty"`
		Definer     common.NullableString `json:"definer,omitempty"`
		SqlSecurity common.NullableString `json:"sqlSecurity,omitempty"`
		CheckOption common.NullableString `json:"checkOption,omitempty"`
		Algorithm   common.NullableString `json:"algorithm,omitempty"`
		Definition  common.NullableString `json:"definition,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"viewName", "database", "replace", "definer", "sqlSecurity", "checkOption", "algorithm", "definition"})
	} else {
		return err
	}
	o.ViewName = all.ViewName
	o.Database = all.Database
	o.Replace = all.Replace
	o.Definer = all.Definer
	o.SqlSecurity = all.SqlSecurity
	o.CheckOption = all.CheckOption
	o.Algorithm = all.Algorithm
	o.Definition = all.Definition

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
