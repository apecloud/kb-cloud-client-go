// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DmsServiceAccountUpdateRequest MinIO service account access-key update request
type DmsServiceAccountUpdateRequest struct {
	// MinIO service account status.
	Status *DmsServiceAccountStatus `json:"status,omitempty"`
	// Optional service account expiration time. Use the Unix epoch to clear expiration.
	Expiration *time.Time `json:"expiration,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsServiceAccountUpdateRequest instantiates a new DmsServiceAccountUpdateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsServiceAccountUpdateRequest() *DmsServiceAccountUpdateRequest {
	this := DmsServiceAccountUpdateRequest{}
	return &this
}

// NewDmsServiceAccountUpdateRequestWithDefaults instantiates a new DmsServiceAccountUpdateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsServiceAccountUpdateRequestWithDefaults() *DmsServiceAccountUpdateRequest {
	this := DmsServiceAccountUpdateRequest{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DmsServiceAccountUpdateRequest) GetStatus() DmsServiceAccountStatus {
	if o == nil || o.Status == nil {
		var ret DmsServiceAccountStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsServiceAccountUpdateRequest) GetStatusOk() (*DmsServiceAccountStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DmsServiceAccountUpdateRequest) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given DmsServiceAccountStatus and assigns it to the Status field.
func (o *DmsServiceAccountUpdateRequest) SetStatus(v DmsServiceAccountStatus) {
	o.Status = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *DmsServiceAccountUpdateRequest) GetExpiration() time.Time {
	if o == nil || o.Expiration == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsServiceAccountUpdateRequest) GetExpirationOk() (*time.Time, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *DmsServiceAccountUpdateRequest) HasExpiration() bool {
	return o != nil && o.Expiration != nil
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *DmsServiceAccountUpdateRequest) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsServiceAccountUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Expiration != nil {
		if o.Expiration.Nanosecond() == 0 {
			toSerialize["expiration"] = o.Expiration.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["expiration"] = o.Expiration.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsServiceAccountUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status     *DmsServiceAccountStatus `json:"status,omitempty"`
		Expiration *time.Time               `json:"expiration,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "expiration"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Expiration = all.Expiration

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
