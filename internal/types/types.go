package types

type Student struct {
	Id    int64
	Name  string `validate:"required"` // term: feild validation
	Email string `validate:"required"`
	Age   int    `validate:"required"`
}
