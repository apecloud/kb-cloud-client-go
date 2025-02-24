// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DmsForeignKey struct {
	// The name of the foreign key
	Name common.NullableString `json:"name,omitempty"`
	// The list of columns that make up the foreign key
	Columns []string `json:"columns,omitempty"`
	// The reference details of the foreign key
	Reference *DmsForeignKeyReference `json:"reference,omitempty"`
	// The action on update (e.g., CASCADE, NO ACTION)
	OnUpdate *string `json:"onUpdate,omitempty"`
	// The action on delete (e.g., CASCADE, NO ACTION)
	OnDelete *string `json:"onDelete,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsForeignKey instantiates a new DmsForeignKey object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsForeignKey() *DmsForeignKey {
	this := DmsForeignKey{}
	return &this
}

// NewDmsForeignKeyWithDefaults instantiates a new DmsForeignKey object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsForeignKeyWithDefaults() *DmsForeignKey {
	this := DmsForeignKey{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsForeignKey) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsForeignKey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DmsForeignKey) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given common.NullableString and assigns it to the Name field.
func (o *DmsForeignKey) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *DmsForeignKey) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *DmsForeignKey) UnsetName() {
	o.Name.Unset()
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *DmsForeignKey) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsForeignKey) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *DmsForeignKey) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *DmsForeignKey) SetColumns(v []string) {
	o.Columns = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *DmsForeignKey) GetReference() DmsForeignKeyReference {
	if o == nil || o.Reference == nil {
		var ret DmsForeignKeyReference
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsForeignKey) GetReferenceOk() (*DmsForeignKeyReference, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *DmsForeignKey) HasReference() bool {
	return o != nil && o.Reference != nil
}

// SetReference gets a reference to the given DmsForeignKeyReference and assigns it to the Reference field.
func (o *DmsForeignKey) SetReference(v DmsForeignKeyReference) {
	o.Reference = &v
}

// GetOnUpdate returns the OnUpdate field value if set, zero value otherwise.
func (o *DmsForeignKey) GetOnUpdate() string {
	if o == nil || o.OnUpdate == nil {
		var ret string
		return ret
	}
	return *o.OnUpdate
}

// GetOnUpdateOk returns a tuple with the OnUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsForeignKey) GetOnUpdateOk() (*string, bool) {
	if o == nil || o.OnUpdate == nil {
		return nil, false
	}
	return o.OnUpdate, true
}

// HasOnUpdate returns a boolean if a field has been set.
func (o *DmsForeignKey) HasOnUpdate() bool {
	return o != nil && o.OnUpdate != nil
}

// SetOnUpdate gets a reference to the given string and assigns it to the OnUpdate field.
func (o *DmsForeignKey) SetOnUpdate(v string) {
	o.OnUpdate = &v
}

// GetOnDelete returns the OnDelete field value if set, zero value otherwise.
func (o *DmsForeignKey) GetOnDelete() string {
	if o == nil || o.OnDelete == nil {
		var ret string
		return ret
	}
	return *o.OnDelete
}

// GetOnDeleteOk returns a tuple with the OnDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsForeignKey) GetOnDeleteOk() (*string, bool) {
	if o == nil || o.OnDelete == nil {
		return nil, false
	}
	return o.OnDelete, true
}

// HasOnDelete returns a boolean if a field has been set.
func (o *DmsForeignKey) HasOnDelete() bool {
	return o != nil && o.OnDelete != nil
}

// SetOnDelete gets a reference to the given string and assigns it to the OnDelete field.
func (o *DmsForeignKey) SetOnDelete(v string) {
	o.OnDelete = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsForeignKey) MarshalJSON() ([]byte, error) {
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
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.OnUpdate != nil {
		toSerialize["onUpdate"] = o.OnUpdate
	}
	if o.OnDelete != nil {
		toSerialize["onDelete"] = o.OnDelete
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsForeignKey) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name      common.NullableString   `json:"name,omitempty"`
		Columns   []string                `json:"columns,omitempty"`
		Reference *DmsForeignKeyReference `json:"reference,omitempty"`
		OnUpdate  *string                 `json:"onUpdate,omitempty"`
		OnDelete  *string                 `json:"onDelete,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "columns", "reference", "onUpdate", "onDelete"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	o.Columns = all.Columns
	if all.Reference != nil && all.Reference.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Reference = all.Reference
	o.OnUpdate = all.OnUpdate
	o.OnDelete = all.OnDelete

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
