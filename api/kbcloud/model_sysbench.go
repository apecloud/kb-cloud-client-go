// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Sysbench sysbench is the sysbench benchmark object
type Sysbench struct {
	// Step of sysbench
	Step *SysbenchStep `json:"step,omitempty"`
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
	// the cluster name
	Cluster string `json:"cluster"`
	// the database name
	Database string `json:"database"`
	// Number of threads to use
	Threads common.NullableInt32 `json:"threads,omitempty"`
	// the seconds of test duration
	Duration common.NullableInt32 `json:"duration,omitempty"`
	// Number of rows per table
	TableSize common.NullableInt32 `json:"tableSize,omitempty"`
	// Number of tables
	TableNum common.NullableInt32 `json:"tableNum,omitempty"`
	// Test type for sysbench
	TestType *SysbenchTestType `json:"testType,omitempty"`
	// Percentage of reads, only used for oltp_read_write_pct
	ReadPercent common.NullableInt32 `json:"readPercent,omitempty"`
	// Percentage of writes, only used for oltp_read_write_pct
	WritePercent common.NullableInt32 `json:"writePercent,omitempty"`
	// Username for database
	Username string `json:"username"`
	// Password for database
	Password string `json:"password"`
	// Address for database
	Address string `json:"address"`
	// Extra arguments for sysbench
	ExtraArgs *string `json:"extraArgs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSysbench instantiates a new Sysbench object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSysbench(cluster string, database string, username string, password string, address string) *Sysbench {
	this := Sysbench{}
	var step SysbenchStep = SysbenchStepAll
	this.Step = &step
	var limitCpu string = "1"
	this.LimitCpu = &limitCpu
	var limitMemory string = "2Gi"
	this.LimitMemory = &limitMemory
	var requestCpu string = "0.25"
	this.RequestCpu = &requestCpu
	var requestMemory string = "0.5Gi"
	this.RequestMemory = &requestMemory
	this.Cluster = cluster
	this.Database = database
	var testType SysbenchTestType = SysbenchTestTypeOltpReadWrite
	this.TestType = &testType
	this.Username = username
	this.Password = password
	this.Address = address
	return &this
}

// NewSysbenchWithDefaults instantiates a new Sysbench object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSysbenchWithDefaults() *Sysbench {
	this := Sysbench{}
	var step SysbenchStep = SysbenchStepAll
	this.Step = &step
	var limitCpu string = "1"
	this.LimitCpu = &limitCpu
	var limitMemory string = "2Gi"
	this.LimitMemory = &limitMemory
	var requestCpu string = "0.25"
	this.RequestCpu = &requestCpu
	var requestMemory string = "0.5Gi"
	this.RequestMemory = &requestMemory
	var testType SysbenchTestType = SysbenchTestTypeOltpReadWrite
	this.TestType = &testType
	return &this
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *Sysbench) GetStep() SysbenchStep {
	if o == nil || o.Step == nil {
		var ret SysbenchStep
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sysbench) GetStepOk() (*SysbenchStep, bool) {
	if o == nil || o.Step == nil {
		return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *Sysbench) HasStep() bool {
	return o != nil && o.Step != nil
}

// SetStep gets a reference to the given SysbenchStep and assigns it to the Step field.
func (o *Sysbench) SetStep(v SysbenchStep) {
	o.Step = &v
}

// GetLimitCpu returns the LimitCpu field value if set, zero value otherwise.
func (o *Sysbench) GetLimitCpu() string {
	if o == nil || o.LimitCpu == nil {
		var ret string
		return ret
	}
	return *o.LimitCpu
}

// GetLimitCpuOk returns a tuple with the LimitCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sysbench) GetLimitCpuOk() (*string, bool) {
	if o == nil || o.LimitCpu == nil {
		return nil, false
	}
	return o.LimitCpu, true
}

// HasLimitCpu returns a boolean if a field has been set.
func (o *Sysbench) HasLimitCpu() bool {
	return o != nil && o.LimitCpu != nil
}

// SetLimitCpu gets a reference to the given string and assigns it to the LimitCpu field.
func (o *Sysbench) SetLimitCpu(v string) {
	o.LimitCpu = &v
}

// GetLimitMemory returns the LimitMemory field value if set, zero value otherwise.
func (o *Sysbench) GetLimitMemory() string {
	if o == nil || o.LimitMemory == nil {
		var ret string
		return ret
	}
	return *o.LimitMemory
}

// GetLimitMemoryOk returns a tuple with the LimitMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sysbench) GetLimitMemoryOk() (*string, bool) {
	if o == nil || o.LimitMemory == nil {
		return nil, false
	}
	return o.LimitMemory, true
}

// HasLimitMemory returns a boolean if a field has been set.
func (o *Sysbench) HasLimitMemory() bool {
	return o != nil && o.LimitMemory != nil
}

// SetLimitMemory gets a reference to the given string and assigns it to the LimitMemory field.
func (o *Sysbench) SetLimitMemory(v string) {
	o.LimitMemory = &v
}

// GetRequestCpu returns the RequestCpu field value if set, zero value otherwise.
func (o *Sysbench) GetRequestCpu() string {
	if o == nil || o.RequestCpu == nil {
		var ret string
		return ret
	}
	return *o.RequestCpu
}

// GetRequestCpuOk returns a tuple with the RequestCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sysbench) GetRequestCpuOk() (*string, bool) {
	if o == nil || o.RequestCpu == nil {
		return nil, false
	}
	return o.RequestCpu, true
}

// HasRequestCpu returns a boolean if a field has been set.
func (o *Sysbench) HasRequestCpu() bool {
	return o != nil && o.RequestCpu != nil
}

// SetRequestCpu gets a reference to the given string and assigns it to the RequestCpu field.
func (o *Sysbench) SetRequestCpu(v string) {
	o.RequestCpu = &v
}

// GetRequestMemory returns the RequestMemory field value if set, zero value otherwise.
func (o *Sysbench) GetRequestMemory() string {
	if o == nil || o.RequestMemory == nil {
		var ret string
		return ret
	}
	return *o.RequestMemory
}

// GetRequestMemoryOk returns a tuple with the RequestMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sysbench) GetRequestMemoryOk() (*string, bool) {
	if o == nil || o.RequestMemory == nil {
		return nil, false
	}
	return o.RequestMemory, true
}

// HasRequestMemory returns a boolean if a field has been set.
func (o *Sysbench) HasRequestMemory() bool {
	return o != nil && o.RequestMemory != nil
}

// SetRequestMemory gets a reference to the given string and assigns it to the RequestMemory field.
func (o *Sysbench) SetRequestMemory(v string) {
	o.RequestMemory = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Sysbench) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sysbench) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Sysbench) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Sysbench) SetName(v string) {
	o.Name = &v
}

// GetCluster returns the Cluster field value.
func (o *Sysbench) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *Sysbench) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value.
func (o *Sysbench) SetCluster(v string) {
	o.Cluster = v
}

// GetDatabase returns the Database field value.
func (o *Sysbench) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *Sysbench) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *Sysbench) SetDatabase(v string) {
	o.Database = v
}

// GetThreads returns the Threads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Sysbench) GetThreads() int32 {
	if o == nil || o.Threads.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Threads.Get()
}

// GetThreadsOk returns a tuple with the Threads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Sysbench) GetThreadsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Threads.Get(), o.Threads.IsSet()
}

// HasThreads returns a boolean if a field has been set.
func (o *Sysbench) HasThreads() bool {
	return o != nil && o.Threads.IsSet()
}

// SetThreads gets a reference to the given common.NullableInt32 and assigns it to the Threads field.
func (o *Sysbench) SetThreads(v int32) {
	o.Threads.Set(&v)
}

// SetThreadsNil sets the value for Threads to be an explicit nil.
func (o *Sysbench) SetThreadsNil() {
	o.Threads.Set(nil)
}

// UnsetThreads ensures that no value is present for Threads, not even an explicit nil.
func (o *Sysbench) UnsetThreads() {
	o.Threads.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Sysbench) GetDuration() int32 {
	if o == nil || o.Duration.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Sysbench) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *Sysbench) HasDuration() bool {
	return o != nil && o.Duration.IsSet()
}

// SetDuration gets a reference to the given common.NullableInt32 and assigns it to the Duration field.
func (o *Sysbench) SetDuration(v int32) {
	o.Duration.Set(&v)
}

// SetDurationNil sets the value for Duration to be an explicit nil.
func (o *Sysbench) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil.
func (o *Sysbench) UnsetDuration() {
	o.Duration.Unset()
}

// GetTableSize returns the TableSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Sysbench) GetTableSize() int32 {
	if o == nil || o.TableSize.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TableSize.Get()
}

// GetTableSizeOk returns a tuple with the TableSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Sysbench) GetTableSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TableSize.Get(), o.TableSize.IsSet()
}

// HasTableSize returns a boolean if a field has been set.
func (o *Sysbench) HasTableSize() bool {
	return o != nil && o.TableSize.IsSet()
}

// SetTableSize gets a reference to the given common.NullableInt32 and assigns it to the TableSize field.
func (o *Sysbench) SetTableSize(v int32) {
	o.TableSize.Set(&v)
}

// SetTableSizeNil sets the value for TableSize to be an explicit nil.
func (o *Sysbench) SetTableSizeNil() {
	o.TableSize.Set(nil)
}

// UnsetTableSize ensures that no value is present for TableSize, not even an explicit nil.
func (o *Sysbench) UnsetTableSize() {
	o.TableSize.Unset()
}

// GetTableNum returns the TableNum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Sysbench) GetTableNum() int32 {
	if o == nil || o.TableNum.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TableNum.Get()
}

// GetTableNumOk returns a tuple with the TableNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Sysbench) GetTableNumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TableNum.Get(), o.TableNum.IsSet()
}

// HasTableNum returns a boolean if a field has been set.
func (o *Sysbench) HasTableNum() bool {
	return o != nil && o.TableNum.IsSet()
}

// SetTableNum gets a reference to the given common.NullableInt32 and assigns it to the TableNum field.
func (o *Sysbench) SetTableNum(v int32) {
	o.TableNum.Set(&v)
}

// SetTableNumNil sets the value for TableNum to be an explicit nil.
func (o *Sysbench) SetTableNumNil() {
	o.TableNum.Set(nil)
}

// UnsetTableNum ensures that no value is present for TableNum, not even an explicit nil.
func (o *Sysbench) UnsetTableNum() {
	o.TableNum.Unset()
}

// GetTestType returns the TestType field value if set, zero value otherwise.
func (o *Sysbench) GetTestType() SysbenchTestType {
	if o == nil || o.TestType == nil {
		var ret SysbenchTestType
		return ret
	}
	return *o.TestType
}

// GetTestTypeOk returns a tuple with the TestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sysbench) GetTestTypeOk() (*SysbenchTestType, bool) {
	if o == nil || o.TestType == nil {
		return nil, false
	}
	return o.TestType, true
}

// HasTestType returns a boolean if a field has been set.
func (o *Sysbench) HasTestType() bool {
	return o != nil && o.TestType != nil
}

// SetTestType gets a reference to the given SysbenchTestType and assigns it to the TestType field.
func (o *Sysbench) SetTestType(v SysbenchTestType) {
	o.TestType = &v
}

// GetReadPercent returns the ReadPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Sysbench) GetReadPercent() int32 {
	if o == nil || o.ReadPercent.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ReadPercent.Get()
}

// GetReadPercentOk returns a tuple with the ReadPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Sysbench) GetReadPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReadPercent.Get(), o.ReadPercent.IsSet()
}

// HasReadPercent returns a boolean if a field has been set.
func (o *Sysbench) HasReadPercent() bool {
	return o != nil && o.ReadPercent.IsSet()
}

// SetReadPercent gets a reference to the given common.NullableInt32 and assigns it to the ReadPercent field.
func (o *Sysbench) SetReadPercent(v int32) {
	o.ReadPercent.Set(&v)
}

// SetReadPercentNil sets the value for ReadPercent to be an explicit nil.
func (o *Sysbench) SetReadPercentNil() {
	o.ReadPercent.Set(nil)
}

// UnsetReadPercent ensures that no value is present for ReadPercent, not even an explicit nil.
func (o *Sysbench) UnsetReadPercent() {
	o.ReadPercent.Unset()
}

// GetWritePercent returns the WritePercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Sysbench) GetWritePercent() int32 {
	if o == nil || o.WritePercent.Get() == nil {
		var ret int32
		return ret
	}
	return *o.WritePercent.Get()
}

// GetWritePercentOk returns a tuple with the WritePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Sysbench) GetWritePercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.WritePercent.Get(), o.WritePercent.IsSet()
}

// HasWritePercent returns a boolean if a field has been set.
func (o *Sysbench) HasWritePercent() bool {
	return o != nil && o.WritePercent.IsSet()
}

// SetWritePercent gets a reference to the given common.NullableInt32 and assigns it to the WritePercent field.
func (o *Sysbench) SetWritePercent(v int32) {
	o.WritePercent.Set(&v)
}

// SetWritePercentNil sets the value for WritePercent to be an explicit nil.
func (o *Sysbench) SetWritePercentNil() {
	o.WritePercent.Set(nil)
}

// UnsetWritePercent ensures that no value is present for WritePercent, not even an explicit nil.
func (o *Sysbench) UnsetWritePercent() {
	o.WritePercent.Unset()
}

// GetUsername returns the Username field value.
func (o *Sysbench) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Sysbench) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value.
func (o *Sysbench) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value.
func (o *Sysbench) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *Sysbench) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *Sysbench) SetPassword(v string) {
	o.Password = v
}

// GetAddress returns the Address field value.
func (o *Sysbench) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *Sysbench) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value.
func (o *Sysbench) SetAddress(v string) {
	o.Address = v
}

// GetExtraArgs returns the ExtraArgs field value if set, zero value otherwise.
func (o *Sysbench) GetExtraArgs() string {
	if o == nil || o.ExtraArgs == nil {
		var ret string
		return ret
	}
	return *o.ExtraArgs
}

// GetExtraArgsOk returns a tuple with the ExtraArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sysbench) GetExtraArgsOk() (*string, bool) {
	if o == nil || o.ExtraArgs == nil {
		return nil, false
	}
	return o.ExtraArgs, true
}

// HasExtraArgs returns a boolean if a field has been set.
func (o *Sysbench) HasExtraArgs() bool {
	return o != nil && o.ExtraArgs != nil
}

// SetExtraArgs gets a reference to the given string and assigns it to the ExtraArgs field.
func (o *Sysbench) SetExtraArgs(v string) {
	o.ExtraArgs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Sysbench) MarshalJSON() ([]byte, error) {
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
	toSerialize["cluster"] = o.Cluster
	toSerialize["database"] = o.Database
	if o.Threads.IsSet() {
		toSerialize["threads"] = o.Threads.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.TableSize.IsSet() {
		toSerialize["tableSize"] = o.TableSize.Get()
	}
	if o.TableNum.IsSet() {
		toSerialize["tableNum"] = o.TableNum.Get()
	}
	if o.TestType != nil {
		toSerialize["testType"] = o.TestType
	}
	if o.ReadPercent.IsSet() {
		toSerialize["readPercent"] = o.ReadPercent.Get()
	}
	if o.WritePercent.IsSet() {
		toSerialize["writePercent"] = o.WritePercent.Get()
	}
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	toSerialize["address"] = o.Address
	if o.ExtraArgs != nil {
		toSerialize["extraArgs"] = o.ExtraArgs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Sysbench) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Step          *SysbenchStep        `json:"step,omitempty"`
		LimitCpu      *string              `json:"limitCpu,omitempty"`
		LimitMemory   *string              `json:"limitMemory,omitempty"`
		RequestCpu    *string              `json:"requestCpu,omitempty"`
		RequestMemory *string              `json:"requestMemory,omitempty"`
		Name          *string              `json:"name,omitempty"`
		Cluster       *string              `json:"cluster"`
		Database      *string              `json:"database"`
		Threads       common.NullableInt32 `json:"threads,omitempty"`
		Duration      common.NullableInt32 `json:"duration,omitempty"`
		TableSize     common.NullableInt32 `json:"tableSize,omitempty"`
		TableNum      common.NullableInt32 `json:"tableNum,omitempty"`
		TestType      *SysbenchTestType    `json:"testType,omitempty"`
		ReadPercent   common.NullableInt32 `json:"readPercent,omitempty"`
		WritePercent  common.NullableInt32 `json:"writePercent,omitempty"`
		Username      *string              `json:"username"`
		Password      *string              `json:"password"`
		Address       *string              `json:"address"`
		ExtraArgs     *string              `json:"extraArgs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Cluster == nil {
		return fmt.Errorf("required field cluster missing")
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
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
		common.DeleteKeys(additionalProperties, &[]string{"step", "limitCpu", "limitMemory", "requestCpu", "requestMemory", "name", "cluster", "database", "threads", "duration", "tableSize", "tableNum", "testType", "readPercent", "writePercent", "username", "password", "address", "extraArgs"})
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
	o.Cluster = *all.Cluster
	o.Database = *all.Database
	o.Threads = all.Threads
	o.Duration = all.Duration
	o.TableSize = all.TableSize
	o.TableNum = all.TableNum
	if all.TestType != nil && !all.TestType.IsValid() {
		hasInvalidField = true
	} else {
		o.TestType = all.TestType
	}
	o.ReadPercent = all.ReadPercent
	o.WritePercent = all.WritePercent
	o.Username = *all.Username
	o.Password = *all.Password
	o.Address = *all.Address
	o.ExtraArgs = all.ExtraArgs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
