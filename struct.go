package main

type FontCask struct {
	Token                         string      `json:"token"`
	FullToken                     string      `json:"full_token"`
	OldTokens                     []string    `json:"old_tokens"`
	Tap                           string      `json:"tap"`
	Name                          []string    `json:"name"`
	Desc                          *string     `json:"desc"`
	Homepage                      string      `json:"homepage"`
	URL                           string      `json:"url"`
	URLSpecs                      URLSpecs    `json:"url_specs"`
	Version                       string      `json:"version"`
	Autobump                      bool        `json:"autobump"`
	NoAutobumpMessage             string      `json:"no_autobump_message"`
	SkipLivecheck                 bool        `json:"skip_livecheck"`
	Installed                     interface{} `json:"installed"`
	InstalledTime                 interface{} `json:"installed_time"`
	BundleVersion                 interface{} `json:"bundle_version"`
	BundleShortVersion            interface{} `json:"bundle_short_version"`
	Outdated                      bool        `json:"outdated"`
	Sha256                        string      `json:"sha256"`
	Artifacts                     []Artifact  `json:"artifacts"`
	Caveats                       interface{} `json:"caveats"`
	DependsOn                     DependsOn   `json:"depends_on"`
	ConflictsWith                 interface{} `json:"conflicts_with"`
	Container                     interface{} `json:"container"`
	Rename                        []string    `json:"rename"`
	AutoUpdates                   interface{} `json:"auto_updates"`
	Deprecated                    bool        `json:"deprecated"`
	DeprecationDate               interface{} `json:"deprecation_date"`
	DeprecationReason             interface{} `json:"deprecation_reason"`
	DeprecationReplacementFormula interface{} `json:"deprecation_replacement_formula"`
	DeprecationReplacementCask    interface{} `json:"deprecation_replacement_cask"`
	Disabled                      bool        `json:"disabled"`
	DisableDate                   interface{} `json:"disable_date"`
	DisableReason                 interface{} `json:"disable_reason"`
	DisableReplacementFormula     interface{} `json:"disable_replacement_formula"`
	DisableReplacementCask        interface{} `json:"disable_replacement_cask"`
	TapGitHead                    string      `json:"tap_git_head"`
	Languages                     []string    `json:"languages"`
	RubySourcePath                string      `json:"ruby_source_path"`
	RubySourceChecksum            Checksum    `json:"ruby_source_checksum"`
	Variations                    interface{} `json:"variations"`
}

type URLSpecs struct {
	Verified string `json:"verified"`
}

type Artifact struct {
	Font []string `json:"font"`
}

type DependsOn struct {
	MacOS MacOSVersion `json:"macos"`
}

type MacOSVersion struct {
	GreaterOrEqual []string `json:">="`
}

type Checksum struct {
	Sha256 string `json:"sha256"`
}
