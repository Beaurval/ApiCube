package models

type type_relation struct {
	IDTYPERELATION int            `gorm:"column:ID_TYPE_RELATION;primary_key"` //
	NOM            sql.NullString `gorm:"column:NOM"`                          //
}

// TableName sets the insert table name for this struct type
func (t *type_relation) TableName() string {
	return "type_relation"
}
