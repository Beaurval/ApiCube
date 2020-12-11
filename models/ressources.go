package models

type ressources struct {
	CONTENU          sql.NullString `gorm:"column:CONTENU"`                  //
	DATECREATION     time.Time      `gorm:"column:DATE_CREATION"`            //
	DATEMODIFICATION time.Time      `gorm:"column:DATE_MODIFICATION"`        //
	IDCITYOEN        int            `gorm:"column:ID_CITYOEN"`               //
	IDRESSOURCE      int            `gorm:"column:ID_RESSOURCE;primary_key"` //
	IDTYPERELATION   int            `gorm:"column:ID_TYPE_RELATION"`         //
	IDTYPERESSOURCE  int            `gorm:"column:ID_TYPE_RESSOURCE"`        //
	TITRE            sql.NullString `gorm:"column:TITRE"`                    //
	VOTE             sql.NullInt64  `gorm:"column:VOTE"`                     //
	VUES             sql.NullInt64  `gorm:"column:VUES"`                     //
}

// TableName sets the insert table name for this struct type
func (r *ressources) TableName() string {
	return "ressources"
}
