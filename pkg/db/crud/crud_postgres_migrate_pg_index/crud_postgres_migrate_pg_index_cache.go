//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_postgres_migrate_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"context"
	"time"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gorm.io/gorm"
)

// cache - кэш с данными
var cache *expirable.LRU[string, postgres_migrate_pg_index.PostgresMigratePgIndex]

// CACHE_SIZE - количество элементов в кэше
const CACHE_SIZE = 0

// CACHE_EXPIRE_MINUTES - время жизни элемента в кэше
const CACHE_EXPIRE_MINUTES = 86400

// init - инициализация кэша
func init() {
	cache = expirable.NewLRU[string, postgres_migrate_pg_index.PostgresMigratePgIndex](CACHE_SIZE, nil, time.Minute*CACHE_EXPIRE_MINUTES)
}

// ReadFromCache - находит запись в кеше или в БД по ID
func (crud Crud_DB) ReadFromCache(Indexrelid int64, VersionID int64) (postgres_migrate_pg_index.PostgresMigratePgIndex, error) {
	var Otvet postgres_migrate_pg_index.PostgresMigratePgIndex
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	Otvet, err = ReadFromCache_ctx(ctx, db, Indexrelid, VersionID)

	return Otvet, err
}

// ReadFromCache_ctx - находит запись в кеше или в БД по ID
func ReadFromCache_ctx(ctx context.Context, db *gorm.DB, Indexrelid int64, VersionID int64) (postgres_migrate_pg_index.PostgresMigratePgIndex, error) {
	var Otvet postgres_migrate_pg_index.PostgresMigratePgIndex
	var err error

	// поищем сначала в кэше
	Identifier := (postgres_migrate_pg_index.StringIdentifier(Indexrelid, VersionID))
	Otvet, ok := cache.Get(Identifier)
	if ok {
		return Otvet, nil
	}

	// поищем в БД
	Otvet.Indexrelid = Indexrelid
	Otvet.VersionID = VersionID

	err = Read_ctx(ctx, db, &Otvet)
	if err == nil {
		cache.Add(Identifier, Otvet)
	}

	return Otvet, err
}
