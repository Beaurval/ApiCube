package models

type RelationCitoyen struct {
	CitoyenID      uint
	CitoyenCibleID uint
	TypeRelationID uint

	Approbation *bool

	Citoyen      Citoyen `gorm:"foreignkey:CitoyenID"`
	CitoyenCible Citoyen `gorm:"foreignkey:CitoyenCibleID"`
	TypeRelation TypeRelation
}

type CreateRelationCitoyenInput struct {
	CitoyenID      uint
	CitoyenCibleID uint
	TypeRelationID uint
	Approbation    *bool
}

type UpdateRelationCitoyenInput struct {
	CitoyenID      uint
	CitoyenCibleID uint
	Approbation    *bool `gorm:"type:boolean"`
}

// TableName sets the insert table name for this struct type
func (r *RelationCitoyen) TableName() string {
	return "citoyen_relations"
}
