package models

type RelationCitoyen struct {
	CitoyenID      uint
	CitoyenCibleID uint
	TypeRelationID uint
	TypeRelation   TypeRelation
	Approbation    bool
}

type CreateRelationCitoyenInput struct {
	CitoyenID      uint
	CitoyenCibleID uint
	TypeRelationID uint
}

type UpdateRelationCitoyenInput struct {
	Approbation bool
}

// TableName sets the insert table name for this struct type
func (r *RelationCitoyen) TableName() string {
	return "citoyen_relations"
}
