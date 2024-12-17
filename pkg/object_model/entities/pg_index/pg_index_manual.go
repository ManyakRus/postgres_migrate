package pg_index

// Crud_manual_PgIndex - объект контроллер crud операций
var Crud_manual_PgIndex ICrud_manual_PgIndex

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_PgIndex interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PgIndex) SetCrudManualInterface(crud ICrud_manual_PgIndex) {
	Crud_manual_PgIndex = crud

	return
}
