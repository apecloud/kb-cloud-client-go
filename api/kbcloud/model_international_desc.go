// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type InternationalDesc struct {
	ZhCn *string `json:"zh-CN,omitempty"`
	EnUs *string `json:"en-US,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInternationalDesc instantiates a new InternationalDesc object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInternationalDesc() *InternationalDesc {
	this := InternationalDesc{}
	return &this
}

// NewInternationalDescWithDefaults instantiates a new InternationalDesc object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInternationalDescWithDefaults() *InternationalDesc {
	this := InternationalDesc{}
	return &this
}

// GetZhCn returns the ZhCn field value if set, zero value otherwise.
func (o *InternationalDesc) GetZhCn() string {
	if o == nil || o.ZhCn == nil {
		var ret string
		return ret
	}
	return *o.ZhCn
}

// GetZhCnOk returns a tuple with the ZhCn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternationalDesc) GetZhCnOk() (*string, bool) {
	if o == nil || o.ZhCn == nil {
		return nil, false
	}
	return o.ZhCn, true
}

// HasZhCn returns a boolean if a field has been set.
func (o *InternationalDesc) HasZhCn() bool {
	return o != nil && o.ZhCn != nil
}

// SetZhCn gets a reference to the given string and assigns it to the ZhCn field.
func (o *InternationalDesc) SetZhCn(v string) {
	o.ZhCn = &v
}

// GetEnUs returns the EnUs field value if set, zero value otherwise.
func (o *InternationalDesc) GetEnUs() string {
	if o == nil || o.EnUs == nil {
		var ret string
		return ret
	}
	return *o.EnUs
}

// GetEnUsOk returns a tuple with the EnUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternationalDesc) GetEnUsOk() (*string, bool) {
	if o == nil || o.EnUs == nil {
		return nil, false
	}
	return o.EnUs, true
}

// HasEnUs returns a boolean if a field has been set.
func (o *InternationalDesc) HasEnUs() bool {
	return o != nil && o.EnUs != nil
}

// SetEnUs gets a reference to the given string and assigns it to the EnUs field.
func (o *InternationalDesc) SetEnUs(v string) {
	o.EnUs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InternationalDesc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ZhCn != nil {
		toSerialize["zh-CN"] = o.ZhCn
	}
	if o.EnUs != nil {
		toSerialize["en-US"] = o.EnUs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InternationalDesc) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ZhCn *string `json:"zh-CN,omitempty"`
		EnUs *string `json:"en-US,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"zh-CN", "en-US"})
	} else {
		return err
	}
	o.ZhCn = all.ZhCn
	o.EnUs = all.EnUs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
