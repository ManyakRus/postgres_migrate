package postgres_migrate_pg_index

// Crud_manual_PostgresMigratePgIndex - объект контроллер crud операций
var Crud_manual_PostgresMigratePgIndex ICrud_manual_PostgresMigratePgIndex

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_PostgresMigratePgIndex interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PostgresMigratePgIndex) SetCrudManualInterface(crud ICrud_manual_PostgresMigratePgIndex) {
	Crud_manual_PostgresMigratePgIndex = crud

	return
}
