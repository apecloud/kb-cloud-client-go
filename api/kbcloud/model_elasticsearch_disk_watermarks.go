// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchDiskWatermarks struct {
	Low        string `json:"low"`
	High       string `json:"high"`
	FloodStage string `json:"floodStage"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchDiskWatermarks instantiates a new ElasticsearchDiskWatermarks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchDiskWatermarks(low string, high string, floodStage string) *ElasticsearchDiskWatermarks {
	this := ElasticsearchDiskWatermarks{}
	this.Low = low
	this.High = high
	this.FloodStage = floodStage
	return &this
}

// NewElasticsearchDiskWatermarksWithDefaults instantiates a new ElasticsearchDiskWatermarks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchDiskWatermarksWithDefaults() *ElasticsearchDiskWatermarks {
	this := ElasticsearchDiskWatermarks{}
	return &this
}

// GetLow returns the Low field value.
func (o *ElasticsearchDiskWatermarks) GetLow() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Low
}

// GetLowOk returns a tuple with the Low field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchDiskWatermarks) GetLowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Low, true
}

// SetLow sets field value.
func (o *ElasticsearchDiskWatermarks) SetLow(v string) {
	o.Low = v
}

// GetHigh returns the High field value.
func (o *ElasticsearchDiskWatermarks) GetHigh() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.High
}

// GetHighOk returns a tuple with the High field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchDiskWatermarks) GetHighOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.High, true
}

// SetHigh sets field value.
func (o *ElasticsearchDiskWatermarks) SetHigh(v string) {
	o.High = v
}

// GetFloodStage returns the FloodStage field value.
func (o *ElasticsearchDiskWatermarks) GetFloodStage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FloodStage
}

// GetFloodStageOk returns a tuple with the FloodStage field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchDiskWatermarks) GetFloodStageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FloodStage, true
}

// SetFloodStage sets field value.
func (o *ElasticsearchDiskWatermarks) SetFloodStage(v string) {
	o.FloodStage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchDiskWatermarks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["low"] = o.Low
	toSerialize["high"] = o.High
	toSerialize["floodStage"] = o.FloodStage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchDiskWatermarks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Low        *string `json:"low"`
		High       *string `json:"high"`
		FloodStage *string `json:"floodStage"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Low == nil {
		return fmt.Errorf("required field low missing")
	}
	if all.High == nil {
		return fmt.Errorf("required field high missing")
	}
	if all.FloodStage == nil {
		return fmt.Errorf("required field floodStage missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"low", "high", "floodStage"})
	} else {
		return err
	}
	o.Low = *all.Low
	o.High = *all.High
	o.FloodStage = *all.FloodStage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
