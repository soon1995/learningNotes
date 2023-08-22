package factory

type ChocolateBoiler struct {
	empty  bool
	boiled bool
}

func NewChocolateBoiler() *ChocolateBoiler {
	return &ChocolateBoiler{true, false}
}

func (f *ChocolateBoiler) Fill() {
	if f.empty {
		f.empty = false
		f.boiled = false
	}
}

func (f *ChocolateBoiler) Drain() {
	if !f.empty && f.boiled {
		f.empty = true
	}
}

func (f *ChocolateBoiler) Boil() {
	if !f.empty && !f.boiled {
		f.boiled = true
	}
}

func (f *ChocolateBoiler) IsEmpty() bool {
	return f.empty
}

func (f *ChocolateBoiler) IsBoiled() bool {
	return f.boiled
}
