package tasks

type Synchronize struct {
	Source string `yaml:"src"`
	Destination string `yaml:"dest"`
	Delete Boolean `yaml:"delete"`
	RSyncOpts []string `yaml:"rsync_opts"`
}