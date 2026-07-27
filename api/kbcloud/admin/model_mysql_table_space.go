// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MysqlTableSpace struct {
	Database string `json:"database"`
	Name     string `json:"name"`
	Engine   string `json:"engine"`
	// Estimated row count for InnoDB tables.
	Rows         int64 `json:"rows"`
	DataBytes    int64 `json:"dataBytes"`
	IndexBytes   int64 `json:"indexBytes"`
	FreeBytes    int64 `json:"freeBytes"`
	TotalBytes   int64 `json:"totalBytes"`
	AvgRowLength int64 `json:"avgRowLength"`
	// freeBytes divided by totalBytes plus freeBytes when available.
	FragmentationRatio common.NullableFloat64 `json:"fragmentationRatio,omitempty"`
	CreateTime         *string                `json:"createTime,omitempty"`
	UpdateTime         *string                `json:"updateTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMysqlTableSpace instantiates a new MysqlTableSpace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMysqlTableSpace(database string, name string, engine string, rows int64, dataBytes int64, indexBytes int64, freeBytes int64, totalBytes int64, avgRowLength int64) *MysqlTableSpace {
	this := MysqlTableSpace{}
	this.Database = database
	this.Name = name
	this.Engine = engine
	this.Rows = rows
	this.DataBytes = dataBytes
	this.IndexBytes = indexBytes
	this.FreeBytes = freeBytes
	this.TotalBytes = totalBytes
	this.AvgRowLength = avgRowLength
	return &this
}

// NewMysqlTableSpaceWithDefaults instantiates a new MysqlTableSpace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMysqlTableSpaceWithDefaults() *MysqlTableSpace {
	this := MysqlTableSpace{}
	return &this
}

// GetDatabase returns the Database field value.
func (o *MysqlTableSpace) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *MysqlTableSpace) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *MysqlTableSpace) SetDatabase(v string) {
	o.Database = v
}

// GetName returns the Name field value.
func (o *MysqlTableSpace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MysqlTableSpace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MysqlTableSpace) SetName(v string) {
	o.Name = v
}

// GetEngine returns the Engine field value.
func (o *MysqlTableSpace) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *MysqlTableSpace) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *MysqlTableSpace) SetEngine(v string) {
	o.Engine = v
}

// GetRows returns the Rows field value.
func (o *MysqlTableSpace) GetRows() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value
// and a boolean to check if the value has been set.
func (o *MysqlTableSpace) GetRowsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rows, true
}

// SetRows sets field value.
func (o *MysqlTableSpace) SetRows(v int64) {
	o.Rows = v
}

// GetDataBytes returns the DataBytes field value.
func (o *MysqlTableSpace) GetDataBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DataBytes
}

// GetDataBytesOk returns a tuple with the DataBytes field value
// and a boolean to check if the value has been set.
func (o *MysqlTableSpace) GetDataBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataBytes, true
}

// SetDataBytes sets field value.
func (o *MysqlTableSpace) SetDataBytes(v int64) {
	o.DataBytes = v
}

// GetIndexBytes returns the IndexBytes field value.
func (o *MysqlTableSpace) GetIndexBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.IndexBytes
}

// GetIndexBytesOk returns a tuple with the IndexBytes field value
// and a boolean to check if the value has been set.
func (o *MysqlTableSpace) GetIndexBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexBytes, true
}

// SetIndexBytes sets field value.
func (o *MysqlTableSpace) SetIndexBytes(v int64) {
	o.IndexBytes = v
}

// GetFreeBytes returns the FreeBytes field value.
func (o *MysqlTableSpace) GetFreeBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FreeBytes
}

// GetFreeBytesOk returns a tuple with the FreeBytes field value
// and a boolean to check if the value has been set.
func (o *MysqlTableSpace) GetFreeBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FreeBytes, true
}

// SetFreeBytes sets field value.
func (o *MysqlTableSpace) SetFreeBytes(v int64) {
	o.FreeBytes = v
}

// GetTotalBytes returns the TotalBytes field value.
func (o *MysqlTableSpace) GetTotalBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value
// and a boolean to check if the value has been set.
func (o *MysqlTableSpace) GetTotalBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalBytes, true
}

// SetTotalBytes sets field value.
func (o *MysqlTableSpace) SetTotalBytes(v int64) {
	o.TotalBytes = v
}

// GetAvgRowLength returns the AvgRowLength field value.
func (o *MysqlTableSpace) GetAvgRowLength() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AvgRowLength
}

// GetAvgRowLengthOk returns a tuple with the AvgRowLength field value
// and a boolean to check if the value has been set.
func (o *MysqlTableSpace) GetAvgRowLengthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgRowLength, true
}

// SetAvgRowLength sets field value.
func (o *MysqlTableSpace) SetAvgRowLength(v int64) {
	o.AvgRowLength = v
}

// GetFragmentationRatio returns the FragmentationRatio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MysqlTableSpace) GetFragmentationRatio() float64 {
	if o == nil || o.FragmentationRatio.Get() == nil {
		var ret float64
		return ret
	}
	return *o.FragmentationRatio.Get()
}

// GetFragmentationRatioOk returns a tuple with the FragmentationRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MysqlTableSpace) GetFragmentationRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FragmentationRatio.Get(), o.FragmentationRatio.IsSet()
}

// HasFragmentationRatio returns a boolean if a field has been set.
func (o *MysqlTableSpace) HasFragmentationRatio() bool {
	return o != nil && o.FragmentationRatio.IsSet()
}

// SetFragmentationRatio gets a reference to the given common.NullableFloat64 and assigns it to the FragmentationRatio field.
func (o *MysqlTableSpace) SetFragmentationRatio(v float64) {
	o.FragmentationRatio.Set(&v)
}

// SetFragmentationRatioNil sets the value for FragmentationRatio to be an explicit nil.
func (o *MysqlTableSpace) SetFragmentationRatioNil() {
	o.FragmentationRatio.Set(nil)
}

// UnsetFragmentationRatio ensures that no value is present for FragmentationRatio, not even an explicit nil.
func (o *MysqlTableSpace) UnsetFragmentationRatio() {
	o.FragmentationRatio.Unset()
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *MysqlTableSpace) GetCreateTime() string {
	if o == nil || o.CreateTime == nil {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlTableSpace) GetCreateTimeOk() (*string, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *MysqlTableSpace) HasCreateTime() bool {
	return o != nil && o.CreateTime != nil
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *MysqlTableSpace) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *MysqlTableSpace) GetUpdateTime() string {
	if o == nil || o.UpdateTime == nil {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlTableSpace) GetUpdateTimeOk() (*string, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *MysqlTableSpace) HasUpdateTime() bool {
	return o != nil && o.UpdateTime != nil
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *MysqlTableSpace) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MysqlTableSpace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["database"] = o.Database
	toSerialize["name"] = o.Name
	toSerialize["engine"] = o.Engine
	toSerialize["rows"] = o.Rows
	toSerialize["dataBytes"] = o.DataBytes
	toSerialize["indexBytes"] = o.IndexBytes
	toSerialize["freeBytes"] = o.FreeBytes
	toSerialize["totalBytes"] = o.TotalBytes
	toSerialize["avgRowLength"] = o.AvgRowLength
	if o.FragmentationRatio.IsSet() {
		toSerialize["fragmentationRatio"] = o.FragmentationRatio.Get()
	}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.UpdateTime != nil {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MysqlTableSpace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database           *string                `json:"database"`
		Name               *string                `json:"name"`
		Engine             *string                `json:"engine"`
		Rows               *int64                 `json:"rows"`
		DataBytes          *int64                 `json:"dataBytes"`
		IndexBytes         *int64                 `json:"indexBytes"`
		FreeBytes          *int64                 `json:"freeBytes"`
		TotalBytes         *int64                 `json:"totalBytes"`
		AvgRowLength       *int64                 `json:"avgRowLength"`
		FragmentationRatio common.NullableFloat64 `json:"fragmentationRatio,omitempty"`
		CreateTime         *string                `json:"createTime,omitempty"`
		UpdateTime         *string                `json:"updateTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	if all.Rows == nil {
		return fmt.Errorf("required field rows missing")
	}
	if all.DataBytes == nil {
		return fmt.Errorf("required field dataBytes missing")
	}
	if all.IndexBytes == nil {
		return fmt.Errorf("required field indexBytes missing")
	}
	if all.FreeBytes == nil {
		return fmt.Errorf("required field freeBytes missing")
	}
	if all.TotalBytes == nil {
		return fmt.Errorf("required field totalBytes missing")
	}
	if all.AvgRowLength == nil {
		return fmt.Errorf("required field avgRowLength missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "name", "engine", "rows", "dataBytes", "indexBytes", "freeBytes", "totalBytes", "avgRowLength", "fragmentationRatio", "createTime", "updateTime"})
	} else {
		return err
	}
	o.Database = *all.Database
	o.Name = *all.Name
	o.Engine = *all.Engine
	o.Rows = *all.Rows
	o.DataBytes = *all.DataBytes
	o.IndexBytes = *all.IndexBytes
	o.FreeBytes = *all.FreeBytes
	o.TotalBytes = *all.TotalBytes
	o.AvgRowLength = *all.AvgRowLength
	o.FragmentationRatio = all.FragmentationRatio
	o.CreateTime = all.CreateTime
	o.UpdateTime = all.UpdateTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
