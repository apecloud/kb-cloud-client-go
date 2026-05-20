// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoDatabaseInfo struct {
	// database name
	Name            *string `json:"name,omitempty"`
	SizeOnDisk      *int64  `json:"sizeOnDisk,omitempty"`
	Empty           *bool   `json:"empty,omitempty"`
	CollectionCount *int64  `json:"collectionCount,omitempty"`
	IsSystem        *bool   `json:"isSystem,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoDatabaseInfo instantiates a new MongoDatabaseInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoDatabaseInfo() *MongoDatabaseInfo {
	this := MongoDatabaseInfo{}
	return &this
}

// NewMongoDatabaseInfoWithDefaults instantiates a new MongoDatabaseInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoDatabaseInfoWithDefaults() *MongoDatabaseInfo {
	this := MongoDatabaseInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MongoDatabaseInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDatabaseInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MongoDatabaseInfo) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MongoDatabaseInfo) SetName(v string) {
	o.Name = &v
}

// GetSizeOnDisk returns the SizeOnDisk field value if set, zero value otherwise.
func (o *MongoDatabaseInfo) GetSizeOnDisk() int64 {
	if o == nil || o.SizeOnDisk == nil {
		var ret int64
		return ret
	}
	return *o.SizeOnDisk
}

// GetSizeOnDiskOk returns a tuple with the SizeOnDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDatabaseInfo) GetSizeOnDiskOk() (*int64, bool) {
	if o == nil || o.SizeOnDisk == nil {
		return nil, false
	}
	return o.SizeOnDisk, true
}

// HasSizeOnDisk returns a boolean if a field has been set.
func (o *MongoDatabaseInfo) HasSizeOnDisk() bool {
	return o != nil && o.SizeOnDisk != nil
}

// SetSizeOnDisk gets a reference to the given int64 and assigns it to the SizeOnDisk field.
func (o *MongoDatabaseInfo) SetSizeOnDisk(v int64) {
	o.SizeOnDisk = &v
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *MongoDatabaseInfo) GetEmpty() bool {
	if o == nil || o.Empty == nil {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDatabaseInfo) GetEmptyOk() (*bool, bool) {
	if o == nil || o.Empty == nil {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *MongoDatabaseInfo) HasEmpty() bool {
	return o != nil && o.Empty != nil
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *MongoDatabaseInfo) SetEmpty(v bool) {
	o.Empty = &v
}

// GetCollectionCount returns the CollectionCount field value if set, zero value otherwise.
func (o *MongoDatabaseInfo) GetCollectionCount() int64 {
	if o == nil || o.CollectionCount == nil {
		var ret int64
		return ret
	}
	return *o.CollectionCount
}

// GetCollectionCountOk returns a tuple with the CollectionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDatabaseInfo) GetCollectionCountOk() (*int64, bool) {
	if o == nil || o.CollectionCount == nil {
		return nil, false
	}
	return o.CollectionCount, true
}

// HasCollectionCount returns a boolean if a field has been set.
func (o *MongoDatabaseInfo) HasCollectionCount() bool {
	return o != nil && o.CollectionCount != nil
}

// SetCollectionCount gets a reference to the given int64 and assigns it to the CollectionCount field.
func (o *MongoDatabaseInfo) SetCollectionCount(v int64) {
	o.CollectionCount = &v
}

// GetIsSystem returns the IsSystem field value if set, zero value otherwise.
func (o *MongoDatabaseInfo) GetIsSystem() bool {
	if o == nil || o.IsSystem == nil {
		var ret bool
		return ret
	}
	return *o.IsSystem
}

// GetIsSystemOk returns a tuple with the IsSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDatabaseInfo) GetIsSystemOk() (*bool, bool) {
	if o == nil || o.IsSystem == nil {
		return nil, false
	}
	return o.IsSystem, true
}

// HasIsSystem returns a boolean if a field has been set.
func (o *MongoDatabaseInfo) HasIsSystem() bool {
	return o != nil && o.IsSystem != nil
}

// SetIsSystem gets a reference to the given bool and assigns it to the IsSystem field.
func (o *MongoDatabaseInfo) SetIsSystem(v bool) {
	o.IsSystem = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoDatabaseInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SizeOnDisk != nil {
		toSerialize["sizeOnDisk"] = o.SizeOnDisk
	}
	if o.Empty != nil {
		toSerialize["empty"] = o.Empty
	}
	if o.CollectionCount != nil {
		toSerialize["collectionCount"] = o.CollectionCount
	}
	if o.IsSystem != nil {
		toSerialize["isSystem"] = o.IsSystem
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoDatabaseInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name            *string `json:"name,omitempty"`
		SizeOnDisk      *int64  `json:"sizeOnDisk,omitempty"`
		Empty           *bool   `json:"empty,omitempty"`
		CollectionCount *int64  `json:"collectionCount,omitempty"`
		IsSystem        *bool   `json:"isSystem,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "sizeOnDisk", "empty", "collectionCount", "isSystem"})
	} else {
		return err
	}
	o.Name = all.Name
	o.SizeOnDisk = all.SizeOnDisk
	o.Empty = all.Empty
	o.CollectionCount = all.CollectionCount
	o.IsSystem = all.IsSystem

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
