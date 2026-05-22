// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsQueryHistory struct {
	// history record id
	Id *int64 `json:"id,omitempty"`
	// organization name
	OrgName *string `json:"orgName,omitempty"`
	// cluster name
	ClusterName *string `json:"clusterName,omitempty"`
	// datasource id
	DatasourceId *int64 `json:"datasourceId,omitempty"`
	// datasource engine
	Engine *string `json:"engine,omitempty"`
	// user input text, such as SQL statements or MongoDB Shell commands
	Sql *string `json:"sql,omitempty"`
	// sql executed massage
	ErrMassage *string `json:"errMassage,omitempty"`
	// sql executed time
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// sql executed duration
	Duration *int64 `json:"duration,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsQueryHistory instantiates a new DmsQueryHistory object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsQueryHistory() *DmsQueryHistory {
	this := DmsQueryHistory{}
	return &this
}

// NewDmsQueryHistoryWithDefaults instantiates a new DmsQueryHistory object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsQueryHistoryWithDefaults() *DmsQueryHistory {
	this := DmsQueryHistory{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DmsQueryHistory) SetId(v int64) {
	o.Id = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *DmsQueryHistory) SetOrgName(v string) {
	o.OrgName = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *DmsQueryHistory) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetDatasourceId returns the DatasourceId field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetDatasourceId() int64 {
	if o == nil || o.DatasourceId == nil {
		var ret int64
		return ret
	}
	return *o.DatasourceId
}

// GetDatasourceIdOk returns a tuple with the DatasourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetDatasourceIdOk() (*int64, bool) {
	if o == nil || o.DatasourceId == nil {
		return nil, false
	}
	return o.DatasourceId, true
}

// HasDatasourceId returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasDatasourceId() bool {
	return o != nil && o.DatasourceId != nil
}

// SetDatasourceId gets a reference to the given int64 and assigns it to the DatasourceId field.
func (o *DmsQueryHistory) SetDatasourceId(v int64) {
	o.DatasourceId = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *DmsQueryHistory) SetEngine(v string) {
	o.Engine = &v
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetSql() string {
	if o == nil || o.Sql == nil {
		var ret string
		return ret
	}
	return *o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetSqlOk() (*string, bool) {
	if o == nil || o.Sql == nil {
		return nil, false
	}
	return o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasSql() bool {
	return o != nil && o.Sql != nil
}

// SetSql gets a reference to the given string and assigns it to the Sql field.
func (o *DmsQueryHistory) SetSql(v string) {
	o.Sql = &v
}

// GetErrMassage returns the ErrMassage field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetErrMassage() string {
	if o == nil || o.ErrMassage == nil {
		var ret string
		return ret
	}
	return *o.ErrMassage
}

// GetErrMassageOk returns a tuple with the ErrMassage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetErrMassageOk() (*string, bool) {
	if o == nil || o.ErrMassage == nil {
		return nil, false
	}
	return o.ErrMassage, true
}

// HasErrMassage returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasErrMassage() bool {
	return o != nil && o.ErrMassage != nil
}

// SetErrMassage gets a reference to the given string and assigns it to the ErrMassage field.
func (o *DmsQueryHistory) SetErrMassage(v string) {
	o.ErrMassage = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DmsQueryHistory) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *DmsQueryHistory) SetDuration(v int64) {
	o.Duration = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsQueryHistory) MarshalJSON() ([]byte, error) {
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
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.DatasourceId != nil {
		toSerialize["datasourceId"] = o.DatasourceId
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.Sql != nil {
		toSerialize["sql"] = o.Sql
	}
	if o.ErrMassage != nil {
		toSerialize["errMassage"] = o.ErrMassage
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsQueryHistory) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id           *int64     `json:"id,omitempty"`
		OrgName      *string    `json:"orgName,omitempty"`
		ClusterName  *string    `json:"clusterName,omitempty"`
		DatasourceId *int64     `json:"datasourceId,omitempty"`
		Engine       *string    `json:"engine,omitempty"`
		Sql          *string    `json:"sql,omitempty"`
		ErrMassage   *string    `json:"errMassage,omitempty"`
		CreatedAt    *time.Time `json:"createdAt,omitempty"`
		Duration     *int64     `json:"duration,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "orgName", "clusterName", "datasourceId", "engine", "sql", "errMassage", "createdAt", "duration"})
	} else {
		return err
	}
	o.Id = all.Id
	o.OrgName = all.OrgName
	o.ClusterName = all.ClusterName
	o.DatasourceId = all.DatasourceId
	o.Engine = all.Engine
	o.Sql = all.Sql
	o.ErrMassage = all.ErrMassage
	o.CreatedAt = all.CreatedAt
	o.Duration = all.Duration

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
