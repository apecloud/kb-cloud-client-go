// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DmsTableIndex struct {
	// The name of the index
	Name *string `json:"name,omitempty"`
	// The type of the index (e.g., UNIQUE, FULLTEXT)
	Type *string `json:"type,omitempty"`
	// List of columns included in the index
	Columns []string `json:"columns,omitempty"`
	// A comment describing the index
	Comment common.NullableString `json:"comment,omitempty"`
	// Whether the index is hidden
	Hidden *bool `json:"hidden,omitempty"`
	// Other options for the index
	OtherOptions common.NullableList[string] `json:"otherOptions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsTableIndex instantiates a new DmsTableIndex object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsTableIndex() *DmsTableIndex {
	this := DmsTableIndex{}
	return &this
}

// NewDmsTableIndexWithDefaults instantiates a new DmsTableIndex object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsTableIndexWithDefaults() *DmsTableIndex {
	this := DmsTableIndex{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DmsTableIndex) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableIndex) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DmsTableIndex) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DmsTableIndex) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DmsTableIndex) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableIndex) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DmsTableIndex) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DmsTableIndex) SetType(v string) {
	o.Type = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *DmsTableIndex) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableIndex) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *DmsTableIndex) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *DmsTableIndex) SetColumns(v []string) {
	o.Columns = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableIndex) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableIndex) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *DmsTableIndex) HasComment() bool {
	return o != nil && o.Comment.IsSet()
}

// SetComment gets a reference to the given common.NullableString and assigns it to the Comment field.
func (o *DmsTableIndex) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil.
func (o *DmsTableIndex) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil.
func (o *DmsTableIndex) UnsetComment() {
	o.Comment.Unset()
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *DmsTableIndex) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableIndex) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *DmsTableIndex) HasHidden() bool {
	return o != nil && o.Hidden != nil
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *DmsTableIndex) SetHidden(v bool) {
	o.Hidden = &v
}

// GetOtherOptions returns the OtherOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableIndex) GetOtherOptions() []string {
	if o == nil || o.OtherOptions.Get() == nil {
		var ret []string
		return ret
	}
	return *o.OtherOptions.Get()
}

// GetOtherOptionsOk returns a tuple with the OtherOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableIndex) GetOtherOptionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OtherOptions.Get(), o.OtherOptions.IsSet()
}

// HasOtherOptions returns a boolean if a field has been set.
func (o *DmsTableIndex) HasOtherOptions() bool {
	return o != nil && o.OtherOptions.IsSet()
}

// SetOtherOptions gets a reference to the given common.NullableList[string] and assigns it to the OtherOptions field.
func (o *DmsTableIndex) SetOtherOptions(v []string) {
	o.OtherOptions.Set(&v)
}

// SetOtherOptionsNil sets the value for OtherOptions to be an explicit nil.
func (o *DmsTableIndex) SetOtherOptionsNil() {
	o.OtherOptions.Set(nil)
}

// UnsetOtherOptions ensures that no value is present for OtherOptions, not even an explicit nil.
func (o *DmsTableIndex) UnsetOtherOptions() {
	o.OtherOptions.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsTableIndex) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.OtherOptions.IsSet() {
		toSerialize["otherOptions"] = o.OtherOptions.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsTableIndex) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         *string                     `json:"name,omitempty"`
		Type         *string                     `json:"type,omitempty"`
		Columns      []string                    `json:"columns,omitempty"`
		Comment      common.NullableString       `json:"comment,omitempty"`
		Hidden       *bool                       `json:"hidden,omitempty"`
		OtherOptions common.NullableList[string] `json:"otherOptions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "type", "columns", "comment", "hidden", "otherOptions"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Type = all.Type
	o.Columns = all.Columns
	o.Comment = all.Comment
	o.Hidden = all.Hidden
	o.OtherOptions = all.OtherOptions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
