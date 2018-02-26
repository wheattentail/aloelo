package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set UserQuerySet

// UserQuerySet is an queryset type for User
type UserQuerySet struct {
	db *gorm.DB
}

// NewUserQuerySet constructs new UserQuerySet
func NewUserQuerySet(db *gorm.DB) UserQuerySet {
	return UserQuerySet{
		db: db.Model(&User{}),
	}
}

func (qs UserQuerySet) w(db *gorm.DB) UserQuerySet {
	return NewUserQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) All(ret *[]User) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *User) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtEq(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtGt(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtGte(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtLt(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtLte(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtNe(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *User) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) Delete() error {
	return qs.db.Delete(User{}).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtEq(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtGt(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtGte(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtIsNotNull() UserQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtIsNull() UserQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtLt(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtLte(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtNe(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) GetUpdater() UserUpdater {
	return NewUserUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDEq(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDGt(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDGte(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDIn(ID uint, IDRest ...uint) UserQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDLt(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDLte(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDNe(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDNotIn(ID uint, IDRest ...uint) UserQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) Limit(limit int) UserQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// LocalBitcoinsKeyEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) LocalBitcoinsKeyEq(localBitcoinsKey string) UserQuerySet {
	return qs.w(qs.db.Where("local_bitcoins_key = ?", localBitcoinsKey))
}

// LocalBitcoinsKeyIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) LocalBitcoinsKeyIn(localBitcoinsKey string, localBitcoinsKeyRest ...string) UserQuerySet {
	iArgs := []interface{}{localBitcoinsKey}
	for _, arg := range localBitcoinsKeyRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("local_bitcoins_key IN (?)", iArgs))
}

// LocalBitcoinsKeyNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) LocalBitcoinsKeyNe(localBitcoinsKey string) UserQuerySet {
	return qs.w(qs.db.Where("local_bitcoins_key != ?", localBitcoinsKey))
}

// LocalBitcoinsKeyNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) LocalBitcoinsKeyNotIn(localBitcoinsKey string, localBitcoinsKeyRest ...string) UserQuerySet {
	iArgs := []interface{}{localBitcoinsKey}
	for _, arg := range localBitcoinsKeyRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("local_bitcoins_key NOT IN (?)", iArgs))
}

// LocalBitcoinsSecretEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) LocalBitcoinsSecretEq(localBitcoinsSecret string) UserQuerySet {
	return qs.w(qs.db.Where("local_bitcoins_secret = ?", localBitcoinsSecret))
}

// LocalBitcoinsSecretIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) LocalBitcoinsSecretIn(localBitcoinsSecret string, localBitcoinsSecretRest ...string) UserQuerySet {
	iArgs := []interface{}{localBitcoinsSecret}
	for _, arg := range localBitcoinsSecretRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("local_bitcoins_secret IN (?)", iArgs))
}

// LocalBitcoinsSecretNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) LocalBitcoinsSecretNe(localBitcoinsSecret string) UserQuerySet {
	return qs.w(qs.db.Where("local_bitcoins_secret != ?", localBitcoinsSecret))
}

// LocalBitcoinsSecretNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) LocalBitcoinsSecretNotIn(localBitcoinsSecret string, localBitcoinsSecretRest ...string) UserQuerySet {
	iArgs := []interface{}{localBitcoinsSecret}
	for _, arg := range localBitcoinsSecretRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("local_bitcoins_secret NOT IN (?)", iArgs))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs UserQuerySet) One(ret *User) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByCreatedAt() UserQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByDeletedAt() UserQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByID() UserQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByUpdatedAt() UserQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByCreatedAt() UserQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByDeletedAt() UserQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByID() UserQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByUpdatedAt() UserQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetCreatedAt(createdAt time.Time) UserUpdater {
	u.fields[string(UserDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetDeletedAt(deletedAt *time.Time) UserUpdater {
	u.fields[string(UserDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetID(ID uint) UserUpdater {
	u.fields[string(UserDBSchema.ID)] = ID
	return u
}

// SetLocalBitcoinsKey is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetLocalBitcoinsKey(localBitcoinsKey string) UserUpdater {
	u.fields[string(UserDBSchema.LocalBitcoinsKey)] = localBitcoinsKey
	return u
}

// SetLocalBitcoinsSecret is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetLocalBitcoinsSecret(localBitcoinsSecret string) UserUpdater {
	u.fields[string(UserDBSchema.LocalBitcoinsSecret)] = localBitcoinsSecret
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetUpdatedAt(updatedAt time.Time) UserUpdater {
	u.fields[string(UserDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetUsername is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetUsername(username string) UserUpdater {
	u.fields[string(UserDBSchema.Username)] = username
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u UserUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u UserUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtEq(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtGt(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtGte(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtLt(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtLte(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtNe(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// UsernameEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UsernameEq(username string) UserQuerySet {
	return qs.w(qs.db.Where("username = ?", username))
}

// UsernameIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UsernameIn(username string, usernameRest ...string) UserQuerySet {
	iArgs := []interface{}{username}
	for _, arg := range usernameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("username IN (?)", iArgs))
}

// UsernameNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UsernameNe(username string) UserQuerySet {
	return qs.w(qs.db.Where("username != ?", username))
}

// UsernameNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UsernameNotIn(username string, usernameRest ...string) UserQuerySet {
	iArgs := []interface{}{username}
	for _, arg := range usernameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("username NOT IN (?)", iArgs))
}

// ===== END of query set UserQuerySet

// ===== BEGIN of User modifiers

type userDBSchemaField string

func (f userDBSchemaField) String() string {
	return string(f)
}

// UserDBSchema stores db field names of User
var UserDBSchema = struct {
	ID                  userDBSchemaField
	CreatedAt           userDBSchemaField
	UpdatedAt           userDBSchemaField
	DeletedAt           userDBSchemaField
	Username            userDBSchemaField
	LocalBitcoinsKey    userDBSchemaField
	LocalBitcoinsSecret userDBSchemaField
}{

	ID:                  userDBSchemaField("id"),
	CreatedAt:           userDBSchemaField("created_at"),
	UpdatedAt:           userDBSchemaField("updated_at"),
	DeletedAt:           userDBSchemaField("deleted_at"),
	Username:            userDBSchemaField("username"),
	LocalBitcoinsKey:    userDBSchemaField("local_bitcoins_key"),
	LocalBitcoinsSecret: userDBSchemaField("local_bitcoins_secret"),
}

// Update updates User fields by primary key
func (o *User) Update(db *gorm.DB, fields ...userDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":                    o.ID,
		"created_at":            o.CreatedAt,
		"updated_at":            o.UpdatedAt,
		"deleted_at":            o.DeletedAt,
		"username":              o.Username,
		"local_bitcoins_key":    o.LocalBitcoinsKey,
		"local_bitcoins_secret": o.LocalBitcoinsSecret,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update User %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// UserUpdater is an User updates manager
type UserUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewUserUpdater creates new User updater
func NewUserUpdater(db *gorm.DB) UserUpdater {
	return UserUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&User{}),
	}
}

// ===== END of User modifiers

// ===== END of all query sets
