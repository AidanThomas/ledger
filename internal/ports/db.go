package ports

type database interface {
	ExecuteCommand(string) error
}
