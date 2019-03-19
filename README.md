# fileroller
[![Build Status](https://travis-ci.com/1800alex/go-utilities-fileroller.svg?branch=master)](https://travis-ci.com/1800alex/go-utilities-fileroller)
[![Coverage Status](https://coveralls.io/repos/github/1800alex/go-utilities-fileroller/badge.svg?branch=master)](https://coveralls.io/github/1800alex/go-utilities-fileroller?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/1800alex/go-utilities-fileroller)](https://goreportcard.com/report/github.com/1800alex/go-utilities-fileroller)
[![GoDoc](https://godoc.org/github.com/1800alex/go-utilities-fileroller?status.svg)](https://godoc.org/github.com/1800alex/go-utilities-fileroller)

Package fileroller is a package capable rolling files when they becoming > a certain size The fileroller has the following features: - Log rolling with configurable size - Configurable number of logs

Download:
```shell
go get github.com/1800alex/go-utilities-fileroller
```

* * *
Package fileroller is a package capable rolling files when they becoming > a certain size
The fileroller has the following features:

```
- Log rolling with configurable size
- Configurable number of logs
```





# Examples

FileRoller Roll
Code:

```
{
	logfile := "out.log"
	fr, err := NewFileRoller(logfile, 1024, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 100; i++ {
		testString := time.Now().Format(time.RFC850) + ": test log message\n"
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
}
```



