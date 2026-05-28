// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSpaceAnalysisThreshold Readonly collection thresholds and limits used by space analysis.
type DiagnosticSpaceAnalysisThreshold struct {
	// Maximum number of database objects returned in topObjects.
	TopN int64 `json:"topN"`
	// Stable ordering key for topObjects.
	SortBy DiagnosticSpaceAnalysisThresholdSortBy `json:"sortBy"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticSpaceAnalysisThreshold instantiates a new DiagnosticSpaceAnalysisThreshold object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticSpaceAnalysisThreshold(topN int64, sortBy DiagnosticSpaceAnalysisThresholdSortBy) *DiagnosticSpaceAnalysisThreshold {
	this := DiagnosticSpaceAnalysisThreshold{}
	this.TopN = topN
	this.SortBy = sortBy
	return &this
}

// NewDiagnosticSpaceAnalysisThresholdWithDefaults instantiates a new DiagnosticSpaceAnalysisThreshold object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticSpaceAnalysisThresholdWithDefaults() *DiagnosticSpaceAnalysisThreshold {
	this := DiagnosticSpaceAnalysisThreshold{}
	return &this
}

// GetTopN returns the TopN field value.
func (o *DiagnosticSpaceAnalysisThreshold) GetTopN() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TopN
}

// GetTopNOk returns a tuple with the TopN field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisThreshold) GetTopNOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopN, true
}

// SetTopN sets field value.
func (o *DiagnosticSpaceAnalysisThreshold) SetTopN(v int64) {
	o.TopN = v
}

// GetSortBy returns the SortBy field value.
func (o *DiagnosticSpaceAnalysisThreshold) GetSortBy() DiagnosticSpaceAnalysisThresholdSortBy {
	if o == nil {
		var ret DiagnosticSpaceAnalysisThresholdSortBy
		return ret
	}
	return o.SortBy
}

// GetSortByOk returns a tuple with the SortBy field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisThreshold) GetSortByOk() (*DiagnosticSpaceAnalysisThresholdSortBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortBy, true
}

// SetSortBy sets field value.
func (o *DiagnosticSpaceAnalysisThreshold) SetSortBy(v DiagnosticSpaceAnalysisThresholdSortBy) {
	o.SortBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticSpaceAnalysisThreshold) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["topN"] = o.TopN
	toSerialize["sortBy"] = o.SortBy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticSpaceAnalysisThreshold) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TopN   *int64                                  `json:"topN"`
		SortBy *DiagnosticSpaceAnalysisThresholdSortBy `json:"sortBy"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TopN == nil {
		return fmt.Errorf("required field topN missing")
	}
	if all.SortBy == nil {
		return fmt.Errorf("required field sortBy missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"topN", "sortBy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.TopN = *all.TopN
	if !all.SortBy.IsValid() {
		hasInvalidField = true
	} else {
		o.SortBy = *all.SortBy
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
