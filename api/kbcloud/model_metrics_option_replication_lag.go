// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

type MetricsOptionReplicationLag struct {
	QueryPattern *string `json:"queryPattern,omitempty"`
	// Specifies the type of metrics query to be performed.
	// 'instant' for a single point in time, 'range' for a time range.
	//
	QueryType *EngineOptionsMetricsQueryType `json:"queryType,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricsOptionReplicationLag instantiates a new MetricsOptionReplicationLag object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricsOptionReplicationLag() *MetricsOptionReplicationLag {
	this := MetricsOptionReplicationLag{}
	return &this
}

// NewMetricsOptionReplicationLagWithDefaults instantiates a new MetricsOptionReplicationLag object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricsOptionReplicationLagWithDefaults() *MetricsOptionReplicationLag {
	this := MetricsOptionReplicationLag{}
	return &this
}

// GetQueryPattern returns the QueryPattern field value if set, zero value otherwise.
func (o *MetricsOptionReplicationLag) GetQueryPattern() string {
	if o == nil || o.QueryPattern == nil {
		var ret string
		return ret
	}
	return *o.QueryPattern
}

// GetQueryPatternOk returns a tuple with the QueryPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsOptionReplicationLag) GetQueryPatternOk() (*string, bool) {
	if o == nil || o.QueryPattern == nil {
		return nil, false
	}
	return o.QueryPattern, true
}

// HasQueryPattern returns a boolean if a field has been set.
func (o *MetricsOptionReplicationLag) HasQueryPattern() bool {
	return o != nil && o.QueryPattern != nil
}

// SetQueryPattern gets a reference to the given string and assigns it to the QueryPattern field.
func (o *MetricsOptionReplicationLag) SetQueryPattern(v string) {
	o.QueryPattern = &v
}

// GetQueryType returns the QueryType field value if set, zero value otherwise.
func (o *MetricsOptionReplicationLag) GetQueryType() EngineOptionsMetricsQueryType {
	if o == nil || o.QueryType == nil {
		var ret EngineOptionsMetricsQueryType
		return ret
	}
	return *o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsOptionReplicationLag) GetQueryTypeOk() (*EngineOptionsMetricsQueryType, bool) {
	if o == nil || o.QueryType == nil {
		return nil, false
	}
	return o.QueryType, true
}

// HasQueryType returns a boolean if a field has been set.
func (o *MetricsOptionReplicationLag) HasQueryType() bool {
	return o != nil && o.QueryType != nil
}

// SetQueryType gets a reference to the given EngineOptionsMetricsQueryType and assigns it to the QueryType field.
func (o *MetricsOptionReplicationLag) SetQueryType(v EngineOptionsMetricsQueryType) {
	o.QueryType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricsOptionReplicationLag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.QueryPattern != nil {
		toSerialize["queryPattern"] = o.QueryPattern
	}
	if o.QueryType != nil {
		toSerialize["queryType"] = o.QueryType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricsOptionReplicationLag) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryPattern *string                        `json:"queryPattern,omitempty"`
		QueryType    *EngineOptionsMetricsQueryType `json:"queryType,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"queryPattern", "queryType"})
	} else {
		return err
	}

	hasInvalidField := false
	o.QueryPattern = all.QueryPattern
	if all.QueryType != nil && !all.QueryType.IsValid() {
		hasInvalidField = true
	} else {
		o.QueryType = all.QueryType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
