// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// BasicBill Basic bill information
type BasicBill struct {
	// The name of aggregationGroupType
	AggregationGroupName *string `json:"aggregationGroupName,omitempty"`
	// The timestamp of data
	DataTime *int64 `json:"dataTime,omitempty"`
	// The total price
	Price *string `json:"price,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBasicBill instantiates a new BasicBill object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBasicBill() *BasicBill {
	this := BasicBill{}
	return &this
}

// NewBasicBillWithDefaults instantiates a new BasicBill object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBasicBillWithDefaults() *BasicBill {
	this := BasicBill{}
	return &this
}

// GetAggregationGroupName returns the AggregationGroupName field value if set, zero value otherwise.
func (o *BasicBill) GetAggregationGroupName() string {
	if o == nil || o.AggregationGroupName == nil {
		var ret string
		return ret
	}
	return *o.AggregationGroupName
}

// GetAggregationGroupNameOk returns a tuple with the AggregationGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicBill) GetAggregationGroupNameOk() (*string, bool) {
	if o == nil || o.AggregationGroupName == nil {
		return nil, false
	}
	return o.AggregationGroupName, true
}

// HasAggregationGroupName returns a boolean if a field has been set.
func (o *BasicBill) HasAggregationGroupName() bool {
	return o != nil && o.AggregationGroupName != nil
}

// SetAggregationGroupName gets a reference to the given string and assigns it to the AggregationGroupName field.
func (o *BasicBill) SetAggregationGroupName(v string) {
	o.AggregationGroupName = &v
}

// GetDataTime returns the DataTime field value if set, zero value otherwise.
func (o *BasicBill) GetDataTime() int64 {
	if o == nil || o.DataTime == nil {
		var ret int64
		return ret
	}
	return *o.DataTime
}

// GetDataTimeOk returns a tuple with the DataTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicBill) GetDataTimeOk() (*int64, bool) {
	if o == nil || o.DataTime == nil {
		return nil, false
	}
	return o.DataTime, true
}

// HasDataTime returns a boolean if a field has been set.
func (o *BasicBill) HasDataTime() bool {
	return o != nil && o.DataTime != nil
}

// SetDataTime gets a reference to the given int64 and assigns it to the DataTime field.
func (o *BasicBill) SetDataTime(v int64) {
	o.DataTime = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BasicBill) GetPrice() string {
	if o == nil || o.Price == nil {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicBill) GetPriceOk() (*string, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BasicBill) HasPrice() bool {
	return o != nil && o.Price != nil
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *BasicBill) SetPrice(v string) {
	o.Price = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BasicBill) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.AggregationGroupName != nil {
		toSerialize["aggregationGroupName"] = o.AggregationGroupName
	}
	if o.DataTime != nil {
		toSerialize["dataTime"] = o.DataTime
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BasicBill) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AggregationGroupName *string `json:"aggregationGroupName,omitempty"`
		DataTime             *int64  `json:"dataTime,omitempty"`
		Price                *string `json:"price,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"aggregationGroupName", "dataTime", "price"})
	} else {
		return err
	}
	o.AggregationGroupName = all.AggregationGroupName
	o.DataTime = all.DataTime
	o.Price = all.Price

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
