package singleton

import "sync"

type singleton struct {
	count int
	mux   sync.Mutex
}

var instance *singleton

//GetInstance get an instance
func GetInstance() *singleton {
	var (
		mux sync.Mutex
		i   int
	)
	mux.Lock()
	if instance == nil {
		i++
		instance = new(singleton)
		println("new instance", i)
	}
	mux.Unlock()
	return instance
}

func (s *singleton) AddOne() {
	s.mux.Lock()
	s.count++
	s.mux.Unlock()
}

func (s *singleton) GetCount() int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.count
}
