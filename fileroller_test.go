package fileroller

import (
	"github.com/1800alex/go-utilities-password"
	"bufio"
	"fmt"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestFileRoller_Roll(t *testing.T) {
	Stubs()

	fr, err := NewFileRoller("out.log", 30, 10)

	assert.NoError(t, err)

	var contents []byte

	// get some random strings
	testString, err := password.Generate(32, true, false, false, true)
	assert.NoError(t, err)

	testString2, err := password.Generate(32, true, false, false, true)
	assert.NoError(t, err)

	afero.WriteFile(DepFS, "out.log", []byte(testString), 0666)

	contents, err = afero.ReadFile(DepFS, "out.log")
	assert.NoError(t, err)

	assert.Equal(t, testString, string(contents))

	fr.Roll(len(testString))
	afero.WriteFile(DepFS, "out.log", []byte(testString2), 0666)

	contents, err = afero.ReadFile(DepFS, "out.log")
	assert.NoError(t, err)

	assert.Equal(t, testString2, string(contents))

	contents, err = afero.ReadFile(DepFS, "out.log0")
	assert.NoError(t, err)

	assert.Equal(t, testString, string(contents))

	StubsRestore()
}

func BenchmarkFileRoller_Roll(b *testing.B) {
	Stubs()

	logfile := "out.log"
	fr, err := NewFileRoller(logfile, 128, 10)

	if err != nil {
		fmt.Println(err)
		return
	}

	testString := time.Now().Format(time.RFC850) + ": test log message\n"

	for n := 0; n < b.N; n++ {
		// Will roll log if it exceeds current size + new string > 1024 bytes
		fr.Roll(len(testString))

		var f afero.File

		f, err = DepFS.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

		if err != nil {
			return
		}

		defer f.Close()

		w := bufio.NewWriter(f)

		var n int
		n, err = w.WriteString(testString)

		if err != nil {
			fmt.Println(err)
			return
		}

		if n != len(testString) {
			fmt.Println("failed to write log")
			return
		}

		w.Flush()
	}

	StubsRestore()
}

func ExampleFileRoller_Roll() {
	logfile := "out.log"
	fr, err := NewFileRoller(logfile, 1024, 3)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 100; i++ {
		testString := time.Now().Format(time.RFC850) + ": test log message\n"

		// Will roll log if it exceeds current size + new string > 1024 bytes
		fr.Roll(len(testString))

		var f afero.File

		f, err = DepFS.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

		if err != nil {
			return
		}

		defer f.Close()

		w := bufio.NewWriter(f)

		var n int
		n, err = w.WriteString(testString)

		if err != nil {
			fmt.Println(err)
			return
		}

		if n != len(testString) {
			fmt.Println("failed to write log")
			return
		}

		w.Flush()
	}

	// //Output:
}
