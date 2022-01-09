package example

import (
	"fmt"
	"reflect"
	"strings"
)

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

var employeeList = []Employee{
	{"Hao", 44, 0, 8000},
	{"Bob", 34, 10, 5000},
	{"Alice", 23, 5, 9000},
	{"Jack", 26, 0, 4000},
	{"Tom", 48, 9, 7500},
	{"Marry", 29, 0, 6000},
	{"Mike", 32, 8, 4000},
}

// EmployeeCountIf 统计满足某个条件的个数
func EmployeeCountIf(list []Employee, fn func(e *Employee) bool) int {
	count := 0
	for _, employee := range list {
		if fn(&employee) {
			count++
		}
	}
	return count
}

// EmployeeSumIf 统计满足某个条件的总数
func EmployeeSumIf(list []Employee, fn func(e *Employee) int) int {
	sum := 0
	for _, employee := range list {
		sum += fn(&employee)
	}
	return sum
}

// EmployeeFilter 按某种条件过虑Employee
func EmployeeFilter(list []Employee, fn func(e *Employee) bool) []Employee {
	var newList []Employee
	for _, employee := range list {
		if fn(&employee) {
			newList = append(newList, employee)
		}
	}
	return newList
}

func runEmployeeExample() {
	// 统计有多少员工大于40岁
	old := EmployeeCountIf(employeeList, func(e *Employee) bool {
		return e.Age > 40
	})
	fmt.Printf("old people: %d\n", old)

	// 统计有多少员工薪水大于6000
	highPay := EmployeeCountIf(employeeList, func(e *Employee) bool {
		return e.Salary >= 6000
	})
	fmt.Printf("High Salary people: %d\n", highPay)

	// 列出有没有休假的员工
	noVacation := EmployeeFilter(employeeList, func(e *Employee) bool {
		return e.Vacation == 0
	})
	fmt.Printf("People no vacation: %v\n", noVacation)

	// 统计所有员工的薪资总和
	totalPay := EmployeeSumIf(employeeList, func(e *Employee) int {
		return e.Salary
	})
	fmt.Printf("Total Salary: %d\n", totalPay)
}

// Map 简单版泛型map
// 通过 reflect.ValueOf() 来获得 interface{} 的值，其中一个是数据 vData，另一个是函数 vfn
// 然后通过 vfn.Call() 方法来调用函数，通过 []refelct.Value{vdata.Index(i)}来获得数据
func Map(data interface{}, fn interface{}) []interface{} {
	vfn := reflect.ValueOf(fn)
	vData := reflect.ValueOf(data)

	result := make([]interface{}, vData.Len())
	for i := 0; i < vData.Len(); i++ {
		result[i] = vfn.Call([]reflect.Value{vData.Index(i)})[0].Interface()
	}
	return result
}

func runGSampleMap() {
	square := func(x int) int {
		return x * x
	}

	nums := []int{1, 2, 3, 4}
	fmt.Println(Map(nums, square))

	upCase := func(s string) string {
		return strings.ToUpper(s)
	}

	strs := []string{"a", "b", "hello"}
	fmt.Println(Map(strs, upCase))

	// will panic
	/*x := Map(5, 5)
	fmt.Println(x)*/

}

// Transform 健壮版的 Map
func Transform(slice, fn interface{}) interface{} {
	return transform(slice, fn, false)
}

// TransformInPlace 原地替换
func TransformInPlace(slice, fn interface{}) interface{} {
	return transform(slice, fn, true)
}

// transform inPlace  是否就地完成
func transform(slice, function interface{}, inPlace bool) interface{} {
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("transform: not slice")
	}

	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("transform: function must be of type func(" + sliceInType.Type().Elem().String() + ") outputElemType")
	}

	sliceOutType := sliceInType
	if !inPlace {
		sliceOutType = reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)), sliceInType.Len(), sliceInType.Len())
	}

	for i := 0; i < sliceInType.Len(); i++ {
		sliceOutType.Index(i).Set(fn.Call([]reflect.Value{sliceInType.Index(i)})[0])
	}

	return sliceOutType.Interface()
}

func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {

	//Check it is a function
	if fn.Kind() != reflect.Func {
		return false
	}

	// NumIn() - returns a function type's input parameter count.
	// NumOut() - returns a function type's output parameter count.
	if fn.Type().NumIn() != len(types)-1 || fn.Type().NumOut() != 1 {
		return false
	}

	// In() - returns the type of a function type's i'th input parameter.
	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}

	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}

// RobustReduce 健壮版的 Generic Reduce
func RobustReduce(slice, pairFunc, zero interface{}) interface{} {
	sliceType := reflect.ValueOf(slice)
	if sliceType.Kind() != reflect.Slice {
		panic("reduce: wrong type, not slice")
	}

	length := sliceType.Len()
	if length == 0 {
		return zero
	}

	elemType := sliceType.Type().Elem()
	fn := reflect.ValueOf(pairFunc)
	if !verifyFuncSignature(fn, elemType, elemType, elemType) {
		t := elemType.String()
		panic("reduce: function must be of type func(" + t + ", " + t + ") " + t)
	}

	var ins [2]reflect.Value
	ins[0] = sliceType.Index(0)
	ins[1] = sliceType.Index(1)
	out := fn.Call(ins[:])[0]

	for i := 2; i < length; i++ {
		ins[0] = out
		ins[1] = sliceType.Index(i)
		out = fn.Call(ins[:])[0]
	}

	return out.Interface()
}

// RobustFilter 健壮版的 Generic Filter
func RobustFilter(slice, function interface{}) interface{} {
	result, _ := filter(slice, function, false)
	return result
}

func FilterInPlace(slicePtr, function interface{}) {
	in := reflect.ValueOf(slicePtr)
	if in.Kind() != reflect.Ptr {
		panic("FilterInPlace: wrong type, " +
			"not a pointer to slice")
	}
	_, n := filter(in.Elem().Interface(), function, true)
	in.Elem().SetLen(n)
}

var boolType = reflect.ValueOf(true).Type()

func filter(slice, function interface{}, inPlace bool) (interface{}, int) {

	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("filter: wrong type, not a slice")
	}

	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, boolType) {
		panic("filter: function must be of type func(" + elemType.String() + ") bool")
	}

	var which []int
	for i := 0; i < sliceInType.Len(); i++ {
		if fn.Call([]reflect.Value{sliceInType.Index(i)})[0].Bool() {
			which = append(which, i)
		}
	}

	out := sliceInType

	if !inPlace {
		out = reflect.MakeSlice(sliceInType.Type(), len(which), len(which))
	}
	for i := range which {
		out.Index(i).Set(sliceInType.Index(which[i]))
	}

	return out.Interface(), len(which)
}
