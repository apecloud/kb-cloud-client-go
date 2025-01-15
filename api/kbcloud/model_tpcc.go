// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Tpcc tpcc is the tpcc benchmark object
type Tpcc struct {
	// Step of sysbench
	Step *TpccStep `json:"step,omitempty"`
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
	// Number of threads to use
	Threads common.NullableInt32 `json:"threads,omitempty"`
	// Number of warehouses
	Warehouses common.NullableInt32 `json:"warehouses,omitempty"`
	// the minutes of test duration
	Duration common.NullableInt32 `json:"duration,omitempty"`
	// limit the number of transactions per minute, 0 means no limit
	LimitTxPerMin *int32 `json:"limitTxPerMin,omitempty"`
	// Percentage of new order transactions
	NewOrderWeight common.NullableInt32 `json:"newOrderWeight,omitempty"`
	// Percentage of payment transactions
	PaymentWeight common.NullableInt32 `json:"paymentWeight,omitempty"`
	// Percentage of order status transactions
	OrderStatusWeight common.NullableInt32 `json:"orderStatusWeight,omitempty"`
	// Percentage of delivery transactions
	DeliveryWeight common.NullableInt32 `json:"deliveryWeight,omitempty"`
	// Percentage of stock level transactions
	StockLevelWeight common.NullableInt32 `json:"stockLevelWeight,omitempty"`
	// Extra arguments for tpcc
	ExtraArgs *string `json:"extraArgs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTpcc instantiates a new Tpcc object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTpcc(cluster string, database string, username string, password string, address string) *Tpcc {
	this := Tpcc{}
	var step TpccStep = TpccStepAll
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
	return &this
}

// NewTpccWithDefaults instantiates a new Tpcc object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTpccWithDefaults() *Tpcc {
	this := Tpcc{}
	var step TpccStep = TpccStepAll
	this.Step = &step
	var limitCpu string = "1"
	this.LimitCpu = &limitCpu
	var limitMemory string = "2Gi"
	this.LimitMemory = &limitMemory
	var requestCpu string = "0.25"
	this.RequestCpu = &requestCpu
	var requestMemory string = "0.5Gi"
	this.RequestMemory = &requestMemory
	return &this
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *Tpcc) GetStep() TpccStep {
	if o == nil || o.Step == nil {
		var ret TpccStep
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tpcc) GetStepOk() (*TpccStep, bool) {
	if o == nil || o.Step == nil {
		return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *Tpcc) HasStep() bool {
	return o != nil && o.Step != nil
}

// SetStep gets a reference to the given TpccStep and assigns it to the Step field.
func (o *Tpcc) SetStep(v TpccStep) {
	o.Step = &v
}

// GetLimitCpu returns the LimitCpu field value if set, zero value otherwise.
func (o *Tpcc) GetLimitCpu() string {
	if o == nil || o.LimitCpu == nil {
		var ret string
		return ret
	}
	return *o.LimitCpu
}

// GetLimitCpuOk returns a tuple with the LimitCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tpcc) GetLimitCpuOk() (*string, bool) {
	if o == nil || o.LimitCpu == nil {
		return nil, false
	}
	return o.LimitCpu, true
}

// HasLimitCpu returns a boolean if a field has been set.
func (o *Tpcc) HasLimitCpu() bool {
	return o != nil && o.LimitCpu != nil
}

// SetLimitCpu gets a reference to the given string and assigns it to the LimitCpu field.
func (o *Tpcc) SetLimitCpu(v string) {
	o.LimitCpu = &v
}

// GetLimitMemory returns the LimitMemory field value if set, zero value otherwise.
func (o *Tpcc) GetLimitMemory() string {
	if o == nil || o.LimitMemory == nil {
		var ret string
		return ret
	}
	return *o.LimitMemory
}

// GetLimitMemoryOk returns a tuple with the LimitMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tpcc) GetLimitMemoryOk() (*string, bool) {
	if o == nil || o.LimitMemory == nil {
		return nil, false
	}
	return o.LimitMemory, true
}

// HasLimitMemory returns a boolean if a field has been set.
func (o *Tpcc) HasLimitMemory() bool {
	return o != nil && o.LimitMemory != nil
}

// SetLimitMemory gets a reference to the given string and assigns it to the LimitMemory field.
func (o *Tpcc) SetLimitMemory(v string) {
	o.LimitMemory = &v
}

// GetRequestCpu returns the RequestCpu field value if set, zero value otherwise.
func (o *Tpcc) GetRequestCpu() string {
	if o == nil || o.RequestCpu == nil {
		var ret string
		return ret
	}
	return *o.RequestCpu
}

// GetRequestCpuOk returns a tuple with the RequestCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tpcc) GetRequestCpuOk() (*string, bool) {
	if o == nil || o.RequestCpu == nil {
		return nil, false
	}
	return o.RequestCpu, true
}

// HasRequestCpu returns a boolean if a field has been set.
func (o *Tpcc) HasRequestCpu() bool {
	return o != nil && o.RequestCpu != nil
}

// SetRequestCpu gets a reference to the given string and assigns it to the RequestCpu field.
func (o *Tpcc) SetRequestCpu(v string) {
	o.RequestCpu = &v
}

// GetRequestMemory returns the RequestMemory field value if set, zero value otherwise.
func (o *Tpcc) GetRequestMemory() string {
	if o == nil || o.RequestMemory == nil {
		var ret string
		return ret
	}
	return *o.RequestMemory
}

// GetRequestMemoryOk returns a tuple with the RequestMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tpcc) GetRequestMemoryOk() (*string, bool) {
	if o == nil || o.RequestMemory == nil {
		return nil, false
	}
	return o.RequestMemory, true
}

// HasRequestMemory returns a boolean if a field has been set.
func (o *Tpcc) HasRequestMemory() bool {
	return o != nil && o.RequestMemory != nil
}

// SetRequestMemory gets a reference to the given string and assigns it to the RequestMemory field.
func (o *Tpcc) SetRequestMemory(v string) {
	o.RequestMemory = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Tpcc) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tpcc) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Tpcc) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Tpcc) SetName(v string) {
	o.Name = &v
}

// GetCluster returns the Cluster field value.
func (o *Tpcc) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *Tpcc) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value.
func (o *Tpcc) SetCluster(v string) {
	o.Cluster = v
}

// GetDatabase returns the Database field value.
func (o *Tpcc) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *Tpcc) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *Tpcc) SetDatabase(v string) {
	o.Database = v
}

// GetUsername returns the Username field value.
func (o *Tpcc) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Tpcc) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value.
func (o *Tpcc) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value.
func (o *Tpcc) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *Tpcc) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *Tpcc) SetPassword(v string) {
	o.Password = v
}

// GetAddress returns the Address field value.
func (o *Tpcc) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *Tpcc) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value.
func (o *Tpcc) SetAddress(v string) {
	o.Address = v
}

// GetThreads returns the Threads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tpcc) GetThreads() int32 {
	if o == nil || o.Threads.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Threads.Get()
}

// GetThreadsOk returns a tuple with the Threads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Tpcc) GetThreadsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Threads.Get(), o.Threads.IsSet()
}

// HasThreads returns a boolean if a field has been set.
func (o *Tpcc) HasThreads() bool {
	return o != nil && o.Threads.IsSet()
}

// SetThreads gets a reference to the given common.NullableInt32 and assigns it to the Threads field.
func (o *Tpcc) SetThreads(v int32) {
	o.Threads.Set(&v)
}

// SetThreadsNil sets the value for Threads to be an explicit nil.
func (o *Tpcc) SetThreadsNil() {
	o.Threads.Set(nil)
}

// UnsetThreads ensures that no value is present for Threads, not even an explicit nil.
func (o *Tpcc) UnsetThreads() {
	o.Threads.Unset()
}

// GetWarehouses returns the Warehouses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tpcc) GetWarehouses() int32 {
	if o == nil || o.Warehouses.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Warehouses.Get()
}

// GetWarehousesOk returns a tuple with the Warehouses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Tpcc) GetWarehousesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warehouses.Get(), o.Warehouses.IsSet()
}

// HasWarehouses returns a boolean if a field has been set.
func (o *Tpcc) HasWarehouses() bool {
	return o != nil && o.Warehouses.IsSet()
}

// SetWarehouses gets a reference to the given common.NullableInt32 and assigns it to the Warehouses field.
func (o *Tpcc) SetWarehouses(v int32) {
	o.Warehouses.Set(&v)
}

// SetWarehousesNil sets the value for Warehouses to be an explicit nil.
func (o *Tpcc) SetWarehousesNil() {
	o.Warehouses.Set(nil)
}

// UnsetWarehouses ensures that no value is present for Warehouses, not even an explicit nil.
func (o *Tpcc) UnsetWarehouses() {
	o.Warehouses.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tpcc) GetDuration() int32 {
	if o == nil || o.Duration.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Tpcc) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *Tpcc) HasDuration() bool {
	return o != nil && o.Duration.IsSet()
}

// SetDuration gets a reference to the given common.NullableInt32 and assigns it to the Duration field.
func (o *Tpcc) SetDuration(v int32) {
	o.Duration.Set(&v)
}

// SetDurationNil sets the value for Duration to be an explicit nil.
func (o *Tpcc) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil.
func (o *Tpcc) UnsetDuration() {
	o.Duration.Unset()
}

// GetLimitTxPerMin returns the LimitTxPerMin field value if set, zero value otherwise.
func (o *Tpcc) GetLimitTxPerMin() int32 {
	if o == nil || o.LimitTxPerMin == nil {
		var ret int32
		return ret
	}
	return *o.LimitTxPerMin
}

// GetLimitTxPerMinOk returns a tuple with the LimitTxPerMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tpcc) GetLimitTxPerMinOk() (*int32, bool) {
	if o == nil || o.LimitTxPerMin == nil {
		return nil, false
	}
	return o.LimitTxPerMin, true
}

// HasLimitTxPerMin returns a boolean if a field has been set.
func (o *Tpcc) HasLimitTxPerMin() bool {
	return o != nil && o.LimitTxPerMin != nil
}

// SetLimitTxPerMin gets a reference to the given int32 and assigns it to the LimitTxPerMin field.
func (o *Tpcc) SetLimitTxPerMin(v int32) {
	o.LimitTxPerMin = &v
}

// GetNewOrderWeight returns the NewOrderWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tpcc) GetNewOrderWeight() int32 {
	if o == nil || o.NewOrderWeight.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NewOrderWeight.Get()
}

// GetNewOrderWeightOk returns a tuple with the NewOrderWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Tpcc) GetNewOrderWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewOrderWeight.Get(), o.NewOrderWeight.IsSet()
}

// HasNewOrderWeight returns a boolean if a field has been set.
func (o *Tpcc) HasNewOrderWeight() bool {
	return o != nil && o.NewOrderWeight.IsSet()
}

// SetNewOrderWeight gets a reference to the given common.NullableInt32 and assigns it to the NewOrderWeight field.
func (o *Tpcc) SetNewOrderWeight(v int32) {
	o.NewOrderWeight.Set(&v)
}

// SetNewOrderWeightNil sets the value for NewOrderWeight to be an explicit nil.
func (o *Tpcc) SetNewOrderWeightNil() {
	o.NewOrderWeight.Set(nil)
}

// UnsetNewOrderWeight ensures that no value is present for NewOrderWeight, not even an explicit nil.
func (o *Tpcc) UnsetNewOrderWeight() {
	o.NewOrderWeight.Unset()
}

// GetPaymentWeight returns the PaymentWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tpcc) GetPaymentWeight() int32 {
	if o == nil || o.PaymentWeight.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PaymentWeight.Get()
}

// GetPaymentWeightOk returns a tuple with the PaymentWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Tpcc) GetPaymentWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentWeight.Get(), o.PaymentWeight.IsSet()
}

// HasPaymentWeight returns a boolean if a field has been set.
func (o *Tpcc) HasPaymentWeight() bool {
	return o != nil && o.PaymentWeight.IsSet()
}

// SetPaymentWeight gets a reference to the given common.NullableInt32 and assigns it to the PaymentWeight field.
func (o *Tpcc) SetPaymentWeight(v int32) {
	o.PaymentWeight.Set(&v)
}

// SetPaymentWeightNil sets the value for PaymentWeight to be an explicit nil.
func (o *Tpcc) SetPaymentWeightNil() {
	o.PaymentWeight.Set(nil)
}

// UnsetPaymentWeight ensures that no value is present for PaymentWeight, not even an explicit nil.
func (o *Tpcc) UnsetPaymentWeight() {
	o.PaymentWeight.Unset()
}

// GetOrderStatusWeight returns the OrderStatusWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tpcc) GetOrderStatusWeight() int32 {
	if o == nil || o.OrderStatusWeight.Get() == nil {
		var ret int32
		return ret
	}
	return *o.OrderStatusWeight.Get()
}

// GetOrderStatusWeightOk returns a tuple with the OrderStatusWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Tpcc) GetOrderStatusWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderStatusWeight.Get(), o.OrderStatusWeight.IsSet()
}

// HasOrderStatusWeight returns a boolean if a field has been set.
func (o *Tpcc) HasOrderStatusWeight() bool {
	return o != nil && o.OrderStatusWeight.IsSet()
}

// SetOrderStatusWeight gets a reference to the given common.NullableInt32 and assigns it to the OrderStatusWeight field.
func (o *Tpcc) SetOrderStatusWeight(v int32) {
	o.OrderStatusWeight.Set(&v)
}

// SetOrderStatusWeightNil sets the value for OrderStatusWeight to be an explicit nil.
func (o *Tpcc) SetOrderStatusWeightNil() {
	o.OrderStatusWeight.Set(nil)
}

// UnsetOrderStatusWeight ensures that no value is present for OrderStatusWeight, not even an explicit nil.
func (o *Tpcc) UnsetOrderStatusWeight() {
	o.OrderStatusWeight.Unset()
}

// GetDeliveryWeight returns the DeliveryWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tpcc) GetDeliveryWeight() int32 {
	if o == nil || o.DeliveryWeight.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DeliveryWeight.Get()
}

// GetDeliveryWeightOk returns a tuple with the DeliveryWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Tpcc) GetDeliveryWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryWeight.Get(), o.DeliveryWeight.IsSet()
}

// HasDeliveryWeight returns a boolean if a field has been set.
func (o *Tpcc) HasDeliveryWeight() bool {
	return o != nil && o.DeliveryWeight.IsSet()
}

// SetDeliveryWeight gets a reference to the given common.NullableInt32 and assigns it to the DeliveryWeight field.
func (o *Tpcc) SetDeliveryWeight(v int32) {
	o.DeliveryWeight.Set(&v)
}

// SetDeliveryWeightNil sets the value for DeliveryWeight to be an explicit nil.
func (o *Tpcc) SetDeliveryWeightNil() {
	o.DeliveryWeight.Set(nil)
}

// UnsetDeliveryWeight ensures that no value is present for DeliveryWeight, not even an explicit nil.
func (o *Tpcc) UnsetDeliveryWeight() {
	o.DeliveryWeight.Unset()
}

// GetStockLevelWeight returns the StockLevelWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tpcc) GetStockLevelWeight() int32 {
	if o == nil || o.StockLevelWeight.Get() == nil {
		var ret int32
		return ret
	}
	return *o.StockLevelWeight.Get()
}

// GetStockLevelWeightOk returns a tuple with the StockLevelWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Tpcc) GetStockLevelWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StockLevelWeight.Get(), o.StockLevelWeight.IsSet()
}

// HasStockLevelWeight returns a boolean if a field has been set.
func (o *Tpcc) HasStockLevelWeight() bool {
	return o != nil && o.StockLevelWeight.IsSet()
}

// SetStockLevelWeight gets a reference to the given common.NullableInt32 and assigns it to the StockLevelWeight field.
func (o *Tpcc) SetStockLevelWeight(v int32) {
	o.StockLevelWeight.Set(&v)
}

// SetStockLevelWeightNil sets the value for StockLevelWeight to be an explicit nil.
func (o *Tpcc) SetStockLevelWeightNil() {
	o.StockLevelWeight.Set(nil)
}

// UnsetStockLevelWeight ensures that no value is present for StockLevelWeight, not even an explicit nil.
func (o *Tpcc) UnsetStockLevelWeight() {
	o.StockLevelWeight.Unset()
}

// GetExtraArgs returns the ExtraArgs field value if set, zero value otherwise.
func (o *Tpcc) GetExtraArgs() string {
	if o == nil || o.ExtraArgs == nil {
		var ret string
		return ret
	}
	return *o.ExtraArgs
}

// GetExtraArgsOk returns a tuple with the ExtraArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tpcc) GetExtraArgsOk() (*string, bool) {
	if o == nil || o.ExtraArgs == nil {
		return nil, false
	}
	return o.ExtraArgs, true
}

// HasExtraArgs returns a boolean if a field has been set.
func (o *Tpcc) HasExtraArgs() bool {
	return o != nil && o.ExtraArgs != nil
}

// SetExtraArgs gets a reference to the given string and assigns it to the ExtraArgs field.
func (o *Tpcc) SetExtraArgs(v string) {
	o.ExtraArgs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Tpcc) MarshalJSON() ([]byte, error) {
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
	if o.Threads.IsSet() {
		toSerialize["threads"] = o.Threads.Get()
	}
	if o.Warehouses.IsSet() {
		toSerialize["warehouses"] = o.Warehouses.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.LimitTxPerMin != nil {
		toSerialize["limitTxPerMin"] = o.LimitTxPerMin
	}
	if o.NewOrderWeight.IsSet() {
		toSerialize["newOrderWeight"] = o.NewOrderWeight.Get()
	}
	if o.PaymentWeight.IsSet() {
		toSerialize["paymentWeight"] = o.PaymentWeight.Get()
	}
	if o.OrderStatusWeight.IsSet() {
		toSerialize["orderStatusWeight"] = o.OrderStatusWeight.Get()
	}
	if o.DeliveryWeight.IsSet() {
		toSerialize["deliveryWeight"] = o.DeliveryWeight.Get()
	}
	if o.StockLevelWeight.IsSet() {
		toSerialize["stockLevelWeight"] = o.StockLevelWeight.Get()
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
func (o *Tpcc) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Step              *TpccStep            `json:"step,omitempty"`
		LimitCpu          *string              `json:"limitCpu,omitempty"`
		LimitMemory       *string              `json:"limitMemory,omitempty"`
		RequestCpu        *string              `json:"requestCpu,omitempty"`
		RequestMemory     *string              `json:"requestMemory,omitempty"`
		Name              *string              `json:"name,omitempty"`
		Cluster           *string              `json:"cluster"`
		Database          *string              `json:"database"`
		Username          *string              `json:"username"`
		Password          *string              `json:"password"`
		Address           *string              `json:"address"`
		Threads           common.NullableInt32 `json:"threads,omitempty"`
		Warehouses        common.NullableInt32 `json:"warehouses,omitempty"`
		Duration          common.NullableInt32 `json:"duration,omitempty"`
		LimitTxPerMin     *int32               `json:"limitTxPerMin,omitempty"`
		NewOrderWeight    common.NullableInt32 `json:"newOrderWeight,omitempty"`
		PaymentWeight     common.NullableInt32 `json:"paymentWeight,omitempty"`
		OrderStatusWeight common.NullableInt32 `json:"orderStatusWeight,omitempty"`
		DeliveryWeight    common.NullableInt32 `json:"deliveryWeight,omitempty"`
		StockLevelWeight  common.NullableInt32 `json:"stockLevelWeight,omitempty"`
		ExtraArgs         *string              `json:"extraArgs,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"step", "limitCpu", "limitMemory", "requestCpu", "requestMemory", "name", "cluster", "database", "username", "password", "address", "threads", "warehouses", "duration", "limitTxPerMin", "newOrderWeight", "paymentWeight", "orderStatusWeight", "deliveryWeight", "stockLevelWeight", "extraArgs"})
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
	o.Threads = all.Threads
	o.Warehouses = all.Warehouses
	o.Duration = all.Duration
	o.LimitTxPerMin = all.LimitTxPerMin
	o.NewOrderWeight = all.NewOrderWeight
	o.PaymentWeight = all.PaymentWeight
	o.OrderStatusWeight = all.OrderStatusWeight
	o.DeliveryWeight = all.DeliveryWeight
	o.StockLevelWeight = all.StockLevelWeight
	o.ExtraArgs = all.ExtraArgs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
