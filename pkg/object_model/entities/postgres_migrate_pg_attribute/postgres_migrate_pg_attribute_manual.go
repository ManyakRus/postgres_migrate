package postgres_migrate_pg_attribute

// Crud_manual_PostgresMigratePgAttribute - объект контроллер crud операций
var Crud_manual_PostgresMigratePgAttribute ICrud_manual_PostgresMigratePgAttribute

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_PostgresMigratePgAttribute interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PostgresMigratePgAttribute) SetCrudManualInterface(crud ICrud_manual_PostgresMigratePgAttribute) {
	Crud_manual_PostgresMigratePgAttribute = crud

	return
}
