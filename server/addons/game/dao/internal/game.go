// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GameDao is the data access object for table ue_game.
type GameDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns GameColumns // columns contains all the column names of Table for convenient usage.
}

// GameColumns defines and stores column names for table ue_game.
type GameColumns struct {
	GameId    string // 自动编号
	GameSku   string // 唯一值
	Name      string // 游戏全名
	Title     string // 游戏名称
	Status    string // 状态
	Sort      string // 排序
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// gameColumns holds the columns for table ue_game.
var gameColumns = GameColumns{
	GameId:    "game_id",
	GameSku:   "game_sku",
	Name:      "name",
	Title:     "title",
	Status:    "status",
	Sort:      "sort",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewGameDao creates and returns a new DAO object for table data access.
func NewGameDao() *GameDao {
	return &GameDao{
		group:   "bee",
		table:   "ue_game",
		columns: gameColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GameDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GameDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GameDao) Columns() GameColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GameDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GameDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GameDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
