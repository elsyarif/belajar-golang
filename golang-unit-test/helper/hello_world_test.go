package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Syarif")

	if result != "Hello Syarif" {
		// Error
		//panic("Result is Not Hello Syarif")
		//t.Fail() // ketika terjadi error akan melanjutkan ke kode program selanjutnya
		t.Error("hasil harus mengandung Hello Syarif")
	}

	fmt.Println("Test Run Done")
}

func TestHelloWorld2(t *testing.T) {
	result := HelloWorld("Hidayatulloh")

	if result != "Hello Hidayatulloh" {
		// Error
		//panic("Result is Not Hello Syarif")
		// t.FailNow() // ketika terjadi error akan berhenti dan tidak melanjutkan kode program selnjutnya tp akan melanjutkan test yang lain
		t.Fatal("asil harus mengandung Hello Hidayatulloh")
	}

	fmt.Println("Test Run Done")
}

// Menggunakan package assertion
func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Syarif")

	assert.Equal(t, "Hi Syarif", result, "hasil harus mengandung Hello Syarif")
	fmt.Println("Test Run Done")
}

// Menggunakan package require
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Syarif")

	require.Equal(t, "Hi Syarif", result)
	fmt.Println("Test Run Done")
}

// Test Skip
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can't Run on Mac Os")
	}

	result := HelloWorld("Syarif")
	require.Equal(t, "Hi Syarif", result)
}
