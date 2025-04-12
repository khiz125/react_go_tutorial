package domain

type TodoFilter struct {
	Completed *bool `bson:"completed,omitempty"`
}