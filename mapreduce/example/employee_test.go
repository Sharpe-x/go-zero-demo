package example

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_runEmployeeExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "runEmployeeExample",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runEmployeeExample()
		})
	}
}

func Test_runGSampleMap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "runGSampleMap",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runGSampleMap()
		})
	}
}

func TestTransform(t *testing.T) {
	type args struct {
		slice interface{}
		fn    interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "用于字符串数组",
			args: args{
				slice: []string{"1", "2", "3", "4", "5", "6"},
				fn: func(s string) string {
					return s + s + s
				},
			},
			want: []string{"111", "222", "333", "444", "555", "666"},
		},

		{
			name: "用于整形数组",
			args: args{
				slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				fn: func(s int) int {
					return s * 3
				},
			},
			want: []int{3, 6, 9, 12, 15, 18, 21, 24, 27},
		},

		{
			name: "用于结构体",
			args: args{
				slice: employeeList,
				fn: func(e Employee) Employee {
					e.Salary += 1000
					e.Age += 1
					return e
				},
			},
			want: buildWantEmployee(employeeList),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Transform(tt.args.slice, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transform() = %v, want %v", got, tt.want)
			} else {
				t.Logf("%s return is : %v", tt.name, got)
			}
		})
	}
	fmt.Printf("employeeList return after : %v", employeeList)
}

func buildWantEmployee(list []Employee) []Employee {
	var newList []Employee
	for _, employee := range list {
		employee.Salary += 1000
		employee.Age += 1
		newList = append(newList, employee)
	}
	return newList
}

func TestTransformInPlace(t *testing.T) {
	type args struct {
		slice interface{}
		fn    interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "用于字符串数组",
			args: args{
				slice: []string{"1", "2", "3", "4", "5", "6"},
				fn: func(s string) string {
					return s + s + s
				},
			},
			want: []string{"111", "222", "333", "444", "555", "666"},
		},

		{
			name: "用于整形数组",
			args: args{
				slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				fn: func(s int) int {
					return s * 3
				},
			},
			want: []int{3, 6, 9, 12, 15, 18, 21, 24, 27},
		},

		{
			name: "用于结构体",
			args: args{
				slice: employeeList,
				fn: func(e Employee) Employee {
					e.Salary += 1000
					e.Age += 1
					return e
				},
			},
			want: buildWantEmployee(employeeList),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransformInPlace(tt.args.slice, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransformInPlace() = %v, want %v", got, tt.want)
			} else {
				t.Logf("%s return is : %v", tt.name, got)
			}
		})
	}

	fmt.Printf("employeeList return after : %v", employeeList)
}

func TestRobustReduce(t *testing.T) {
	type args struct {
		slice    interface{}
		pairFunc interface{}
		zero     interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Reduce",
			args: args{
				slice: []int{1, 2, 3, 4},
				pairFunc: func(a, b int) int {
					return a + b
				},
				zero: 0,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RobustReduce(tt.args.slice, tt.args.pairFunc, tt.args.zero); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RobustReduce() = %v, want %v", got, tt.want)
			} else {
				t.Logf("%s return is : %v", tt.name, got)
			}
		})
	}
}

func TestRobustFilter(t *testing.T) {
	type args struct {
		slice    interface{}
		function interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Reduce",
			args: args{
				slice: []int{1, 2, 3, 4},
				function: func(a int) bool {
					return a%2 == 0
				},
			},
			want: []int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RobustFilter(tt.args.slice, tt.args.function); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RobustFilter() = %v, want %v", got, tt.want)
			} else {
				t.Logf("%s return is : %v", tt.name, got)
			}
		})
	}
}

func TestFilterInPlace(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	type args struct {
		slicePtr interface{}
		function interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Reduce",
			args: args{
				slicePtr: &nums,
				function: func(a int) bool {
					return a%2 == 0
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FilterInPlace(tt.args.slicePtr, tt.args.function)
		})
		fmt.Printf("after slicePtr = %v", tt.args.slicePtr)
	}
}
