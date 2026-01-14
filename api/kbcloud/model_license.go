// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// License License info
type License struct {
	// The kubernetes cluster ID
	ClusterId string `json:"clusterID"`
	// The license expiration time
	NotAfter time.Time `json:"notAfter"`
	// The license start time
	NotBefore time.Time `json:"notBefore"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLicense instantiates a new License object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLicense(clusterId string, notAfter time.Time, notBefore time.Time) *License {
	this := License{}
	this.ClusterId = clusterId
	this.NotAfter = notAfter
	this.NotBefore = notBefore
	return &this
}

// NewLicenseWithDefaults instantiates a new License object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLicenseWithDefaults() *License {
	this := License{}
	return &this
}

// GetClusterId returns the ClusterId field value.
func (o *License) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *License) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value.
func (o *License) SetClusterId(v string) {
	o.ClusterId = v
}

// GetNotAfter returns the NotAfter field value.
func (o *License) GetNotAfter() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value
// and a boolean to check if the value has been set.
func (o *License) GetNotAfterOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotAfter, true
}

// SetNotAfter sets field value.
func (o *License) SetNotAfter(v time.Time) {
	o.NotAfter = v
}

// GetNotBefore returns the NotBefore field value.
func (o *License) GetNotBefore() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value
// and a boolean to check if the value has been set.
func (o *License) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotBefore, true
}

// SetNotBefore sets field value.
func (o *License) SetNotBefore(v time.Time) {
	o.NotBefore = v
}

// MarshalJSON serializes the struct using spec logic.
func (o License) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["clusterID"] = o.ClusterId
	if o.NotAfter.Nanosecond() == 0 {
		toSerialize["notAfter"] = o.NotAfter.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["notAfter"] = o.NotAfter.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.NotBefore.Nanosecond() == 0 {
		toSerialize["notBefore"] = o.NotBefore.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["notBefore"] = o.NotBefore.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *License) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusterId *string    `json:"clusterID"`
		NotAfter  *time.Time `json:"notAfter"`
		NotBefore *time.Time `json:"notBefore"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ClusterId == nil {
		return fmt.Errorf("required field clusterID missing")
	}
	if all.NotAfter == nil {
		return fmt.Errorf("required field notAfter missing")
	}
	if all.NotBefore == nil {
		return fmt.Errorf("required field notBefore missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clusterID", "notAfter", "notBefore"})
	} else {
		return err
	}
	o.ClusterId = *all.ClusterId
	o.NotAfter = *all.NotAfter
	o.NotBefore = *all.NotBefore

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
