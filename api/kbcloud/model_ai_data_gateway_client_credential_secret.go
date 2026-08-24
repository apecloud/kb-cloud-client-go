// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiDataGatewayClientCredentialSecret struct {
	// Public access key ID used to identify this client credential.
	AccessKeyId   *string    `json:"accessKeyId,omitempty"`
	GatewayId     *string    `json:"gatewayId,omitempty"`
	OrgName       *string    `json:"orgName,omitempty"`
	Name          *string    `json:"name,omitempty"`
	Description   *string    `json:"description,omitempty"`
	Status        *string    `json:"status,omitempty"`
	CreatedBy     *string    `json:"createdBy,omitempty"`
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
	LastRotatedAt *time.Time `json:"lastRotatedAt,omitempty"`
	LastUsedAt    *time.Time `json:"lastUsedAt,omitempty"`
	DisabledAt    *time.Time `json:"disabledAt,omitempty"`
	RevokedAt     *time.Time `json:"revokedAt,omitempty"`
	// One-time returned secret key. It is never returned by list, get, runtime config, or audit APIs.
	SecretKey *string `json:"secretKey,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiDataGatewayClientCredentialSecret instantiates a new AiDataGatewayClientCredentialSecret object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiDataGatewayClientCredentialSecret() *AiDataGatewayClientCredentialSecret {
	this := AiDataGatewayClientCredentialSecret{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *AiDataGatewayClientCredentialSecret) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayClientCredentialSecret) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *AiDataGatewayClientCredentialSecret) HasAccessKeyId() bool {
	return o != nil && o.AccessKeyId != nil
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *AiDataGatewayClientCredentialSecret) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *AiDataGatewayClientCredentialSecret) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayClientCredentialSecret) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *AiDataGatewayClientCredentialSecret) HasGatewayId() bool {
	return o != nil && o.GatewayId != nil
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *AiDataGatewayClientCredentialSecret) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *AiDataGatewayClientCredentialSecret) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayClientCredentialSecret) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *AiDataGatewayClientCredentialSecret) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *AiDataGatewayClientCredentialSecret) SetOrgName(v string) {
	o.OrgName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AiDataGatewayClientCredentialSecret) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayClientCredentialSecret) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AiDataGatewayClientCredentialSecret) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AiDataGatewayClientCredentialSecret) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AiDataGatewayClientCredentialSecret) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayClientCredentialSecret) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AiDataGatewayClientCredentialSecret) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AiDataGatewayClientCredentialSecret) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiDataGatewayClientCredentialSecret) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayClientCredentialSecret) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiDataGatewayClientCredentialSecret) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AiDataGatewayClientCredentialSecret) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *AiDataGatewayClientCredentialSecret) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayClientCredentialSecret) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AiDataGatewayClientCredentialSecret) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *AiDataGatewayClientCredentialSecret) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AiDataGatewayClientCredentialSecret) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayClientCredentialSecret) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AiDataGatewayClientCredentialSecret) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AiDataGatewayClientCredentialSecret) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AiDataGatewayClientCredentialSecret) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayClientCredentialSecret) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AiDataGatewayClientCredentialSecret) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AiDataGatewayClientCredentialSecret) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLastRotatedAt returns the LastRotatedAt field value if set, zero value otherwise.
func (o *AiDataGatewayClientCredentialSecret) GetLastRotatedAt() time.Time {
	if o == nil || o.LastRotatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastRotatedAt
}

// GetLastRotatedAtOk returns a tuple with the LastRotatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayClientCredentialSecret) GetLastRotatedAtOk() (*time.Time, bool) {
	if o == nil || o.LastRotatedAt == nil {
		return nil, false
	}
	return o.LastRotatedAt, true
}

// HasLastRotatedAt returns a boolean if a field has been set.
func (o *AiDataGatewayClientCredentialSecret) HasLastRotatedAt() bool {
	return o != nil && o.LastRotatedAt != nil
}

// SetLastRotatedAt gets a reference to the given time.Time and assigns it to the LastRotatedAt field.
func (o *AiDataGatewayClientCredentialSecret) SetLastRotatedAt(v time.Time) {
	o.LastRotatedAt = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *AiDataGatewayClientCredentialSecret) GetLastUsedAt() time.Time {
	if o == nil || o.LastUsedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayClientCredentialSecret) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil || o.LastUsedAt == nil {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *AiDataGatewayClientCredentialSecret) HasLastUsedAt() bool {
	return o != nil && o.LastUsedAt != nil
}

// SetLastUsedAt gets a reference to the given time.Time and assigns it to the LastUsedAt field.
func (o *AiDataGatewayClientCredentialSecret) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = &v
}

// GetDisabledAt returns the DisabledAt field value if set, zero value otherwise.
func (o *AiDataGatewayClientCredentialSecret) GetDisabledAt() time.Time {
	if o == nil || o.DisabledAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayClientCredentialSecret) GetDisabledAtOk() (*time.Time, bool) {
	if o == nil || o.DisabledAt == nil {
		return nil, false
	}
	return o.DisabledAt, true
}

// HasDisabledAt returns a boolean if a field has been set.
func (o *AiDataGatewayClientCredentialSecret) HasDisabledAt() bool {
	return o != nil && o.DisabledAt != nil
}

// SetDisabledAt gets a reference to the given time.Time and assigns it to the DisabledAt field.
func (o *AiDataGatewayClientCredentialSecret) SetDisabledAt(v time.Time) {
	o.DisabledAt = &v
}

// GetRevokedAt returns the RevokedAt field value if set, zero value otherwise.
func (o *AiDataGatewayClientCredentialSecret) GetRevokedAt() time.Time {
	if o == nil || o.RevokedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.RevokedAt
}

// GetRevokedAtOk returns a tuple with the RevokedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayClientCredentialSecret) GetRevokedAtOk() (*time.Time, bool) {
	if o == nil || o.RevokedAt == nil {
		return nil, false
	}
	return o.RevokedAt, true
}

// HasRevokedAt returns a boolean if a field has been set.
func (o *AiDataGatewayClientCredentialSecret) HasRevokedAt() bool {
	return o != nil && o.RevokedAt != nil
}

// SetRevokedAt gets a reference to the given time.Time and assigns it to the RevokedAt field.
func (o *AiDataGatewayClientCredentialSecret) SetRevokedAt(v time.Time) {
	o.RevokedAt = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *AiDataGatewayClientCredentialSecret) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayClientCredentialSecret) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *AiDataGatewayClientCredentialSecret) HasSecretKey() bool {
	return o != nil && o.SecretKey != nil
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *AiDataGatewayClientCredentialSecret) SetSecretKey(v string) {
	o.SecretKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiDataGatewayClientCredentialSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.AccessKeyId != nil {
		toSerialize["accessKeyId"] = o.AccessKeyId
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
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
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
	if o.LastRotatedAt != nil {
		if o.LastRotatedAt.Nanosecond() == 0 {
			toSerialize["lastRotatedAt"] = o.LastRotatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["lastRotatedAt"] = o.LastRotatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.LastUsedAt != nil {
		if o.LastUsedAt.Nanosecond() == 0 {
			toSerialize["lastUsedAt"] = o.LastUsedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["lastUsedAt"] = o.LastUsedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.DisabledAt != nil {
		if o.DisabledAt.Nanosecond() == 0 {
			toSerialize["disabledAt"] = o.DisabledAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["disabledAt"] = o.DisabledAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.RevokedAt != nil {
		if o.RevokedAt.Nanosecond() == 0 {
			toSerialize["revokedAt"] = o.RevokedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["revokedAt"] = o.RevokedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.SecretKey != nil {
		toSerialize["secretKey"] = o.SecretKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiDataGatewayClientCredentialSecret) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessKeyId   *string    `json:"accessKeyId,omitempty"`
		GatewayId     *string    `json:"gatewayId,omitempty"`
		OrgName       *string    `json:"orgName,omitempty"`
		Name          *string    `json:"name,omitempty"`
		Description   *string    `json:"description,omitempty"`
		Status        *string    `json:"status,omitempty"`
		CreatedBy     *string    `json:"createdBy,omitempty"`
		CreatedAt     *time.Time `json:"createdAt,omitempty"`
		UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
		LastRotatedAt *time.Time `json:"lastRotatedAt,omitempty"`
		LastUsedAt    *time.Time `json:"lastUsedAt,omitempty"`
		DisabledAt    *time.Time `json:"disabledAt,omitempty"`
		RevokedAt     *time.Time `json:"revokedAt,omitempty"`
		SecretKey     *string    `json:"secretKey,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"accessKeyId", "gatewayId", "orgName", "name", "description", "status", "createdBy", "createdAt", "updatedAt", "lastRotatedAt", "lastUsedAt", "disabledAt", "revokedAt", "secretKey"})
	} else {
		return err
	}
	o.AccessKeyId = all.AccessKeyId
	o.GatewayId = all.GatewayId
	o.OrgName = all.OrgName
	o.Name = all.Name
	o.Description = all.Description
	o.Status = all.Status
	o.CreatedBy = all.CreatedBy
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt
	o.LastRotatedAt = all.LastRotatedAt
	o.LastUsedAt = all.LastUsedAt
	o.DisabledAt = all.DisabledAt
	o.RevokedAt = all.RevokedAt
	o.SecretKey = all.SecretKey
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
