// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// RedisAnalysisResponse Instant Redis key analysis report sampled by bounded SCAN. This aligns with RedisInsight-style analysis dimensions but is not a persisted historical report.
type RedisAnalysisResponse struct {
	Db     *int64               `json:"db,omitempty"`
	Filter *RedisAnalysisFilter `json:"filter,omitempty"`
	// Actual delimiter used to aggregate top namespaces.
	Delimiter *string                `json:"delimiter,omitempty"`
	Progress  *RedisAnalysisProgress `json:"progress,omitempty"`
	CreatedAt *time.Time             `json:"createdAt,omitempty"`
	// Connected client count from Redis INFO when available.
	ConnectedClients *int64                          `json:"connectedClients,omitempty"`
	TotalKeys        *RedisAnalysisTypedSummary      `json:"totalKeys,omitempty"`
	TotalMemory      *RedisAnalysisTypedSummary      `json:"totalMemory,omitempty"`
	TopKeysNsp       []RedisAnalysisNamespaceSummary `json:"topKeysNsp,omitempty"`
	TopMemoryNsp     []RedisAnalysisNamespaceSummary `json:"topMemoryNsp,omitempty"`
	TopKeysLength    []RedisAnalysisKey              `json:"topKeysLength,omitempty"`
	TopKeysMemory    []RedisAnalysisKey              `json:"topKeysMemory,omitempty"`
	ExpirationGroups []RedisAnalysisExpirationGroup  `json:"expirationGroups,omitempty"`
	Warnings         []string                        `json:"warnings,omitempty"`
	Recommendations  []string                        `json:"recommendations,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisAnalysisResponse instantiates a new RedisAnalysisResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisAnalysisResponse() *RedisAnalysisResponse {
	this := RedisAnalysisResponse{}
	return &this
}

// NewRedisAnalysisResponseWithDefaults instantiates a new RedisAnalysisResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisAnalysisResponseWithDefaults() *RedisAnalysisResponse {
	this := RedisAnalysisResponse{}
	return &this
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *RedisAnalysisResponse) GetDb() int64 {
	if o == nil || o.Db == nil {
		var ret int64
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisResponse) GetDbOk() (*int64, bool) {
	if o == nil || o.Db == nil {
		return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *RedisAnalysisResponse) HasDb() bool {
	return o != nil && o.Db != nil
}

// SetDb gets a reference to the given int64 and assigns it to the Db field.
func (o *RedisAnalysisResponse) SetDb(v int64) {
	o.Db = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *RedisAnalysisResponse) GetFilter() RedisAnalysisFilter {
	if o == nil || o.Filter == nil {
		var ret RedisAnalysisFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisResponse) GetFilterOk() (*RedisAnalysisFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *RedisAnalysisResponse) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given RedisAnalysisFilter and assigns it to the Filter field.
func (o *RedisAnalysisResponse) SetFilter(v RedisAnalysisFilter) {
	o.Filter = &v
}

// GetDelimiter returns the Delimiter field value if set, zero value otherwise.
func (o *RedisAnalysisResponse) GetDelimiter() string {
	if o == nil || o.Delimiter == nil {
		var ret string
		return ret
	}
	return *o.Delimiter
}

// GetDelimiterOk returns a tuple with the Delimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisResponse) GetDelimiterOk() (*string, bool) {
	if o == nil || o.Delimiter == nil {
		return nil, false
	}
	return o.Delimiter, true
}

// HasDelimiter returns a boolean if a field has been set.
func (o *RedisAnalysisResponse) HasDelimiter() bool {
	return o != nil && o.Delimiter != nil
}

// SetDelimiter gets a reference to the given string and assigns it to the Delimiter field.
func (o *RedisAnalysisResponse) SetDelimiter(v string) {
	o.Delimiter = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *RedisAnalysisResponse) GetProgress() RedisAnalysisProgress {
	if o == nil || o.Progress == nil {
		var ret RedisAnalysisProgress
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisResponse) GetProgressOk() (*RedisAnalysisProgress, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *RedisAnalysisResponse) HasProgress() bool {
	return o != nil && o.Progress != nil
}

// SetProgress gets a reference to the given RedisAnalysisProgress and assigns it to the Progress field.
func (o *RedisAnalysisResponse) SetProgress(v RedisAnalysisProgress) {
	o.Progress = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RedisAnalysisResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RedisAnalysisResponse) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RedisAnalysisResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetConnectedClients returns the ConnectedClients field value if set, zero value otherwise.
func (o *RedisAnalysisResponse) GetConnectedClients() int64 {
	if o == nil || o.ConnectedClients == nil {
		var ret int64
		return ret
	}
	return *o.ConnectedClients
}

// GetConnectedClientsOk returns a tuple with the ConnectedClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisResponse) GetConnectedClientsOk() (*int64, bool) {
	if o == nil || o.ConnectedClients == nil {
		return nil, false
	}
	return o.ConnectedClients, true
}

// HasConnectedClients returns a boolean if a field has been set.
func (o *RedisAnalysisResponse) HasConnectedClients() bool {
	return o != nil && o.ConnectedClients != nil
}

// SetConnectedClients gets a reference to the given int64 and assigns it to the ConnectedClients field.
func (o *RedisAnalysisResponse) SetConnectedClients(v int64) {
	o.ConnectedClients = &v
}

// GetTotalKeys returns the TotalKeys field value if set, zero value otherwise.
func (o *RedisAnalysisResponse) GetTotalKeys() RedisAnalysisTypedSummary {
	if o == nil || o.TotalKeys == nil {
		var ret RedisAnalysisTypedSummary
		return ret
	}
	return *o.TotalKeys
}

// GetTotalKeysOk returns a tuple with the TotalKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisResponse) GetTotalKeysOk() (*RedisAnalysisTypedSummary, bool) {
	if o == nil || o.TotalKeys == nil {
		return nil, false
	}
	return o.TotalKeys, true
}

// HasTotalKeys returns a boolean if a field has been set.
func (o *RedisAnalysisResponse) HasTotalKeys() bool {
	return o != nil && o.TotalKeys != nil
}

// SetTotalKeys gets a reference to the given RedisAnalysisTypedSummary and assigns it to the TotalKeys field.
func (o *RedisAnalysisResponse) SetTotalKeys(v RedisAnalysisTypedSummary) {
	o.TotalKeys = &v
}

// GetTotalMemory returns the TotalMemory field value if set, zero value otherwise.
func (o *RedisAnalysisResponse) GetTotalMemory() RedisAnalysisTypedSummary {
	if o == nil || o.TotalMemory == nil {
		var ret RedisAnalysisTypedSummary
		return ret
	}
	return *o.TotalMemory
}

// GetTotalMemoryOk returns a tuple with the TotalMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisResponse) GetTotalMemoryOk() (*RedisAnalysisTypedSummary, bool) {
	if o == nil || o.TotalMemory == nil {
		return nil, false
	}
	return o.TotalMemory, true
}

// HasTotalMemory returns a boolean if a field has been set.
func (o *RedisAnalysisResponse) HasTotalMemory() bool {
	return o != nil && o.TotalMemory != nil
}

// SetTotalMemory gets a reference to the given RedisAnalysisTypedSummary and assigns it to the TotalMemory field.
func (o *RedisAnalysisResponse) SetTotalMemory(v RedisAnalysisTypedSummary) {
	o.TotalMemory = &v
}

// GetTopKeysNsp returns the TopKeysNsp field value if set, zero value otherwise.
func (o *RedisAnalysisResponse) GetTopKeysNsp() []RedisAnalysisNamespaceSummary {
	if o == nil || o.TopKeysNsp == nil {
		var ret []RedisAnalysisNamespaceSummary
		return ret
	}
	return o.TopKeysNsp
}

// GetTopKeysNspOk returns a tuple with the TopKeysNsp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisResponse) GetTopKeysNspOk() (*[]RedisAnalysisNamespaceSummary, bool) {
	if o == nil || o.TopKeysNsp == nil {
		return nil, false
	}
	return &o.TopKeysNsp, true
}

// HasTopKeysNsp returns a boolean if a field has been set.
func (o *RedisAnalysisResponse) HasTopKeysNsp() bool {
	return o != nil && o.TopKeysNsp != nil
}

// SetTopKeysNsp gets a reference to the given []RedisAnalysisNamespaceSummary and assigns it to the TopKeysNsp field.
func (o *RedisAnalysisResponse) SetTopKeysNsp(v []RedisAnalysisNamespaceSummary) {
	o.TopKeysNsp = v
}

// GetTopMemoryNsp returns the TopMemoryNsp field value if set, zero value otherwise.
func (o *RedisAnalysisResponse) GetTopMemoryNsp() []RedisAnalysisNamespaceSummary {
	if o == nil || o.TopMemoryNsp == nil {
		var ret []RedisAnalysisNamespaceSummary
		return ret
	}
	return o.TopMemoryNsp
}

// GetTopMemoryNspOk returns a tuple with the TopMemoryNsp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisResponse) GetTopMemoryNspOk() (*[]RedisAnalysisNamespaceSummary, bool) {
	if o == nil || o.TopMemoryNsp == nil {
		return nil, false
	}
	return &o.TopMemoryNsp, true
}

// HasTopMemoryNsp returns a boolean if a field has been set.
func (o *RedisAnalysisResponse) HasTopMemoryNsp() bool {
	return o != nil && o.TopMemoryNsp != nil
}

// SetTopMemoryNsp gets a reference to the given []RedisAnalysisNamespaceSummary and assigns it to the TopMemoryNsp field.
func (o *RedisAnalysisResponse) SetTopMemoryNsp(v []RedisAnalysisNamespaceSummary) {
	o.TopMemoryNsp = v
}

// GetTopKeysLength returns the TopKeysLength field value if set, zero value otherwise.
func (o *RedisAnalysisResponse) GetTopKeysLength() []RedisAnalysisKey {
	if o == nil || o.TopKeysLength == nil {
		var ret []RedisAnalysisKey
		return ret
	}
	return o.TopKeysLength
}

// GetTopKeysLengthOk returns a tuple with the TopKeysLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisResponse) GetTopKeysLengthOk() (*[]RedisAnalysisKey, bool) {
	if o == nil || o.TopKeysLength == nil {
		return nil, false
	}
	return &o.TopKeysLength, true
}

// HasTopKeysLength returns a boolean if a field has been set.
func (o *RedisAnalysisResponse) HasTopKeysLength() bool {
	return o != nil && o.TopKeysLength != nil
}

// SetTopKeysLength gets a reference to the given []RedisAnalysisKey and assigns it to the TopKeysLength field.
func (o *RedisAnalysisResponse) SetTopKeysLength(v []RedisAnalysisKey) {
	o.TopKeysLength = v
}

// GetTopKeysMemory returns the TopKeysMemory field value if set, zero value otherwise.
func (o *RedisAnalysisResponse) GetTopKeysMemory() []RedisAnalysisKey {
	if o == nil || o.TopKeysMemory == nil {
		var ret []RedisAnalysisKey
		return ret
	}
	return o.TopKeysMemory
}

// GetTopKeysMemoryOk returns a tuple with the TopKeysMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisResponse) GetTopKeysMemoryOk() (*[]RedisAnalysisKey, bool) {
	if o == nil || o.TopKeysMemory == nil {
		return nil, false
	}
	return &o.TopKeysMemory, true
}

// HasTopKeysMemory returns a boolean if a field has been set.
func (o *RedisAnalysisResponse) HasTopKeysMemory() bool {
	return o != nil && o.TopKeysMemory != nil
}

// SetTopKeysMemory gets a reference to the given []RedisAnalysisKey and assigns it to the TopKeysMemory field.
func (o *RedisAnalysisResponse) SetTopKeysMemory(v []RedisAnalysisKey) {
	o.TopKeysMemory = v
}

// GetExpirationGroups returns the ExpirationGroups field value if set, zero value otherwise.
func (o *RedisAnalysisResponse) GetExpirationGroups() []RedisAnalysisExpirationGroup {
	if o == nil || o.ExpirationGroups == nil {
		var ret []RedisAnalysisExpirationGroup
		return ret
	}
	return o.ExpirationGroups
}

// GetExpirationGroupsOk returns a tuple with the ExpirationGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisResponse) GetExpirationGroupsOk() (*[]RedisAnalysisExpirationGroup, bool) {
	if o == nil || o.ExpirationGroups == nil {
		return nil, false
	}
	return &o.ExpirationGroups, true
}

// HasExpirationGroups returns a boolean if a field has been set.
func (o *RedisAnalysisResponse) HasExpirationGroups() bool {
	return o != nil && o.ExpirationGroups != nil
}

// SetExpirationGroups gets a reference to the given []RedisAnalysisExpirationGroup and assigns it to the ExpirationGroups field.
func (o *RedisAnalysisResponse) SetExpirationGroups(v []RedisAnalysisExpirationGroup) {
	o.ExpirationGroups = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *RedisAnalysisResponse) GetWarnings() []string {
	if o == nil || o.Warnings == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisResponse) GetWarningsOk() (*[]string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *RedisAnalysisResponse) HasWarnings() bool {
	return o != nil && o.Warnings != nil
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *RedisAnalysisResponse) SetWarnings(v []string) {
	o.Warnings = v
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise.
func (o *RedisAnalysisResponse) GetRecommendations() []string {
	if o == nil || o.Recommendations == nil {
		var ret []string
		return ret
	}
	return o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisResponse) GetRecommendationsOk() (*[]string, bool) {
	if o == nil || o.Recommendations == nil {
		return nil, false
	}
	return &o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *RedisAnalysisResponse) HasRecommendations() bool {
	return o != nil && o.Recommendations != nil
}

// SetRecommendations gets a reference to the given []string and assigns it to the Recommendations field.
func (o *RedisAnalysisResponse) SetRecommendations(v []string) {
	o.Recommendations = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisAnalysisResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Db != nil {
		toSerialize["db"] = o.Db
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Delimiter != nil {
		toSerialize["delimiter"] = o.Delimiter
	}
	if o.Progress != nil {
		toSerialize["progress"] = o.Progress
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ConnectedClients != nil {
		toSerialize["connectedClients"] = o.ConnectedClients
	}
	if o.TotalKeys != nil {
		toSerialize["totalKeys"] = o.TotalKeys
	}
	if o.TotalMemory != nil {
		toSerialize["totalMemory"] = o.TotalMemory
	}
	if o.TopKeysNsp != nil {
		toSerialize["topKeysNsp"] = o.TopKeysNsp
	}
	if o.TopMemoryNsp != nil {
		toSerialize["topMemoryNsp"] = o.TopMemoryNsp
	}
	if o.TopKeysLength != nil {
		toSerialize["topKeysLength"] = o.TopKeysLength
	}
	if o.TopKeysMemory != nil {
		toSerialize["topKeysMemory"] = o.TopKeysMemory
	}
	if o.ExpirationGroups != nil {
		toSerialize["expirationGroups"] = o.ExpirationGroups
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if o.Recommendations != nil {
		toSerialize["recommendations"] = o.Recommendations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisAnalysisResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Db               *int64                          `json:"db,omitempty"`
		Filter           *RedisAnalysisFilter            `json:"filter,omitempty"`
		Delimiter        *string                         `json:"delimiter,omitempty"`
		Progress         *RedisAnalysisProgress          `json:"progress,omitempty"`
		CreatedAt        *time.Time                      `json:"createdAt,omitempty"`
		ConnectedClients *int64                          `json:"connectedClients,omitempty"`
		TotalKeys        *RedisAnalysisTypedSummary      `json:"totalKeys,omitempty"`
		TotalMemory      *RedisAnalysisTypedSummary      `json:"totalMemory,omitempty"`
		TopKeysNsp       []RedisAnalysisNamespaceSummary `json:"topKeysNsp,omitempty"`
		TopMemoryNsp     []RedisAnalysisNamespaceSummary `json:"topMemoryNsp,omitempty"`
		TopKeysLength    []RedisAnalysisKey              `json:"topKeysLength,omitempty"`
		TopKeysMemory    []RedisAnalysisKey              `json:"topKeysMemory,omitempty"`
		ExpirationGroups []RedisAnalysisExpirationGroup  `json:"expirationGroups,omitempty"`
		Warnings         []string                        `json:"warnings,omitempty"`
		Recommendations  []string                        `json:"recommendations,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"db", "filter", "delimiter", "progress", "createdAt", "connectedClients", "totalKeys", "totalMemory", "topKeysNsp", "topMemoryNsp", "topKeysLength", "topKeysMemory", "expirationGroups", "warnings", "recommendations"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Db = all.Db
	if all.Filter != nil && all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = all.Filter
	o.Delimiter = all.Delimiter
	if all.Progress != nil && all.Progress.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Progress = all.Progress
	o.CreatedAt = all.CreatedAt
	o.ConnectedClients = all.ConnectedClients
	if all.TotalKeys != nil && all.TotalKeys.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TotalKeys = all.TotalKeys
	if all.TotalMemory != nil && all.TotalMemory.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TotalMemory = all.TotalMemory
	o.TopKeysNsp = all.TopKeysNsp
	o.TopMemoryNsp = all.TopMemoryNsp
	o.TopKeysLength = all.TopKeysLength
	o.TopKeysMemory = all.TopKeysMemory
	o.ExpirationGroups = all.ExpirationGroups
	o.Warnings = all.Warnings
	o.Recommendations = all.Recommendations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
