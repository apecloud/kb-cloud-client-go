// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// MssqlSpaceAnalysis Current database files and on-disk table allocations. Free pages are reusable inside the database and do not imply immediately reclaimable filesystem bytes. Historical disk usage is provided by the cluster storage usage history API.
type MssqlSpaceAnalysis struct {
	CapturedAt      string               `json:"capturedAt"`
	Database        string               `json:"database"`
	Databases       []MssqlDatabaseSpace `json:"databases"`
	Files           []MssqlFileSpace     `json:"files"`
	Tables          []MssqlTableSpace    `json:"tables"`
	TablesTruncated bool                 `json:"tablesTruncated"`
	DataBytes       int64                `json:"dataBytes"`
	DataUsedBytes   int64                `json:"dataUsedBytes"`
	LogBytes        int64                `json:"logBytes"`
	LogUsedBytes    int64                `json:"logUsedBytes"`
	LogReuseWait    string               `json:"logReuseWait"`
	Recommendations []string             `json:"recommendations"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMssqlSpaceAnalysis instantiates a new MssqlSpaceAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMssqlSpaceAnalysis(capturedAt string, database string, databases []MssqlDatabaseSpace, files []MssqlFileSpace, tables []MssqlTableSpace, tablesTruncated bool, dataBytes int64, dataUsedBytes int64, logBytes int64, logUsedBytes int64, logReuseWait string, recommendations []string) *MssqlSpaceAnalysis {
	this := MssqlSpaceAnalysis{}
	this.CapturedAt = capturedAt
	this.Database = database
	this.Databases = databases
	this.Files = files
	this.Tables = tables
	this.TablesTruncated = tablesTruncated
	this.DataBytes = dataBytes
	this.DataUsedBytes = dataUsedBytes
	this.LogBytes = logBytes
	this.LogUsedBytes = logUsedBytes
	this.LogReuseWait = logReuseWait
	this.Recommendations = recommendations
	return &this
}

// NewMssqlSpaceAnalysisWithDefaults instantiates a new MssqlSpaceAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMssqlSpaceAnalysisWithDefaults() *MssqlSpaceAnalysis {
	this := MssqlSpaceAnalysis{}
	return &this
}

// GetCapturedAt returns the CapturedAt field value.
func (o *MssqlSpaceAnalysis) GetCapturedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CapturedAt
}

// GetCapturedAtOk returns a tuple with the CapturedAt field value
// and a boolean to check if the value has been set.
func (o *MssqlSpaceAnalysis) GetCapturedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapturedAt, true
}

// SetCapturedAt sets field value.
func (o *MssqlSpaceAnalysis) SetCapturedAt(v string) {
	o.CapturedAt = v
}

// GetDatabase returns the Database field value.
func (o *MssqlSpaceAnalysis) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *MssqlSpaceAnalysis) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *MssqlSpaceAnalysis) SetDatabase(v string) {
	o.Database = v
}

// GetDatabases returns the Databases field value.
func (o *MssqlSpaceAnalysis) GetDatabases() []MssqlDatabaseSpace {
	if o == nil {
		var ret []MssqlDatabaseSpace
		return ret
	}
	return o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value
// and a boolean to check if the value has been set.
func (o *MssqlSpaceAnalysis) GetDatabasesOk() (*[]MssqlDatabaseSpace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Databases, true
}

// SetDatabases sets field value.
func (o *MssqlSpaceAnalysis) SetDatabases(v []MssqlDatabaseSpace) {
	o.Databases = v
}

// GetFiles returns the Files field value.
func (o *MssqlSpaceAnalysis) GetFiles() []MssqlFileSpace {
	if o == nil {
		var ret []MssqlFileSpace
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *MssqlSpaceAnalysis) GetFilesOk() (*[]MssqlFileSpace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Files, true
}

// SetFiles sets field value.
func (o *MssqlSpaceAnalysis) SetFiles(v []MssqlFileSpace) {
	o.Files = v
}

// GetTables returns the Tables field value.
func (o *MssqlSpaceAnalysis) GetTables() []MssqlTableSpace {
	if o == nil {
		var ret []MssqlTableSpace
		return ret
	}
	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value
// and a boolean to check if the value has been set.
func (o *MssqlSpaceAnalysis) GetTablesOk() (*[]MssqlTableSpace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tables, true
}

// SetTables sets field value.
func (o *MssqlSpaceAnalysis) SetTables(v []MssqlTableSpace) {
	o.Tables = v
}

// GetTablesTruncated returns the TablesTruncated field value.
func (o *MssqlSpaceAnalysis) GetTablesTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.TablesTruncated
}

// GetTablesTruncatedOk returns a tuple with the TablesTruncated field value
// and a boolean to check if the value has been set.
func (o *MssqlSpaceAnalysis) GetTablesTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TablesTruncated, true
}

// SetTablesTruncated sets field value.
func (o *MssqlSpaceAnalysis) SetTablesTruncated(v bool) {
	o.TablesTruncated = v
}

// GetDataBytes returns the DataBytes field value.
func (o *MssqlSpaceAnalysis) GetDataBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DataBytes
}

// GetDataBytesOk returns a tuple with the DataBytes field value
// and a boolean to check if the value has been set.
func (o *MssqlSpaceAnalysis) GetDataBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataBytes, true
}

// SetDataBytes sets field value.
func (o *MssqlSpaceAnalysis) SetDataBytes(v int64) {
	o.DataBytes = v
}

// GetDataUsedBytes returns the DataUsedBytes field value.
func (o *MssqlSpaceAnalysis) GetDataUsedBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DataUsedBytes
}

// GetDataUsedBytesOk returns a tuple with the DataUsedBytes field value
// and a boolean to check if the value has been set.
func (o *MssqlSpaceAnalysis) GetDataUsedBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataUsedBytes, true
}

// SetDataUsedBytes sets field value.
func (o *MssqlSpaceAnalysis) SetDataUsedBytes(v int64) {
	o.DataUsedBytes = v
}

// GetLogBytes returns the LogBytes field value.
func (o *MssqlSpaceAnalysis) GetLogBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LogBytes
}

// GetLogBytesOk returns a tuple with the LogBytes field value
// and a boolean to check if the value has been set.
func (o *MssqlSpaceAnalysis) GetLogBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogBytes, true
}

// SetLogBytes sets field value.
func (o *MssqlSpaceAnalysis) SetLogBytes(v int64) {
	o.LogBytes = v
}

// GetLogUsedBytes returns the LogUsedBytes field value.
func (o *MssqlSpaceAnalysis) GetLogUsedBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LogUsedBytes
}

// GetLogUsedBytesOk returns a tuple with the LogUsedBytes field value
// and a boolean to check if the value has been set.
func (o *MssqlSpaceAnalysis) GetLogUsedBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogUsedBytes, true
}

// SetLogUsedBytes sets field value.
func (o *MssqlSpaceAnalysis) SetLogUsedBytes(v int64) {
	o.LogUsedBytes = v
}

// GetLogReuseWait returns the LogReuseWait field value.
func (o *MssqlSpaceAnalysis) GetLogReuseWait() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LogReuseWait
}

// GetLogReuseWaitOk returns a tuple with the LogReuseWait field value
// and a boolean to check if the value has been set.
func (o *MssqlSpaceAnalysis) GetLogReuseWaitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogReuseWait, true
}

// SetLogReuseWait sets field value.
func (o *MssqlSpaceAnalysis) SetLogReuseWait(v string) {
	o.LogReuseWait = v
}

// GetRecommendations returns the Recommendations field value.
func (o *MssqlSpaceAnalysis) GetRecommendations() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value
// and a boolean to check if the value has been set.
func (o *MssqlSpaceAnalysis) GetRecommendationsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recommendations, true
}

// SetRecommendations sets field value.
func (o *MssqlSpaceAnalysis) SetRecommendations(v []string) {
	o.Recommendations = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MssqlSpaceAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["capturedAt"] = o.CapturedAt
	toSerialize["database"] = o.Database
	toSerialize["databases"] = o.Databases
	toSerialize["files"] = o.Files
	toSerialize["tables"] = o.Tables
	toSerialize["tablesTruncated"] = o.TablesTruncated
	toSerialize["dataBytes"] = o.DataBytes
	toSerialize["dataUsedBytes"] = o.DataUsedBytes
	toSerialize["logBytes"] = o.LogBytes
	toSerialize["logUsedBytes"] = o.LogUsedBytes
	toSerialize["logReuseWait"] = o.LogReuseWait
	toSerialize["recommendations"] = o.Recommendations

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MssqlSpaceAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CapturedAt      *string               `json:"capturedAt"`
		Database        *string               `json:"database"`
		Databases       *[]MssqlDatabaseSpace `json:"databases"`
		Files           *[]MssqlFileSpace     `json:"files"`
		Tables          *[]MssqlTableSpace    `json:"tables"`
		TablesTruncated *bool                 `json:"tablesTruncated"`
		DataBytes       *int64                `json:"dataBytes"`
		DataUsedBytes   *int64                `json:"dataUsedBytes"`
		LogBytes        *int64                `json:"logBytes"`
		LogUsedBytes    *int64                `json:"logUsedBytes"`
		LogReuseWait    *string               `json:"logReuseWait"`
		Recommendations *[]string             `json:"recommendations"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CapturedAt == nil {
		return fmt.Errorf("required field capturedAt missing")
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	if all.Databases == nil {
		return fmt.Errorf("required field databases missing")
	}
	if all.Files == nil {
		return fmt.Errorf("required field files missing")
	}
	if all.Tables == nil {
		return fmt.Errorf("required field tables missing")
	}
	if all.TablesTruncated == nil {
		return fmt.Errorf("required field tablesTruncated missing")
	}
	if all.DataBytes == nil {
		return fmt.Errorf("required field dataBytes missing")
	}
	if all.DataUsedBytes == nil {
		return fmt.Errorf("required field dataUsedBytes missing")
	}
	if all.LogBytes == nil {
		return fmt.Errorf("required field logBytes missing")
	}
	if all.LogUsedBytes == nil {
		return fmt.Errorf("required field logUsedBytes missing")
	}
	if all.LogReuseWait == nil {
		return fmt.Errorf("required field logReuseWait missing")
	}
	if all.Recommendations == nil {
		return fmt.Errorf("required field recommendations missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"capturedAt", "database", "databases", "files", "tables", "tablesTruncated", "dataBytes", "dataUsedBytes", "logBytes", "logUsedBytes", "logReuseWait", "recommendations"})
	} else {
		return err
	}
	o.CapturedAt = *all.CapturedAt
	o.Database = *all.Database
	o.Databases = *all.Databases
	o.Files = *all.Files
	o.Tables = *all.Tables
	o.TablesTruncated = *all.TablesTruncated
	o.DataBytes = *all.DataBytes
	o.DataUsedBytes = *all.DataUsedBytes
	o.LogBytes = *all.LogBytes
	o.LogUsedBytes = *all.LogUsedBytes
	o.LogReuseWait = *all.LogReuseWait
	o.Recommendations = *all.Recommendations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
