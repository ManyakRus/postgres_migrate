package pg_attribute

// Crud_manual_PgAttribute - объект контроллер crud операций
var Crud_manual_PgAttribute ICrud_manual_PgAttribute

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_PgAttribute interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PgAttribute) SetCrudManualInterface(crud ICrud_manual_PgAttribute) {
	Crud_manual_PgAttribute = crud

	return
}
