// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"Nft-Go/user/internal/model"
)

func newNotice(db *gorm.DB, opts ...gen.DOOption) notice {
	_notice := notice{}

	_notice.noticeDo.UseDB(db, opts...)
	_notice.noticeDo.UseModel(&model.Notice{})

	tableName := _notice.noticeDo.TableName()
	_notice.ALL = field.NewAsterisk(tableName)
	_notice.Id = field.NewInt32(tableName, "id")
	_notice.Title = field.NewString(tableName, "title")
	_notice.Description = field.NewString(tableName, "description")
	_notice.PublishTime = field.NewString(tableName, "publish_time")
	_notice.UserAddress = field.NewString(tableName, "user_address")
	_notice.Type = field.NewInt(tableName, "type")

	_notice.fillFieldMap()

	return _notice
}

type notice struct {
	noticeDo

	ALL         field.Asterisk
	Id          field.Int32
	Title       field.String
	Description field.String
	PublishTime field.String
	UserAddress field.String
	Type        field.Int

	fieldMap map[string]field.Expr
}

func (n notice) Table(newTableName string) *notice {
	n.noticeDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n notice) As(alias string) *notice {
	n.noticeDo.DO = *(n.noticeDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *notice) updateTableName(table string) *notice {
	n.ALL = field.NewAsterisk(table)
	n.Id = field.NewInt32(table, "id")
	n.Title = field.NewString(table, "title")
	n.Description = field.NewString(table, "description")
	n.PublishTime = field.NewString(table, "publish_time")
	n.UserAddress = field.NewString(table, "user_address")
	n.Type = field.NewInt(table, "type")

	n.fillFieldMap()

	return n
}

func (n *notice) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *notice) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 6)
	n.fieldMap["id"] = n.Id
	n.fieldMap["title"] = n.Title
	n.fieldMap["description"] = n.Description
	n.fieldMap["publish_time"] = n.PublishTime
	n.fieldMap["user_address"] = n.UserAddress
	n.fieldMap["type"] = n.Type
}

func (n notice) clone(db *gorm.DB) notice {
	n.noticeDo.ReplaceConnPool(db.Statement.ConnPool)
	return n
}

func (n notice) replaceDB(db *gorm.DB) notice {
	n.noticeDo.ReplaceDB(db)
	return n
}

type noticeDo struct{ gen.DO }

type INoticeDo interface {
	gen.SubQuery
	Debug() INoticeDo
	WithContext(ctx context.Context) INoticeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() INoticeDo
	WriteDB() INoticeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) INoticeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) INoticeDo
	Not(conds ...gen.Condition) INoticeDo
	Or(conds ...gen.Condition) INoticeDo
	Select(conds ...field.Expr) INoticeDo
	Where(conds ...gen.Condition) INoticeDo
	Order(conds ...field.Expr) INoticeDo
	Distinct(cols ...field.Expr) INoticeDo
	Omit(cols ...field.Expr) INoticeDo
	Join(table schema.Tabler, on ...field.Expr) INoticeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) INoticeDo
	RightJoin(table schema.Tabler, on ...field.Expr) INoticeDo
	Group(cols ...field.Expr) INoticeDo
	Having(conds ...gen.Condition) INoticeDo
	Limit(limit int) INoticeDo
	Offset(offset int) INoticeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) INoticeDo
	Unscoped() INoticeDo
	Create(values ...*model.Notice) error
	CreateInBatches(values []*model.Notice, batchSize int) error
	Save(values ...*model.Notice) error
	First() (*model.Notice, error)
	Take() (*model.Notice, error)
	Last() (*model.Notice, error)
	Find() ([]*model.Notice, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Notice, err error)
	FindInBatches(result *[]*model.Notice, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Notice) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) INoticeDo
	Assign(attrs ...field.AssignExpr) INoticeDo
	Joins(fields ...field.RelationField) INoticeDo
	Preload(fields ...field.RelationField) INoticeDo
	FirstOrInit() (*model.Notice, error)
	FirstOrCreate() (*model.Notice, error)
	FindByPage(offset int, limit int) (result []*model.Notice, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) INoticeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (n noticeDo) Debug() INoticeDo {
	return n.withDO(n.DO.Debug())
}

func (n noticeDo) WithContext(ctx context.Context) INoticeDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n noticeDo) ReadDB() INoticeDo {
	return n.Clauses(dbresolver.Read)
}

func (n noticeDo) WriteDB() INoticeDo {
	return n.Clauses(dbresolver.Write)
}

func (n noticeDo) Session(config *gorm.Session) INoticeDo {
	return n.withDO(n.DO.Session(config))
}

func (n noticeDo) Clauses(conds ...clause.Expression) INoticeDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n noticeDo) Returning(value interface{}, columns ...string) INoticeDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n noticeDo) Not(conds ...gen.Condition) INoticeDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n noticeDo) Or(conds ...gen.Condition) INoticeDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n noticeDo) Select(conds ...field.Expr) INoticeDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n noticeDo) Where(conds ...gen.Condition) INoticeDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n noticeDo) Order(conds ...field.Expr) INoticeDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n noticeDo) Distinct(cols ...field.Expr) INoticeDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n noticeDo) Omit(cols ...field.Expr) INoticeDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n noticeDo) Join(table schema.Tabler, on ...field.Expr) INoticeDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n noticeDo) LeftJoin(table schema.Tabler, on ...field.Expr) INoticeDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n noticeDo) RightJoin(table schema.Tabler, on ...field.Expr) INoticeDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n noticeDo) Group(cols ...field.Expr) INoticeDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n noticeDo) Having(conds ...gen.Condition) INoticeDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n noticeDo) Limit(limit int) INoticeDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n noticeDo) Offset(offset int) INoticeDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n noticeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) INoticeDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n noticeDo) Unscoped() INoticeDo {
	return n.withDO(n.DO.Unscoped())
}

func (n noticeDo) Create(values ...*model.Notice) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n noticeDo) CreateInBatches(values []*model.Notice, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n noticeDo) Save(values ...*model.Notice) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n noticeDo) First() (*model.Notice, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Notice), nil
	}
}

func (n noticeDo) Take() (*model.Notice, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Notice), nil
	}
}

func (n noticeDo) Last() (*model.Notice, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Notice), nil
	}
}

func (n noticeDo) Find() ([]*model.Notice, error) {
	result, err := n.DO.Find()
	return result.([]*model.Notice), err
}

func (n noticeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Notice, err error) {
	buf := make([]*model.Notice, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n noticeDo) FindInBatches(result *[]*model.Notice, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n noticeDo) Attrs(attrs ...field.AssignExpr) INoticeDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n noticeDo) Assign(attrs ...field.AssignExpr) INoticeDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n noticeDo) Joins(fields ...field.RelationField) INoticeDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n noticeDo) Preload(fields ...field.RelationField) INoticeDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n noticeDo) FirstOrInit() (*model.Notice, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Notice), nil
	}
}

func (n noticeDo) FirstOrCreate() (*model.Notice, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Notice), nil
	}
}

func (n noticeDo) FindByPage(offset int, limit int) (result []*model.Notice, count int64, err error) {
	result, err = n.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = n.Offset(-1).Limit(-1).Count()
	return
}

func (n noticeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n noticeDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n noticeDo) Delete(models ...*model.Notice) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *noticeDo) withDO(do gen.Dao) *noticeDo {
	n.DO = *do.(*gen.DO)
	return n
}
