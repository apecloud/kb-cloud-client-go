// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataEndpointResponse struct {
	Status       *EndpointStatusEnum `json:"status,omitempty"`
	EndpointId   *string             `json:"endpointID,omitempty"`
	EndpointName *string             `json:"endpointName,omitempty"`
	CreatedAt    *time.Time          `json:"createdAt,omitempty"`
	UpdatedAt    common.NullableTime `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataEndpointResponse instantiates a new DataEndpointResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataEndpointResponse() *DataEndpointResponse {
	this := DataEndpointResponse{}
	return &this
}

// NewDataEndpointResponseWithDefaults instantiates a new DataEndpointResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataEndpointResponseWithDefaults() *DataEndpointResponse {
	this := DataEndpointResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DataEndpointResponse) GetStatus() EndpointStatusEnum {
	if o == nil || o.Status == nil {
		var ret EndpointStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataEndpointResponse) GetStatusOk() (*EndpointStatusEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DataEndpointResponse) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given EndpointStatusEnum and assigns it to the Status field.
func (o *DataEndpointResponse) SetStatus(v EndpointStatusEnum) {
	o.Status = &v
}

// GetEndpointId returns the EndpointId field value if set, zero value otherwise.
func (o *DataEndpointResponse) GetEndpointId() string {
	if o == nil || o.EndpointId == nil {
		var ret string
		return ret
	}
	return *o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataEndpointResponse) GetEndpointIdOk() (*string, bool) {
	if o == nil || o.EndpointId == nil {
		return nil, false
	}
	return o.EndpointId, true
}

// HasEndpointId returns a boolean if a field has been set.
func (o *DataEndpointResponse) HasEndpointId() bool {
	return o != nil && o.EndpointId != nil
}

// SetEndpointId gets a reference to the given string and assigns it to the EndpointId field.
func (o *DataEndpointResponse) SetEndpointId(v string) {
	o.EndpointId = &v
}

// GetEndpointName returns the EndpointName field value if set, zero value otherwise.
func (o *DataEndpointResponse) GetEndpointName() string {
	if o == nil || o.EndpointName == nil {
		var ret string
		return ret
	}
	return *o.EndpointName
}

// GetEndpointNameOk returns a tuple with the EndpointName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataEndpointResponse) GetEndpointNameOk() (*string, bool) {
	if o == nil || o.EndpointName == nil {
		return nil, false
	}
	return o.EndpointName, true
}

// HasEndpointName returns a boolean if a field has been set.
func (o *DataEndpointResponse) HasEndpointName() bool {
	return o != nil && o.EndpointName != nil
}

// SetEndpointName gets a reference to the given string and assigns it to the EndpointName field.
func (o *DataEndpointResponse) SetEndpointName(v string) {
	o.EndpointName = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DataEndpointResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataEndpointResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DataEndpointResponse) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DataEndpointResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataEndpointResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataEndpointResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DataEndpointResponse) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt.IsSet()
}

// SetUpdatedAt gets a reference to the given common.NullableTime and assigns it to the UpdatedAt field.
func (o *DataEndpointResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil.
func (o *DataEndpointResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil.
func (o *DataEndpointResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DataEndpointResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.EndpointId != nil {
		toSerialize["endpointID"] = o.EndpointId
	}
	if o.EndpointName != nil {
		toSerialize["endpointName"] = o.EndpointName
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updatedAt"] = o.UpdatedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataEndpointResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status       *EndpointStatusEnum `json:"status,omitempty"`
		EndpointId   *string             `json:"endpointID,omitempty"`
		EndpointName *string             `json:"endpointName,omitempty"`
		CreatedAt    *time.Time          `json:"createdAt,omitempty"`
		UpdatedAt    common.NullableTime `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "endpointID", "endpointName", "createdAt", "updatedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.EndpointId = all.EndpointId
	o.EndpointName = all.EndpointName
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
