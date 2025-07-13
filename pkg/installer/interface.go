package installer

type Installer interface {
	Install(path string) error
	Uninstall() error
	GetStatus() string
}
