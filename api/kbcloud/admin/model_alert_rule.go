// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AlertRule Alert rule information
// NODESCRIPTION AlertRule
//
// Deprecated: This model is deprecated.
type AlertRule struct {
	// NODESCRIPTION Description
	Description *string `json:"description,omitempty"`
	// NODESCRIPTION Summary
	Summary *string `json:"summary,omitempty"`
	// NODESCRIPTION AlertName
	AlertName string `json:"alertName"`
	// NODESCRIPTION Expr
	Expr *string `json:"expr,omitempty"`
	// NODESCRIPTION For
	For *string `json:"for,omitempty"`
	// NODESCRIPTION GroupName
	GroupName *string `json:"groupName,omitempty"`
	// NODESCRIPTION Disabled
	Disabled *bool `json:"disabled,omitempty"`
	// NODESCRIPTION Severity
	Severity *AlertSeverity `json:"severity,omitempty"`
	// NODESCRIPTION CreatedAt
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// NODESCRIPTION UpdatedAt
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// Alert metric information
	Metric *AlertMetric `json:"metric,omitempty"`
	// NODESCRIPTION OrgName
	OrgName *string `json:"orgName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertRule instantiates a new AlertRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertRule(alertName string) *AlertRule {
	this := AlertRule{}
	this.AlertName = alertName
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// NewAlertRuleWithDefaults instantiates a new AlertRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertRuleWithDefaults() *AlertRule {
	this := AlertRule{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlertRule) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlertRule) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlertRule) SetDescription(v string) {
	o.Description = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AlertRule) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AlertRule) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AlertRule) SetSummary(v string) {
	o.Summary = &v
}

// GetAlertName returns the AlertName field value.
func (o *AlertRule) GetAlertName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AlertName
}

// GetAlertNameOk returns a tuple with the AlertName field value
// and a boolean to check if the value has been set.
func (o *AlertRule) GetAlertNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertName, true
}

// SetAlertName sets field value.
func (o *AlertRule) SetAlertName(v string) {
	o.AlertName = v
}

// GetExpr returns the Expr field value if set, zero value otherwise.
func (o *AlertRule) GetExpr() string {
	if o == nil || o.Expr == nil {
		var ret string
		return ret
	}
	return *o.Expr
}

// GetExprOk returns a tuple with the Expr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetExprOk() (*string, bool) {
	if o == nil || o.Expr == nil {
		return nil, false
	}
	return o.Expr, true
}

// HasExpr returns a boolean if a field has been set.
func (o *AlertRule) HasExpr() bool {
	return o != nil && o.Expr != nil
}

// SetExpr gets a reference to the given string and assigns it to the Expr field.
func (o *AlertRule) SetExpr(v string) {
	o.Expr = &v
}

// GetFor returns the For field value if set, zero value otherwise.
func (o *AlertRule) GetFor() string {
	if o == nil || o.For == nil {
		var ret string
		return ret
	}
	return *o.For
}

// GetForOk returns a tuple with the For field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetForOk() (*string, bool) {
	if o == nil || o.For == nil {
		return nil, false
	}
	return o.For, true
}

// HasFor returns a boolean if a field has been set.
func (o *AlertRule) HasFor() bool {
	return o != nil && o.For != nil
}

// SetFor gets a reference to the given string and assigns it to the For field.
func (o *AlertRule) SetFor(v string) {
	o.For = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *AlertRule) GetGroupName() string {
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *AlertRule) HasGroupName() bool {
	return o != nil && o.GroupName != nil
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *AlertRule) SetGroupName(v string) {
	o.GroupName = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *AlertRule) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *AlertRule) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *AlertRule) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *AlertRule) GetSeverity() AlertSeverity {
	if o == nil || o.Severity == nil {
		var ret AlertSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetSeverityOk() (*AlertSeverity, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *AlertRule) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given AlertSeverity and assigns it to the Severity field.
func (o *AlertRule) SetSeverity(v AlertSeverity) {
	o.Severity = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AlertRule) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AlertRule) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AlertRule) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AlertRule) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AlertRule) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AlertRule) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *AlertRule) GetMetric() AlertMetric {
	if o == nil || o.Metric == nil {
		var ret AlertMetric
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetMetricOk() (*AlertMetric, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *AlertRule) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given AlertMetric and assigns it to the Metric field.
func (o *AlertRule) SetMetric(v AlertMetric) {
	o.Metric = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *AlertRule) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *AlertRule) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *AlertRule) SetOrgName(v string) {
	o.OrgName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	toSerialize["alertName"] = o.AlertName
	if o.Expr != nil {
		toSerialize["expr"] = o.Expr
	}
	if o.For != nil {
		toSerialize["for"] = o.For
	}
	if o.GroupName != nil {
		toSerialize["groupName"] = o.GroupName
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
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
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
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
func (o *AlertRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string        `json:"description,omitempty"`
		Summary     *string        `json:"summary,omitempty"`
		AlertName   *string        `json:"alertName"`
		Expr        *string        `json:"expr,omitempty"`
		For         *string        `json:"for,omitempty"`
		GroupName   *string        `json:"groupName,omitempty"`
		Disabled    *bool          `json:"disabled,omitempty"`
		Severity    *AlertSeverity `json:"severity,omitempty"`
		CreatedAt   *time.Time     `json:"createdAt,omitempty"`
		UpdatedAt   *time.Time     `json:"updatedAt,omitempty"`
		Metric      *AlertMetric   `json:"metric,omitempty"`
		OrgName     *string        `json:"orgName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AlertName == nil {
		return fmt.Errorf("required field alertName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"description", "summary", "alertName", "expr", "for", "groupName", "disabled", "severity", "createdAt", "updatedAt", "metric", "orgName"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.Summary = all.Summary
	o.AlertName = *all.AlertName
	o.Expr = all.Expr
	o.For = all.For
	o.GroupName = all.GroupName
	o.Disabled = all.Disabled
	if all.Severity != nil && !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = all.Severity
	}
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt
	if all.Metric != nil && all.Metric.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metric = all.Metric
	o.OrgName = all.OrgName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
