package models

type tags struct {
	IDTAG int            `gorm:"column:ID_TAG;primary_key"` //
	NOM   sql.NullString `gorm:"column:NOM"`                //
}

// TableName sets the insert table name for this struct type
func (t *tags) TableName() string {
	return "tags"
}
