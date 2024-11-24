// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

// Ycsb ycsb is the ycsb benchmark object
type Ycsb struct {
	// Step of benchmark
	Step *YcsbStep `json:"step,omitempty"`
	// the cpu limit for test container
	LimitCpu *string `json:"limitCpu,omitempty"`
	// the memory limit for test container
	LimitMemory *string `json:"limitMemory,omitempty"`
	// the cpu request for test container
	RequestCpu *string `json:"requestCpu,omitempty"`
	// the memory request for test container
	RequestMemory *string `json:"requestMemory,omitempty"`
	// the name of benchmark
	Name *string `json:"name,omitempty"`
	// the database name
	Database *string `json:"database,omitempty"`
	// the cluster name
	Cluster string `json:"cluster"`
	// Username for database
	Username string `json:"username"`
	// Password for database
	Password string `json:"password"`
	// Address for database
	Address string `json:"address"`
	// Number of records to load into database
	RecordCount common.NullableInt32 `json:"recordCount,omitempty"`
	// Number of operations to perform
	OperationCount common.NullableInt32 `json:"operationCount,omitempty"`
	// Percentage of read operations
	ReadProportion common.NullableInt32 `json:"readProportion,omitempty"`
	// Percentage of update operations
	UpdateProportion common.NullableInt32 `json:"updateProportion,omitempty"`
	// Percentage of insert operations
	InsertProportion *int32 `json:"insertProportion,omitempty"`
	// Percentage of read-modify-write operations
	ReadModifyWriteProportion *int32 `json:"readModifyWriteProportion,omitempty"`
	// Percentage of scan operations
	ScanProportion *int32 `json:"scanProportion,omitempty"`
	// Number of client threads to run
	Threads common.NullableInt32 `json:"threads,omitempty"`
	// Extra arguments for ycsb
	ExtraArgs *string `json:"extraArgs,omitempty"`
	// Redis mode
	RedisMode *YcsbRedisMode `json:"redisMode,omitempty"`
	// Redis sentinel master
	RedisSentinelMaster *string `json:"redisSentinelMaster,omitempty"`
	// Redis sentinel username
	RedisSentinelUsername *string `json:"redisSentinelUsername,omitempty"`
	// Redis sentinel password
	RedisSentinelPassword *string `json:"redisSentinelPassword,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewYcsb instantiates a new Ycsb object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewYcsb(cluster string, username string, password string, address string) *Ycsb {
	this := Ycsb{}
	var step YcsbStep = YcsbStepAll
	this.Step = &step
	var limitCpu string = "1"
	this.LimitCpu = &limitCpu
	var limitMemory string = "2Gi"
	this.LimitMemory = &limitMemory
	var requestCpu string = "0.25"
	this.RequestCpu = &requestCpu
	var requestMemory string = "0.5Gi"
	this.RequestMemory = &requestMemory
	var database string = "test"
	this.Database = &database
	this.Cluster = cluster
	this.Username = username
	this.Password = password
	this.Address = address
	return &this
}

// NewYcsbWithDefaults instantiates a new Ycsb object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewYcsbWithDefaults() *Ycsb {
	this := Ycsb{}
	var step YcsbStep = YcsbStepAll
	this.Step = &step
	var limitCpu string = "1"
	this.LimitCpu = &limitCpu
	var limitMemory string = "2Gi"
	this.LimitMemory = &limitMemory
	var requestCpu string = "0.25"
	this.RequestCpu = &requestCpu
	var requestMemory string = "0.5Gi"
	this.RequestMemory = &requestMemory
	var database string = "test"
	this.Database = &database
	return &this
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *Ycsb) GetStep() YcsbStep {
	if o == nil || o.Step == nil {
		var ret YcsbStep
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ycsb) GetStepOk() (*YcsbStep, bool) {
	if o == nil || o.Step == nil {
		return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *Ycsb) HasStep() bool {
	return o != nil && o.Step != nil
}

// SetStep gets a reference to the given YcsbStep and assigns it to the Step field.
func (o *Ycsb) SetStep(v YcsbStep) {
	o.Step = &v
}

// GetLimitCpu returns the LimitCpu field value if set, zero value otherwise.
func (o *Ycsb) GetLimitCpu() string {
	if o == nil || o.LimitCpu == nil {
		var ret string
		return ret
	}
	return *o.LimitCpu
}

// GetLimitCpuOk returns a tuple with the LimitCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ycsb) GetLimitCpuOk() (*string, bool) {
	if o == nil || o.LimitCpu == nil {
		return nil, false
	}
	return o.LimitCpu, true
}

// HasLimitCpu returns a boolean if a field has been set.
func (o *Ycsb) HasLimitCpu() bool {
	return o != nil && o.LimitCpu != nil
}

// SetLimitCpu gets a reference to the given string and assigns it to the LimitCpu field.
func (o *Ycsb) SetLimitCpu(v string) {
	o.LimitCpu = &v
}

// GetLimitMemory returns the LimitMemory field value if set, zero value otherwise.
func (o *Ycsb) GetLimitMemory() string {
	if o == nil || o.LimitMemory == nil {
		var ret string
		return ret
	}
	return *o.LimitMemory
}

// GetLimitMemoryOk returns a tuple with the LimitMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ycsb) GetLimitMemoryOk() (*string, bool) {
	if o == nil || o.LimitMemory == nil {
		return nil, false
	}
	return o.LimitMemory, true
}

// HasLimitMemory returns a boolean if a field has been set.
func (o *Ycsb) HasLimitMemory() bool {
	return o != nil && o.LimitMemory != nil
}

// SetLimitMemory gets a reference to the given string and assigns it to the LimitMemory field.
func (o *Ycsb) SetLimitMemory(v string) {
	o.LimitMemory = &v
}

// GetRequestCpu returns the RequestCpu field value if set, zero value otherwise.
func (o *Ycsb) GetRequestCpu() string {
	if o == nil || o.RequestCpu == nil {
		var ret string
		return ret
	}
	return *o.RequestCpu
}

// GetRequestCpuOk returns a tuple with the RequestCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ycsb) GetRequestCpuOk() (*string, bool) {
	if o == nil || o.RequestCpu == nil {
		return nil, false
	}
	return o.RequestCpu, true
}

// HasRequestCpu returns a boolean if a field has been set.
func (o *Ycsb) HasRequestCpu() bool {
	return o != nil && o.RequestCpu != nil
}

// SetRequestCpu gets a reference to the given string and assigns it to the RequestCpu field.
func (o *Ycsb) SetRequestCpu(v string) {
	o.RequestCpu = &v
}

// GetRequestMemory returns the RequestMemory field value if set, zero value otherwise.
func (o *Ycsb) GetRequestMemory() string {
	if o == nil || o.RequestMemory == nil {
		var ret string
		return ret
	}
	return *o.RequestMemory
}

// GetRequestMemoryOk returns a tuple with the RequestMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ycsb) GetRequestMemoryOk() (*string, bool) {
	if o == nil || o.RequestMemory == nil {
		return nil, false
	}
	return o.RequestMemory, true
}

// HasRequestMemory returns a boolean if a field has been set.
func (o *Ycsb) HasRequestMemory() bool {
	return o != nil && o.RequestMemory != nil
}

// SetRequestMemory gets a reference to the given string and assigns it to the RequestMemory field.
func (o *Ycsb) SetRequestMemory(v string) {
	o.RequestMemory = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Ycsb) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ycsb) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ycsb) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ycsb) SetName(v string) {
	o.Name = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *Ycsb) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ycsb) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *Ycsb) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *Ycsb) SetDatabase(v string) {
	o.Database = &v
}

// GetCluster returns the Cluster field value.
func (o *Ycsb) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *Ycsb) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value.
func (o *Ycsb) SetCluster(v string) {
	o.Cluster = v
}

// GetUsername returns the Username field value.
func (o *Ycsb) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Ycsb) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value.
func (o *Ycsb) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value.
func (o *Ycsb) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *Ycsb) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *Ycsb) SetPassword(v string) {
	o.Password = v
}

// GetAddress returns the Address field value.
func (o *Ycsb) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *Ycsb) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value.
func (o *Ycsb) SetAddress(v string) {
	o.Address = v
}

// GetRecordCount returns the RecordCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ycsb) GetRecordCount() int32 {
	if o == nil || o.RecordCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.RecordCount.Get()
}

// GetRecordCountOk returns a tuple with the RecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Ycsb) GetRecordCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecordCount.Get(), o.RecordCount.IsSet()
}

// HasRecordCount returns a boolean if a field has been set.
func (o *Ycsb) HasRecordCount() bool {
	return o != nil && o.RecordCount.IsSet()
}

// SetRecordCount gets a reference to the given common.NullableInt32 and assigns it to the RecordCount field.
func (o *Ycsb) SetRecordCount(v int32) {
	o.RecordCount.Set(&v)
}

// SetRecordCountNil sets the value for RecordCount to be an explicit nil.
func (o *Ycsb) SetRecordCountNil() {
	o.RecordCount.Set(nil)
}

// UnsetRecordCount ensures that no value is present for RecordCount, not even an explicit nil.
func (o *Ycsb) UnsetRecordCount() {
	o.RecordCount.Unset()
}

// GetOperationCount returns the OperationCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ycsb) GetOperationCount() int32 {
	if o == nil || o.OperationCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.OperationCount.Get()
}

// GetOperationCountOk returns a tuple with the OperationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Ycsb) GetOperationCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperationCount.Get(), o.OperationCount.IsSet()
}

// HasOperationCount returns a boolean if a field has been set.
func (o *Ycsb) HasOperationCount() bool {
	return o != nil && o.OperationCount.IsSet()
}

// SetOperationCount gets a reference to the given common.NullableInt32 and assigns it to the OperationCount field.
func (o *Ycsb) SetOperationCount(v int32) {
	o.OperationCount.Set(&v)
}

// SetOperationCountNil sets the value for OperationCount to be an explicit nil.
func (o *Ycsb) SetOperationCountNil() {
	o.OperationCount.Set(nil)
}

// UnsetOperationCount ensures that no value is present for OperationCount, not even an explicit nil.
func (o *Ycsb) UnsetOperationCount() {
	o.OperationCount.Unset()
}

// GetReadProportion returns the ReadProportion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ycsb) GetReadProportion() int32 {
	if o == nil || o.ReadProportion.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ReadProportion.Get()
}

// GetReadProportionOk returns a tuple with the ReadProportion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Ycsb) GetReadProportionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReadProportion.Get(), o.ReadProportion.IsSet()
}

// HasReadProportion returns a boolean if a field has been set.
func (o *Ycsb) HasReadProportion() bool {
	return o != nil && o.ReadProportion.IsSet()
}

// SetReadProportion gets a reference to the given common.NullableInt32 and assigns it to the ReadProportion field.
func (o *Ycsb) SetReadProportion(v int32) {
	o.ReadProportion.Set(&v)
}

// SetReadProportionNil sets the value for ReadProportion to be an explicit nil.
func (o *Ycsb) SetReadProportionNil() {
	o.ReadProportion.Set(nil)
}

// UnsetReadProportion ensures that no value is present for ReadProportion, not even an explicit nil.
func (o *Ycsb) UnsetReadProportion() {
	o.ReadProportion.Unset()
}

// GetUpdateProportion returns the UpdateProportion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ycsb) GetUpdateProportion() int32 {
	if o == nil || o.UpdateProportion.Get() == nil {
		var ret int32
		return ret
	}
	return *o.UpdateProportion.Get()
}

// GetUpdateProportionOk returns a tuple with the UpdateProportion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Ycsb) GetUpdateProportionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdateProportion.Get(), o.UpdateProportion.IsSet()
}

// HasUpdateProportion returns a boolean if a field has been set.
func (o *Ycsb) HasUpdateProportion() bool {
	return o != nil && o.UpdateProportion.IsSet()
}

// SetUpdateProportion gets a reference to the given common.NullableInt32 and assigns it to the UpdateProportion field.
func (o *Ycsb) SetUpdateProportion(v int32) {
	o.UpdateProportion.Set(&v)
}

// SetUpdateProportionNil sets the value for UpdateProportion to be an explicit nil.
func (o *Ycsb) SetUpdateProportionNil() {
	o.UpdateProportion.Set(nil)
}

// UnsetUpdateProportion ensures that no value is present for UpdateProportion, not even an explicit nil.
func (o *Ycsb) UnsetUpdateProportion() {
	o.UpdateProportion.Unset()
}

// GetInsertProportion returns the InsertProportion field value if set, zero value otherwise.
func (o *Ycsb) GetInsertProportion() int32 {
	if o == nil || o.InsertProportion == nil {
		var ret int32
		return ret
	}
	return *o.InsertProportion
}

// GetInsertProportionOk returns a tuple with the InsertProportion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ycsb) GetInsertProportionOk() (*int32, bool) {
	if o == nil || o.InsertProportion == nil {
		return nil, false
	}
	return o.InsertProportion, true
}

// HasInsertProportion returns a boolean if a field has been set.
func (o *Ycsb) HasInsertProportion() bool {
	return o != nil && o.InsertProportion != nil
}

// SetInsertProportion gets a reference to the given int32 and assigns it to the InsertProportion field.
func (o *Ycsb) SetInsertProportion(v int32) {
	o.InsertProportion = &v
}

// GetReadModifyWriteProportion returns the ReadModifyWriteProportion field value if set, zero value otherwise.
func (o *Ycsb) GetReadModifyWriteProportion() int32 {
	if o == nil || o.ReadModifyWriteProportion == nil {
		var ret int32
		return ret
	}
	return *o.ReadModifyWriteProportion
}

// GetReadModifyWriteProportionOk returns a tuple with the ReadModifyWriteProportion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ycsb) GetReadModifyWriteProportionOk() (*int32, bool) {
	if o == nil || o.ReadModifyWriteProportion == nil {
		return nil, false
	}
	return o.ReadModifyWriteProportion, true
}

// HasReadModifyWriteProportion returns a boolean if a field has been set.
func (o *Ycsb) HasReadModifyWriteProportion() bool {
	return o != nil && o.ReadModifyWriteProportion != nil
}

// SetReadModifyWriteProportion gets a reference to the given int32 and assigns it to the ReadModifyWriteProportion field.
func (o *Ycsb) SetReadModifyWriteProportion(v int32) {
	o.ReadModifyWriteProportion = &v
}

// GetScanProportion returns the ScanProportion field value if set, zero value otherwise.
func (o *Ycsb) GetScanProportion() int32 {
	if o == nil || o.ScanProportion == nil {
		var ret int32
		return ret
	}
	return *o.ScanProportion
}

// GetScanProportionOk returns a tuple with the ScanProportion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ycsb) GetScanProportionOk() (*int32, bool) {
	if o == nil || o.ScanProportion == nil {
		return nil, false
	}
	return o.ScanProportion, true
}

// HasScanProportion returns a boolean if a field has been set.
func (o *Ycsb) HasScanProportion() bool {
	return o != nil && o.ScanProportion != nil
}

// SetScanProportion gets a reference to the given int32 and assigns it to the ScanProportion field.
func (o *Ycsb) SetScanProportion(v int32) {
	o.ScanProportion = &v
}

// GetThreads returns the Threads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ycsb) GetThreads() int32 {
	if o == nil || o.Threads.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Threads.Get()
}

// GetThreadsOk returns a tuple with the Threads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Ycsb) GetThreadsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Threads.Get(), o.Threads.IsSet()
}

// HasThreads returns a boolean if a field has been set.
func (o *Ycsb) HasThreads() bool {
	return o != nil && o.Threads.IsSet()
}

// SetThreads gets a reference to the given common.NullableInt32 and assigns it to the Threads field.
func (o *Ycsb) SetThreads(v int32) {
	o.Threads.Set(&v)
}

// SetThreadsNil sets the value for Threads to be an explicit nil.
func (o *Ycsb) SetThreadsNil() {
	o.Threads.Set(nil)
}

// UnsetThreads ensures that no value is present for Threads, not even an explicit nil.
func (o *Ycsb) UnsetThreads() {
	o.Threads.Unset()
}

// GetExtraArgs returns the ExtraArgs field value if set, zero value otherwise.
func (o *Ycsb) GetExtraArgs() string {
	if o == nil || o.ExtraArgs == nil {
		var ret string
		return ret
	}
	return *o.ExtraArgs
}

// GetExtraArgsOk returns a tuple with the ExtraArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ycsb) GetExtraArgsOk() (*string, bool) {
	if o == nil || o.ExtraArgs == nil {
		return nil, false
	}
	return o.ExtraArgs, true
}

// HasExtraArgs returns a boolean if a field has been set.
func (o *Ycsb) HasExtraArgs() bool {
	return o != nil && o.ExtraArgs != nil
}

// SetExtraArgs gets a reference to the given string and assigns it to the ExtraArgs field.
func (o *Ycsb) SetExtraArgs(v string) {
	o.ExtraArgs = &v
}

// GetRedisMode returns the RedisMode field value if set, zero value otherwise.
func (o *Ycsb) GetRedisMode() YcsbRedisMode {
	if o == nil || o.RedisMode == nil {
		var ret YcsbRedisMode
		return ret
	}
	return *o.RedisMode
}

// GetRedisModeOk returns a tuple with the RedisMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ycsb) GetRedisModeOk() (*YcsbRedisMode, bool) {
	if o == nil || o.RedisMode == nil {
		return nil, false
	}
	return o.RedisMode, true
}

// HasRedisMode returns a boolean if a field has been set.
func (o *Ycsb) HasRedisMode() bool {
	return o != nil && o.RedisMode != nil
}

// SetRedisMode gets a reference to the given YcsbRedisMode and assigns it to the RedisMode field.
func (o *Ycsb) SetRedisMode(v YcsbRedisMode) {
	o.RedisMode = &v
}

// GetRedisSentinelMaster returns the RedisSentinelMaster field value if set, zero value otherwise.
func (o *Ycsb) GetRedisSentinelMaster() string {
	if o == nil || o.RedisSentinelMaster == nil {
		var ret string
		return ret
	}
	return *o.RedisSentinelMaster
}

// GetRedisSentinelMasterOk returns a tuple with the RedisSentinelMaster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ycsb) GetRedisSentinelMasterOk() (*string, bool) {
	if o == nil || o.RedisSentinelMaster == nil {
		return nil, false
	}
	return o.RedisSentinelMaster, true
}

// HasRedisSentinelMaster returns a boolean if a field has been set.
func (o *Ycsb) HasRedisSentinelMaster() bool {
	return o != nil && o.RedisSentinelMaster != nil
}

// SetRedisSentinelMaster gets a reference to the given string and assigns it to the RedisSentinelMaster field.
func (o *Ycsb) SetRedisSentinelMaster(v string) {
	o.RedisSentinelMaster = &v
}

// GetRedisSentinelUsername returns the RedisSentinelUsername field value if set, zero value otherwise.
func (o *Ycsb) GetRedisSentinelUsername() string {
	if o == nil || o.RedisSentinelUsername == nil {
		var ret string
		return ret
	}
	return *o.RedisSentinelUsername
}

// GetRedisSentinelUsernameOk returns a tuple with the RedisSentinelUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ycsb) GetRedisSentinelUsernameOk() (*string, bool) {
	if o == nil || o.RedisSentinelUsername == nil {
		return nil, false
	}
	return o.RedisSentinelUsername, true
}

// HasRedisSentinelUsername returns a boolean if a field has been set.
func (o *Ycsb) HasRedisSentinelUsername() bool {
	return o != nil && o.RedisSentinelUsername != nil
}

// SetRedisSentinelUsername gets a reference to the given string and assigns it to the RedisSentinelUsername field.
func (o *Ycsb) SetRedisSentinelUsername(v string) {
	o.RedisSentinelUsername = &v
}

// GetRedisSentinelPassword returns the RedisSentinelPassword field value if set, zero value otherwise.
func (o *Ycsb) GetRedisSentinelPassword() string {
	if o == nil || o.RedisSentinelPassword == nil {
		var ret string
		return ret
	}
	return *o.RedisSentinelPassword
}

// GetRedisSentinelPasswordOk returns a tuple with the RedisSentinelPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ycsb) GetRedisSentinelPasswordOk() (*string, bool) {
	if o == nil || o.RedisSentinelPassword == nil {
		return nil, false
	}
	return o.RedisSentinelPassword, true
}

// HasRedisSentinelPassword returns a boolean if a field has been set.
func (o *Ycsb) HasRedisSentinelPassword() bool {
	return o != nil && o.RedisSentinelPassword != nil
}

// SetRedisSentinelPassword gets a reference to the given string and assigns it to the RedisSentinelPassword field.
func (o *Ycsb) SetRedisSentinelPassword(v string) {
	o.RedisSentinelPassword = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Ycsb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Step != nil {
		toSerialize["step"] = o.Step
	}
	if o.LimitCpu != nil {
		toSerialize["limitCpu"] = o.LimitCpu
	}
	if o.LimitMemory != nil {
		toSerialize["limitMemory"] = o.LimitMemory
	}
	if o.RequestCpu != nil {
		toSerialize["requestCpu"] = o.RequestCpu
	}
	if o.RequestMemory != nil {
		toSerialize["requestMemory"] = o.RequestMemory
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	toSerialize["cluster"] = o.Cluster
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	toSerialize["address"] = o.Address
	if o.RecordCount.IsSet() {
		toSerialize["recordCount"] = o.RecordCount.Get()
	}
	if o.OperationCount.IsSet() {
		toSerialize["operationCount"] = o.OperationCount.Get()
	}
	if o.ReadProportion.IsSet() {
		toSerialize["readProportion"] = o.ReadProportion.Get()
	}
	if o.UpdateProportion.IsSet() {
		toSerialize["updateProportion"] = o.UpdateProportion.Get()
	}
	if o.InsertProportion != nil {
		toSerialize["insertProportion"] = o.InsertProportion
	}
	if o.ReadModifyWriteProportion != nil {
		toSerialize["readModifyWriteProportion"] = o.ReadModifyWriteProportion
	}
	if o.ScanProportion != nil {
		toSerialize["scanProportion"] = o.ScanProportion
	}
	if o.Threads.IsSet() {
		toSerialize["threads"] = o.Threads.Get()
	}
	if o.ExtraArgs != nil {
		toSerialize["extraArgs"] = o.ExtraArgs
	}
	if o.RedisMode != nil {
		toSerialize["redisMode"] = o.RedisMode
	}
	if o.RedisSentinelMaster != nil {
		toSerialize["redisSentinelMaster"] = o.RedisSentinelMaster
	}
	if o.RedisSentinelUsername != nil {
		toSerialize["redisSentinelUsername"] = o.RedisSentinelUsername
	}
	if o.RedisSentinelPassword != nil {
		toSerialize["redisSentinelPassword"] = o.RedisSentinelPassword
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Ycsb) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Step                      *YcsbStep            `json:"step,omitempty"`
		LimitCpu                  *string              `json:"limitCpu,omitempty"`
		LimitMemory               *string              `json:"limitMemory,omitempty"`
		RequestCpu                *string              `json:"requestCpu,omitempty"`
		RequestMemory             *string              `json:"requestMemory,omitempty"`
		Name                      *string              `json:"name,omitempty"`
		Database                  *string              `json:"database,omitempty"`
		Cluster                   *string              `json:"cluster"`
		Username                  *string              `json:"username"`
		Password                  *string              `json:"password"`
		Address                   *string              `json:"address"`
		RecordCount               common.NullableInt32 `json:"recordCount,omitempty"`
		OperationCount            common.NullableInt32 `json:"operationCount,omitempty"`
		ReadProportion            common.NullableInt32 `json:"readProportion,omitempty"`
		UpdateProportion          common.NullableInt32 `json:"updateProportion,omitempty"`
		InsertProportion          *int32               `json:"insertProportion,omitempty"`
		ReadModifyWriteProportion *int32               `json:"readModifyWriteProportion,omitempty"`
		ScanProportion            *int32               `json:"scanProportion,omitempty"`
		Threads                   common.NullableInt32 `json:"threads,omitempty"`
		ExtraArgs                 *string              `json:"extraArgs,omitempty"`
		RedisMode                 *YcsbRedisMode       `json:"redisMode,omitempty"`
		RedisSentinelMaster       *string              `json:"redisSentinelMaster,omitempty"`
		RedisSentinelUsername     *string              `json:"redisSentinelUsername,omitempty"`
		RedisSentinelPassword     *string              `json:"redisSentinelPassword,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Cluster == nil {
		return fmt.Errorf("required field cluster missing")
	}
	if all.Username == nil {
		return fmt.Errorf("required field username missing")
	}
	if all.Password == nil {
		return fmt.Errorf("required field password missing")
	}
	if all.Address == nil {
		return fmt.Errorf("required field address missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"step", "limitCpu", "limitMemory", "requestCpu", "requestMemory", "name", "database", "cluster", "username", "password", "address", "recordCount", "operationCount", "readProportion", "updateProportion", "insertProportion", "readModifyWriteProportion", "scanProportion", "threads", "extraArgs", "redisMode", "redisSentinelMaster", "redisSentinelUsername", "redisSentinelPassword"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Step != nil && !all.Step.IsValid() {
		hasInvalidField = true
	} else {
		o.Step = all.Step
	}
	o.LimitCpu = all.LimitCpu
	o.LimitMemory = all.LimitMemory
	o.RequestCpu = all.RequestCpu
	o.RequestMemory = all.RequestMemory
	o.Name = all.Name
	o.Database = all.Database
	o.Cluster = *all.Cluster
	o.Username = *all.Username
	o.Password = *all.Password
	o.Address = *all.Address
	o.RecordCount = all.RecordCount
	o.OperationCount = all.OperationCount
	o.ReadProportion = all.ReadProportion
	o.UpdateProportion = all.UpdateProportion
	o.InsertProportion = all.InsertProportion
	o.ReadModifyWriteProportion = all.ReadModifyWriteProportion
	o.ScanProportion = all.ScanProportion
	o.Threads = all.Threads
	o.ExtraArgs = all.ExtraArgs
	if all.RedisMode != nil && !all.RedisMode.IsValid() {
		hasInvalidField = true
	} else {
		o.RedisMode = all.RedisMode
	}
	o.RedisSentinelMaster = all.RedisSentinelMaster
	o.RedisSentinelUsername = all.RedisSentinelUsername
	o.RedisSentinelPassword = all.RedisSentinelPassword

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
