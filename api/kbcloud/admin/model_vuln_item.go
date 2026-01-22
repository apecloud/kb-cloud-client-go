// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// VulnItem a vulnerability which affected the cluster
type VulnItem struct {
	// the CVE ID of the vulnerability
	CveId string `json:"cveId"`
	// the detail of the vulnerability
	Detail *string `json:"detail,omitempty"`
	// the severity of the vulnerability
	Severity string `json:"severity"`
	// the CVSS vector and score of the vulnerability
	CvssVector *string `json:"cvssVector,omitempty"`
	// the published time of the vulnerability
	PublishedAt time.Time `json:"publishedAt"`
	// the modified time of the vulnerability
	ModifiedAt time.Time `json:"modifiedAt"`
	// the references of the vulnerability
	Refs []RefItem `json:"refs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVulnItem instantiates a new VulnItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVulnItem(cveId string, severity string, publishedAt time.Time, modifiedAt time.Time) *VulnItem {
	this := VulnItem{}
	this.CveId = cveId
	this.Severity = severity
	this.PublishedAt = publishedAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewVulnItemWithDefaults instantiates a new VulnItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVulnItemWithDefaults() *VulnItem {
	this := VulnItem{}
	return &this
}

// GetCveId returns the CveId field value.
func (o *VulnItem) GetCveId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CveId
}

// GetCveIdOk returns a tuple with the CveId field value
// and a boolean to check if the value has been set.
func (o *VulnItem) GetCveIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CveId, true
}

// SetCveId sets field value.
func (o *VulnItem) SetCveId(v string) {
	o.CveId = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *VulnItem) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnItem) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *VulnItem) HasDetail() bool {
	return o != nil && o.Detail != nil
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *VulnItem) SetDetail(v string) {
	o.Detail = &v
}

// GetSeverity returns the Severity field value.
func (o *VulnItem) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *VulnItem) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *VulnItem) SetSeverity(v string) {
	o.Severity = v
}

// GetCvssVector returns the CvssVector field value if set, zero value otherwise.
func (o *VulnItem) GetCvssVector() string {
	if o == nil || o.CvssVector == nil {
		var ret string
		return ret
	}
	return *o.CvssVector
}

// GetCvssVectorOk returns a tuple with the CvssVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnItem) GetCvssVectorOk() (*string, bool) {
	if o == nil || o.CvssVector == nil {
		return nil, false
	}
	return o.CvssVector, true
}

// HasCvssVector returns a boolean if a field has been set.
func (o *VulnItem) HasCvssVector() bool {
	return o != nil && o.CvssVector != nil
}

// SetCvssVector gets a reference to the given string and assigns it to the CvssVector field.
func (o *VulnItem) SetCvssVector(v string) {
	o.CvssVector = &v
}

// GetPublishedAt returns the PublishedAt field value.
func (o *VulnItem) GetPublishedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value
// and a boolean to check if the value has been set.
func (o *VulnItem) GetPublishedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishedAt, true
}

// SetPublishedAt sets field value.
func (o *VulnItem) SetPublishedAt(v time.Time) {
	o.PublishedAt = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *VulnItem) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *VulnItem) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *VulnItem) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetRefs returns the Refs field value if set, zero value otherwise.
func (o *VulnItem) GetRefs() []RefItem {
	if o == nil || o.Refs == nil {
		var ret []RefItem
		return ret
	}
	return o.Refs
}

// GetRefsOk returns a tuple with the Refs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnItem) GetRefsOk() (*[]RefItem, bool) {
	if o == nil || o.Refs == nil {
		return nil, false
	}
	return &o.Refs, true
}

// HasRefs returns a boolean if a field has been set.
func (o *VulnItem) HasRefs() bool {
	return o != nil && o.Refs != nil
}

// SetRefs gets a reference to the given []RefItem and assigns it to the Refs field.
func (o *VulnItem) SetRefs(v []RefItem) {
	o.Refs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o VulnItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["cveId"] = o.CveId
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	toSerialize["severity"] = o.Severity
	if o.CvssVector != nil {
		toSerialize["cvssVector"] = o.CvssVector
	}
	if o.PublishedAt.Nanosecond() == 0 {
		toSerialize["publishedAt"] = o.PublishedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["publishedAt"] = o.PublishedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modifiedAt"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modifiedAt"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Refs != nil {
		toSerialize["refs"] = o.Refs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VulnItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CveId       *string    `json:"cveId"`
		Detail      *string    `json:"detail,omitempty"`
		Severity    *string    `json:"severity"`
		CvssVector  *string    `json:"cvssVector,omitempty"`
		PublishedAt *time.Time `json:"publishedAt"`
		ModifiedAt  *time.Time `json:"modifiedAt"`
		Refs        []RefItem  `json:"refs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CveId == nil {
		return fmt.Errorf("required field cveId missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.PublishedAt == nil {
		return fmt.Errorf("required field publishedAt missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modifiedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cveId", "detail", "severity", "cvssVector", "publishedAt", "modifiedAt", "refs"})
	} else {
		return err
	}
	o.CveId = *all.CveId
	o.Detail = all.Detail
	o.Severity = *all.Severity
	o.CvssVector = all.CvssVector
	o.PublishedAt = *all.PublishedAt
	o.ModifiedAt = *all.ModifiedAt
	o.Refs = all.Refs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
