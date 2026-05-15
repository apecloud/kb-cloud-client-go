// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoIndexInfo struct {
	Name   *string     `json:"name,omitempty"`
	Key    interface{} `json:"key,omitempty"`
	Type   *string     `json:"type,omitempty"`
	Unique *bool       `json:"unique,omitempty"`
	Sparse *bool       `json:"sparse,omitempty"`
	Ttl    *int32      `json:"ttl,omitempty"`
	Size   *int64      `json:"size,omitempty"`
	Usage  interface{} `json:"usage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoIndexInfo instantiates a new MongoIndexInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoIndexInfo() *MongoIndexInfo {
	this := MongoIndexInfo{}
	return &this
}

// NewMongoIndexInfoWithDefaults instantiates a new MongoIndexInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoIndexInfoWithDefaults() *MongoIndexInfo {
	this := MongoIndexInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MongoIndexInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoIndexInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MongoIndexInfo) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MongoIndexInfo) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *MongoIndexInfo) GetKey() interface{} {
	if o == nil || o.Key == nil {
		var ret interface{}
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoIndexInfo) GetKeyOk() (*interface{}, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return &o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *MongoIndexInfo) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given interface{} and assigns it to the Key field.
func (o *MongoIndexInfo) SetKey(v interface{}) {
	o.Key = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MongoIndexInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoIndexInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MongoIndexInfo) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MongoIndexInfo) SetType(v string) {
	o.Type = &v
}

// GetUnique returns the Unique field value if set, zero value otherwise.
func (o *MongoIndexInfo) GetUnique() bool {
	if o == nil || o.Unique == nil {
		var ret bool
		return ret
	}
	return *o.Unique
}

// GetUniqueOk returns a tuple with the Unique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoIndexInfo) GetUniqueOk() (*bool, bool) {
	if o == nil || o.Unique == nil {
		return nil, false
	}
	return o.Unique, true
}

// HasUnique returns a boolean if a field has been set.
func (o *MongoIndexInfo) HasUnique() bool {
	return o != nil && o.Unique != nil
}

// SetUnique gets a reference to the given bool and assigns it to the Unique field.
func (o *MongoIndexInfo) SetUnique(v bool) {
	o.Unique = &v
}

// GetSparse returns the Sparse field value if set, zero value otherwise.
func (o *MongoIndexInfo) GetSparse() bool {
	if o == nil || o.Sparse == nil {
		var ret bool
		return ret
	}
	return *o.Sparse
}

// GetSparseOk returns a tuple with the Sparse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoIndexInfo) GetSparseOk() (*bool, bool) {
	if o == nil || o.Sparse == nil {
		return nil, false
	}
	return o.Sparse, true
}

// HasSparse returns a boolean if a field has been set.
func (o *MongoIndexInfo) HasSparse() bool {
	return o != nil && o.Sparse != nil
}

// SetSparse gets a reference to the given bool and assigns it to the Sparse field.
func (o *MongoIndexInfo) SetSparse(v bool) {
	o.Sparse = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *MongoIndexInfo) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoIndexInfo) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *MongoIndexInfo) HasTtl() bool {
	return o != nil && o.Ttl != nil
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *MongoIndexInfo) SetTtl(v int32) {
	o.Ttl = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *MongoIndexInfo) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoIndexInfo) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *MongoIndexInfo) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *MongoIndexInfo) SetSize(v int64) {
	o.Size = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *MongoIndexInfo) GetUsage() interface{} {
	if o == nil || o.Usage == nil {
		var ret interface{}
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoIndexInfo) GetUsageOk() (*interface{}, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return &o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *MongoIndexInfo) HasUsage() bool {
	return o != nil && o.Usage != nil
}

// SetUsage gets a reference to the given interface{} and assigns it to the Usage field.
func (o *MongoIndexInfo) SetUsage(v interface{}) {
	o.Usage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoIndexInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Unique != nil {
		toSerialize["unique"] = o.Unique
	}
	if o.Sparse != nil {
		toSerialize["sparse"] = o.Sparse
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoIndexInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name   *string     `json:"name,omitempty"`
		Key    interface{} `json:"key,omitempty"`
		Type   *string     `json:"type,omitempty"`
		Unique *bool       `json:"unique,omitempty"`
		Sparse *bool       `json:"sparse,omitempty"`
		Ttl    *int32      `json:"ttl,omitempty"`
		Size   *int64      `json:"size,omitempty"`
		Usage  interface{} `json:"usage,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "key", "type", "unique", "sparse", "ttl", "size", "usage"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Key = all.Key
	o.Type = all.Type
	o.Unique = all.Unique
	o.Sparse = all.Sparse
	o.Ttl = all.Ttl
	o.Size = all.Size
	o.Usage = all.Usage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
