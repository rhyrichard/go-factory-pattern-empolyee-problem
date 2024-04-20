package employee

// Staff is a struct for staff
// It embeds Employee
// This means that Staff has all the fields and methods of Employee
type Staff struct {
	Employee
}

// NewStaff creates a new staff
// It returns a pointer to the staff
// Creational method
func NewStaff() *Staff {
	staff := &Staff{
		Employee: Employee{
			Name:   "Staff",
			Salary: 500,
		},
	}
	return staff
}

func (s *Staff) GetBonus() float64 {
	return float64(s.Salary) * 0.1
}
