// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MysqlSQLFingerprint struct {
	// MySQL Performance Schema DIGEST value.
	QueryId string `json:"queryID"`
	// Stable SQL fingerprint identifier. M1 uses DIGEST.
	Fingerprint string `json:"fingerprint"`
	// Normalized DIGEST_TEXT. Raw SQL text is intentionally not returned.
	QuerySummary string `json:"querySummary"`
	// Number of executions accumulated in the current digest row.
	Calls int64 `json:"calls"`
	// Total statement wait time in milliseconds.
	TotalTimeMs float64 `json:"totalTimeMs"`
	// Mean statement wait time in milliseconds.
	MeanTimeMs float64 `json:"meanTimeMs"`
	// Maximum statement wait time in milliseconds.
	MaxTimeMs float64 `json:"maxTimeMs"`
	// Sum of rows sent and rows affected.
	Rows int64 `json:"rows"`
	// Total rows sent to clients.
	RowsSent int64 `json:"rowsSent"`
	// Total rows affected by statements.
	RowsAffected int64 `json:"rowsAffected"`
	// Total rows examined by statements.
	RowsExamined int64 `json:"rowsExamined"`
	// Total statement lock time in milliseconds.
	LockTimeMs float64 `json:"lockTimeMs"`
	// Total statement errors.
	Errors int64 `json:"errors"`
	// Total statement warnings.
	Warnings int64 `json:"warnings"`
	// Total internal temporary tables created.
	TmpTables int64 `json:"tmpTables"`
	// Total internal on-disk temporary tables created.
	TmpDiskTables int64 `json:"tmpDiskTables"`
	// Executions that used no index.
	NoIndexUsed int64 `json:"noIndexUsed"`
	// Executions for which no good index was found.
	NoGoodIndexUsed int64 `json:"noGoodIndexUsed"`
	// Default schema name associated with the digest row.
	Database string `json:"database"`
	// UTC timestamp when the digest was first observed.
	FirstSeen *string `json:"firstSeen,omitempty"`
	// UTC timestamp when the digest was most recently observed.
	LastSeen *string `json:"lastSeen,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMysqlSQLFingerprint instantiates a new MysqlSQLFingerprint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMysqlSQLFingerprint(queryId string, fingerprint string, querySummary string, calls int64, totalTimeMs float64, meanTimeMs float64, maxTimeMs float64, rows int64, rowsSent int64, rowsAffected int64, rowsExamined int64, lockTimeMs float64, errors int64, warnings int64, tmpTables int64, tmpDiskTables int64, noIndexUsed int64, noGoodIndexUsed int64, database string) *MysqlSQLFingerprint {
	this := MysqlSQLFingerprint{}
	this.QueryId = queryId
	this.Fingerprint = fingerprint
	this.QuerySummary = querySummary
	this.Calls = calls
	this.TotalTimeMs = totalTimeMs
	this.MeanTimeMs = meanTimeMs
	this.MaxTimeMs = maxTimeMs
	this.Rows = rows
	this.RowsSent = rowsSent
	this.RowsAffected = rowsAffected
	this.RowsExamined = rowsExamined
	this.LockTimeMs = lockTimeMs
	this.Errors = errors
	this.Warnings = warnings
	this.TmpTables = tmpTables
	this.TmpDiskTables = tmpDiskTables
	this.NoIndexUsed = noIndexUsed
	this.NoGoodIndexUsed = noGoodIndexUsed
	this.Database = database
	return &this
}

// NewMysqlSQLFingerprintWithDefaults instantiates a new MysqlSQLFingerprint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMysqlSQLFingerprintWithDefaults() *MysqlSQLFingerprint {
	this := MysqlSQLFingerprint{}
	return &this
}

// GetQueryId returns the QueryId field value.
func (o *MysqlSQLFingerprint) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value.
func (o *MysqlSQLFingerprint) SetQueryId(v string) {
	o.QueryId = v
}

// GetFingerprint returns the Fingerprint field value.
func (o *MysqlSQLFingerprint) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value.
func (o *MysqlSQLFingerprint) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetQuerySummary returns the QuerySummary field value.
func (o *MysqlSQLFingerprint) GetQuerySummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QuerySummary
}

// GetQuerySummaryOk returns a tuple with the QuerySummary field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetQuerySummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuerySummary, true
}

// SetQuerySummary sets field value.
func (o *MysqlSQLFingerprint) SetQuerySummary(v string) {
	o.QuerySummary = v
}

// GetCalls returns the Calls field value.
func (o *MysqlSQLFingerprint) GetCalls() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Calls
}

// GetCallsOk returns a tuple with the Calls field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetCallsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Calls, true
}

// SetCalls sets field value.
func (o *MysqlSQLFingerprint) SetCalls(v int64) {
	o.Calls = v
}

// GetTotalTimeMs returns the TotalTimeMs field value.
func (o *MysqlSQLFingerprint) GetTotalTimeMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.TotalTimeMs
}

// GetTotalTimeMsOk returns a tuple with the TotalTimeMs field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetTotalTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTimeMs, true
}

// SetTotalTimeMs sets field value.
func (o *MysqlSQLFingerprint) SetTotalTimeMs(v float64) {
	o.TotalTimeMs = v
}

// GetMeanTimeMs returns the MeanTimeMs field value.
func (o *MysqlSQLFingerprint) GetMeanTimeMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MeanTimeMs
}

// GetMeanTimeMsOk returns a tuple with the MeanTimeMs field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetMeanTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeanTimeMs, true
}

// SetMeanTimeMs sets field value.
func (o *MysqlSQLFingerprint) SetMeanTimeMs(v float64) {
	o.MeanTimeMs = v
}

// GetMaxTimeMs returns the MaxTimeMs field value.
func (o *MysqlSQLFingerprint) GetMaxTimeMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MaxTimeMs
}

// GetMaxTimeMsOk returns a tuple with the MaxTimeMs field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetMaxTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxTimeMs, true
}

// SetMaxTimeMs sets field value.
func (o *MysqlSQLFingerprint) SetMaxTimeMs(v float64) {
	o.MaxTimeMs = v
}

// GetRows returns the Rows field value.
func (o *MysqlSQLFingerprint) GetRows() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetRowsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rows, true
}

// SetRows sets field value.
func (o *MysqlSQLFingerprint) SetRows(v int64) {
	o.Rows = v
}

// GetRowsSent returns the RowsSent field value.
func (o *MysqlSQLFingerprint) GetRowsSent() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RowsSent
}

// GetRowsSentOk returns a tuple with the RowsSent field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetRowsSentOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowsSent, true
}

// SetRowsSent sets field value.
func (o *MysqlSQLFingerprint) SetRowsSent(v int64) {
	o.RowsSent = v
}

// GetRowsAffected returns the RowsAffected field value.
func (o *MysqlSQLFingerprint) GetRowsAffected() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RowsAffected
}

// GetRowsAffectedOk returns a tuple with the RowsAffected field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetRowsAffectedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowsAffected, true
}

// SetRowsAffected sets field value.
func (o *MysqlSQLFingerprint) SetRowsAffected(v int64) {
	o.RowsAffected = v
}

// GetRowsExamined returns the RowsExamined field value.
func (o *MysqlSQLFingerprint) GetRowsExamined() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RowsExamined
}

// GetRowsExaminedOk returns a tuple with the RowsExamined field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetRowsExaminedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowsExamined, true
}

// SetRowsExamined sets field value.
func (o *MysqlSQLFingerprint) SetRowsExamined(v int64) {
	o.RowsExamined = v
}

// GetLockTimeMs returns the LockTimeMs field value.
func (o *MysqlSQLFingerprint) GetLockTimeMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.LockTimeMs
}

// GetLockTimeMsOk returns a tuple with the LockTimeMs field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetLockTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockTimeMs, true
}

// SetLockTimeMs sets field value.
func (o *MysqlSQLFingerprint) SetLockTimeMs(v float64) {
	o.LockTimeMs = v
}

// GetErrors returns the Errors field value.
func (o *MysqlSQLFingerprint) GetErrors() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetErrorsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value.
func (o *MysqlSQLFingerprint) SetErrors(v int64) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value.
func (o *MysqlSQLFingerprint) GetWarnings() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetWarningsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// SetWarnings sets field value.
func (o *MysqlSQLFingerprint) SetWarnings(v int64) {
	o.Warnings = v
}

// GetTmpTables returns the TmpTables field value.
func (o *MysqlSQLFingerprint) GetTmpTables() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TmpTables
}

// GetTmpTablesOk returns a tuple with the TmpTables field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetTmpTablesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TmpTables, true
}

// SetTmpTables sets field value.
func (o *MysqlSQLFingerprint) SetTmpTables(v int64) {
	o.TmpTables = v
}

// GetTmpDiskTables returns the TmpDiskTables field value.
func (o *MysqlSQLFingerprint) GetTmpDiskTables() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TmpDiskTables
}

// GetTmpDiskTablesOk returns a tuple with the TmpDiskTables field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetTmpDiskTablesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TmpDiskTables, true
}

// SetTmpDiskTables sets field value.
func (o *MysqlSQLFingerprint) SetTmpDiskTables(v int64) {
	o.TmpDiskTables = v
}

// GetNoIndexUsed returns the NoIndexUsed field value.
func (o *MysqlSQLFingerprint) GetNoIndexUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.NoIndexUsed
}

// GetNoIndexUsedOk returns a tuple with the NoIndexUsed field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetNoIndexUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoIndexUsed, true
}

// SetNoIndexUsed sets field value.
func (o *MysqlSQLFingerprint) SetNoIndexUsed(v int64) {
	o.NoIndexUsed = v
}

// GetNoGoodIndexUsed returns the NoGoodIndexUsed field value.
func (o *MysqlSQLFingerprint) GetNoGoodIndexUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.NoGoodIndexUsed
}

// GetNoGoodIndexUsedOk returns a tuple with the NoGoodIndexUsed field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetNoGoodIndexUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoGoodIndexUsed, true
}

// SetNoGoodIndexUsed sets field value.
func (o *MysqlSQLFingerprint) SetNoGoodIndexUsed(v int64) {
	o.NoGoodIndexUsed = v
}

// GetDatabase returns the Database field value.
func (o *MysqlSQLFingerprint) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *MysqlSQLFingerprint) SetDatabase(v string) {
	o.Database = v
}

// GetFirstSeen returns the FirstSeen field value if set, zero value otherwise.
func (o *MysqlSQLFingerprint) GetFirstSeen() string {
	if o == nil || o.FirstSeen == nil {
		var ret string
		return ret
	}
	return *o.FirstSeen
}

// GetFirstSeenOk returns a tuple with the FirstSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetFirstSeenOk() (*string, bool) {
	if o == nil || o.FirstSeen == nil {
		return nil, false
	}
	return o.FirstSeen, true
}

// HasFirstSeen returns a boolean if a field has been set.
func (o *MysqlSQLFingerprint) HasFirstSeen() bool {
	return o != nil && o.FirstSeen != nil
}

// SetFirstSeen gets a reference to the given string and assigns it to the FirstSeen field.
func (o *MysqlSQLFingerprint) SetFirstSeen(v string) {
	o.FirstSeen = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *MysqlSQLFingerprint) GetLastSeen() string {
	if o == nil || o.LastSeen == nil {
		var ret string
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlSQLFingerprint) GetLastSeenOk() (*string, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *MysqlSQLFingerprint) HasLastSeen() bool {
	return o != nil && o.LastSeen != nil
}

// SetLastSeen gets a reference to the given string and assigns it to the LastSeen field.
func (o *MysqlSQLFingerprint) SetLastSeen(v string) {
	o.LastSeen = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MysqlSQLFingerprint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["queryID"] = o.QueryId
	toSerialize["fingerprint"] = o.Fingerprint
	toSerialize["querySummary"] = o.QuerySummary
	toSerialize["calls"] = o.Calls
	toSerialize["totalTimeMs"] = o.TotalTimeMs
	toSerialize["meanTimeMs"] = o.MeanTimeMs
	toSerialize["maxTimeMs"] = o.MaxTimeMs
	toSerialize["rows"] = o.Rows
	toSerialize["rowsSent"] = o.RowsSent
	toSerialize["rowsAffected"] = o.RowsAffected
	toSerialize["rowsExamined"] = o.RowsExamined
	toSerialize["lockTimeMs"] = o.LockTimeMs
	toSerialize["errors"] = o.Errors
	toSerialize["warnings"] = o.Warnings
	toSerialize["tmpTables"] = o.TmpTables
	toSerialize["tmpDiskTables"] = o.TmpDiskTables
	toSerialize["noIndexUsed"] = o.NoIndexUsed
	toSerialize["noGoodIndexUsed"] = o.NoGoodIndexUsed
	toSerialize["database"] = o.Database
	if o.FirstSeen != nil {
		toSerialize["firstSeen"] = o.FirstSeen
	}
	if o.LastSeen != nil {
		toSerialize["lastSeen"] = o.LastSeen
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MysqlSQLFingerprint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryId         *string  `json:"queryID"`
		Fingerprint     *string  `json:"fingerprint"`
		QuerySummary    *string  `json:"querySummary"`
		Calls           *int64   `json:"calls"`
		TotalTimeMs     *float64 `json:"totalTimeMs"`
		MeanTimeMs      *float64 `json:"meanTimeMs"`
		MaxTimeMs       *float64 `json:"maxTimeMs"`
		Rows            *int64   `json:"rows"`
		RowsSent        *int64   `json:"rowsSent"`
		RowsAffected    *int64   `json:"rowsAffected"`
		RowsExamined    *int64   `json:"rowsExamined"`
		LockTimeMs      *float64 `json:"lockTimeMs"`
		Errors          *int64   `json:"errors"`
		Warnings        *int64   `json:"warnings"`
		TmpTables       *int64   `json:"tmpTables"`
		TmpDiskTables   *int64   `json:"tmpDiskTables"`
		NoIndexUsed     *int64   `json:"noIndexUsed"`
		NoGoodIndexUsed *int64   `json:"noGoodIndexUsed"`
		Database        *string  `json:"database"`
		FirstSeen       *string  `json:"firstSeen,omitempty"`
		LastSeen        *string  `json:"lastSeen,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.QueryId == nil {
		return fmt.Errorf("required field queryID missing")
	}
	if all.Fingerprint == nil {
		return fmt.Errorf("required field fingerprint missing")
	}
	if all.QuerySummary == nil {
		return fmt.Errorf("required field querySummary missing")
	}
	if all.Calls == nil {
		return fmt.Errorf("required field calls missing")
	}
	if all.TotalTimeMs == nil {
		return fmt.Errorf("required field totalTimeMs missing")
	}
	if all.MeanTimeMs == nil {
		return fmt.Errorf("required field meanTimeMs missing")
	}
	if all.MaxTimeMs == nil {
		return fmt.Errorf("required field maxTimeMs missing")
	}
	if all.Rows == nil {
		return fmt.Errorf("required field rows missing")
	}
	if all.RowsSent == nil {
		return fmt.Errorf("required field rowsSent missing")
	}
	if all.RowsAffected == nil {
		return fmt.Errorf("required field rowsAffected missing")
	}
	if all.RowsExamined == nil {
		return fmt.Errorf("required field rowsExamined missing")
	}
	if all.LockTimeMs == nil {
		return fmt.Errorf("required field lockTimeMs missing")
	}
	if all.Errors == nil {
		return fmt.Errorf("required field errors missing")
	}
	if all.Warnings == nil {
		return fmt.Errorf("required field warnings missing")
	}
	if all.TmpTables == nil {
		return fmt.Errorf("required field tmpTables missing")
	}
	if all.TmpDiskTables == nil {
		return fmt.Errorf("required field tmpDiskTables missing")
	}
	if all.NoIndexUsed == nil {
		return fmt.Errorf("required field noIndexUsed missing")
	}
	if all.NoGoodIndexUsed == nil {
		return fmt.Errorf("required field noGoodIndexUsed missing")
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"queryID", "fingerprint", "querySummary", "calls", "totalTimeMs", "meanTimeMs", "maxTimeMs", "rows", "rowsSent", "rowsAffected", "rowsExamined", "lockTimeMs", "errors", "warnings", "tmpTables", "tmpDiskTables", "noIndexUsed", "noGoodIndexUsed", "database", "firstSeen", "lastSeen"})
	} else {
		return err
	}
	o.QueryId = *all.QueryId
	o.Fingerprint = *all.Fingerprint
	o.QuerySummary = *all.QuerySummary
	o.Calls = *all.Calls
	o.TotalTimeMs = *all.TotalTimeMs
	o.MeanTimeMs = *all.MeanTimeMs
	o.MaxTimeMs = *all.MaxTimeMs
	o.Rows = *all.Rows
	o.RowsSent = *all.RowsSent
	o.RowsAffected = *all.RowsAffected
	o.RowsExamined = *all.RowsExamined
	o.LockTimeMs = *all.LockTimeMs
	o.Errors = *all.Errors
	o.Warnings = *all.Warnings
	o.TmpTables = *all.TmpTables
	o.TmpDiskTables = *all.TmpDiskTables
	o.NoIndexUsed = *all.NoIndexUsed
	o.NoGoodIndexUsed = *all.NoGoodIndexUsed
	o.Database = *all.Database
	o.FirstSeen = all.FirstSeen
	o.LastSeen = all.LastSeen

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
