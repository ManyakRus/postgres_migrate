package version

// Crud_manual_Version - объект контроллер crud операций
var Crud_manual_Version ICrud_manual_Version

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_Version interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Version) SetCrudManualInterface(crud ICrud_manual_Version) {
	Crud_manual_Version = crud

	return
}
