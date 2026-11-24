// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OrgDeletionPreview struct {
	OrgId             string                    `json:"orgID"`
	Name              string                    `json:"name"`
	Status            string                    `json:"status"`
	ExpiredAt         common.NullableTime       `json:"expiredAt"`
	PurgeError        *string                   `json:"purgeError,omitempty"`
	Action            string                    `json:"action"`
	CanProceed        bool                      `json:"canProceed"`
	Blockers          []OrgDeletionBlocker      `json:"blockers"`
	Configurations    []OrgConfigurationSummary `json:"configurations"`
	RetainedResources []OrgRetainedResource     `json:"retainedResources"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgDeletionPreview instantiates a new OrgDeletionPreview object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgDeletionPreview(orgId string, name string, status string, expiredAt common.NullableTime, action string, canProceed bool, blockers []OrgDeletionBlocker, configurations []OrgConfigurationSummary, retainedResources []OrgRetainedResource) *OrgDeletionPreview {
	this := OrgDeletionPreview{}
	this.OrgId = orgId
	this.Name = name
	this.Status = status
	this.ExpiredAt = expiredAt
	this.Action = action
	this.CanProceed = canProceed
	this.Blockers = blockers
	this.Configurations = configurations
	this.RetainedResources = retainedResources
	return &this
}

// NewOrgDeletionPreviewWithDefaults instantiates a new OrgDeletionPreview object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgDeletionPreviewWithDefaults() *OrgDeletionPreview {
	this := OrgDeletionPreview{}
	return &this
}

// GetOrgId returns the OrgId field value.
func (o *OrgDeletionPreview) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *OrgDeletionPreview) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *OrgDeletionPreview) SetOrgId(v string) {
	o.OrgId = v
}

// GetName returns the Name field value.
func (o *OrgDeletionPreview) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrgDeletionPreview) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OrgDeletionPreview) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value.
func (o *OrgDeletionPreview) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrgDeletionPreview) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *OrgDeletionPreview) SetStatus(v string) {
	o.Status = v
}

// GetExpiredAt returns the ExpiredAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *OrgDeletionPreview) GetExpiredAt() time.Time {
	if o == nil || o.ExpiredAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt.Get()
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OrgDeletionPreview) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiredAt.Get(), o.ExpiredAt.IsSet()
}

// SetExpiredAt sets field value.
func (o *OrgDeletionPreview) SetExpiredAt(v time.Time) {
	o.ExpiredAt.Set(&v)
}

// GetPurgeError returns the PurgeError field value if set, zero value otherwise.
func (o *OrgDeletionPreview) GetPurgeError() string {
	if o == nil || o.PurgeError == nil {
		var ret string
		return ret
	}
	return *o.PurgeError
}

// GetPurgeErrorOk returns a tuple with the PurgeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgDeletionPreview) GetPurgeErrorOk() (*string, bool) {
	if o == nil || o.PurgeError == nil {
		return nil, false
	}
	return o.PurgeError, true
}

// HasPurgeError returns a boolean if a field has been set.
func (o *OrgDeletionPreview) HasPurgeError() bool {
	return o != nil && o.PurgeError != nil
}

// SetPurgeError gets a reference to the given string and assigns it to the PurgeError field.
func (o *OrgDeletionPreview) SetPurgeError(v string) {
	o.PurgeError = &v
}

// GetAction returns the Action field value.
func (o *OrgDeletionPreview) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *OrgDeletionPreview) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *OrgDeletionPreview) SetAction(v string) {
	o.Action = v
}

// GetCanProceed returns the CanProceed field value.
func (o *OrgDeletionPreview) GetCanProceed() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.CanProceed
}

// GetCanProceedOk returns a tuple with the CanProceed field value
// and a boolean to check if the value has been set.
func (o *OrgDeletionPreview) GetCanProceedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanProceed, true
}

// SetCanProceed sets field value.
func (o *OrgDeletionPreview) SetCanProceed(v bool) {
	o.CanProceed = v
}

// GetBlockers returns the Blockers field value.
func (o *OrgDeletionPreview) GetBlockers() []OrgDeletionBlocker {
	if o == nil {
		var ret []OrgDeletionBlocker
		return ret
	}
	return o.Blockers
}

// GetBlockersOk returns a tuple with the Blockers field value
// and a boolean to check if the value has been set.
func (o *OrgDeletionPreview) GetBlockersOk() (*[]OrgDeletionBlocker, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blockers, true
}

// SetBlockers sets field value.
func (o *OrgDeletionPreview) SetBlockers(v []OrgDeletionBlocker) {
	o.Blockers = v
}

// GetConfigurations returns the Configurations field value.
func (o *OrgDeletionPreview) GetConfigurations() []OrgConfigurationSummary {
	if o == nil {
		var ret []OrgConfigurationSummary
		return ret
	}
	return o.Configurations
}

// GetConfigurationsOk returns a tuple with the Configurations field value
// and a boolean to check if the value has been set.
func (o *OrgDeletionPreview) GetConfigurationsOk() (*[]OrgConfigurationSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configurations, true
}

// SetConfigurations sets field value.
func (o *OrgDeletionPreview) SetConfigurations(v []OrgConfigurationSummary) {
	o.Configurations = v
}

// GetRetainedResources returns the RetainedResources field value.
func (o *OrgDeletionPreview) GetRetainedResources() []OrgRetainedResource {
	if o == nil {
		var ret []OrgRetainedResource
		return ret
	}
	return o.RetainedResources
}

// GetRetainedResourcesOk returns a tuple with the RetainedResources field value
// and a boolean to check if the value has been set.
func (o *OrgDeletionPreview) GetRetainedResourcesOk() (*[]OrgRetainedResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetainedResources, true
}

// SetRetainedResources sets field value.
func (o *OrgDeletionPreview) SetRetainedResources(v []OrgRetainedResource) {
	o.RetainedResources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgDeletionPreview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["orgID"] = o.OrgId
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["expiredAt"] = o.ExpiredAt.Get()
	if o.PurgeError != nil {
		toSerialize["purgeError"] = o.PurgeError
	}
	toSerialize["action"] = o.Action
	toSerialize["canProceed"] = o.CanProceed
	toSerialize["blockers"] = o.Blockers
	toSerialize["configurations"] = o.Configurations
	toSerialize["retainedResources"] = o.RetainedResources

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgDeletionPreview) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgId             *string                    `json:"orgID"`
		Name              *string                    `json:"name"`
		Status            *string                    `json:"status"`
		ExpiredAt         common.NullableTime        `json:"expiredAt"`
		PurgeError        *string                    `json:"purgeError,omitempty"`
		Action            *string                    `json:"action"`
		CanProceed        *bool                      `json:"canProceed"`
		Blockers          *[]OrgDeletionBlocker      `json:"blockers"`
		Configurations    *[]OrgConfigurationSummary `json:"configurations"`
		RetainedResources *[]OrgRetainedResource     `json:"retainedResources"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field orgID missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if !all.ExpiredAt.IsSet() {
		return fmt.Errorf("required field expiredAt missing")
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.CanProceed == nil {
		return fmt.Errorf("required field canProceed missing")
	}
	if all.Blockers == nil {
		return fmt.Errorf("required field blockers missing")
	}
	if all.Configurations == nil {
		return fmt.Errorf("required field configurations missing")
	}
	if all.RetainedResources == nil {
		return fmt.Errorf("required field retainedResources missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"orgID", "name", "status", "expiredAt", "purgeError", "action", "canProceed", "blockers", "configurations", "retainedResources"})
	} else {
		return err
	}
	o.OrgId = *all.OrgId
	o.Name = *all.Name
	o.Status = *all.Status
	o.ExpiredAt = all.ExpiredAt
	o.PurgeError = all.PurgeError
	o.Action = *all.Action
	o.CanProceed = *all.CanProceed
	o.Blockers = *all.Blockers
	o.Configurations = *all.Configurations
	o.RetainedResources = *all.RetainedResources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
