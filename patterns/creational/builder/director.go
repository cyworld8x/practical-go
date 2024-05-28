package builder

type director struct {
	builder IComputerBuilder
}

func NewDirector(builder IComputerBuilder) *director {
	return &director{
		builder: builder,
	}
}

func (d *director) setBuilder(builder IComputerBuilder) {
	d.builder = builder
}

func (d *director) buildComputer() *Computer {
	return d.builder.Build()
}
