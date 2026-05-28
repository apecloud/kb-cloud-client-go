// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSpaceAnalysisTotalSize Total PostgreSQL database size observed by the readonly collector.
type DiagnosticSpaceAnalysisTotalSize struct {
	// PostgreSQL database used for this readonly space analysis.
	DatabaseName string `json:"databaseName"`
	// pg_database_size(current_database()) in bytes.
	TotalSizeBytes int64 `json:"totalSizeBytes"`
	// Human-readable database size produced from totalSizeBytes.
	TotalSizeHuman string `json:"totalSizeHuman"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticSpaceAnalysisTotalSize instantiates a new DiagnosticSpaceAnalysisTotalSize object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticSpaceAnalysisTotalSize(databaseName string, totalSizeBytes int64, totalSizeHuman string) *DiagnosticSpaceAnalysisTotalSize {
	this := DiagnosticSpaceAnalysisTotalSize{}
	this.DatabaseName = databaseName
	this.TotalSizeBytes = totalSizeBytes
	this.TotalSizeHuman = totalSizeHuman
	return &this
}

// NewDiagnosticSpaceAnalysisTotalSizeWithDefaults instantiates a new DiagnosticSpaceAnalysisTotalSize object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticSpaceAnalysisTotalSizeWithDefaults() *DiagnosticSpaceAnalysisTotalSize {
	this := DiagnosticSpaceAnalysisTotalSize{}
	return &this
}

// GetDatabaseName returns the DatabaseName field value.
func (o *DiagnosticSpaceAnalysisTotalSize) GetDatabaseName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisTotalSize) GetDatabaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseName, true
}

// SetDatabaseName sets field value.
func (o *DiagnosticSpaceAnalysisTotalSize) SetDatabaseName(v string) {
	o.DatabaseName = v
}

// GetTotalSizeBytes returns the TotalSizeBytes field value.
func (o *DiagnosticSpaceAnalysisTotalSize) GetTotalSizeBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalSizeBytes
}

// GetTotalSizeBytesOk returns a tuple with the TotalSizeBytes field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisTotalSize) GetTotalSizeBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSizeBytes, true
}

// SetTotalSizeBytes sets field value.
func (o *DiagnosticSpaceAnalysisTotalSize) SetTotalSizeBytes(v int64) {
	o.TotalSizeBytes = v
}

// GetTotalSizeHuman returns the TotalSizeHuman field value.
func (o *DiagnosticSpaceAnalysisTotalSize) GetTotalSizeHuman() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TotalSizeHuman
}

// GetTotalSizeHumanOk returns a tuple with the TotalSizeHuman field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisTotalSize) GetTotalSizeHumanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSizeHuman, true
}

// SetTotalSizeHuman sets field value.
func (o *DiagnosticSpaceAnalysisTotalSize) SetTotalSizeHuman(v string) {
	o.TotalSizeHuman = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticSpaceAnalysisTotalSize) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["databaseName"] = o.DatabaseName
	toSerialize["totalSizeBytes"] = o.TotalSizeBytes
	toSerialize["totalSizeHuman"] = o.TotalSizeHuman

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticSpaceAnalysisTotalSize) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatabaseName   *string `json:"databaseName"`
		TotalSizeBytes *int64  `json:"totalSizeBytes"`
		TotalSizeHuman *string `json:"totalSizeHuman"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.DatabaseName == nil {
		return fmt.Errorf("required field databaseName missing")
	}
	if all.TotalSizeBytes == nil {
		return fmt.Errorf("required field totalSizeBytes missing")
	}
	if all.TotalSizeHuman == nil {
		return fmt.Errorf("required field totalSizeHuman missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"databaseName", "totalSizeBytes", "totalSizeHuman"})
	} else {
		return err
	}
	o.DatabaseName = *all.DatabaseName
	o.TotalSizeBytes = *all.TotalSizeBytes
	o.TotalSizeHuman = *all.TotalSizeHuman

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
