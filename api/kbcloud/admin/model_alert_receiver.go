// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AlertReceiver Alert receiver information

type AlertReceiver struct {
	// NODESCRIPTION CreatedAt
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// NODESCRIPTION Id
	Id *string `json:"id,omitempty"`
	// NODESCRIPTION Name
	Name *string `json:"name,omitempty"`
	// NODESCRIPTION Category
	Category *AlertReceiverCategory `json:"category,omitempty"`
	// NODESCRIPTION UpdatedAt
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// NODESCRIPTION UserGroup
	UserGroup *AlertReceiverUserGroup `json:"userGroup,omitempty"`
	// Webhook config of alert receiver
	WebhookConfig *WebhookConfig `json:"webhookConfig,omitempty"`
	// NODESCRIPTION OrgName
	OrgName *string `json:"orgName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertReceiver instantiates a new AlertReceiver object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertReceiver() *AlertReceiver {
	this := AlertReceiver{}
	return &this
}

// NewAlertReceiverWithDefaults instantiates a new AlertReceiver object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertReceiverWithDefaults() *AlertReceiver {
	this := AlertReceiver{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AlertReceiver) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertReceiver) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AlertReceiver) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AlertReceiver) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertReceiver) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertReceiver) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertReceiver) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AlertReceiver) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlertReceiver) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertReceiver) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlertReceiver) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlertReceiver) SetName(v string) {
	o.Name = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AlertReceiver) GetCategory() AlertReceiverCategory {
	if o == nil || o.Category == nil {
		var ret AlertReceiverCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertReceiver) GetCategoryOk() (*AlertReceiverCategory, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AlertReceiver) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given AlertReceiverCategory and assigns it to the Category field.
func (o *AlertReceiver) SetCategory(v AlertReceiverCategory) {
	o.Category = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AlertReceiver) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertReceiver) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AlertReceiver) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AlertReceiver) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUserGroup returns the UserGroup field value if set, zero value otherwise.
func (o *AlertReceiver) GetUserGroup() AlertReceiverUserGroup {
	if o == nil || o.UserGroup == nil {
		var ret AlertReceiverUserGroup
		return ret
	}
	return *o.UserGroup
}

// GetUserGroupOk returns a tuple with the UserGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertReceiver) GetUserGroupOk() (*AlertReceiverUserGroup, bool) {
	if o == nil || o.UserGroup == nil {
		return nil, false
	}
	return o.UserGroup, true
}

// HasUserGroup returns a boolean if a field has been set.
func (o *AlertReceiver) HasUserGroup() bool {
	return o != nil && o.UserGroup != nil
}

// SetUserGroup gets a reference to the given AlertReceiverUserGroup and assigns it to the UserGroup field.
func (o *AlertReceiver) SetUserGroup(v AlertReceiverUserGroup) {
	o.UserGroup = &v
}

// GetWebhookConfig returns the WebhookConfig field value if set, zero value otherwise.
func (o *AlertReceiver) GetWebhookConfig() WebhookConfig {
	if o == nil || o.WebhookConfig == nil {
		var ret WebhookConfig
		return ret
	}
	return *o.WebhookConfig
}

// GetWebhookConfigOk returns a tuple with the WebhookConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertReceiver) GetWebhookConfigOk() (*WebhookConfig, bool) {
	if o == nil || o.WebhookConfig == nil {
		return nil, false
	}
	return o.WebhookConfig, true
}

// HasWebhookConfig returns a boolean if a field has been set.
func (o *AlertReceiver) HasWebhookConfig() bool {
	return o != nil && o.WebhookConfig != nil
}

// SetWebhookConfig gets a reference to the given WebhookConfig and assigns it to the WebhookConfig field.
func (o *AlertReceiver) SetWebhookConfig(v WebhookConfig) {
	o.WebhookConfig = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *AlertReceiver) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertReceiver) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *AlertReceiver) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *AlertReceiver) SetOrgName(v string) {
	o.OrgName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertReceiver) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UserGroup != nil {
		toSerialize["userGroup"] = o.UserGroup
	}
	if o.WebhookConfig != nil {
		toSerialize["webhookConfig"] = o.WebhookConfig
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertReceiver) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt     *time.Time              `json:"createdAt,omitempty"`
		Id            *string                 `json:"id,omitempty"`
		Name          *string                 `json:"name,omitempty"`
		Category      *AlertReceiverCategory  `json:"category,omitempty"`
		UpdatedAt     *time.Time              `json:"updatedAt,omitempty"`
		UserGroup     *AlertReceiverUserGroup `json:"userGroup,omitempty"`
		WebhookConfig *WebhookConfig          `json:"webhookConfig,omitempty"`
		OrgName       *string                 `json:"orgName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"createdAt", "id", "name", "category", "updatedAt", "userGroup", "webhookConfig", "orgName"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.Id = all.Id
	o.Name = all.Name
	if all.Category != nil && !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = all.Category
	}
	o.UpdatedAt = all.UpdatedAt
	if all.UserGroup != nil && all.UserGroup.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UserGroup = all.UserGroup
	if all.WebhookConfig != nil && all.WebhookConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.WebhookConfig = all.WebhookConfig
	o.OrgName = all.OrgName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
