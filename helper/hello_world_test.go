package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Roki",
			request: "Roki",
		},
		{
			name:    "Nanda",
			request: "Nanda",
		},
		{
			name:    "RokiNandaSaputra",
			request: "Roki Nanda Saputra",
		},
		{
			name:    "Budi",
			request: "Budi Nugraha",
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

func BenchmarkSub(b *testing.B) {
	b.Run("Eko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Roki")
		}
	})
	b.Run("Kurniawan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Saputra")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Roki")
	}
}

func BenchmarkHelloWorldKurniawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Saputra")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Roki",
			request:  "Roki",
			expected: "Hello Roki",
		},
		{
			name:     "Saputra",
			request:  "Saputra",
			expected: "Hello Saputra",
		},
		{
			name:     "NandaSaputra",
			request:  "NandaSaputra",
			expected: "Hello NandaSaputra",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Nova",
			request:  "Nova",
			expected: "Hello Nova",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Nanda", func(t *testing.T) {
		result := HelloWorld("Nanda")
		require.Equal(t, "Hello Nanda", result, "Result must be 'Hello Nanda'")
	})
	t.Run("Saputra", func(t *testing.T) {
		result := HelloWorld("Saputra")
		require.Equal(t, "Hello Saputra", result, "Result must be 'Hello Saputra'")
	})
	t.Run("Bambang", func(t *testing.T) {
		result := HelloWorld("Bambang")
		require.Equal(t, "Hello Bambang", result, "Result must be 'Hello Bambang'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Roki")
	require.Equal(t, "Hello Roki", result, "Result must be 'Hello Roki'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Roki")
	require.Equal(t, "Hello Roki", result, "Result must be 'Hello Roki'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Roki")
	assert.Equal(t, "Hello Roki", result, "Result must be 'Hello Roki'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldRoki(t *testing.T) {
	result := HelloWorld("Roki")

	if result != "Hello Roki" {
		// error
		t.Error("Result must be 'Hello Roki'")
	}

	fmt.Println("TestHelloWorldRoki Done")
}

func TestHelloWorldSaputra(t *testing.T) {
	result := HelloWorld("Saputra")

	if result != "Hello Saputra" {
		// error
		t.Fatal("Result must be 'Hello Saputra'")
	}

	fmt.Println("TestHelloWorldSaputra Done")
}

func TestHelloWorldNova(t *testing.T) {
	result := HelloWorld("Nova")

	if result != "Hello Nova" {
		// error
		t.Fatal("Result must be 'Hello Nova'")
	}

	fmt.Println("TestHelloWorldNova Done")
}
