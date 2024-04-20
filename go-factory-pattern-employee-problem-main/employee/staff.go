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
			Bonus:  500.0 * 0.1,
		},
	}
	return staff
}
