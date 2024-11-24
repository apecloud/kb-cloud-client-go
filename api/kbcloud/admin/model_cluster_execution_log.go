// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// ClusterExecutionLog Log entry for cluster execution
type ClusterExecutionLog struct {
	// Timestamp of the execution
	Timestamp int64 `json:"timestamp"`
	// Client identifier
	Client string `json:"client"`
	// Database name
	DbName string `json:"dbName"`
	// User who executed the command
	User string `json:"user"`
	// Time taken for execution
	ExecutionTime float64 `json:"executionTime"`
	// Command that was executed
	Command string `json:"command"`
	// Additional information
	Extra map[string]interface{} `json:"extra"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterExecutionLog instantiates a new ClusterExecutionLog object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterExecutionLog(timestamp int64, client string, dbName string, user string, executionTime float64, command string, extra map[string]interface{}) *ClusterExecutionLog {
	this := ClusterExecutionLog{}
	this.Timestamp = timestamp
	this.Client = client
	this.DbName = dbName
	this.User = user
	this.ExecutionTime = executionTime
	this.Command = command
	this.Extra = extra
	return &this
}

// NewClusterExecutionLogWithDefaults instantiates a new ClusterExecutionLog object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterExecutionLogWithDefaults() *ClusterExecutionLog {
	this := ClusterExecutionLog{}
	return &this
}

// GetTimestamp returns the Timestamp field value.
func (o *ClusterExecutionLog) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLog) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *ClusterExecutionLog) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetClient returns the Client field value.
func (o *ClusterExecutionLog) GetClient() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Client
}

// GetClientOk returns a tuple with the Client field value
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLog) GetClientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Client, true
}

// SetClient sets field value.
func (o *ClusterExecutionLog) SetClient(v string) {
	o.Client = v
}

// GetDbName returns the DbName field value.
func (o *ClusterExecutionLog) GetDbName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLog) GetDbNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbName, true
}

// SetDbName sets field value.
func (o *ClusterExecutionLog) SetDbName(v string) {
	o.DbName = v
}

// GetUser returns the User field value.
func (o *ClusterExecutionLog) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLog) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value.
func (o *ClusterExecutionLog) SetUser(v string) {
	o.User = v
}

// GetExecutionTime returns the ExecutionTime field value.
func (o *ClusterExecutionLog) GetExecutionTime() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ExecutionTime
}

// GetExecutionTimeOk returns a tuple with the ExecutionTime field value
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLog) GetExecutionTimeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionTime, true
}

// SetExecutionTime sets field value.
func (o *ClusterExecutionLog) SetExecutionTime(v float64) {
	o.ExecutionTime = v
}

// GetCommand returns the Command field value.
func (o *ClusterExecutionLog) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLog) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value.
func (o *ClusterExecutionLog) SetCommand(v string) {
	o.Command = v
}

// GetExtra returns the Extra field value.
func (o *ClusterExecutionLog) GetExtra() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLog) GetExtraOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Extra, true
}

// SetExtra sets field value.
func (o *ClusterExecutionLog) SetExtra(v map[string]interface{}) {
	o.Extra = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterExecutionLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["client"] = o.Client
	toSerialize["dbName"] = o.DbName
	toSerialize["user"] = o.User
	toSerialize["executionTime"] = o.ExecutionTime
	toSerialize["command"] = o.Command
	toSerialize["extra"] = o.Extra

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterExecutionLog) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Timestamp     *int64                  `json:"timestamp"`
		Client        *string                 `json:"client"`
		DbName        *string                 `json:"dbName"`
		User          *string                 `json:"user"`
		ExecutionTime *float64                `json:"executionTime"`
		Command       *string                 `json:"command"`
		Extra         *map[string]interface{} `json:"extra"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Timestamp == nil {
		return fmt.Errorf("required field timestamp missing")
	}
	if all.Client == nil {
		return fmt.Errorf("required field client missing")
	}
	if all.DbName == nil {
		return fmt.Errorf("required field dbName missing")
	}
	if all.User == nil {
		return fmt.Errorf("required field user missing")
	}
	if all.ExecutionTime == nil {
		return fmt.Errorf("required field executionTime missing")
	}
	if all.Command == nil {
		return fmt.Errorf("required field command missing")
	}
	if all.Extra == nil {
		return fmt.Errorf("required field extra missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"timestamp", "client", "dbName", "user", "executionTime", "command", "extra"})
	} else {
		return err
	}
	o.Timestamp = *all.Timestamp
	o.Client = *all.Client
	o.DbName = *all.DbName
	o.User = *all.User
	o.ExecutionTime = *all.ExecutionTime
	o.Command = *all.Command
	o.Extra = *all.Extra

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
