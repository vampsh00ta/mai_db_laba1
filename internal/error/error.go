package error

type Error struct {
	Err         error
	FuncName    string
	PackageName string
}
