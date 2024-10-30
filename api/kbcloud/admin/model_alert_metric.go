// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// AlertMetric Alert metric information 
type AlertMetric struct {
	Key string `json:"key"`
	Threshold int32 `json:"threshold"`
	Notation string `json:"notation"`
	Category *string `json:"category,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewAlertMetric instantiates a new AlertMetric object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertMetric(key string, threshold int32, notation string) *AlertMetric {
	this := AlertMetric{}
	this.Key = key
	this.Threshold = threshold
	this.Notation = notation
	return &this
}

// NewAlertMetricWithDefaults instantiates a new AlertMetric object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertMetricWithDefaults() *AlertMetric {
	this := AlertMetric{}
	return &this
}
// GetKey returns the Key field value.
func (o *AlertMetric) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *AlertMetric) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *AlertMetric) SetKey(v string) {
	o.Key = v
}


// GetThreshold returns the Threshold field value.
func (o *AlertMetric) GetThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *AlertMetric) GetThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value.
func (o *AlertMetric) SetThreshold(v int32) {
	o.Threshold = v
}


// GetNotation returns the Notation field value.
func (o *AlertMetric) GetNotation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Notation
}

// GetNotationOk returns a tuple with the Notation field value
// and a boolean to check if the value has been set.
func (o *AlertMetric) GetNotationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notation, true
}

// SetNotation sets field value.
func (o *AlertMetric) SetNotation(v string) {
	o.Notation = v
}


// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AlertMetric) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertMetric) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AlertMetric) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *AlertMetric) SetCategory(v string) {
	o.Category = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o AlertMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["key"] = o.Key
	toSerialize["threshold"] = o.Threshold
	toSerialize["notation"] = o.Notation
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertMetric) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key *string `json:"key"`
		Threshold *int32 `json:"threshold"`
		Notation *string `json:"notation"`
		Category *string `json:"category,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Threshold == nil {
		return fmt.Errorf("required field threshold missing")
	}
	if all.Notation == nil {
		return fmt.Errorf("required field notation missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "key", "threshold", "notation", "category",  })
	} else {
		return err
	}
	o.Key = *all.Key
	o.Threshold = *all.Threshold
	o.Notation = *all.Notation
	o.Category = all.Category

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
