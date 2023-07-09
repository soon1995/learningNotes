package main

var Single = &singleton{}

type singleton struct {
	uniqueInstance *singleton
}

func (s *singleton) GetInstance() *singleton {
	if s.uniqueInstance == nil {
		s.uniqueInstance = &singleton{}
	}
	return s.uniqueInstance
}
