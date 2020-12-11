package models

type tager struct {
	IDRESSOURCE int `gorm:"column:ID_RESSOURCE;primary_key"` //
	IDTAG       int `gorm:"column:ID_TAG;primary_key"`       //
}

// TableName sets the insert table name for this struct type
func (t *tager) TableName() string {
	return "tager"
}
