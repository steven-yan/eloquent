package models

// !!! DO NOT EDIT THIS FILE

import (
	"context"
	"database/sql"
	"github.com/iancoleman/strcase"
	"github.com/mylxsw/eloquent"
	"github.com/mylxsw/eloquent/query"
	"gopkg.in/guregu/null.v3"
	"time"
)

func init() {

}

// Organization is a Organization object
type Organization struct {
	original          *organizationOriginal
	organizationModel *OrganizationModel

	Id        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// SetModel set model for Organization
func (inst *Organization) SetModel(organizationModel *OrganizationModel) {
	inst.organizationModel = organizationModel
}

// organizationOriginal is an object which stores original Organization from database
type organizationOriginal struct {
	Id        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Staled identify whether the object has been modified
func (inst *Organization) Staled() bool {
	if inst.original == nil {
		inst.original = &organizationOriginal{}
	}

	if inst.Id != inst.original.Id {
		return true
	}
	if inst.Name != inst.original.Name {
		return true
	}
	if inst.CreatedAt != inst.original.CreatedAt {
		return true
	}
	if inst.UpdatedAt != inst.original.UpdatedAt {
		return true
	}

	return false
}

// StaledKV return all fields has been modified
func (inst *Organization) StaledKV() query.KV {
	kv := make(query.KV, 0)

	if inst.original == nil {
		inst.original = &organizationOriginal{}
	}

	if inst.Id != inst.original.Id {
		kv["id"] = inst.Id
	}
	if inst.Name != inst.original.Name {
		kv["name"] = inst.Name
	}
	if inst.CreatedAt != inst.original.CreatedAt {
		kv["created_at"] = inst.CreatedAt
	}
	if inst.UpdatedAt != inst.original.UpdatedAt {
		kv["updated_at"] = inst.UpdatedAt
	}

	return kv
}

// Save create a new model or update it
func (inst *Organization) Save() error {
	if inst.organizationModel == nil {
		return query.ErrModelNotSet
	}

	id, _, err := inst.organizationModel.SaveOrUpdate(*inst)
	if err != nil {
		return err
	}

	inst.Id = id
	return nil
}

// Delete remove a organization
func (inst *Organization) Delete() error {
	if inst.organizationModel == nil {
		return query.ErrModelNotSet
	}

	_, err := inst.organizationModel.DeleteById(inst.Id)
	if err != nil {
		return err
	}

	return nil
}

func (inst *Organization) Users() *OrganizationBelongsToManyUserRel {
	return &OrganizationBelongsToManyUserRel{
		source:     inst,
		pivotTable: "user_organization_ref",
		relModel:   NewUserModel(inst.organizationModel.GetDB()),
	}
}

type OrganizationBelongsToManyUserRel struct {
	source     *Organization
	pivotTable string
	relModel   *UserModel
}

func (rel *OrganizationBelongsToManyUserRel) Get(builders ...query.SQLBuilder) ([]User, error) {
	res, err := eloquent.DB(rel.relModel.GetDB()).Query(
		query.Builder().Table(rel.pivotTable).Select("user_id").Where("organization_id", rel.source.Id),
		func(row *sql.Rows) (interface{}, error) {
			var k interface{}
			if err := row.Scan(&k); err != nil {
				return nil, err
			}

			return k, nil
		},
	)

	if err != nil {
		return nil, err
	}

	resArr, _ := res.ToArray()
	return rel.relModel.Get(query.Builder().Merge(builders...).WhereIn("id", resArr...))
}

func (rel *OrganizationBelongsToManyUserRel) Count(builders ...query.SQLBuilder) (int64, error) {
	res, err := eloquent.DB(rel.relModel.GetDB()).Query(
		query.Builder().Table(rel.pivotTable).Select(query.Raw("COUNT(1) as c")).Where("organization_id", rel.source.Id),
		func(row *sql.Rows) (interface{}, error) {
			var k int64
			if err := row.Scan(&k); err != nil {
				return nil, err
			}

			return k, nil
		},
	)

	if err != nil {
		return 0, err
	}

	return res.Index(0).(int64), nil
}

func (rel *OrganizationBelongsToManyUserRel) Exists(builders ...query.SQLBuilder) (bool, error) {
	c, err := rel.Count(builders...)
	if err != nil {
		return false, err
	}

	return c > 0, nil
}

func (rel *OrganizationBelongsToManyUserRel) Attach(target User) error {
	_, err := eloquent.DB(rel.relModel.GetDB()).Insert(rel.pivotTable, query.KV{
		"user_id":         target.Id,
		"organization_id": rel.source.Id,
	})

	return err
}

func (rel *OrganizationBelongsToManyUserRel) Detach(target User) error {
	_, err := eloquent.DB(rel.relModel.GetDB()).
		Delete(eloquent.Build(rel.pivotTable).
			Where("user_id", target.Id).
			Where("organization_id", rel.source.Id))

	return err
}

func (rel *OrganizationBelongsToManyUserRel) DetachAll() error {
	_, err := eloquent.DB(rel.relModel.GetDB()).
		Delete(eloquent.Build(rel.pivotTable).
			Where("organization_id", rel.source.Id))
	return err
}

func (rel *OrganizationBelongsToManyUserRel) Create(target User, builders ...query.SQLBuilder) (int64, error) {
	targetId, err := rel.relModel.Save(target)
	if err != nil {
		return 0, err
	}

	target.Id = targetId

	err = rel.Attach(target)

	return targetId, err
}

type organizationScope struct {
	name  string
	apply func(builder query.Condition)
}

var organizationGlobalScopes = make([]organizationScope, 0)
var organizationLocalScopes = make([]organizationScope, 0)

// AddGlobalScopeForOrganization assign a global scope to a model
func AddGlobalScopeForOrganization(name string, apply func(builder query.Condition)) {
	organizationGlobalScopes = append(organizationGlobalScopes, organizationScope{name: name, apply: apply})
}

// AddLocalScopeForOrganization assign a local scope to a model
func AddLocalScopeForOrganization(name string, apply func(builder query.Condition)) {
	organizationLocalScopes = append(organizationLocalScopes, organizationScope{name: name, apply: apply})
}

func (m *OrganizationModel) applyScope() query.Condition {
	scopeCond := query.ConditionBuilder()
	for _, g := range organizationGlobalScopes {
		if m.globalScopeEnabled(g.name) {
			g.apply(scopeCond)
		}
	}

	for _, s := range organizationLocalScopes {
		if m.localScopeEnabled(s.name) {
			s.apply(scopeCond)
		}
	}

	return scopeCond
}

func (m *OrganizationModel) localScopeEnabled(name string) bool {
	for _, n := range m.includeLocalScopes {
		if name == n {
			return true
		}
	}

	return false
}

func (m *OrganizationModel) globalScopeEnabled(name string) bool {
	for _, n := range m.excludeGlobalScopes {
		if name == n {
			return false
		}
	}

	return true
}

type organizationWrap struct {
	Id        null.Int
	Name      null.String
	CreatedAt null.Time
	UpdatedAt null.Time
}

func (w organizationWrap) ToOrganization() Organization {
	return Organization{
		original: &organizationOriginal{
			Id:        w.Id.Int64,
			Name:      w.Name.String,
			CreatedAt: w.CreatedAt.Time,
			UpdatedAt: w.UpdatedAt.Time,
		},

		Id:        w.Id.Int64,
		Name:      w.Name.String,
		CreatedAt: w.CreatedAt.Time,
		UpdatedAt: w.UpdatedAt.Time,
	}
}

// OrganizationModel is a model which encapsulates the operations of the object
type OrganizationModel struct {
	db        *query.DatabaseWrap
	tableName string

	excludeGlobalScopes []string
	includeLocalScopes  []string

	query query.SQLBuilder
}

var organizationTableName = "wz_organization"

func SetOrganizationTable(tableName string) {
	organizationTableName = tableName
}

// NewOrganizationModel create a OrganizationModel
func NewOrganizationModel(db query.Database) *OrganizationModel {
	return &OrganizationModel{
		db:                  query.NewDatabaseWrap(db),
		tableName:           organizationTableName,
		excludeGlobalScopes: make([]string, 0),
		includeLocalScopes:  make([]string, 0),
		query:               query.Builder(),
	}
}

// GetDB return database instance
func (m *OrganizationModel) GetDB() query.Database {
	return m.db.GetDB()
}

func (m *OrganizationModel) clone() *OrganizationModel {
	return &OrganizationModel{
		db:                  m.db,
		tableName:           m.tableName,
		excludeGlobalScopes: append([]string{}, m.excludeGlobalScopes...),
		includeLocalScopes:  append([]string{}, m.includeLocalScopes...),
		query:               m.query,
	}
}

// WithoutGlobalScopes remove a global scope for given query
func (m *OrganizationModel) WithoutGlobalScopes(names ...string) *OrganizationModel {
	mc := m.clone()
	mc.excludeGlobalScopes = append(mc.excludeGlobalScopes, names...)

	return mc
}

// WithLocalScopes add a local scope for given query
func (m *OrganizationModel) WithLocalScopes(names ...string) *OrganizationModel {
	mc := m.clone()
	mc.includeLocalScopes = append(mc.includeLocalScopes, names...)

	return mc
}

// Query add query builder to model
func (m *OrganizationModel) Query(builder query.SQLBuilder) *OrganizationModel {
	mm := m.clone()
	mm.query = mm.query.Merge(builder)

	return mm
}

// Find retrieve a model by its primary key
func (m *OrganizationModel) Find(id int64) (Organization, error) {
	return m.First(m.query.Where("id", "=", id))
}

// Exists return whether the records exists for a given query
func (m *OrganizationModel) Exists(builders ...query.SQLBuilder) (bool, error) {
	count, err := m.Count(builders...)
	return count > 0, err
}

// Count return model count for a given query
func (m *OrganizationModel) Count(builders ...query.SQLBuilder) (int64, error) {
	sqlStr, params := m.query.
		Merge(builders...).
		Table(m.tableName).
		AppendCondition(m.applyScope()).
		ResolveCount()

	rows, err := m.db.QueryContext(context.Background(), sqlStr, params...)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	rows.Next()
	var res int64
	if err := rows.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (m *OrganizationModel) Paginate(page int64, perPage int64, builders ...query.SQLBuilder) ([]Organization, query.PaginateMeta, error) {
	if page <= 0 {
		page = 1
	}

	if perPage <= 0 {
		perPage = 15
	}

	meta := query.PaginateMeta{
		PerPage: perPage,
		Page:    page,
	}

	count, err := m.Count(builders...)
	if err != nil {
		return nil, meta, err
	}

	meta.Total = count
	meta.LastPage = count / perPage
	if count%perPage != 0 {
		meta.LastPage += 1
	}

	res, err := m.Get(append([]query.SQLBuilder{query.Builder().Limit(perPage).Offset((page - 1) * perPage)}, builders...)...)
	if err != nil {
		return res, meta, err
	}

	return res, meta, nil
}

// Get retrieve all results for given query
func (m *OrganizationModel) Get(builders ...query.SQLBuilder) ([]Organization, error) {
	b := m.query.Merge(builders...).Table(m.tableName).AppendCondition(m.applyScope())
	if len(b.GetFields()) == 0 {
		b = b.Select(
			"id",
			"name",
			"created_at",
			"updated_at",
		)
	}

	fields := b.GetFields()
	selectFields := make([]query.Expr, 0)

	for _, f := range fields {
		switch strcase.ToSnake(f.Value) {

		case "id":
			selectFields = append(selectFields, f)
		case "name":
			selectFields = append(selectFields, f)
		case "created_at":
			selectFields = append(selectFields, f)
		case "updated_at":
			selectFields = append(selectFields, f)
		}
	}

	var createScanVar = func(fields []query.Expr) (*organizationWrap, []interface{}) {
		var organizationVar organizationWrap
		scanFields := make([]interface{}, 0)

		for _, f := range fields {
			switch strcase.ToSnake(f.Value) {

			case "id":
				scanFields = append(scanFields, &organizationVar.Id)
			case "name":
				scanFields = append(scanFields, &organizationVar.Name)
			case "created_at":
				scanFields = append(scanFields, &organizationVar.CreatedAt)
			case "updated_at":
				scanFields = append(scanFields, &organizationVar.UpdatedAt)
			}
		}

		return &organizationVar, scanFields
	}

	sqlStr, params := b.Fields(selectFields...).ResolveQuery()

	rows, err := m.db.QueryContext(context.Background(), sqlStr, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	organizations := make([]Organization, 0)
	for rows.Next() {
		organizationVar, scanFields := createScanVar(fields)
		if err := rows.Scan(scanFields...); err != nil {
			return nil, err
		}

		organizationReal := organizationVar.ToOrganization()
		organizationReal.SetModel(m)
		organizations = append(organizations, organizationReal)
	}

	return organizations, nil
}

// First return first result for given query
func (m *OrganizationModel) First(builders ...query.SQLBuilder) (Organization, error) {
	res, err := m.Get(append(builders, query.Builder().Limit(1))...)
	if err != nil {
		return Organization{}, err
	}

	if len(res) == 0 {
		return Organization{}, query.ErrNoResult
	}

	return res[0], nil
}

// Create save a new organization to database
func (m *OrganizationModel) Create(kv query.KV) (int64, error) {

	if _, ok := kv["created_at"]; !ok {
		kv["created_at"] = time.Now()
	}

	if _, ok := kv["updated_at"]; !ok {
		kv["updated_at"] = time.Now()
	}

	sqlStr, params := m.query.Table(m.tableName).ResolveInsert(kv)

	res, err := m.db.ExecContext(context.Background(), sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// SaveAll save all organizations to database
func (m *OrganizationModel) SaveAll(organizations []Organization) ([]int64, error) {
	ids := make([]int64, 0)
	for _, organization := range organizations {
		id, err := m.Save(organization)
		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Save save a organization to database
func (m *OrganizationModel) Save(organization Organization) (int64, error) {
	return m.Create(organization.StaledKV())
}

// SaveOrUpdate save a new organization or update it when it has a id > 0
func (m *OrganizationModel) SaveOrUpdate(organization Organization) (id int64, updated bool, err error) {
	if organization.Id > 0 {
		_, _err := m.UpdateById(organization.Id, organization)
		return organization.Id, true, _err
	}

	_id, _err := m.Save(organization)
	return _id, false, _err
}

// UpdateFields update kv for a given query
func (m *OrganizationModel) UpdateFields(kv query.KV, builders ...query.SQLBuilder) (int64, error) {
	if len(kv) == 0 {
		return 0, nil
	}

	kv["updated_at"] = time.Now()

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).
		Table(m.tableName).
		ResolveUpdate(kv)

	res, err := m.db.ExecContext(context.Background(), sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Update update a model for given query
func (m *OrganizationModel) Update(organization Organization) (int64, error) {
	return m.UpdateFields(organization.StaledKV())
}

// UpdateById update a model by id
func (m *OrganizationModel) UpdateById(id int64, organization Organization) (int64, error) {
	return m.Query(query.Builder().Where("id", "=", id)).Update(organization)
}

// Delete remove a model
func (m *OrganizationModel) Delete(builders ...query.SQLBuilder) (int64, error) {

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).Table(m.tableName).ResolveDelete()

	res, err := m.db.ExecContext(context.Background(), sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}

// DeleteById remove a model by id
func (m *OrganizationModel) DeleteById(id int64) (int64, error) {
	return m.Query(query.Builder().Where("id", "=", id)).Delete()
}
