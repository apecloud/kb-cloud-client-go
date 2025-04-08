// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// IoQuantity IO Quantity describes IOPS and BPS of a volume
type IoQuantity struct {
	ReadIops  common.NullableInt64 `json:"readIOPS,omitempty"`
	WriteIops common.NullableInt64 `json:"writeIOPS,omitempty"`
	ReadBps   common.NullableInt64 `json:"readBPS,omitempty"`
	WriteBps  common.NullableInt64 `json:"writeBPS,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIoQuantity instantiates a new IoQuantity object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIoQuantity() *IoQuantity {
	this := IoQuantity{}
	return &this
}

// NewIoQuantityWithDefaults instantiates a new IoQuantity object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIoQuantityWithDefaults() *IoQuantity {
	this := IoQuantity{}
	return &this
}

// GetReadIops returns the ReadIops field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IoQuantity) GetReadIops() int64 {
	if o == nil || o.ReadIops.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ReadIops.Get()
}

// GetReadIopsOk returns a tuple with the ReadIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IoQuantity) GetReadIopsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReadIops.Get(), o.ReadIops.IsSet()
}

// HasReadIops returns a boolean if a field has been set.
func (o *IoQuantity) HasReadIops() bool {
	return o != nil && o.ReadIops.IsSet()
}

// SetReadIops gets a reference to the given common.NullableInt64 and assigns it to the ReadIops field.
func (o *IoQuantity) SetReadIops(v int64) {
	o.ReadIops.Set(&v)
}

// SetReadIopsNil sets the value for ReadIops to be an explicit nil.
func (o *IoQuantity) SetReadIopsNil() {
	o.ReadIops.Set(nil)
}

// UnsetReadIops ensures that no value is present for ReadIops, not even an explicit nil.
func (o *IoQuantity) UnsetReadIops() {
	o.ReadIops.Unset()
}

// GetWriteIops returns the WriteIops field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IoQuantity) GetWriteIops() int64 {
	if o == nil || o.WriteIops.Get() == nil {
		var ret int64
		return ret
	}
	return *o.WriteIops.Get()
}

// GetWriteIopsOk returns a tuple with the WriteIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IoQuantity) GetWriteIopsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.WriteIops.Get(), o.WriteIops.IsSet()
}

// HasWriteIops returns a boolean if a field has been set.
func (o *IoQuantity) HasWriteIops() bool {
	return o != nil && o.WriteIops.IsSet()
}

// SetWriteIops gets a reference to the given common.NullableInt64 and assigns it to the WriteIops field.
func (o *IoQuantity) SetWriteIops(v int64) {
	o.WriteIops.Set(&v)
}

// SetWriteIopsNil sets the value for WriteIops to be an explicit nil.
func (o *IoQuantity) SetWriteIopsNil() {
	o.WriteIops.Set(nil)
}

// UnsetWriteIops ensures that no value is present for WriteIops, not even an explicit nil.
func (o *IoQuantity) UnsetWriteIops() {
	o.WriteIops.Unset()
}

// GetReadBps returns the ReadBps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IoQuantity) GetReadBps() int64 {
	if o == nil || o.ReadBps.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ReadBps.Get()
}

// GetReadBpsOk returns a tuple with the ReadBps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IoQuantity) GetReadBpsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReadBps.Get(), o.ReadBps.IsSet()
}

// HasReadBps returns a boolean if a field has been set.
func (o *IoQuantity) HasReadBps() bool {
	return o != nil && o.ReadBps.IsSet()
}

// SetReadBps gets a reference to the given common.NullableInt64 and assigns it to the ReadBps field.
func (o *IoQuantity) SetReadBps(v int64) {
	o.ReadBps.Set(&v)
}

// SetReadBpsNil sets the value for ReadBps to be an explicit nil.
func (o *IoQuantity) SetReadBpsNil() {
	o.ReadBps.Set(nil)
}

// UnsetReadBps ensures that no value is present for ReadBps, not even an explicit nil.
func (o *IoQuantity) UnsetReadBps() {
	o.ReadBps.Unset()
}

// GetWriteBps returns the WriteBps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IoQuantity) GetWriteBps() int64 {
	if o == nil || o.WriteBps.Get() == nil {
		var ret int64
		return ret
	}
	return *o.WriteBps.Get()
}

// GetWriteBpsOk returns a tuple with the WriteBps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IoQuantity) GetWriteBpsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.WriteBps.Get(), o.WriteBps.IsSet()
}

// HasWriteBps returns a boolean if a field has been set.
func (o *IoQuantity) HasWriteBps() bool {
	return o != nil && o.WriteBps.IsSet()
}

// SetWriteBps gets a reference to the given common.NullableInt64 and assigns it to the WriteBps field.
func (o *IoQuantity) SetWriteBps(v int64) {
	o.WriteBps.Set(&v)
}

// SetWriteBpsNil sets the value for WriteBps to be an explicit nil.
func (o *IoQuantity) SetWriteBpsNil() {
	o.WriteBps.Set(nil)
}

// UnsetWriteBps ensures that no value is present for WriteBps, not even an explicit nil.
func (o *IoQuantity) UnsetWriteBps() {
	o.WriteBps.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o IoQuantity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ReadIops.IsSet() {
		toSerialize["readIOPS"] = o.ReadIops.Get()
	}
	if o.WriteIops.IsSet() {
		toSerialize["writeIOPS"] = o.WriteIops.Get()
	}
	if o.ReadBps.IsSet() {
		toSerialize["readBPS"] = o.ReadBps.Get()
	}
	if o.WriteBps.IsSet() {
		toSerialize["writeBPS"] = o.WriteBps.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IoQuantity) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ReadIops  common.NullableInt64 `json:"readIOPS,omitempty"`
		WriteIops common.NullableInt64 `json:"writeIOPS,omitempty"`
		ReadBps   common.NullableInt64 `json:"readBPS,omitempty"`
		WriteBps  common.NullableInt64 `json:"writeBPS,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"readIOPS", "writeIOPS", "readBPS", "writeBPS"})
	} else {
		return err
	}
	o.ReadIops = all.ReadIops
	o.WriteIops = all.WriteIops
	o.ReadBps = all.ReadBps
	o.WriteBps = all.WriteBps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
