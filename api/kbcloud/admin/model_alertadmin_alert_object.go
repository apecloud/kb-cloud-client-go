// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Alertadmin_alertObject Alert object information
type Alertadmin_alertObject struct {
	Id          *int32         `json:"id,omitempty"`
	AlertName   *string        `json:"alertName,omitempty"`
	GroupName   *string        `json:"groupName,omitempty"`
	Expr        *string        `json:"expr,omitempty"`
	Engine      *string        `json:"engine,omitempty"`
	Namespace   *string        `json:"namespace,omitempty"`
	Pod         *string        `json:"pod,omitempty"`
	Severity    *AlertSeverity `json:"severity,omitempty"`
	Description *string        `json:"description,omitempty"`
	Fingerprint *string        `json:"fingerprint,omitempty"`
	StartsAt    *time.Time     `json:"startsAt,omitempty"`
	EndsAt      *time.Time     `json:"endsAt,omitempty"`
	Status      *AlertStatus   `json:"status,omitempty"`
	Count       *int32         `json:"count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertadmin_alertObject instantiates a new Alertadmin_alertObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertadmin_alertObject() *Alertadmin_alertObject {
	this := Alertadmin_alertObject{}
	return &this
}

// NewAlertadmin_alertObjectWithDefaults instantiates a new Alertadmin_alertObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertadmin_alertObjectWithDefaults() *Alertadmin_alertObject {
	this := Alertadmin_alertObject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Alertadmin_alertObject) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alertadmin_alertObject) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Alertadmin_alertObject) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Alertadmin_alertObject) SetId(v int32) {
	o.Id = &v
}

// GetAlertName returns the AlertName field value if set, zero value otherwise.
func (o *Alertadmin_alertObject) GetAlertName() string {
	if o == nil || o.AlertName == nil {
		var ret string
		return ret
	}
	return *o.AlertName
}

// GetAlertNameOk returns a tuple with the AlertName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alertadmin_alertObject) GetAlertNameOk() (*string, bool) {
	if o == nil || o.AlertName == nil {
		return nil, false
	}
	return o.AlertName, true
}

// HasAlertName returns a boolean if a field has been set.
func (o *Alertadmin_alertObject) HasAlertName() bool {
	return o != nil && o.AlertName != nil
}

// SetAlertName gets a reference to the given string and assigns it to the AlertName field.
func (o *Alertadmin_alertObject) SetAlertName(v string) {
	o.AlertName = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *Alertadmin_alertObject) GetGroupName() string {
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alertadmin_alertObject) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *Alertadmin_alertObject) HasGroupName() bool {
	return o != nil && o.GroupName != nil
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *Alertadmin_alertObject) SetGroupName(v string) {
	o.GroupName = &v
}

// GetExpr returns the Expr field value if set, zero value otherwise.
func (o *Alertadmin_alertObject) GetExpr() string {
	if o == nil || o.Expr == nil {
		var ret string
		return ret
	}
	return *o.Expr
}

// GetExprOk returns a tuple with the Expr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alertadmin_alertObject) GetExprOk() (*string, bool) {
	if o == nil || o.Expr == nil {
		return nil, false
	}
	return o.Expr, true
}

// HasExpr returns a boolean if a field has been set.
func (o *Alertadmin_alertObject) HasExpr() bool {
	return o != nil && o.Expr != nil
}

// SetExpr gets a reference to the given string and assigns it to the Expr field.
func (o *Alertadmin_alertObject) SetExpr(v string) {
	o.Expr = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *Alertadmin_alertObject) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alertadmin_alertObject) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *Alertadmin_alertObject) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *Alertadmin_alertObject) SetEngine(v string) {
	o.Engine = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *Alertadmin_alertObject) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alertadmin_alertObject) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *Alertadmin_alertObject) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *Alertadmin_alertObject) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPod returns the Pod field value if set, zero value otherwise.
func (o *Alertadmin_alertObject) GetPod() string {
	if o == nil || o.Pod == nil {
		var ret string
		return ret
	}
	return *o.Pod
}

// GetPodOk returns a tuple with the Pod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alertadmin_alertObject) GetPodOk() (*string, bool) {
	if o == nil || o.Pod == nil {
		return nil, false
	}
	return o.Pod, true
}

// HasPod returns a boolean if a field has been set.
func (o *Alertadmin_alertObject) HasPod() bool {
	return o != nil && o.Pod != nil
}

// SetPod gets a reference to the given string and assigns it to the Pod field.
func (o *Alertadmin_alertObject) SetPod(v string) {
	o.Pod = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *Alertadmin_alertObject) GetSeverity() AlertSeverity {
	if o == nil || o.Severity == nil {
		var ret AlertSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alertadmin_alertObject) GetSeverityOk() (*AlertSeverity, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *Alertadmin_alertObject) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given AlertSeverity and assigns it to the Severity field.
func (o *Alertadmin_alertObject) SetSeverity(v AlertSeverity) {
	o.Severity = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Alertadmin_alertObject) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alertadmin_alertObject) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Alertadmin_alertObject) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Alertadmin_alertObject) SetDescription(v string) {
	o.Description = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *Alertadmin_alertObject) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alertadmin_alertObject) GetFingerprintOk() (*string, bool) {
	if o == nil || o.Fingerprint == nil {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *Alertadmin_alertObject) HasFingerprint() bool {
	return o != nil && o.Fingerprint != nil
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *Alertadmin_alertObject) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *Alertadmin_alertObject) GetStartsAt() time.Time {
	if o == nil || o.StartsAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alertadmin_alertObject) GetStartsAtOk() (*time.Time, bool) {
	if o == nil || o.StartsAt == nil {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *Alertadmin_alertObject) HasStartsAt() bool {
	return o != nil && o.StartsAt != nil
}

// SetStartsAt gets a reference to the given time.Time and assigns it to the StartsAt field.
func (o *Alertadmin_alertObject) SetStartsAt(v time.Time) {
	o.StartsAt = &v
}

// GetEndsAt returns the EndsAt field value if set, zero value otherwise.
func (o *Alertadmin_alertObject) GetEndsAt() time.Time {
	if o == nil || o.EndsAt == nil {
		var ret time.Time
		return ret
	}
	return *o.EndsAt
}

// GetEndsAtOk returns a tuple with the EndsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alertadmin_alertObject) GetEndsAtOk() (*time.Time, bool) {
	if o == nil || o.EndsAt == nil {
		return nil, false
	}
	return o.EndsAt, true
}

// HasEndsAt returns a boolean if a field has been set.
func (o *Alertadmin_alertObject) HasEndsAt() bool {
	return o != nil && o.EndsAt != nil
}

// SetEndsAt gets a reference to the given time.Time and assigns it to the EndsAt field.
func (o *Alertadmin_alertObject) SetEndsAt(v time.Time) {
	o.EndsAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Alertadmin_alertObject) GetStatus() AlertStatus {
	if o == nil || o.Status == nil {
		var ret AlertStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alertadmin_alertObject) GetStatusOk() (*AlertStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Alertadmin_alertObject) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given AlertStatus and assigns it to the Status field.
func (o *Alertadmin_alertObject) SetStatus(v AlertStatus) {
	o.Status = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Alertadmin_alertObject) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alertadmin_alertObject) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Alertadmin_alertObject) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *Alertadmin_alertObject) SetCount(v int32) {
	o.Count = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Alertadmin_alertObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AlertName != nil {
		toSerialize["alertName"] = o.AlertName
	}
	if o.GroupName != nil {
		toSerialize["groupName"] = o.GroupName
	}
	if o.Expr != nil {
		toSerialize["expr"] = o.Expr
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Pod != nil {
		toSerialize["pod"] = o.Pod
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Fingerprint != nil {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if o.StartsAt != nil {
		if o.StartsAt.Nanosecond() == 0 {
			toSerialize["startsAt"] = o.StartsAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["startsAt"] = o.StartsAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.EndsAt != nil {
		if o.EndsAt.Nanosecond() == 0 {
			toSerialize["endsAt"] = o.EndsAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["endsAt"] = o.EndsAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Alertadmin_alertObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *int32         `json:"id,omitempty"`
		AlertName   *string        `json:"alertName,omitempty"`
		GroupName   *string        `json:"groupName,omitempty"`
		Expr        *string        `json:"expr,omitempty"`
		Engine      *string        `json:"engine,omitempty"`
		Namespace   *string        `json:"namespace,omitempty"`
		Pod         *string        `json:"pod,omitempty"`
		Severity    *AlertSeverity `json:"severity,omitempty"`
		Description *string        `json:"description,omitempty"`
		Fingerprint *string        `json:"fingerprint,omitempty"`
		StartsAt    *time.Time     `json:"startsAt,omitempty"`
		EndsAt      *time.Time     `json:"endsAt,omitempty"`
		Status      *AlertStatus   `json:"status,omitempty"`
		Count       *int32         `json:"count,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "alertName", "groupName", "expr", "engine", "namespace", "pod", "severity", "description", "fingerprint", "startsAt", "endsAt", "status", "count"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.AlertName = all.AlertName
	o.GroupName = all.GroupName
	o.Expr = all.Expr
	o.Engine = all.Engine
	o.Namespace = all.Namespace
	o.Pod = all.Pod
	if all.Severity != nil && !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = all.Severity
	}
	o.Description = all.Description
	o.Fingerprint = all.Fingerprint
	o.StartsAt = all.StartsAt
	o.EndsAt = all.EndsAt
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Count = all.Count

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
