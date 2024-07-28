package teamcity

import "fmt"

// PullRequestsGithubOptions represents parameters used to create Github Commit Status Publisher Feature
type PullRequestsGithubOptions struct {
	//AuthenticationType can be 'password' or 'token'
	AuthenticationType string
	//Username is required if AuthenticationType is 'password'
	Username string
	//Password is required if AuthenticationType is 'password'
	Password string
	//AccessToken is required if AuthenticationType is 'token'
	AccessToken string
	// additional
	FilterAuthorRole   string
	FilterSourceBranch []string
	FilterTargetBranch []string
}

// NewPullRequestsGithubOptionsPassword returns options created for AuthenticationType = 'password'. No validation is performed, parameters indicate mandatory fields.
func NewPullRequestsGithubOptionsPassword(host string, username string, password string, filterAuthorRole string) PullRequestsGithubOptions {
	return PullRequestsGithubOptions{
		FilterAuthorRole:   filterAuthorRole,
		FilterSourceBranch: []string{},
		FilterTargetBranch: []string{},
		AuthenticationType: "password",
		Username:           username,
		Password:           password,
	}
}

// NewPullRequestsGithubOptionsToken returns options created for AuthenticationType = 'token'. No validation is performed, parameters indicate mandatory fields.
func NewPullRequestsGithubOptionsToken(host string, accessToken string, filterAuthorRole string) PullRequestsGithubOptions {
	return PullRequestsGithubOptions{
		FilterAuthorRole:   filterAuthorRole,
		FilterSourceBranch: []string{},
		FilterTargetBranch: []string{},
		AuthenticationType: "token",
		AccessToken:        accessToken,
	}
}

// NewFeaturePullRequestsGithub creates a Build Feature Commit status Publisher to Github with the given options and validates the required properties.
// VcsRootID is optional - if empty, it will apply the commit publisher feature to all VCS roots.
func NewFeaturePullRequestsGithub(opt PullRequestsGithubOptions, vcsRootID string) (*FeaturePullRequests, error) {
	if opt.AuthenticationType == "" {
		return nil, fmt.Errorf("AuthenticationType is required")
	}

	if opt.AuthenticationType != "password" && opt.AuthenticationType != "token" {
		return nil, fmt.Errorf("invalid AuthenticationType, must be 'password' or 'token'")
	}

	if opt.AuthenticationType == "password" {
		if opt.Username == "" || opt.Password == "" {
			return nil, fmt.Errorf("username/password required for auth type 'password'")
		}
	}

	if opt.AuthenticationType == "token" {
		if opt.AccessToken == "" {
			return nil, fmt.Errorf("access token required for auth type 'token'")
		}
	}

	out := &FeaturePullRequests{
		Options:    opt,
		properties: opt.Properties(),
	}

	if vcsRootID != "" {
		out.vcsRootID = vcsRootID
	}

	return out, nil
}

// Properties returns a *Properties collection with properties filled related to this commit publisher parameters to be used in build features
func (s PullRequestsGithubOptions) Properties() *Properties {
	props := NewPropertiesEmpty()

	props.AddOrReplaceValue("provider_type", "github")
	props.AddOrReplaceValue("authentication_type", s.AuthenticationType)
	props.AddOrReplaceValue("filter_author_role", s.FilterAuthorRole)

	if s.AuthenticationType == "password" {
		props.AddOrReplaceValue("username", s.Username)
		props.AddOrReplaceValue("secure:password", s.Password)
	}

	if s.AuthenticationType == "token" {
		props.AddOrReplaceValue("secure:access_token", s.AccessToken)
	}

	return props
}

// PullRequestsGithubOptionsFromProperties grabs a Properties collection and transforms back to a PullRequestsGithubOptions
func PullRequestsGithubOptionsFromProperties(p *Properties) (*PullRequestsGithubOptions, error) {
	var out PullRequestsGithubOptions

	if authType, ok := p.GetOk("authentication_type"); ok {
		out.AuthenticationType = authType
		switch authType {
		case "password":
			u, _ := p.GetOk("username")
			out.Username = u

			// Password or AccessToken is never returned from properties, because it is secure. Once set, we cannot read it back
		}
	} else {
		return nil, fmt.Errorf("Properties do not have 'access_token' key")
	}

	return &out, nil
}
