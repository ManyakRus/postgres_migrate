//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_postgres_migrate_pg_index

import (
	"context"
	"errors"
	"fmt"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_index"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gorm.io/gorm"
	"time"
)

// UpdateManyFields - изменяет несколько полей в базе данных
// MassNeedUpdateFields - список полей структуры golang для обновления
func (crud Crud_DB) UpdateManyFields(m *postgres_migrate_pg_index.PostgresMigratePgIndex, MassNeedUpdateFields []string) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = UpdateManyFields_ctx(ctx, db, m, MassNeedUpdateFields)
	return err
}

// UpdateManyFields_ctx - изменяет несколько полей в базе данных
// с учётом контекста и соединения к БД
// MassNeedUpdateFields - список полей структуры golang для обновления
func UpdateManyFields_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex, MassNeedUpdateFields []string) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` UpdateManyFields() error: ID=0`)
		return err
	}

	err = create_update_ctx(ctx, db, m, MassNeedUpdateFields)
	return err
}

// Update_Indcheckxmin - изменяет 1 поле Indcheckxmin в базе данных
func (crud Crud_DB) Update_Indcheckxmin(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indcheckxmin_ctx(ctx, db, m)
	return err
}

// Update_Indcheckxmin_ctx - изменяет 1 поле Indcheckxmin в базе данных
// с учётом контекста и соединения к БД
func Update_Indcheckxmin_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indcheckxmin() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indcheckxmin
	tx = db.Model(&m).Update("Indcheckxmin", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indcheckxmin() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indclass - изменяет 1 поле Indclass в базе данных
func (crud Crud_DB) Update_Indclass(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indclass_ctx(ctx, db, m)
	return err
}

// Update_Indclass_ctx - изменяет 1 поле Indclass в базе данных
// с учётом контекста и соединения к БД
func Update_Indclass_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indclass() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indclass
	tx = db.Model(&m).Update("Indclass", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indclass() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indcollation - изменяет 1 поле Indcollation в базе данных
func (crud Crud_DB) Update_Indcollation(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indcollation_ctx(ctx, db, m)
	return err
}

// Update_Indcollation_ctx - изменяет 1 поле Indcollation в базе данных
// с учётом контекста и соединения к БД
func Update_Indcollation_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indcollation() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indcollation
	tx = db.Model(&m).Update("Indcollation", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indcollation() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indexprs - изменяет 1 поле Indexprs в базе данных
func (crud Crud_DB) Update_Indexprs(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indexprs_ctx(ctx, db, m)
	return err
}

// Update_Indexprs_ctx - изменяет 1 поле Indexprs в базе данных
// с учётом контекста и соединения к БД
func Update_Indexprs_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indexprs() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indexprs
	tx = db.Model(&m).Update("Indexprs", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indexprs() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indexrelid - изменяет 1 поле Indexrelid в базе данных
func (crud Crud_DB) Update_Indexrelid(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indexrelid_ctx(ctx, db, m)
	return err
}

// Update_Indexrelid_ctx - изменяет 1 поле Indexrelid в базе данных
// с учётом контекста и соединения к БД
func Update_Indexrelid_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indexrelid() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indexrelid
	tx = db.Model(&m).Update("Indexrelid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indexrelid() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indimmediate - изменяет 1 поле Indimmediate в базе данных
func (crud Crud_DB) Update_Indimmediate(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indimmediate_ctx(ctx, db, m)
	return err
}

// Update_Indimmediate_ctx - изменяет 1 поле Indimmediate в базе данных
// с учётом контекста и соединения к БД
func Update_Indimmediate_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indimmediate() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indimmediate
	tx = db.Model(&m).Update("Indimmediate", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indimmediate() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indisclustered - изменяет 1 поле Indisclustered в базе данных
func (crud Crud_DB) Update_Indisclustered(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indisclustered_ctx(ctx, db, m)
	return err
}

// Update_Indisclustered_ctx - изменяет 1 поле Indisclustered в базе данных
// с учётом контекста и соединения к БД
func Update_Indisclustered_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indisclustered() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indisclustered
	tx = db.Model(&m).Update("Indisclustered", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indisclustered() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indisexclusion - изменяет 1 поле Indisexclusion в базе данных
func (crud Crud_DB) Update_Indisexclusion(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indisexclusion_ctx(ctx, db, m)
	return err
}

// Update_Indisexclusion_ctx - изменяет 1 поле Indisexclusion в базе данных
// с учётом контекста и соединения к БД
func Update_Indisexclusion_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indisexclusion() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indisexclusion
	tx = db.Model(&m).Update("Indisexclusion", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indisexclusion() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indislive - изменяет 1 поле Indislive в базе данных
func (crud Crud_DB) Update_Indislive(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indislive_ctx(ctx, db, m)
	return err
}

// Update_Indislive_ctx - изменяет 1 поле Indislive в базе данных
// с учётом контекста и соединения к БД
func Update_Indislive_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indislive() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indislive
	tx = db.Model(&m).Update("Indislive", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indislive() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indisprimary - изменяет 1 поле Indisprimary в базе данных
func (crud Crud_DB) Update_Indisprimary(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indisprimary_ctx(ctx, db, m)
	return err
}

// Update_Indisprimary_ctx - изменяет 1 поле Indisprimary в базе данных
// с учётом контекста и соединения к БД
func Update_Indisprimary_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indisprimary() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indisprimary
	tx = db.Model(&m).Update("Indisprimary", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indisprimary() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indisready - изменяет 1 поле Indisready в базе данных
func (crud Crud_DB) Update_Indisready(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indisready_ctx(ctx, db, m)
	return err
}

// Update_Indisready_ctx - изменяет 1 поле Indisready в базе данных
// с учётом контекста и соединения к БД
func Update_Indisready_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indisready() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indisready
	tx = db.Model(&m).Update("Indisready", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indisready() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indisreplident - изменяет 1 поле Indisreplident в базе данных
func (crud Crud_DB) Update_Indisreplident(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indisreplident_ctx(ctx, db, m)
	return err
}

// Update_Indisreplident_ctx - изменяет 1 поле Indisreplident в базе данных
// с учётом контекста и соединения к БД
func Update_Indisreplident_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indisreplident() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indisreplident
	tx = db.Model(&m).Update("Indisreplident", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indisreplident() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indisunique - изменяет 1 поле Indisunique в базе данных
func (crud Crud_DB) Update_Indisunique(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indisunique_ctx(ctx, db, m)
	return err
}

// Update_Indisunique_ctx - изменяет 1 поле Indisunique в базе данных
// с учётом контекста и соединения к БД
func Update_Indisunique_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indisunique() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indisunique
	tx = db.Model(&m).Update("Indisunique", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indisunique() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indisvalid - изменяет 1 поле Indisvalid в базе данных
func (crud Crud_DB) Update_Indisvalid(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indisvalid_ctx(ctx, db, m)
	return err
}

// Update_Indisvalid_ctx - изменяет 1 поле Indisvalid в базе данных
// с учётом контекста и соединения к БД
func Update_Indisvalid_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indisvalid() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indisvalid
	tx = db.Model(&m).Update("Indisvalid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indisvalid() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indkey - изменяет 1 поле Indkey в базе данных
func (crud Crud_DB) Update_Indkey(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indkey_ctx(ctx, db, m)
	return err
}

// Update_Indkey_ctx - изменяет 1 поле Indkey в базе данных
// с учётом контекста и соединения к БД
func Update_Indkey_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indkey() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indkey
	tx = db.Model(&m).Update("Indkey", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indkey() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indnatts - изменяет 1 поле Indnatts в базе данных
func (crud Crud_DB) Update_Indnatts(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indnatts_ctx(ctx, db, m)
	return err
}

// Update_Indnatts_ctx - изменяет 1 поле Indnatts в базе данных
// с учётом контекста и соединения к БД
func Update_Indnatts_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indnatts() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indnatts
	tx = db.Model(&m).Update("Indnatts", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indnatts() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indnkeyatts - изменяет 1 поле Indnkeyatts в базе данных
func (crud Crud_DB) Update_Indnkeyatts(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indnkeyatts_ctx(ctx, db, m)
	return err
}

// Update_Indnkeyatts_ctx - изменяет 1 поле Indnkeyatts в базе данных
// с учётом контекста и соединения к БД
func Update_Indnkeyatts_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indnkeyatts() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indnkeyatts
	tx = db.Model(&m).Update("Indnkeyatts", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indnkeyatts() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indoption - изменяет 1 поле Indoption в базе данных
func (crud Crud_DB) Update_Indoption(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indoption_ctx(ctx, db, m)
	return err
}

// Update_Indoption_ctx - изменяет 1 поле Indoption в базе данных
// с учётом контекста и соединения к БД
func Update_Indoption_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indoption() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indoption
	tx = db.Model(&m).Update("Indoption", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indoption() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indpred - изменяет 1 поле Indpred в базе данных
func (crud Crud_DB) Update_Indpred(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indpred_ctx(ctx, db, m)
	return err
}

// Update_Indpred_ctx - изменяет 1 поле Indpred в базе данных
// с учётом контекста и соединения к БД
func Update_Indpred_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indpred() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indpred
	tx = db.Model(&m).Update("Indpred", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indpred() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_Indrelid - изменяет 1 поле Indrelid в базе данных
func (crud Crud_DB) Update_Indrelid(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Indrelid_ctx(ctx, db, m)
	return err
}

// Update_Indrelid_ctx - изменяет 1 поле Indrelid в базе данных
// с учётом контекста и соединения к БД
func Update_Indrelid_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Indrelid() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Indrelid
	tx = db.Model(&m).Update("Indrelid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Indrelid() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}

// Update_VersionID - изменяет 1 поле VersionID в базе данных
func (crud Crud_DB) Update_VersionID(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_VersionID_ctx(ctx, db, m)
	return err
}

// Update_VersionID_ctx - изменяет 1 поле VersionID в базе данных
// с учётом контекста и соединения к БД
func Update_VersionID_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Indexrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_VersionID() error: Indexrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.VersionID
	tx = db.Model(&m).Update("VersionID", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_VersionID() id: %v, error: %w", m.Indexrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	return err
}
