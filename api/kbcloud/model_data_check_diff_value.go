// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataCheckDiffValue struct {
	Src     common.NullableString `json:"src,omitempty"`
	Dst     common.NullableString `json:"dst,omitempty"`
	SrcType common.NullableString `json:"src_type,omitempty"`
	DstType common.NullableString `json:"dst_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataCheckDiffValue instantiates a new DataCheckDiffValue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataCheckDiffValue() *DataCheckDiffValue {
	this := DataCheckDiffValue{}
	return &this
}

// NewDataCheckDiffValueWithDefaults instantiates a new DataCheckDiffValue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataCheckDiffValueWithDefaults() *DataCheckDiffValue {
	this := DataCheckDiffValue{}
	return &this
}

// GetSrc returns the Src field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataCheckDiffValue) GetSrc() string {
	if o == nil || o.Src.Get() == nil {
		var ret string
		return ret
	}
	return *o.Src.Get()
}

// GetSrcOk returns a tuple with the Src field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataCheckDiffValue) GetSrcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Src.Get(), o.Src.IsSet()
}

// HasSrc returns a boolean if a field has been set.
func (o *DataCheckDiffValue) HasSrc() bool {
	return o != nil && o.Src.IsSet()
}

// SetSrc gets a reference to the given common.NullableString and assigns it to the Src field.
func (o *DataCheckDiffValue) SetSrc(v string) {
	o.Src.Set(&v)
}

// SetSrcNil sets the value for Src to be an explicit nil.
func (o *DataCheckDiffValue) SetSrcNil() {
	o.Src.Set(nil)
}

// UnsetSrc ensures that no value is present for Src, not even an explicit nil.
func (o *DataCheckDiffValue) UnsetSrc() {
	o.Src.Unset()
}

// GetDst returns the Dst field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataCheckDiffValue) GetDst() string {
	if o == nil || o.Dst.Get() == nil {
		var ret string
		return ret
	}
	return *o.Dst.Get()
}

// GetDstOk returns a tuple with the Dst field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataCheckDiffValue) GetDstOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dst.Get(), o.Dst.IsSet()
}

// HasDst returns a boolean if a field has been set.
func (o *DataCheckDiffValue) HasDst() bool {
	return o != nil && o.Dst.IsSet()
}

// SetDst gets a reference to the given common.NullableString and assigns it to the Dst field.
func (o *DataCheckDiffValue) SetDst(v string) {
	o.Dst.Set(&v)
}

// SetDstNil sets the value for Dst to be an explicit nil.
func (o *DataCheckDiffValue) SetDstNil() {
	o.Dst.Set(nil)
}

// UnsetDst ensures that no value is present for Dst, not even an explicit nil.
func (o *DataCheckDiffValue) UnsetDst() {
	o.Dst.Unset()
}

// GetSrcType returns the SrcType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataCheckDiffValue) GetSrcType() string {
	if o == nil || o.SrcType.Get() == nil {
		var ret string
		return ret
	}
	return *o.SrcType.Get()
}

// GetSrcTypeOk returns a tuple with the SrcType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataCheckDiffValue) GetSrcTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SrcType.Get(), o.SrcType.IsSet()
}

// HasSrcType returns a boolean if a field has been set.
func (o *DataCheckDiffValue) HasSrcType() bool {
	return o != nil && o.SrcType.IsSet()
}

// SetSrcType gets a reference to the given common.NullableString and assigns it to the SrcType field.
func (o *DataCheckDiffValue) SetSrcType(v string) {
	o.SrcType.Set(&v)
}

// SetSrcTypeNil sets the value for SrcType to be an explicit nil.
func (o *DataCheckDiffValue) SetSrcTypeNil() {
	o.SrcType.Set(nil)
}

// UnsetSrcType ensures that no value is present for SrcType, not even an explicit nil.
func (o *DataCheckDiffValue) UnsetSrcType() {
	o.SrcType.Unset()
}

// GetDstType returns the DstType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataCheckDiffValue) GetDstType() string {
	if o == nil || o.DstType.Get() == nil {
		var ret string
		return ret
	}
	return *o.DstType.Get()
}

// GetDstTypeOk returns a tuple with the DstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataCheckDiffValue) GetDstTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DstType.Get(), o.DstType.IsSet()
}

// HasDstType returns a boolean if a field has been set.
func (o *DataCheckDiffValue) HasDstType() bool {
	return o != nil && o.DstType.IsSet()
}

// SetDstType gets a reference to the given common.NullableString and assigns it to the DstType field.
func (o *DataCheckDiffValue) SetDstType(v string) {
	o.DstType.Set(&v)
}

// SetDstTypeNil sets the value for DstType to be an explicit nil.
func (o *DataCheckDiffValue) SetDstTypeNil() {
	o.DstType.Set(nil)
}

// UnsetDstType ensures that no value is present for DstType, not even an explicit nil.
func (o *DataCheckDiffValue) UnsetDstType() {
	o.DstType.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DataCheckDiffValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Src.IsSet() {
		toSerialize["src"] = o.Src.Get()
	}
	if o.Dst.IsSet() {
		toSerialize["dst"] = o.Dst.Get()
	}
	if o.SrcType.IsSet() {
		toSerialize["src_type"] = o.SrcType.Get()
	}
	if o.DstType.IsSet() {
		toSerialize["dst_type"] = o.DstType.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataCheckDiffValue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Src     common.NullableString `json:"src,omitempty"`
		Dst     common.NullableString `json:"dst,omitempty"`
		SrcType common.NullableString `json:"src_type,omitempty"`
		DstType common.NullableString `json:"dst_type,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"src", "dst", "src_type", "dst_type"})
	} else {
		return err
	}
	o.Src = all.Src
	o.Dst = all.Dst
	o.SrcType = all.SrcType
	o.DstType = all.DstType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
