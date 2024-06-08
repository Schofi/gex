// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/luxun9527/gex/app/admin/api/internal/dao/admin/model"
)

func newCoin(db *gorm.DB, opts ...gen.DOOption) coin {
	_coin := coin{}

	_coin.coinDo.UseDB(db, opts...)
	_coin.coinDo.UseModel(&model.Coin{})

	tableName := _coin.coinDo.TableName()
	_coin.ALL = field.NewAsterisk(tableName)
	_coin.ID = field.NewUint32(tableName, "id")
	_coin.CoinID = field.NewInt32(tableName, "coin_id")
	_coin.CoinName = field.NewString(tableName, "coin_name")
	_coin.Prec = field.NewInt32(tableName, "prec")
	_coin.CreatedAt = field.NewUint32(tableName, "created_at")
	_coin.UpdatedAt = field.NewUint32(tableName, "updated_at")
	_coin.DeletedAt = field.NewUint32(tableName, "deleted_at")

	_coin.fillFieldMap()

	return _coin
}

type coin struct {
	coinDo coinDo

	ALL       field.Asterisk
	ID        field.Uint32 // 主键ID
	CoinID    field.Int32  // 币种ID
	CoinName  field.String // 币种名称
	Prec      field.Int32  // 币种精度，小数位保留多少
	CreatedAt field.Uint32 // 创建时间
	UpdatedAt field.Uint32 // 修改时间
	DeletedAt field.Uint32 // 删除时间

	fieldMap map[string]field.Expr
}

func (c coin) Table(newTableName string) *coin {
	c.coinDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c coin) As(alias string) *coin {
	c.coinDo.DO = *(c.coinDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *coin) updateTableName(table string) *coin {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint32(table, "id")
	c.CoinID = field.NewInt32(table, "coin_id")
	c.CoinName = field.NewString(table, "coin_name")
	c.Prec = field.NewInt32(table, "prec")
	c.CreatedAt = field.NewUint32(table, "created_at")
	c.UpdatedAt = field.NewUint32(table, "updated_at")
	c.DeletedAt = field.NewUint32(table, "deleted_at")

	c.fillFieldMap()

	return c
}

func (c *coin) WithContext(ctx context.Context) *coinDo { return c.coinDo.WithContext(ctx) }

func (c coin) TableName() string { return c.coinDo.TableName() }

func (c coin) Alias() string { return c.coinDo.Alias() }

func (c coin) Columns(cols ...field.Expr) gen.Columns { return c.coinDo.Columns(cols...) }

func (c *coin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *coin) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.ID
	c.fieldMap["coin_id"] = c.CoinID
	c.fieldMap["coin_name"] = c.CoinName
	c.fieldMap["prec"] = c.Prec
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
}

func (c coin) clone(db *gorm.DB) coin {
	c.coinDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c coin) replaceDB(db *gorm.DB) coin {
	c.coinDo.ReplaceDB(db)
	return c
}

type coinDo struct{ gen.DO }

func (c coinDo) Debug() *coinDo {
	return c.withDO(c.DO.Debug())
}

func (c coinDo) WithContext(ctx context.Context) *coinDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c coinDo) ReadDB() *coinDo {
	return c.Clauses(dbresolver.Read)
}

func (c coinDo) WriteDB() *coinDo {
	return c.Clauses(dbresolver.Write)
}

func (c coinDo) Session(config *gorm.Session) *coinDo {
	return c.withDO(c.DO.Session(config))
}

func (c coinDo) Clauses(conds ...clause.Expression) *coinDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c coinDo) Returning(value interface{}, columns ...string) *coinDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c coinDo) Not(conds ...gen.Condition) *coinDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c coinDo) Or(conds ...gen.Condition) *coinDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c coinDo) Select(conds ...field.Expr) *coinDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c coinDo) Where(conds ...gen.Condition) *coinDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c coinDo) Order(conds ...field.Expr) *coinDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c coinDo) Distinct(cols ...field.Expr) *coinDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c coinDo) Omit(cols ...field.Expr) *coinDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c coinDo) Join(table schema.Tabler, on ...field.Expr) *coinDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c coinDo) LeftJoin(table schema.Tabler, on ...field.Expr) *coinDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c coinDo) RightJoin(table schema.Tabler, on ...field.Expr) *coinDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c coinDo) Group(cols ...field.Expr) *coinDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c coinDo) Having(conds ...gen.Condition) *coinDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c coinDo) Limit(limit int) *coinDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c coinDo) Offset(offset int) *coinDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c coinDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *coinDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c coinDo) Unscoped() *coinDo {
	return c.withDO(c.DO.Unscoped())
}

func (c coinDo) Create(values ...*model.Coin) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c coinDo) CreateInBatches(values []*model.Coin, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c coinDo) Save(values ...*model.Coin) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c coinDo) First() (*model.Coin, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Coin), nil
	}
}

func (c coinDo) Take() (*model.Coin, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Coin), nil
	}
}

func (c coinDo) Last() (*model.Coin, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Coin), nil
	}
}

func (c coinDo) Find() ([]*model.Coin, error) {
	result, err := c.DO.Find()
	return result.([]*model.Coin), err
}

func (c coinDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Coin, err error) {
	buf := make([]*model.Coin, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c coinDo) FindInBatches(result *[]*model.Coin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c coinDo) Attrs(attrs ...field.AssignExpr) *coinDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c coinDo) Assign(attrs ...field.AssignExpr) *coinDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c coinDo) Joins(fields ...field.RelationField) *coinDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c coinDo) Preload(fields ...field.RelationField) *coinDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c coinDo) FirstOrInit() (*model.Coin, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Coin), nil
	}
}

func (c coinDo) FirstOrCreate() (*model.Coin, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Coin), nil
	}
}

func (c coinDo) FindByPage(offset int, limit int) (result []*model.Coin, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c coinDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c coinDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c coinDo) Delete(models ...*model.Coin) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *coinDo) withDO(do gen.Dao) *coinDo {
	c.DO = *do.(*gen.DO)
	return c
}