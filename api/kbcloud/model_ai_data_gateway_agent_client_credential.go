// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiDataGatewayAgentClientCredential struct {
	ClientId  *string `json:"clientId,omitempty"`
	GatewayId *string `json:"gatewayId,omitempty"`
	OrgName   *string `json:"orgName,omitempty"`
	Name      *string `json:"name,omitempty"`
	Status    *string `json:"status,omitempty"`
	// Safe prefix of the client secret. The full secret is only returned on create or rotate.
	SecretPrefix  *string    `json:"secretPrefix,omitempty"`
	CreatedBy     *string    `json:"createdBy,omitempty"`
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
	LastRotatedAt *time.Time `json:"lastRotatedAt,omitempty"`
	DisabledAt    *time.Time `json:"disabledAt,omitempty"`
	// One-time returned client secret.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiDataGatewayAgentClientCredential instantiates a new AiDataGatewayAgentClientCredential object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiDataGatewayAgentClientCredential() *AiDataGatewayAgentClientCredential {
	this := AiDataGatewayAgentClientCredential{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AiDataGatewayAgentClientCredential) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayAgentClientCredential) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AiDataGatewayAgentClientCredential) HasClientId() bool {
	return o != nil && o.ClientId != nil
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AiDataGatewayAgentClientCredential) SetClientId(v string) {
	o.ClientId = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *AiDataGatewayAgentClientCredential) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayAgentClientCredential) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *AiDataGatewayAgentClientCredential) HasGatewayId() bool {
	return o != nil && o.GatewayId != nil
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *AiDataGatewayAgentClientCredential) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *AiDataGatewayAgentClientCredential) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayAgentClientCredential) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *AiDataGatewayAgentClientCredential) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *AiDataGatewayAgentClientCredential) SetOrgName(v string) {
	o.OrgName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AiDataGatewayAgentClientCredential) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayAgentClientCredential) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AiDataGatewayAgentClientCredential) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AiDataGatewayAgentClientCredential) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiDataGatewayAgentClientCredential) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayAgentClientCredential) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiDataGatewayAgentClientCredential) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AiDataGatewayAgentClientCredential) SetStatus(v string) {
	o.Status = &v
}

// GetSecretPrefix returns the SecretPrefix field value if set, zero value otherwise.
func (o *AiDataGatewayAgentClientCredential) GetSecretPrefix() string {
	if o == nil || o.SecretPrefix == nil {
		var ret string
		return ret
	}
	return *o.SecretPrefix
}

// GetSecretPrefixOk returns a tuple with the SecretPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayAgentClientCredential) GetSecretPrefixOk() (*string, bool) {
	if o == nil || o.SecretPrefix == nil {
		return nil, false
	}
	return o.SecretPrefix, true
}

// HasSecretPrefix returns a boolean if a field has been set.
func (o *AiDataGatewayAgentClientCredential) HasSecretPrefix() bool {
	return o != nil && o.SecretPrefix != nil
}

// SetSecretPrefix gets a reference to the given string and assigns it to the SecretPrefix field.
func (o *AiDataGatewayAgentClientCredential) SetSecretPrefix(v string) {
	o.SecretPrefix = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *AiDataGatewayAgentClientCredential) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayAgentClientCredential) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AiDataGatewayAgentClientCredential) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *AiDataGatewayAgentClientCredential) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AiDataGatewayAgentClientCredential) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayAgentClientCredential) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AiDataGatewayAgentClientCredential) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AiDataGatewayAgentClientCredential) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AiDataGatewayAgentClientCredential) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayAgentClientCredential) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AiDataGatewayAgentClientCredential) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AiDataGatewayAgentClientCredential) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLastRotatedAt returns the LastRotatedAt field value if set, zero value otherwise.
func (o *AiDataGatewayAgentClientCredential) GetLastRotatedAt() time.Time {
	if o == nil || o.LastRotatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastRotatedAt
}

// GetLastRotatedAtOk returns a tuple with the LastRotatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayAgentClientCredential) GetLastRotatedAtOk() (*time.Time, bool) {
	if o == nil || o.LastRotatedAt == nil {
		return nil, false
	}
	return o.LastRotatedAt, true
}

// HasLastRotatedAt returns a boolean if a field has been set.
func (o *AiDataGatewayAgentClientCredential) HasLastRotatedAt() bool {
	return o != nil && o.LastRotatedAt != nil
}

// SetLastRotatedAt gets a reference to the given time.Time and assigns it to the LastRotatedAt field.
func (o *AiDataGatewayAgentClientCredential) SetLastRotatedAt(v time.Time) {
	o.LastRotatedAt = &v
}

// GetDisabledAt returns the DisabledAt field value if set, zero value otherwise.
func (o *AiDataGatewayAgentClientCredential) GetDisabledAt() time.Time {
	if o == nil || o.DisabledAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayAgentClientCredential) GetDisabledAtOk() (*time.Time, bool) {
	if o == nil || o.DisabledAt == nil {
		return nil, false
	}
	return o.DisabledAt, true
}

// HasDisabledAt returns a boolean if a field has been set.
func (o *AiDataGatewayAgentClientCredential) HasDisabledAt() bool {
	return o != nil && o.DisabledAt != nil
}

// SetDisabledAt gets a reference to the given time.Time and assigns it to the DisabledAt field.
func (o *AiDataGatewayAgentClientCredential) SetDisabledAt(v time.Time) {
	o.DisabledAt = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *AiDataGatewayAgentClientCredential) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayAgentClientCredential) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *AiDataGatewayAgentClientCredential) HasClientSecret() bool {
	return o != nil && o.ClientSecret != nil
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *AiDataGatewayAgentClientCredential) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiDataGatewayAgentClientCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
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
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SecretPrefix != nil {
		toSerialize["secretPrefix"] = o.SecretPrefix
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
	if o.DisabledAt != nil {
		if o.DisabledAt.Nanosecond() == 0 {
			toSerialize["disabledAt"] = o.DisabledAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["disabledAt"] = o.DisabledAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ClientSecret != nil {
		toSerialize["clientSecret"] = o.ClientSecret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiDataGatewayAgentClientCredential) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClientId      *string    `json:"clientId,omitempty"`
		GatewayId     *string    `json:"gatewayId,omitempty"`
		OrgName       *string    `json:"orgName,omitempty"`
		Name          *string    `json:"name,omitempty"`
		Status        *string    `json:"status,omitempty"`
		SecretPrefix  *string    `json:"secretPrefix,omitempty"`
		CreatedBy     *string    `json:"createdBy,omitempty"`
		CreatedAt     *time.Time `json:"createdAt,omitempty"`
		UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
		LastRotatedAt *time.Time `json:"lastRotatedAt,omitempty"`
		DisabledAt    *time.Time `json:"disabledAt,omitempty"`
		ClientSecret  *string    `json:"clientSecret,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clientId", "gatewayId", "orgName", "name", "status", "secretPrefix", "createdBy", "createdAt", "updatedAt", "lastRotatedAt", "disabledAt", "clientSecret"})
	} else {
		return err
	}
	o.ClientId = all.ClientId
	o.GatewayId = all.GatewayId
	o.OrgName = all.OrgName
	o.Name = all.Name
	o.Status = all.Status
	o.SecretPrefix = all.SecretPrefix
	o.CreatedBy = all.CreatedBy
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt
	o.LastRotatedAt = all.LastRotatedAt
	o.DisabledAt = all.DisabledAt
	o.ClientSecret = all.ClientSecret
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
