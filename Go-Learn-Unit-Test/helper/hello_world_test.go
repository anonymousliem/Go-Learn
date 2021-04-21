package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Jokowi")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Jokowi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Jokowi")
		}
	})

	b.Run("Prabowo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Prabowo")
		}
	})

}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "Jokowi",
			request: "Jokowi",
			expected: "Halo Jokowi",
		},
		{
			name: "Prabowo",
			request: "Prabowo",
			expected: "Halo Prabowo",
		},
		{
			name: "Sandi",
			request: "Sandi",
			expected: "Halo Sandi",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
func TestTableHelloWorld(t *testing.T) {
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name : "Jokowi",
			request: "Jokowi",
			expected: "Halo Jokowi",
		},
		{
			name : "Prabowo",
			request: "Prabowo",
			expected: "Halo Prabowo",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := HelloWorld(test.request)
			require.Equal(t, test.expected, results, "Tidak sama")
		})
	}
}
func TestSubTest(t *testing.T)  {
	t.Run("Jokowi", func(t *testing.T) {
		result := HelloWorld("Jokowi")
		require.Equal(t, "Halo Jokowi", result, "Result is not Halo Jokowi")
	})
	
	t.Run("Prabowo", func(t *testing.T) {
		result := HelloWorld("Prabowo")
		require.Equal(t, "Halo Prabowo", result, "Result is not Halo Prabowo")
	})
}

func TestHelloWorld(t *testing.T) {
	 result := HelloWorld("Jokowi")

	 fmt.Println(result)
	 if result != "Halo Jokowi" {

	 	t.Error("Result is not Jokowi")
	 }
}

func  TestHelloWorldAsset(t *testing.T) {
	result := HelloWorld("Jokowi")
	assert.Equal(t, "Halo Jokowi", result, "Result is not Halo Jokowi")
	fmt.Println("Test Assert Is Done")
}

func  TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Jokowi")
	require.Equal(t, "Halo Jokowi", result, "Result is not Halo Jokowi")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Only Run On Windows OS")
	}
	result := HelloWorld("Jokowi")
	assert.Equal(t, "Halo Jokowi", result, "Result is not Halo Jokowi")
}

func TestMain(m *testing.M)  {
	fmt.Println("Before Unit Test")
	m.Run()
	fmt.Println("After Unit Test")
}