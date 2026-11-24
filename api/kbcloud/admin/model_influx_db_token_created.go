// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InfluxDBTokenCreated struct {
	Id          string               `json:"id"`
	UserId      string               `json:"userID"`
	OrgId       string               `json:"orgID"`
	Description string               `json:"description"`
	Status      InfluxDBTokenStatus  `json:"status"`
	Permissions []InfluxDBPermission `json:"permissions"`
	// Only explicit bucket-scoped tokens can be changed.
	Editable bool `json:"editable"`
	// New token secret, returned only at creation. Save it securely.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInfluxDBTokenCreated instantiates a new InfluxDBTokenCreated object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInfluxDBTokenCreated(id string, userId string, orgId string, description string, status InfluxDBTokenStatus, permissions []InfluxDBPermission, editable bool, value string) *InfluxDBTokenCreated {
	this := InfluxDBTokenCreated{}
	this.Id = id
	this.UserId = userId
	this.OrgId = orgId
	this.Description = description
	this.Status = status
	this.Permissions = permissions
	this.Editable = editable
	this.Value = value
	return &this
}

// NewInfluxDBTokenCreatedWithDefaults instantiates a new InfluxDBTokenCreated object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInfluxDBTokenCreatedWithDefaults() *InfluxDBTokenCreated {
	this := InfluxDBTokenCreated{}
	return &this
}

// GetId returns the Id field value.
func (o *InfluxDBTokenCreated) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InfluxDBTokenCreated) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *InfluxDBTokenCreated) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value.
func (o *InfluxDBTokenCreated) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *InfluxDBTokenCreated) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value.
func (o *InfluxDBTokenCreated) SetUserId(v string) {
	o.UserId = v
}

// GetOrgId returns the OrgId field value.
func (o *InfluxDBTokenCreated) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *InfluxDBTokenCreated) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *InfluxDBTokenCreated) SetOrgId(v string) {
	o.OrgId = v
}

// GetDescription returns the Description field value.
func (o *InfluxDBTokenCreated) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InfluxDBTokenCreated) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *InfluxDBTokenCreated) SetDescription(v string) {
	o.Description = v
}

// GetStatus returns the Status field value.
func (o *InfluxDBTokenCreated) GetStatus() InfluxDBTokenStatus {
	if o == nil {
		var ret InfluxDBTokenStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InfluxDBTokenCreated) GetStatusOk() (*InfluxDBTokenStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *InfluxDBTokenCreated) SetStatus(v InfluxDBTokenStatus) {
	o.Status = v
}

// GetPermissions returns the Permissions field value.
func (o *InfluxDBTokenCreated) GetPermissions() []InfluxDBPermission {
	if o == nil {
		var ret []InfluxDBPermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *InfluxDBTokenCreated) GetPermissionsOk() (*[]InfluxDBPermission, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value.
func (o *InfluxDBTokenCreated) SetPermissions(v []InfluxDBPermission) {
	o.Permissions = v
}

// GetEditable returns the Editable field value.
func (o *InfluxDBTokenCreated) GetEditable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Editable
}

// GetEditableOk returns a tuple with the Editable field value
// and a boolean to check if the value has been set.
func (o *InfluxDBTokenCreated) GetEditableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Editable, true
}

// SetEditable sets field value.
func (o *InfluxDBTokenCreated) SetEditable(v bool) {
	o.Editable = v
}

// GetValue returns the Value field value.
func (o *InfluxDBTokenCreated) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *InfluxDBTokenCreated) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *InfluxDBTokenCreated) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InfluxDBTokenCreated) MarshalJSON() ([]byte, error) {
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
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InfluxDBTokenCreated) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *string               `json:"id"`
		UserId      *string               `json:"userID"`
		OrgId       *string               `json:"orgID"`
		Description *string               `json:"description"`
		Status      *InfluxDBTokenStatus  `json:"status"`
		Permissions *[]InfluxDBPermission `json:"permissions"`
		Editable    *bool                 `json:"editable"`
		Value       *string               `json:"value"`
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
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "userID", "orgID", "description", "status", "permissions", "editable", "value"})
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
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
