// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiDataGateway struct {
	GatewayId     *string    `json:"gatewayId,omitempty"`
	OrgName       *string    `json:"orgName,omitempty"`
	Name          *string    `json:"name,omitempty"`
	Description   *string    `json:"description,omitempty"`
	Status        *string    `json:"status,omitempty"`
	ConfigVersion *int64     `json:"configVersion,omitempty"`
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiDataGateway instantiates a new AiDataGateway object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiDataGateway() *AiDataGateway {
	this := AiDataGateway{}
	return &this
}

// NewAiDataGatewayWithDefaults instantiates a new AiDataGateway object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiDataGatewayWithDefaults() *AiDataGateway {
	this := AiDataGateway{}
	return &this
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *AiDataGateway) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGateway) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *AiDataGateway) HasGatewayId() bool {
	return o != nil && o.GatewayId != nil
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *AiDataGateway) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *AiDataGateway) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGateway) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *AiDataGateway) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *AiDataGateway) SetOrgName(v string) {
	o.OrgName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AiDataGateway) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGateway) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AiDataGateway) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AiDataGateway) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AiDataGateway) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGateway) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AiDataGateway) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AiDataGateway) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiDataGateway) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGateway) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiDataGateway) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AiDataGateway) SetStatus(v string) {
	o.Status = &v
}

// GetConfigVersion returns the ConfigVersion field value if set, zero value otherwise.
func (o *AiDataGateway) GetConfigVersion() int64 {
	if o == nil || o.ConfigVersion == nil {
		var ret int64
		return ret
	}
	return *o.ConfigVersion
}

// GetConfigVersionOk returns a tuple with the ConfigVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGateway) GetConfigVersionOk() (*int64, bool) {
	if o == nil || o.ConfigVersion == nil {
		return nil, false
	}
	return o.ConfigVersion, true
}

// HasConfigVersion returns a boolean if a field has been set.
func (o *AiDataGateway) HasConfigVersion() bool {
	return o != nil && o.ConfigVersion != nil
}

// SetConfigVersion gets a reference to the given int64 and assigns it to the ConfigVersion field.
func (o *AiDataGateway) SetConfigVersion(v int64) {
	o.ConfigVersion = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AiDataGateway) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGateway) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AiDataGateway) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AiDataGateway) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AiDataGateway) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGateway) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AiDataGateway) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AiDataGateway) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiDataGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.GatewayId != nil {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ConfigVersion != nil {
		toSerialize["configVersion"] = o.ConfigVersion
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiDataGateway) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GatewayId     *string    `json:"gatewayId,omitempty"`
		OrgName       *string    `json:"orgName,omitempty"`
		Name          *string    `json:"name,omitempty"`
		Description   *string    `json:"description,omitempty"`
		Status        *string    `json:"status,omitempty"`
		ConfigVersion *int64     `json:"configVersion,omitempty"`
		CreatedAt     *time.Time `json:"createdAt,omitempty"`
		UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"gatewayId", "orgName", "name", "description", "status", "configVersion", "createdAt", "updatedAt"})
	} else {
		return err
	}
	o.GatewayId = all.GatewayId
	o.OrgName = all.OrgName
	o.Name = all.Name
	o.Description = all.Description
	o.Status = all.Status
	o.ConfigVersion = all.ConfigVersion
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
