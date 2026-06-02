// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
	"github.com/google/uuid"
)

type EngineNetworkMode struct {
	// record id
	Id string `json:"id"`
	// engine name
	EngineName string `json:"engineName"`
	// engine mode
	Mode string `json:"mode"`
	// network modes configured for this engine+mode+env combination
	NetworkModes []string `json:"networkModes"`
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

// NewEngineNetworkMode instantiates a new EngineNetworkMode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineNetworkMode(id string, engineName string, mode string, networkModes []string, envId uuid.UUID, envName string) *EngineNetworkMode {
	this := EngineNetworkMode{}
	this.Id = id
	this.EngineName = engineName
	this.Mode = mode
	this.NetworkModes = networkModes
	this.EnvId = envId
	this.EnvName = envName
	return &this
}

// NewEngineNetworkModeWithDefaults instantiates a new EngineNetworkMode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineNetworkModeWithDefaults() *EngineNetworkMode {
	this := EngineNetworkMode{}
	return &this
}

// GetId returns the Id field value.
func (o *EngineNetworkMode) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EngineNetworkMode) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *EngineNetworkMode) SetId(v string) {
	o.Id = v
}

// GetEngineName returns the EngineName field value.
func (o *EngineNetworkMode) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *EngineNetworkMode) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *EngineNetworkMode) SetEngineName(v string) {
	o.EngineName = v
}

// GetMode returns the Mode field value.
func (o *EngineNetworkMode) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *EngineNetworkMode) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *EngineNetworkMode) SetMode(v string) {
	o.Mode = v
}

// GetNetworkModes returns the NetworkModes field value.
func (o *EngineNetworkMode) GetNetworkModes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NetworkModes
}

// GetNetworkModesOk returns a tuple with the NetworkModes field value
// and a boolean to check if the value has been set.
func (o *EngineNetworkMode) GetNetworkModesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkModes, true
}

// SetNetworkModes sets field value.
func (o *EngineNetworkMode) SetNetworkModes(v []string) {
	o.NetworkModes = v
}

// GetEnvId returns the EnvId field value.
func (o *EngineNetworkMode) GetEnvId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.EnvId
}

// GetEnvIdOk returns a tuple with the EnvId field value
// and a boolean to check if the value has been set.
func (o *EngineNetworkMode) GetEnvIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvId, true
}

// SetEnvId sets field value.
func (o *EngineNetworkMode) SetEnvId(v uuid.UUID) {
	o.EnvId = v
}

// GetEnvName returns the EnvName field value.
func (o *EngineNetworkMode) GetEnvName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnvName
}

// GetEnvNameOk returns a tuple with the EnvName field value
// and a boolean to check if the value has been set.
func (o *EngineNetworkMode) GetEnvNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvName, true
}

// SetEnvName sets field value.
func (o *EngineNetworkMode) SetEnvName(v string) {
	o.EnvName = v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *EngineNetworkMode) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineNetworkMode) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *EngineNetworkMode) HasCreateTime() bool {
	return o != nil && o.CreateTime != nil
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *EngineNetworkMode) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *EngineNetworkMode) GetUpdateTime() time.Time {
	if o == nil || o.UpdateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineNetworkMode) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *EngineNetworkMode) HasUpdateTime() bool {
	return o != nil && o.UpdateTime != nil
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *EngineNetworkMode) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineNetworkMode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["engineName"] = o.EngineName
	toSerialize["mode"] = o.Mode
	toSerialize["networkModes"] = o.NetworkModes
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
func (o *EngineNetworkMode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id           *string    `json:"id"`
		EngineName   *string    `json:"engineName"`
		Mode         *string    `json:"mode"`
		NetworkModes *[]string  `json:"networkModes"`
		EnvId        *uuid.UUID `json:"envId"`
		EnvName      *string    `json:"envName"`
		CreateTime   *time.Time `json:"createTime,omitempty"`
		UpdateTime   *time.Time `json:"updateTime,omitempty"`
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
	if all.NetworkModes == nil {
		return fmt.Errorf("required field networkModes missing")
	}
	if all.EnvId == nil {
		return fmt.Errorf("required field envId missing")
	}
	if all.EnvName == nil {
		return fmt.Errorf("required field envName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "engineName", "mode", "networkModes", "envId", "envName", "createTime", "updateTime"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.EngineName = *all.EngineName
	o.Mode = *all.Mode
	o.NetworkModes = *all.NetworkModes
	o.EnvId = *all.EnvId
	o.EnvName = *all.EnvName
	o.CreateTime = all.CreateTime
	o.UpdateTime = all.UpdateTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
