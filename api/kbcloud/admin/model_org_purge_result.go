// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OrgPurgeResult struct {
	OrgId     *string    `json:"orgID,omitempty"`
	Status    *string    `json:"status,omitempty"`
	ExpiredAt *time.Time `json:"expiredAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgPurgeResult instantiates a new OrgPurgeResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgPurgeResult() *OrgPurgeResult {
	this := OrgPurgeResult{}
	return &this
}

// NewOrgPurgeResultWithDefaults instantiates a new OrgPurgeResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgPurgeResultWithDefaults() *OrgPurgeResult {
	this := OrgPurgeResult{}
	return &this
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *OrgPurgeResult) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgPurgeResult) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *OrgPurgeResult) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *OrgPurgeResult) SetOrgId(v string) {
	o.OrgId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrgPurgeResult) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgPurgeResult) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrgPurgeResult) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrgPurgeResult) SetStatus(v string) {
	o.Status = &v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise.
func (o *OrgPurgeResult) GetExpiredAt() time.Time {
	if o == nil || o.ExpiredAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgPurgeResult) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiredAt == nil {
		return nil, false
	}
	return o.ExpiredAt, true
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *OrgPurgeResult) HasExpiredAt() bool {
	return o != nil && o.ExpiredAt != nil
}

// SetExpiredAt gets a reference to the given time.Time and assigns it to the ExpiredAt field.
func (o *OrgPurgeResult) SetExpiredAt(v time.Time) {
	o.ExpiredAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgPurgeResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.OrgId != nil {
		toSerialize["orgID"] = o.OrgId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ExpiredAt != nil {
		if o.ExpiredAt.Nanosecond() == 0 {
			toSerialize["expiredAt"] = o.ExpiredAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["expiredAt"] = o.ExpiredAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgPurgeResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgId     *string    `json:"orgID,omitempty"`
		Status    *string    `json:"status,omitempty"`
		ExpiredAt *time.Time `json:"expiredAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"orgID", "status", "expiredAt"})
	} else {
		return err
	}
	o.OrgId = all.OrgId
	o.Status = all.Status
	o.ExpiredAt = all.ExpiredAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
