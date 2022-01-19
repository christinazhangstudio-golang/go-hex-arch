package ports

//Add arithmetic results to history table in db
type DbPort interface {
	CloseDbConnection()
	AddToHistory(answer int32, operation string) error
}