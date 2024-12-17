package pg_class

// Crud_manual_PgClass - объект контроллер crud операций
var Crud_manual_PgClass ICrud_manual_PgClass

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_PgClass interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PgClass) SetCrudManualInterface(crud ICrud_manual_PgClass) {
	Crud_manual_PgClass = crud

	return
}
