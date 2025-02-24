// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DmsUniqueKey struct {
	// The name of the unique key
	Name common.NullableString `json:"name,omitempty"`
	// The list of columns that make up the unique key
	Columns []string `json:"columns,omitempty"`
	// Additional columns included in the unique key
	Include common.NullableList[string] `json:"include,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsUniqueKey instantiates a new DmsUniqueKey object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsUniqueKey() *DmsUniqueKey {
	this := DmsUniqueKey{}
	return &this
}

// NewDmsUniqueKeyWithDefaults instantiates a new DmsUniqueKey object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsUniqueKeyWithDefaults() *DmsUniqueKey {
	this := DmsUniqueKey{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsUniqueKey) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsUniqueKey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DmsUniqueKey) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given common.NullableString and assigns it to the Name field.
func (o *DmsUniqueKey) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *DmsUniqueKey) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *DmsUniqueKey) UnsetName() {
	o.Name.Unset()
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *DmsUniqueKey) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsUniqueKey) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *DmsUniqueKey) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *DmsUniqueKey) SetColumns(v []string) {
	o.Columns = v
}

// GetInclude returns the Include field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsUniqueKey) GetInclude() []string {
	if o == nil || o.Include.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Include.Get()
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsUniqueKey) GetIncludeOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Include.Get(), o.Include.IsSet()
}

// HasInclude returns a boolean if a field has been set.
func (o *DmsUniqueKey) HasInclude() bool {
	return o != nil && o.Include.IsSet()
}

// SetInclude gets a reference to the given common.NullableList[string] and assigns it to the Include field.
func (o *DmsUniqueKey) SetInclude(v []string) {
	o.Include.Set(&v)
}

// SetIncludeNil sets the value for Include to be an explicit nil.
func (o *DmsUniqueKey) SetIncludeNil() {
	o.Include.Set(nil)
}

// UnsetInclude ensures that no value is present for Include, not even an explicit nil.
func (o *DmsUniqueKey) UnsetInclude() {
	o.Include.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsUniqueKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Include.IsSet() {
		toSerialize["include"] = o.Include.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsUniqueKey) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name    common.NullableString       `json:"name,omitempty"`
		Columns []string                    `json:"columns,omitempty"`
		Include common.NullableList[string] `json:"include,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "columns", "include"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Columns = all.Columns
	o.Include = all.Include

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
