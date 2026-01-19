// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BlueGreenDeploymentStatusResponse blueGreenDeploymentStatusResponse represents the response containing the blue-green deployment status.
type BlueGreenDeploymentStatusResponse struct {
	// The ID of the blue-green deployment.
	DeploymentId string `json:"deploymentID"`
	// The status of a blue-green deployment.
	DeploymentStatus *BlueGreenDeploymentStatus `json:"deploymentStatus,omitempty"`
	// The lag of the event.
	EventLag *string `json:"eventLag,omitempty"`
	// The lag of the timestamp.
	TimestampLag common.NullableString `json:"timestampLag,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBlueGreenDeploymentStatusResponse instantiates a new BlueGreenDeploymentStatusResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBlueGreenDeploymentStatusResponse(deploymentId string) *BlueGreenDeploymentStatusResponse {
	this := BlueGreenDeploymentStatusResponse{}
	this.DeploymentId = deploymentId
	return &this
}

// NewBlueGreenDeploymentStatusResponseWithDefaults instantiates a new BlueGreenDeploymentStatusResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBlueGreenDeploymentStatusResponseWithDefaults() *BlueGreenDeploymentStatusResponse {
	this := BlueGreenDeploymentStatusResponse{}
	return &this
}

// GetDeploymentId returns the DeploymentId field value.
func (o *BlueGreenDeploymentStatusResponse) GetDeploymentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentStatusResponse) GetDeploymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentId, true
}

// SetDeploymentId sets field value.
func (o *BlueGreenDeploymentStatusResponse) SetDeploymentId(v string) {
	o.DeploymentId = v
}

// GetDeploymentStatus returns the DeploymentStatus field value if set, zero value otherwise.
func (o *BlueGreenDeploymentStatusResponse) GetDeploymentStatus() BlueGreenDeploymentStatus {
	if o == nil || o.DeploymentStatus == nil {
		var ret BlueGreenDeploymentStatus
		return ret
	}
	return *o.DeploymentStatus
}

// GetDeploymentStatusOk returns a tuple with the DeploymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentStatusResponse) GetDeploymentStatusOk() (*BlueGreenDeploymentStatus, bool) {
	if o == nil || o.DeploymentStatus == nil {
		return nil, false
	}
	return o.DeploymentStatus, true
}

// HasDeploymentStatus returns a boolean if a field has been set.
func (o *BlueGreenDeploymentStatusResponse) HasDeploymentStatus() bool {
	return o != nil && o.DeploymentStatus != nil
}

// SetDeploymentStatus gets a reference to the given BlueGreenDeploymentStatus and assigns it to the DeploymentStatus field.
func (o *BlueGreenDeploymentStatusResponse) SetDeploymentStatus(v BlueGreenDeploymentStatus) {
	o.DeploymentStatus = &v
}

// GetEventLag returns the EventLag field value if set, zero value otherwise.
func (o *BlueGreenDeploymentStatusResponse) GetEventLag() string {
	if o == nil || o.EventLag == nil {
		var ret string
		return ret
	}
	return *o.EventLag
}

// GetEventLagOk returns a tuple with the EventLag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentStatusResponse) GetEventLagOk() (*string, bool) {
	if o == nil || o.EventLag == nil {
		return nil, false
	}
	return o.EventLag, true
}

// HasEventLag returns a boolean if a field has been set.
func (o *BlueGreenDeploymentStatusResponse) HasEventLag() bool {
	return o != nil && o.EventLag != nil
}

// SetEventLag gets a reference to the given string and assigns it to the EventLag field.
func (o *BlueGreenDeploymentStatusResponse) SetEventLag(v string) {
	o.EventLag = &v
}

// GetTimestampLag returns the TimestampLag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlueGreenDeploymentStatusResponse) GetTimestampLag() string {
	if o == nil || o.TimestampLag.Get() == nil {
		var ret string
		return ret
	}
	return *o.TimestampLag.Get()
}

// GetTimestampLagOk returns a tuple with the TimestampLag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BlueGreenDeploymentStatusResponse) GetTimestampLagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimestampLag.Get(), o.TimestampLag.IsSet()
}

// HasTimestampLag returns a boolean if a field has been set.
func (o *BlueGreenDeploymentStatusResponse) HasTimestampLag() bool {
	return o != nil && o.TimestampLag.IsSet()
}

// SetTimestampLag gets a reference to the given common.NullableString and assigns it to the TimestampLag field.
func (o *BlueGreenDeploymentStatusResponse) SetTimestampLag(v string) {
	o.TimestampLag.Set(&v)
}

// SetTimestampLagNil sets the value for TimestampLag to be an explicit nil.
func (o *BlueGreenDeploymentStatusResponse) SetTimestampLagNil() {
	o.TimestampLag.Set(nil)
}

// UnsetTimestampLag ensures that no value is present for TimestampLag, not even an explicit nil.
func (o *BlueGreenDeploymentStatusResponse) UnsetTimestampLag() {
	o.TimestampLag.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o BlueGreenDeploymentStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["deploymentID"] = o.DeploymentId
	if o.DeploymentStatus != nil {
		toSerialize["deploymentStatus"] = o.DeploymentStatus
	}
	if o.EventLag != nil {
		toSerialize["eventLag"] = o.EventLag
	}
	if o.TimestampLag.IsSet() {
		toSerialize["timestampLag"] = o.TimestampLag.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BlueGreenDeploymentStatusResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeploymentId     *string                    `json:"deploymentID"`
		DeploymentStatus *BlueGreenDeploymentStatus `json:"deploymentStatus,omitempty"`
		EventLag         *string                    `json:"eventLag,omitempty"`
		TimestampLag     common.NullableString      `json:"timestampLag,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.DeploymentId == nil {
		return fmt.Errorf("required field deploymentID missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"deploymentID", "deploymentStatus", "eventLag", "timestampLag"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DeploymentId = *all.DeploymentId
	if all.DeploymentStatus != nil && !all.DeploymentStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.DeploymentStatus = all.DeploymentStatus
	}
	o.EventLag = all.EventLag
	o.TimestampLag = all.TimestampLag

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
