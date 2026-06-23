// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type PostgresqlTableSpace struct {
	Database         *string `json:"database,omitempty"`
	Schema           *string `json:"schema,omitempty"`
	Name             *string `json:"name,omitempty"`
	RelationOid      *int64  `json:"relationOid,omitempty"`
	TotalBytes       *int64  `json:"totalBytes,omitempty"`
	TableBytes       *int64  `json:"tableBytes,omitempty"`
	IndexBytes       *int64  `json:"indexBytes,omitempty"`
	ToastBytes       *int64  `json:"toastBytes,omitempty"`
	ToastRelationOid *int64  `json:"toastRelationOid,omitempty"`
	LiveTuples       *int64  `json:"liveTuples,omitempty"`
	DeadTuples       *int64  `json:"deadTuples,omitempty"`
	SeqScan          *int64  `json:"seqScan,omitempty"`
	IndexScan        *int64  `json:"indexScan,omitempty"`
	LastVacuum       *string `json:"lastVacuum,omitempty"`
	LastAutoVacuum   *string `json:"lastAutoVacuum,omitempty"`
	LastAnalyze      *string `json:"lastAnalyze,omitempty"`
	LastAutoAnalyze  *string `json:"lastAutoAnalyze,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlTableSpace instantiates a new PostgresqlTableSpace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlTableSpace() *PostgresqlTableSpace {
	this := PostgresqlTableSpace{}
	return &this
}

// NewPostgresqlTableSpaceWithDefaults instantiates a new PostgresqlTableSpace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlTableSpaceWithDefaults() *PostgresqlTableSpace {
	this := PostgresqlTableSpace{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *PostgresqlTableSpace) SetDatabase(v string) {
	o.Database = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *PostgresqlTableSpace) SetSchema(v string) {
	o.Schema = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostgresqlTableSpace) SetName(v string) {
	o.Name = &v
}

// GetRelationOid returns the RelationOid field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetRelationOid() int64 {
	if o == nil || o.RelationOid == nil {
		var ret int64
		return ret
	}
	return *o.RelationOid
}

// GetRelationOidOk returns a tuple with the RelationOid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetRelationOidOk() (*int64, bool) {
	if o == nil || o.RelationOid == nil {
		return nil, false
	}
	return o.RelationOid, true
}

// HasRelationOid returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasRelationOid() bool {
	return o != nil && o.RelationOid != nil
}

// SetRelationOid gets a reference to the given int64 and assigns it to the RelationOid field.
func (o *PostgresqlTableSpace) SetRelationOid(v int64) {
	o.RelationOid = &v
}

// GetTotalBytes returns the TotalBytes field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetTotalBytes() int64 {
	if o == nil || o.TotalBytes == nil {
		var ret int64
		return ret
	}
	return *o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetTotalBytesOk() (*int64, bool) {
	if o == nil || o.TotalBytes == nil {
		return nil, false
	}
	return o.TotalBytes, true
}

// HasTotalBytes returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasTotalBytes() bool {
	return o != nil && o.TotalBytes != nil
}

// SetTotalBytes gets a reference to the given int64 and assigns it to the TotalBytes field.
func (o *PostgresqlTableSpace) SetTotalBytes(v int64) {
	o.TotalBytes = &v
}

// GetTableBytes returns the TableBytes field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetTableBytes() int64 {
	if o == nil || o.TableBytes == nil {
		var ret int64
		return ret
	}
	return *o.TableBytes
}

// GetTableBytesOk returns a tuple with the TableBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetTableBytesOk() (*int64, bool) {
	if o == nil || o.TableBytes == nil {
		return nil, false
	}
	return o.TableBytes, true
}

// HasTableBytes returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasTableBytes() bool {
	return o != nil && o.TableBytes != nil
}

// SetTableBytes gets a reference to the given int64 and assigns it to the TableBytes field.
func (o *PostgresqlTableSpace) SetTableBytes(v int64) {
	o.TableBytes = &v
}

// GetIndexBytes returns the IndexBytes field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetIndexBytes() int64 {
	if o == nil || o.IndexBytes == nil {
		var ret int64
		return ret
	}
	return *o.IndexBytes
}

// GetIndexBytesOk returns a tuple with the IndexBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetIndexBytesOk() (*int64, bool) {
	if o == nil || o.IndexBytes == nil {
		return nil, false
	}
	return o.IndexBytes, true
}

// HasIndexBytes returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasIndexBytes() bool {
	return o != nil && o.IndexBytes != nil
}

// SetIndexBytes gets a reference to the given int64 and assigns it to the IndexBytes field.
func (o *PostgresqlTableSpace) SetIndexBytes(v int64) {
	o.IndexBytes = &v
}

// GetToastBytes returns the ToastBytes field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetToastBytes() int64 {
	if o == nil || o.ToastBytes == nil {
		var ret int64
		return ret
	}
	return *o.ToastBytes
}

// GetToastBytesOk returns a tuple with the ToastBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetToastBytesOk() (*int64, bool) {
	if o == nil || o.ToastBytes == nil {
		return nil, false
	}
	return o.ToastBytes, true
}

// HasToastBytes returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasToastBytes() bool {
	return o != nil && o.ToastBytes != nil
}

// SetToastBytes gets a reference to the given int64 and assigns it to the ToastBytes field.
func (o *PostgresqlTableSpace) SetToastBytes(v int64) {
	o.ToastBytes = &v
}

// GetToastRelationOid returns the ToastRelationOid field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetToastRelationOid() int64 {
	if o == nil || o.ToastRelationOid == nil {
		var ret int64
		return ret
	}
	return *o.ToastRelationOid
}

// GetToastRelationOidOk returns a tuple with the ToastRelationOid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetToastRelationOidOk() (*int64, bool) {
	if o == nil || o.ToastRelationOid == nil {
		return nil, false
	}
	return o.ToastRelationOid, true
}

// HasToastRelationOid returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasToastRelationOid() bool {
	return o != nil && o.ToastRelationOid != nil
}

// SetToastRelationOid gets a reference to the given int64 and assigns it to the ToastRelationOid field.
func (o *PostgresqlTableSpace) SetToastRelationOid(v int64) {
	o.ToastRelationOid = &v
}

// GetLiveTuples returns the LiveTuples field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetLiveTuples() int64 {
	if o == nil || o.LiveTuples == nil {
		var ret int64
		return ret
	}
	return *o.LiveTuples
}

// GetLiveTuplesOk returns a tuple with the LiveTuples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetLiveTuplesOk() (*int64, bool) {
	if o == nil || o.LiveTuples == nil {
		return nil, false
	}
	return o.LiveTuples, true
}

// HasLiveTuples returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasLiveTuples() bool {
	return o != nil && o.LiveTuples != nil
}

// SetLiveTuples gets a reference to the given int64 and assigns it to the LiveTuples field.
func (o *PostgresqlTableSpace) SetLiveTuples(v int64) {
	o.LiveTuples = &v
}

// GetDeadTuples returns the DeadTuples field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetDeadTuples() int64 {
	if o == nil || o.DeadTuples == nil {
		var ret int64
		return ret
	}
	return *o.DeadTuples
}

// GetDeadTuplesOk returns a tuple with the DeadTuples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetDeadTuplesOk() (*int64, bool) {
	if o == nil || o.DeadTuples == nil {
		return nil, false
	}
	return o.DeadTuples, true
}

// HasDeadTuples returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasDeadTuples() bool {
	return o != nil && o.DeadTuples != nil
}

// SetDeadTuples gets a reference to the given int64 and assigns it to the DeadTuples field.
func (o *PostgresqlTableSpace) SetDeadTuples(v int64) {
	o.DeadTuples = &v
}

// GetSeqScan returns the SeqScan field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetSeqScan() int64 {
	if o == nil || o.SeqScan == nil {
		var ret int64
		return ret
	}
	return *o.SeqScan
}

// GetSeqScanOk returns a tuple with the SeqScan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetSeqScanOk() (*int64, bool) {
	if o == nil || o.SeqScan == nil {
		return nil, false
	}
	return o.SeqScan, true
}

// HasSeqScan returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasSeqScan() bool {
	return o != nil && o.SeqScan != nil
}

// SetSeqScan gets a reference to the given int64 and assigns it to the SeqScan field.
func (o *PostgresqlTableSpace) SetSeqScan(v int64) {
	o.SeqScan = &v
}

// GetIndexScan returns the IndexScan field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetIndexScan() int64 {
	if o == nil || o.IndexScan == nil {
		var ret int64
		return ret
	}
	return *o.IndexScan
}

// GetIndexScanOk returns a tuple with the IndexScan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetIndexScanOk() (*int64, bool) {
	if o == nil || o.IndexScan == nil {
		return nil, false
	}
	return o.IndexScan, true
}

// HasIndexScan returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasIndexScan() bool {
	return o != nil && o.IndexScan != nil
}

// SetIndexScan gets a reference to the given int64 and assigns it to the IndexScan field.
func (o *PostgresqlTableSpace) SetIndexScan(v int64) {
	o.IndexScan = &v
}

// GetLastVacuum returns the LastVacuum field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetLastVacuum() string {
	if o == nil || o.LastVacuum == nil {
		var ret string
		return ret
	}
	return *o.LastVacuum
}

// GetLastVacuumOk returns a tuple with the LastVacuum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetLastVacuumOk() (*string, bool) {
	if o == nil || o.LastVacuum == nil {
		return nil, false
	}
	return o.LastVacuum, true
}

// HasLastVacuum returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasLastVacuum() bool {
	return o != nil && o.LastVacuum != nil
}

// SetLastVacuum gets a reference to the given string and assigns it to the LastVacuum field.
func (o *PostgresqlTableSpace) SetLastVacuum(v string) {
	o.LastVacuum = &v
}

// GetLastAutoVacuum returns the LastAutoVacuum field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetLastAutoVacuum() string {
	if o == nil || o.LastAutoVacuum == nil {
		var ret string
		return ret
	}
	return *o.LastAutoVacuum
}

// GetLastAutoVacuumOk returns a tuple with the LastAutoVacuum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetLastAutoVacuumOk() (*string, bool) {
	if o == nil || o.LastAutoVacuum == nil {
		return nil, false
	}
	return o.LastAutoVacuum, true
}

// HasLastAutoVacuum returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasLastAutoVacuum() bool {
	return o != nil && o.LastAutoVacuum != nil
}

// SetLastAutoVacuum gets a reference to the given string and assigns it to the LastAutoVacuum field.
func (o *PostgresqlTableSpace) SetLastAutoVacuum(v string) {
	o.LastAutoVacuum = &v
}

// GetLastAnalyze returns the LastAnalyze field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetLastAnalyze() string {
	if o == nil || o.LastAnalyze == nil {
		var ret string
		return ret
	}
	return *o.LastAnalyze
}

// GetLastAnalyzeOk returns a tuple with the LastAnalyze field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetLastAnalyzeOk() (*string, bool) {
	if o == nil || o.LastAnalyze == nil {
		return nil, false
	}
	return o.LastAnalyze, true
}

// HasLastAnalyze returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasLastAnalyze() bool {
	return o != nil && o.LastAnalyze != nil
}

// SetLastAnalyze gets a reference to the given string and assigns it to the LastAnalyze field.
func (o *PostgresqlTableSpace) SetLastAnalyze(v string) {
	o.LastAnalyze = &v
}

// GetLastAutoAnalyze returns the LastAutoAnalyze field value if set, zero value otherwise.
func (o *PostgresqlTableSpace) GetLastAutoAnalyze() string {
	if o == nil || o.LastAutoAnalyze == nil {
		var ret string
		return ret
	}
	return *o.LastAutoAnalyze
}

// GetLastAutoAnalyzeOk returns a tuple with the LastAutoAnalyze field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlTableSpace) GetLastAutoAnalyzeOk() (*string, bool) {
	if o == nil || o.LastAutoAnalyze == nil {
		return nil, false
	}
	return o.LastAutoAnalyze, true
}

// HasLastAutoAnalyze returns a boolean if a field has been set.
func (o *PostgresqlTableSpace) HasLastAutoAnalyze() bool {
	return o != nil && o.LastAutoAnalyze != nil
}

// SetLastAutoAnalyze gets a reference to the given string and assigns it to the LastAutoAnalyze field.
func (o *PostgresqlTableSpace) SetLastAutoAnalyze(v string) {
	o.LastAutoAnalyze = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlTableSpace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RelationOid != nil {
		toSerialize["relationOid"] = o.RelationOid
	}
	if o.TotalBytes != nil {
		toSerialize["totalBytes"] = o.TotalBytes
	}
	if o.TableBytes != nil {
		toSerialize["tableBytes"] = o.TableBytes
	}
	if o.IndexBytes != nil {
		toSerialize["indexBytes"] = o.IndexBytes
	}
	if o.ToastBytes != nil {
		toSerialize["toastBytes"] = o.ToastBytes
	}
	if o.ToastRelationOid != nil {
		toSerialize["toastRelationOid"] = o.ToastRelationOid
	}
	if o.LiveTuples != nil {
		toSerialize["liveTuples"] = o.LiveTuples
	}
	if o.DeadTuples != nil {
		toSerialize["deadTuples"] = o.DeadTuples
	}
	if o.SeqScan != nil {
		toSerialize["seqScan"] = o.SeqScan
	}
	if o.IndexScan != nil {
		toSerialize["indexScan"] = o.IndexScan
	}
	if o.LastVacuum != nil {
		toSerialize["lastVacuum"] = o.LastVacuum
	}
	if o.LastAutoVacuum != nil {
		toSerialize["lastAutoVacuum"] = o.LastAutoVacuum
	}
	if o.LastAnalyze != nil {
		toSerialize["lastAnalyze"] = o.LastAnalyze
	}
	if o.LastAutoAnalyze != nil {
		toSerialize["lastAutoAnalyze"] = o.LastAutoAnalyze
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlTableSpace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database         *string `json:"database,omitempty"`
		Schema           *string `json:"schema,omitempty"`
		Name             *string `json:"name,omitempty"`
		RelationOid      *int64  `json:"relationOid,omitempty"`
		TotalBytes       *int64  `json:"totalBytes,omitempty"`
		TableBytes       *int64  `json:"tableBytes,omitempty"`
		IndexBytes       *int64  `json:"indexBytes,omitempty"`
		ToastBytes       *int64  `json:"toastBytes,omitempty"`
		ToastRelationOid *int64  `json:"toastRelationOid,omitempty"`
		LiveTuples       *int64  `json:"liveTuples,omitempty"`
		DeadTuples       *int64  `json:"deadTuples,omitempty"`
		SeqScan          *int64  `json:"seqScan,omitempty"`
		IndexScan        *int64  `json:"indexScan,omitempty"`
		LastVacuum       *string `json:"lastVacuum,omitempty"`
		LastAutoVacuum   *string `json:"lastAutoVacuum,omitempty"`
		LastAnalyze      *string `json:"lastAnalyze,omitempty"`
		LastAutoAnalyze  *string `json:"lastAutoAnalyze,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "schema", "name", "relationOid", "totalBytes", "tableBytes", "indexBytes", "toastBytes", "toastRelationOid", "liveTuples", "deadTuples", "seqScan", "indexScan", "lastVacuum", "lastAutoVacuum", "lastAnalyze", "lastAutoAnalyze"})
	} else {
		return err
	}
	o.Database = all.Database
	o.Schema = all.Schema
	o.Name = all.Name
	o.RelationOid = all.RelationOid
	o.TotalBytes = all.TotalBytes
	o.TableBytes = all.TableBytes
	o.IndexBytes = all.IndexBytes
	o.ToastBytes = all.ToastBytes
	o.ToastRelationOid = all.ToastRelationOid
	o.LiveTuples = all.LiveTuples
	o.DeadTuples = all.DeadTuples
	o.SeqScan = all.SeqScan
	o.IndexScan = all.IndexScan
	o.LastVacuum = all.LastVacuum
	o.LastAutoVacuum = all.LastAutoVacuum
	o.LastAnalyze = all.LastAnalyze
	o.LastAutoAnalyze = all.LastAutoAnalyze

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
