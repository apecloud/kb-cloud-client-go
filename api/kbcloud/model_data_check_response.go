// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataCheckResponse struct {
	CheckId   *string    `json:"checkID,omitempty"`
	CheckName *string    `json:"checkName,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataCheckResponse instantiates a new DataCheckResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataCheckResponse() *DataCheckResponse {
	this := DataCheckResponse{}
	return &this
}

// NewDataCheckResponseWithDefaults instantiates a new DataCheckResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataCheckResponseWithDefaults() *DataCheckResponse {
	this := DataCheckResponse{}
	return &this
}

// GetCheckId returns the CheckId field value if set, zero value otherwise.
func (o *DataCheckResponse) GetCheckId() string {
	if o == nil || o.CheckId == nil {
		var ret string
		return ret
	}
	return *o.CheckId
}

// GetCheckIdOk returns a tuple with the CheckId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckResponse) GetCheckIdOk() (*string, bool) {
	if o == nil || o.CheckId == nil {
		return nil, false
	}
	return o.CheckId, true
}

// HasCheckId returns a boolean if a field has been set.
func (o *DataCheckResponse) HasCheckId() bool {
	return o != nil && o.CheckId != nil
}

// SetCheckId gets a reference to the given string and assigns it to the CheckId field.
func (o *DataCheckResponse) SetCheckId(v string) {
	o.CheckId = &v
}

// GetCheckName returns the CheckName field value if set, zero value otherwise.
func (o *DataCheckResponse) GetCheckName() string {
	if o == nil || o.CheckName == nil {
		var ret string
		return ret
	}
	return *o.CheckName
}

// GetCheckNameOk returns a tuple with the CheckName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckResponse) GetCheckNameOk() (*string, bool) {
	if o == nil || o.CheckName == nil {
		return nil, false
	}
	return o.CheckName, true
}

// HasCheckName returns a boolean if a field has been set.
func (o *DataCheckResponse) HasCheckName() bool {
	return o != nil && o.CheckName != nil
}

// SetCheckName gets a reference to the given string and assigns it to the CheckName field.
func (o *DataCheckResponse) SetCheckName(v string) {
	o.CheckName = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DataCheckResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DataCheckResponse) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DataCheckResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CheckId != nil {
		toSerialize["checkID"] = o.CheckId
	}
	if o.CheckName != nil {
		toSerialize["checkName"] = o.CheckName
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataCheckResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CheckId   *string    `json:"checkID,omitempty"`
		CheckName *string    `json:"checkName,omitempty"`
		CreatedAt *time.Time `json:"createdAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"checkID", "checkName", "createdAt"})
	} else {
		return err
	}
	o.CheckId = all.CheckId
	o.CheckName = all.CheckName
	o.CreatedAt = all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
