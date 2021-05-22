package sqlstore

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	mysqlmigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/mattermost/gorp"
	"github.com/mattermost/mattermost-server/v5/db/migrations"
	"github.com/zengqiang96/mattermostclone/model"
	"github.com/zengqiang96/mattermostclone/store"
)

type SqlStoreStores struct {
	channel store.ChannelStore
	post    store.PostStore
	session store.SessionStore
	team    store.TeamStore
	user    store.UserStore
}

type SqlStore struct {
	rrCounter int64
	srCounter int64
	stores    SqlStoreStores
	master    *gorp.DbMap
	Replicas  []*gorp.DbMap
	settings  *model.SqlSettings
}

func New(settings model.SqlSettings) *SqlStore {
	store := SqlStore{
		settings: &settings,
	}

	store.initConnection()
	err := store.migrate()
	if err != nil {
		fmt.Printf("数据库migrate失败 err: %v", err)
		os.Exit(1)
	}
	store.stores.user = newSqlUserStore(&store)

	return &store
}

func (ss *SqlStore) User() store.UserStore {
	return ss.stores.user
}

func (ss *SqlStore) Channel() store.ChannelStore {
	return ss.stores.channel
}

func (ss *SqlStore) Post() store.PostStore {
	return ss.stores.post
}

func (ss *SqlStore) Team() store.TeamStore {
	return ss.stores.team
}

func (ss *SqlStore) Session() store.SessionStore {
	return ss.stores.session
}

func (ss *SqlStore) GetMaster() *gorp.DbMap {
	return ss.master
}

func (ss *SqlStore) getQueryBuilder() sq.StatementBuilderType {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Question)
	return builder
}

func (ss *SqlStore) GetReplica() *gorp.DbMap {
	if len(ss.settings.DataSourceRelicas) == 0 {
		return ss.GetMaster()
	}
	rrNum := atomic.AddInt64(&ss.rrCounter, 1) % int64(len(ss.Replicas))
	return ss.Replicas[rrNum]
}

func (ss *SqlStore) GetAllConns() []*gorp.DbMap {
	all := make([]*gorp.DbMap, len(ss.Replicas)+1)
	copy(all, ss.Replicas)
	all[len(ss.Replicas)] = ss.master
	return all
}

func (ss *SqlStore) initConnection() {
	dataSource := *ss.settings.DataSource

	ss.master = setupConnection("master", dataSource, ss.settings)

	if len(ss.settings.DataSourceRelicas) > 0 {
		ss.Replicas = make([]*gorp.DbMap, len(ss.settings.DataSourceRelicas))
		for i, replica := range ss.settings.DataSourceRelicas {
			ss.Replicas[i] = setupConnection(fmt.Sprintf("replica-%v", i), replica, ss.settings)
		}
	}
}

func (ss *SqlStore) migrate() error {
	var driver database.Driver
	var err error

	dataSource, err := ss.appendMultipleStatementsFlag(*ss.settings.DataSource)
	if err != nil {
		return err
	}

	conn := setupConnection("migrations", dataSource, ss.settings)
	defer conn.Db.Close()

	driver, err = mysqlmigrate.WithInstance(conn.Db, &mysqlmigrate.Config{})
	if err != nil {
		return err
	}

	var assetNamesForDriver []string
	for _, assetName := range migrations.AssetNames() {
		if strings.HasPrefix(assetName, ss.DriverName()) {
			assetNamesForDriver = append(assetNamesForDriver, filepath.Base(assetName))
		}
	}

	source := bindata.Resource(assetNamesForDriver, func(name string) ([]byte, error) {
		return migrations.Asset(filepath.Join(ss.DriverName(), name))
	})

	sourceDriver, err := bindata.WithInstance(source)
	if err != nil {
		return err
	}

	migrations, err := migrate.NewWithInstance("go-bindata",
		sourceDriver,
		ss.DriverName(),
		driver)

	if err != nil {
		return err
	}
	defer migrations.Close()

	migrations.Up()

	return nil
}

func (ss *SqlStore) appendMultipleStatementsFlag(dataSource string) (string, error) {
	config, err := mysql.ParseDSN(dataSource)
	if err != nil {
		return "", err
	}

	config.Params["multiStatements"] = "true"
	return config.FormatDSN(), nil
}

func setupConnection(connType string, dataSource string, settings *model.SqlSettings) *gorp.DbMap {
	db, err := sql.Open(*settings.DriverName, dataSource)
	if err != nil {
		fmt.Printf("open SQL 连接失败 err: %v\n", err)
		os.Exit(1)
	}
	dbmap := &gorp.DbMap{Db: db}

	return dbmap
}

func (ss *SqlStore) DriverName() string {
	return *ss.settings.DriverName
}
