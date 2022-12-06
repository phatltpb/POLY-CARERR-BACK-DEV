package role

const (
	IsStudent  = 1 << iota // 0001
	IsEmployer             // 0010
	IsADmin
)

const (
	Student  = IsStudent
	Employer = IsEmployer
	Admin    = IsADmin
)

func CheckRole(userRole string, validateRole string) error {
	return nil
}
