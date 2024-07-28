package teamcity

import (
	"encoding/json"
)

// FeaturePullRequestsOptions represents options needed to create a commit status publisher build feature
type FeaturePullRequestsOptions interface {
	Properties() *Properties
}

// FeaturePullRequests represents a commit status publisher build feature. Implements BuildFeature interface
type FeaturePullRequests struct {
	id          string
	vcsRootID   string
	disabled    bool
	Options     FeaturePullRequestsOptions
	buildTypeID string

	properties *Properties
}

// ID returns the ID for this instance.
func (f *FeaturePullRequests) ID() string {
	return f.id
}

// SetID sets the ID for this instance.
func (f *FeaturePullRequests) SetID(value string) {
	f.id = value
}

// Type returns the "commit-status-publisher", the keyed-type for this build feature instance
func (f *FeaturePullRequests) Type() string {
	return "pullRequests"
}

// VcsRootID returns the VCS Root ID that this build feature is associated with.
func (f *FeaturePullRequests) VcsRootID() string {
	return f.vcsRootID
}

// SetVcsRootID sets the VCS Root ID that this build feature is associated with.
func (f *FeaturePullRequests) SetVcsRootID(value string) {
	f.vcsRootID = value
}

// Disabled returns whether this build feature is disabled or not.
func (f *FeaturePullRequests) Disabled() bool {
	return f.disabled
}

// SetDisabled sets whether this build feature is disabled or not.
func (f *FeaturePullRequests) SetDisabled(value bool) {
	f.disabled = value
}

// BuildTypeID is a getter for the Build Type ID associated with this build feature.
func (f *FeaturePullRequests) BuildTypeID() string {
	return f.buildTypeID
}

// SetBuildTypeID is a setter for the Build Type ID associated with this build feature.
func (f *FeaturePullRequests) SetBuildTypeID(value string) {
	f.buildTypeID = value
}

// Properties returns a *Properties instance representing a serializable collection to be used.
func (f *FeaturePullRequests) Properties() *Properties {
	return f.properties
}

// MarshalJSON implements JSON serialization for FeaturePullRequests
func (f *FeaturePullRequests) MarshalJSON() ([]byte, error) {
	out := &buildFeatureJSON{
		ID:         f.id,
		Disabled:   NewBool(f.disabled),
		Properties: f.properties,
		Inherited:  NewFalse(),
		Type:       f.Type(),
	}

	if f.vcsRootID != "" {
		out.Properties.AddOrReplaceValue("vcsRootId", f.vcsRootID)
	}
	return json.Marshal(out)
}

// UnmarshalJSON implements JSON deserialization for FeaturePullRequests
func (f *FeaturePullRequests) UnmarshalJSON(data []byte) error {
	var aux buildFeatureJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	f.id = aux.ID

	disabled := aux.Disabled
	if disabled == nil {
		disabled = NewFalse()
	}
	f.disabled = *disabled
	f.properties = NewProperties(aux.Properties.Items...)

	opt, err := PullRequestsGithubOptionsFromProperties(f.properties)
	if err != nil {
		return err
	}

	if v, ok := f.properties.GetOk("vcsRootId"); ok {
		f.vcsRootID = v
	}
	f.Options = opt

	return nil
}
