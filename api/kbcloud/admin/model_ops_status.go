// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type OpsStatus struct {
	Status  interface{} `json:"Status,omitempty"`
	Message *string     `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsStatus instantiates a new OpsStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsStatus() *OpsStatus {
	this := OpsStatus{}
	return &this
}

// NewOpsStatusWithDefaults instantiates a new OpsStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsStatusWithDefaults() *OpsStatus {
	this := OpsStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OpsStatus) GetStatus() interface{} {
	if o == nil || o.Status == nil {
		var ret interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsStatus) GetStatusOk() (*interface{}, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OpsStatus) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given interface{} and assigns it to the Status field.
func (o *OpsStatus) SetStatus(v interface{}) {
	o.Status = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OpsStatus) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsStatus) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OpsStatus) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OpsStatus) SetMessage(v string) {
	o.Message = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status  interface{} `json:"Status,omitempty"`
		Message *string     `json:"message,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"Status", "message"})
	} else {
		return err
	}
	o.Status = all.Status
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
