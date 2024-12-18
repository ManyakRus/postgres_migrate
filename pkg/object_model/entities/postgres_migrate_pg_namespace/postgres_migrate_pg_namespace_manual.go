package postgres_migrate_pg_namespace

// Crud_manual_PostgresMigratePgNamespace - объект контроллер crud операций
var Crud_manual_PostgresMigratePgNamespace ICrud_manual_PostgresMigratePgNamespace

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_PostgresMigratePgNamespace interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PostgresMigratePgNamespace) SetCrudManualInterface(crud ICrud_manual_PostgresMigratePgNamespace) {
	Crud_manual_PostgresMigratePgNamespace = crud

	return
}
