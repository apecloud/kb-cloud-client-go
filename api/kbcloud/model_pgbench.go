// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Pgbench pgbench is the pgbench benchmark object
type Pgbench struct {
	// Step of pgbench
	Step *PgbenchStep `json:"step,omitempty"`
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
	// Username for database
	Username string `json:"username"`
	// Password for database
	Password string `json:"password"`
	// Address for database
	Address string `json:"address"`
	// Scale of pgbench
	Scale *int32 `json:"scale,omitempty"`
	// Number of clients to run
	Clients *int32 `json:"clients,omitempty"`
	// Number of threads to use
	Threads *int32 `json:"threads,omitempty"`
	// the seconds of test duration
	Duration *int32 `json:"duration,omitempty"`
	// Run select only test
	SelectOnly *bool `json:"selectOnly,omitempty"`
	// Extra arguments for pgbench
	ExtraArgs *string `json:"extraArgs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPgbench instantiates a new Pgbench object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPgbench(cluster string, database string, username string, password string, address string) *Pgbench {
	this := Pgbench{}
	var step PgbenchStep = PgbenchStepAll
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
	this.Username = username
	this.Password = password
	this.Address = address
	var selectOnly bool = false
	this.SelectOnly = &selectOnly
	return &this
}

// NewPgbenchWithDefaults instantiates a new Pgbench object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPgbenchWithDefaults() *Pgbench {
	this := Pgbench{}
	var step PgbenchStep = PgbenchStepAll
	this.Step = &step
	var limitCpu string = "1"
	this.LimitCpu = &limitCpu
	var limitMemory string = "2Gi"
	this.LimitMemory = &limitMemory
	var requestCpu string = "0.25"
	this.RequestCpu = &requestCpu
	var requestMemory string = "0.5Gi"
	this.RequestMemory = &requestMemory
	var selectOnly bool = false
	this.SelectOnly = &selectOnly
	return &this
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *Pgbench) GetStep() PgbenchStep {
	if o == nil || o.Step == nil {
		var ret PgbenchStep
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pgbench) GetStepOk() (*PgbenchStep, bool) {
	if o == nil || o.Step == nil {
		return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *Pgbench) HasStep() bool {
	return o != nil && o.Step != nil
}

// SetStep gets a reference to the given PgbenchStep and assigns it to the Step field.
func (o *Pgbench) SetStep(v PgbenchStep) {
	o.Step = &v
}

// GetLimitCpu returns the LimitCpu field value if set, zero value otherwise.
func (o *Pgbench) GetLimitCpu() string {
	if o == nil || o.LimitCpu == nil {
		var ret string
		return ret
	}
	return *o.LimitCpu
}

// GetLimitCpuOk returns a tuple with the LimitCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pgbench) GetLimitCpuOk() (*string, bool) {
	if o == nil || o.LimitCpu == nil {
		return nil, false
	}
	return o.LimitCpu, true
}

// HasLimitCpu returns a boolean if a field has been set.
func (o *Pgbench) HasLimitCpu() bool {
	return o != nil && o.LimitCpu != nil
}

// SetLimitCpu gets a reference to the given string and assigns it to the LimitCpu field.
func (o *Pgbench) SetLimitCpu(v string) {
	o.LimitCpu = &v
}

// GetLimitMemory returns the LimitMemory field value if set, zero value otherwise.
func (o *Pgbench) GetLimitMemory() string {
	if o == nil || o.LimitMemory == nil {
		var ret string
		return ret
	}
	return *o.LimitMemory
}

// GetLimitMemoryOk returns a tuple with the LimitMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pgbench) GetLimitMemoryOk() (*string, bool) {
	if o == nil || o.LimitMemory == nil {
		return nil, false
	}
	return o.LimitMemory, true
}

// HasLimitMemory returns a boolean if a field has been set.
func (o *Pgbench) HasLimitMemory() bool {
	return o != nil && o.LimitMemory != nil
}

// SetLimitMemory gets a reference to the given string and assigns it to the LimitMemory field.
func (o *Pgbench) SetLimitMemory(v string) {
	o.LimitMemory = &v
}

// GetRequestCpu returns the RequestCpu field value if set, zero value otherwise.
func (o *Pgbench) GetRequestCpu() string {
	if o == nil || o.RequestCpu == nil {
		var ret string
		return ret
	}
	return *o.RequestCpu
}

// GetRequestCpuOk returns a tuple with the RequestCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pgbench) GetRequestCpuOk() (*string, bool) {
	if o == nil || o.RequestCpu == nil {
		return nil, false
	}
	return o.RequestCpu, true
}

// HasRequestCpu returns a boolean if a field has been set.
func (o *Pgbench) HasRequestCpu() bool {
	return o != nil && o.RequestCpu != nil
}

// SetRequestCpu gets a reference to the given string and assigns it to the RequestCpu field.
func (o *Pgbench) SetRequestCpu(v string) {
	o.RequestCpu = &v
}

// GetRequestMemory returns the RequestMemory field value if set, zero value otherwise.
func (o *Pgbench) GetRequestMemory() string {
	if o == nil || o.RequestMemory == nil {
		var ret string
		return ret
	}
	return *o.RequestMemory
}

// GetRequestMemoryOk returns a tuple with the RequestMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pgbench) GetRequestMemoryOk() (*string, bool) {
	if o == nil || o.RequestMemory == nil {
		return nil, false
	}
	return o.RequestMemory, true
}

// HasRequestMemory returns a boolean if a field has been set.
func (o *Pgbench) HasRequestMemory() bool {
	return o != nil && o.RequestMemory != nil
}

// SetRequestMemory gets a reference to the given string and assigns it to the RequestMemory field.
func (o *Pgbench) SetRequestMemory(v string) {
	o.RequestMemory = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Pgbench) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pgbench) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Pgbench) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Pgbench) SetName(v string) {
	o.Name = &v
}

// GetCluster returns the Cluster field value.
func (o *Pgbench) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *Pgbench) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value.
func (o *Pgbench) SetCluster(v string) {
	o.Cluster = v
}

// GetDatabase returns the Database field value.
func (o *Pgbench) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *Pgbench) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *Pgbench) SetDatabase(v string) {
	o.Database = v
}

// GetUsername returns the Username field value.
func (o *Pgbench) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Pgbench) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value.
func (o *Pgbench) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value.
func (o *Pgbench) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *Pgbench) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *Pgbench) SetPassword(v string) {
	o.Password = v
}

// GetAddress returns the Address field value.
func (o *Pgbench) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *Pgbench) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value.
func (o *Pgbench) SetAddress(v string) {
	o.Address = v
}

// GetScale returns the Scale field value if set, zero value otherwise.
func (o *Pgbench) GetScale() int32 {
	if o == nil || o.Scale == nil {
		var ret int32
		return ret
	}
	return *o.Scale
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pgbench) GetScaleOk() (*int32, bool) {
	if o == nil || o.Scale == nil {
		return nil, false
	}
	return o.Scale, true
}

// HasScale returns a boolean if a field has been set.
func (o *Pgbench) HasScale() bool {
	return o != nil && o.Scale != nil
}

// SetScale gets a reference to the given int32 and assigns it to the Scale field.
func (o *Pgbench) SetScale(v int32) {
	o.Scale = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *Pgbench) GetClients() int32 {
	if o == nil || o.Clients == nil {
		var ret int32
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pgbench) GetClientsOk() (*int32, bool) {
	if o == nil || o.Clients == nil {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *Pgbench) HasClients() bool {
	return o != nil && o.Clients != nil
}

// SetClients gets a reference to the given int32 and assigns it to the Clients field.
func (o *Pgbench) SetClients(v int32) {
	o.Clients = &v
}

// GetThreads returns the Threads field value if set, zero value otherwise.
func (o *Pgbench) GetThreads() int32 {
	if o == nil || o.Threads == nil {
		var ret int32
		return ret
	}
	return *o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pgbench) GetThreadsOk() (*int32, bool) {
	if o == nil || o.Threads == nil {
		return nil, false
	}
	return o.Threads, true
}

// HasThreads returns a boolean if a field has been set.
func (o *Pgbench) HasThreads() bool {
	return o != nil && o.Threads != nil
}

// SetThreads gets a reference to the given int32 and assigns it to the Threads field.
func (o *Pgbench) SetThreads(v int32) {
	o.Threads = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Pgbench) GetDuration() int32 {
	if o == nil || o.Duration == nil {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pgbench) GetDurationOk() (*int32, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Pgbench) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *Pgbench) SetDuration(v int32) {
	o.Duration = &v
}

// GetSelectOnly returns the SelectOnly field value if set, zero value otherwise.
func (o *Pgbench) GetSelectOnly() bool {
	if o == nil || o.SelectOnly == nil {
		var ret bool
		return ret
	}
	return *o.SelectOnly
}

// GetSelectOnlyOk returns a tuple with the SelectOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pgbench) GetSelectOnlyOk() (*bool, bool) {
	if o == nil || o.SelectOnly == nil {
		return nil, false
	}
	return o.SelectOnly, true
}

// HasSelectOnly returns a boolean if a field has been set.
func (o *Pgbench) HasSelectOnly() bool {
	return o != nil && o.SelectOnly != nil
}

// SetSelectOnly gets a reference to the given bool and assigns it to the SelectOnly field.
func (o *Pgbench) SetSelectOnly(v bool) {
	o.SelectOnly = &v
}

// GetExtraArgs returns the ExtraArgs field value if set, zero value otherwise.
func (o *Pgbench) GetExtraArgs() string {
	if o == nil || o.ExtraArgs == nil {
		var ret string
		return ret
	}
	return *o.ExtraArgs
}

// GetExtraArgsOk returns a tuple with the ExtraArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pgbench) GetExtraArgsOk() (*string, bool) {
	if o == nil || o.ExtraArgs == nil {
		return nil, false
	}
	return o.ExtraArgs, true
}

// HasExtraArgs returns a boolean if a field has been set.
func (o *Pgbench) HasExtraArgs() bool {
	return o != nil && o.ExtraArgs != nil
}

// SetExtraArgs gets a reference to the given string and assigns it to the ExtraArgs field.
func (o *Pgbench) SetExtraArgs(v string) {
	o.ExtraArgs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Pgbench) MarshalJSON() ([]byte, error) {
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
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	toSerialize["address"] = o.Address
	if o.Scale != nil {
		toSerialize["scale"] = o.Scale
	}
	if o.Clients != nil {
		toSerialize["clients"] = o.Clients
	}
	if o.Threads != nil {
		toSerialize["threads"] = o.Threads
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.SelectOnly != nil {
		toSerialize["selectOnly"] = o.SelectOnly
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
func (o *Pgbench) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Step          *PgbenchStep `json:"step,omitempty"`
		LimitCpu      *string      `json:"limitCpu,omitempty"`
		LimitMemory   *string      `json:"limitMemory,omitempty"`
		RequestCpu    *string      `json:"requestCpu,omitempty"`
		RequestMemory *string      `json:"requestMemory,omitempty"`
		Name          *string      `json:"name,omitempty"`
		Cluster       *string      `json:"cluster"`
		Database      *string      `json:"database"`
		Username      *string      `json:"username"`
		Password      *string      `json:"password"`
		Address       *string      `json:"address"`
		Scale         *int32       `json:"scale,omitempty"`
		Clients       *int32       `json:"clients,omitempty"`
		Threads       *int32       `json:"threads,omitempty"`
		Duration      *int32       `json:"duration,omitempty"`
		SelectOnly    *bool        `json:"selectOnly,omitempty"`
		ExtraArgs     *string      `json:"extraArgs,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"step", "limitCpu", "limitMemory", "requestCpu", "requestMemory", "name", "cluster", "database", "username", "password", "address", "scale", "clients", "threads", "duration", "selectOnly", "extraArgs"})
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
	o.Username = *all.Username
	o.Password = *all.Password
	o.Address = *all.Address
	o.Scale = all.Scale
	o.Clients = all.Clients
	o.Threads = all.Threads
	o.Duration = all.Duration
	o.SelectOnly = all.SelectOnly
	o.ExtraArgs = all.ExtraArgs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
