package creational_patterns

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Obj struct {
	name string
}

func TestStructMultiInstance(t *testing.T)  {
	obj1 := Obj{
		name: "obj",
	}
	obj2 := Obj{
		name: "obj",
	}
	r1 := &obj1
	obj3 := *r1

	assert.True(t, obj1 == obj2)
	assert.False(t, &obj1 == &obj2)
	assert.True(t, obj1 == obj3)
	assert.False(t, &obj1 == &obj3)
}

func TestSingleton(t *testing.T)  {
	obj1 := GetInstance()
	obj2 := GetInstance()

	fmt.Printf("%p\n", &obj1)
	fmt.Printf("%p\n", &obj2)
	obj1.Count = 2

	assert.True(t, obj1 == obj2)
	assert.Equal(t, obj1, obj2)
	assert.Equal(t, obj1.Count, obj2.Count)
}
