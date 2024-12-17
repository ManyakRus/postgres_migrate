package pg_constraint

// Crud_manual_PgConstraint - объект контроллер crud операций
var Crud_manual_PgConstraint ICrud_manual_PgConstraint

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_PgConstraint interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PgConstraint) SetCrudManualInterface(crud ICrud_manual_PgConstraint) {
	Crud_manual_PgConstraint = crud

	return
}
