// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

type VolumeExpand struct {
	// the new volume size of cluster
	Size *string `json:"size,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVolumeExpand instantiates a new VolumeExpand object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVolumeExpand() *VolumeExpand {
	this := VolumeExpand{}
	return &this
}

// NewVolumeExpandWithDefaults instantiates a new VolumeExpand object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVolumeExpandWithDefaults() *VolumeExpand {
	this := VolumeExpand{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *VolumeExpand) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeExpand) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *VolumeExpand) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *VolumeExpand) SetSize(v string) {
	o.Size = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o VolumeExpand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VolumeExpand) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Size *string `json:"size,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"size"})
	} else {
		return err
	}
	o.Size = all.Size

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
