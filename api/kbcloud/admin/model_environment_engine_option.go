// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EnvironmentEngineOption struct {
	// record id
	Id string `json:"id"`
	// engine name
	EngineName string `json:"engineName"`
	// engine mode
	Mode string `json:"mode"`
	// network modes configured for this engine+mode+env combination
	NetworkModes []string `json:"networkModes,omitempty"`
	// Number of data volume templates supported by this engine+mode in the environment.
	// For RustFS, this mirrors the addon installation value dataVolumeCount/maxDataVolumeCount
	// so the frontend can render data0..dataN volume inputs without changing installed addons.
	//
	DataVolumeCount common.NullableInt64 `json:"dataVolumeCount,omitempty"`
	// environment id
	EnvId uuid.UUID `json:"envId"`
	// environment name
	EnvName string `json:"envName"`
	// create time
	CreateTime *time.Time `json:"createTime,omitempty"`
	// update time
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentEngineOption instantiates a new EnvironmentEngineOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentEngineOption(id string, engineName string, mode string, envId uuid.UUID, envName string) *EnvironmentEngineOption {
	this := EnvironmentEngineOption{}
	this.Id = id
	this.EngineName = engineName
	this.Mode = mode
	this.EnvId = envId
	this.EnvName = envName
	return &this
}

// NewEnvironmentEngineOptionWithDefaults instantiates a new EnvironmentEngineOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentEngineOptionWithDefaults() *EnvironmentEngineOption {
	this := EnvironmentEngineOption{}
	return &this
}

// GetId returns the Id field value.
func (o *EnvironmentEngineOption) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvironmentEngineOption) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *EnvironmentEngineOption) SetId(v string) {
	o.Id = v
}

// GetEngineName returns the EngineName field value.
func (o *EnvironmentEngineOption) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *EnvironmentEngineOption) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *EnvironmentEngineOption) SetEngineName(v string) {
	o.EngineName = v
}

// GetMode returns the Mode field value.
func (o *EnvironmentEngineOption) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *EnvironmentEngineOption) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *EnvironmentEngineOption) SetMode(v string) {
	o.Mode = v
}

// GetNetworkModes returns the NetworkModes field value if set, zero value otherwise.
func (o *EnvironmentEngineOption) GetNetworkModes() []string {
	if o == nil || o.NetworkModes == nil {
		var ret []string
		return ret
	}
	return o.NetworkModes
}

// GetNetworkModesOk returns a tuple with the NetworkModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentEngineOption) GetNetworkModesOk() (*[]string, bool) {
	if o == nil || o.NetworkModes == nil {
		return nil, false
	}
	return &o.NetworkModes, true
}

// HasNetworkModes returns a boolean if a field has been set.
func (o *EnvironmentEngineOption) HasNetworkModes() bool {
	return o != nil && o.NetworkModes != nil
}

// SetNetworkModes gets a reference to the given []string and assigns it to the NetworkModes field.
func (o *EnvironmentEngineOption) SetNetworkModes(v []string) {
	o.NetworkModes = v
}

// GetDataVolumeCount returns the DataVolumeCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentEngineOption) GetDataVolumeCount() int64 {
	if o == nil || o.DataVolumeCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DataVolumeCount.Get()
}

// GetDataVolumeCountOk returns a tuple with the DataVolumeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentEngineOption) GetDataVolumeCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataVolumeCount.Get(), o.DataVolumeCount.IsSet()
}

// HasDataVolumeCount returns a boolean if a field has been set.
func (o *EnvironmentEngineOption) HasDataVolumeCount() bool {
	return o != nil && o.DataVolumeCount.IsSet()
}

// SetDataVolumeCount gets a reference to the given common.NullableInt64 and assigns it to the DataVolumeCount field.
func (o *EnvironmentEngineOption) SetDataVolumeCount(v int64) {
	o.DataVolumeCount.Set(&v)
}

// SetDataVolumeCountNil sets the value for DataVolumeCount to be an explicit nil.
func (o *EnvironmentEngineOption) SetDataVolumeCountNil() {
	o.DataVolumeCount.Set(nil)
}

// UnsetDataVolumeCount ensures that no value is present for DataVolumeCount, not even an explicit nil.
func (o *EnvironmentEngineOption) UnsetDataVolumeCount() {
	o.DataVolumeCount.Unset()
}

// GetEnvId returns the EnvId field value.
func (o *EnvironmentEngineOption) GetEnvId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.EnvId
}

// GetEnvIdOk returns a tuple with the EnvId field value
// and a boolean to check if the value has been set.
func (o *EnvironmentEngineOption) GetEnvIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvId, true
}

// SetEnvId sets field value.
func (o *EnvironmentEngineOption) SetEnvId(v uuid.UUID) {
	o.EnvId = v
}

// GetEnvName returns the EnvName field value.
func (o *EnvironmentEngineOption) GetEnvName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnvName
}

// GetEnvNameOk returns a tuple with the EnvName field value
// and a boolean to check if the value has been set.
func (o *EnvironmentEngineOption) GetEnvNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvName, true
}

// SetEnvName sets field value.
func (o *EnvironmentEngineOption) SetEnvName(v string) {
	o.EnvName = v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *EnvironmentEngineOption) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentEngineOption) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *EnvironmentEngineOption) HasCreateTime() bool {
	return o != nil && o.CreateTime != nil
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *EnvironmentEngineOption) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *EnvironmentEngineOption) GetUpdateTime() time.Time {
	if o == nil || o.UpdateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentEngineOption) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *EnvironmentEngineOption) HasUpdateTime() bool {
	return o != nil && o.UpdateTime != nil
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *EnvironmentEngineOption) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentEngineOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["engineName"] = o.EngineName
	toSerialize["mode"] = o.Mode
	if o.NetworkModes != nil {
		toSerialize["networkModes"] = o.NetworkModes
	}
	if o.DataVolumeCount.IsSet() {
		toSerialize["dataVolumeCount"] = o.DataVolumeCount.Get()
	}
	toSerialize["envId"] = o.EnvId
	toSerialize["envName"] = o.EnvName
	if o.CreateTime != nil {
		if o.CreateTime.Nanosecond() == 0 {
			toSerialize["createTime"] = o.CreateTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createTime"] = o.CreateTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdateTime != nil {
		if o.UpdateTime.Nanosecond() == 0 {
			toSerialize["updateTime"] = o.UpdateTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updateTime"] = o.UpdateTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentEngineOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id              *string              `json:"id"`
		EngineName      *string              `json:"engineName"`
		Mode            *string              `json:"mode"`
		NetworkModes    []string             `json:"networkModes,omitempty"`
		DataVolumeCount common.NullableInt64 `json:"dataVolumeCount,omitempty"`
		EnvId           *uuid.UUID           `json:"envId"`
		EnvName         *string              `json:"envName"`
		CreateTime      *time.Time           `json:"createTime,omitempty"`
		UpdateTime      *time.Time           `json:"updateTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.EngineName == nil {
		return fmt.Errorf("required field engineName missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.EnvId == nil {
		return fmt.Errorf("required field envId missing")
	}
	if all.EnvName == nil {
		return fmt.Errorf("required field envName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "engineName", "mode", "networkModes", "dataVolumeCount", "envId", "envName", "createTime", "updateTime"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.EngineName = *all.EngineName
	o.Mode = *all.Mode
	o.NetworkModes = all.NetworkModes
	o.DataVolumeCount = all.DataVolumeCount
	o.EnvId = *all.EnvId
	o.EnvName = *all.EnvName
	o.CreateTime = all.CreateTime
	o.UpdateTime = all.UpdateTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
