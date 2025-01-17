package employee

// Manager is a struct for manager
// It embeds Employee
// This means that Manager has all the fields and methods of Employee
type Manager struct {
	Employee
}

// NewManager creates a new manager
// It returns a pointer to the manager
// Creational method
func NewManager() *Manager {
	manager := &Manager{
		Employee: Employee{
			Name:   "Manager",
			Salary: 1000,
		},
	}
	return manager
}

func (m *Manager) GetBonus() float64 {
	return float64(m.Salary) * 0.2
}
