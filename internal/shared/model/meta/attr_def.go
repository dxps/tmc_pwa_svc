package meta

type AttributeValueType int

const (
	Text     AttributeValueType = iota // Mapped to PostgreSQL's `text` type.
	Number                             // Mapped to PostgreSQL's `numeric` type.
	Date                               // Mapped to PostgreSQL's `date` type.
	DateTime                           // Mapped to PostgreSQL's `timestamptz` (with time zone) type.
	Bool                               // Mapped to PostgreSQL's `boolean` type.
)

type AttributeDef struct {
	Id           Id                 `db:"id"              json:"id"`
	Name         string             `db:"name"            json:"name"`
	Description  string             `db:"description"     json:"description"`
	ValueType    AttributeValueType `db:"value_type"      json:"valueType"`
	DefaultValue string             `db:"default_value"   json:"defaultValue"`
	IsRequired   bool               `db:"required"        json:"isRequired"`
}
