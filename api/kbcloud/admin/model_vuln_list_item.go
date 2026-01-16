// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// VulnListItem a vulnerability which affected the cluster
type VulnListItem struct {
	// the CVE ID of the vulnerability
	CveId string `json:"cveId"`
	// the engine name of the vulnerability
	EngineName string `json:"engineName"`
	// the product name of the vulnerability
	ProductName string `json:"productName"`
	// the version of the product which is affected by the vulnerability
	Version string `json:"version"`
	// the detail of the vulnerability
	Detail *string `json:"detail,omitempty"`
	// the severity of the vulnerability
	Severity string `json:"severity"`
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

// NewVulnListItem instantiates a new VulnListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVulnListItem(cveId string, engineName string, productName string, version string, severity string, publishedAt time.Time, modifiedAt time.Time) *VulnListItem {
	this := VulnListItem{}
	this.CveId = cveId
	this.EngineName = engineName
	this.ProductName = productName
	this.Version = version
	this.Severity = severity
	this.PublishedAt = publishedAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewVulnListItemWithDefaults instantiates a new VulnListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVulnListItemWithDefaults() *VulnListItem {
	this := VulnListItem{}
	return &this
}

// GetCveId returns the CveId field value.
func (o *VulnListItem) GetCveId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CveId
}

// GetCveIdOk returns a tuple with the CveId field value
// and a boolean to check if the value has been set.
func (o *VulnListItem) GetCveIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CveId, true
}

// SetCveId sets field value.
func (o *VulnListItem) SetCveId(v string) {
	o.CveId = v
}

// GetEngineName returns the EngineName field value.
func (o *VulnListItem) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *VulnListItem) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *VulnListItem) SetEngineName(v string) {
	o.EngineName = v
}

// GetProductName returns the ProductName field value.
func (o *VulnListItem) GetProductName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value
// and a boolean to check if the value has been set.
func (o *VulnListItem) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductName, true
}

// SetProductName sets field value.
func (o *VulnListItem) SetProductName(v string) {
	o.ProductName = v
}

// GetVersion returns the Version field value.
func (o *VulnListItem) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *VulnListItem) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *VulnListItem) SetVersion(v string) {
	o.Version = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *VulnListItem) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnListItem) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *VulnListItem) HasDetail() bool {
	return o != nil && o.Detail != nil
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *VulnListItem) SetDetail(v string) {
	o.Detail = &v
}

// GetSeverity returns the Severity field value.
func (o *VulnListItem) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *VulnListItem) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *VulnListItem) SetSeverity(v string) {
	o.Severity = v
}

// GetPublishedAt returns the PublishedAt field value.
func (o *VulnListItem) GetPublishedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value
// and a boolean to check if the value has been set.
func (o *VulnListItem) GetPublishedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishedAt, true
}

// SetPublishedAt sets field value.
func (o *VulnListItem) SetPublishedAt(v time.Time) {
	o.PublishedAt = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *VulnListItem) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *VulnListItem) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *VulnListItem) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetRefs returns the Refs field value if set, zero value otherwise.
func (o *VulnListItem) GetRefs() []RefItem {
	if o == nil || o.Refs == nil {
		var ret []RefItem
		return ret
	}
	return o.Refs
}

// GetRefsOk returns a tuple with the Refs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnListItem) GetRefsOk() (*[]RefItem, bool) {
	if o == nil || o.Refs == nil {
		return nil, false
	}
	return &o.Refs, true
}

// HasRefs returns a boolean if a field has been set.
func (o *VulnListItem) HasRefs() bool {
	return o != nil && o.Refs != nil
}

// SetRefs gets a reference to the given []RefItem and assigns it to the Refs field.
func (o *VulnListItem) SetRefs(v []RefItem) {
	o.Refs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o VulnListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["cveId"] = o.CveId
	toSerialize["engineName"] = o.EngineName
	toSerialize["productName"] = o.ProductName
	toSerialize["version"] = o.Version
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	toSerialize["severity"] = o.Severity
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
func (o *VulnListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CveId       *string    `json:"cveId"`
		EngineName  *string    `json:"engineName"`
		ProductName *string    `json:"productName"`
		Version     *string    `json:"version"`
		Detail      *string    `json:"detail,omitempty"`
		Severity    *string    `json:"severity"`
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
	if all.EngineName == nil {
		return fmt.Errorf("required field engineName missing")
	}
	if all.ProductName == nil {
		return fmt.Errorf("required field productName missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
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
		common.DeleteKeys(additionalProperties, &[]string{"cveId", "engineName", "productName", "version", "detail", "severity", "publishedAt", "modifiedAt", "refs"})
	} else {
		return err
	}
	o.CveId = *all.CveId
	o.EngineName = *all.EngineName
	o.ProductName = *all.ProductName
	o.Version = *all.Version
	o.Detail = all.Detail
	o.Severity = *all.Severity
	o.PublishedAt = *all.PublishedAt
	o.ModifiedAt = *all.ModifiedAt
	o.Refs = all.Refs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
