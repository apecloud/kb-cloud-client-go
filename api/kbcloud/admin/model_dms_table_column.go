// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsTableColumn struct {
	// The name of the column
	Name *string `json:"name,omitempty"`
	// The data type of the column
	Type *string `json:"type,omitempty"`
	// The length of the column if applicable
	Length common.NullableInt32 `json:"length,omitempty"`
	// The number of decimal places for numeric columns
	Scale common.NullableInt32 `json:"scale,omitempty"`
	// Whether the column can be null
	NotNull *bool `json:"notNull,omitempty"`
	// The default value for the column
	Default common.NullableString `json:"default,omitempty"`
	// A comment describing the column
	Comment common.NullableString `json:"comment,omitempty"`
	// Whether the column is auto-incremented
	AutoIncrement *bool `json:"autoIncrement,omitempty"`
	// used for autoIncrement
	Seed common.NullableInt32 `json:"seed,omitempty"`
	// used for autoIncrement
	Increment common.NullableInt32 `json:"increment,omitempty"`
	// Whether the column is unsigned
	Unsigned *bool `json:"unsigned,omitempty"`
	// Whether the column updates to the current timestamp on update
	OnUpdateCurrentTimestamp *bool `json:"onUpdateCurrentTimestamp,omitempty"`
	// Whether the column is hidden
	Hidden *bool `json:"hidden,omitempty"`
	// Generated column information
	Generated NullableDmsTableColumnGenerated `json:"generated,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsTableColumn instantiates a new DmsTableColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsTableColumn() *DmsTableColumn {
	this := DmsTableColumn{}
	return &this
}

// NewDmsTableColumnWithDefaults instantiates a new DmsTableColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsTableColumnWithDefaults() *DmsTableColumn {
	this := DmsTableColumn{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DmsTableColumn) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableColumn) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DmsTableColumn) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DmsTableColumn) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DmsTableColumn) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableColumn) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DmsTableColumn) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DmsTableColumn) SetType(v string) {
	o.Type = &v
}

// GetLength returns the Length field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableColumn) GetLength() int32 {
	if o == nil || o.Length.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Length.Get()
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableColumn) GetLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Length.Get(), o.Length.IsSet()
}

// HasLength returns a boolean if a field has been set.
func (o *DmsTableColumn) HasLength() bool {
	return o != nil && o.Length.IsSet()
}

// SetLength gets a reference to the given common.NullableInt32 and assigns it to the Length field.
func (o *DmsTableColumn) SetLength(v int32) {
	o.Length.Set(&v)
}

// SetLengthNil sets the value for Length to be an explicit nil.
func (o *DmsTableColumn) SetLengthNil() {
	o.Length.Set(nil)
}

// UnsetLength ensures that no value is present for Length, not even an explicit nil.
func (o *DmsTableColumn) UnsetLength() {
	o.Length.Unset()
}

// GetScale returns the Scale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableColumn) GetScale() int32 {
	if o == nil || o.Scale.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Scale.Get()
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableColumn) GetScaleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scale.Get(), o.Scale.IsSet()
}

// HasScale returns a boolean if a field has been set.
func (o *DmsTableColumn) HasScale() bool {
	return o != nil && o.Scale.IsSet()
}

// SetScale gets a reference to the given common.NullableInt32 and assigns it to the Scale field.
func (o *DmsTableColumn) SetScale(v int32) {
	o.Scale.Set(&v)
}

// SetScaleNil sets the value for Scale to be an explicit nil.
func (o *DmsTableColumn) SetScaleNil() {
	o.Scale.Set(nil)
}

// UnsetScale ensures that no value is present for Scale, not even an explicit nil.
func (o *DmsTableColumn) UnsetScale() {
	o.Scale.Unset()
}

// GetNotNull returns the NotNull field value if set, zero value otherwise.
func (o *DmsTableColumn) GetNotNull() bool {
	if o == nil || o.NotNull == nil {
		var ret bool
		return ret
	}
	return *o.NotNull
}

// GetNotNullOk returns a tuple with the NotNull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableColumn) GetNotNullOk() (*bool, bool) {
	if o == nil || o.NotNull == nil {
		return nil, false
	}
	return o.NotNull, true
}

// HasNotNull returns a boolean if a field has been set.
func (o *DmsTableColumn) HasNotNull() bool {
	return o != nil && o.NotNull != nil
}

// SetNotNull gets a reference to the given bool and assigns it to the NotNull field.
func (o *DmsTableColumn) SetNotNull(v bool) {
	o.NotNull = &v
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableColumn) GetDefault() string {
	if o == nil || o.Default.Get() == nil {
		var ret string
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableColumn) GetDefaultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *DmsTableColumn) HasDefault() bool {
	return o != nil && o.Default.IsSet()
}

// SetDefault gets a reference to the given common.NullableString and assigns it to the Default field.
func (o *DmsTableColumn) SetDefault(v string) {
	o.Default.Set(&v)
}

// SetDefaultNil sets the value for Default to be an explicit nil.
func (o *DmsTableColumn) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil.
func (o *DmsTableColumn) UnsetDefault() {
	o.Default.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableColumn) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableColumn) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *DmsTableColumn) HasComment() bool {
	return o != nil && o.Comment.IsSet()
}

// SetComment gets a reference to the given common.NullableString and assigns it to the Comment field.
func (o *DmsTableColumn) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil.
func (o *DmsTableColumn) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil.
func (o *DmsTableColumn) UnsetComment() {
	o.Comment.Unset()
}

// GetAutoIncrement returns the AutoIncrement field value if set, zero value otherwise.
func (o *DmsTableColumn) GetAutoIncrement() bool {
	if o == nil || o.AutoIncrement == nil {
		var ret bool
		return ret
	}
	return *o.AutoIncrement
}

// GetAutoIncrementOk returns a tuple with the AutoIncrement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableColumn) GetAutoIncrementOk() (*bool, bool) {
	if o == nil || o.AutoIncrement == nil {
		return nil, false
	}
	return o.AutoIncrement, true
}

// HasAutoIncrement returns a boolean if a field has been set.
func (o *DmsTableColumn) HasAutoIncrement() bool {
	return o != nil && o.AutoIncrement != nil
}

// SetAutoIncrement gets a reference to the given bool and assigns it to the AutoIncrement field.
func (o *DmsTableColumn) SetAutoIncrement(v bool) {
	o.AutoIncrement = &v
}

// GetSeed returns the Seed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableColumn) GetSeed() int32 {
	if o == nil || o.Seed.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Seed.Get()
}

// GetSeedOk returns a tuple with the Seed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableColumn) GetSeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Seed.Get(), o.Seed.IsSet()
}

// HasSeed returns a boolean if a field has been set.
func (o *DmsTableColumn) HasSeed() bool {
	return o != nil && o.Seed.IsSet()
}

// SetSeed gets a reference to the given common.NullableInt32 and assigns it to the Seed field.
func (o *DmsTableColumn) SetSeed(v int32) {
	o.Seed.Set(&v)
}

// SetSeedNil sets the value for Seed to be an explicit nil.
func (o *DmsTableColumn) SetSeedNil() {
	o.Seed.Set(nil)
}

// UnsetSeed ensures that no value is present for Seed, not even an explicit nil.
func (o *DmsTableColumn) UnsetSeed() {
	o.Seed.Unset()
}

// GetIncrement returns the Increment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableColumn) GetIncrement() int32 {
	if o == nil || o.Increment.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Increment.Get()
}

// GetIncrementOk returns a tuple with the Increment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableColumn) GetIncrementOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Increment.Get(), o.Increment.IsSet()
}

// HasIncrement returns a boolean if a field has been set.
func (o *DmsTableColumn) HasIncrement() bool {
	return o != nil && o.Increment.IsSet()
}

// SetIncrement gets a reference to the given common.NullableInt32 and assigns it to the Increment field.
func (o *DmsTableColumn) SetIncrement(v int32) {
	o.Increment.Set(&v)
}

// SetIncrementNil sets the value for Increment to be an explicit nil.
func (o *DmsTableColumn) SetIncrementNil() {
	o.Increment.Set(nil)
}

// UnsetIncrement ensures that no value is present for Increment, not even an explicit nil.
func (o *DmsTableColumn) UnsetIncrement() {
	o.Increment.Unset()
}

// GetUnsigned returns the Unsigned field value if set, zero value otherwise.
func (o *DmsTableColumn) GetUnsigned() bool {
	if o == nil || o.Unsigned == nil {
		var ret bool
		return ret
	}
	return *o.Unsigned
}

// GetUnsignedOk returns a tuple with the Unsigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableColumn) GetUnsignedOk() (*bool, bool) {
	if o == nil || o.Unsigned == nil {
		return nil, false
	}
	return o.Unsigned, true
}

// HasUnsigned returns a boolean if a field has been set.
func (o *DmsTableColumn) HasUnsigned() bool {
	return o != nil && o.Unsigned != nil
}

// SetUnsigned gets a reference to the given bool and assigns it to the Unsigned field.
func (o *DmsTableColumn) SetUnsigned(v bool) {
	o.Unsigned = &v
}

// GetOnUpdateCurrentTimestamp returns the OnUpdateCurrentTimestamp field value if set, zero value otherwise.
func (o *DmsTableColumn) GetOnUpdateCurrentTimestamp() bool {
	if o == nil || o.OnUpdateCurrentTimestamp == nil {
		var ret bool
		return ret
	}
	return *o.OnUpdateCurrentTimestamp
}

// GetOnUpdateCurrentTimestampOk returns a tuple with the OnUpdateCurrentTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableColumn) GetOnUpdateCurrentTimestampOk() (*bool, bool) {
	if o == nil || o.OnUpdateCurrentTimestamp == nil {
		return nil, false
	}
	return o.OnUpdateCurrentTimestamp, true
}

// HasOnUpdateCurrentTimestamp returns a boolean if a field has been set.
func (o *DmsTableColumn) HasOnUpdateCurrentTimestamp() bool {
	return o != nil && o.OnUpdateCurrentTimestamp != nil
}

// SetOnUpdateCurrentTimestamp gets a reference to the given bool and assigns it to the OnUpdateCurrentTimestamp field.
func (o *DmsTableColumn) SetOnUpdateCurrentTimestamp(v bool) {
	o.OnUpdateCurrentTimestamp = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *DmsTableColumn) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableColumn) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *DmsTableColumn) HasHidden() bool {
	return o != nil && o.Hidden != nil
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *DmsTableColumn) SetHidden(v bool) {
	o.Hidden = &v
}

// GetGenerated returns the Generated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableColumn) GetGenerated() DmsTableColumnGenerated {
	if o == nil || o.Generated.Get() == nil {
		var ret DmsTableColumnGenerated
		return ret
	}
	return *o.Generated.Get()
}

// GetGeneratedOk returns a tuple with the Generated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableColumn) GetGeneratedOk() (*DmsTableColumnGenerated, bool) {
	if o == nil {
		return nil, false
	}
	return o.Generated.Get(), o.Generated.IsSet()
}

// HasGenerated returns a boolean if a field has been set.
func (o *DmsTableColumn) HasGenerated() bool {
	return o != nil && o.Generated.IsSet()
}

// SetGenerated gets a reference to the given NullableDmsTableColumnGenerated and assigns it to the Generated field.
func (o *DmsTableColumn) SetGenerated(v DmsTableColumnGenerated) {
	o.Generated.Set(&v)
}

// SetGeneratedNil sets the value for Generated to be an explicit nil.
func (o *DmsTableColumn) SetGeneratedNil() {
	o.Generated.Set(nil)
}

// UnsetGenerated ensures that no value is present for Generated, not even an explicit nil.
func (o *DmsTableColumn) UnsetGenerated() {
	o.Generated.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsTableColumn) MarshalJSON() ([]byte, error) {
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
	if o.Length.IsSet() {
		toSerialize["length"] = o.Length.Get()
	}
	if o.Scale.IsSet() {
		toSerialize["scale"] = o.Scale.Get()
	}
	if o.NotNull != nil {
		toSerialize["notNull"] = o.NotNull
	}
	if o.Default.IsSet() {
		toSerialize["default"] = o.Default.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.AutoIncrement != nil {
		toSerialize["autoIncrement"] = o.AutoIncrement
	}
	if o.Seed.IsSet() {
		toSerialize["seed"] = o.Seed.Get()
	}
	if o.Increment.IsSet() {
		toSerialize["increment"] = o.Increment.Get()
	}
	if o.Unsigned != nil {
		toSerialize["unsigned"] = o.Unsigned
	}
	if o.OnUpdateCurrentTimestamp != nil {
		toSerialize["onUpdateCurrentTimestamp"] = o.OnUpdateCurrentTimestamp
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.Generated.IsSet() {
		toSerialize["generated"] = o.Generated.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsTableColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                     *string                         `json:"name,omitempty"`
		Type                     *string                         `json:"type,omitempty"`
		Length                   common.NullableInt32            `json:"length,omitempty"`
		Scale                    common.NullableInt32            `json:"scale,omitempty"`
		NotNull                  *bool                           `json:"notNull,omitempty"`
		Default                  common.NullableString           `json:"default,omitempty"`
		Comment                  common.NullableString           `json:"comment,omitempty"`
		AutoIncrement            *bool                           `json:"autoIncrement,omitempty"`
		Seed                     common.NullableInt32            `json:"seed,omitempty"`
		Increment                common.NullableInt32            `json:"increment,omitempty"`
		Unsigned                 *bool                           `json:"unsigned,omitempty"`
		OnUpdateCurrentTimestamp *bool                           `json:"onUpdateCurrentTimestamp,omitempty"`
		Hidden                   *bool                           `json:"hidden,omitempty"`
		Generated                NullableDmsTableColumnGenerated `json:"generated,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "type", "length", "scale", "notNull", "default", "comment", "autoIncrement", "seed", "increment", "unsigned", "onUpdateCurrentTimestamp", "hidden", "generated"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Type = all.Type
	o.Length = all.Length
	o.Scale = all.Scale
	o.NotNull = all.NotNull
	o.Default = all.Default
	o.Comment = all.Comment
	o.AutoIncrement = all.AutoIncrement
	o.Seed = all.Seed
	o.Increment = all.Increment
	o.Unsigned = all.Unsigned
	o.OnUpdateCurrentTimestamp = all.OnUpdateCurrentTimestamp
	o.Hidden = all.Hidden
	o.Generated = all.Generated

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
