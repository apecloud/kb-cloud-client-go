// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DmsExcludeConstraint struct {
	// The name of the exclude constraint
	Name common.NullableString `json:"name,omitempty"`
	// The index method (e.g., GIST, BTREE)
	IndexMethod *string                           `json:"indexMethod,omitempty"`
	Exclude     []DmsExcludeConstraintExcludeItem `json:"exclude,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsExcludeConstraint instantiates a new DmsExcludeConstraint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsExcludeConstraint() *DmsExcludeConstraint {
	this := DmsExcludeConstraint{}
	return &this
}

// NewDmsExcludeConstraintWithDefaults instantiates a new DmsExcludeConstraint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsExcludeConstraintWithDefaults() *DmsExcludeConstraint {
	this := DmsExcludeConstraint{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExcludeConstraint) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExcludeConstraint) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DmsExcludeConstraint) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given common.NullableString and assigns it to the Name field.
func (o *DmsExcludeConstraint) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *DmsExcludeConstraint) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *DmsExcludeConstraint) UnsetName() {
	o.Name.Unset()
}

// GetIndexMethod returns the IndexMethod field value if set, zero value otherwise.
func (o *DmsExcludeConstraint) GetIndexMethod() string {
	if o == nil || o.IndexMethod == nil {
		var ret string
		return ret
	}
	return *o.IndexMethod
}

// GetIndexMethodOk returns a tuple with the IndexMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExcludeConstraint) GetIndexMethodOk() (*string, bool) {
	if o == nil || o.IndexMethod == nil {
		return nil, false
	}
	return o.IndexMethod, true
}

// HasIndexMethod returns a boolean if a field has been set.
func (o *DmsExcludeConstraint) HasIndexMethod() bool {
	return o != nil && o.IndexMethod != nil
}

// SetIndexMethod gets a reference to the given string and assigns it to the IndexMethod field.
func (o *DmsExcludeConstraint) SetIndexMethod(v string) {
	o.IndexMethod = &v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *DmsExcludeConstraint) GetExclude() []DmsExcludeConstraintExcludeItem {
	if o == nil || o.Exclude == nil {
		var ret []DmsExcludeConstraintExcludeItem
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExcludeConstraint) GetExcludeOk() (*[]DmsExcludeConstraintExcludeItem, bool) {
	if o == nil || o.Exclude == nil {
		return nil, false
	}
	return &o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *DmsExcludeConstraint) HasExclude() bool {
	return o != nil && o.Exclude != nil
}

// SetExclude gets a reference to the given []DmsExcludeConstraintExcludeItem and assigns it to the Exclude field.
func (o *DmsExcludeConstraint) SetExclude(v []DmsExcludeConstraintExcludeItem) {
	o.Exclude = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsExcludeConstraint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.IndexMethod != nil {
		toSerialize["indexMethod"] = o.IndexMethod
	}
	if o.Exclude != nil {
		toSerialize["exclude"] = o.Exclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsExcludeConstraint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        common.NullableString             `json:"name,omitempty"`
		IndexMethod *string                           `json:"indexMethod,omitempty"`
		Exclude     []DmsExcludeConstraintExcludeItem `json:"exclude,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "indexMethod", "exclude"})
	} else {
		return err
	}
	o.Name = all.Name
	o.IndexMethod = all.IndexMethod
	o.Exclude = all.Exclude

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
