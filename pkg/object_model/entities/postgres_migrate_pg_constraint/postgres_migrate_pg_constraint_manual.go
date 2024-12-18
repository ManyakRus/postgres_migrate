package postgres_migrate_pg_constraint

// Crud_manual_PostgresMigratePgConstraint - объект контроллер crud операций
var Crud_manual_PostgresMigratePgConstraint ICrud_manual_PostgresMigratePgConstraint

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_PostgresMigratePgConstraint interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PostgresMigratePgConstraint) SetCrudManualInterface(crud ICrud_manual_PostgresMigratePgConstraint) {
	Crud_manual_PostgresMigratePgConstraint = crud

	return
}
