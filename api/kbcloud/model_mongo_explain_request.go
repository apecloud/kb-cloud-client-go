// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoExplainRequest struct {
	// database name
	Database *string `json:"database,omitempty"`
	// collection name
	Collection *string `json:"collection,omitempty"`
	// query filter to explain
	Filter interface{} `json:"filter,omitempty"`
	// sort specification
	Sort interface{} `json:"sort,omitempty"`
	// aggregation pipeline to explain
	Pipeline []interface{} `json:"pipeline,omitempty"`
	// MongoDB collation options
	Collation interface{} `json:"collation,omitempty"`
	// maximum query execution time in milliseconds
	MaxTimeMs *int64 `json:"maxTimeMS,omitempty"`
	// explain verbosity level (queryPlanner, executionStats, or allPlansExecution)
	Verbosity *string `json:"verbosity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoExplainRequest instantiates a new MongoExplainRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoExplainRequest() *MongoExplainRequest {
	this := MongoExplainRequest{}
	return &this
}

// NewMongoExplainRequestWithDefaults instantiates a new MongoExplainRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoExplainRequestWithDefaults() *MongoExplainRequest {
	this := MongoExplainRequest{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *MongoExplainRequest) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoExplainRequest) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *MongoExplainRequest) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *MongoExplainRequest) SetDatabase(v string) {
	o.Database = &v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *MongoExplainRequest) GetCollection() string {
	if o == nil || o.Collection == nil {
		var ret string
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoExplainRequest) GetCollectionOk() (*string, bool) {
	if o == nil || o.Collection == nil {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *MongoExplainRequest) HasCollection() bool {
	return o != nil && o.Collection != nil
}

// SetCollection gets a reference to the given string and assigns it to the Collection field.
func (o *MongoExplainRequest) SetCollection(v string) {
	o.Collection = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MongoExplainRequest) GetFilter() interface{} {
	if o == nil || o.Filter == nil {
		var ret interface{}
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoExplainRequest) GetFilterOk() (*interface{}, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return &o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MongoExplainRequest) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given interface{} and assigns it to the Filter field.
func (o *MongoExplainRequest) SetFilter(v interface{}) {
	o.Filter = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *MongoExplainRequest) GetSort() interface{} {
	if o == nil || o.Sort == nil {
		var ret interface{}
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoExplainRequest) GetSortOk() (*interface{}, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return &o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *MongoExplainRequest) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given interface{} and assigns it to the Sort field.
func (o *MongoExplainRequest) SetSort(v interface{}) {
	o.Sort = v
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *MongoExplainRequest) GetPipeline() []interface{} {
	if o == nil || o.Pipeline == nil {
		var ret []interface{}
		return ret
	}
	return o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoExplainRequest) GetPipelineOk() (*[]interface{}, bool) {
	if o == nil || o.Pipeline == nil {
		return nil, false
	}
	return &o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *MongoExplainRequest) HasPipeline() bool {
	return o != nil && o.Pipeline != nil
}

// SetPipeline gets a reference to the given []interface{} and assigns it to the Pipeline field.
func (o *MongoExplainRequest) SetPipeline(v []interface{}) {
	o.Pipeline = v
}

// GetCollation returns the Collation field value if set, zero value otherwise.
func (o *MongoExplainRequest) GetCollation() interface{} {
	if o == nil || o.Collation == nil {
		var ret interface{}
		return ret
	}
	return o.Collation
}

// GetCollationOk returns a tuple with the Collation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoExplainRequest) GetCollationOk() (*interface{}, bool) {
	if o == nil || o.Collation == nil {
		return nil, false
	}
	return &o.Collation, true
}

// HasCollation returns a boolean if a field has been set.
func (o *MongoExplainRequest) HasCollation() bool {
	return o != nil && o.Collation != nil
}

// SetCollation gets a reference to the given interface{} and assigns it to the Collation field.
func (o *MongoExplainRequest) SetCollation(v interface{}) {
	o.Collation = v
}

// GetMaxTimeMs returns the MaxTimeMs field value if set, zero value otherwise.
func (o *MongoExplainRequest) GetMaxTimeMs() int64 {
	if o == nil || o.MaxTimeMs == nil {
		var ret int64
		return ret
	}
	return *o.MaxTimeMs
}

// GetMaxTimeMsOk returns a tuple with the MaxTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoExplainRequest) GetMaxTimeMsOk() (*int64, bool) {
	if o == nil || o.MaxTimeMs == nil {
		return nil, false
	}
	return o.MaxTimeMs, true
}

// HasMaxTimeMs returns a boolean if a field has been set.
func (o *MongoExplainRequest) HasMaxTimeMs() bool {
	return o != nil && o.MaxTimeMs != nil
}

// SetMaxTimeMs gets a reference to the given int64 and assigns it to the MaxTimeMs field.
func (o *MongoExplainRequest) SetMaxTimeMs(v int64) {
	o.MaxTimeMs = &v
}

// GetVerbosity returns the Verbosity field value if set, zero value otherwise.
func (o *MongoExplainRequest) GetVerbosity() string {
	if o == nil || o.Verbosity == nil {
		var ret string
		return ret
	}
	return *o.Verbosity
}

// GetVerbosityOk returns a tuple with the Verbosity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoExplainRequest) GetVerbosityOk() (*string, bool) {
	if o == nil || o.Verbosity == nil {
		return nil, false
	}
	return o.Verbosity, true
}

// HasVerbosity returns a boolean if a field has been set.
func (o *MongoExplainRequest) HasVerbosity() bool {
	return o != nil && o.Verbosity != nil
}

// SetVerbosity gets a reference to the given string and assigns it to the Verbosity field.
func (o *MongoExplainRequest) SetVerbosity(v string) {
	o.Verbosity = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoExplainRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.Collection != nil {
		toSerialize["collection"] = o.Collection
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Pipeline != nil {
		toSerialize["pipeline"] = o.Pipeline
	}
	if o.Collation != nil {
		toSerialize["collation"] = o.Collation
	}
	if o.MaxTimeMs != nil {
		toSerialize["maxTimeMS"] = o.MaxTimeMs
	}
	if o.Verbosity != nil {
		toSerialize["verbosity"] = o.Verbosity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoExplainRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database   *string       `json:"database,omitempty"`
		Collection *string       `json:"collection,omitempty"`
		Filter     interface{}   `json:"filter,omitempty"`
		Sort       interface{}   `json:"sort,omitempty"`
		Pipeline   []interface{} `json:"pipeline,omitempty"`
		Collation  interface{}   `json:"collation,omitempty"`
		MaxTimeMs  *int64        `json:"maxTimeMS,omitempty"`
		Verbosity  *string       `json:"verbosity,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "collection", "filter", "sort", "pipeline", "collation", "maxTimeMS", "verbosity"})
	} else {
		return err
	}
	o.Database = all.Database
	o.Collection = all.Collection
	o.Filter = all.Filter
	o.Sort = all.Sort
	o.Pipeline = all.Pipeline
	o.Collation = all.Collation
	o.MaxTimeMs = all.MaxTimeMs
	o.Verbosity = all.Verbosity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
