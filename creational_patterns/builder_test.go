package creational_patterns

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuilder(t *testing.T) {
	employeeBuilder := &EmployeeBuilderImpl{}
	employee := employeeBuilder.setName("DoanThai").setSalary(5.00).build()

	assert.Equal(t, employee.Name, "DoanThai")
	assert.Equal(t, employee.Salary, float32(5.00))
}
