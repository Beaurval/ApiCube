package models

type citoyens struct {
	ADRESSE         sql.NullString `gorm:"column:ADRESSE"`                //
	CP              sql.NullString `gorm:"column:CP"`                     //
	DATEINSCRIPTION time.Time      `gorm:"column:DATE_INSCRIPTION"`       //
	GENRE           sql.NullString `gorm:"column:GENRE"`                  //
	IDCITYOEN       int            `gorm:"column:ID_CITYOEN;primary_key"` //
	IDRANG          int            `gorm:"column:ID_RANG"`                //
	MAIL            sql.NullString `gorm:"column:MAIL"`                   //
	MOTDEPASSE      sql.NullString `gorm:"column:MOT_DE_PASSE"`           //
	NOM             sql.NullString `gorm:"column:NOM"`                    //
	PRENOM          sql.NullString `gorm:"column:PRENOM"`                 //
	PSEUDO          sql.NullString `gorm:"column:PSEUDO"`                 //
	TELEPHONE       sql.NullString `gorm:"column:TELEPHONE"`              //
	VILLE           sql.NullString `gorm:"column:VILLE"`                  //
}

// TableName sets the insert table name for this struct type
func (c *citoyens) TableName() string {
	return "citoyens"
}
