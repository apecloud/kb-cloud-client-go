// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisKeyEntry struct {
	Key                *string `json:"Key,omitempty"`
	Bytes              *int64  `json:"Bytes,omitempty"`
	Type               *string `json:"Type,omitempty"`
	NumOfElem          *int64  `json:"NumOfElem,omitempty"`
	LenOfLargestElem   *int64  `json:"LenOfLargestElem,omitempty"`
	FieldOfLargestElem *string `json:"FieldOfLargestElem,omitempty"`
	Db                 *int32  `json:"Db,omitempty"`
	Encoding           *string `json:"Encoding,omitempty"`
	Expiration         *int64  `json:"Expiration,omitempty"`
	LruIdle            *int64  `json:"LruIdle,omitempty"`
	LfuFreq            *int32  `json:"LfuFreq,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisKeyEntry instantiates a new RedisKeyEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisKeyEntry() *RedisKeyEntry {
	this := RedisKeyEntry{}
	return &this
}

// NewRedisKeyEntryWithDefaults instantiates a new RedisKeyEntry object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisKeyEntryWithDefaults() *RedisKeyEntry {
	this := RedisKeyEntry{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RedisKeyEntry) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyEntry) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RedisKeyEntry) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RedisKeyEntry) SetKey(v string) {
	o.Key = &v
}

// GetBytes returns the Bytes field value if set, zero value otherwise.
func (o *RedisKeyEntry) GetBytes() int64 {
	if o == nil || o.Bytes == nil {
		var ret int64
		return ret
	}
	return *o.Bytes
}

// GetBytesOk returns a tuple with the Bytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyEntry) GetBytesOk() (*int64, bool) {
	if o == nil || o.Bytes == nil {
		return nil, false
	}
	return o.Bytes, true
}

// HasBytes returns a boolean if a field has been set.
func (o *RedisKeyEntry) HasBytes() bool {
	return o != nil && o.Bytes != nil
}

// SetBytes gets a reference to the given int64 and assigns it to the Bytes field.
func (o *RedisKeyEntry) SetBytes(v int64) {
	o.Bytes = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RedisKeyEntry) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyEntry) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RedisKeyEntry) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RedisKeyEntry) SetType(v string) {
	o.Type = &v
}

// GetNumOfElem returns the NumOfElem field value if set, zero value otherwise.
func (o *RedisKeyEntry) GetNumOfElem() int64 {
	if o == nil || o.NumOfElem == nil {
		var ret int64
		return ret
	}
	return *o.NumOfElem
}

// GetNumOfElemOk returns a tuple with the NumOfElem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyEntry) GetNumOfElemOk() (*int64, bool) {
	if o == nil || o.NumOfElem == nil {
		return nil, false
	}
	return o.NumOfElem, true
}

// HasNumOfElem returns a boolean if a field has been set.
func (o *RedisKeyEntry) HasNumOfElem() bool {
	return o != nil && o.NumOfElem != nil
}

// SetNumOfElem gets a reference to the given int64 and assigns it to the NumOfElem field.
func (o *RedisKeyEntry) SetNumOfElem(v int64) {
	o.NumOfElem = &v
}

// GetLenOfLargestElem returns the LenOfLargestElem field value if set, zero value otherwise.
func (o *RedisKeyEntry) GetLenOfLargestElem() int64 {
	if o == nil || o.LenOfLargestElem == nil {
		var ret int64
		return ret
	}
	return *o.LenOfLargestElem
}

// GetLenOfLargestElemOk returns a tuple with the LenOfLargestElem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyEntry) GetLenOfLargestElemOk() (*int64, bool) {
	if o == nil || o.LenOfLargestElem == nil {
		return nil, false
	}
	return o.LenOfLargestElem, true
}

// HasLenOfLargestElem returns a boolean if a field has been set.
func (o *RedisKeyEntry) HasLenOfLargestElem() bool {
	return o != nil && o.LenOfLargestElem != nil
}

// SetLenOfLargestElem gets a reference to the given int64 and assigns it to the LenOfLargestElem field.
func (o *RedisKeyEntry) SetLenOfLargestElem(v int64) {
	o.LenOfLargestElem = &v
}

// GetFieldOfLargestElem returns the FieldOfLargestElem field value if set, zero value otherwise.
func (o *RedisKeyEntry) GetFieldOfLargestElem() string {
	if o == nil || o.FieldOfLargestElem == nil {
		var ret string
		return ret
	}
	return *o.FieldOfLargestElem
}

// GetFieldOfLargestElemOk returns a tuple with the FieldOfLargestElem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyEntry) GetFieldOfLargestElemOk() (*string, bool) {
	if o == nil || o.FieldOfLargestElem == nil {
		return nil, false
	}
	return o.FieldOfLargestElem, true
}

// HasFieldOfLargestElem returns a boolean if a field has been set.
func (o *RedisKeyEntry) HasFieldOfLargestElem() bool {
	return o != nil && o.FieldOfLargestElem != nil
}

// SetFieldOfLargestElem gets a reference to the given string and assigns it to the FieldOfLargestElem field.
func (o *RedisKeyEntry) SetFieldOfLargestElem(v string) {
	o.FieldOfLargestElem = &v
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *RedisKeyEntry) GetDb() int32 {
	if o == nil || o.Db == nil {
		var ret int32
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyEntry) GetDbOk() (*int32, bool) {
	if o == nil || o.Db == nil {
		return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *RedisKeyEntry) HasDb() bool {
	return o != nil && o.Db != nil
}

// SetDb gets a reference to the given int32 and assigns it to the Db field.
func (o *RedisKeyEntry) SetDb(v int32) {
	o.Db = &v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *RedisKeyEntry) GetEncoding() string {
	if o == nil || o.Encoding == nil {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyEntry) GetEncodingOk() (*string, bool) {
	if o == nil || o.Encoding == nil {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *RedisKeyEntry) HasEncoding() bool {
	return o != nil && o.Encoding != nil
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *RedisKeyEntry) SetEncoding(v string) {
	o.Encoding = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *RedisKeyEntry) GetExpiration() int64 {
	if o == nil || o.Expiration == nil {
		var ret int64
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyEntry) GetExpirationOk() (*int64, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *RedisKeyEntry) HasExpiration() bool {
	return o != nil && o.Expiration != nil
}

// SetExpiration gets a reference to the given int64 and assigns it to the Expiration field.
func (o *RedisKeyEntry) SetExpiration(v int64) {
	o.Expiration = &v
}

// GetLruIdle returns the LruIdle field value if set, zero value otherwise.
func (o *RedisKeyEntry) GetLruIdle() int64 {
	if o == nil || o.LruIdle == nil {
		var ret int64
		return ret
	}
	return *o.LruIdle
}

// GetLruIdleOk returns a tuple with the LruIdle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyEntry) GetLruIdleOk() (*int64, bool) {
	if o == nil || o.LruIdle == nil {
		return nil, false
	}
	return o.LruIdle, true
}

// HasLruIdle returns a boolean if a field has been set.
func (o *RedisKeyEntry) HasLruIdle() bool {
	return o != nil && o.LruIdle != nil
}

// SetLruIdle gets a reference to the given int64 and assigns it to the LruIdle field.
func (o *RedisKeyEntry) SetLruIdle(v int64) {
	o.LruIdle = &v
}

// GetLfuFreq returns the LfuFreq field value if set, zero value otherwise.
func (o *RedisKeyEntry) GetLfuFreq() int32 {
	if o == nil || o.LfuFreq == nil {
		var ret int32
		return ret
	}
	return *o.LfuFreq
}

// GetLfuFreqOk returns a tuple with the LfuFreq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyEntry) GetLfuFreqOk() (*int32, bool) {
	if o == nil || o.LfuFreq == nil {
		return nil, false
	}
	return o.LfuFreq, true
}

// HasLfuFreq returns a boolean if a field has been set.
func (o *RedisKeyEntry) HasLfuFreq() bool {
	return o != nil && o.LfuFreq != nil
}

// SetLfuFreq gets a reference to the given int32 and assigns it to the LfuFreq field.
func (o *RedisKeyEntry) SetLfuFreq(v int32) {
	o.LfuFreq = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisKeyEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Bytes != nil {
		toSerialize["Bytes"] = o.Bytes
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.NumOfElem != nil {
		toSerialize["NumOfElem"] = o.NumOfElem
	}
	if o.LenOfLargestElem != nil {
		toSerialize["LenOfLargestElem"] = o.LenOfLargestElem
	}
	if o.FieldOfLargestElem != nil {
		toSerialize["FieldOfLargestElem"] = o.FieldOfLargestElem
	}
	if o.Db != nil {
		toSerialize["Db"] = o.Db
	}
	if o.Encoding != nil {
		toSerialize["Encoding"] = o.Encoding
	}
	if o.Expiration != nil {
		toSerialize["Expiration"] = o.Expiration
	}
	if o.LruIdle != nil {
		toSerialize["LruIdle"] = o.LruIdle
	}
	if o.LfuFreq != nil {
		toSerialize["LfuFreq"] = o.LfuFreq
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisKeyEntry) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key                *string `json:"Key,omitempty"`
		Bytes              *int64  `json:"Bytes,omitempty"`
		Type               *string `json:"Type,omitempty"`
		NumOfElem          *int64  `json:"NumOfElem,omitempty"`
		LenOfLargestElem   *int64  `json:"LenOfLargestElem,omitempty"`
		FieldOfLargestElem *string `json:"FieldOfLargestElem,omitempty"`
		Db                 *int32  `json:"Db,omitempty"`
		Encoding           *string `json:"Encoding,omitempty"`
		Expiration         *int64  `json:"Expiration,omitempty"`
		LruIdle            *int64  `json:"LruIdle,omitempty"`
		LfuFreq            *int32  `json:"LfuFreq,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Key", "Bytes", "Type", "NumOfElem", "LenOfLargestElem", "FieldOfLargestElem", "Db", "Encoding", "Expiration", "LruIdle", "LfuFreq"})
	} else {
		return err
	}
	o.Key = all.Key
	o.Bytes = all.Bytes
	o.Type = all.Type
	o.NumOfElem = all.NumOfElem
	o.LenOfLargestElem = all.LenOfLargestElem
	o.FieldOfLargestElem = all.FieldOfLargestElem
	o.Db = all.Db
	o.Encoding = all.Encoding
	o.Expiration = all.Expiration
	o.LruIdle = all.LruIdle
	o.LfuFreq = all.LfuFreq

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
