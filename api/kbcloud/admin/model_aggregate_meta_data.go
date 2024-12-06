// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// AggregateMetaData A series of data points for meta data, including a count and a timestamp.
type AggregateMetaData struct {
	// The total count of the metaData items.
	Total int32 `json:"total"`
	// The list of series items aggregated in the time range.
	Items []SeriesItem `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAggregateMetaData instantiates a new AggregateMetaData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregateMetaData(total int32, items []SeriesItem) *AggregateMetaData {
	this := AggregateMetaData{}
	this.Total = total
	this.Items = items
	return &this
}

// NewAggregateMetaDataWithDefaults instantiates a new AggregateMetaData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregateMetaDataWithDefaults() *AggregateMetaData {
	this := AggregateMetaData{}
	return &this
}

// GetTotal returns the Total field value.
func (o *AggregateMetaData) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *AggregateMetaData) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *AggregateMetaData) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value.
func (o *AggregateMetaData) GetItems() []SeriesItem {
	if o == nil {
		var ret []SeriesItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *AggregateMetaData) GetItemsOk() (*[]SeriesItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *AggregateMetaData) SetItems(v []SeriesItem) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregateMetaData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["total"] = o.Total
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregateMetaData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Total *int32        `json:"total"`
		Items *[]SeriesItem `json:"items"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"total", "items"})
	} else {
		return err
	}
	o.Total = *all.Total
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
