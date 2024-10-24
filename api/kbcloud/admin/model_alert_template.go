// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AlertTemplate Alert template
// NODESCRIPTION AlertTemplate
//
// Deprecated: This model is deprecated.
type AlertTemplate struct {
	// NODESCRIPTION CreatedAt
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// NODESCRIPTION Description
	Description *string `json:"description,omitempty"`
	// NODESCRIPTION EmailSubject
	EmailSubject *string `json:"emailSubject,omitempty"`
	// NODESCRIPTION EmailText
	EmailText *string `json:"emailText,omitempty"`
	// NODESCRIPTION FeishuText
	FeishuText *string `json:"feishuText,omitempty"`
	// NODESCRIPTION FeishuTitle
	FeishuTitle *string `json:"feishuTitle,omitempty"`
	// NODESCRIPTION DingdingText
	DingdingText *string `json:"dingdingText,omitempty"`
	// NODESCRIPTION DingdingTitle
	DingdingTitle *string `json:"dingdingTitle,omitempty"`
	// NODESCRIPTION WeixinText
	WeixinText *string `json:"weixinText,omitempty"`
	// NODESCRIPTION WeixinTitle
	WeixinTitle *string `json:"weixinTitle,omitempty"`
	// NODESCRIPTION WebhookText
	WebhookText *string `json:"webhookText,omitempty"`
	// NODESCRIPTION WebhookTitle
	WebhookTitle *string `json:"webhookTitle,omitempty"`
	// NODESCRIPTION Id
	Id *string `json:"id,omitempty"`
	// NODESCRIPTION Name
	Name string `json:"name"`
	// NODESCRIPTION UpdatedAt
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// NODESCRIPTION IsDefault
	IsDefault *bool `json:"isDefault,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertTemplate instantiates a new AlertTemplate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertTemplate(name string) *AlertTemplate {
	this := AlertTemplate{}
	this.Name = name
	return &this
}

// NewAlertTemplateWithDefaults instantiates a new AlertTemplate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertTemplateWithDefaults() *AlertTemplate {
	this := AlertTemplate{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AlertTemplate) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AlertTemplate) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AlertTemplate) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlertTemplate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlertTemplate) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlertTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetEmailSubject returns the EmailSubject field value if set, zero value otherwise.
func (o *AlertTemplate) GetEmailSubject() string {
	if o == nil || o.EmailSubject == nil {
		var ret string
		return ret
	}
	return *o.EmailSubject
}

// GetEmailSubjectOk returns a tuple with the EmailSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplate) GetEmailSubjectOk() (*string, bool) {
	if o == nil || o.EmailSubject == nil {
		return nil, false
	}
	return o.EmailSubject, true
}

// HasEmailSubject returns a boolean if a field has been set.
func (o *AlertTemplate) HasEmailSubject() bool {
	return o != nil && o.EmailSubject != nil
}

// SetEmailSubject gets a reference to the given string and assigns it to the EmailSubject field.
func (o *AlertTemplate) SetEmailSubject(v string) {
	o.EmailSubject = &v
}

// GetEmailText returns the EmailText field value if set, zero value otherwise.
func (o *AlertTemplate) GetEmailText() string {
	if o == nil || o.EmailText == nil {
		var ret string
		return ret
	}
	return *o.EmailText
}

// GetEmailTextOk returns a tuple with the EmailText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplate) GetEmailTextOk() (*string, bool) {
	if o == nil || o.EmailText == nil {
		return nil, false
	}
	return o.EmailText, true
}

// HasEmailText returns a boolean if a field has been set.
func (o *AlertTemplate) HasEmailText() bool {
	return o != nil && o.EmailText != nil
}

// SetEmailText gets a reference to the given string and assigns it to the EmailText field.
func (o *AlertTemplate) SetEmailText(v string) {
	o.EmailText = &v
}

// GetFeishuText returns the FeishuText field value if set, zero value otherwise.
func (o *AlertTemplate) GetFeishuText() string {
	if o == nil || o.FeishuText == nil {
		var ret string
		return ret
	}
	return *o.FeishuText
}

// GetFeishuTextOk returns a tuple with the FeishuText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplate) GetFeishuTextOk() (*string, bool) {
	if o == nil || o.FeishuText == nil {
		return nil, false
	}
	return o.FeishuText, true
}

// HasFeishuText returns a boolean if a field has been set.
func (o *AlertTemplate) HasFeishuText() bool {
	return o != nil && o.FeishuText != nil
}

// SetFeishuText gets a reference to the given string and assigns it to the FeishuText field.
func (o *AlertTemplate) SetFeishuText(v string) {
	o.FeishuText = &v
}

// GetFeishuTitle returns the FeishuTitle field value if set, zero value otherwise.
func (o *AlertTemplate) GetFeishuTitle() string {
	if o == nil || o.FeishuTitle == nil {
		var ret string
		return ret
	}
	return *o.FeishuTitle
}

// GetFeishuTitleOk returns a tuple with the FeishuTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplate) GetFeishuTitleOk() (*string, bool) {
	if o == nil || o.FeishuTitle == nil {
		return nil, false
	}
	return o.FeishuTitle, true
}

// HasFeishuTitle returns a boolean if a field has been set.
func (o *AlertTemplate) HasFeishuTitle() bool {
	return o != nil && o.FeishuTitle != nil
}

// SetFeishuTitle gets a reference to the given string and assigns it to the FeishuTitle field.
func (o *AlertTemplate) SetFeishuTitle(v string) {
	o.FeishuTitle = &v
}

// GetDingdingText returns the DingdingText field value if set, zero value otherwise.
func (o *AlertTemplate) GetDingdingText() string {
	if o == nil || o.DingdingText == nil {
		var ret string
		return ret
	}
	return *o.DingdingText
}

// GetDingdingTextOk returns a tuple with the DingdingText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplate) GetDingdingTextOk() (*string, bool) {
	if o == nil || o.DingdingText == nil {
		return nil, false
	}
	return o.DingdingText, true
}

// HasDingdingText returns a boolean if a field has been set.
func (o *AlertTemplate) HasDingdingText() bool {
	return o != nil && o.DingdingText != nil
}

// SetDingdingText gets a reference to the given string and assigns it to the DingdingText field.
func (o *AlertTemplate) SetDingdingText(v string) {
	o.DingdingText = &v
}

// GetDingdingTitle returns the DingdingTitle field value if set, zero value otherwise.
func (o *AlertTemplate) GetDingdingTitle() string {
	if o == nil || o.DingdingTitle == nil {
		var ret string
		return ret
	}
	return *o.DingdingTitle
}

// GetDingdingTitleOk returns a tuple with the DingdingTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplate) GetDingdingTitleOk() (*string, bool) {
	if o == nil || o.DingdingTitle == nil {
		return nil, false
	}
	return o.DingdingTitle, true
}

// HasDingdingTitle returns a boolean if a field has been set.
func (o *AlertTemplate) HasDingdingTitle() bool {
	return o != nil && o.DingdingTitle != nil
}

// SetDingdingTitle gets a reference to the given string and assigns it to the DingdingTitle field.
func (o *AlertTemplate) SetDingdingTitle(v string) {
	o.DingdingTitle = &v
}

// GetWeixinText returns the WeixinText field value if set, zero value otherwise.
func (o *AlertTemplate) GetWeixinText() string {
	if o == nil || o.WeixinText == nil {
		var ret string
		return ret
	}
	return *o.WeixinText
}

// GetWeixinTextOk returns a tuple with the WeixinText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplate) GetWeixinTextOk() (*string, bool) {
	if o == nil || o.WeixinText == nil {
		return nil, false
	}
	return o.WeixinText, true
}

// HasWeixinText returns a boolean if a field has been set.
func (o *AlertTemplate) HasWeixinText() bool {
	return o != nil && o.WeixinText != nil
}

// SetWeixinText gets a reference to the given string and assigns it to the WeixinText field.
func (o *AlertTemplate) SetWeixinText(v string) {
	o.WeixinText = &v
}

// GetWeixinTitle returns the WeixinTitle field value if set, zero value otherwise.
func (o *AlertTemplate) GetWeixinTitle() string {
	if o == nil || o.WeixinTitle == nil {
		var ret string
		return ret
	}
	return *o.WeixinTitle
}

// GetWeixinTitleOk returns a tuple with the WeixinTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplate) GetWeixinTitleOk() (*string, bool) {
	if o == nil || o.WeixinTitle == nil {
		return nil, false
	}
	return o.WeixinTitle, true
}

// HasWeixinTitle returns a boolean if a field has been set.
func (o *AlertTemplate) HasWeixinTitle() bool {
	return o != nil && o.WeixinTitle != nil
}

// SetWeixinTitle gets a reference to the given string and assigns it to the WeixinTitle field.
func (o *AlertTemplate) SetWeixinTitle(v string) {
	o.WeixinTitle = &v
}

// GetWebhookText returns the WebhookText field value if set, zero value otherwise.
func (o *AlertTemplate) GetWebhookText() string {
	if o == nil || o.WebhookText == nil {
		var ret string
		return ret
	}
	return *o.WebhookText
}

// GetWebhookTextOk returns a tuple with the WebhookText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplate) GetWebhookTextOk() (*string, bool) {
	if o == nil || o.WebhookText == nil {
		return nil, false
	}
	return o.WebhookText, true
}

// HasWebhookText returns a boolean if a field has been set.
func (o *AlertTemplate) HasWebhookText() bool {
	return o != nil && o.WebhookText != nil
}

// SetWebhookText gets a reference to the given string and assigns it to the WebhookText field.
func (o *AlertTemplate) SetWebhookText(v string) {
	o.WebhookText = &v
}

// GetWebhookTitle returns the WebhookTitle field value if set, zero value otherwise.
func (o *AlertTemplate) GetWebhookTitle() string {
	if o == nil || o.WebhookTitle == nil {
		var ret string
		return ret
	}
	return *o.WebhookTitle
}

// GetWebhookTitleOk returns a tuple with the WebhookTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplate) GetWebhookTitleOk() (*string, bool) {
	if o == nil || o.WebhookTitle == nil {
		return nil, false
	}
	return o.WebhookTitle, true
}

// HasWebhookTitle returns a boolean if a field has been set.
func (o *AlertTemplate) HasWebhookTitle() bool {
	return o != nil && o.WebhookTitle != nil
}

// SetWebhookTitle gets a reference to the given string and assigns it to the WebhookTitle field.
func (o *AlertTemplate) SetWebhookTitle(v string) {
	o.WebhookTitle = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertTemplate) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplate) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertTemplate) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AlertTemplate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value.
func (o *AlertTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AlertTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AlertTemplate) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AlertTemplate) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplate) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AlertTemplate) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AlertTemplate) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *AlertTemplate) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTemplate) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *AlertTemplate) HasIsDefault() bool {
	return o != nil && o.IsDefault != nil
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *AlertTemplate) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertTemplate) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EmailSubject != nil {
		toSerialize["emailSubject"] = o.EmailSubject
	}
	if o.EmailText != nil {
		toSerialize["emailText"] = o.EmailText
	}
	if o.FeishuText != nil {
		toSerialize["feishuText"] = o.FeishuText
	}
	if o.FeishuTitle != nil {
		toSerialize["feishuTitle"] = o.FeishuTitle
	}
	if o.DingdingText != nil {
		toSerialize["dingdingText"] = o.DingdingText
	}
	if o.DingdingTitle != nil {
		toSerialize["dingdingTitle"] = o.DingdingTitle
	}
	if o.WeixinText != nil {
		toSerialize["weixinText"] = o.WeixinText
	}
	if o.WeixinTitle != nil {
		toSerialize["weixinTitle"] = o.WeixinTitle
	}
	if o.WebhookText != nil {
		toSerialize["webhookText"] = o.WebhookText
	}
	if o.WebhookTitle != nil {
		toSerialize["webhookTitle"] = o.WebhookTitle
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertTemplate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt     *time.Time `json:"createdAt,omitempty"`
		Description   *string    `json:"description,omitempty"`
		EmailSubject  *string    `json:"emailSubject,omitempty"`
		EmailText     *string    `json:"emailText,omitempty"`
		FeishuText    *string    `json:"feishuText,omitempty"`
		FeishuTitle   *string    `json:"feishuTitle,omitempty"`
		DingdingText  *string    `json:"dingdingText,omitempty"`
		DingdingTitle *string    `json:"dingdingTitle,omitempty"`
		WeixinText    *string    `json:"weixinText,omitempty"`
		WeixinTitle   *string    `json:"weixinTitle,omitempty"`
		WebhookText   *string    `json:"webhookText,omitempty"`
		WebhookTitle  *string    `json:"webhookTitle,omitempty"`
		Id            *string    `json:"id,omitempty"`
		Name          *string    `json:"name"`
		UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
		IsDefault     *bool      `json:"isDefault,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"createdAt", "description", "emailSubject", "emailText", "feishuText", "feishuTitle", "dingdingText", "dingdingTitle", "weixinText", "weixinTitle", "webhookText", "webhookTitle", "id", "name", "updatedAt", "isDefault"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.Description = all.Description
	o.EmailSubject = all.EmailSubject
	o.EmailText = all.EmailText
	o.FeishuText = all.FeishuText
	o.FeishuTitle = all.FeishuTitle
	o.DingdingText = all.DingdingText
	o.DingdingTitle = all.DingdingTitle
	o.WeixinText = all.WeixinText
	o.WeixinTitle = all.WeixinTitle
	o.WebhookText = all.WebhookText
	o.WebhookTitle = all.WebhookTitle
	o.Id = all.Id
	o.Name = *all.Name
	o.UpdatedAt = all.UpdatedAt
	o.IsDefault = all.IsDefault

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
