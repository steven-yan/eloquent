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

}

// Role is a Role object
type Role struct {
	original  *roleOriginal
	roleModel *RoleModel

	Name        null.String
	Description null.String
	Id          null.Int
	CreatedAt   null.Time
	UpdatedAt   null.Time
}

// As convert object to other type
// dst must be a pointer to struct
func (inst *Role) As(dst interface{}) error {
	return coll.CopyProperties(inst, dst)
}

// SetModel set model for Role
func (inst *Role) SetModel(roleModel *RoleModel) {
	inst.roleModel = roleModel
}

// roleOriginal is an object which stores original Role from database
type roleOriginal struct {
	Name        null.String
	Description null.String
	Id          null.Int
	CreatedAt   null.Time
	UpdatedAt   null.Time
}

// Staled identify whether the object has been modified
func (inst *Role) Staled(onlyFields ...string) bool {
	if inst.original == nil {
		inst.original = &roleOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Name != inst.original.Name {
			return true
		}
		if inst.Description != inst.original.Description {
			return true
		}
		if inst.Id != inst.original.Id {
			return true
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			return true
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			return true
		}
	} else {
		switch strcase.ToSnake(f) {

		case "name":
			if inst.Name != inst.original.Name {
				return true
			}
		case "description":
			if inst.Description != inst.original.Description {
				return true
			}
		case "id":
			if inst.Id != inst.original.Id {
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
		default:
		}
	}

	return false
}

// StaledKV return all fields has been modified
func (inst *Role) StaledKV(onlyFields ...string) query.KV {
	kv := make(query.KV, 0)

	if inst.original == nil {
		inst.original = &roleOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Name != inst.original.Name {
			kv["name"] = inst.Name
		}
		if inst.Description != inst.original.Description {
			kv["description"] = inst.Description
		}
		if inst.Id != inst.original.Id {
			kv["id"] = inst.Id
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			kv["created_at"] = inst.CreatedAt
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			kv["updated_at"] = inst.UpdatedAt
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "name":
				if inst.Name != inst.original.Name {
					kv["name"] = inst.Name
				}
			case "description":
				if inst.Description != inst.original.Description {
					kv["description"] = inst.Description
				}
			case "id":
				if inst.Id != inst.original.Id {
					kv["id"] = inst.Id
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					kv["created_at"] = inst.CreatedAt
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					kv["updated_at"] = inst.UpdatedAt
				}
			default:
			}
		}
	}

	return kv
}

// Save create a new model or update it
func (inst *Role) Save() error {
	if inst.roleModel == nil {
		return query.ErrModelNotSet
	}

	id, _, err := inst.roleModel.SaveOrUpdate(*inst)
	if err != nil {
		return err
	}

	inst.Id = null.IntFrom(id)
	return nil
}

// Delete remove a Role
func (inst *Role) Delete() error {
	if inst.roleModel == nil {
		return query.ErrModelNotSet
	}

	_, err := inst.roleModel.DeleteById(inst.Id.Int64)
	if err != nil {
		return err
	}

	return nil
}

// String convert instance to json string
func (inst *Role) String() string {
	rs, _ := json.Marshal(inst)
	return string(rs)
}

func (inst *Role) Users() *RoleHasManyUserRel {
	return &RoleHasManyUserRel{
		source:   inst,
		relModel: NewUserModel(inst.roleModel.GetDB()),
	}
}

type RoleHasManyUserRel struct {
	source   *Role
	relModel *UserModel
}

func (rel *RoleHasManyUserRel) Get(builders ...query.SQLBuilder) ([]User, error) {
	builder := query.Builder().Where("role_id", rel.source.Id).Merge(builders...)

	return rel.relModel.Get(builder)
}

func (rel *RoleHasManyUserRel) Count(builders ...query.SQLBuilder) (int64, error) {
	builder := query.Builder().Where("role_id", rel.source.Id).Merge(builders...)

	return rel.relModel.Count(builder)
}

func (rel *RoleHasManyUserRel) Exists(builders ...query.SQLBuilder) (bool, error) {
	builder := query.Builder().Where("role_id", rel.source.Id).Merge(builders...)

	return rel.relModel.Exists(builder)
}

func (rel *RoleHasManyUserRel) First(builders ...query.SQLBuilder) (User, error) {
	builder := query.Builder().Where("role_id", rel.source.Id).Limit(1).Merge(builders...)
	return rel.relModel.First(builder)
}

func (rel *RoleHasManyUserRel) Create(target User) (int64, error) {
	target.RoleId = rel.source.Id
	return rel.relModel.Save(target)
}

type roleScope struct {
	name  string
	apply func(builder query.Condition)
}

var roleGlobalScopes = make([]roleScope, 0)
var roleLocalScopes = make([]roleScope, 0)

// AddGlobalScopeForRole assign a global scope to a model
func AddGlobalScopeForRole(name string, apply func(builder query.Condition)) {
	roleGlobalScopes = append(roleGlobalScopes, roleScope{name: name, apply: apply})
}

// AddLocalScopeForRole assign a local scope to a model
func AddLocalScopeForRole(name string, apply func(builder query.Condition)) {
	roleLocalScopes = append(roleLocalScopes, roleScope{name: name, apply: apply})
}

func (m *RoleModel) applyScope() query.Condition {
	scopeCond := query.ConditionBuilder()
	for _, g := range roleGlobalScopes {
		if m.globalScopeEnabled(g.name) {
			g.apply(scopeCond)
		}
	}

	for _, s := range roleLocalScopes {
		if m.localScopeEnabled(s.name) {
			s.apply(scopeCond)
		}
	}

	return scopeCond
}

func (m *RoleModel) localScopeEnabled(name string) bool {
	for _, n := range m.includeLocalScopes {
		if name == n {
			return true
		}
	}

	return false
}

func (m *RoleModel) globalScopeEnabled(name string) bool {
	for _, n := range m.excludeGlobalScopes {
		if name == n {
			return false
		}
	}

	return true
}

type RolePlain struct {
	Name        string
	Description string
	Id          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (w RolePlain) ToRole(allows ...string) Role {
	if len(allows) == 0 {
		return Role{

			Name:        null.StringFrom(w.Name),
			Description: null.StringFrom(w.Description),
			Id:          null.IntFrom(int64(w.Id)),
			CreatedAt:   null.TimeFrom(w.CreatedAt),
			UpdatedAt:   null.TimeFrom(w.UpdatedAt),
		}
	}

	res := Role{}
	for _, al := range allows {
		switch strcase.ToSnake(al) {

		case "name":
			res.Name = null.StringFrom(w.Name)
		case "description":
			res.Description = null.StringFrom(w.Description)
		case "id":
			res.Id = null.IntFrom(int64(w.Id))
		case "created_at":
			res.CreatedAt = null.TimeFrom(w.CreatedAt)
		case "updated_at":
			res.UpdatedAt = null.TimeFrom(w.UpdatedAt)
		default:
		}
	}

	return res
}

// As convert object to other type
// dst must be a pointer to struct
func (w RolePlain) As(dst interface{}) error {
	return coll.CopyProperties(w, dst)
}

func (w *Role) ToRolePlain() RolePlain {
	return RolePlain{

		Name:        w.Name.String,
		Description: w.Description.String,
		Id:          w.Id.Int64,
		CreatedAt:   w.CreatedAt.Time,
		UpdatedAt:   w.UpdatedAt.Time,
	}
}

// RoleModel is a model which encapsulates the operations of the object
type RoleModel struct {
	db        *query.DatabaseWrap
	tableName string

	excludeGlobalScopes []string
	includeLocalScopes  []string

	query query.SQLBuilder
}

var roleTableName = "wz_role"

const (
	RoleFieldName        = "name"
	RoleFieldDescription = "description"
	RoleFieldId          = "id"
	RoleFieldCreatedAt   = "created_at"
	RoleFieldUpdatedAt   = "updated_at"
)

const RoleFields = []string{
	"name",
	"description",
	"id",
	"created_at",
	"updated_at",
}

func SetRoleTable(tableName string) {
	roleTableName = tableName
}

// NewRoleModel create a RoleModel
func NewRoleModel(db query.Database) *RoleModel {
	return &RoleModel{
		db:                  query.NewDatabaseWrap(db),
		tableName:           roleTableName,
		excludeGlobalScopes: make([]string, 0),
		includeLocalScopes:  make([]string, 0),
		query:               query.Builder(),
	}
}

// GetDB return database instance
func (m *RoleModel) GetDB() query.Database {
	return m.db.GetDB()
}

func (m *RoleModel) clone() *RoleModel {
	return &RoleModel{
		db:                  m.db,
		tableName:           m.tableName,
		excludeGlobalScopes: append([]string{}, m.excludeGlobalScopes...),
		includeLocalScopes:  append([]string{}, m.includeLocalScopes...),
		query:               m.query,
	}
}

// WithoutGlobalScopes remove a global scope for given query
func (m *RoleModel) WithoutGlobalScopes(names ...string) *RoleModel {
	mc := m.clone()
	mc.excludeGlobalScopes = append(mc.excludeGlobalScopes, names...)

	return mc
}

// WithLocalScopes add a local scope for given query
func (m *RoleModel) WithLocalScopes(names ...string) *RoleModel {
	mc := m.clone()
	mc.includeLocalScopes = append(mc.includeLocalScopes, names...)

	return mc
}

// Condition add query builder to model
func (m *RoleModel) Condition(builder query.SQLBuilder) *RoleModel {
	mm := m.clone()
	mm.query = mm.query.Merge(builder)

	return mm
}

// Find retrieve a model by its primary key
func (m *RoleModel) Find(id int64) (Role, error) {
	return m.First(m.query.Where("id", "=", id))
}

// Exists return whether the records exists for a given query
func (m *RoleModel) Exists(builders ...query.SQLBuilder) (bool, error) {
	count, err := m.Count(builders...)
	return count > 0, err
}

// Count return model count for a given query
func (m *RoleModel) Count(builders ...query.SQLBuilder) (int64, error) {
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

func (m *RoleModel) Paginate(page int64, perPage int64, builders ...query.SQLBuilder) ([]Role, query.PaginateMeta, error) {
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
func (m *RoleModel) Get(builders ...query.SQLBuilder) ([]Role, error) {
	b := m.query.Merge(builders...).Table(m.tableName).AppendCondition(m.applyScope())
	if len(b.GetFields()) == 0 {
		b = b.Select(
			"name",
			"description",
			"id",
			"created_at",
			"updated_at",
		)
	}

	fields := b.GetFields()
	selectFields := make([]query.Expr, 0)

	for _, f := range fields {
		switch strcase.ToSnake(f.Value) {

		case "name":
			selectFields = append(selectFields, f)
		case "description":
			selectFields = append(selectFields, f)
		case "id":
			selectFields = append(selectFields, f)
		case "created_at":
			selectFields = append(selectFields, f)
		case "updated_at":
			selectFields = append(selectFields, f)
		}
	}

	var createScanVar = func(fields []query.Expr) (*Role, []interface{}) {
		var roleVar Role
		scanFields := make([]interface{}, 0)

		for _, f := range fields {
			switch strcase.ToSnake(f.Value) {

			case "name":
				scanFields = append(scanFields, &roleVar.Name)
			case "description":
				scanFields = append(scanFields, &roleVar.Description)
			case "id":
				scanFields = append(scanFields, &roleVar.Id)
			case "created_at":
				scanFields = append(scanFields, &roleVar.CreatedAt)
			case "updated_at":
				scanFields = append(scanFields, &roleVar.UpdatedAt)
			}
		}

		return &roleVar, scanFields
	}

	sqlStr, params := b.Fields(selectFields...).ResolveQuery()

	rows, err := m.db.QueryContext(context.Background(), sqlStr, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	roles := make([]Role, 0)
	for rows.Next() {
		roleReal, scanFields := createScanVar(fields)
		if err := rows.Scan(scanFields...); err != nil {
			return nil, err
		}

		roleReal.SetModel(m)
		roles = append(roles, *roleReal)
	}

	return roles, nil
}

// First return first result for given query
func (m *RoleModel) First(builders ...query.SQLBuilder) (Role, error) {
	res, err := m.Get(append(builders, query.Builder().Limit(1))...)
	if err != nil {
		return Role{}, err
	}

	if len(res) == 0 {
		return Role{}, query.ErrNoResult
	}

	return res[0], nil
}

// Create save a new Role to database
func (m *RoleModel) Create(kv query.KV) (int64, error) {

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

// SaveAll save all Roles to database
func (m *RoleModel) SaveAll(roles []Role) ([]int64, error) {
	ids := make([]int64, 0)
	for _, role := range roles {
		id, err := m.Save(role)
		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Save save a Role to database
func (m *RoleModel) Save(role Role, onlyFields ...string) (int64, error) {
	return m.Create(role.StaledKV(onlyFields...))
}

// SaveOrUpdate save a new Role or update it when it has a id > 0
func (m *RoleModel) SaveOrUpdate(role Role, onlyFields ...string) (id int64, updated bool, err error) {
	if role.Id.Int64 > 0 {
		_, _err := m.UpdateById(role.Id.Int64, role, onlyFields...)
		return role.Id.Int64, true, _err
	}

	_id, _err := m.Save(role, onlyFields...)
	return _id, false, _err
}

// UpdateFields update kv for a given query
func (m *RoleModel) UpdateFields(kv query.KV, builders ...query.SQLBuilder) (int64, error) {
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
func (m *RoleModel) Update(role Role, builders ...query.SQLBuilder) (int64, error) {
	return m.UpdateFields(role.StaledKV(), builders...)
}

// UpdatePart update a model for given query
func (m *RoleModel) UpdatePart(role Role, onlyFields []string, builders ...query.SQLBuilder) (int64, error) {
	return m.UpdateFields(role.StaledKV(onlyFields...), builders...)
}

// UpdateById update a model by id
func (m *RoleModel) UpdateById(id int64, role Role, onlyFields ...string) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).UpdateFields(role.StaledKV(onlyFields...), builders...)
}

// Delete remove a model
func (m *RoleModel) Delete(builders ...query.SQLBuilder) (int64, error) {

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).Table(m.tableName).ResolveDelete()

	res, err := m.db.ExecContext(context.Background(), sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}

// DeleteById remove a model by id
func (m *RoleModel) DeleteById(id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Delete()
}
