// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OrgLifecycleRequest struct {
	OrgId             string    `json:"orgID"`
	ExpectedExpiredAt time.Time `json:"expectedExpiredAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgLifecycleRequest instantiates a new OrgLifecycleRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgLifecycleRequest(orgId string, expectedExpiredAt time.Time) *OrgLifecycleRequest {
	this := OrgLifecycleRequest{}
	this.OrgId = orgId
	this.ExpectedExpiredAt = expectedExpiredAt
	return &this
}

// NewOrgLifecycleRequestWithDefaults instantiates a new OrgLifecycleRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgLifecycleRequestWithDefaults() *OrgLifecycleRequest {
	this := OrgLifecycleRequest{}
	return &this
}

// GetOrgId returns the OrgId field value.
func (o *OrgLifecycleRequest) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *OrgLifecycleRequest) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *OrgLifecycleRequest) SetOrgId(v string) {
	o.OrgId = v
}

// GetExpectedExpiredAt returns the ExpectedExpiredAt field value.
func (o *OrgLifecycleRequest) GetExpectedExpiredAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ExpectedExpiredAt
}

// GetExpectedExpiredAtOk returns a tuple with the ExpectedExpiredAt field value
// and a boolean to check if the value has been set.
func (o *OrgLifecycleRequest) GetExpectedExpiredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpectedExpiredAt, true
}

// SetExpectedExpiredAt sets field value.
func (o *OrgLifecycleRequest) SetExpectedExpiredAt(v time.Time) {
	o.ExpectedExpiredAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgLifecycleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["orgID"] = o.OrgId
	if o.ExpectedExpiredAt.Nanosecond() == 0 {
		toSerialize["expectedExpiredAt"] = o.ExpectedExpiredAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["expectedExpiredAt"] = o.ExpectedExpiredAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgLifecycleRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgId             *string    `json:"orgID"`
		ExpectedExpiredAt *time.Time `json:"expectedExpiredAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field orgID missing")
	}
	if all.ExpectedExpiredAt == nil {
		return fmt.Errorf("required field expectedExpiredAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"orgID", "expectedExpiredAt"})
	} else {
		return err
	}
	o.OrgId = *all.OrgId
	o.ExpectedExpiredAt = *all.ExpectedExpiredAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
