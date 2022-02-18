package playbook


type tasks struct {
	Name string `yaml:"name"`
	Fail struct {
		Msg string `yaml:"msg"`
	} `yaml:"fail"`
	When string `yaml:"when"`
	WithItems []string `yaml:with_items`
}

type Playbook struct {
	Name string `yaml:"name"`
	Hosts string `yaml:"hosts"`
	GatherFacts bool `yaml:"gather_facts"`
	Become bool `yaml:"become"`
	BecomeMethod string `yaml:"become_method"`
	Tasks []tasks `yaml:"tasks"`
	Roles []string `yaml:"roles"`
}

