package Builder

type Builder interface {
	Build()
}
type Director struct {
	builder Builder
}

func NewBuilder(b Builder) Director {
	return Director{builder: b}
}
func (d *Director) Construct() {
	d.builder.Build()
}
func NewDirector(b Builder) Director {
	return Director{builder: b}
}

type ConcreateBuilder struct {
	built bool
}

func NewConcreateBuilder() ConcreateBuilder {
	return ConcreateBuilder{false}
}

func (b *ConcreateBuilder) Build() {
	b.built = true
}

type Product struct {
	Built bool
}

func (b *ConcreateBuilder) GetResult() *Product {
	return &Product{b.built}
}
