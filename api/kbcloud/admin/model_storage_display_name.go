// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// StorageDisplayName the storage displayName
type StorageDisplayName struct {
	Zh *string `json:"zh,omitempty"`
	En *string `json:"en,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageDisplayName instantiates a new StorageDisplayName object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageDisplayName() *StorageDisplayName {
	this := StorageDisplayName{}
	return &this
}

// NewStorageDisplayNameWithDefaults instantiates a new StorageDisplayName object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageDisplayNameWithDefaults() *StorageDisplayName {
	this := StorageDisplayName{}
	return &this
}

// GetZh returns the Zh field value if set, zero value otherwise.
func (o *StorageDisplayName) GetZh() string {
	if o == nil || o.Zh == nil {
		var ret string
		return ret
	}
	return *o.Zh
}

// GetZhOk returns a tuple with the Zh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDisplayName) GetZhOk() (*string, bool) {
	if o == nil || o.Zh == nil {
		return nil, false
	}
	return o.Zh, true
}

// HasZh returns a boolean if a field has been set.
func (o *StorageDisplayName) HasZh() bool {
	return o != nil && o.Zh != nil
}

// SetZh gets a reference to the given string and assigns it to the Zh field.
func (o *StorageDisplayName) SetZh(v string) {
	o.Zh = &v
}

// GetEn returns the En field value if set, zero value otherwise.
func (o *StorageDisplayName) GetEn() string {
	if o == nil || o.En == nil {
		var ret string
		return ret
	}
	return *o.En
}

// GetEnOk returns a tuple with the En field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDisplayName) GetEnOk() (*string, bool) {
	if o == nil || o.En == nil {
		return nil, false
	}
	return o.En, true
}

// HasEn returns a boolean if a field has been set.
func (o *StorageDisplayName) HasEn() bool {
	return o != nil && o.En != nil
}

// SetEn gets a reference to the given string and assigns it to the En field.
func (o *StorageDisplayName) SetEn(v string) {
	o.En = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageDisplayName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Zh != nil {
		toSerialize["zh"] = o.Zh
	}
	if o.En != nil {
		toSerialize["en"] = o.En
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageDisplayName) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Zh *string `json:"zh,omitempty"`
		En *string `json:"en,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"zh", "en"})
	} else {
		return err
	}
	o.Zh = all.Zh
	o.En = all.En

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
