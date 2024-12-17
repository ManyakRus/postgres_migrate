package pg_namespace

// Crud_manual_PgNamespace - объект контроллер crud операций
var Crud_manual_PgNamespace ICrud_manual_PgNamespace

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_PgNamespace interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PgNamespace) SetCrudManualInterface(crud ICrud_manual_PgNamespace) {
	Crud_manual_PgNamespace = crud

	return
}
