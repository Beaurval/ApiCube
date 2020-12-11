package models

type action_ressource struct {
	EXPLOITE    sql.NullInt64 `gorm:"column:EXPLOITE"`                 //
	FAVORIS     sql.NullInt64 `gorm:"column:FAVORIS"`                  //
	IDCITYOEN   int           `gorm:"column:ID_CITYOEN;primary_key"`   //
	IDRESSOURCE int           `gorm:"column:ID_RESSOURCE;primary_key"` //
	MISDECOTE   sql.NullInt64 `gorm:"column:MIS_DE_COTE"`              //
}

// TableName sets the insert table name for this struct type
func (a *action_ressource) TableName() string {
	return "action_ressource"
}
