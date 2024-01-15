// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dbinteractors

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
	"mvp-student-project-management-system/services/orm/models"
)

func newProjectStatus(db *gorm.DB, opts ...gen.DOOption) projectStatus {
	_projectStatus := projectStatus{}

	_projectStatus.projectStatusDo.UseDB(db, opts...)
	_projectStatus.projectStatusDo.UseModel(&model.ProjectStatus{})

	tableName := _projectStatus.projectStatusDo.TableName()
	_projectStatus.ALL = field.NewAsterisk(tableName)
	_projectStatus.ID = field.NewInt32(tableName, "id")
	_projectStatus.Name = field.NewString(tableName, "name")

	_projectStatus.fillFieldMap()

	return _projectStatus
}

type projectStatus struct {
	projectStatusDo

	ALL  field.Asterisk
	ID   field.Int32
	Name field.String

	fieldMap map[string]field.Expr
}

func (p projectStatus) Table(newTableName string) *projectStatus {
	p.projectStatusDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p projectStatus) As(alias string) *projectStatus {
	p.projectStatusDo.DO = *(p.projectStatusDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *projectStatus) updateTableName(table string) *projectStatus {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.Name = field.NewString(table, "name")

	p.fillFieldMap()

	return p
}

func (p *projectStatus) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *projectStatus) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 2)
	p.fieldMap["id"] = p.ID
	p.fieldMap["name"] = p.Name
}

func (p projectStatus) clone(db *gorm.DB) projectStatus {
	p.projectStatusDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p projectStatus) replaceDB(db *gorm.DB) projectStatus {
	p.projectStatusDo.ReplaceDB(db)
	return p
}

type projectStatusDo struct{ gen.DO }

type IProjectStatusDo interface {
	gen.SubQuery
	Debug() IProjectStatusDo
	WithContext(ctx context.Context) IProjectStatusDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProjectStatusDo
	WriteDB() IProjectStatusDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProjectStatusDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProjectStatusDo
	Not(conds ...gen.Condition) IProjectStatusDo
	Or(conds ...gen.Condition) IProjectStatusDo
	Select(conds ...field.Expr) IProjectStatusDo
	Where(conds ...gen.Condition) IProjectStatusDo
	Order(conds ...field.Expr) IProjectStatusDo
	Distinct(cols ...field.Expr) IProjectStatusDo
	Omit(cols ...field.Expr) IProjectStatusDo
	Join(table schema.Tabler, on ...field.Expr) IProjectStatusDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProjectStatusDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProjectStatusDo
	Group(cols ...field.Expr) IProjectStatusDo
	Having(conds ...gen.Condition) IProjectStatusDo
	Limit(limit int) IProjectStatusDo
	Offset(offset int) IProjectStatusDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProjectStatusDo
	Unscoped() IProjectStatusDo
	Create(values ...*model.ProjectStatus) error
	CreateInBatches(values []*model.ProjectStatus, batchSize int) error
	Save(values ...*model.ProjectStatus) error
	First() (*model.ProjectStatus, error)
	Take() (*model.ProjectStatus, error)
	Last() (*model.ProjectStatus, error)
	Find() ([]*model.ProjectStatus, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProjectStatus, err error)
	FindInBatches(result *[]*model.ProjectStatus, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ProjectStatus) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProjectStatusDo
	Assign(attrs ...field.AssignExpr) IProjectStatusDo
	Joins(fields ...field.RelationField) IProjectStatusDo
	Preload(fields ...field.RelationField) IProjectStatusDo
	FirstOrInit() (*model.ProjectStatus, error)
	FirstOrCreate() (*model.ProjectStatus, error)
	FindByPage(offset int, limit int) (result []*model.ProjectStatus, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProjectStatusDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p projectStatusDo) Debug() IProjectStatusDo {
	return p.withDO(p.DO.Debug())
}

func (p projectStatusDo) WithContext(ctx context.Context) IProjectStatusDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p projectStatusDo) ReadDB() IProjectStatusDo {
	return p.Clauses(dbresolver.Read)
}

func (p projectStatusDo) WriteDB() IProjectStatusDo {
	return p.Clauses(dbresolver.Write)
}

func (p projectStatusDo) Session(config *gorm.Session) IProjectStatusDo {
	return p.withDO(p.DO.Session(config))
}

func (p projectStatusDo) Clauses(conds ...clause.Expression) IProjectStatusDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p projectStatusDo) Returning(value interface{}, columns ...string) IProjectStatusDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p projectStatusDo) Not(conds ...gen.Condition) IProjectStatusDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p projectStatusDo) Or(conds ...gen.Condition) IProjectStatusDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p projectStatusDo) Select(conds ...field.Expr) IProjectStatusDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p projectStatusDo) Where(conds ...gen.Condition) IProjectStatusDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p projectStatusDo) Order(conds ...field.Expr) IProjectStatusDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p projectStatusDo) Distinct(cols ...field.Expr) IProjectStatusDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p projectStatusDo) Omit(cols ...field.Expr) IProjectStatusDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p projectStatusDo) Join(table schema.Tabler, on ...field.Expr) IProjectStatusDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p projectStatusDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProjectStatusDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p projectStatusDo) RightJoin(table schema.Tabler, on ...field.Expr) IProjectStatusDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p projectStatusDo) Group(cols ...field.Expr) IProjectStatusDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p projectStatusDo) Having(conds ...gen.Condition) IProjectStatusDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p projectStatusDo) Limit(limit int) IProjectStatusDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p projectStatusDo) Offset(offset int) IProjectStatusDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p projectStatusDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProjectStatusDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p projectStatusDo) Unscoped() IProjectStatusDo {
	return p.withDO(p.DO.Unscoped())
}

func (p projectStatusDo) Create(values ...*model.ProjectStatus) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p projectStatusDo) CreateInBatches(values []*model.ProjectStatus, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p projectStatusDo) Save(values ...*model.ProjectStatus) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p projectStatusDo) First() (*model.ProjectStatus, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectStatus), nil
	}
}

func (p projectStatusDo) Take() (*model.ProjectStatus, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectStatus), nil
	}
}

func (p projectStatusDo) Last() (*model.ProjectStatus, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectStatus), nil
	}
}

func (p projectStatusDo) Find() ([]*model.ProjectStatus, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProjectStatus), err
}

func (p projectStatusDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProjectStatus, err error) {
	buf := make([]*model.ProjectStatus, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p projectStatusDo) FindInBatches(result *[]*model.ProjectStatus, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p projectStatusDo) Attrs(attrs ...field.AssignExpr) IProjectStatusDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p projectStatusDo) Assign(attrs ...field.AssignExpr) IProjectStatusDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p projectStatusDo) Joins(fields ...field.RelationField) IProjectStatusDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p projectStatusDo) Preload(fields ...field.RelationField) IProjectStatusDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p projectStatusDo) FirstOrInit() (*model.ProjectStatus, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectStatus), nil
	}
}

func (p projectStatusDo) FirstOrCreate() (*model.ProjectStatus, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectStatus), nil
	}
}

func (p projectStatusDo) FindByPage(offset int, limit int) (result []*model.ProjectStatus, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p projectStatusDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p projectStatusDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p projectStatusDo) Delete(models ...*model.ProjectStatus) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *projectStatusDo) withDO(do gen.Dao) *projectStatusDo {
	p.DO = *do.(*gen.DO)
	return p
}