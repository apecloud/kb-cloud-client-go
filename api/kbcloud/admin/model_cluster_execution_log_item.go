// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterExecutionLogItem Cluster execution log item represents a single log entry
type ClusterExecutionLogItem struct {
	Client        *string                `json:"client,omitempty"`
	Command       *string                `json:"command,omitempty"`
	DbName        *string                `json:"dbName,omitempty"`
	ExecutionTime *float64               `json:"executionTime,omitempty"`
	Extra         map[string]interface{} `json:"extra,omitempty"`
	Timestamp     *int32                 `json:"timestamp,omitempty"`
	User          *string                `json:"user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterExecutionLogItem instantiates a new ClusterExecutionLogItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterExecutionLogItem() *ClusterExecutionLogItem {
	this := ClusterExecutionLogItem{}
	return &this
}

// NewClusterExecutionLogItemWithDefaults instantiates a new ClusterExecutionLogItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterExecutionLogItemWithDefaults() *ClusterExecutionLogItem {
	this := ClusterExecutionLogItem{}
	return &this
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *ClusterExecutionLogItem) GetClient() string {
	if o == nil || o.Client == nil {
		var ret string
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLogItem) GetClientOk() (*string, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *ClusterExecutionLogItem) HasClient() bool {
	return o != nil && o.Client != nil
}

// SetClient gets a reference to the given string and assigns it to the Client field.
func (o *ClusterExecutionLogItem) SetClient(v string) {
	o.Client = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *ClusterExecutionLogItem) GetCommand() string {
	if o == nil || o.Command == nil {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLogItem) GetCommandOk() (*string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *ClusterExecutionLogItem) HasCommand() bool {
	return o != nil && o.Command != nil
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *ClusterExecutionLogItem) SetCommand(v string) {
	o.Command = &v
}

// GetDbName returns the DbName field value if set, zero value otherwise.
func (o *ClusterExecutionLogItem) GetDbName() string {
	if o == nil || o.DbName == nil {
		var ret string
		return ret
	}
	return *o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLogItem) GetDbNameOk() (*string, bool) {
	if o == nil || o.DbName == nil {
		return nil, false
	}
	return o.DbName, true
}

// HasDbName returns a boolean if a field has been set.
func (o *ClusterExecutionLogItem) HasDbName() bool {
	return o != nil && o.DbName != nil
}

// SetDbName gets a reference to the given string and assigns it to the DbName field.
func (o *ClusterExecutionLogItem) SetDbName(v string) {
	o.DbName = &v
}

// GetExecutionTime returns the ExecutionTime field value if set, zero value otherwise.
func (o *ClusterExecutionLogItem) GetExecutionTime() float64 {
	if o == nil || o.ExecutionTime == nil {
		var ret float64
		return ret
	}
	return *o.ExecutionTime
}

// GetExecutionTimeOk returns a tuple with the ExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLogItem) GetExecutionTimeOk() (*float64, bool) {
	if o == nil || o.ExecutionTime == nil {
		return nil, false
	}
	return o.ExecutionTime, true
}

// HasExecutionTime returns a boolean if a field has been set.
func (o *ClusterExecutionLogItem) HasExecutionTime() bool {
	return o != nil && o.ExecutionTime != nil
}

// SetExecutionTime gets a reference to the given float64 and assigns it to the ExecutionTime field.
func (o *ClusterExecutionLogItem) SetExecutionTime(v float64) {
	o.ExecutionTime = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *ClusterExecutionLogItem) GetExtra() map[string]interface{} {
	if o == nil || o.Extra == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLogItem) GetExtraOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return &o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *ClusterExecutionLogItem) HasExtra() bool {
	return o != nil && o.Extra != nil
}

// SetExtra gets a reference to the given map[string]interface{} and assigns it to the Extra field.
func (o *ClusterExecutionLogItem) SetExtra(v map[string]interface{}) {
	o.Extra = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ClusterExecutionLogItem) GetTimestamp() int32 {
	if o == nil || o.Timestamp == nil {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLogItem) GetTimestampOk() (*int32, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ClusterExecutionLogItem) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *ClusterExecutionLogItem) SetTimestamp(v int32) {
	o.Timestamp = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ClusterExecutionLogItem) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLogItem) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ClusterExecutionLogItem) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ClusterExecutionLogItem) SetUser(v string) {
	o.User = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterExecutionLogItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Client != nil {
		toSerialize["client"] = o.Client
	}
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	if o.DbName != nil {
		toSerialize["dbName"] = o.DbName
	}
	if o.ExecutionTime != nil {
		toSerialize["executionTime"] = o.ExecutionTime
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterExecutionLogItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Client        *string                `json:"client,omitempty"`
		Command       *string                `json:"command,omitempty"`
		DbName        *string                `json:"dbName,omitempty"`
		ExecutionTime *float64               `json:"executionTime,omitempty"`
		Extra         map[string]interface{} `json:"extra,omitempty"`
		Timestamp     *int32                 `json:"timestamp,omitempty"`
		User          *string                `json:"user,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"client", "command", "dbName", "executionTime", "extra", "timestamp", "user"})
	} else {
		return err
	}
	o.Client = all.Client
	o.Command = all.Command
	o.DbName = all.DbName
	o.ExecutionTime = all.ExecutionTime
	o.Extra = all.Extra
	o.Timestamp = all.Timestamp
	o.User = all.User

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
