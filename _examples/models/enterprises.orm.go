package models

// !!! DO NOT EDIT THIS FILE

import (
	"context"
	"encoding/json"
	"github.com/iancoleman/strcase"
	"github.com/mylxsw/coll"
	"github.com/mylxsw/eloquent/query"
	"gopkg.in/guregu/null.v3"
	"time"
)

func init() {

	// AddEnterpriseGlobalScope assign a global scope to a model for soft delete
	AddGlobalScopeForEnterprise("soft_delete", func(builder query.Condition) {
		builder.WhereNull("deleted_at")
	})

}

// Enterprise is a Enterprise object
type Enterprise struct {
	original        *enterpriseOriginal
	enterpriseModel *EnterpriseModel

	Id        null.Int
	Name      null.String
	Address   null.String
	Status    null.Int
	CreatedAt null.Time
	UpdatedAt null.Time
	DeletedAt null.Time
}

// As convert object to other type
// dst must be a pointer to struct
func (inst *Enterprise) As(dst interface{}) error {
	return coll.CopyProperties(inst, dst)
}

// SetModel set model for Enterprise
func (inst *Enterprise) SetModel(enterpriseModel *EnterpriseModel) {
	inst.enterpriseModel = enterpriseModel
}

// enterpriseOriginal is an object which stores original Enterprise from database
type enterpriseOriginal struct {
	Id        null.Int
	Name      null.String
	Address   null.String
	Status    null.Int
	CreatedAt null.Time
	UpdatedAt null.Time
	DeletedAt null.Time
}

// Staled identify whether the object has been modified
func (inst *Enterprise) Staled(onlyFields ...string) bool {
	if inst.original == nil {
		inst.original = &enterpriseOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			return true
		}
		if inst.Name != inst.original.Name {
			return true
		}
		if inst.Address != inst.original.Address {
			return true
		}
		if inst.Status != inst.original.Status {
			return true
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			return true
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			return true
		}
		if inst.DeletedAt != inst.original.DeletedAt {
			return true
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					return true
				}
			case "name":
				if inst.Name != inst.original.Name {
					return true
				}
			case "address":
				if inst.Address != inst.original.Address {
					return true
				}
			case "status":
				if inst.Status != inst.original.Status {
					return true
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					return true
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					return true
				}
			case "deleted_at":
				if inst.DeletedAt != inst.original.DeletedAt {
					return true
				}
			default:
			}
		}
	}

	return false
}

// StaledKV return all fields has been modified
func (inst *Enterprise) StaledKV(onlyFields ...string) query.KV {
	kv := make(query.KV, 0)

	if inst.original == nil {
		inst.original = &enterpriseOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			kv["id"] = inst.Id
		}
		if inst.Name != inst.original.Name {
			kv["name"] = inst.Name
		}
		if inst.Address != inst.original.Address {
			kv["address"] = inst.Address
		}
		if inst.Status != inst.original.Status {
			kv["status"] = inst.Status
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			kv["created_at"] = inst.CreatedAt
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			kv["updated_at"] = inst.UpdatedAt
		}
		if inst.DeletedAt != inst.original.DeletedAt {
			kv["deleted_at"] = inst.DeletedAt
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					kv["id"] = inst.Id
				}
			case "name":
				if inst.Name != inst.original.Name {
					kv["name"] = inst.Name
				}
			case "address":
				if inst.Address != inst.original.Address {
					kv["address"] = inst.Address
				}
			case "status":
				if inst.Status != inst.original.Status {
					kv["status"] = inst.Status
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					kv["created_at"] = inst.CreatedAt
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					kv["updated_at"] = inst.UpdatedAt
				}
			case "deleted_at":
				if inst.DeletedAt != inst.original.DeletedAt {
					kv["deleted_at"] = inst.DeletedAt
				}
			default:
			}
		}
	}

	return kv
}

// Save create a new model or update it
func (inst *Enterprise) Save() error {
	if inst.enterpriseModel == nil {
		return query.ErrModelNotSet
	}

	id, _, err := inst.enterpriseModel.SaveOrUpdate(*inst)
	if err != nil {
		return err
	}

	inst.Id = null.IntFrom(id)
	return nil
}

// Delete remove a enterprise
func (inst *Enterprise) Delete() error {
	if inst.enterpriseModel == nil {
		return query.ErrModelNotSet
	}

	_, err := inst.enterpriseModel.DeleteById(inst.Id.Int64)
	if err != nil {
		return err
	}

	return nil
}

// String convert instance to json string
func (inst *Enterprise) String() string {
	rs, _ := json.Marshal(inst)
	return string(rs)
}

func (inst *Enterprise) Users() *EnterpriseHasManyUserRel {
	return &EnterpriseHasManyUserRel{
		source:   inst,
		relModel: NewUserModel(inst.enterpriseModel.GetDB()),
	}
}

type EnterpriseHasManyUserRel struct {
	source   *Enterprise
	relModel *UserModel
}

func (rel *EnterpriseHasManyUserRel) Get(builders ...query.SQLBuilder) ([]User, error) {
	builder := query.Builder().Where("enterprise_id", rel.source.Id).Merge(builders...)

	return rel.relModel.Get(builder)
}

func (rel *EnterpriseHasManyUserRel) Count(builders ...query.SQLBuilder) (int64, error) {
	builder := query.Builder().Where("enterprise_id", rel.source.Id).Merge(builders...)

	return rel.relModel.Count(builder)
}

func (rel *EnterpriseHasManyUserRel) Exists(builders ...query.SQLBuilder) (bool, error) {
	builder := query.Builder().Where("enterprise_id", rel.source.Id).Merge(builders...)

	return rel.relModel.Exists(builder)
}

func (rel *EnterpriseHasManyUserRel) First(builders ...query.SQLBuilder) (User, error) {
	builder := query.Builder().Where("enterprise_id", rel.source.Id).Limit(1).Merge(builders...)
	return rel.relModel.First(builder)
}

func (rel *EnterpriseHasManyUserRel) Create(target User) (int64, error) {
	target.EnterpriseId = rel.source.Id
	return rel.relModel.Save(target)
}

type enterpriseScope struct {
	name  string
	apply func(builder query.Condition)
}

var enterpriseGlobalScopes = make([]enterpriseScope, 0)
var enterpriseLocalScopes = make([]enterpriseScope, 0)

// AddGlobalScopeForEnterprise assign a global scope to a model
func AddGlobalScopeForEnterprise(name string, apply func(builder query.Condition)) {
	enterpriseGlobalScopes = append(enterpriseGlobalScopes, enterpriseScope{name: name, apply: apply})
}

// AddLocalScopeForEnterprise assign a local scope to a model
func AddLocalScopeForEnterprise(name string, apply func(builder query.Condition)) {
	enterpriseLocalScopes = append(enterpriseLocalScopes, enterpriseScope{name: name, apply: apply})
}

func (m *EnterpriseModel) applyScope() query.Condition {
	scopeCond := query.ConditionBuilder()
	for _, g := range enterpriseGlobalScopes {
		if m.globalScopeEnabled(g.name) {
			g.apply(scopeCond)
		}
	}

	for _, s := range enterpriseLocalScopes {
		if m.localScopeEnabled(s.name) {
			s.apply(scopeCond)
		}
	}

	return scopeCond
}

func (m *EnterpriseModel) localScopeEnabled(name string) bool {
	for _, n := range m.includeLocalScopes {
		if name == n {
			return true
		}
	}

	return false
}

func (m *EnterpriseModel) globalScopeEnabled(name string) bool {
	for _, n := range m.excludeGlobalScopes {
		if name == n {
			return false
		}
	}

	return true
}

type EnterprisePlain struct {
	Id        int64
	Name      string
	Address   string
	Status    int8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (w EnterprisePlain) ToEnterprise(allows ...string) Enterprise {
	if len(allows) == 0 {
		return Enterprise{

			Id:        null.IntFrom(int64(w.Id)),
			Name:      null.StringFrom(w.Name),
			Address:   null.StringFrom(w.Address),
			Status:    null.IntFrom(int64(w.Status)),
			CreatedAt: null.TimeFrom(w.CreatedAt),
			UpdatedAt: null.TimeFrom(w.UpdatedAt),
			DeletedAt: null.TimeFrom(w.DeletedAt),
		}
	}

	res := Enterprise{}
	for _, al := range allows {
		switch strcase.ToSnake(al) {

		case "id":
			res.Id = null.IntFrom(int64(w.Id))
		case "name":
			res.Name = null.StringFrom(w.Name)
		case "address":
			res.Address = null.StringFrom(w.Address)
		case "status":
			res.Status = null.IntFrom(int64(w.Status))
		case "created_at":
			res.CreatedAt = null.TimeFrom(w.CreatedAt)
		case "updated_at":
			res.UpdatedAt = null.TimeFrom(w.UpdatedAt)
		case "deleted_at":
			res.DeletedAt = null.TimeFrom(w.DeletedAt)
		default:
		}
	}

	return res
}

// As convert object to other type
// dst must be a pointer to struct
func (w EnterprisePlain) As(dst interface{}) error {
	return coll.CopyProperties(w, dst)
}

func (w *Enterprise) ToEnterprisePlain() EnterprisePlain {
	return EnterprisePlain{

		Id:        w.Id.Int64,
		Name:      w.Name.String,
		Address:   w.Address.String,
		Status:    int8(w.Status.Int64),
		CreatedAt: w.CreatedAt.Time,
		UpdatedAt: w.UpdatedAt.Time,
		DeletedAt: w.DeletedAt.Time,
	}
}

// EnterpriseModel is a model which encapsulates the operations of the object
type EnterpriseModel struct {
	db        *query.DatabaseWrap
	tableName string

	excludeGlobalScopes []string
	includeLocalScopes  []string

	query query.SQLBuilder
}

var enterpriseTableName = "wz_enterprise"

const (
	EnterpriseFieldId        = "id"
	EnterpriseFieldName      = "name"
	EnterpriseFieldAddress   = "address"
	EnterpriseFieldStatus    = "status"
	EnterpriseFieldCreatedAt = "created_at"
	EnterpriseFieldUpdatedAt = "updated_at"
	EnterpriseFieldDeletedAt = "deleted_at"
)

// EnterpriseFields return all fields in Enterprise model
func EnterpriseFields() []string {
	return []string{
		"id",
		"name",
		"address",
		"status",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

func SetEnterpriseTable(tableName string) {
	enterpriseTableName = tableName
}

// NewEnterpriseModel create a EnterpriseModel
func NewEnterpriseModel(db query.Database) *EnterpriseModel {
	return &EnterpriseModel{
		db:                  query.NewDatabaseWrap(db),
		tableName:           enterpriseTableName,
		excludeGlobalScopes: make([]string, 0),
		includeLocalScopes:  make([]string, 0),
		query:               query.Builder(),
	}
}

// GetDB return database instance
func (m *EnterpriseModel) GetDB() query.Database {
	return m.db.GetDB()
}

// WithTrashed force soft deleted models to appear in a result set
func (m *EnterpriseModel) WithTrashed() *EnterpriseModel {
	return m.WithoutGlobalScopes("soft_delete")
}

func (m *EnterpriseModel) clone() *EnterpriseModel {
	return &EnterpriseModel{
		db:                  m.db,
		tableName:           m.tableName,
		excludeGlobalScopes: append([]string{}, m.excludeGlobalScopes...),
		includeLocalScopes:  append([]string{}, m.includeLocalScopes...),
		query:               m.query,
	}
}

// WithoutGlobalScopes remove a global scope for given query
func (m *EnterpriseModel) WithoutGlobalScopes(names ...string) *EnterpriseModel {
	mc := m.clone()
	mc.excludeGlobalScopes = append(mc.excludeGlobalScopes, names...)

	return mc
}

// WithLocalScopes add a local scope for given query
func (m *EnterpriseModel) WithLocalScopes(names ...string) *EnterpriseModel {
	mc := m.clone()
	mc.includeLocalScopes = append(mc.includeLocalScopes, names...)

	return mc
}

// Condition add query builder to model
func (m *EnterpriseModel) Condition(builder query.SQLBuilder) *EnterpriseModel {
	mm := m.clone()
	mm.query = mm.query.Merge(builder)

	return mm
}

// Find retrieve a model by its primary key
func (m *EnterpriseModel) Find(id int64) (Enterprise, error) {
	return m.First(m.query.Where("id", "=", id))
}

// Exists return whether the records exists for a given query
func (m *EnterpriseModel) Exists(builders ...query.SQLBuilder) (bool, error) {
	count, err := m.Count(builders...)
	return count > 0, err
}

// Count return model count for a given query
func (m *EnterpriseModel) Count(builders ...query.SQLBuilder) (int64, error) {
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

func (m *EnterpriseModel) Paginate(page int64, perPage int64, builders ...query.SQLBuilder) ([]Enterprise, query.PaginateMeta, error) {
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
func (m *EnterpriseModel) Get(builders ...query.SQLBuilder) ([]Enterprise, error) {
	b := m.query.Merge(builders...).Table(m.tableName).AppendCondition(m.applyScope())
	if len(b.GetFields()) == 0 {
		b = b.Select(
			"id",
			"name",
			"address",
			"status",
			"created_at",
			"updated_at",
			"deleted_at",
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
		case "address":
			selectFields = append(selectFields, f)
		case "status":
			selectFields = append(selectFields, f)
		case "created_at":
			selectFields = append(selectFields, f)
		case "updated_at":
			selectFields = append(selectFields, f)
		case "deleted_at":
			selectFields = append(selectFields, f)
		}
	}

	var createScanVar = func(fields []query.Expr) (*Enterprise, []interface{}) {
		var enterpriseVar Enterprise
		scanFields := make([]interface{}, 0)

		for _, f := range fields {
			switch strcase.ToSnake(f.Value) {

			case "id":
				scanFields = append(scanFields, &enterpriseVar.Id)
			case "name":
				scanFields = append(scanFields, &enterpriseVar.Name)
			case "address":
				scanFields = append(scanFields, &enterpriseVar.Address)
			case "status":
				scanFields = append(scanFields, &enterpriseVar.Status)
			case "created_at":
				scanFields = append(scanFields, &enterpriseVar.CreatedAt)
			case "updated_at":
				scanFields = append(scanFields, &enterpriseVar.UpdatedAt)
			case "deleted_at":
				scanFields = append(scanFields, &enterpriseVar.DeletedAt)
			}
		}

		return &enterpriseVar, scanFields
	}

	sqlStr, params := b.Fields(selectFields...).ResolveQuery()

	rows, err := m.db.QueryContext(context.Background(), sqlStr, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	enterprises := make([]Enterprise, 0)
	for rows.Next() {
		enterpriseReal, scanFields := createScanVar(fields)
		if err := rows.Scan(scanFields...); err != nil {
			return nil, err
		}

		enterpriseReal.SetModel(m)
		enterprises = append(enterprises, *enterpriseReal)
	}

	return enterprises, nil
}

// First return first result for given query
func (m *EnterpriseModel) First(builders ...query.SQLBuilder) (Enterprise, error) {
	res, err := m.Get(append(builders, query.Builder().Limit(1))...)
	if err != nil {
		return Enterprise{}, err
	}

	if len(res) == 0 {
		return Enterprise{}, query.ErrNoResult
	}

	return res[0], nil
}

// Create save a new enterprise to database
func (m *EnterpriseModel) Create(kv query.KV) (int64, error) {

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

// SaveAll save all enterprises to database
func (m *EnterpriseModel) SaveAll(enterprises []Enterprise) ([]int64, error) {
	ids := make([]int64, 0)
	for _, enterprise := range enterprises {
		id, err := m.Save(enterprise)
		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Save save a enterprise to database
func (m *EnterpriseModel) Save(enterprise Enterprise, onlyFields ...string) (int64, error) {
	return m.Create(enterprise.StaledKV(onlyFields...))
}

// SaveOrUpdate save a new enterprise or update it when it has a id > 0
func (m *EnterpriseModel) SaveOrUpdate(enterprise Enterprise, onlyFields ...string) (id int64, updated bool, err error) {
	if enterprise.Id.Int64 > 0 {
		_, _err := m.UpdateById(enterprise.Id.Int64, enterprise, onlyFields...)
		return enterprise.Id.Int64, true, _err
	}

	_id, _err := m.Save(enterprise, onlyFields...)
	return _id, false, _err
}

// UpdateFields update kv for a given query
func (m *EnterpriseModel) UpdateFields(kv query.KV, builders ...query.SQLBuilder) (int64, error) {
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
func (m *EnterpriseModel) Update(enterprise Enterprise, builders ...query.SQLBuilder) (int64, error) {
	return m.UpdateFields(enterprise.StaledKV(), builders...)
}

// UpdatePart update a model for given query
func (m *EnterpriseModel) UpdatePart(enterprise Enterprise, onlyFields ...string) (int64, error) {
	return m.UpdateFields(enterprise.StaledKV(onlyFields...))
}

// UpdateById update a model by id
func (m *EnterpriseModel) UpdateById(id int64, enterprise Enterprise, onlyFields ...string) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).UpdateFields(enterprise.StaledKV(onlyFields...), builders...)
}

// ForceDelete permanently remove a soft deleted model from the database
func (m *EnterpriseModel) ForceDelete(builders ...query.SQLBuilder) (int64, error) {
	m2 := m.WithTrashed()

	sqlStr, params := m2.query.Merge(builders...).AppendCondition(m2.applyScope()).Table(m2.tableName).ResolveDelete()

	res, err := m2.db.ExecContext(context.Background(), sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// ForceDeleteById permanently remove a soft deleted model from the database by id
func (m *EnterpriseModel) ForceDeleteById(id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).ForceDelete()
}

// Restore restore a soft deleted model into an active state
func (m *EnterpriseModel) Restore(builders ...query.SQLBuilder) (int64, error) {
	m2 := m.WithTrashed()
	return m2.UpdateFields(query.KV{
		"deleted_at": nil,
	}, builders...)
}

// RestoreById restore a soft deleted model into an active state by id
func (m *EnterpriseModel) RestoreById(id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Restore()
}

// Delete remove a model
func (m *EnterpriseModel) Delete(builders ...query.SQLBuilder) (int64, error) {

	return m.UpdateFields(query.KV{
		"deleted_at": time.Now(),
	}, builders...)

}

// DeleteById remove a model by id
func (m *EnterpriseModel) DeleteById(id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Delete()
}
