package models

type type_ressource struct {
	IDTYPERESSOURCE int            `gorm:"column:ID_TYPE_RESSOURCE;primary_key"` //
	NOM             sql.NullString `gorm:"column:NOM"`                           //
}

// TableName sets the insert table name for this struct type
func (t *type_ressource) TableName() string {
	return "type_ressource"
}
