// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ClickhouseMutation struct {
	// Database name.
	Database *string `json:"database,omitempty"`
	// Table name.
	Table *string `json:"table,omitempty"`
	// ClickHouse server host running the mutation.
	HostName *string `json:"hostName,omitempty"`
	// Mutation identifier.
	MutationId *string `json:"mutationId,omitempty"`
	// Mutation command (ALTER/UPDATE/DELETE) with string and numeric literals and comments masked as ?.
	Command *string `json:"command,omitempty"`
	// Mutation creation time in RFC3339 UTC.
	CreateTime *string `json:"createTime,omitempty"`
	// Parts still waiting to be mutated.
	PartsToDo *int64 `json:"partsToDo,omitempty"`
	// Most recent part that failed mutation.
	LatestFailedPart *string `json:"latestFailedPart,omitempty"`
	// Most recent mutation failure time in RFC3339 UTC.
	LatestFailTime *string `json:"latestFailTime,omitempty"`
	// Most recent mutation failure reason.
	LatestFailReason *string `json:"latestFailReason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseMutation instantiates a new ClickhouseMutation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseMutation() *ClickhouseMutation {
	this := ClickhouseMutation{}
	return &this
}

// NewClickhouseMutationWithDefaults instantiates a new ClickhouseMutation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseMutationWithDefaults() *ClickhouseMutation {
	this := ClickhouseMutation{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ClickhouseMutation) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMutation) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ClickhouseMutation) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *ClickhouseMutation) SetDatabase(v string) {
	o.Database = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *ClickhouseMutation) GetTable() string {
	if o == nil || o.Table == nil {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMutation) GetTableOk() (*string, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *ClickhouseMutation) HasTable() bool {
	return o != nil && o.Table != nil
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *ClickhouseMutation) SetTable(v string) {
	o.Table = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *ClickhouseMutation) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMutation) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *ClickhouseMutation) HasHostName() bool {
	return o != nil && o.HostName != nil
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *ClickhouseMutation) SetHostName(v string) {
	o.HostName = &v
}

// GetMutationId returns the MutationId field value if set, zero value otherwise.
func (o *ClickhouseMutation) GetMutationId() string {
	if o == nil || o.MutationId == nil {
		var ret string
		return ret
	}
	return *o.MutationId
}

// GetMutationIdOk returns a tuple with the MutationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMutation) GetMutationIdOk() (*string, bool) {
	if o == nil || o.MutationId == nil {
		return nil, false
	}
	return o.MutationId, true
}

// HasMutationId returns a boolean if a field has been set.
func (o *ClickhouseMutation) HasMutationId() bool {
	return o != nil && o.MutationId != nil
}

// SetMutationId gets a reference to the given string and assigns it to the MutationId field.
func (o *ClickhouseMutation) SetMutationId(v string) {
	o.MutationId = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *ClickhouseMutation) GetCommand() string {
	if o == nil || o.Command == nil {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMutation) GetCommandOk() (*string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *ClickhouseMutation) HasCommand() bool {
	return o != nil && o.Command != nil
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *ClickhouseMutation) SetCommand(v string) {
	o.Command = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *ClickhouseMutation) GetCreateTime() string {
	if o == nil || o.CreateTime == nil {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMutation) GetCreateTimeOk() (*string, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *ClickhouseMutation) HasCreateTime() bool {
	return o != nil && o.CreateTime != nil
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *ClickhouseMutation) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetPartsToDo returns the PartsToDo field value if set, zero value otherwise.
func (o *ClickhouseMutation) GetPartsToDo() int64 {
	if o == nil || o.PartsToDo == nil {
		var ret int64
		return ret
	}
	return *o.PartsToDo
}

// GetPartsToDoOk returns a tuple with the PartsToDo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMutation) GetPartsToDoOk() (*int64, bool) {
	if o == nil || o.PartsToDo == nil {
		return nil, false
	}
	return o.PartsToDo, true
}

// HasPartsToDo returns a boolean if a field has been set.
func (o *ClickhouseMutation) HasPartsToDo() bool {
	return o != nil && o.PartsToDo != nil
}

// SetPartsToDo gets a reference to the given int64 and assigns it to the PartsToDo field.
func (o *ClickhouseMutation) SetPartsToDo(v int64) {
	o.PartsToDo = &v
}

// GetLatestFailedPart returns the LatestFailedPart field value if set, zero value otherwise.
func (o *ClickhouseMutation) GetLatestFailedPart() string {
	if o == nil || o.LatestFailedPart == nil {
		var ret string
		return ret
	}
	return *o.LatestFailedPart
}

// GetLatestFailedPartOk returns a tuple with the LatestFailedPart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMutation) GetLatestFailedPartOk() (*string, bool) {
	if o == nil || o.LatestFailedPart == nil {
		return nil, false
	}
	return o.LatestFailedPart, true
}

// HasLatestFailedPart returns a boolean if a field has been set.
func (o *ClickhouseMutation) HasLatestFailedPart() bool {
	return o != nil && o.LatestFailedPart != nil
}

// SetLatestFailedPart gets a reference to the given string and assigns it to the LatestFailedPart field.
func (o *ClickhouseMutation) SetLatestFailedPart(v string) {
	o.LatestFailedPart = &v
}

// GetLatestFailTime returns the LatestFailTime field value if set, zero value otherwise.
func (o *ClickhouseMutation) GetLatestFailTime() string {
	if o == nil || o.LatestFailTime == nil {
		var ret string
		return ret
	}
	return *o.LatestFailTime
}

// GetLatestFailTimeOk returns a tuple with the LatestFailTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMutation) GetLatestFailTimeOk() (*string, bool) {
	if o == nil || o.LatestFailTime == nil {
		return nil, false
	}
	return o.LatestFailTime, true
}

// HasLatestFailTime returns a boolean if a field has been set.
func (o *ClickhouseMutation) HasLatestFailTime() bool {
	return o != nil && o.LatestFailTime != nil
}

// SetLatestFailTime gets a reference to the given string and assigns it to the LatestFailTime field.
func (o *ClickhouseMutation) SetLatestFailTime(v string) {
	o.LatestFailTime = &v
}

// GetLatestFailReason returns the LatestFailReason field value if set, zero value otherwise.
func (o *ClickhouseMutation) GetLatestFailReason() string {
	if o == nil || o.LatestFailReason == nil {
		var ret string
		return ret
	}
	return *o.LatestFailReason
}

// GetLatestFailReasonOk returns a tuple with the LatestFailReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseMutation) GetLatestFailReasonOk() (*string, bool) {
	if o == nil || o.LatestFailReason == nil {
		return nil, false
	}
	return o.LatestFailReason, true
}

// HasLatestFailReason returns a boolean if a field has been set.
func (o *ClickhouseMutation) HasLatestFailReason() bool {
	return o != nil && o.LatestFailReason != nil
}

// SetLatestFailReason gets a reference to the given string and assigns it to the LatestFailReason field.
func (o *ClickhouseMutation) SetLatestFailReason(v string) {
	o.LatestFailReason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseMutation) MarshalJSON() ([]byte, error) {
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
	if o.MutationId != nil {
		toSerialize["mutationId"] = o.MutationId
	}
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.PartsToDo != nil {
		toSerialize["partsToDo"] = o.PartsToDo
	}
	if o.LatestFailedPart != nil {
		toSerialize["latestFailedPart"] = o.LatestFailedPart
	}
	if o.LatestFailTime != nil {
		toSerialize["latestFailTime"] = o.LatestFailTime
	}
	if o.LatestFailReason != nil {
		toSerialize["latestFailReason"] = o.LatestFailReason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseMutation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database         *string `json:"database,omitempty"`
		Table            *string `json:"table,omitempty"`
		HostName         *string `json:"hostName,omitempty"`
		MutationId       *string `json:"mutationId,omitempty"`
		Command          *string `json:"command,omitempty"`
		CreateTime       *string `json:"createTime,omitempty"`
		PartsToDo        *int64  `json:"partsToDo,omitempty"`
		LatestFailedPart *string `json:"latestFailedPart,omitempty"`
		LatestFailTime   *string `json:"latestFailTime,omitempty"`
		LatestFailReason *string `json:"latestFailReason,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "table", "hostName", "mutationId", "command", "createTime", "partsToDo", "latestFailedPart", "latestFailTime", "latestFailReason"})
	} else {
		return err
	}
	o.Database = all.Database
	o.Table = all.Table
	o.HostName = all.HostName
	o.MutationId = all.MutationId
	o.Command = all.Command
	o.CreateTime = all.CreateTime
	o.PartsToDo = all.PartsToDo
	o.LatestFailedPart = all.LatestFailedPart
	o.LatestFailTime = all.LatestFailTime
	o.LatestFailReason = all.LatestFailReason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
