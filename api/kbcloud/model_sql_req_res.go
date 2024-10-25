// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// NODESCRIPTION SqlReqRes
type SqlReqRes struct {
	// SQL request or response body
	Body *string `json:"body,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSqlReqRes instantiates a new SqlReqRes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSqlReqRes() *SqlReqRes {
	this := SqlReqRes{}
	return &this
}

// NewSqlReqResWithDefaults instantiates a new SqlReqRes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSqlReqResWithDefaults() *SqlReqRes {
	this := SqlReqRes{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *SqlReqRes) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlReqRes) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *SqlReqRes) HasBody() bool {
	return o != nil && o.Body != nil
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *SqlReqRes) SetBody(v string) {
	o.Body = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SqlReqRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SqlReqRes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Body *string `json:"body,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"body"})
	} else {
		return err
	}
	o.Body = all.Body

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
