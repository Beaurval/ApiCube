package models

type fichiers struct {
	DATEAJOUT   time.Time      `gorm:"column:DATE_AJOUT"`             //
	IDFICHIER   int            `gorm:"column:ID_FICHIER;primary_key"` //
	IDRESSOURCE int            `gorm:"column:ID_RESSOURCE"`           //
	NOM         sql.NullString `gorm:"column:NOM"`                    //
	TAILLE      sql.NullInt64  `gorm:"column:TAILLE"`                 //
	TYPE        sql.NullString `gorm:"column:TYPE"`                   //
	URL         sql.NullString `gorm:"column:URL"`                    //
}

// TableName sets the insert table name for this struct type
func (f *fichiers) TableName() string {
	return "fichiers"
}
