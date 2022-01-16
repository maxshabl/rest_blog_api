package domaininterface

type CommandInterface interface {
	Validate() (bool, error)
}
