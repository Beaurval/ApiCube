package models

type rangs struct {
	IDRANG  int            `gorm:"column:ID_RANG;primary_key"` //
	NOMRANG sql.NullString `gorm:"column:NOM_RANG"`            //
}

// TableName sets the insert table name for this struct type
func (r *rangs) TableName() string {
	return "rangs"
}
