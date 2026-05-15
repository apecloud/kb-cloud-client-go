// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoExplainResponse struct {
	QueryPlanner   interface{} `json:"queryPlanner,omitempty"`
	ExecutionStats interface{} `json:"executionStats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoExplainResponse instantiates a new MongoExplainResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoExplainResponse() *MongoExplainResponse {
	this := MongoExplainResponse{}
	return &this
}

// NewMongoExplainResponseWithDefaults instantiates a new MongoExplainResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoExplainResponseWithDefaults() *MongoExplainResponse {
	this := MongoExplainResponse{}
	return &this
}

// GetQueryPlanner returns the QueryPlanner field value if set, zero value otherwise.
func (o *MongoExplainResponse) GetQueryPlanner() interface{} {
	if o == nil || o.QueryPlanner == nil {
		var ret interface{}
		return ret
	}
	return o.QueryPlanner
}

// GetQueryPlannerOk returns a tuple with the QueryPlanner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoExplainResponse) GetQueryPlannerOk() (*interface{}, bool) {
	if o == nil || o.QueryPlanner == nil {
		return nil, false
	}
	return &o.QueryPlanner, true
}

// HasQueryPlanner returns a boolean if a field has been set.
func (o *MongoExplainResponse) HasQueryPlanner() bool {
	return o != nil && o.QueryPlanner != nil
}

// SetQueryPlanner gets a reference to the given interface{} and assigns it to the QueryPlanner field.
func (o *MongoExplainResponse) SetQueryPlanner(v interface{}) {
	o.QueryPlanner = v
}

// GetExecutionStats returns the ExecutionStats field value if set, zero value otherwise.
func (o *MongoExplainResponse) GetExecutionStats() interface{} {
	if o == nil || o.ExecutionStats == nil {
		var ret interface{}
		return ret
	}
	return o.ExecutionStats
}

// GetExecutionStatsOk returns a tuple with the ExecutionStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoExplainResponse) GetExecutionStatsOk() (*interface{}, bool) {
	if o == nil || o.ExecutionStats == nil {
		return nil, false
	}
	return &o.ExecutionStats, true
}

// HasExecutionStats returns a boolean if a field has been set.
func (o *MongoExplainResponse) HasExecutionStats() bool {
	return o != nil && o.ExecutionStats != nil
}

// SetExecutionStats gets a reference to the given interface{} and assigns it to the ExecutionStats field.
func (o *MongoExplainResponse) SetExecutionStats(v interface{}) {
	o.ExecutionStats = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoExplainResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.QueryPlanner != nil {
		toSerialize["queryPlanner"] = o.QueryPlanner
	}
	if o.ExecutionStats != nil {
		toSerialize["executionStats"] = o.ExecutionStats
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoExplainResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryPlanner   interface{} `json:"queryPlanner,omitempty"`
		ExecutionStats interface{} `json:"executionStats,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"queryPlanner", "executionStats"})
	} else {
		return err
	}
	o.QueryPlanner = all.QueryPlanner
	o.ExecutionStats = all.ExecutionStats

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
