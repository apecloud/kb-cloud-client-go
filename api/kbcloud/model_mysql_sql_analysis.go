// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MysqlSQLAnalysis struct {
	// Data source. M1 uses performance_schema.events_statements_summary_by_digest.
	Source string `json:"source"`
	// Source status. Expected values are available or unavailable.
	Status string `json:"status"`
	// Reason when status is unavailable, such as performance_schema_disabled, digest_consumer_disabled, permission_denied, or query_failed.
	UnavailableReason *string `json:"unavailableReason,omitempty"`
	// Backend collection timestamp in UTC. It is not a sampling-window end time.
	CollectedAt string `json:"collectedAt"`
	// Earliest FIRST_SEEN timestamp across currently visible digest rows, when available.
	FirstSeen *string `json:"firstSeen,omitempty"`
	// Latest LAST_SEEN timestamp across currently visible digest rows, when available.
	LastSeen *string `json:"lastSeen,omitempty"`
	// Effective row limit after backend normalization.
	Limit int64 `json:"limit"`
	// Effective sort key. Expected values are totalTime, meanTime, maxTime, or calls.
	OrderBy string `json:"orderBy"`
	// Total execution time in milliseconds across all visible digest rows, not limited to the returned top-N items.
	TotalTimeMsAll *float64 `json:"totalTimeMsAll,omitempty"`
	// Total execution count across all visible digest rows, not limited to the returned top-N items.
	CallsAll *int64 `json:"callsAll,omitempty"`
	// Number of statement digests not recorded because the digest table was full, when exposed by the server.
	DigestLost *int64 `json:"digestLost,omitempty"`
	// SQL fingerprint ranking rows from Performance Schema. The list may be empty when no statement statistics exist.
	Items []MysqlSQLFingerprint `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMysqlSQLAnalysis instantiates a new MysqlSQLAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMysqlSQLAnalysis(source string, status string, collectedAt string, limit int64, orderBy string, items []MysqlSQLFingerprint) *MysqlSQLAnalysis {
	this := MysqlSQLAnalysis{}
	this.Source = source
	this.Status = status
	this.CollectedAt = collectedAt
	this.Limit = limit
	this.OrderBy = orderBy
	this.Items = items
	return &this
}

// NewMysqlSQLAnalysisWithDefaults instantiates a new MysqlSQLAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMysqlSQLAnalysisWithDefaults() *MysqlSQLAnalysis {
	this := MysqlSQLAnalysis{}
	return &this
}

// GetSource returns the Source field value.
func (o *MysqlSQLAnalysis) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLAnalysis) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *MysqlSQLAnalysis) SetSource(v string) {
	o.Source = v
}

// GetStatus returns the Status field value.
func (o *MysqlSQLAnalysis) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLAnalysis) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *MysqlSQLAnalysis) SetStatus(v string) {
	o.Status = v
}

// GetUnavailableReason returns the UnavailableReason field value if set, zero value otherwise.
func (o *MysqlSQLAnalysis) GetUnavailableReason() string {
	if o == nil || o.UnavailableReason == nil {
		var ret string
		return ret
	}
	return *o.UnavailableReason
}

// GetUnavailableReasonOk returns a tuple with the UnavailableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlSQLAnalysis) GetUnavailableReasonOk() (*string, bool) {
	if o == nil || o.UnavailableReason == nil {
		return nil, false
	}
	return o.UnavailableReason, true
}

// HasUnavailableReason returns a boolean if a field has been set.
func (o *MysqlSQLAnalysis) HasUnavailableReason() bool {
	return o != nil && o.UnavailableReason != nil
}

// SetUnavailableReason gets a reference to the given string and assigns it to the UnavailableReason field.
func (o *MysqlSQLAnalysis) SetUnavailableReason(v string) {
	o.UnavailableReason = &v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *MysqlSQLAnalysis) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLAnalysis) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *MysqlSQLAnalysis) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// GetFirstSeen returns the FirstSeen field value if set, zero value otherwise.
func (o *MysqlSQLAnalysis) GetFirstSeen() string {
	if o == nil || o.FirstSeen == nil {
		var ret string
		return ret
	}
	return *o.FirstSeen
}

// GetFirstSeenOk returns a tuple with the FirstSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlSQLAnalysis) GetFirstSeenOk() (*string, bool) {
	if o == nil || o.FirstSeen == nil {
		return nil, false
	}
	return o.FirstSeen, true
}

// HasFirstSeen returns a boolean if a field has been set.
func (o *MysqlSQLAnalysis) HasFirstSeen() bool {
	return o != nil && o.FirstSeen != nil
}

// SetFirstSeen gets a reference to the given string and assigns it to the FirstSeen field.
func (o *MysqlSQLAnalysis) SetFirstSeen(v string) {
	o.FirstSeen = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *MysqlSQLAnalysis) GetLastSeen() string {
	if o == nil || o.LastSeen == nil {
		var ret string
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlSQLAnalysis) GetLastSeenOk() (*string, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *MysqlSQLAnalysis) HasLastSeen() bool {
	return o != nil && o.LastSeen != nil
}

// SetLastSeen gets a reference to the given string and assigns it to the LastSeen field.
func (o *MysqlSQLAnalysis) SetLastSeen(v string) {
	o.LastSeen = &v
}

// GetLimit returns the Limit field value.
func (o *MysqlSQLAnalysis) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLAnalysis) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *MysqlSQLAnalysis) SetLimit(v int64) {
	o.Limit = v
}

// GetOrderBy returns the OrderBy field value.
func (o *MysqlSQLAnalysis) GetOrderBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLAnalysis) GetOrderByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderBy, true
}

// SetOrderBy sets field value.
func (o *MysqlSQLAnalysis) SetOrderBy(v string) {
	o.OrderBy = v
}

// GetTotalTimeMsAll returns the TotalTimeMsAll field value if set, zero value otherwise.
func (o *MysqlSQLAnalysis) GetTotalTimeMsAll() float64 {
	if o == nil || o.TotalTimeMsAll == nil {
		var ret float64
		return ret
	}
	return *o.TotalTimeMsAll
}

// GetTotalTimeMsAllOk returns a tuple with the TotalTimeMsAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlSQLAnalysis) GetTotalTimeMsAllOk() (*float64, bool) {
	if o == nil || o.TotalTimeMsAll == nil {
		return nil, false
	}
	return o.TotalTimeMsAll, true
}

// HasTotalTimeMsAll returns a boolean if a field has been set.
func (o *MysqlSQLAnalysis) HasTotalTimeMsAll() bool {
	return o != nil && o.TotalTimeMsAll != nil
}

// SetTotalTimeMsAll gets a reference to the given float64 and assigns it to the TotalTimeMsAll field.
func (o *MysqlSQLAnalysis) SetTotalTimeMsAll(v float64) {
	o.TotalTimeMsAll = &v
}

// GetCallsAll returns the CallsAll field value if set, zero value otherwise.
func (o *MysqlSQLAnalysis) GetCallsAll() int64 {
	if o == nil || o.CallsAll == nil {
		var ret int64
		return ret
	}
	return *o.CallsAll
}

// GetCallsAllOk returns a tuple with the CallsAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlSQLAnalysis) GetCallsAllOk() (*int64, bool) {
	if o == nil || o.CallsAll == nil {
		return nil, false
	}
	return o.CallsAll, true
}

// HasCallsAll returns a boolean if a field has been set.
func (o *MysqlSQLAnalysis) HasCallsAll() bool {
	return o != nil && o.CallsAll != nil
}

// SetCallsAll gets a reference to the given int64 and assigns it to the CallsAll field.
func (o *MysqlSQLAnalysis) SetCallsAll(v int64) {
	o.CallsAll = &v
}

// GetDigestLost returns the DigestLost field value if set, zero value otherwise.
func (o *MysqlSQLAnalysis) GetDigestLost() int64 {
	if o == nil || o.DigestLost == nil {
		var ret int64
		return ret
	}
	return *o.DigestLost
}

// GetDigestLostOk returns a tuple with the DigestLost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlSQLAnalysis) GetDigestLostOk() (*int64, bool) {
	if o == nil || o.DigestLost == nil {
		return nil, false
	}
	return o.DigestLost, true
}

// HasDigestLost returns a boolean if a field has been set.
func (o *MysqlSQLAnalysis) HasDigestLost() bool {
	return o != nil && o.DigestLost != nil
}

// SetDigestLost gets a reference to the given int64 and assigns it to the DigestLost field.
func (o *MysqlSQLAnalysis) SetDigestLost(v int64) {
	o.DigestLost = &v
}

// GetItems returns the Items field value.
func (o *MysqlSQLAnalysis) GetItems() []MysqlSQLFingerprint {
	if o == nil {
		var ret []MysqlSQLFingerprint
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLAnalysis) GetItemsOk() (*[]MysqlSQLFingerprint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *MysqlSQLAnalysis) SetItems(v []MysqlSQLFingerprint) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MysqlSQLAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["source"] = o.Source
	toSerialize["status"] = o.Status
	if o.UnavailableReason != nil {
		toSerialize["unavailableReason"] = o.UnavailableReason
	}
	toSerialize["collectedAt"] = o.CollectedAt
	if o.FirstSeen != nil {
		toSerialize["firstSeen"] = o.FirstSeen
	}
	if o.LastSeen != nil {
		toSerialize["lastSeen"] = o.LastSeen
	}
	toSerialize["limit"] = o.Limit
	toSerialize["orderBy"] = o.OrderBy
	if o.TotalTimeMsAll != nil {
		toSerialize["totalTimeMsAll"] = o.TotalTimeMsAll
	}
	if o.CallsAll != nil {
		toSerialize["callsAll"] = o.CallsAll
	}
	if o.DigestLost != nil {
		toSerialize["digestLost"] = o.DigestLost
	}
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MysqlSQLAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Source            *string                `json:"source"`
		Status            *string                `json:"status"`
		UnavailableReason *string                `json:"unavailableReason,omitempty"`
		CollectedAt       *string                `json:"collectedAt"`
		FirstSeen         *string                `json:"firstSeen,omitempty"`
		LastSeen          *string                `json:"lastSeen,omitempty"`
		Limit             *int64                 `json:"limit"`
		OrderBy           *string                `json:"orderBy"`
		TotalTimeMsAll    *float64               `json:"totalTimeMsAll,omitempty"`
		CallsAll          *int64                 `json:"callsAll,omitempty"`
		DigestLost        *int64                 `json:"digestLost,omitempty"`
		Items             *[]MysqlSQLFingerprint `json:"items"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	if all.OrderBy == nil {
		return fmt.Errorf("required field orderBy missing")
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"source", "status", "unavailableReason", "collectedAt", "firstSeen", "lastSeen", "limit", "orderBy", "totalTimeMsAll", "callsAll", "digestLost", "items"})
	} else {
		return err
	}
	o.Source = *all.Source
	o.Status = *all.Status
	o.UnavailableReason = all.UnavailableReason
	o.CollectedAt = *all.CollectedAt
	o.FirstSeen = all.FirstSeen
	o.LastSeen = all.LastSeen
	o.Limit = *all.Limit
	o.OrderBy = *all.OrderBy
	o.TotalTimeMsAll = all.TotalTimeMsAll
	o.CallsAll = all.CallsAll
	o.DigestLost = all.DigestLost
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
