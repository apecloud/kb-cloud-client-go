// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DmsQueryResponse struct {
	// result set of query
	Data []DmsResult `json:"data,omitempty"`
	// error message set of query
	ErrMessage common.NullableList[string] `json:"errMessage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsQueryResponse instantiates a new DmsQueryResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsQueryResponse() *DmsQueryResponse {
	this := DmsQueryResponse{}
	return &this
}

// NewDmsQueryResponseWithDefaults instantiates a new DmsQueryResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsQueryResponseWithDefaults() *DmsQueryResponse {
	this := DmsQueryResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DmsQueryResponse) GetData() []DmsResult {
	if o == nil || o.Data == nil {
		var ret []DmsResult
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryResponse) GetDataOk() (*[]DmsResult, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DmsQueryResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []DmsResult and assigns it to the Data field.
func (o *DmsQueryResponse) SetData(v []DmsResult) {
	o.Data = v
}

// GetErrMessage returns the ErrMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsQueryResponse) GetErrMessage() []string {
	if o == nil || o.ErrMessage.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ErrMessage.Get()
}

// GetErrMessageOk returns a tuple with the ErrMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsQueryResponse) GetErrMessageOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrMessage.Get(), o.ErrMessage.IsSet()
}

// HasErrMessage returns a boolean if a field has been set.
func (o *DmsQueryResponse) HasErrMessage() bool {
	return o != nil && o.ErrMessage.IsSet()
}

// SetErrMessage gets a reference to the given common.NullableList[string] and assigns it to the ErrMessage field.
func (o *DmsQueryResponse) SetErrMessage(v []string) {
	o.ErrMessage.Set(&v)
}

// SetErrMessageNil sets the value for ErrMessage to be an explicit nil.
func (o *DmsQueryResponse) SetErrMessageNil() {
	o.ErrMessage.Set(nil)
}

// UnsetErrMessage ensures that no value is present for ErrMessage, not even an explicit nil.
func (o *DmsQueryResponse) UnsetErrMessage() {
	o.ErrMessage.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsQueryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.ErrMessage.IsSet() {
		toSerialize["errMessage"] = o.ErrMessage.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsQueryResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data       []DmsResult                 `json:"data,omitempty"`
		ErrMessage common.NullableList[string] `json:"errMessage,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"data", "errMessage"})
	} else {
		return err
	}
	o.Data = all.Data
	o.ErrMessage = all.ErrMessage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
