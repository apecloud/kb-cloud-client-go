// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ClickhouseReplicaHealth struct {
	// Database name.
	Database *string `json:"database,omitempty"`
	// Table name.
	Table *string `json:"table,omitempty"`
	// ClickHouse server host that reported this replica.
	HostName *string `json:"hostName,omitempty"`
	// Replica is in readonly mode. Usually indicates Keeper or network problems.
	IsReadonly *bool `json:"isReadonly,omitempty"`
	// Keeper session expired.
	IsSessionExpired *bool `json:"isSessionExpired,omitempty"`
	// absolute_delay from system.replicas.
	AbsoluteDelaySeconds *float64 `json:"absoluteDelaySeconds,omitempty"`
	// replication queue size.
	QueueSize *int64 `json:"queueSize,omitempty"`
	// Number of inserts waiting in the replication queue.
	InsertsInQueue *int64 `json:"insertsInQueue,omitempty"`
	// Number of merges waiting in the replication queue.
	MergesInQueue *int64 `json:"mergesInQueue,omitempty"`
	// Number of future parts.
	FutureParts *int64 `json:"futureParts,omitempty"`
	// Number of parts to check.
	PartsToCheck *int64 `json:"partsToCheck,omitempty"`
	// total_replicas for the table.
	TotalReplicas *int64 `json:"totalReplicas,omitempty"`
	// active_replicas for the table.
	ActiveReplicas *int64 `json:"activeReplicas,omitempty"`
	// Maximum log index in Keeper.
	LogMaxIndex *int64 `json:"logMaxIndex,omitempty"`
	// Last applied log index.
	LogPointer *int64 `json:"logPointer,omitempty"`
	// Last replication queue update time in RFC3339 UTC.
	LastQueueUpdate *string `json:"lastQueueUpdate,omitempty"`
	// Keeper path of the table.
	ZookeeperPath *string `json:"zookeeperPath,omitempty"`
	// Replica name.
	ReplicaName *string `json:"replicaName,omitempty"`
	// Replica path in Keeper.
	ReplicaPath *string `json:"replicaPath,omitempty"`
	// Last Keeper exception, empty when healthy.
	ZookeeperException *string `json:"zookeeperException,omitempty"`
	// Last replication queue update exception, empty when healthy.
	LastQueueUpdateException *string `json:"lastQueueUpdateException,omitempty"`
	// normal, warning or critical.
	Risk *string `json:"risk,omitempty"`
	// Detected issue codes such as readonly, session_expired, replication_delay, queue_backlog, missing_replicas, keeper_error or queue_update_error.
	Issues []string `json:"issues,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseReplicaHealth instantiates a new ClickhouseReplicaHealth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseReplicaHealth() *ClickhouseReplicaHealth {
	this := ClickhouseReplicaHealth{}
	return &this
}

// NewClickhouseReplicaHealthWithDefaults instantiates a new ClickhouseReplicaHealth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseReplicaHealthWithDefaults() *ClickhouseReplicaHealth {
	this := ClickhouseReplicaHealth{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *ClickhouseReplicaHealth) SetDatabase(v string) {
	o.Database = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetTable() string {
	if o == nil || o.Table == nil {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetTableOk() (*string, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasTable() bool {
	return o != nil && o.Table != nil
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *ClickhouseReplicaHealth) SetTable(v string) {
	o.Table = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasHostName() bool {
	return o != nil && o.HostName != nil
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *ClickhouseReplicaHealth) SetHostName(v string) {
	o.HostName = &v
}

// GetIsReadonly returns the IsReadonly field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetIsReadonly() bool {
	if o == nil || o.IsReadonly == nil {
		var ret bool
		return ret
	}
	return *o.IsReadonly
}

// GetIsReadonlyOk returns a tuple with the IsReadonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetIsReadonlyOk() (*bool, bool) {
	if o == nil || o.IsReadonly == nil {
		return nil, false
	}
	return o.IsReadonly, true
}

// HasIsReadonly returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasIsReadonly() bool {
	return o != nil && o.IsReadonly != nil
}

// SetIsReadonly gets a reference to the given bool and assigns it to the IsReadonly field.
func (o *ClickhouseReplicaHealth) SetIsReadonly(v bool) {
	o.IsReadonly = &v
}

// GetIsSessionExpired returns the IsSessionExpired field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetIsSessionExpired() bool {
	if o == nil || o.IsSessionExpired == nil {
		var ret bool
		return ret
	}
	return *o.IsSessionExpired
}

// GetIsSessionExpiredOk returns a tuple with the IsSessionExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetIsSessionExpiredOk() (*bool, bool) {
	if o == nil || o.IsSessionExpired == nil {
		return nil, false
	}
	return o.IsSessionExpired, true
}

// HasIsSessionExpired returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasIsSessionExpired() bool {
	return o != nil && o.IsSessionExpired != nil
}

// SetIsSessionExpired gets a reference to the given bool and assigns it to the IsSessionExpired field.
func (o *ClickhouseReplicaHealth) SetIsSessionExpired(v bool) {
	o.IsSessionExpired = &v
}

// GetAbsoluteDelaySeconds returns the AbsoluteDelaySeconds field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetAbsoluteDelaySeconds() float64 {
	if o == nil || o.AbsoluteDelaySeconds == nil {
		var ret float64
		return ret
	}
	return *o.AbsoluteDelaySeconds
}

// GetAbsoluteDelaySecondsOk returns a tuple with the AbsoluteDelaySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetAbsoluteDelaySecondsOk() (*float64, bool) {
	if o == nil || o.AbsoluteDelaySeconds == nil {
		return nil, false
	}
	return o.AbsoluteDelaySeconds, true
}

// HasAbsoluteDelaySeconds returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasAbsoluteDelaySeconds() bool {
	return o != nil && o.AbsoluteDelaySeconds != nil
}

// SetAbsoluteDelaySeconds gets a reference to the given float64 and assigns it to the AbsoluteDelaySeconds field.
func (o *ClickhouseReplicaHealth) SetAbsoluteDelaySeconds(v float64) {
	o.AbsoluteDelaySeconds = &v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetQueueSize() int64 {
	if o == nil || o.QueueSize == nil {
		var ret int64
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetQueueSizeOk() (*int64, bool) {
	if o == nil || o.QueueSize == nil {
		return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasQueueSize() bool {
	return o != nil && o.QueueSize != nil
}

// SetQueueSize gets a reference to the given int64 and assigns it to the QueueSize field.
func (o *ClickhouseReplicaHealth) SetQueueSize(v int64) {
	o.QueueSize = &v
}

// GetInsertsInQueue returns the InsertsInQueue field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetInsertsInQueue() int64 {
	if o == nil || o.InsertsInQueue == nil {
		var ret int64
		return ret
	}
	return *o.InsertsInQueue
}

// GetInsertsInQueueOk returns a tuple with the InsertsInQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetInsertsInQueueOk() (*int64, bool) {
	if o == nil || o.InsertsInQueue == nil {
		return nil, false
	}
	return o.InsertsInQueue, true
}

// HasInsertsInQueue returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasInsertsInQueue() bool {
	return o != nil && o.InsertsInQueue != nil
}

// SetInsertsInQueue gets a reference to the given int64 and assigns it to the InsertsInQueue field.
func (o *ClickhouseReplicaHealth) SetInsertsInQueue(v int64) {
	o.InsertsInQueue = &v
}

// GetMergesInQueue returns the MergesInQueue field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetMergesInQueue() int64 {
	if o == nil || o.MergesInQueue == nil {
		var ret int64
		return ret
	}
	return *o.MergesInQueue
}

// GetMergesInQueueOk returns a tuple with the MergesInQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetMergesInQueueOk() (*int64, bool) {
	if o == nil || o.MergesInQueue == nil {
		return nil, false
	}
	return o.MergesInQueue, true
}

// HasMergesInQueue returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasMergesInQueue() bool {
	return o != nil && o.MergesInQueue != nil
}

// SetMergesInQueue gets a reference to the given int64 and assigns it to the MergesInQueue field.
func (o *ClickhouseReplicaHealth) SetMergesInQueue(v int64) {
	o.MergesInQueue = &v
}

// GetFutureParts returns the FutureParts field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetFutureParts() int64 {
	if o == nil || o.FutureParts == nil {
		var ret int64
		return ret
	}
	return *o.FutureParts
}

// GetFuturePartsOk returns a tuple with the FutureParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetFuturePartsOk() (*int64, bool) {
	if o == nil || o.FutureParts == nil {
		return nil, false
	}
	return o.FutureParts, true
}

// HasFutureParts returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasFutureParts() bool {
	return o != nil && o.FutureParts != nil
}

// SetFutureParts gets a reference to the given int64 and assigns it to the FutureParts field.
func (o *ClickhouseReplicaHealth) SetFutureParts(v int64) {
	o.FutureParts = &v
}

// GetPartsToCheck returns the PartsToCheck field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetPartsToCheck() int64 {
	if o == nil || o.PartsToCheck == nil {
		var ret int64
		return ret
	}
	return *o.PartsToCheck
}

// GetPartsToCheckOk returns a tuple with the PartsToCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetPartsToCheckOk() (*int64, bool) {
	if o == nil || o.PartsToCheck == nil {
		return nil, false
	}
	return o.PartsToCheck, true
}

// HasPartsToCheck returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasPartsToCheck() bool {
	return o != nil && o.PartsToCheck != nil
}

// SetPartsToCheck gets a reference to the given int64 and assigns it to the PartsToCheck field.
func (o *ClickhouseReplicaHealth) SetPartsToCheck(v int64) {
	o.PartsToCheck = &v
}

// GetTotalReplicas returns the TotalReplicas field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetTotalReplicas() int64 {
	if o == nil || o.TotalReplicas == nil {
		var ret int64
		return ret
	}
	return *o.TotalReplicas
}

// GetTotalReplicasOk returns a tuple with the TotalReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetTotalReplicasOk() (*int64, bool) {
	if o == nil || o.TotalReplicas == nil {
		return nil, false
	}
	return o.TotalReplicas, true
}

// HasTotalReplicas returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasTotalReplicas() bool {
	return o != nil && o.TotalReplicas != nil
}

// SetTotalReplicas gets a reference to the given int64 and assigns it to the TotalReplicas field.
func (o *ClickhouseReplicaHealth) SetTotalReplicas(v int64) {
	o.TotalReplicas = &v
}

// GetActiveReplicas returns the ActiveReplicas field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetActiveReplicas() int64 {
	if o == nil || o.ActiveReplicas == nil {
		var ret int64
		return ret
	}
	return *o.ActiveReplicas
}

// GetActiveReplicasOk returns a tuple with the ActiveReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetActiveReplicasOk() (*int64, bool) {
	if o == nil || o.ActiveReplicas == nil {
		return nil, false
	}
	return o.ActiveReplicas, true
}

// HasActiveReplicas returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasActiveReplicas() bool {
	return o != nil && o.ActiveReplicas != nil
}

// SetActiveReplicas gets a reference to the given int64 and assigns it to the ActiveReplicas field.
func (o *ClickhouseReplicaHealth) SetActiveReplicas(v int64) {
	o.ActiveReplicas = &v
}

// GetLogMaxIndex returns the LogMaxIndex field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetLogMaxIndex() int64 {
	if o == nil || o.LogMaxIndex == nil {
		var ret int64
		return ret
	}
	return *o.LogMaxIndex
}

// GetLogMaxIndexOk returns a tuple with the LogMaxIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetLogMaxIndexOk() (*int64, bool) {
	if o == nil || o.LogMaxIndex == nil {
		return nil, false
	}
	return o.LogMaxIndex, true
}

// HasLogMaxIndex returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasLogMaxIndex() bool {
	return o != nil && o.LogMaxIndex != nil
}

// SetLogMaxIndex gets a reference to the given int64 and assigns it to the LogMaxIndex field.
func (o *ClickhouseReplicaHealth) SetLogMaxIndex(v int64) {
	o.LogMaxIndex = &v
}

// GetLogPointer returns the LogPointer field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetLogPointer() int64 {
	if o == nil || o.LogPointer == nil {
		var ret int64
		return ret
	}
	return *o.LogPointer
}

// GetLogPointerOk returns a tuple with the LogPointer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetLogPointerOk() (*int64, bool) {
	if o == nil || o.LogPointer == nil {
		return nil, false
	}
	return o.LogPointer, true
}

// HasLogPointer returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasLogPointer() bool {
	return o != nil && o.LogPointer != nil
}

// SetLogPointer gets a reference to the given int64 and assigns it to the LogPointer field.
func (o *ClickhouseReplicaHealth) SetLogPointer(v int64) {
	o.LogPointer = &v
}

// GetLastQueueUpdate returns the LastQueueUpdate field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetLastQueueUpdate() string {
	if o == nil || o.LastQueueUpdate == nil {
		var ret string
		return ret
	}
	return *o.LastQueueUpdate
}

// GetLastQueueUpdateOk returns a tuple with the LastQueueUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetLastQueueUpdateOk() (*string, bool) {
	if o == nil || o.LastQueueUpdate == nil {
		return nil, false
	}
	return o.LastQueueUpdate, true
}

// HasLastQueueUpdate returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasLastQueueUpdate() bool {
	return o != nil && o.LastQueueUpdate != nil
}

// SetLastQueueUpdate gets a reference to the given string and assigns it to the LastQueueUpdate field.
func (o *ClickhouseReplicaHealth) SetLastQueueUpdate(v string) {
	o.LastQueueUpdate = &v
}

// GetZookeeperPath returns the ZookeeperPath field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetZookeeperPath() string {
	if o == nil || o.ZookeeperPath == nil {
		var ret string
		return ret
	}
	return *o.ZookeeperPath
}

// GetZookeeperPathOk returns a tuple with the ZookeeperPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetZookeeperPathOk() (*string, bool) {
	if o == nil || o.ZookeeperPath == nil {
		return nil, false
	}
	return o.ZookeeperPath, true
}

// HasZookeeperPath returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasZookeeperPath() bool {
	return o != nil && o.ZookeeperPath != nil
}

// SetZookeeperPath gets a reference to the given string and assigns it to the ZookeeperPath field.
func (o *ClickhouseReplicaHealth) SetZookeeperPath(v string) {
	o.ZookeeperPath = &v
}

// GetReplicaName returns the ReplicaName field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetReplicaName() string {
	if o == nil || o.ReplicaName == nil {
		var ret string
		return ret
	}
	return *o.ReplicaName
}

// GetReplicaNameOk returns a tuple with the ReplicaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetReplicaNameOk() (*string, bool) {
	if o == nil || o.ReplicaName == nil {
		return nil, false
	}
	return o.ReplicaName, true
}

// HasReplicaName returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasReplicaName() bool {
	return o != nil && o.ReplicaName != nil
}

// SetReplicaName gets a reference to the given string and assigns it to the ReplicaName field.
func (o *ClickhouseReplicaHealth) SetReplicaName(v string) {
	o.ReplicaName = &v
}

// GetReplicaPath returns the ReplicaPath field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetReplicaPath() string {
	if o == nil || o.ReplicaPath == nil {
		var ret string
		return ret
	}
	return *o.ReplicaPath
}

// GetReplicaPathOk returns a tuple with the ReplicaPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetReplicaPathOk() (*string, bool) {
	if o == nil || o.ReplicaPath == nil {
		return nil, false
	}
	return o.ReplicaPath, true
}

// HasReplicaPath returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasReplicaPath() bool {
	return o != nil && o.ReplicaPath != nil
}

// SetReplicaPath gets a reference to the given string and assigns it to the ReplicaPath field.
func (o *ClickhouseReplicaHealth) SetReplicaPath(v string) {
	o.ReplicaPath = &v
}

// GetZookeeperException returns the ZookeeperException field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetZookeeperException() string {
	if o == nil || o.ZookeeperException == nil {
		var ret string
		return ret
	}
	return *o.ZookeeperException
}

// GetZookeeperExceptionOk returns a tuple with the ZookeeperException field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetZookeeperExceptionOk() (*string, bool) {
	if o == nil || o.ZookeeperException == nil {
		return nil, false
	}
	return o.ZookeeperException, true
}

// HasZookeeperException returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasZookeeperException() bool {
	return o != nil && o.ZookeeperException != nil
}

// SetZookeeperException gets a reference to the given string and assigns it to the ZookeeperException field.
func (o *ClickhouseReplicaHealth) SetZookeeperException(v string) {
	o.ZookeeperException = &v
}

// GetLastQueueUpdateException returns the LastQueueUpdateException field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetLastQueueUpdateException() string {
	if o == nil || o.LastQueueUpdateException == nil {
		var ret string
		return ret
	}
	return *o.LastQueueUpdateException
}

// GetLastQueueUpdateExceptionOk returns a tuple with the LastQueueUpdateException field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetLastQueueUpdateExceptionOk() (*string, bool) {
	if o == nil || o.LastQueueUpdateException == nil {
		return nil, false
	}
	return o.LastQueueUpdateException, true
}

// HasLastQueueUpdateException returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasLastQueueUpdateException() bool {
	return o != nil && o.LastQueueUpdateException != nil
}

// SetLastQueueUpdateException gets a reference to the given string and assigns it to the LastQueueUpdateException field.
func (o *ClickhouseReplicaHealth) SetLastQueueUpdateException(v string) {
	o.LastQueueUpdateException = &v
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetRisk() string {
	if o == nil || o.Risk == nil {
		var ret string
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetRiskOk() (*string, bool) {
	if o == nil || o.Risk == nil {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasRisk() bool {
	return o != nil && o.Risk != nil
}

// SetRisk gets a reference to the given string and assigns it to the Risk field.
func (o *ClickhouseReplicaHealth) SetRisk(v string) {
	o.Risk = &v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *ClickhouseReplicaHealth) GetIssues() []string {
	if o == nil || o.Issues == nil {
		var ret []string
		return ret
	}
	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseReplicaHealth) GetIssuesOk() (*[]string, bool) {
	if o == nil || o.Issues == nil {
		return nil, false
	}
	return &o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *ClickhouseReplicaHealth) HasIssues() bool {
	return o != nil && o.Issues != nil
}

// SetIssues gets a reference to the given []string and assigns it to the Issues field.
func (o *ClickhouseReplicaHealth) SetIssues(v []string) {
	o.Issues = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseReplicaHealth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.Table != nil {
		toSerialize["table"] = o.Table
	}
	if o.HostName != nil {
		toSerialize["hostName"] = o.HostName
	}
	if o.IsReadonly != nil {
		toSerialize["isReadonly"] = o.IsReadonly
	}
	if o.IsSessionExpired != nil {
		toSerialize["isSessionExpired"] = o.IsSessionExpired
	}
	if o.AbsoluteDelaySeconds != nil {
		toSerialize["absoluteDelaySeconds"] = o.AbsoluteDelaySeconds
	}
	if o.QueueSize != nil {
		toSerialize["queueSize"] = o.QueueSize
	}
	if o.InsertsInQueue != nil {
		toSerialize["insertsInQueue"] = o.InsertsInQueue
	}
	if o.MergesInQueue != nil {
		toSerialize["mergesInQueue"] = o.MergesInQueue
	}
	if o.FutureParts != nil {
		toSerialize["futureParts"] = o.FutureParts
	}
	if o.PartsToCheck != nil {
		toSerialize["partsToCheck"] = o.PartsToCheck
	}
	if o.TotalReplicas != nil {
		toSerialize["totalReplicas"] = o.TotalReplicas
	}
	if o.ActiveReplicas != nil {
		toSerialize["activeReplicas"] = o.ActiveReplicas
	}
	if o.LogMaxIndex != nil {
		toSerialize["logMaxIndex"] = o.LogMaxIndex
	}
	if o.LogPointer != nil {
		toSerialize["logPointer"] = o.LogPointer
	}
	if o.LastQueueUpdate != nil {
		toSerialize["lastQueueUpdate"] = o.LastQueueUpdate
	}
	if o.ZookeeperPath != nil {
		toSerialize["zookeeperPath"] = o.ZookeeperPath
	}
	if o.ReplicaName != nil {
		toSerialize["replicaName"] = o.ReplicaName
	}
	if o.ReplicaPath != nil {
		toSerialize["replicaPath"] = o.ReplicaPath
	}
	if o.ZookeeperException != nil {
		toSerialize["zookeeperException"] = o.ZookeeperException
	}
	if o.LastQueueUpdateException != nil {
		toSerialize["lastQueueUpdateException"] = o.LastQueueUpdateException
	}
	if o.Risk != nil {
		toSerialize["risk"] = o.Risk
	}
	if o.Issues != nil {
		toSerialize["issues"] = o.Issues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseReplicaHealth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database                 *string  `json:"database,omitempty"`
		Table                    *string  `json:"table,omitempty"`
		HostName                 *string  `json:"hostName,omitempty"`
		IsReadonly               *bool    `json:"isReadonly,omitempty"`
		IsSessionExpired         *bool    `json:"isSessionExpired,omitempty"`
		AbsoluteDelaySeconds     *float64 `json:"absoluteDelaySeconds,omitempty"`
		QueueSize                *int64   `json:"queueSize,omitempty"`
		InsertsInQueue           *int64   `json:"insertsInQueue,omitempty"`
		MergesInQueue            *int64   `json:"mergesInQueue,omitempty"`
		FutureParts              *int64   `json:"futureParts,omitempty"`
		PartsToCheck             *int64   `json:"partsToCheck,omitempty"`
		TotalReplicas            *int64   `json:"totalReplicas,omitempty"`
		ActiveReplicas           *int64   `json:"activeReplicas,omitempty"`
		LogMaxIndex              *int64   `json:"logMaxIndex,omitempty"`
		LogPointer               *int64   `json:"logPointer,omitempty"`
		LastQueueUpdate          *string  `json:"lastQueueUpdate,omitempty"`
		ZookeeperPath            *string  `json:"zookeeperPath,omitempty"`
		ReplicaName              *string  `json:"replicaName,omitempty"`
		ReplicaPath              *string  `json:"replicaPath,omitempty"`
		ZookeeperException       *string  `json:"zookeeperException,omitempty"`
		LastQueueUpdateException *string  `json:"lastQueueUpdateException,omitempty"`
		Risk                     *string  `json:"risk,omitempty"`
		Issues                   []string `json:"issues,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "table", "hostName", "isReadonly", "isSessionExpired", "absoluteDelaySeconds", "queueSize", "insertsInQueue", "mergesInQueue", "futureParts", "partsToCheck", "totalReplicas", "activeReplicas", "logMaxIndex", "logPointer", "lastQueueUpdate", "zookeeperPath", "replicaName", "replicaPath", "zookeeperException", "lastQueueUpdateException", "risk", "issues"})
	} else {
		return err
	}
	o.Database = all.Database
	o.Table = all.Table
	o.HostName = all.HostName
	o.IsReadonly = all.IsReadonly
	o.IsSessionExpired = all.IsSessionExpired
	o.AbsoluteDelaySeconds = all.AbsoluteDelaySeconds
	o.QueueSize = all.QueueSize
	o.InsertsInQueue = all.InsertsInQueue
	o.MergesInQueue = all.MergesInQueue
	o.FutureParts = all.FutureParts
	o.PartsToCheck = all.PartsToCheck
	o.TotalReplicas = all.TotalReplicas
	o.ActiveReplicas = all.ActiveReplicas
	o.LogMaxIndex = all.LogMaxIndex
	o.LogPointer = all.LogPointer
	o.LastQueueUpdate = all.LastQueueUpdate
	o.ZookeeperPath = all.ZookeeperPath
	o.ReplicaName = all.ReplicaName
	o.ReplicaPath = all.ReplicaPath
	o.ZookeeperException = all.ZookeeperException
	o.LastQueueUpdateException = all.LastQueueUpdateException
	o.Risk = all.Risk
	o.Issues = all.Issues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
