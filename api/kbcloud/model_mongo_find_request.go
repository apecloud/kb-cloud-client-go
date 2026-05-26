// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoFindRequest struct {
	// MongoDB query filter
	Filter interface{} `json:"filter,omitempty"`
	// sort specification
	Sort interface{} `json:"sort,omitempty"`
	// field projection
	Projection interface{} `json:"projection,omitempty"`
	// MongoDB collation options
	Collation interface{} `json:"collation,omitempty"`
	// number of documents to skip
	Skip *int64 `json:"skip,omitempty"`
	// maximum number of documents to return
	Limit *int64 `json:"limit,omitempty"`
	// maximum query execution time in milliseconds
	MaxTimeMs *int64 `json:"maxTimeMS,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoFindRequest instantiates a new MongoFindRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoFindRequest() *MongoFindRequest {
	this := MongoFindRequest{}
	return &this
}

// NewMongoFindRequestWithDefaults instantiates a new MongoFindRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoFindRequestWithDefaults() *MongoFindRequest {
	this := MongoFindRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MongoFindRequest) GetFilter() interface{} {
	if o == nil || o.Filter == nil {
		var ret interface{}
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoFindRequest) GetFilterOk() (*interface{}, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return &o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MongoFindRequest) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given interface{} and assigns it to the Filter field.
func (o *MongoFindRequest) SetFilter(v interface{}) {
	o.Filter = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *MongoFindRequest) GetSort() interface{} {
	if o == nil || o.Sort == nil {
		var ret interface{}
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoFindRequest) GetSortOk() (*interface{}, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return &o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *MongoFindRequest) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given interface{} and assigns it to the Sort field.
func (o *MongoFindRequest) SetSort(v interface{}) {
	o.Sort = v
}

// GetProjection returns the Projection field value if set, zero value otherwise.
func (o *MongoFindRequest) GetProjection() interface{} {
	if o == nil || o.Projection == nil {
		var ret interface{}
		return ret
	}
	return o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoFindRequest) GetProjectionOk() (*interface{}, bool) {
	if o == nil || o.Projection == nil {
		return nil, false
	}
	return &o.Projection, true
}

// HasProjection returns a boolean if a field has been set.
func (o *MongoFindRequest) HasProjection() bool {
	return o != nil && o.Projection != nil
}

// SetProjection gets a reference to the given interface{} and assigns it to the Projection field.
func (o *MongoFindRequest) SetProjection(v interface{}) {
	o.Projection = v
}

// GetCollation returns the Collation field value if set, zero value otherwise.
func (o *MongoFindRequest) GetCollation() interface{} {
	if o == nil || o.Collation == nil {
		var ret interface{}
		return ret
	}
	return o.Collation
}

// GetCollationOk returns a tuple with the Collation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoFindRequest) GetCollationOk() (*interface{}, bool) {
	if o == nil || o.Collation == nil {
		return nil, false
	}
	return &o.Collation, true
}

// HasCollation returns a boolean if a field has been set.
func (o *MongoFindRequest) HasCollation() bool {
	return o != nil && o.Collation != nil
}

// SetCollation gets a reference to the given interface{} and assigns it to the Collation field.
func (o *MongoFindRequest) SetCollation(v interface{}) {
	o.Collation = v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *MongoFindRequest) GetSkip() int64 {
	if o == nil || o.Skip == nil {
		var ret int64
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoFindRequest) GetSkipOk() (*int64, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *MongoFindRequest) HasSkip() bool {
	return o != nil && o.Skip != nil
}

// SetSkip gets a reference to the given int64 and assigns it to the Skip field.
func (o *MongoFindRequest) SetSkip(v int64) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *MongoFindRequest) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoFindRequest) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *MongoFindRequest) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *MongoFindRequest) SetLimit(v int64) {
	o.Limit = &v
}

// GetMaxTimeMs returns the MaxTimeMs field value if set, zero value otherwise.
func (o *MongoFindRequest) GetMaxTimeMs() int64 {
	if o == nil || o.MaxTimeMs == nil {
		var ret int64
		return ret
	}
	return *o.MaxTimeMs
}

// GetMaxTimeMsOk returns a tuple with the MaxTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoFindRequest) GetMaxTimeMsOk() (*int64, bool) {
	if o == nil || o.MaxTimeMs == nil {
		return nil, false
	}
	return o.MaxTimeMs, true
}

// HasMaxTimeMs returns a boolean if a field has been set.
func (o *MongoFindRequest) HasMaxTimeMs() bool {
	return o != nil && o.MaxTimeMs != nil
}

// SetMaxTimeMs gets a reference to the given int64 and assigns it to the MaxTimeMs field.
func (o *MongoFindRequest) SetMaxTimeMs(v int64) {
	o.MaxTimeMs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoFindRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Projection != nil {
		toSerialize["projection"] = o.Projection
	}
	if o.Collation != nil {
		toSerialize["collation"] = o.Collation
	}
	if o.Skip != nil {
		toSerialize["skip"] = o.Skip
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.MaxTimeMs != nil {
		toSerialize["maxTimeMS"] = o.MaxTimeMs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoFindRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter     interface{} `json:"filter,omitempty"`
		Sort       interface{} `json:"sort,omitempty"`
		Projection interface{} `json:"projection,omitempty"`
		Collation  interface{} `json:"collation,omitempty"`
		Skip       *int64      `json:"skip,omitempty"`
		Limit      *int64      `json:"limit,omitempty"`
		MaxTimeMs  *int64      `json:"maxTimeMS,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"filter", "sort", "projection", "collation", "skip", "limit", "maxTimeMS"})
	} else {
		return err
	}
	o.Filter = all.Filter
	o.Sort = all.Sort
	o.Projection = all.Projection
	o.Collation = all.Collation
	o.Skip = all.Skip
	o.Limit = all.Limit
	o.MaxTimeMs = all.MaxTimeMs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
