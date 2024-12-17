package pg_description

// Crud_manual_PgDescription - объект контроллер crud операций
var Crud_manual_PgDescription ICrud_manual_PgDescription

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_PgDescription interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PgDescription) SetCrudManualInterface(crud ICrud_manual_PgDescription) {
	Crud_manual_PgDescription = crud

	return
}
