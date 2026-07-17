// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ESSecurityRemoteIndexPrivileges struct {
	Clusters               []string                 `json:"clusters"`
	Names                  []string                 `json:"names"`
	Privileges             []string                 `json:"privileges"`
	FieldSecurity          *ESSecurityFieldSecurity `json:"field_security,omitempty"`
	Query                  common.NullableString    `json:"query,omitempty"`
	AllowRestrictedIndices common.NullableBool      `json:"allow_restricted_indices,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewESSecurityRemoteIndexPrivileges instantiates a new ESSecurityRemoteIndexPrivileges object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewESSecurityRemoteIndexPrivileges(clusters []string, names []string, privileges []string) *ESSecurityRemoteIndexPrivileges {
	this := ESSecurityRemoteIndexPrivileges{}
	this.Clusters = clusters
	this.Names = names
	this.Privileges = privileges
	return &this
}

// NewESSecurityRemoteIndexPrivilegesWithDefaults instantiates a new ESSecurityRemoteIndexPrivileges object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewESSecurityRemoteIndexPrivilegesWithDefaults() *ESSecurityRemoteIndexPrivileges {
	this := ESSecurityRemoteIndexPrivileges{}
	return &this
}

// GetClusters returns the Clusters field value.
func (o *ESSecurityRemoteIndexPrivileges) GetClusters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value
// and a boolean to check if the value has been set.
func (o *ESSecurityRemoteIndexPrivileges) GetClustersOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clusters, true
}

// SetClusters sets field value.
func (o *ESSecurityRemoteIndexPrivileges) SetClusters(v []string) {
	o.Clusters = v
}

// GetNames returns the Names field value.
func (o *ESSecurityRemoteIndexPrivileges) GetNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value
// and a boolean to check if the value has been set.
func (o *ESSecurityRemoteIndexPrivileges) GetNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Names, true
}

// SetNames sets field value.
func (o *ESSecurityRemoteIndexPrivileges) SetNames(v []string) {
	o.Names = v
}

// GetPrivileges returns the Privileges field value.
func (o *ESSecurityRemoteIndexPrivileges) GetPrivileges() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value
// and a boolean to check if the value has been set.
func (o *ESSecurityRemoteIndexPrivileges) GetPrivilegesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// SetPrivileges sets field value.
func (o *ESSecurityRemoteIndexPrivileges) SetPrivileges(v []string) {
	o.Privileges = v
}

// GetFieldSecurity returns the FieldSecurity field value if set, zero value otherwise.
func (o *ESSecurityRemoteIndexPrivileges) GetFieldSecurity() ESSecurityFieldSecurity {
	if o == nil || o.FieldSecurity == nil {
		var ret ESSecurityFieldSecurity
		return ret
	}
	return *o.FieldSecurity
}

// GetFieldSecurityOk returns a tuple with the FieldSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRemoteIndexPrivileges) GetFieldSecurityOk() (*ESSecurityFieldSecurity, bool) {
	if o == nil || o.FieldSecurity == nil {
		return nil, false
	}
	return o.FieldSecurity, true
}

// HasFieldSecurity returns a boolean if a field has been set.
func (o *ESSecurityRemoteIndexPrivileges) HasFieldSecurity() bool {
	return o != nil && o.FieldSecurity != nil
}

// SetFieldSecurity gets a reference to the given ESSecurityFieldSecurity and assigns it to the FieldSecurity field.
func (o *ESSecurityRemoteIndexPrivileges) SetFieldSecurity(v ESSecurityFieldSecurity) {
	o.FieldSecurity = &v
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ESSecurityRemoteIndexPrivileges) GetQuery() string {
	if o == nil || o.Query.Get() == nil {
		var ret string
		return ret
	}
	return *o.Query.Get()
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ESSecurityRemoteIndexPrivileges) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Query.Get(), o.Query.IsSet()
}

// HasQuery returns a boolean if a field has been set.
func (o *ESSecurityRemoteIndexPrivileges) HasQuery() bool {
	return o != nil && o.Query.IsSet()
}

// SetQuery gets a reference to the given common.NullableString and assigns it to the Query field.
func (o *ESSecurityRemoteIndexPrivileges) SetQuery(v string) {
	o.Query.Set(&v)
}

// SetQueryNil sets the value for Query to be an explicit nil.
func (o *ESSecurityRemoteIndexPrivileges) SetQueryNil() {
	o.Query.Set(nil)
}

// UnsetQuery ensures that no value is present for Query, not even an explicit nil.
func (o *ESSecurityRemoteIndexPrivileges) UnsetQuery() {
	o.Query.Unset()
}

// GetAllowRestrictedIndices returns the AllowRestrictedIndices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ESSecurityRemoteIndexPrivileges) GetAllowRestrictedIndices() bool {
	if o == nil || o.AllowRestrictedIndices.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AllowRestrictedIndices.Get()
}

// GetAllowRestrictedIndicesOk returns a tuple with the AllowRestrictedIndices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ESSecurityRemoteIndexPrivileges) GetAllowRestrictedIndicesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowRestrictedIndices.Get(), o.AllowRestrictedIndices.IsSet()
}

// HasAllowRestrictedIndices returns a boolean if a field has been set.
func (o *ESSecurityRemoteIndexPrivileges) HasAllowRestrictedIndices() bool {
	return o != nil && o.AllowRestrictedIndices.IsSet()
}

// SetAllowRestrictedIndices gets a reference to the given common.NullableBool and assigns it to the AllowRestrictedIndices field.
func (o *ESSecurityRemoteIndexPrivileges) SetAllowRestrictedIndices(v bool) {
	o.AllowRestrictedIndices.Set(&v)
}

// SetAllowRestrictedIndicesNil sets the value for AllowRestrictedIndices to be an explicit nil.
func (o *ESSecurityRemoteIndexPrivileges) SetAllowRestrictedIndicesNil() {
	o.AllowRestrictedIndices.Set(nil)
}

// UnsetAllowRestrictedIndices ensures that no value is present for AllowRestrictedIndices, not even an explicit nil.
func (o *ESSecurityRemoteIndexPrivileges) UnsetAllowRestrictedIndices() {
	o.AllowRestrictedIndices.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ESSecurityRemoteIndexPrivileges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["clusters"] = o.Clusters
	toSerialize["names"] = o.Names
	toSerialize["privileges"] = o.Privileges
	if o.FieldSecurity != nil {
		toSerialize["field_security"] = o.FieldSecurity
	}
	if o.Query.IsSet() {
		toSerialize["query"] = o.Query.Get()
	}
	if o.AllowRestrictedIndices.IsSet() {
		toSerialize["allow_restricted_indices"] = o.AllowRestrictedIndices.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ESSecurityRemoteIndexPrivileges) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Clusters               *[]string                `json:"clusters"`
		Names                  *[]string                `json:"names"`
		Privileges             *[]string                `json:"privileges"`
		FieldSecurity          *ESSecurityFieldSecurity `json:"field_security,omitempty"`
		Query                  common.NullableString    `json:"query,omitempty"`
		AllowRestrictedIndices common.NullableBool      `json:"allow_restricted_indices,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Clusters == nil {
		return fmt.Errorf("required field clusters missing")
	}
	if all.Names == nil {
		return fmt.Errorf("required field names missing")
	}
	if all.Privileges == nil {
		return fmt.Errorf("required field privileges missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clusters", "names", "privileges", "field_security", "query", "allow_restricted_indices"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Clusters = *all.Clusters
	o.Names = *all.Names
	o.Privileges = *all.Privileges
	if all.FieldSecurity != nil && all.FieldSecurity.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.FieldSecurity = all.FieldSecurity
	o.Query = all.Query
	o.AllowRestrictedIndices = all.AllowRestrictedIndices

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
