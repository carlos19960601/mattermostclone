package sqlstore

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zengqiang96/mattermostclone/model"
	"github.com/zengqiang96/mattermostclone/store"
)

type SqlUserStore struct {
	*SqlStore

	usersQuery sq.SelectBuilder
}

func newSqlUserStore(sqlStore *SqlStore) store.UserStore {
	us := SqlUserStore{
		SqlStore: sqlStore,
	}

	us.usersQuery = us.getQueryBuilder().
		Select("u.Id", "u.CreateAt", "u.UpdateAt", "u.DeleteAt", "u.Username", "u.Password").
		From("Users u")

	for _, db := range sqlStore.GetAllConns() {
		table := db.AddTableWithName(model.User{}, "Users").SetKeys(false, "Id")
		table.ColMap("Id").SetMaxSize(26)
		table.ColMap("Username").SetMaxSize(64).SetUnique(true)
		table.ColMap("Password").SetMaxSize(128)
	}
	return us
}

func (us SqlUserStore) Count(options model.UserCountOptions) (int64, error) {
	query := us.getQueryBuilder().Select("COUNT(DISTINCT u.Id)").From("Users as u")

	if !options.IncludeDeleted {
		query = query.Where("u.DeleteAt = 0")
	}

	queryString, args, err := query.ToSql()

	count, err := us.GetReplica().SelectInt(queryString, args...)
	if err != nil {
		return int64(0), fmt.Errorf("count user失败 err: %w", err)
	}
	return count, nil
}

func (us SqlUserStore) Get(ctx context.Context, id string) (*model.User, error) {
	query := us.usersQuery.Where("Id = ?", id)
	queryString, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("users构建sql失败 err: %w", err)
	}
	row := us.SqlStore.DbFromContext(ctx).Db.QueryRow(queryString, args...)

	var user model.User
	err = row.Scan(&user.Id, &user.CreateAt, &user.UpdateAt, &user.DeleteAt, &user.Username, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("根据id获取User失败 userId: %s, err: %w", id, err)
	}
	return &user, nil
}

func (us SqlUserStore) GetForLogin(loginId string) (*model.User, error) {
	query := us.usersQuery
	query = query.Where("Username = lower(?)", loginId)

	queryString, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("GetForLogin 构建sql失败 err: %w", err)
	}
	users := []*model.User{}
	if _, err := us.GetReplica().Select(&users, queryString, args...); err != nil {
		return nil, fmt.Errorf("查找用户失败 err: %w", err)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("未找到用户")
	}

	if len(users) > 1 {
		return nil, fmt.Errorf("查找到多个用户")
	}
	return users[0], nil
}

func (us SqlUserStore) Save(user *model.User) (*model.User, error) {
	if user.Id != "" {
		return nil, store.NewErrInvalidInput("User", "id", user.Id)
	}

	user.PreSave()

	if err := us.GetMaster().Insert(user); err != nil {

		return nil, fmt.Errorf("保存用户失败 userId: %s, err: %w", user.Id, err)
	}
	return user, nil
}

func (us SqlUserStore) IsEmpty(excludeBots bool) (bool, error) {
	var hasRow bool
	builder := us.getQueryBuilder().
		Select("1").
		Prefix("SELECT EXISTS (").
		From("Users")

	if excludeBots {
		builder = builder.LeftJoin("Bots ON Users.Id = Bots.UserId").Where("Bots.UserId IS NULL")
	}

	builder = builder.Suffix(")")

	query, args, err := builder.ToSql()
	if err != nil {
		return false, errors.Wrapf(err, "users_is_empty_to_sql")
	}

	if err = us.GetReplica().SelectOne(&hasRow, query, args...); err != nil {
		return false, errors.Wrap(err, "failed to check if table is empty")
	}
	return !hasRow, nil
}
