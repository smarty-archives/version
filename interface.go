package version

type Repository interface {
	CurrentVersion() (Number, error)
	UpdateVersion(Number) error
}
