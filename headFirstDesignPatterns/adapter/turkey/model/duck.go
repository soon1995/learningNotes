package model

type Duck interface {
	Quack()
	Fly()
}

type turkeyAdapter struct {
	turkey Turkey
}

func NewTurkeyAdapter(turkey Turkey) *turkeyAdapter {
  return &turkeyAdapter{
    turkey: turkey,
  }
}

func (t *turkeyAdapter) Quack() {
  t.turkey.Gobble()
}

func (t *turkeyAdapter) Fly() {
  for i := 0; i < 5; i++ {
    t.turkey.Fly()
  }
}
