// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// NetworkChaosLoss specify the loss in the chaos action
// NODESCRIPTION NetworkChaosLoss
//
// Deprecated: This model is deprecated.
type NetworkChaosLoss struct {
	// specify the loss in the chaos action
	Loss *string `json:"loss,omitempty"`
	// specify the correlation in the chaos action
	Correlation *string `json:"correlation,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNetworkChaosLoss instantiates a new NetworkChaosLoss object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNetworkChaosLoss() *NetworkChaosLoss {
	this := NetworkChaosLoss{}
	return &this
}

// NewNetworkChaosLossWithDefaults instantiates a new NetworkChaosLoss object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNetworkChaosLossWithDefaults() *NetworkChaosLoss {
	this := NetworkChaosLoss{}
	return &this
}

// GetLoss returns the Loss field value if set, zero value otherwise.
func (o *NetworkChaosLoss) GetLoss() string {
	if o == nil || o.Loss == nil {
		var ret string
		return ret
	}
	return *o.Loss
}

// GetLossOk returns a tuple with the Loss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkChaosLoss) GetLossOk() (*string, bool) {
	if o == nil || o.Loss == nil {
		return nil, false
	}
	return o.Loss, true
}

// HasLoss returns a boolean if a field has been set.
func (o *NetworkChaosLoss) HasLoss() bool {
	return o != nil && o.Loss != nil
}

// SetLoss gets a reference to the given string and assigns it to the Loss field.
func (o *NetworkChaosLoss) SetLoss(v string) {
	o.Loss = &v
}

// GetCorrelation returns the Correlation field value if set, zero value otherwise.
func (o *NetworkChaosLoss) GetCorrelation() string {
	if o == nil || o.Correlation == nil {
		var ret string
		return ret
	}
	return *o.Correlation
}

// GetCorrelationOk returns a tuple with the Correlation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkChaosLoss) GetCorrelationOk() (*string, bool) {
	if o == nil || o.Correlation == nil {
		return nil, false
	}
	return o.Correlation, true
}

// HasCorrelation returns a boolean if a field has been set.
func (o *NetworkChaosLoss) HasCorrelation() bool {
	return o != nil && o.Correlation != nil
}

// SetCorrelation gets a reference to the given string and assigns it to the Correlation field.
func (o *NetworkChaosLoss) SetCorrelation(v string) {
	o.Correlation = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NetworkChaosLoss) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Loss != nil {
		toSerialize["loss"] = o.Loss
	}
	if o.Correlation != nil {
		toSerialize["correlation"] = o.Correlation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NetworkChaosLoss) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Loss        *string `json:"loss,omitempty"`
		Correlation *string `json:"correlation,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"loss", "correlation"})
	} else {
		return err
	}
	o.Loss = all.Loss
	o.Correlation = all.Correlation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
