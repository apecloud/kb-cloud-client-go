// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InfluxDBToken struct {
	Id          string               `json:"id"`
	UserId      string               `json:"userID"`
	OrgId       string               `json:"orgID"`
	Description string               `json:"description"`
	Status      InfluxDBTokenStatus  `json:"status"`
	Permissions []InfluxDBPermission `json:"permissions"`
	// Only explicit bucket-scoped tokens can be changed.
	Editable bool `json:"editable"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInfluxDBToken instantiates a new InfluxDBToken object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInfluxDBToken(id string, userId string, orgId string, description string, status InfluxDBTokenStatus, permissions []InfluxDBPermission, editable bool) *InfluxDBToken {
	this := InfluxDBToken{}
	this.Id = id
	this.UserId = userId
	this.OrgId = orgId
	this.Description = description
	this.Status = status
	this.Permissions = permissions
	this.Editable = editable
	return &this
}

// NewInfluxDBTokenWithDefaults instantiates a new InfluxDBToken object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInfluxDBTokenWithDefaults() *InfluxDBToken {
	this := InfluxDBToken{}
	return &this
}

// GetId returns the Id field value.
func (o *InfluxDBToken) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InfluxDBToken) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *InfluxDBToken) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value.
func (o *InfluxDBToken) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *InfluxDBToken) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value.
func (o *InfluxDBToken) SetUserId(v string) {
	o.UserId = v
}

// GetOrgId returns the OrgId field value.
func (o *InfluxDBToken) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *InfluxDBToken) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *InfluxDBToken) SetOrgId(v string) {
	o.OrgId = v
}

// GetDescription returns the Description field value.
func (o *InfluxDBToken) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InfluxDBToken) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *InfluxDBToken) SetDescription(v string) {
	o.Description = v
}

// GetStatus returns the Status field value.
func (o *InfluxDBToken) GetStatus() InfluxDBTokenStatus {
	if o == nil {
		var ret InfluxDBTokenStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InfluxDBToken) GetStatusOk() (*InfluxDBTokenStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *InfluxDBToken) SetStatus(v InfluxDBTokenStatus) {
	o.Status = v
}

// GetPermissions returns the Permissions field value.
func (o *InfluxDBToken) GetPermissions() []InfluxDBPermission {
	if o == nil {
		var ret []InfluxDBPermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *InfluxDBToken) GetPermissionsOk() (*[]InfluxDBPermission, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value.
func (o *InfluxDBToken) SetPermissions(v []InfluxDBPermission) {
	o.Permissions = v
}

// GetEditable returns the Editable field value.
func (o *InfluxDBToken) GetEditable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Editable
}

// GetEditableOk returns a tuple with the Editable field value
// and a boolean to check if the value has been set.
func (o *InfluxDBToken) GetEditableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Editable, true
}

// SetEditable sets field value.
func (o *InfluxDBToken) SetEditable(v bool) {
	o.Editable = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InfluxDBToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["userID"] = o.UserId
	toSerialize["orgID"] = o.OrgId
	toSerialize["description"] = o.Description
	toSerialize["status"] = o.Status
	toSerialize["permissions"] = o.Permissions
	toSerialize["editable"] = o.Editable

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InfluxDBToken) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *string               `json:"id"`
		UserId      *string               `json:"userID"`
		OrgId       *string               `json:"orgID"`
		Description *string               `json:"description"`
		Status      *InfluxDBTokenStatus  `json:"status"`
		Permissions *[]InfluxDBPermission `json:"permissions"`
		Editable    *bool                 `json:"editable"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.UserId == nil {
		return fmt.Errorf("required field userID missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field orgID missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Permissions == nil {
		return fmt.Errorf("required field permissions missing")
	}
	if all.Editable == nil {
		return fmt.Errorf("required field editable missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "userID", "orgID", "description", "status", "permissions", "editable"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.UserId = *all.UserId
	o.OrgId = *all.OrgId
	o.Description = *all.Description
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Permissions = *all.Permissions
	o.Editable = *all.Editable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
