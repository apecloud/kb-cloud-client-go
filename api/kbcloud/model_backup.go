// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Backup backup is the payload for KubeBlocks cluster backup
type Backup struct {
	// autoBackup or not
	AutoBackup bool `json:"autoBackup"`
	// Backup Method.
	BackupMethod string `json:"backupMethod"`
	// Which backupPolicy is applied to perform this backup
	BackupPolicyName string `json:"backupPolicyName"`
	// backupRepo is the name of backupRepo and it is used to store the backup data
	BackupRepo *string `json:"backupRepo,omitempty"`
	// the type of backup
	BackupType BackupType `json:"backupType"`
	// Date/time when the backup finished being processed.
	CompletionTimestamp common.NullableTime `json:"completionTimestamp,omitempty"`
	// Date/time when the backup was created.
	CreationTimestamp time.Time `json:"creationTimestamp"`
	// The duration time of backup execution. When converted to a string, the form is "1h2m0.5s".
	Duration *string `json:"duration,omitempty"`
	// name of the backup
	Name string `json:"name"`
	// orgName records the organization name for this backup.
	OrgName string `json:"orgName"`
	// snapshotVolumes specifies whether to take snapshots of persistent volumes to back up
	SnapshotVolumes bool `json:"snapshotVolumes"`
	// sourceCluster records the source cluster information for this backup.
	SourceCluster string `json:"sourceCluster"`
	// Date/time when the backup started being processed.
	StartTimestamp common.NullableTime `json:"startTimestamp,omitempty"`
	// The current status. Valid values are New, InProgress, Completed, Failed.
	Status BackupStatus `json:"status"`
	// timeRangeEnd records the end time of the backup.
	TimeRangeEnd common.NullableTime `json:"timeRangeEnd,omitempty"`
	// timeRangeStart records the start time of the backup.
	TimeRangeStart common.NullableTime `json:"timeRangeStart,omitempty"`
	// Backup total size. A string with capacity units in the form of "1Gi", "1Mi", "1Ki".
	TotalSize     string  `json:"totalSize"`
	FailureReason *string `json:"failureReason,omitempty"`
	Extras        *string `json:"extras,omitempty"`
	// backup target pods
	TargetPods []string `json:"targetPods,omitempty"`
	// the path of backup files
	Path *string `json:"path,omitempty"`
	// determines a duration up to which the backup should be kept
	RetentionPeriod string `json:"retentionPeriod"`
	// indicates when this backup becomes eligible for garbage collection
	Expiration common.NullableTime `json:"expiration,omitempty"`
	// the backup id
	Id *string `json:"id,omitempty"`
	// the id of cluster that backup belong to
	ClusterId *string `json:"clusterId,omitempty"`
	// the cloud provider
	CloudProvider string `json:"cloudProvider"`
	// the cloud region
	CloudRegion string `json:"cloudRegion"`
	// the environment name
	EnvironmentName string `json:"environmentName"`
	// the cluster engine
	Engine string `json:"engine"`
	// the parent backup name
	ParentBackupName *string `json:"parentBackupName,omitempty"`
	// the base backup name
	BaseBackupName *string `json:"baseBackupName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackup instantiates a new Backup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackup(autoBackup bool, backupMethod string, backupPolicyName string, backupType BackupType, creationTimestamp time.Time, name string, orgName string, snapshotVolumes bool, sourceCluster string, status BackupStatus, totalSize string, retentionPeriod string, cloudProvider string, cloudRegion string, environmentName string, engine string) *Backup {
	this := Backup{}
	this.AutoBackup = autoBackup
	this.BackupMethod = backupMethod
	this.BackupPolicyName = backupPolicyName
	this.BackupType = backupType
	this.CreationTimestamp = creationTimestamp
	this.Name = name
	this.OrgName = orgName
	this.SnapshotVolumes = snapshotVolumes
	this.SourceCluster = sourceCluster
	this.Status = status
	this.TotalSize = totalSize
	this.RetentionPeriod = retentionPeriod
	this.CloudProvider = cloudProvider
	this.CloudRegion = cloudRegion
	this.EnvironmentName = environmentName
	this.Engine = engine
	return &this
}

// NewBackupWithDefaults instantiates a new Backup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupWithDefaults() *Backup {
	this := Backup{}
	return &this
}

// GetAutoBackup returns the AutoBackup field value.
func (o *Backup) GetAutoBackup() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AutoBackup
}

// GetAutoBackupOk returns a tuple with the AutoBackup field value
// and a boolean to check if the value has been set.
func (o *Backup) GetAutoBackupOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoBackup, true
}

// SetAutoBackup sets field value.
func (o *Backup) SetAutoBackup(v bool) {
	o.AutoBackup = v
}

// GetBackupMethod returns the BackupMethod field value.
func (o *Backup) GetBackupMethod() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BackupMethod
}

// GetBackupMethodOk returns a tuple with the BackupMethod field value
// and a boolean to check if the value has been set.
func (o *Backup) GetBackupMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackupMethod, true
}

// SetBackupMethod sets field value.
func (o *Backup) SetBackupMethod(v string) {
	o.BackupMethod = v
}

// GetBackupPolicyName returns the BackupPolicyName field value.
func (o *Backup) GetBackupPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BackupPolicyName
}

// GetBackupPolicyNameOk returns a tuple with the BackupPolicyName field value
// and a boolean to check if the value has been set.
func (o *Backup) GetBackupPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackupPolicyName, true
}

// SetBackupPolicyName sets field value.
func (o *Backup) SetBackupPolicyName(v string) {
	o.BackupPolicyName = v
}

// GetBackupRepo returns the BackupRepo field value if set, zero value otherwise.
func (o *Backup) GetBackupRepo() string {
	if o == nil || o.BackupRepo == nil {
		var ret string
		return ret
	}
	return *o.BackupRepo
}

// GetBackupRepoOk returns a tuple with the BackupRepo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetBackupRepoOk() (*string, bool) {
	if o == nil || o.BackupRepo == nil {
		return nil, false
	}
	return o.BackupRepo, true
}

// HasBackupRepo returns a boolean if a field has been set.
func (o *Backup) HasBackupRepo() bool {
	return o != nil && o.BackupRepo != nil
}

// SetBackupRepo gets a reference to the given string and assigns it to the BackupRepo field.
func (o *Backup) SetBackupRepo(v string) {
	o.BackupRepo = &v
}

// GetBackupType returns the BackupType field value.
func (o *Backup) GetBackupType() BackupType {
	if o == nil {
		var ret BackupType
		return ret
	}
	return o.BackupType
}

// GetBackupTypeOk returns a tuple with the BackupType field value
// and a boolean to check if the value has been set.
func (o *Backup) GetBackupTypeOk() (*BackupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackupType, true
}

// SetBackupType sets field value.
func (o *Backup) SetBackupType(v BackupType) {
	o.BackupType = v
}

// GetCompletionTimestamp returns the CompletionTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Backup) GetCompletionTimestamp() time.Time {
	if o == nil || o.CompletionTimestamp.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletionTimestamp.Get()
}

// GetCompletionTimestampOk returns a tuple with the CompletionTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Backup) GetCompletionTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletionTimestamp.Get(), o.CompletionTimestamp.IsSet()
}

// HasCompletionTimestamp returns a boolean if a field has been set.
func (o *Backup) HasCompletionTimestamp() bool {
	return o != nil && o.CompletionTimestamp.IsSet()
}

// SetCompletionTimestamp gets a reference to the given common.NullableTime and assigns it to the CompletionTimestamp field.
func (o *Backup) SetCompletionTimestamp(v time.Time) {
	o.CompletionTimestamp.Set(&v)
}

// SetCompletionTimestampNil sets the value for CompletionTimestamp to be an explicit nil.
func (o *Backup) SetCompletionTimestampNil() {
	o.CompletionTimestamp.Set(nil)
}

// UnsetCompletionTimestamp ensures that no value is present for CompletionTimestamp, not even an explicit nil.
func (o *Backup) UnsetCompletionTimestamp() {
	o.CompletionTimestamp.Unset()
}

// GetCreationTimestamp returns the CreationTimestamp field value.
func (o *Backup) GetCreationTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value
// and a boolean to check if the value has been set.
func (o *Backup) GetCreationTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTimestamp, true
}

// SetCreationTimestamp sets field value.
func (o *Backup) SetCreationTimestamp(v time.Time) {
	o.CreationTimestamp = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Backup) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Backup) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *Backup) SetDuration(v string) {
	o.Duration = &v
}

// GetName returns the Name field value.
func (o *Backup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Backup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Backup) SetName(v string) {
	o.Name = v
}

// GetOrgName returns the OrgName field value.
func (o *Backup) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *Backup) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value.
func (o *Backup) SetOrgName(v string) {
	o.OrgName = v
}

// GetSnapshotVolumes returns the SnapshotVolumes field value.
func (o *Backup) GetSnapshotVolumes() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.SnapshotVolumes
}

// GetSnapshotVolumesOk returns a tuple with the SnapshotVolumes field value
// and a boolean to check if the value has been set.
func (o *Backup) GetSnapshotVolumesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotVolumes, true
}

// SetSnapshotVolumes sets field value.
func (o *Backup) SetSnapshotVolumes(v bool) {
	o.SnapshotVolumes = v
}

// GetSourceCluster returns the SourceCluster field value.
func (o *Backup) GetSourceCluster() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceCluster
}

// GetSourceClusterOk returns a tuple with the SourceCluster field value
// and a boolean to check if the value has been set.
func (o *Backup) GetSourceClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceCluster, true
}

// SetSourceCluster sets field value.
func (o *Backup) SetSourceCluster(v string) {
	o.SourceCluster = v
}

// GetStartTimestamp returns the StartTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Backup) GetStartTimestamp() time.Time {
	if o == nil || o.StartTimestamp.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTimestamp.Get()
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Backup) GetStartTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTimestamp.Get(), o.StartTimestamp.IsSet()
}

// HasStartTimestamp returns a boolean if a field has been set.
func (o *Backup) HasStartTimestamp() bool {
	return o != nil && o.StartTimestamp.IsSet()
}

// SetStartTimestamp gets a reference to the given common.NullableTime and assigns it to the StartTimestamp field.
func (o *Backup) SetStartTimestamp(v time.Time) {
	o.StartTimestamp.Set(&v)
}

// SetStartTimestampNil sets the value for StartTimestamp to be an explicit nil.
func (o *Backup) SetStartTimestampNil() {
	o.StartTimestamp.Set(nil)
}

// UnsetStartTimestamp ensures that no value is present for StartTimestamp, not even an explicit nil.
func (o *Backup) UnsetStartTimestamp() {
	o.StartTimestamp.Unset()
}

// GetStatus returns the Status field value.
func (o *Backup) GetStatus() BackupStatus {
	if o == nil {
		var ret BackupStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Backup) GetStatusOk() (*BackupStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *Backup) SetStatus(v BackupStatus) {
	o.Status = v
}

// GetTimeRangeEnd returns the TimeRangeEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Backup) GetTimeRangeEnd() time.Time {
	if o == nil || o.TimeRangeEnd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.TimeRangeEnd.Get()
}

// GetTimeRangeEndOk returns a tuple with the TimeRangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Backup) GetTimeRangeEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeRangeEnd.Get(), o.TimeRangeEnd.IsSet()
}

// HasTimeRangeEnd returns a boolean if a field has been set.
func (o *Backup) HasTimeRangeEnd() bool {
	return o != nil && o.TimeRangeEnd.IsSet()
}

// SetTimeRangeEnd gets a reference to the given common.NullableTime and assigns it to the TimeRangeEnd field.
func (o *Backup) SetTimeRangeEnd(v time.Time) {
	o.TimeRangeEnd.Set(&v)
}

// SetTimeRangeEndNil sets the value for TimeRangeEnd to be an explicit nil.
func (o *Backup) SetTimeRangeEndNil() {
	o.TimeRangeEnd.Set(nil)
}

// UnsetTimeRangeEnd ensures that no value is present for TimeRangeEnd, not even an explicit nil.
func (o *Backup) UnsetTimeRangeEnd() {
	o.TimeRangeEnd.Unset()
}

// GetTimeRangeStart returns the TimeRangeStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Backup) GetTimeRangeStart() time.Time {
	if o == nil || o.TimeRangeStart.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.TimeRangeStart.Get()
}

// GetTimeRangeStartOk returns a tuple with the TimeRangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Backup) GetTimeRangeStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeRangeStart.Get(), o.TimeRangeStart.IsSet()
}

// HasTimeRangeStart returns a boolean if a field has been set.
func (o *Backup) HasTimeRangeStart() bool {
	return o != nil && o.TimeRangeStart.IsSet()
}

// SetTimeRangeStart gets a reference to the given common.NullableTime and assigns it to the TimeRangeStart field.
func (o *Backup) SetTimeRangeStart(v time.Time) {
	o.TimeRangeStart.Set(&v)
}

// SetTimeRangeStartNil sets the value for TimeRangeStart to be an explicit nil.
func (o *Backup) SetTimeRangeStartNil() {
	o.TimeRangeStart.Set(nil)
}

// UnsetTimeRangeStart ensures that no value is present for TimeRangeStart, not even an explicit nil.
func (o *Backup) UnsetTimeRangeStart() {
	o.TimeRangeStart.Unset()
}

// GetTotalSize returns the TotalSize field value.
func (o *Backup) GetTotalSize() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value
// and a boolean to check if the value has been set.
func (o *Backup) GetTotalSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSize, true
}

// SetTotalSize sets field value.
func (o *Backup) SetTotalSize(v string) {
	o.TotalSize = v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *Backup) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *Backup) HasFailureReason() bool {
	return o != nil && o.FailureReason != nil
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *Backup) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *Backup) GetExtras() string {
	if o == nil || o.Extras == nil {
		var ret string
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetExtrasOk() (*string, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *Backup) HasExtras() bool {
	return o != nil && o.Extras != nil
}

// SetExtras gets a reference to the given string and assigns it to the Extras field.
func (o *Backup) SetExtras(v string) {
	o.Extras = &v
}

// GetTargetPods returns the TargetPods field value if set, zero value otherwise.
func (o *Backup) GetTargetPods() []string {
	if o == nil || o.TargetPods == nil {
		var ret []string
		return ret
	}
	return o.TargetPods
}

// GetTargetPodsOk returns a tuple with the TargetPods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetTargetPodsOk() (*[]string, bool) {
	if o == nil || o.TargetPods == nil {
		return nil, false
	}
	return &o.TargetPods, true
}

// HasTargetPods returns a boolean if a field has been set.
func (o *Backup) HasTargetPods() bool {
	return o != nil && o.TargetPods != nil
}

// SetTargetPods gets a reference to the given []string and assigns it to the TargetPods field.
func (o *Backup) SetTargetPods(v []string) {
	o.TargetPods = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Backup) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Backup) HasPath() bool {
	return o != nil && o.Path != nil
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Backup) SetPath(v string) {
	o.Path = &v
}

// GetRetentionPeriod returns the RetentionPeriod field value.
func (o *Backup) GetRetentionPeriod() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value
// and a boolean to check if the value has been set.
func (o *Backup) GetRetentionPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionPeriod, true
}

// SetRetentionPeriod sets field value.
func (o *Backup) SetRetentionPeriod(v string) {
	o.RetentionPeriod = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Backup) GetExpiration() time.Time {
	if o == nil || o.Expiration.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Backup) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// HasExpiration returns a boolean if a field has been set.
func (o *Backup) HasExpiration() bool {
	return o != nil && o.Expiration.IsSet()
}

// SetExpiration gets a reference to the given common.NullableTime and assigns it to the Expiration field.
func (o *Backup) SetExpiration(v time.Time) {
	o.Expiration.Set(&v)
}

// SetExpirationNil sets the value for Expiration to be an explicit nil.
func (o *Backup) SetExpirationNil() {
	o.Expiration.Set(nil)
}

// UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil.
func (o *Backup) UnsetExpiration() {
	o.Expiration.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Backup) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Backup) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Backup) SetId(v string) {
	o.Id = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *Backup) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *Backup) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *Backup) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetCloudProvider returns the CloudProvider field value.
func (o *Backup) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *Backup) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value.
func (o *Backup) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetCloudRegion returns the CloudRegion field value.
func (o *Backup) GetCloudRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CloudRegion
}

// GetCloudRegionOk returns a tuple with the CloudRegion field value
// and a boolean to check if the value has been set.
func (o *Backup) GetCloudRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudRegion, true
}

// SetCloudRegion sets field value.
func (o *Backup) SetCloudRegion(v string) {
	o.CloudRegion = v
}

// GetEnvironmentName returns the EnvironmentName field value.
func (o *Backup) GetEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value
// and a boolean to check if the value has been set.
func (o *Backup) GetEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentName, true
}

// SetEnvironmentName sets field value.
func (o *Backup) SetEnvironmentName(v string) {
	o.EnvironmentName = v
}

// GetEngine returns the Engine field value.
func (o *Backup) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *Backup) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *Backup) SetEngine(v string) {
	o.Engine = v
}

// GetParentBackupName returns the ParentBackupName field value if set, zero value otherwise.
func (o *Backup) GetParentBackupName() string {
	if o == nil || o.ParentBackupName == nil {
		var ret string
		return ret
	}
	return *o.ParentBackupName
}

// GetParentBackupNameOk returns a tuple with the ParentBackupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetParentBackupNameOk() (*string, bool) {
	if o == nil || o.ParentBackupName == nil {
		return nil, false
	}
	return o.ParentBackupName, true
}

// HasParentBackupName returns a boolean if a field has been set.
func (o *Backup) HasParentBackupName() bool {
	return o != nil && o.ParentBackupName != nil
}

// SetParentBackupName gets a reference to the given string and assigns it to the ParentBackupName field.
func (o *Backup) SetParentBackupName(v string) {
	o.ParentBackupName = &v
}

// GetBaseBackupName returns the BaseBackupName field value if set, zero value otherwise.
func (o *Backup) GetBaseBackupName() string {
	if o == nil || o.BaseBackupName == nil {
		var ret string
		return ret
	}
	return *o.BaseBackupName
}

// GetBaseBackupNameOk returns a tuple with the BaseBackupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetBaseBackupNameOk() (*string, bool) {
	if o == nil || o.BaseBackupName == nil {
		return nil, false
	}
	return o.BaseBackupName, true
}

// HasBaseBackupName returns a boolean if a field has been set.
func (o *Backup) HasBaseBackupName() bool {
	return o != nil && o.BaseBackupName != nil
}

// SetBaseBackupName gets a reference to the given string and assigns it to the BaseBackupName field.
func (o *Backup) SetBaseBackupName(v string) {
	o.BaseBackupName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Backup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["autoBackup"] = o.AutoBackup
	toSerialize["backupMethod"] = o.BackupMethod
	toSerialize["backupPolicyName"] = o.BackupPolicyName
	if o.BackupRepo != nil {
		toSerialize["backupRepo"] = o.BackupRepo
	}
	toSerialize["backupType"] = o.BackupType
	if o.CompletionTimestamp.IsSet() {
		toSerialize["completionTimestamp"] = o.CompletionTimestamp.Get()
	}
	if o.CreationTimestamp.Nanosecond() == 0 {
		toSerialize["creationTimestamp"] = o.CreationTimestamp.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["creationTimestamp"] = o.CreationTimestamp.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	toSerialize["name"] = o.Name
	toSerialize["orgName"] = o.OrgName
	toSerialize["snapshotVolumes"] = o.SnapshotVolumes
	toSerialize["sourceCluster"] = o.SourceCluster
	if o.StartTimestamp.IsSet() {
		toSerialize["startTimestamp"] = o.StartTimestamp.Get()
	}
	toSerialize["status"] = o.Status
	if o.TimeRangeEnd.IsSet() {
		toSerialize["timeRangeEnd"] = o.TimeRangeEnd.Get()
	}
	if o.TimeRangeStart.IsSet() {
		toSerialize["timeRangeStart"] = o.TimeRangeStart.Get()
	}
	toSerialize["totalSize"] = o.TotalSize
	if o.FailureReason != nil {
		toSerialize["failureReason"] = o.FailureReason
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.TargetPods != nil {
		toSerialize["targetPods"] = o.TargetPods
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	toSerialize["retentionPeriod"] = o.RetentionPeriod
	if o.Expiration.IsSet() {
		toSerialize["expiration"] = o.Expiration.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ClusterId != nil {
		toSerialize["clusterId"] = o.ClusterId
	}
	toSerialize["cloudProvider"] = o.CloudProvider
	toSerialize["cloudRegion"] = o.CloudRegion
	toSerialize["environmentName"] = o.EnvironmentName
	toSerialize["engine"] = o.Engine
	if o.ParentBackupName != nil {
		toSerialize["parentBackupName"] = o.ParentBackupName
	}
	if o.BaseBackupName != nil {
		toSerialize["baseBackupName"] = o.BaseBackupName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Backup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoBackup          *bool               `json:"autoBackup"`
		BackupMethod        *string             `json:"backupMethod"`
		BackupPolicyName    *string             `json:"backupPolicyName"`
		BackupRepo          *string             `json:"backupRepo,omitempty"`
		BackupType          *BackupType         `json:"backupType"`
		CompletionTimestamp common.NullableTime `json:"completionTimestamp,omitempty"`
		CreationTimestamp   *time.Time          `json:"creationTimestamp"`
		Duration            *string             `json:"duration,omitempty"`
		Name                *string             `json:"name"`
		OrgName             *string             `json:"orgName"`
		SnapshotVolumes     *bool               `json:"snapshotVolumes"`
		SourceCluster       *string             `json:"sourceCluster"`
		StartTimestamp      common.NullableTime `json:"startTimestamp,omitempty"`
		Status              *BackupStatus       `json:"status"`
		TimeRangeEnd        common.NullableTime `json:"timeRangeEnd,omitempty"`
		TimeRangeStart      common.NullableTime `json:"timeRangeStart,omitempty"`
		TotalSize           *string             `json:"totalSize"`
		FailureReason       *string             `json:"failureReason,omitempty"`
		Extras              *string             `json:"extras,omitempty"`
		TargetPods          []string            `json:"targetPods,omitempty"`
		Path                *string             `json:"path,omitempty"`
		RetentionPeriod     *string             `json:"retentionPeriod"`
		Expiration          common.NullableTime `json:"expiration,omitempty"`
		Id                  *string             `json:"id,omitempty"`
		ClusterId           *string             `json:"clusterId,omitempty"`
		CloudProvider       *string             `json:"cloudProvider"`
		CloudRegion         *string             `json:"cloudRegion"`
		EnvironmentName     *string             `json:"environmentName"`
		Engine              *string             `json:"engine"`
		ParentBackupName    *string             `json:"parentBackupName,omitempty"`
		BaseBackupName      *string             `json:"baseBackupName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.AutoBackup == nil {
		return fmt.Errorf("required field autoBackup missing")
	}
	if all.BackupMethod == nil {
		return fmt.Errorf("required field backupMethod missing")
	}
	if all.BackupPolicyName == nil {
		return fmt.Errorf("required field backupPolicyName missing")
	}
	if all.BackupType == nil {
		return fmt.Errorf("required field backupType missing")
	}
	if all.CreationTimestamp == nil {
		return fmt.Errorf("required field creationTimestamp missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OrgName == nil {
		return fmt.Errorf("required field orgName missing")
	}
	if all.SnapshotVolumes == nil {
		return fmt.Errorf("required field snapshotVolumes missing")
	}
	if all.SourceCluster == nil {
		return fmt.Errorf("required field sourceCluster missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.TotalSize == nil {
		return fmt.Errorf("required field totalSize missing")
	}
	if all.RetentionPeriod == nil {
		return fmt.Errorf("required field retentionPeriod missing")
	}
	if all.CloudProvider == nil {
		return fmt.Errorf("required field cloudProvider missing")
	}
	if all.CloudRegion == nil {
		return fmt.Errorf("required field cloudRegion missing")
	}
	if all.EnvironmentName == nil {
		return fmt.Errorf("required field environmentName missing")
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"autoBackup", "backupMethod", "backupPolicyName", "backupRepo", "backupType", "completionTimestamp", "creationTimestamp", "duration", "name", "orgName", "snapshotVolumes", "sourceCluster", "startTimestamp", "status", "timeRangeEnd", "timeRangeStart", "totalSize", "failureReason", "extras", "targetPods", "path", "retentionPeriod", "expiration", "id", "clusterId", "cloudProvider", "cloudRegion", "environmentName", "engine", "parentBackupName", "baseBackupName"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AutoBackup = *all.AutoBackup
	o.BackupMethod = *all.BackupMethod
	o.BackupPolicyName = *all.BackupPolicyName
	o.BackupRepo = all.BackupRepo
	if !all.BackupType.IsValid() {
		hasInvalidField = true
	} else {
		o.BackupType = *all.BackupType
	}
	o.CompletionTimestamp = all.CompletionTimestamp
	o.CreationTimestamp = *all.CreationTimestamp
	o.Duration = all.Duration
	o.Name = *all.Name
	o.OrgName = *all.OrgName
	o.SnapshotVolumes = *all.SnapshotVolumes
	o.SourceCluster = *all.SourceCluster
	o.StartTimestamp = all.StartTimestamp
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.TimeRangeEnd = all.TimeRangeEnd
	o.TimeRangeStart = all.TimeRangeStart
	o.TotalSize = *all.TotalSize
	o.FailureReason = all.FailureReason
	o.Extras = all.Extras
	o.TargetPods = all.TargetPods
	o.Path = all.Path
	o.RetentionPeriod = *all.RetentionPeriod
	o.Expiration = all.Expiration
	o.Id = all.Id
	o.ClusterId = all.ClusterId
	o.CloudProvider = *all.CloudProvider
	o.CloudRegion = *all.CloudRegion
	o.EnvironmentName = *all.EnvironmentName
	o.Engine = *all.Engine
	o.ParentBackupName = all.ParentBackupName
	o.BaseBackupName = all.BaseBackupName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
