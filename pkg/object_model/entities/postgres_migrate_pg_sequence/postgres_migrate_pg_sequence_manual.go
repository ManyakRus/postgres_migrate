package postgres_migrate_pg_sequence

// Crud_manual_PostgresMigratePgSequence - объект контроллер crud операций
var Crud_manual_PostgresMigratePgSequence ICrud_manual_PostgresMigratePgSequence

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_PostgresMigratePgSequence interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PostgresMigratePgSequence) SetCrudManualInterface(crud ICrud_manual_PostgresMigratePgSequence) {
	Crud_manual_PostgresMigratePgSequence = crud

	return
}
