// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DmsVolume struct {
	Name      *string `json:"name,omitempty"`
	Type      *string `json:"type,omitempty"`
	IsDefault *string `json:"isDefault,omitempty"`
	Location  *string `json:"location,omitempty"`
	Params    *string `json:"params,omitempty"`
	Comment   *string `json:"comment,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsVolume instantiates a new DmsVolume object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsVolume() *DmsVolume {
	this := DmsVolume{}
	return &this
}

// NewDmsVolumeWithDefaults instantiates a new DmsVolume object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsVolumeWithDefaults() *DmsVolume {
	this := DmsVolume{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DmsVolume) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsVolume) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DmsVolume) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DmsVolume) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DmsVolume) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsVolume) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DmsVolume) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DmsVolume) SetType(v string) {
	o.Type = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *DmsVolume) GetIsDefault() string {
	if o == nil || o.IsDefault == nil {
		var ret string
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsVolume) GetIsDefaultOk() (*string, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *DmsVolume) HasIsDefault() bool {
	return o != nil && o.IsDefault != nil
}

// SetIsDefault gets a reference to the given string and assigns it to the IsDefault field.
func (o *DmsVolume) SetIsDefault(v string) {
	o.IsDefault = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *DmsVolume) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsVolume) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *DmsVolume) HasLocation() bool {
	return o != nil && o.Location != nil
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *DmsVolume) SetLocation(v string) {
	o.Location = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *DmsVolume) GetParams() string {
	if o == nil || o.Params == nil {
		var ret string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsVolume) GetParamsOk() (*string, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *DmsVolume) HasParams() bool {
	return o != nil && o.Params != nil
}

// SetParams gets a reference to the given string and assigns it to the Params field.
func (o *DmsVolume) SetParams(v string) {
	o.Params = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *DmsVolume) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsVolume) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *DmsVolume) HasComment() bool {
	return o != nil && o.Comment != nil
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *DmsVolume) SetComment(v string) {
	o.Comment = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsVolume) MarshalJSON() ([]byte, error) {
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
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsVolume) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name      *string `json:"name,omitempty"`
		Type      *string `json:"type,omitempty"`
		IsDefault *string `json:"isDefault,omitempty"`
		Location  *string `json:"location,omitempty"`
		Params    *string `json:"params,omitempty"`
		Comment   *string `json:"comment,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "type", "isDefault", "location", "params", "comment"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Type = all.Type
	o.IsDefault = all.IsDefault
	o.Location = all.Location
	o.Params = all.Params
	o.Comment = all.Comment

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
