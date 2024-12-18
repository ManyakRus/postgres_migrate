package postgres_migrate_pg_description

// Crud_manual_PostgresMigratePgDescription - объект контроллер crud операций
var Crud_manual_PostgresMigratePgDescription ICrud_manual_PostgresMigratePgDescription

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_PostgresMigratePgDescription interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PostgresMigratePgDescription) SetCrudManualInterface(crud ICrud_manual_PostgresMigratePgDescription) {
	Crud_manual_PostgresMigratePgDescription = crud

	return
}
