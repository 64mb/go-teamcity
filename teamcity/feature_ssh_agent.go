package teamcity

import (
	"encoding/json"
)

// FeatureSshAgentOptions represents options needed to create a commit status publisher build feature
type FeatureSshAgentOptions interface {
	Properties() *Properties
}

// FeatureSshAgent represents a commit status publisher build feature. Implements BuildFeature interface
type FeatureSshAgent struct {
	id          string
	disabled    bool
	Options     FeatureSshAgentOptions
	buildTypeID string

	properties *Properties
}

// ID returns the ID for this instance.
func (f *FeatureSshAgent) ID() string {
	return f.id
}

// SetID sets the ID for this instance.
func (f *FeatureSshAgent) SetID(value string) {
	f.id = value
}

// Type returns the "commit-status-publisher", the keyed-type for this build feature instance
func (f *FeatureSshAgent) Type() string {
	return "ssh-agent-build-feature"
}

// Disabled returns whether this build feature is disabled or not.
func (f *FeatureSshAgent) Disabled() bool {
	return f.disabled
}

// SetDisabled sets whether this build feature is disabled or not.
func (f *FeatureSshAgent) SetDisabled(value bool) {
	f.disabled = value
}

// BuildTypeID is a getter for the Build Type ID associated with this build feature.
func (f *FeatureSshAgent) BuildTypeID() string {
	return f.buildTypeID
}

// SetBuildTypeID is a setter for the Build Type ID associated with this build feature.
func (f *FeatureSshAgent) SetBuildTypeID(value string) {
	f.buildTypeID = value
}

// Properties returns a *Properties instance representing a serializable collection to be used.
func (f *FeatureSshAgent) Properties() *Properties {
	return f.properties
}

// MarshalJSON implements JSON serialization for FeatureSshAgent
func (f *FeatureSshAgent) MarshalJSON() ([]byte, error) {
	out := &buildFeatureJSON{
		ID:         f.id,
		Disabled:   NewBool(f.disabled),
		Properties: f.properties,
		Inherited:  NewFalse(),
		Type:       f.Type(),
	}

	return json.Marshal(out)
}

// UnmarshalJSON implements JSON deserialization for FeatureSshAgent
func (f *FeatureSshAgent) UnmarshalJSON(data []byte) error {
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

	opt, err := SshAgentOptionsFromProperties(f.properties)
	if err != nil {
		return err
	}

	f.Options = opt

	return nil
}
