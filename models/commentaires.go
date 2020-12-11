package models

type commentaires struct {
	CONTENU       sql.NullString `gorm:"column:CONTENU"`                    //
	DATEECRITURE  time.Time      `gorm:"column:DATE_ECRITURE"`              //
	IDCITYOEN     int            `gorm:"column:ID_CITYOEN"`                 //
	IDCOMMENTAIRE int            `gorm:"column:ID_COMMENTAIRE;primary_key"` //
	IDRESSOURCE   int            `gorm:"column:ID_RESSOURCE"`               //
	PARENT        sql.NullInt64  `gorm:"column:PARENT"`                     //
	VOTE          sql.NullInt64  `gorm:"column:VOTE"`                       //
}

// TableName sets the insert table name for this struct type
func (c *commentaires) TableName() string {
	return "commentaires"
}
