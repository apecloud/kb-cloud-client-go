// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCPUAnalysisPodCPUEvidence Pod/container CPU evidence from Kubernetes metrics and resource requests/limits.
type DiagnosticCPUAnalysisPodCPUEvidence struct {
	// PostgreSQL primary Pod used for the CPU analysis.
	PodName string `json:"podName"`
	// Pod namespace.
	Namespace string `json:"namespace"`
	// PostgreSQL container name.
	ContainerName string `json:"containerName"`
	// Current CPU usage in millicores, when metrics are available.
	UsageMilliCores int64 `json:"usageMilliCores"`
	// CPU request in millicores, 0 when not configured.
	RequestMilliCores int64 `json:"requestMilliCores"`
	// CPU limit in millicores, 0 when not configured.
	LimitMilliCores int64 `json:"limitMilliCores"`
	// usage / request * 100 when request is non-zero.
	UsagePercentOfRequest *float64 `json:"usagePercentOfRequest,omitempty"`
	// usage / limit * 100 when limit is non-zero.
	UsagePercentOfLimit *float64 `json:"usagePercentOfLimit,omitempty"`
	// Whether Kubernetes metrics for the primary Pod were available.
	MetricsAvailable bool `json:"metricsAvailable"`
	// Redacted reason when CPU metrics are unavailable.
	MissingReason string `json:"missingReason"`
	// Collection timestamp.
	CollectedAt time.Time `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticCPUAnalysisPodCPUEvidence instantiates a new DiagnosticCPUAnalysisPodCPUEvidence object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticCPUAnalysisPodCPUEvidence(podName string, namespace string, containerName string, usageMilliCores int64, requestMilliCores int64, limitMilliCores int64, metricsAvailable bool, missingReason string, collectedAt time.Time) *DiagnosticCPUAnalysisPodCPUEvidence {
	this := DiagnosticCPUAnalysisPodCPUEvidence{}
	this.PodName = podName
	this.Namespace = namespace
	this.ContainerName = containerName
	this.UsageMilliCores = usageMilliCores
	this.RequestMilliCores = requestMilliCores
	this.LimitMilliCores = limitMilliCores
	this.MetricsAvailable = metricsAvailable
	this.MissingReason = missingReason
	this.CollectedAt = collectedAt
	return &this
}

// NewDiagnosticCPUAnalysisPodCPUEvidenceWithDefaults instantiates a new DiagnosticCPUAnalysisPodCPUEvidence object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticCPUAnalysisPodCPUEvidenceWithDefaults() *DiagnosticCPUAnalysisPodCPUEvidence {
	this := DiagnosticCPUAnalysisPodCPUEvidence{}
	return &this
}

// GetPodName returns the PodName field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetPodName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetPodNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodName, true
}

// SetPodName sets field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) SetPodName(v string) {
	o.PodName = v
}

// GetNamespace returns the Namespace field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) SetNamespace(v string) {
	o.Namespace = v
}

// GetContainerName returns the ContainerName field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetContainerName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetContainerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerName, true
}

// SetContainerName sets field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) SetContainerName(v string) {
	o.ContainerName = v
}

// GetUsageMilliCores returns the UsageMilliCores field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetUsageMilliCores() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UsageMilliCores
}

// GetUsageMilliCoresOk returns a tuple with the UsageMilliCores field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetUsageMilliCoresOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageMilliCores, true
}

// SetUsageMilliCores sets field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) SetUsageMilliCores(v int64) {
	o.UsageMilliCores = v
}

// GetRequestMilliCores returns the RequestMilliCores field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetRequestMilliCores() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RequestMilliCores
}

// GetRequestMilliCoresOk returns a tuple with the RequestMilliCores field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetRequestMilliCoresOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestMilliCores, true
}

// SetRequestMilliCores sets field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) SetRequestMilliCores(v int64) {
	o.RequestMilliCores = v
}

// GetLimitMilliCores returns the LimitMilliCores field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetLimitMilliCores() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LimitMilliCores
}

// GetLimitMilliCoresOk returns a tuple with the LimitMilliCores field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetLimitMilliCoresOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LimitMilliCores, true
}

// SetLimitMilliCores sets field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) SetLimitMilliCores(v int64) {
	o.LimitMilliCores = v
}

// GetUsagePercentOfRequest returns the UsagePercentOfRequest field value if set, zero value otherwise.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetUsagePercentOfRequest() float64 {
	if o == nil || o.UsagePercentOfRequest == nil {
		var ret float64
		return ret
	}
	return *o.UsagePercentOfRequest
}

// GetUsagePercentOfRequestOk returns a tuple with the UsagePercentOfRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetUsagePercentOfRequestOk() (*float64, bool) {
	if o == nil || o.UsagePercentOfRequest == nil {
		return nil, false
	}
	return o.UsagePercentOfRequest, true
}

// HasUsagePercentOfRequest returns a boolean if a field has been set.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) HasUsagePercentOfRequest() bool {
	return o != nil && o.UsagePercentOfRequest != nil
}

// SetUsagePercentOfRequest gets a reference to the given float64 and assigns it to the UsagePercentOfRequest field.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) SetUsagePercentOfRequest(v float64) {
	o.UsagePercentOfRequest = &v
}

// GetUsagePercentOfLimit returns the UsagePercentOfLimit field value if set, zero value otherwise.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetUsagePercentOfLimit() float64 {
	if o == nil || o.UsagePercentOfLimit == nil {
		var ret float64
		return ret
	}
	return *o.UsagePercentOfLimit
}

// GetUsagePercentOfLimitOk returns a tuple with the UsagePercentOfLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetUsagePercentOfLimitOk() (*float64, bool) {
	if o == nil || o.UsagePercentOfLimit == nil {
		return nil, false
	}
	return o.UsagePercentOfLimit, true
}

// HasUsagePercentOfLimit returns a boolean if a field has been set.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) HasUsagePercentOfLimit() bool {
	return o != nil && o.UsagePercentOfLimit != nil
}

// SetUsagePercentOfLimit gets a reference to the given float64 and assigns it to the UsagePercentOfLimit field.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) SetUsagePercentOfLimit(v float64) {
	o.UsagePercentOfLimit = &v
}

// GetMetricsAvailable returns the MetricsAvailable field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetMetricsAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.MetricsAvailable
}

// GetMetricsAvailableOk returns a tuple with the MetricsAvailable field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetMetricsAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricsAvailable, true
}

// SetMetricsAvailable sets field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) SetMetricsAvailable(v bool) {
	o.MetricsAvailable = v
}

// GetMissingReason returns the MissingReason field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetMissingReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MissingReason
}

// GetMissingReasonOk returns a tuple with the MissingReason field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetMissingReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MissingReason, true
}

// SetMissingReason sets field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) SetMissingReason(v string) {
	o.MissingReason = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetCollectedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) GetCollectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) SetCollectedAt(v time.Time) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticCPUAnalysisPodCPUEvidence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["podName"] = o.PodName
	toSerialize["namespace"] = o.Namespace
	toSerialize["containerName"] = o.ContainerName
	toSerialize["usageMilliCores"] = o.UsageMilliCores
	toSerialize["requestMilliCores"] = o.RequestMilliCores
	toSerialize["limitMilliCores"] = o.LimitMilliCores
	if o.UsagePercentOfRequest != nil {
		toSerialize["usagePercentOfRequest"] = o.UsagePercentOfRequest
	}
	if o.UsagePercentOfLimit != nil {
		toSerialize["usagePercentOfLimit"] = o.UsagePercentOfLimit
	}
	toSerialize["metricsAvailable"] = o.MetricsAvailable
	toSerialize["missingReason"] = o.MissingReason
	if o.CollectedAt.Nanosecond() == 0 {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticCPUAnalysisPodCPUEvidence) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PodName               *string    `json:"podName"`
		Namespace             *string    `json:"namespace"`
		ContainerName         *string    `json:"containerName"`
		UsageMilliCores       *int64     `json:"usageMilliCores"`
		RequestMilliCores     *int64     `json:"requestMilliCores"`
		LimitMilliCores       *int64     `json:"limitMilliCores"`
		UsagePercentOfRequest *float64   `json:"usagePercentOfRequest,omitempty"`
		UsagePercentOfLimit   *float64   `json:"usagePercentOfLimit,omitempty"`
		MetricsAvailable      *bool      `json:"metricsAvailable"`
		MissingReason         *string    `json:"missingReason"`
		CollectedAt           *time.Time `json:"collectedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.PodName == nil {
		return fmt.Errorf("required field podName missing")
	}
	if all.Namespace == nil {
		return fmt.Errorf("required field namespace missing")
	}
	if all.ContainerName == nil {
		return fmt.Errorf("required field containerName missing")
	}
	if all.UsageMilliCores == nil {
		return fmt.Errorf("required field usageMilliCores missing")
	}
	if all.RequestMilliCores == nil {
		return fmt.Errorf("required field requestMilliCores missing")
	}
	if all.LimitMilliCores == nil {
		return fmt.Errorf("required field limitMilliCores missing")
	}
	if all.MetricsAvailable == nil {
		return fmt.Errorf("required field metricsAvailable missing")
	}
	if all.MissingReason == nil {
		return fmt.Errorf("required field missingReason missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"podName", "namespace", "containerName", "usageMilliCores", "requestMilliCores", "limitMilliCores", "usagePercentOfRequest", "usagePercentOfLimit", "metricsAvailable", "missingReason", "collectedAt"})
	} else {
		return err
	}
	o.PodName = *all.PodName
	o.Namespace = *all.Namespace
	o.ContainerName = *all.ContainerName
	o.UsageMilliCores = *all.UsageMilliCores
	o.RequestMilliCores = *all.RequestMilliCores
	o.LimitMilliCores = *all.LimitMilliCores
	o.UsagePercentOfRequest = all.UsagePercentOfRequest
	o.UsagePercentOfLimit = all.UsagePercentOfLimit
	o.MetricsAvailable = *all.MetricsAvailable
	o.MissingReason = *all.MissingReason
	o.CollectedAt = *all.CollectedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
