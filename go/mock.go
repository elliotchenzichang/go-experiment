package _go

//go:generate mockery --name=Man
type Man interface {
	GetName() string
	IsHandSomeBoy() bool
}
