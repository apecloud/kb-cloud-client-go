// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PlatformAlert_alertRule Alert rule information
type PlatformAlert_alertRule struct {
	Description *string        `json:"description,omitempty"`
	Summary     *string        `json:"summary,omitempty"`
	GroupName   *string        `json:"groupName,omitempty"`
	AlertName   *string        `json:"alertName,omitempty"`
	Expr        *string        `json:"expr,omitempty"`
	For         *string        `json:"for,omitempty"`
	Disabled    *bool          `json:"disabled,omitempty"`
	Severity    *AlertSeverity `json:"severity,omitempty"`
	CreatedAt   *time.Time     `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time     `json:"updatedAt,omitempty"`
	// Alert metric information
	Metric *AlertMetric `json:"metric,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPlatformAlert_alertRule instantiates a new PlatformAlert_alertRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPlatformAlert_alertRule() *PlatformAlert_alertRule {
	this := PlatformAlert_alertRule{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// NewPlatformAlert_alertRuleWithDefaults instantiates a new PlatformAlert_alertRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPlatformAlert_alertRuleWithDefaults() *PlatformAlert_alertRule {
	this := PlatformAlert_alertRule{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PlatformAlert_alertRule) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformAlert_alertRule) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PlatformAlert_alertRule) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PlatformAlert_alertRule) SetDescription(v string) {
	o.Description = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *PlatformAlert_alertRule) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformAlert_alertRule) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *PlatformAlert_alertRule) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *PlatformAlert_alertRule) SetSummary(v string) {
	o.Summary = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *PlatformAlert_alertRule) GetGroupName() string {
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformAlert_alertRule) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *PlatformAlert_alertRule) HasGroupName() bool {
	return o != nil && o.GroupName != nil
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *PlatformAlert_alertRule) SetGroupName(v string) {
	o.GroupName = &v
}

// GetAlertName returns the AlertName field value if set, zero value otherwise.
func (o *PlatformAlert_alertRule) GetAlertName() string {
	if o == nil || o.AlertName == nil {
		var ret string
		return ret
	}
	return *o.AlertName
}

// GetAlertNameOk returns a tuple with the AlertName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformAlert_alertRule) GetAlertNameOk() (*string, bool) {
	if o == nil || o.AlertName == nil {
		return nil, false
	}
	return o.AlertName, true
}

// HasAlertName returns a boolean if a field has been set.
func (o *PlatformAlert_alertRule) HasAlertName() bool {
	return o != nil && o.AlertName != nil
}

// SetAlertName gets a reference to the given string and assigns it to the AlertName field.
func (o *PlatformAlert_alertRule) SetAlertName(v string) {
	o.AlertName = &v
}

// GetExpr returns the Expr field value if set, zero value otherwise.
func (o *PlatformAlert_alertRule) GetExpr() string {
	if o == nil || o.Expr == nil {
		var ret string
		return ret
	}
	return *o.Expr
}

// GetExprOk returns a tuple with the Expr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformAlert_alertRule) GetExprOk() (*string, bool) {
	if o == nil || o.Expr == nil {
		return nil, false
	}
	return o.Expr, true
}

// HasExpr returns a boolean if a field has been set.
func (o *PlatformAlert_alertRule) HasExpr() bool {
	return o != nil && o.Expr != nil
}

// SetExpr gets a reference to the given string and assigns it to the Expr field.
func (o *PlatformAlert_alertRule) SetExpr(v string) {
	o.Expr = &v
}

// GetFor returns the For field value if set, zero value otherwise.
func (o *PlatformAlert_alertRule) GetFor() string {
	if o == nil || o.For == nil {
		var ret string
		return ret
	}
	return *o.For
}

// GetForOk returns a tuple with the For field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformAlert_alertRule) GetForOk() (*string, bool) {
	if o == nil || o.For == nil {
		return nil, false
	}
	return o.For, true
}

// HasFor returns a boolean if a field has been set.
func (o *PlatformAlert_alertRule) HasFor() bool {
	return o != nil && o.For != nil
}

// SetFor gets a reference to the given string and assigns it to the For field.
func (o *PlatformAlert_alertRule) SetFor(v string) {
	o.For = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *PlatformAlert_alertRule) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformAlert_alertRule) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *PlatformAlert_alertRule) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *PlatformAlert_alertRule) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *PlatformAlert_alertRule) GetSeverity() AlertSeverity {
	if o == nil || o.Severity == nil {
		var ret AlertSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformAlert_alertRule) GetSeverityOk() (*AlertSeverity, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *PlatformAlert_alertRule) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given AlertSeverity and assigns it to the Severity field.
func (o *PlatformAlert_alertRule) SetSeverity(v AlertSeverity) {
	o.Severity = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PlatformAlert_alertRule) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformAlert_alertRule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PlatformAlert_alertRule) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PlatformAlert_alertRule) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PlatformAlert_alertRule) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformAlert_alertRule) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PlatformAlert_alertRule) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PlatformAlert_alertRule) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *PlatformAlert_alertRule) GetMetric() AlertMetric {
	if o == nil || o.Metric == nil {
		var ret AlertMetric
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformAlert_alertRule) GetMetricOk() (*AlertMetric, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *PlatformAlert_alertRule) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given AlertMetric and assigns it to the Metric field.
func (o *PlatformAlert_alertRule) SetMetric(v AlertMetric) {
	o.Metric = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PlatformAlert_alertRule) MarshalJSON() ([]byte, error) {
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
	if o.GroupName != nil {
		toSerialize["groupName"] = o.GroupName
	}
	if o.AlertName != nil {
		toSerialize["alertName"] = o.AlertName
	}
	if o.Expr != nil {
		toSerialize["expr"] = o.Expr
	}
	if o.For != nil {
		toSerialize["for"] = o.For
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PlatformAlert_alertRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string        `json:"description,omitempty"`
		Summary     *string        `json:"summary,omitempty"`
		GroupName   *string        `json:"groupName,omitempty"`
		AlertName   *string        `json:"alertName,omitempty"`
		Expr        *string        `json:"expr,omitempty"`
		For         *string        `json:"for,omitempty"`
		Disabled    *bool          `json:"disabled,omitempty"`
		Severity    *AlertSeverity `json:"severity,omitempty"`
		CreatedAt   *time.Time     `json:"createdAt,omitempty"`
		UpdatedAt   *time.Time     `json:"updatedAt,omitempty"`
		Metric      *AlertMetric   `json:"metric,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"description", "summary", "groupName", "alertName", "expr", "for", "disabled", "severity", "createdAt", "updatedAt", "metric"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.Summary = all.Summary
	o.GroupName = all.GroupName
	o.AlertName = all.AlertName
	o.Expr = all.Expr
	o.For = all.For
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
