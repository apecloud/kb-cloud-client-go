// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBNodeLicense struct {
	Name        string                `json:"name"`
	Mode        int32                 `json:"mode"`
	Observation DolphinDBHealthCheck  `json:"observation"`
	Effective   *DolphinDBLicenseInfo `json:"effective,omitempty"`
	// UTC calendar date difference; negative means expired and zero means today.
	DaysUntilExpiration *int64 `json:"daysUntilExpiration,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBNodeLicense instantiates a new DolphinDBNodeLicense object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBNodeLicense(name string, mode int32, observation DolphinDBHealthCheck) *DolphinDBNodeLicense {
	this := DolphinDBNodeLicense{}
	this.Name = name
	this.Mode = mode
	this.Observation = observation
	return &this
}

// NewDolphinDBNodeLicenseWithDefaults instantiates a new DolphinDBNodeLicense object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBNodeLicenseWithDefaults() *DolphinDBNodeLicense {
	this := DolphinDBNodeLicense{}
	return &this
}

// GetName returns the Name field value.
func (o *DolphinDBNodeLicense) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DolphinDBNodeLicense) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DolphinDBNodeLicense) SetName(v string) {
	o.Name = v
}

// GetMode returns the Mode field value.
func (o *DolphinDBNodeLicense) GetMode() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *DolphinDBNodeLicense) GetModeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *DolphinDBNodeLicense) SetMode(v int32) {
	o.Mode = v
}

// GetObservation returns the Observation field value.
func (o *DolphinDBNodeLicense) GetObservation() DolphinDBHealthCheck {
	if o == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return o.Observation
}

// GetObservationOk returns a tuple with the Observation field value
// and a boolean to check if the value has been set.
func (o *DolphinDBNodeLicense) GetObservationOk() (*DolphinDBHealthCheck, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Observation, true
}

// SetObservation sets field value.
func (o *DolphinDBNodeLicense) SetObservation(v DolphinDBHealthCheck) {
	o.Observation = v
}

// GetEffective returns the Effective field value if set, zero value otherwise.
func (o *DolphinDBNodeLicense) GetEffective() DolphinDBLicenseInfo {
	if o == nil || o.Effective == nil {
		var ret DolphinDBLicenseInfo
		return ret
	}
	return *o.Effective
}

// GetEffectiveOk returns a tuple with the Effective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBNodeLicense) GetEffectiveOk() (*DolphinDBLicenseInfo, bool) {
	if o == nil || o.Effective == nil {
		return nil, false
	}
	return o.Effective, true
}

// HasEffective returns a boolean if a field has been set.
func (o *DolphinDBNodeLicense) HasEffective() bool {
	return o != nil && o.Effective != nil
}

// SetEffective gets a reference to the given DolphinDBLicenseInfo and assigns it to the Effective field.
func (o *DolphinDBNodeLicense) SetEffective(v DolphinDBLicenseInfo) {
	o.Effective = &v
}

// GetDaysUntilExpiration returns the DaysUntilExpiration field value if set, zero value otherwise.
func (o *DolphinDBNodeLicense) GetDaysUntilExpiration() int64 {
	if o == nil || o.DaysUntilExpiration == nil {
		var ret int64
		return ret
	}
	return *o.DaysUntilExpiration
}

// GetDaysUntilExpirationOk returns a tuple with the DaysUntilExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBNodeLicense) GetDaysUntilExpirationOk() (*int64, bool) {
	if o == nil || o.DaysUntilExpiration == nil {
		return nil, false
	}
	return o.DaysUntilExpiration, true
}

// HasDaysUntilExpiration returns a boolean if a field has been set.
func (o *DolphinDBNodeLicense) HasDaysUntilExpiration() bool {
	return o != nil && o.DaysUntilExpiration != nil
}

// SetDaysUntilExpiration gets a reference to the given int64 and assigns it to the DaysUntilExpiration field.
func (o *DolphinDBNodeLicense) SetDaysUntilExpiration(v int64) {
	o.DaysUntilExpiration = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBNodeLicense) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["mode"] = o.Mode
	toSerialize["observation"] = o.Observation
	if o.Effective != nil {
		toSerialize["effective"] = o.Effective
	}
	if o.DaysUntilExpiration != nil {
		toSerialize["daysUntilExpiration"] = o.DaysUntilExpiration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBNodeLicense) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                *string               `json:"name"`
		Mode                *int32                `json:"mode"`
		Observation         *DolphinDBHealthCheck `json:"observation"`
		Effective           *DolphinDBLicenseInfo `json:"effective,omitempty"`
		DaysUntilExpiration *int64                `json:"daysUntilExpiration,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.Observation == nil {
		return fmt.Errorf("required field observation missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "mode", "observation", "effective", "daysUntilExpiration"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.Mode = *all.Mode
	if all.Observation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Observation = *all.Observation
	if all.Effective != nil && all.Effective.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Effective = all.Effective
	o.DaysUntilExpiration = all.DaysUntilExpiration

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
