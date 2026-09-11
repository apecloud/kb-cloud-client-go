// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MssqlLock struct {
	SessionId    int64  `json:"sessionId"`
	RequestId    int64  `json:"requestId"`
	Database     string `json:"database"`
	ResourceType string `json:"resourceType"`
	Resource     string `json:"resource"`
	EntityId     string `json:"entityId"`
	Mode         string `json:"mode"`
	Status       string `json:"status"`
	OwnerType    string `json:"ownerType"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMssqlLock instantiates a new MssqlLock object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMssqlLock(sessionId int64, requestId int64, database string, resourceType string, resource string, entityId string, mode string, status string, ownerType string) *MssqlLock {
	this := MssqlLock{}
	this.SessionId = sessionId
	this.RequestId = requestId
	this.Database = database
	this.ResourceType = resourceType
	this.Resource = resource
	this.EntityId = entityId
	this.Mode = mode
	this.Status = status
	this.OwnerType = ownerType
	return &this
}

// NewMssqlLockWithDefaults instantiates a new MssqlLock object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMssqlLockWithDefaults() *MssqlLock {
	this := MssqlLock{}
	return &this
}

// GetSessionId returns the SessionId field value.
func (o *MssqlLock) GetSessionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *MssqlLock) GetSessionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value.
func (o *MssqlLock) SetSessionId(v int64) {
	o.SessionId = v
}

// GetRequestId returns the RequestId field value.
func (o *MssqlLock) GetRequestId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *MssqlLock) GetRequestIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value.
func (o *MssqlLock) SetRequestId(v int64) {
	o.RequestId = v
}

// GetDatabase returns the Database field value.
func (o *MssqlLock) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *MssqlLock) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *MssqlLock) SetDatabase(v string) {
	o.Database = v
}

// GetResourceType returns the ResourceType field value.
func (o *MssqlLock) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *MssqlLock) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *MssqlLock) SetResourceType(v string) {
	o.ResourceType = v
}

// GetResource returns the Resource field value.
func (o *MssqlLock) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *MssqlLock) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value.
func (o *MssqlLock) SetResource(v string) {
	o.Resource = v
}

// GetEntityId returns the EntityId field value.
func (o *MssqlLock) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *MssqlLock) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value.
func (o *MssqlLock) SetEntityId(v string) {
	o.EntityId = v
}

// GetMode returns the Mode field value.
func (o *MssqlLock) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *MssqlLock) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *MssqlLock) SetMode(v string) {
	o.Mode = v
}

// GetStatus returns the Status field value.
func (o *MssqlLock) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MssqlLock) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *MssqlLock) SetStatus(v string) {
	o.Status = v
}

// GetOwnerType returns the OwnerType field value.
func (o *MssqlLock) GetOwnerType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value
// and a boolean to check if the value has been set.
func (o *MssqlLock) GetOwnerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerType, true
}

// SetOwnerType sets field value.
func (o *MssqlLock) SetOwnerType(v string) {
	o.OwnerType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MssqlLock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["sessionId"] = o.SessionId
	toSerialize["requestId"] = o.RequestId
	toSerialize["database"] = o.Database
	toSerialize["resourceType"] = o.ResourceType
	toSerialize["resource"] = o.Resource
	toSerialize["entityId"] = o.EntityId
	toSerialize["mode"] = o.Mode
	toSerialize["status"] = o.Status
	toSerialize["ownerType"] = o.OwnerType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MssqlLock) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SessionId    *int64  `json:"sessionId"`
		RequestId    *int64  `json:"requestId"`
		Database     *string `json:"database"`
		ResourceType *string `json:"resourceType"`
		Resource     *string `json:"resource"`
		EntityId     *string `json:"entityId"`
		Mode         *string `json:"mode"`
		Status       *string `json:"status"`
		OwnerType    *string `json:"ownerType"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SessionId == nil {
		return fmt.Errorf("required field sessionId missing")
	}
	if all.RequestId == nil {
		return fmt.Errorf("required field requestId missing")
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	if all.ResourceType == nil {
		return fmt.Errorf("required field resourceType missing")
	}
	if all.Resource == nil {
		return fmt.Errorf("required field resource missing")
	}
	if all.EntityId == nil {
		return fmt.Errorf("required field entityId missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.OwnerType == nil {
		return fmt.Errorf("required field ownerType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sessionId", "requestId", "database", "resourceType", "resource", "entityId", "mode", "status", "ownerType"})
	} else {
		return err
	}
	o.SessionId = *all.SessionId
	o.RequestId = *all.RequestId
	o.Database = *all.Database
	o.ResourceType = *all.ResourceType
	o.Resource = *all.Resource
	o.EntityId = *all.EntityId
	o.Mode = *all.Mode
	o.Status = *all.Status
	o.OwnerType = *all.OwnerType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
