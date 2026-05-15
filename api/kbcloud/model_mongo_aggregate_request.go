// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoAggregateRequest struct {
	// aggregation pipeline; write stages such as $out/$merge are rejected by DMS in S1
	Pipeline []interface{} `json:"pipeline,omitempty"`
	// maximum aggregation execution time in milliseconds
	MaxTimeMs *int64 `json:"maxTimeMS,omitempty"`
	// optional preview result limit appended when pipeline has no terminal $limit
	Limit *int64 `json:"limit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoAggregateRequest instantiates a new MongoAggregateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoAggregateRequest() *MongoAggregateRequest {
	this := MongoAggregateRequest{}
	return &this
}

// NewMongoAggregateRequestWithDefaults instantiates a new MongoAggregateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoAggregateRequestWithDefaults() *MongoAggregateRequest {
	this := MongoAggregateRequest{}
	return &this
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *MongoAggregateRequest) GetPipeline() []interface{} {
	if o == nil || o.Pipeline == nil {
		var ret []interface{}
		return ret
	}
	return o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoAggregateRequest) GetPipelineOk() (*[]interface{}, bool) {
	if o == nil || o.Pipeline == nil {
		return nil, false
	}
	return &o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *MongoAggregateRequest) HasPipeline() bool {
	return o != nil && o.Pipeline != nil
}

// SetPipeline gets a reference to the given []interface{} and assigns it to the Pipeline field.
func (o *MongoAggregateRequest) SetPipeline(v []interface{}) {
	o.Pipeline = v
}

// GetMaxTimeMs returns the MaxTimeMs field value if set, zero value otherwise.
func (o *MongoAggregateRequest) GetMaxTimeMs() int64 {
	if o == nil || o.MaxTimeMs == nil {
		var ret int64
		return ret
	}
	return *o.MaxTimeMs
}

// GetMaxTimeMsOk returns a tuple with the MaxTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoAggregateRequest) GetMaxTimeMsOk() (*int64, bool) {
	if o == nil || o.MaxTimeMs == nil {
		return nil, false
	}
	return o.MaxTimeMs, true
}

// HasMaxTimeMs returns a boolean if a field has been set.
func (o *MongoAggregateRequest) HasMaxTimeMs() bool {
	return o != nil && o.MaxTimeMs != nil
}

// SetMaxTimeMs gets a reference to the given int64 and assigns it to the MaxTimeMs field.
func (o *MongoAggregateRequest) SetMaxTimeMs(v int64) {
	o.MaxTimeMs = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *MongoAggregateRequest) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoAggregateRequest) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *MongoAggregateRequest) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *MongoAggregateRequest) SetLimit(v int64) {
	o.Limit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoAggregateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Pipeline != nil {
		toSerialize["pipeline"] = o.Pipeline
	}
	if o.MaxTimeMs != nil {
		toSerialize["maxTimeMS"] = o.MaxTimeMs
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoAggregateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pipeline  []interface{} `json:"pipeline,omitempty"`
		MaxTimeMs *int64        `json:"maxTimeMS,omitempty"`
		Limit     *int64        `json:"limit,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"pipeline", "maxTimeMS", "limit"})
	} else {
		return err
	}
	o.Pipeline = all.Pipeline
	o.MaxTimeMs = all.MaxTimeMs
	o.Limit = all.Limit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
