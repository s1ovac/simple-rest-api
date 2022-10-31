package storage

import (
	"fmt"
	"sync"
)

type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

type Storage interface {
	Insert(e *Employee)
	Get(id int) (Employee, error)
	Update(id int, e *Employee) error
	Delete(id int)
	GetAll() ([]Employee, error)
}

type MemoryStorage struct {
	counter int
	data    map[int]Employee
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		counter: 0,
		data:    make(map[int]Employee),
	}
}

func (s *MemoryStorage) Insert(e *Employee) {
	s.Lock()

	e.ID = s.counter
	s.data[e.ID] = *e

	s.counter++
	s.Unlock()
}

func (s *MemoryStorage) Get(id int) (Employee, error) {
	s.Lock()
	defer s.Unlock()

	employee, ok := s.data[id]
	if !ok {
		return employee, fmt.Errorf("employee not found")
	}

	return employee, nil
}

func (s *MemoryStorage) Update(id int, e *Employee) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.data[id]; !ok {
		return fmt.Errorf("employee not found")
	}
	s.data[id] = *e
	return nil
}

func (s *MemoryStorage) Delete(id int) {
	s.Lock()
	defer s.Unlock()

	delete(s.data, id)
}

func (s *MemoryStorage) GetAll() ([]Employee, error) {
	s.Lock()
	defer s.Unlock()
	var employees []Employee

	for num, emp := range s.data {
		if _, ok := s.data[num]; !ok {
			return employees, fmt.Errorf("employee with this id not found")
		}
		employees = append(employees, emp)
	}
	return employees, nil
}
