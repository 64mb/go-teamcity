package teamcity

// SshAgentOptions represents parameters used to create Github Commit Status Publisher Feature
type SshAgentOptions struct {
	SshKey string
}

// NewSshAgentOptions returns options created for AuthenticationType = 'token'. No validation is performed, parameters indicate mandatory fields.
func NewSshAgentOptions(sshKey string) SshAgentOptions {
	return SshAgentOptions{
		SshKey: sshKey,
	}
}

// NewFeatureSshAgent creates a Build Feature Commit status Publisher to Github with the given options and validates the required properties.
// VcsRootID is optional - if empty, it will apply the commit publisher feature to all VCS roots.
func NewFeatureSshAgent(opt SshAgentOptions) (*FeatureSshAgent, error) {
	out := &FeatureSshAgent{
		Options:    opt,
		properties: opt.Properties(),
	}

	return out, nil
}

// Properties returns a *Properties collection with properties filled related to this commit publisher parameters to be used in build features
func (s SshAgentOptions) Properties() *Properties {
	props := NewPropertiesEmpty()

	props.AddOrReplaceValue("teamcitySshKey", s.SshKey)

	return props
}

// SshAgentOptionsFromProperties grabs a Properties collection and transforms back to a SshAgentOptions
func SshAgentOptionsFromProperties(p *Properties) (*SshAgentOptions, error) {
	var out SshAgentOptions

	if v, ok := p.GetOk("teamcitySshKey"); ok {
		out.SshKey = v
	}

	return &out, nil
}
