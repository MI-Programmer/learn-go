package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTableSum(b *testing.B) {
	benchmarks := []struct {
		Name    string
		Request []int
	}{
		{"SUM 60", []int{30, 30}},
		{"SUM 1000", []int{500, 500}},
		{"SUM 700", []int{350, 350}},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Sum(benchmark.Request[0], benchmark.Request[1])
			}
		})
	}
}

func BenchmarkTableHelloWorld(b *testing.B) {
	benchmarks := []struct {
		Name    string
		Request string
	}{
		{"MI", "MI"},
		{"Programmer", "Programmer"},
		{"Ganteng", "Ganteng"},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.Request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("MI", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("MI")
		}
	})

	b.Run("Programmer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Programmer")
		}
	})
}

func BenchmarkHelloWorldMI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("MI")
	}
}

func BenchmarkHelloWorldProgrammer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Programmer")
	}
}

type TestTable struct {
	Name     string
	Request  string
	Expected string
}

type TestsSlice []TestTable

func TestTableHelloWorld(t *testing.T) {
	tests := TestsSlice{
		{"MI", "MI", "Hello MI"},
		{"Programmer", "Programmer", "Hello Programmer"},
		{"Ganteng", "Ganteng", "Hello Ganteng"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			result := HelloWorld(test.Request)
			assert.Equal(t, test.Expected, result)
		})
	}
}

func TestTableFullName(t *testing.T) {
	tests := []struct {
		Name     string
		Request  []string
		Expected string
	}{
		{
			Name:     "Peter Parker",
			Request:  []string{"Peter", "Parker"},
			Expected: "Peter Parker",
		},
		{
			Name:     "Vincent Rompies",
			Request:  []string{"Vincent", "Rompies"},
			Expected: "Vincent Rompies",
		},
		{
			Name:     "MI Programmer",
			Request:  []string{"MI", "Programmer"},
			Expected: "MI Programmer",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := FullName(test.Request[0], test.Request[1])
			assert.Equal(t, test.Expected, result)
		})
	}
}

func TestTableSum(t *testing.T) {
	tests := []struct {
		Name     string
		Request  []int
		Expected int
	}{
		{"Sum 1000", []int{500, 500}, 1000},
		{"Sum 200", []int{100, 100}, 200},
		{"Sum 60", []int{30, 30}, 60},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := Sum(test.Request[0], test.Request[1])
			assert.Equal(t, test.Expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("MI", func(t *testing.T) {
		result := HelloWorld("MI")

		require.Equal(t, "Hello MI", result, "result must be 'Hello MI'")
	})

	t.Run("Programmer", func(t *testing.T) {
		result := HelloWorld("Programmer")

		require.Equal(t, "Hello Programmer", result, "result must be 'Hello Programmer'")
	})

}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE unit test")

	m.Run()

	fmt.Println("AFTER unit test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}

	result := HelloWorld("ilham")

	require.Equal(t, "Hello ilham", result, "result must be 'Hello ilham'")

	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("ilham")

	require.Equal(t, "Hello ilham", result, "result must be 'Hello ilham'")

	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("ilham")

	assert.Equal(t, "Hello ilham", result, "result must be 'Hello ilham'")

	fmt.Println("TestHelloWorld with Assert Done")
}

func TestFullNameAssert(t *testing.T) {
	result := FullName("MI", "Programmer")

	assert.Equal(t, "MI Programmer", result, "result must be 'MI Programmer'")

	fmt.Println("TestFullName with Assert Done")
}

func TestSumAssert(t *testing.T) {
	result := Sum(10, 10)

	assert.Equal(t, 20, result, "result must be '20'")

	fmt.Println("TestSum with Assert Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("ilham")

	if result != "Hello ilham" {
		// t.Fail()
		t.Error("result must be 'Hello ilham'")
	}

	fmt.Println("Done testing for TestHelloWorld")
}

func TestFullName(t *testing.T) {
	result := FullName("MI", "Programmer")

	if result != "MI Programmer" {
		// t.FailNow()
		t.Fatal("result must be 'MI Programmer'")
	}

	fmt.Println("Done testing for TestFullName")
}

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		t.Fatal("result must be '20'")
	}

	fmt.Println("Done testing for TestSum")
}
