// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type PostgresqlStorageInstanceUsage struct {
	// Kubernetes pod name for this PostgreSQL replica storage sample.
	InstanceName *string `json:"instanceName,omitempty"`
	// Kubernetes pod name for compatibility with Kubernetes-level troubleshooting.
	PodName *string `json:"podName,omitempty"`
	// PersistentVolumeClaim name when it can be read from metrics labels.
	PvcName *string `json:"pvcName,omitempty"`
	// PostgreSQL replica role normalized from metrics labels.
	Role *PostgresqlStorageInstanceRole `json:"role,omitempty"`
	// Optional KubeBlocks component name when it can be read from metrics labels.
	ComponentName *string `json:"componentName,omitempty"`
	// Physical PVC capacity in bytes for this replica storage sample.
	TotalBytes *int64 `json:"totalBytes,omitempty"`
	// Physical PVC used bytes for this replica storage sample.
	UsedBytes *int64 `json:"usedBytes,omitempty"`
	// Available bytes derived from totalBytes - usedBytes when both values are available.
	AvailableBytes *int64 `json:"availableBytes,omitempty"`
	// usedBytes / totalBytes for this replica storage sample when both values are available.
	UsageRatio *float64 `json:"usageRatio,omitempty"`
	// Metrics collection timestamp in UTC.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	Source    *string `json:"source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlStorageInstanceUsage instantiates a new PostgresqlStorageInstanceUsage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlStorageInstanceUsage() *PostgresqlStorageInstanceUsage {
	this := PostgresqlStorageInstanceUsage{}
	return &this
}

// NewPostgresqlStorageInstanceUsageWithDefaults instantiates a new PostgresqlStorageInstanceUsage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlStorageInstanceUsageWithDefaults() *PostgresqlStorageInstanceUsage {
	this := PostgresqlStorageInstanceUsage{}
	return &this
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *PostgresqlStorageInstanceUsage) GetInstanceName() string {
	if o == nil || o.InstanceName == nil {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageInstanceUsage) GetInstanceNameOk() (*string, bool) {
	if o == nil || o.InstanceName == nil {
		return nil, false
	}
	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *PostgresqlStorageInstanceUsage) HasInstanceName() bool {
	return o != nil && o.InstanceName != nil
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *PostgresqlStorageInstanceUsage) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetPodName returns the PodName field value if set, zero value otherwise.
func (o *PostgresqlStorageInstanceUsage) GetPodName() string {
	if o == nil || o.PodName == nil {
		var ret string
		return ret
	}
	return *o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageInstanceUsage) GetPodNameOk() (*string, bool) {
	if o == nil || o.PodName == nil {
		return nil, false
	}
	return o.PodName, true
}

// HasPodName returns a boolean if a field has been set.
func (o *PostgresqlStorageInstanceUsage) HasPodName() bool {
	return o != nil && o.PodName != nil
}

// SetPodName gets a reference to the given string and assigns it to the PodName field.
func (o *PostgresqlStorageInstanceUsage) SetPodName(v string) {
	o.PodName = &v
}

// GetPvcName returns the PvcName field value if set, zero value otherwise.
func (o *PostgresqlStorageInstanceUsage) GetPvcName() string {
	if o == nil || o.PvcName == nil {
		var ret string
		return ret
	}
	return *o.PvcName
}

// GetPvcNameOk returns a tuple with the PvcName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageInstanceUsage) GetPvcNameOk() (*string, bool) {
	if o == nil || o.PvcName == nil {
		return nil, false
	}
	return o.PvcName, true
}

// HasPvcName returns a boolean if a field has been set.
func (o *PostgresqlStorageInstanceUsage) HasPvcName() bool {
	return o != nil && o.PvcName != nil
}

// SetPvcName gets a reference to the given string and assigns it to the PvcName field.
func (o *PostgresqlStorageInstanceUsage) SetPvcName(v string) {
	o.PvcName = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *PostgresqlStorageInstanceUsage) GetRole() PostgresqlStorageInstanceRole {
	if o == nil || o.Role == nil {
		var ret PostgresqlStorageInstanceRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageInstanceUsage) GetRoleOk() (*PostgresqlStorageInstanceRole, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *PostgresqlStorageInstanceUsage) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given PostgresqlStorageInstanceRole and assigns it to the Role field.
func (o *PostgresqlStorageInstanceUsage) SetRole(v PostgresqlStorageInstanceRole) {
	o.Role = &v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *PostgresqlStorageInstanceUsage) GetComponentName() string {
	if o == nil || o.ComponentName == nil {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageInstanceUsage) GetComponentNameOk() (*string, bool) {
	if o == nil || o.ComponentName == nil {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *PostgresqlStorageInstanceUsage) HasComponentName() bool {
	return o != nil && o.ComponentName != nil
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *PostgresqlStorageInstanceUsage) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetTotalBytes returns the TotalBytes field value if set, zero value otherwise.
func (o *PostgresqlStorageInstanceUsage) GetTotalBytes() int64 {
	if o == nil || o.TotalBytes == nil {
		var ret int64
		return ret
	}
	return *o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageInstanceUsage) GetTotalBytesOk() (*int64, bool) {
	if o == nil || o.TotalBytes == nil {
		return nil, false
	}
	return o.TotalBytes, true
}

// HasTotalBytes returns a boolean if a field has been set.
func (o *PostgresqlStorageInstanceUsage) HasTotalBytes() bool {
	return o != nil && o.TotalBytes != nil
}

// SetTotalBytes gets a reference to the given int64 and assigns it to the TotalBytes field.
func (o *PostgresqlStorageInstanceUsage) SetTotalBytes(v int64) {
	o.TotalBytes = &v
}

// GetUsedBytes returns the UsedBytes field value if set, zero value otherwise.
func (o *PostgresqlStorageInstanceUsage) GetUsedBytes() int64 {
	if o == nil || o.UsedBytes == nil {
		var ret int64
		return ret
	}
	return *o.UsedBytes
}

// GetUsedBytesOk returns a tuple with the UsedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageInstanceUsage) GetUsedBytesOk() (*int64, bool) {
	if o == nil || o.UsedBytes == nil {
		return nil, false
	}
	return o.UsedBytes, true
}

// HasUsedBytes returns a boolean if a field has been set.
func (o *PostgresqlStorageInstanceUsage) HasUsedBytes() bool {
	return o != nil && o.UsedBytes != nil
}

// SetUsedBytes gets a reference to the given int64 and assigns it to the UsedBytes field.
func (o *PostgresqlStorageInstanceUsage) SetUsedBytes(v int64) {
	o.UsedBytes = &v
}

// GetAvailableBytes returns the AvailableBytes field value if set, zero value otherwise.
func (o *PostgresqlStorageInstanceUsage) GetAvailableBytes() int64 {
	if o == nil || o.AvailableBytes == nil {
		var ret int64
		return ret
	}
	return *o.AvailableBytes
}

// GetAvailableBytesOk returns a tuple with the AvailableBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageInstanceUsage) GetAvailableBytesOk() (*int64, bool) {
	if o == nil || o.AvailableBytes == nil {
		return nil, false
	}
	return o.AvailableBytes, true
}

// HasAvailableBytes returns a boolean if a field has been set.
func (o *PostgresqlStorageInstanceUsage) HasAvailableBytes() bool {
	return o != nil && o.AvailableBytes != nil
}

// SetAvailableBytes gets a reference to the given int64 and assigns it to the AvailableBytes field.
func (o *PostgresqlStorageInstanceUsage) SetAvailableBytes(v int64) {
	o.AvailableBytes = &v
}

// GetUsageRatio returns the UsageRatio field value if set, zero value otherwise.
func (o *PostgresqlStorageInstanceUsage) GetUsageRatio() float64 {
	if o == nil || o.UsageRatio == nil {
		var ret float64
		return ret
	}
	return *o.UsageRatio
}

// GetUsageRatioOk returns a tuple with the UsageRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageInstanceUsage) GetUsageRatioOk() (*float64, bool) {
	if o == nil || o.UsageRatio == nil {
		return nil, false
	}
	return o.UsageRatio, true
}

// HasUsageRatio returns a boolean if a field has been set.
func (o *PostgresqlStorageInstanceUsage) HasUsageRatio() bool {
	return o != nil && o.UsageRatio != nil
}

// SetUsageRatio gets a reference to the given float64 and assigns it to the UsageRatio field.
func (o *PostgresqlStorageInstanceUsage) SetUsageRatio(v float64) {
	o.UsageRatio = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PostgresqlStorageInstanceUsage) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageInstanceUsage) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PostgresqlStorageInstanceUsage) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *PostgresqlStorageInstanceUsage) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PostgresqlStorageInstanceUsage) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageInstanceUsage) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PostgresqlStorageInstanceUsage) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *PostgresqlStorageInstanceUsage) SetSource(v string) {
	o.Source = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlStorageInstanceUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.InstanceName != nil {
		toSerialize["instanceName"] = o.InstanceName
	}
	if o.PodName != nil {
		toSerialize["podName"] = o.PodName
	}
	if o.PvcName != nil {
		toSerialize["pvcName"] = o.PvcName
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.ComponentName != nil {
		toSerialize["componentName"] = o.ComponentName
	}
	if o.TotalBytes != nil {
		toSerialize["totalBytes"] = o.TotalBytes
	}
	if o.UsedBytes != nil {
		toSerialize["usedBytes"] = o.UsedBytes
	}
	if o.AvailableBytes != nil {
		toSerialize["availableBytes"] = o.AvailableBytes
	}
	if o.UsageRatio != nil {
		toSerialize["usageRatio"] = o.UsageRatio
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlStorageInstanceUsage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceName   *string                        `json:"instanceName,omitempty"`
		PodName        *string                        `json:"podName,omitempty"`
		PvcName        *string                        `json:"pvcName,omitempty"`
		Role           *PostgresqlStorageInstanceRole `json:"role,omitempty"`
		ComponentName  *string                        `json:"componentName,omitempty"`
		TotalBytes     *int64                         `json:"totalBytes,omitempty"`
		UsedBytes      *int64                         `json:"usedBytes,omitempty"`
		AvailableBytes *int64                         `json:"availableBytes,omitempty"`
		UsageRatio     *float64                       `json:"usageRatio,omitempty"`
		UpdatedAt      *string                        `json:"updatedAt,omitempty"`
		Source         *string                        `json:"source,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"instanceName", "podName", "pvcName", "role", "componentName", "totalBytes", "usedBytes", "availableBytes", "usageRatio", "updatedAt", "source"})
	} else {
		return err
	}

	hasInvalidField := false
	o.InstanceName = all.InstanceName
	o.PodName = all.PodName
	o.PvcName = all.PvcName
	if all.Role != nil && !all.Role.IsValid() {
		hasInvalidField = true
	} else {
		o.Role = all.Role
	}
	o.ComponentName = all.ComponentName
	o.TotalBytes = all.TotalBytes
	o.UsedBytes = all.UsedBytes
	o.AvailableBytes = all.AvailableBytes
	o.UsageRatio = all.UsageRatio
	o.UpdatedAt = all.UpdatedAt
	o.Source = all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
