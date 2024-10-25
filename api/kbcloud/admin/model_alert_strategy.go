// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AlertStrategy Alert strategy information

type AlertStrategy struct {
	// NODESCRIPTION Id
	Id *int32 `json:"id,omitempty"`
	// NODESCRIPTION OrgName
	OrgName *string `json:"orgName,omitempty"`
	// NODESCRIPTION Name
	Name *string `json:"name,omitempty"`
	// NODESCRIPTION Description
	Description *string `json:"description,omitempty"`
	// NODESCRIPTION CreatedAt
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// NODESCRIPTION UpdatedAt
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// NODESCRIPTION ReceiverIds
	ReceiverIds []int32 `json:"receiverIds"`
	// NODESCRIPTION Receivers
	Receivers []AlertReceiver `json:"receivers,omitempty"`
	// NODESCRIPTION Envs
	Envs []string `json:"envs,omitempty"`
	// NODESCRIPTION Severities
	Severities []string `json:"severities,omitempty"`
	// NODESCRIPTION Rules
	Rules []string `json:"rules,omitempty"`
	// NODESCRIPTION RuleObjs
	RuleObjs []AlertRule `json:"ruleObjs,omitempty"`
	// NODESCRIPTION Engines
	Engines []string `json:"engines,omitempty"`
	// NODESCRIPTION Clusters
	Clusters []string `json:"clusters,omitempty"`
	// NODESCRIPTION Disabled
	Disabled *bool `json:"disabled,omitempty"`
	// NODESCRIPTION RepeatInterval
	RepeatInterval *string `json:"repeatInterval,omitempty"`
	// NODESCRIPTION MuteTimeInterval
	MuteTimeInterval *AlertStrategyMuteTimeInterval `json:"muteTimeInterval,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertStrategy instantiates a new AlertStrategy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertStrategy(receiverIds []int32) *AlertStrategy {
	this := AlertStrategy{}
	this.ReceiverIds = receiverIds
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// NewAlertStrategyWithDefaults instantiates a new AlertStrategy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertStrategyWithDefaults() *AlertStrategy {
	this := AlertStrategy{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertStrategy) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertStrategy) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AlertStrategy) SetId(v int32) {
	o.Id = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *AlertStrategy) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *AlertStrategy) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *AlertStrategy) SetOrgName(v string) {
	o.OrgName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlertStrategy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlertStrategy) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlertStrategy) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlertStrategy) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlertStrategy) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlertStrategy) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AlertStrategy) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AlertStrategy) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AlertStrategy) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AlertStrategy) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AlertStrategy) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AlertStrategy) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetReceiverIds returns the ReceiverIds field value.
func (o *AlertStrategy) GetReceiverIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.ReceiverIds
}

// GetReceiverIdsOk returns a tuple with the ReceiverIds field value
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetReceiverIdsOk() (*[]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiverIds, true
}

// SetReceiverIds sets field value.
func (o *AlertStrategy) SetReceiverIds(v []int32) {
	o.ReceiverIds = v
}

// GetReceivers returns the Receivers field value if set, zero value otherwise.
func (o *AlertStrategy) GetReceivers() []AlertReceiver {
	if o == nil || o.Receivers == nil {
		var ret []AlertReceiver
		return ret
	}
	return o.Receivers
}

// GetReceiversOk returns a tuple with the Receivers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetReceiversOk() (*[]AlertReceiver, bool) {
	if o == nil || o.Receivers == nil {
		return nil, false
	}
	return &o.Receivers, true
}

// HasReceivers returns a boolean if a field has been set.
func (o *AlertStrategy) HasReceivers() bool {
	return o != nil && o.Receivers != nil
}

// SetReceivers gets a reference to the given []AlertReceiver and assigns it to the Receivers field.
func (o *AlertStrategy) SetReceivers(v []AlertReceiver) {
	o.Receivers = v
}

// GetEnvs returns the Envs field value if set, zero value otherwise.
func (o *AlertStrategy) GetEnvs() []string {
	if o == nil || o.Envs == nil {
		var ret []string
		return ret
	}
	return o.Envs
}

// GetEnvsOk returns a tuple with the Envs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetEnvsOk() (*[]string, bool) {
	if o == nil || o.Envs == nil {
		return nil, false
	}
	return &o.Envs, true
}

// HasEnvs returns a boolean if a field has been set.
func (o *AlertStrategy) HasEnvs() bool {
	return o != nil && o.Envs != nil
}

// SetEnvs gets a reference to the given []string and assigns it to the Envs field.
func (o *AlertStrategy) SetEnvs(v []string) {
	o.Envs = v
}

// GetSeverities returns the Severities field value if set, zero value otherwise.
func (o *AlertStrategy) GetSeverities() []string {
	if o == nil || o.Severities == nil {
		var ret []string
		return ret
	}
	return o.Severities
}

// GetSeveritiesOk returns a tuple with the Severities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetSeveritiesOk() (*[]string, bool) {
	if o == nil || o.Severities == nil {
		return nil, false
	}
	return &o.Severities, true
}

// HasSeverities returns a boolean if a field has been set.
func (o *AlertStrategy) HasSeverities() bool {
	return o != nil && o.Severities != nil
}

// SetSeverities gets a reference to the given []string and assigns it to the Severities field.
func (o *AlertStrategy) SetSeverities(v []string) {
	o.Severities = v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *AlertStrategy) GetRules() []string {
	if o == nil || o.Rules == nil {
		var ret []string
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetRulesOk() (*[]string, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return &o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *AlertStrategy) HasRules() bool {
	return o != nil && o.Rules != nil
}

// SetRules gets a reference to the given []string and assigns it to the Rules field.
func (o *AlertStrategy) SetRules(v []string) {
	o.Rules = v
}

// GetRuleObjs returns the RuleObjs field value if set, zero value otherwise.
func (o *AlertStrategy) GetRuleObjs() []AlertRule {
	if o == nil || o.RuleObjs == nil {
		var ret []AlertRule
		return ret
	}
	return o.RuleObjs
}

// GetRuleObjsOk returns a tuple with the RuleObjs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetRuleObjsOk() (*[]AlertRule, bool) {
	if o == nil || o.RuleObjs == nil {
		return nil, false
	}
	return &o.RuleObjs, true
}

// HasRuleObjs returns a boolean if a field has been set.
func (o *AlertStrategy) HasRuleObjs() bool {
	return o != nil && o.RuleObjs != nil
}

// SetRuleObjs gets a reference to the given []AlertRule and assigns it to the RuleObjs field.
func (o *AlertStrategy) SetRuleObjs(v []AlertRule) {
	o.RuleObjs = v
}

// GetEngines returns the Engines field value if set, zero value otherwise.
func (o *AlertStrategy) GetEngines() []string {
	if o == nil || o.Engines == nil {
		var ret []string
		return ret
	}
	return o.Engines
}

// GetEnginesOk returns a tuple with the Engines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetEnginesOk() (*[]string, bool) {
	if o == nil || o.Engines == nil {
		return nil, false
	}
	return &o.Engines, true
}

// HasEngines returns a boolean if a field has been set.
func (o *AlertStrategy) HasEngines() bool {
	return o != nil && o.Engines != nil
}

// SetEngines gets a reference to the given []string and assigns it to the Engines field.
func (o *AlertStrategy) SetEngines(v []string) {
	o.Engines = v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *AlertStrategy) GetClusters() []string {
	if o == nil || o.Clusters == nil {
		var ret []string
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetClustersOk() (*[]string, bool) {
	if o == nil || o.Clusters == nil {
		return nil, false
	}
	return &o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *AlertStrategy) HasClusters() bool {
	return o != nil && o.Clusters != nil
}

// SetClusters gets a reference to the given []string and assigns it to the Clusters field.
func (o *AlertStrategy) SetClusters(v []string) {
	o.Clusters = v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *AlertStrategy) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *AlertStrategy) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *AlertStrategy) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetRepeatInterval returns the RepeatInterval field value if set, zero value otherwise.
func (o *AlertStrategy) GetRepeatInterval() string {
	if o == nil || o.RepeatInterval == nil {
		var ret string
		return ret
	}
	return *o.RepeatInterval
}

// GetRepeatIntervalOk returns a tuple with the RepeatInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetRepeatIntervalOk() (*string, bool) {
	if o == nil || o.RepeatInterval == nil {
		return nil, false
	}
	return o.RepeatInterval, true
}

// HasRepeatInterval returns a boolean if a field has been set.
func (o *AlertStrategy) HasRepeatInterval() bool {
	return o != nil && o.RepeatInterval != nil
}

// SetRepeatInterval gets a reference to the given string and assigns it to the RepeatInterval field.
func (o *AlertStrategy) SetRepeatInterval(v string) {
	o.RepeatInterval = &v
}

// GetMuteTimeInterval returns the MuteTimeInterval field value if set, zero value otherwise.
func (o *AlertStrategy) GetMuteTimeInterval() AlertStrategyMuteTimeInterval {
	if o == nil || o.MuteTimeInterval == nil {
		var ret AlertStrategyMuteTimeInterval
		return ret
	}
	return *o.MuteTimeInterval
}

// GetMuteTimeIntervalOk returns a tuple with the MuteTimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetMuteTimeIntervalOk() (*AlertStrategyMuteTimeInterval, bool) {
	if o == nil || o.MuteTimeInterval == nil {
		return nil, false
	}
	return o.MuteTimeInterval, true
}

// HasMuteTimeInterval returns a boolean if a field has been set.
func (o *AlertStrategy) HasMuteTimeInterval() bool {
	return o != nil && o.MuteTimeInterval != nil
}

// SetMuteTimeInterval gets a reference to the given AlertStrategyMuteTimeInterval and assigns it to the MuteTimeInterval field.
func (o *AlertStrategy) SetMuteTimeInterval(v AlertStrategyMuteTimeInterval) {
	o.MuteTimeInterval = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertStrategy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
	toSerialize["receiverIds"] = o.ReceiverIds
	if o.Receivers != nil {
		toSerialize["receivers"] = o.Receivers
	}
	if o.Envs != nil {
		toSerialize["envs"] = o.Envs
	}
	if o.Severities != nil {
		toSerialize["severities"] = o.Severities
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.RuleObjs != nil {
		toSerialize["ruleObjs"] = o.RuleObjs
	}
	if o.Engines != nil {
		toSerialize["engines"] = o.Engines
	}
	if o.Clusters != nil {
		toSerialize["clusters"] = o.Clusters
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.RepeatInterval != nil {
		toSerialize["repeatInterval"] = o.RepeatInterval
	}
	if o.MuteTimeInterval != nil {
		toSerialize["muteTimeInterval"] = o.MuteTimeInterval
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertStrategy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id               *int32                         `json:"id,omitempty"`
		OrgName          *string                        `json:"orgName,omitempty"`
		Name             *string                        `json:"name,omitempty"`
		Description      *string                        `json:"description,omitempty"`
		CreatedAt        *time.Time                     `json:"createdAt,omitempty"`
		UpdatedAt        *time.Time                     `json:"updatedAt,omitempty"`
		ReceiverIds      *[]int32                       `json:"receiverIds"`
		Receivers        []AlertReceiver                `json:"receivers,omitempty"`
		Envs             []string                       `json:"envs,omitempty"`
		Severities       []string                       `json:"severities,omitempty"`
		Rules            []string                       `json:"rules,omitempty"`
		RuleObjs         []AlertRule                    `json:"ruleObjs,omitempty"`
		Engines          []string                       `json:"engines,omitempty"`
		Clusters         []string                       `json:"clusters,omitempty"`
		Disabled         *bool                          `json:"disabled,omitempty"`
		RepeatInterval   *string                        `json:"repeatInterval,omitempty"`
		MuteTimeInterval *AlertStrategyMuteTimeInterval `json:"muteTimeInterval,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ReceiverIds == nil {
		return fmt.Errorf("required field receiverIds missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "orgName", "name", "description", "createdAt", "updatedAt", "receiverIds", "receivers", "envs", "severities", "rules", "ruleObjs", "engines", "clusters", "disabled", "repeatInterval", "muteTimeInterval"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.OrgName = all.OrgName
	o.Name = all.Name
	o.Description = all.Description
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt
	o.ReceiverIds = *all.ReceiverIds
	o.Receivers = all.Receivers
	o.Envs = all.Envs
	o.Severities = all.Severities
	o.Rules = all.Rules
	o.RuleObjs = all.RuleObjs
	o.Engines = all.Engines
	o.Clusters = all.Clusters
	o.Disabled = all.Disabled
	o.RepeatInterval = all.RepeatInterval
	if all.MuteTimeInterval != nil && all.MuteTimeInterval.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MuteTimeInterval = all.MuteTimeInterval

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
