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
	Id           string             `json:"id"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	ValueType    AttributeValueType `json:"valueType"`
	DefaultValue string             `json:"defaultValue"`
	IsRequired   bool               `json:"isRequired"`
}
