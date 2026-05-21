// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Esrally esrally is the Elasticsearch Rally benchmark object
type Esrally struct {
	// Step of benchmark
	Step *BenchmarkStep `json:"step,omitempty"`
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
	// Elasticsearch index name
	Database *string `json:"database,omitempty"`
	// Username for Elasticsearch
	Username string `json:"username"`
	// Password for Elasticsearch
	Password *string `json:"password,omitempty"`
	// Address for Elasticsearch
	Address string `json:"address"`
	// Elasticsearch target version hint for generated workload compatibility
	TargetVersion *string `json:"targetVersion,omitempty"`
	// Rally request error handling behavior
	OnError *EsrallyOnError `json:"onError,omitempty"`
	// Rally telemetry devices
	Telemetry []EsrallyTelemetry `json:"telemetry,omitempty"`
	// Generated Elasticsearch dataset shape
	DataProfile *EsrallyDataProfile `json:"dataProfile,omitempty"`
	// Number of generated documents
	DocumentCount common.NullableInt32 `json:"documentCount,omitempty"`
	// Generated Rally workload profile
	Workload *EsrallyWorkload `json:"workload,omitempty"`
	// Extra arguments for esrally
	ExtraArgs *string `json:"extraArgs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEsrally instantiates a new Esrally object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEsrally(cluster string, username string, address string) *Esrally {
	this := Esrally{}
	var step BenchmarkStep = BenchmarkStepAll
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
	var database string = "kubebench"
	this.Database = &database
	this.Username = username
	this.Address = address
	var onError EsrallyOnError = EsrallyOnErrorAbort
	this.OnError = &onError
	var dataProfile EsrallyDataProfile = EsrallyDataProfileLogs
	this.DataProfile = &dataProfile
	var documentCount int32 = 10000
	this.DocumentCount = *common.NewNullableInt32(&documentCount)
	var workload EsrallyWorkload = EsrallyWorkloadAll
	this.Workload = &workload
	return &this
}

// NewEsrallyWithDefaults instantiates a new Esrally object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEsrallyWithDefaults() *Esrally {
	this := Esrally{}
	var step BenchmarkStep = BenchmarkStepAll
	this.Step = &step
	var limitCpu string = "1"
	this.LimitCpu = &limitCpu
	var limitMemory string = "2Gi"
	this.LimitMemory = &limitMemory
	var requestCpu string = "0.25"
	this.RequestCpu = &requestCpu
	var requestMemory string = "0.5Gi"
	this.RequestMemory = &requestMemory
	var database string = "kubebench"
	this.Database = &database
	var onError EsrallyOnError = EsrallyOnErrorAbort
	this.OnError = &onError
	var dataProfile EsrallyDataProfile = EsrallyDataProfileLogs
	this.DataProfile = &dataProfile
	var documentCount int32 = 10000
	this.DocumentCount = *common.NewNullableInt32(&documentCount)
	var workload EsrallyWorkload = EsrallyWorkloadAll
	this.Workload = &workload
	return &this
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *Esrally) GetStep() BenchmarkStep {
	if o == nil || o.Step == nil {
		var ret BenchmarkStep
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Esrally) GetStepOk() (*BenchmarkStep, bool) {
	if o == nil || o.Step == nil {
		return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *Esrally) HasStep() bool {
	return o != nil && o.Step != nil
}

// SetStep gets a reference to the given BenchmarkStep and assigns it to the Step field.
func (o *Esrally) SetStep(v BenchmarkStep) {
	o.Step = &v
}

// GetLimitCpu returns the LimitCpu field value if set, zero value otherwise.
func (o *Esrally) GetLimitCpu() string {
	if o == nil || o.LimitCpu == nil {
		var ret string
		return ret
	}
	return *o.LimitCpu
}

// GetLimitCpuOk returns a tuple with the LimitCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Esrally) GetLimitCpuOk() (*string, bool) {
	if o == nil || o.LimitCpu == nil {
		return nil, false
	}
	return o.LimitCpu, true
}

// HasLimitCpu returns a boolean if a field has been set.
func (o *Esrally) HasLimitCpu() bool {
	return o != nil && o.LimitCpu != nil
}

// SetLimitCpu gets a reference to the given string and assigns it to the LimitCpu field.
func (o *Esrally) SetLimitCpu(v string) {
	o.LimitCpu = &v
}

// GetLimitMemory returns the LimitMemory field value if set, zero value otherwise.
func (o *Esrally) GetLimitMemory() string {
	if o == nil || o.LimitMemory == nil {
		var ret string
		return ret
	}
	return *o.LimitMemory
}

// GetLimitMemoryOk returns a tuple with the LimitMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Esrally) GetLimitMemoryOk() (*string, bool) {
	if o == nil || o.LimitMemory == nil {
		return nil, false
	}
	return o.LimitMemory, true
}

// HasLimitMemory returns a boolean if a field has been set.
func (o *Esrally) HasLimitMemory() bool {
	return o != nil && o.LimitMemory != nil
}

// SetLimitMemory gets a reference to the given string and assigns it to the LimitMemory field.
func (o *Esrally) SetLimitMemory(v string) {
	o.LimitMemory = &v
}

// GetRequestCpu returns the RequestCpu field value if set, zero value otherwise.
func (o *Esrally) GetRequestCpu() string {
	if o == nil || o.RequestCpu == nil {
		var ret string
		return ret
	}
	return *o.RequestCpu
}

// GetRequestCpuOk returns a tuple with the RequestCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Esrally) GetRequestCpuOk() (*string, bool) {
	if o == nil || o.RequestCpu == nil {
		return nil, false
	}
	return o.RequestCpu, true
}

// HasRequestCpu returns a boolean if a field has been set.
func (o *Esrally) HasRequestCpu() bool {
	return o != nil && o.RequestCpu != nil
}

// SetRequestCpu gets a reference to the given string and assigns it to the RequestCpu field.
func (o *Esrally) SetRequestCpu(v string) {
	o.RequestCpu = &v
}

// GetRequestMemory returns the RequestMemory field value if set, zero value otherwise.
func (o *Esrally) GetRequestMemory() string {
	if o == nil || o.RequestMemory == nil {
		var ret string
		return ret
	}
	return *o.RequestMemory
}

// GetRequestMemoryOk returns a tuple with the RequestMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Esrally) GetRequestMemoryOk() (*string, bool) {
	if o == nil || o.RequestMemory == nil {
		return nil, false
	}
	return o.RequestMemory, true
}

// HasRequestMemory returns a boolean if a field has been set.
func (o *Esrally) HasRequestMemory() bool {
	return o != nil && o.RequestMemory != nil
}

// SetRequestMemory gets a reference to the given string and assigns it to the RequestMemory field.
func (o *Esrally) SetRequestMemory(v string) {
	o.RequestMemory = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Esrally) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Esrally) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Esrally) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Esrally) SetName(v string) {
	o.Name = &v
}

// GetCluster returns the Cluster field value.
func (o *Esrally) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *Esrally) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value.
func (o *Esrally) SetCluster(v string) {
	o.Cluster = v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *Esrally) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Esrally) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *Esrally) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *Esrally) SetDatabase(v string) {
	o.Database = &v
}

// GetUsername returns the Username field value.
func (o *Esrally) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Esrally) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value.
func (o *Esrally) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *Esrally) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Esrally) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *Esrally) HasPassword() bool {
	return o != nil && o.Password != nil
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *Esrally) SetPassword(v string) {
	o.Password = &v
}

// GetAddress returns the Address field value.
func (o *Esrally) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *Esrally) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value.
func (o *Esrally) SetAddress(v string) {
	o.Address = v
}

// GetTargetVersion returns the TargetVersion field value if set, zero value otherwise.
func (o *Esrally) GetTargetVersion() string {
	if o == nil || o.TargetVersion == nil {
		var ret string
		return ret
	}
	return *o.TargetVersion
}

// GetTargetVersionOk returns a tuple with the TargetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Esrally) GetTargetVersionOk() (*string, bool) {
	if o == nil || o.TargetVersion == nil {
		return nil, false
	}
	return o.TargetVersion, true
}

// HasTargetVersion returns a boolean if a field has been set.
func (o *Esrally) HasTargetVersion() bool {
	return o != nil && o.TargetVersion != nil
}

// SetTargetVersion gets a reference to the given string and assigns it to the TargetVersion field.
func (o *Esrally) SetTargetVersion(v string) {
	o.TargetVersion = &v
}

// GetOnError returns the OnError field value if set, zero value otherwise.
func (o *Esrally) GetOnError() EsrallyOnError {
	if o == nil || o.OnError == nil {
		var ret EsrallyOnError
		return ret
	}
	return *o.OnError
}

// GetOnErrorOk returns a tuple with the OnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Esrally) GetOnErrorOk() (*EsrallyOnError, bool) {
	if o == nil || o.OnError == nil {
		return nil, false
	}
	return o.OnError, true
}

// HasOnError returns a boolean if a field has been set.
func (o *Esrally) HasOnError() bool {
	return o != nil && o.OnError != nil
}

// SetOnError gets a reference to the given EsrallyOnError and assigns it to the OnError field.
func (o *Esrally) SetOnError(v EsrallyOnError) {
	o.OnError = &v
}

// GetTelemetry returns the Telemetry field value if set, zero value otherwise.
func (o *Esrally) GetTelemetry() []EsrallyTelemetry {
	if o == nil || o.Telemetry == nil {
		var ret []EsrallyTelemetry
		return ret
	}
	return o.Telemetry
}

// GetTelemetryOk returns a tuple with the Telemetry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Esrally) GetTelemetryOk() (*[]EsrallyTelemetry, bool) {
	if o == nil || o.Telemetry == nil {
		return nil, false
	}
	return &o.Telemetry, true
}

// HasTelemetry returns a boolean if a field has been set.
func (o *Esrally) HasTelemetry() bool {
	return o != nil && o.Telemetry != nil
}

// SetTelemetry gets a reference to the given []EsrallyTelemetry and assigns it to the Telemetry field.
func (o *Esrally) SetTelemetry(v []EsrallyTelemetry) {
	o.Telemetry = v
}

// GetDataProfile returns the DataProfile field value if set, zero value otherwise.
func (o *Esrally) GetDataProfile() EsrallyDataProfile {
	if o == nil || o.DataProfile == nil {
		var ret EsrallyDataProfile
		return ret
	}
	return *o.DataProfile
}

// GetDataProfileOk returns a tuple with the DataProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Esrally) GetDataProfileOk() (*EsrallyDataProfile, bool) {
	if o == nil || o.DataProfile == nil {
		return nil, false
	}
	return o.DataProfile, true
}

// HasDataProfile returns a boolean if a field has been set.
func (o *Esrally) HasDataProfile() bool {
	return o != nil && o.DataProfile != nil
}

// SetDataProfile gets a reference to the given EsrallyDataProfile and assigns it to the DataProfile field.
func (o *Esrally) SetDataProfile(v EsrallyDataProfile) {
	o.DataProfile = &v
}

// GetDocumentCount returns the DocumentCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Esrally) GetDocumentCount() int32 {
	if o == nil || o.DocumentCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DocumentCount.Get()
}

// GetDocumentCountOk returns a tuple with the DocumentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Esrally) GetDocumentCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentCount.Get(), o.DocumentCount.IsSet()
}

// HasDocumentCount returns a boolean if a field has been set.
func (o *Esrally) HasDocumentCount() bool {
	return o != nil && o.DocumentCount.IsSet()
}

// SetDocumentCount gets a reference to the given common.NullableInt32 and assigns it to the DocumentCount field.
func (o *Esrally) SetDocumentCount(v int32) {
	o.DocumentCount.Set(&v)
}

// SetDocumentCountNil sets the value for DocumentCount to be an explicit nil.
func (o *Esrally) SetDocumentCountNil() {
	o.DocumentCount.Set(nil)
}

// UnsetDocumentCount ensures that no value is present for DocumentCount, not even an explicit nil.
func (o *Esrally) UnsetDocumentCount() {
	o.DocumentCount.Unset()
}

// GetWorkload returns the Workload field value if set, zero value otherwise.
func (o *Esrally) GetWorkload() EsrallyWorkload {
	if o == nil || o.Workload == nil {
		var ret EsrallyWorkload
		return ret
	}
	return *o.Workload
}

// GetWorkloadOk returns a tuple with the Workload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Esrally) GetWorkloadOk() (*EsrallyWorkload, bool) {
	if o == nil || o.Workload == nil {
		return nil, false
	}
	return o.Workload, true
}

// HasWorkload returns a boolean if a field has been set.
func (o *Esrally) HasWorkload() bool {
	return o != nil && o.Workload != nil
}

// SetWorkload gets a reference to the given EsrallyWorkload and assigns it to the Workload field.
func (o *Esrally) SetWorkload(v EsrallyWorkload) {
	o.Workload = &v
}

// GetExtraArgs returns the ExtraArgs field value if set, zero value otherwise.
func (o *Esrally) GetExtraArgs() string {
	if o == nil || o.ExtraArgs == nil {
		var ret string
		return ret
	}
	return *o.ExtraArgs
}

// GetExtraArgsOk returns a tuple with the ExtraArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Esrally) GetExtraArgsOk() (*string, bool) {
	if o == nil || o.ExtraArgs == nil {
		return nil, false
	}
	return o.ExtraArgs, true
}

// HasExtraArgs returns a boolean if a field has been set.
func (o *Esrally) HasExtraArgs() bool {
	return o != nil && o.ExtraArgs != nil
}

// SetExtraArgs gets a reference to the given string and assigns it to the ExtraArgs field.
func (o *Esrally) SetExtraArgs(v string) {
	o.ExtraArgs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Esrally) MarshalJSON() ([]byte, error) {
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
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	toSerialize["username"] = o.Username
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	toSerialize["address"] = o.Address
	if o.TargetVersion != nil {
		toSerialize["targetVersion"] = o.TargetVersion
	}
	if o.OnError != nil {
		toSerialize["onError"] = o.OnError
	}
	if o.Telemetry != nil {
		toSerialize["telemetry"] = o.Telemetry
	}
	if o.DataProfile != nil {
		toSerialize["dataProfile"] = o.DataProfile
	}
	if o.DocumentCount.IsSet() {
		toSerialize["documentCount"] = o.DocumentCount.Get()
	}
	if o.Workload != nil {
		toSerialize["workload"] = o.Workload
	}
	if o.ExtraArgs != nil {
		toSerialize["extraArgs"] = o.ExtraArgs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Esrally) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Step          *BenchmarkStep       `json:"step,omitempty"`
		LimitCpu      *string              `json:"limitCpu,omitempty"`
		LimitMemory   *string              `json:"limitMemory,omitempty"`
		RequestCpu    *string              `json:"requestCpu,omitempty"`
		RequestMemory *string              `json:"requestMemory,omitempty"`
		Name          *string              `json:"name,omitempty"`
		Cluster       *string              `json:"cluster"`
		Database      *string              `json:"database,omitempty"`
		Username      *string              `json:"username"`
		Password      *string              `json:"password,omitempty"`
		Address       *string              `json:"address"`
		TargetVersion *string              `json:"targetVersion,omitempty"`
		OnError       *EsrallyOnError      `json:"onError,omitempty"`
		Telemetry     []EsrallyTelemetry   `json:"telemetry,omitempty"`
		DataProfile   *EsrallyDataProfile  `json:"dataProfile,omitempty"`
		DocumentCount common.NullableInt32 `json:"documentCount,omitempty"`
		Workload      *EsrallyWorkload     `json:"workload,omitempty"`
		ExtraArgs     *string              `json:"extraArgs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Cluster == nil {
		return fmt.Errorf("required field cluster missing")
	}
	if all.Username == nil {
		return fmt.Errorf("required field username missing")
	}
	if all.Address == nil {
		return fmt.Errorf("required field address missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"step", "limitCpu", "limitMemory", "requestCpu", "requestMemory", "name", "cluster", "database", "username", "password", "address", "targetVersion", "onError", "telemetry", "dataProfile", "documentCount", "workload", "extraArgs"})
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
	o.Database = all.Database
	o.Username = *all.Username
	o.Password = all.Password
	o.Address = *all.Address
	o.TargetVersion = all.TargetVersion
	if all.OnError != nil && !all.OnError.IsValid() {
		hasInvalidField = true
	} else {
		o.OnError = all.OnError
	}
	o.Telemetry = all.Telemetry
	if all.DataProfile != nil && !all.DataProfile.IsValid() {
		hasInvalidField = true
	} else {
		o.DataProfile = all.DataProfile
	}
	o.DocumentCount = all.DocumentCount
	if all.Workload != nil && !all.Workload.IsValid() {
		hasInvalidField = true
	} else {
		o.Workload = all.Workload
	}
	o.ExtraArgs = all.ExtraArgs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
