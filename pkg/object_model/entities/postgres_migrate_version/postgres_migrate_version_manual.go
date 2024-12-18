package postgres_migrate_version

// Crud_manual_PostgresMigrateVersion - объект контроллер crud операций
var Crud_manual_PostgresMigrateVersion ICrud_manual_PostgresMigrateVersion

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_PostgresMigrateVersion interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PostgresMigrateVersion) SetCrudManualInterface(crud ICrud_manual_PostgresMigrateVersion) {
	Crud_manual_PostgresMigrateVersion = crud

	return
}
