// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DeleteBenchOption struct {
	Ids []string `json:"ids,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeleteBenchOption instantiates a new DeleteBenchOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeleteBenchOption() *DeleteBenchOption {
	this := DeleteBenchOption{}
	return &this
}

// NewDeleteBenchOptionWithDefaults instantiates a new DeleteBenchOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeleteBenchOptionWithDefaults() *DeleteBenchOption {
	this := DeleteBenchOption{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *DeleteBenchOption) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteBenchOption) GetIdsOk() (*[]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return &o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *DeleteBenchOption) HasIds() bool {
	return o != nil && o.Ids != nil
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *DeleteBenchOption) SetIds(v []string) {
	o.Ids = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeleteBenchOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeleteBenchOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ids []string `json:"ids,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ids"})
	} else {
		return err
	}
	o.Ids = all.Ids

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
