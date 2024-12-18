package postgres_migrate_pg_class

// Crud_manual_PostgresMigratePgClass - объект контроллер crud операций
var Crud_manual_PostgresMigratePgClass ICrud_manual_PostgresMigratePgClass

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_PostgresMigratePgClass interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PostgresMigratePgClass) SetCrudManualInterface(crud ICrud_manual_PostgresMigratePgClass) {
	Crud_manual_PostgresMigratePgClass = crud

	return
}
