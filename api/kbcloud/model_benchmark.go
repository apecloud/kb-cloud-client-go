// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "time"

// Benchmark Benchmark is the benchmark object
type Benchmark struct {
	// ID of benchmark
	Id *string `json:"id,omitempty"`
	// Name of benchmark
	Name *string `json:"name,omitempty"`
	// Type of benchmark
	Type *BenchmarkType `json:"type,omitempty"`
	// Config of benchmark
	Config *string `json:"config,omitempty"`
	// the log of benchmark in prepare stage
	PrepareLog *string `json:"prepareLog,omitempty"`
	// the log of benchmark in run stage
	RunLog *string `json:"runLog,omitempty"`
	// the log of benchmark in cleanup stage
	CleanupLog *string `json:"cleanupLog,omitempty"`
	// the result of benchmark, only available when benchmark is complete
	Result *string `json:"result,omitempty"`
	// the message of benchmark, only available when benchmark is failed
	Message *string `json:"message,omitempty"`
	// the cluster name
	Cluster *string `json:"cluster,omitempty"`
	// the cluster id
	ClusterId *string `json:"clusterID,omitempty"`
	// the database name
	Database *string `json:"database,omitempty"`
	// the completion timestamp of benchmark
	CompletionTimestamp *time.Time `json:"completionTimestamp,omitempty"`
	// the create timestamp of benchmark
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// the status of benchmark
	Status *BenchmarkStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBenchmark instantiates a new Benchmark object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBenchmark() *Benchmark {
	this := Benchmark{}
	return &this
}

// NewBenchmarkWithDefaults instantiates a new Benchmark object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBenchmarkWithDefaults() *Benchmark {
	this := Benchmark{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Benchmark) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Benchmark) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Benchmark) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Benchmark) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Benchmark) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Benchmark) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Benchmark) GetType() BenchmarkType {
	if o == nil || o.Type == nil {
		var ret BenchmarkType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetTypeOk() (*BenchmarkType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Benchmark) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given BenchmarkType and assigns it to the Type field.
func (o *Benchmark) SetType(v BenchmarkType) {
	o.Type = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Benchmark) GetConfig() string {
	if o == nil || o.Config == nil {
		var ret string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetConfigOk() (*string, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Benchmark) HasConfig() bool {
	return o != nil && o.Config != nil
}

// SetConfig gets a reference to the given string and assigns it to the Config field.
func (o *Benchmark) SetConfig(v string) {
	o.Config = &v
}

// GetPrepareLog returns the PrepareLog field value if set, zero value otherwise.
func (o *Benchmark) GetPrepareLog() string {
	if o == nil || o.PrepareLog == nil {
		var ret string
		return ret
	}
	return *o.PrepareLog
}

// GetPrepareLogOk returns a tuple with the PrepareLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetPrepareLogOk() (*string, bool) {
	if o == nil || o.PrepareLog == nil {
		return nil, false
	}
	return o.PrepareLog, true
}

// HasPrepareLog returns a boolean if a field has been set.
func (o *Benchmark) HasPrepareLog() bool {
	return o != nil && o.PrepareLog != nil
}

// SetPrepareLog gets a reference to the given string and assigns it to the PrepareLog field.
func (o *Benchmark) SetPrepareLog(v string) {
	o.PrepareLog = &v
}

// GetRunLog returns the RunLog field value if set, zero value otherwise.
func (o *Benchmark) GetRunLog() string {
	if o == nil || o.RunLog == nil {
		var ret string
		return ret
	}
	return *o.RunLog
}

// GetRunLogOk returns a tuple with the RunLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetRunLogOk() (*string, bool) {
	if o == nil || o.RunLog == nil {
		return nil, false
	}
	return o.RunLog, true
}

// HasRunLog returns a boolean if a field has been set.
func (o *Benchmark) HasRunLog() bool {
	return o != nil && o.RunLog != nil
}

// SetRunLog gets a reference to the given string and assigns it to the RunLog field.
func (o *Benchmark) SetRunLog(v string) {
	o.RunLog = &v
}

// GetCleanupLog returns the CleanupLog field value if set, zero value otherwise.
func (o *Benchmark) GetCleanupLog() string {
	if o == nil || o.CleanupLog == nil {
		var ret string
		return ret
	}
	return *o.CleanupLog
}

// GetCleanupLogOk returns a tuple with the CleanupLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetCleanupLogOk() (*string, bool) {
	if o == nil || o.CleanupLog == nil {
		return nil, false
	}
	return o.CleanupLog, true
}

// HasCleanupLog returns a boolean if a field has been set.
func (o *Benchmark) HasCleanupLog() bool {
	return o != nil && o.CleanupLog != nil
}

// SetCleanupLog gets a reference to the given string and assigns it to the CleanupLog field.
func (o *Benchmark) SetCleanupLog(v string) {
	o.CleanupLog = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *Benchmark) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *Benchmark) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *Benchmark) SetResult(v string) {
	o.Result = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Benchmark) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Benchmark) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Benchmark) SetMessage(v string) {
	o.Message = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *Benchmark) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *Benchmark) HasCluster() bool {
	return o != nil && o.Cluster != nil
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *Benchmark) SetCluster(v string) {
	o.Cluster = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *Benchmark) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *Benchmark) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *Benchmark) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *Benchmark) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *Benchmark) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *Benchmark) SetDatabase(v string) {
	o.Database = &v
}

// GetCompletionTimestamp returns the CompletionTimestamp field value if set, zero value otherwise.
func (o *Benchmark) GetCompletionTimestamp() time.Time {
	if o == nil || o.CompletionTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletionTimestamp
}

// GetCompletionTimestampOk returns a tuple with the CompletionTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetCompletionTimestampOk() (*time.Time, bool) {
	if o == nil || o.CompletionTimestamp == nil {
		return nil, false
	}
	return o.CompletionTimestamp, true
}

// HasCompletionTimestamp returns a boolean if a field has been set.
func (o *Benchmark) HasCompletionTimestamp() bool {
	return o != nil && o.CompletionTimestamp != nil
}

// SetCompletionTimestamp gets a reference to the given time.Time and assigns it to the CompletionTimestamp field.
func (o *Benchmark) SetCompletionTimestamp(v time.Time) {
	o.CompletionTimestamp = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Benchmark) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Benchmark) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Benchmark) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Benchmark) GetStatus() BenchmarkStatus {
	if o == nil || o.Status == nil {
		var ret BenchmarkStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Benchmark) GetStatusOk() (*BenchmarkStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Benchmark) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given BenchmarkStatus and assigns it to the Status field.
func (o *Benchmark) SetStatus(v BenchmarkStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Benchmark) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.PrepareLog != nil {
		toSerialize["prepareLog"] = o.PrepareLog
	}
	if o.RunLog != nil {
		toSerialize["runLog"] = o.RunLog
	}
	if o.CleanupLog != nil {
		toSerialize["cleanupLog"] = o.CleanupLog
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.ClusterId != nil {
		toSerialize["clusterID"] = o.ClusterId
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.CompletionTimestamp != nil {
		if o.CompletionTimestamp.Nanosecond() == 0 {
			toSerialize["completionTimestamp"] = o.CompletionTimestamp.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["completionTimestamp"] = o.CompletionTimestamp.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Benchmark) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                  *string          `json:"id,omitempty"`
		Name                *string          `json:"name,omitempty"`
		Type                *BenchmarkType   `json:"type,omitempty"`
		Config              *string          `json:"config,omitempty"`
		PrepareLog          *string          `json:"prepareLog,omitempty"`
		RunLog              *string          `json:"runLog,omitempty"`
		CleanupLog          *string          `json:"cleanupLog,omitempty"`
		Result              *string          `json:"result,omitempty"`
		Message             *string          `json:"message,omitempty"`
		Cluster             *string          `json:"cluster,omitempty"`
		ClusterId           *string          `json:"clusterID,omitempty"`
		Database            *string          `json:"database,omitempty"`
		CompletionTimestamp *time.Time       `json:"completionTimestamp,omitempty"`
		CreatedAt           *time.Time       `json:"createdAt,omitempty"`
		Status              *BenchmarkStatus `json:"status,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "type", "config", "prepareLog", "runLog", "cleanupLog", "result", "message", "cluster", "clusterID", "database", "completionTimestamp", "createdAt", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Name = all.Name
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.Config = all.Config
	o.PrepareLog = all.PrepareLog
	o.RunLog = all.RunLog
	o.CleanupLog = all.CleanupLog
	o.Result = all.Result
	o.Message = all.Message
	o.Cluster = all.Cluster
	o.ClusterId = all.ClusterId
	o.Database = all.Database
	o.CompletionTimestamp = all.CompletionTimestamp
	o.CreatedAt = all.CreatedAt
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
