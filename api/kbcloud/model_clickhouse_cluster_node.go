// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ClickhouseClusterNode struct {
	// Cluster name.
	Cluster *string `json:"cluster,omitempty"`
	// Shard number.
	ShardNum *int64 `json:"shardNum,omitempty"`
	// Shard weight.
	ShardWeight *int64 `json:"shardWeight,omitempty"`
	// Replica number inside the shard.
	ReplicaNum *int64 `json:"replicaNum,omitempty"`
	// Configured host name.
	HostName *string `json:"hostName,omitempty"`
	// Configured host address.
	HostAddress *string `json:"hostAddress,omitempty"`
	// Configured TCP port.
	Port *int64 `json:"port,omitempty"`
	// Whether this node is the connected node.
	IsLocal *bool `json:"isLocal,omitempty"`
	// Configured user.
	User *string `json:"user,omitempty"`
	// Configured default database.
	DefaultDatabase *string `json:"defaultDatabase,omitempty"`
	// Whether the cluster uses internal replication.
	InternalReplication *bool `json:"internalReplication,omitempty"`
	// Connection errors observed for this node.
	ErrorsCount *int64 `json:"errorsCount,omitempty"`
	// Slowdown events observed for this node.
	SlowdownsCount *int64 `json:"slowdownsCount,omitempty"`
	// Estimated recovery time in seconds, zero when the node is healthy.
	EstimatedRecoveryTime *int64 `json:"estimatedRecoveryTime,omitempty"`
	// normal or warning.
	Risk *string `json:"risk,omitempty"`
	// Detected issue codes such as connection_errors, slowdowns or estimated_recovery_time.
	Issues []string `json:"issues,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseClusterNode instantiates a new ClickhouseClusterNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseClusterNode() *ClickhouseClusterNode {
	this := ClickhouseClusterNode{}
	return &this
}

// NewClickhouseClusterNodeWithDefaults instantiates a new ClickhouseClusterNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseClusterNodeWithDefaults() *ClickhouseClusterNode {
	this := ClickhouseClusterNode{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ClickhouseClusterNode) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterNode) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ClickhouseClusterNode) HasCluster() bool {
	return o != nil && o.Cluster != nil
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *ClickhouseClusterNode) SetCluster(v string) {
	o.Cluster = &v
}

// GetShardNum returns the ShardNum field value if set, zero value otherwise.
func (o *ClickhouseClusterNode) GetShardNum() int64 {
	if o == nil || o.ShardNum == nil {
		var ret int64
		return ret
	}
	return *o.ShardNum
}

// GetShardNumOk returns a tuple with the ShardNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterNode) GetShardNumOk() (*int64, bool) {
	if o == nil || o.ShardNum == nil {
		return nil, false
	}
	return o.ShardNum, true
}

// HasShardNum returns a boolean if a field has been set.
func (o *ClickhouseClusterNode) HasShardNum() bool {
	return o != nil && o.ShardNum != nil
}

// SetShardNum gets a reference to the given int64 and assigns it to the ShardNum field.
func (o *ClickhouseClusterNode) SetShardNum(v int64) {
	o.ShardNum = &v
}

// GetShardWeight returns the ShardWeight field value if set, zero value otherwise.
func (o *ClickhouseClusterNode) GetShardWeight() int64 {
	if o == nil || o.ShardWeight == nil {
		var ret int64
		return ret
	}
	return *o.ShardWeight
}

// GetShardWeightOk returns a tuple with the ShardWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterNode) GetShardWeightOk() (*int64, bool) {
	if o == nil || o.ShardWeight == nil {
		return nil, false
	}
	return o.ShardWeight, true
}

// HasShardWeight returns a boolean if a field has been set.
func (o *ClickhouseClusterNode) HasShardWeight() bool {
	return o != nil && o.ShardWeight != nil
}

// SetShardWeight gets a reference to the given int64 and assigns it to the ShardWeight field.
func (o *ClickhouseClusterNode) SetShardWeight(v int64) {
	o.ShardWeight = &v
}

// GetReplicaNum returns the ReplicaNum field value if set, zero value otherwise.
func (o *ClickhouseClusterNode) GetReplicaNum() int64 {
	if o == nil || o.ReplicaNum == nil {
		var ret int64
		return ret
	}
	return *o.ReplicaNum
}

// GetReplicaNumOk returns a tuple with the ReplicaNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterNode) GetReplicaNumOk() (*int64, bool) {
	if o == nil || o.ReplicaNum == nil {
		return nil, false
	}
	return o.ReplicaNum, true
}

// HasReplicaNum returns a boolean if a field has been set.
func (o *ClickhouseClusterNode) HasReplicaNum() bool {
	return o != nil && o.ReplicaNum != nil
}

// SetReplicaNum gets a reference to the given int64 and assigns it to the ReplicaNum field.
func (o *ClickhouseClusterNode) SetReplicaNum(v int64) {
	o.ReplicaNum = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *ClickhouseClusterNode) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterNode) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *ClickhouseClusterNode) HasHostName() bool {
	return o != nil && o.HostName != nil
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *ClickhouseClusterNode) SetHostName(v string) {
	o.HostName = &v
}

// GetHostAddress returns the HostAddress field value if set, zero value otherwise.
func (o *ClickhouseClusterNode) GetHostAddress() string {
	if o == nil || o.HostAddress == nil {
		var ret string
		return ret
	}
	return *o.HostAddress
}

// GetHostAddressOk returns a tuple with the HostAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterNode) GetHostAddressOk() (*string, bool) {
	if o == nil || o.HostAddress == nil {
		return nil, false
	}
	return o.HostAddress, true
}

// HasHostAddress returns a boolean if a field has been set.
func (o *ClickhouseClusterNode) HasHostAddress() bool {
	return o != nil && o.HostAddress != nil
}

// SetHostAddress gets a reference to the given string and assigns it to the HostAddress field.
func (o *ClickhouseClusterNode) SetHostAddress(v string) {
	o.HostAddress = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ClickhouseClusterNode) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterNode) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ClickhouseClusterNode) HasPort() bool {
	return o != nil && o.Port != nil
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *ClickhouseClusterNode) SetPort(v int64) {
	o.Port = &v
}

// GetIsLocal returns the IsLocal field value if set, zero value otherwise.
func (o *ClickhouseClusterNode) GetIsLocal() bool {
	if o == nil || o.IsLocal == nil {
		var ret bool
		return ret
	}
	return *o.IsLocal
}

// GetIsLocalOk returns a tuple with the IsLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterNode) GetIsLocalOk() (*bool, bool) {
	if o == nil || o.IsLocal == nil {
		return nil, false
	}
	return o.IsLocal, true
}

// HasIsLocal returns a boolean if a field has been set.
func (o *ClickhouseClusterNode) HasIsLocal() bool {
	return o != nil && o.IsLocal != nil
}

// SetIsLocal gets a reference to the given bool and assigns it to the IsLocal field.
func (o *ClickhouseClusterNode) SetIsLocal(v bool) {
	o.IsLocal = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ClickhouseClusterNode) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterNode) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ClickhouseClusterNode) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ClickhouseClusterNode) SetUser(v string) {
	o.User = &v
}

// GetDefaultDatabase returns the DefaultDatabase field value if set, zero value otherwise.
func (o *ClickhouseClusterNode) GetDefaultDatabase() string {
	if o == nil || o.DefaultDatabase == nil {
		var ret string
		return ret
	}
	return *o.DefaultDatabase
}

// GetDefaultDatabaseOk returns a tuple with the DefaultDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterNode) GetDefaultDatabaseOk() (*string, bool) {
	if o == nil || o.DefaultDatabase == nil {
		return nil, false
	}
	return o.DefaultDatabase, true
}

// HasDefaultDatabase returns a boolean if a field has been set.
func (o *ClickhouseClusterNode) HasDefaultDatabase() bool {
	return o != nil && o.DefaultDatabase != nil
}

// SetDefaultDatabase gets a reference to the given string and assigns it to the DefaultDatabase field.
func (o *ClickhouseClusterNode) SetDefaultDatabase(v string) {
	o.DefaultDatabase = &v
}

// GetInternalReplication returns the InternalReplication field value if set, zero value otherwise.
func (o *ClickhouseClusterNode) GetInternalReplication() bool {
	if o == nil || o.InternalReplication == nil {
		var ret bool
		return ret
	}
	return *o.InternalReplication
}

// GetInternalReplicationOk returns a tuple with the InternalReplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterNode) GetInternalReplicationOk() (*bool, bool) {
	if o == nil || o.InternalReplication == nil {
		return nil, false
	}
	return o.InternalReplication, true
}

// HasInternalReplication returns a boolean if a field has been set.
func (o *ClickhouseClusterNode) HasInternalReplication() bool {
	return o != nil && o.InternalReplication != nil
}

// SetInternalReplication gets a reference to the given bool and assigns it to the InternalReplication field.
func (o *ClickhouseClusterNode) SetInternalReplication(v bool) {
	o.InternalReplication = &v
}

// GetErrorsCount returns the ErrorsCount field value if set, zero value otherwise.
func (o *ClickhouseClusterNode) GetErrorsCount() int64 {
	if o == nil || o.ErrorsCount == nil {
		var ret int64
		return ret
	}
	return *o.ErrorsCount
}

// GetErrorsCountOk returns a tuple with the ErrorsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterNode) GetErrorsCountOk() (*int64, bool) {
	if o == nil || o.ErrorsCount == nil {
		return nil, false
	}
	return o.ErrorsCount, true
}

// HasErrorsCount returns a boolean if a field has been set.
func (o *ClickhouseClusterNode) HasErrorsCount() bool {
	return o != nil && o.ErrorsCount != nil
}

// SetErrorsCount gets a reference to the given int64 and assigns it to the ErrorsCount field.
func (o *ClickhouseClusterNode) SetErrorsCount(v int64) {
	o.ErrorsCount = &v
}

// GetSlowdownsCount returns the SlowdownsCount field value if set, zero value otherwise.
func (o *ClickhouseClusterNode) GetSlowdownsCount() int64 {
	if o == nil || o.SlowdownsCount == nil {
		var ret int64
		return ret
	}
	return *o.SlowdownsCount
}

// GetSlowdownsCountOk returns a tuple with the SlowdownsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterNode) GetSlowdownsCountOk() (*int64, bool) {
	if o == nil || o.SlowdownsCount == nil {
		return nil, false
	}
	return o.SlowdownsCount, true
}

// HasSlowdownsCount returns a boolean if a field has been set.
func (o *ClickhouseClusterNode) HasSlowdownsCount() bool {
	return o != nil && o.SlowdownsCount != nil
}

// SetSlowdownsCount gets a reference to the given int64 and assigns it to the SlowdownsCount field.
func (o *ClickhouseClusterNode) SetSlowdownsCount(v int64) {
	o.SlowdownsCount = &v
}

// GetEstimatedRecoveryTime returns the EstimatedRecoveryTime field value if set, zero value otherwise.
func (o *ClickhouseClusterNode) GetEstimatedRecoveryTime() int64 {
	if o == nil || o.EstimatedRecoveryTime == nil {
		var ret int64
		return ret
	}
	return *o.EstimatedRecoveryTime
}

// GetEstimatedRecoveryTimeOk returns a tuple with the EstimatedRecoveryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterNode) GetEstimatedRecoveryTimeOk() (*int64, bool) {
	if o == nil || o.EstimatedRecoveryTime == nil {
		return nil, false
	}
	return o.EstimatedRecoveryTime, true
}

// HasEstimatedRecoveryTime returns a boolean if a field has been set.
func (o *ClickhouseClusterNode) HasEstimatedRecoveryTime() bool {
	return o != nil && o.EstimatedRecoveryTime != nil
}

// SetEstimatedRecoveryTime gets a reference to the given int64 and assigns it to the EstimatedRecoveryTime field.
func (o *ClickhouseClusterNode) SetEstimatedRecoveryTime(v int64) {
	o.EstimatedRecoveryTime = &v
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *ClickhouseClusterNode) GetRisk() string {
	if o == nil || o.Risk == nil {
		var ret string
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterNode) GetRiskOk() (*string, bool) {
	if o == nil || o.Risk == nil {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *ClickhouseClusterNode) HasRisk() bool {
	return o != nil && o.Risk != nil
}

// SetRisk gets a reference to the given string and assigns it to the Risk field.
func (o *ClickhouseClusterNode) SetRisk(v string) {
	o.Risk = &v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *ClickhouseClusterNode) GetIssues() []string {
	if o == nil || o.Issues == nil {
		var ret []string
		return ret
	}
	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterNode) GetIssuesOk() (*[]string, bool) {
	if o == nil || o.Issues == nil {
		return nil, false
	}
	return &o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *ClickhouseClusterNode) HasIssues() bool {
	return o != nil && o.Issues != nil
}

// SetIssues gets a reference to the given []string and assigns it to the Issues field.
func (o *ClickhouseClusterNode) SetIssues(v []string) {
	o.Issues = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseClusterNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.ShardNum != nil {
		toSerialize["shardNum"] = o.ShardNum
	}
	if o.ShardWeight != nil {
		toSerialize["shardWeight"] = o.ShardWeight
	}
	if o.ReplicaNum != nil {
		toSerialize["replicaNum"] = o.ReplicaNum
	}
	if o.HostName != nil {
		toSerialize["hostName"] = o.HostName
	}
	if o.HostAddress != nil {
		toSerialize["hostAddress"] = o.HostAddress
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.IsLocal != nil {
		toSerialize["isLocal"] = o.IsLocal
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.DefaultDatabase != nil {
		toSerialize["defaultDatabase"] = o.DefaultDatabase
	}
	if o.InternalReplication != nil {
		toSerialize["internalReplication"] = o.InternalReplication
	}
	if o.ErrorsCount != nil {
		toSerialize["errorsCount"] = o.ErrorsCount
	}
	if o.SlowdownsCount != nil {
		toSerialize["slowdownsCount"] = o.SlowdownsCount
	}
	if o.EstimatedRecoveryTime != nil {
		toSerialize["estimatedRecoveryTime"] = o.EstimatedRecoveryTime
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
func (o *ClickhouseClusterNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cluster               *string  `json:"cluster,omitempty"`
		ShardNum              *int64   `json:"shardNum,omitempty"`
		ShardWeight           *int64   `json:"shardWeight,omitempty"`
		ReplicaNum            *int64   `json:"replicaNum,omitempty"`
		HostName              *string  `json:"hostName,omitempty"`
		HostAddress           *string  `json:"hostAddress,omitempty"`
		Port                  *int64   `json:"port,omitempty"`
		IsLocal               *bool    `json:"isLocal,omitempty"`
		User                  *string  `json:"user,omitempty"`
		DefaultDatabase       *string  `json:"defaultDatabase,omitempty"`
		InternalReplication   *bool    `json:"internalReplication,omitempty"`
		ErrorsCount           *int64   `json:"errorsCount,omitempty"`
		SlowdownsCount        *int64   `json:"slowdownsCount,omitempty"`
		EstimatedRecoveryTime *int64   `json:"estimatedRecoveryTime,omitempty"`
		Risk                  *string  `json:"risk,omitempty"`
		Issues                []string `json:"issues,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cluster", "shardNum", "shardWeight", "replicaNum", "hostName", "hostAddress", "port", "isLocal", "user", "defaultDatabase", "internalReplication", "errorsCount", "slowdownsCount", "estimatedRecoveryTime", "risk", "issues"})
	} else {
		return err
	}
	o.Cluster = all.Cluster
	o.ShardNum = all.ShardNum
	o.ShardWeight = all.ShardWeight
	o.ReplicaNum = all.ReplicaNum
	o.HostName = all.HostName
	o.HostAddress = all.HostAddress
	o.Port = all.Port
	o.IsLocal = all.IsLocal
	o.User = all.User
	o.DefaultDatabase = all.DefaultDatabase
	o.InternalReplication = all.InternalReplication
	o.ErrorsCount = all.ErrorsCount
	o.SlowdownsCount = all.SlowdownsCount
	o.EstimatedRecoveryTime = all.EstimatedRecoveryTime
	o.Risk = all.Risk
	o.Issues = all.Issues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
