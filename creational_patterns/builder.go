package creational_patterns

type Employee struct {
	Name string
	Salary float32
}

type EmployeeBuilder interface {
	setName(name string) EmployeeBuilder
	setSalary(salary float32) EmployeeBuilder
	build() Employee
}

type EmployeeBuilderImpl struct {
	Employee
}

func (builder *EmployeeBuilderImpl) setName(name string) EmployeeBuilder  {
	builder.Name = name
	return builder
}

func (builder *EmployeeBuilderImpl) setSalary(salary float32) EmployeeBuilder {
	builder.Salary = salary
	return builder
}

func (builder *EmployeeBuilderImpl) build() Employee  {
	return builder.Employee
}