# fileroller [![GoDoc](https://godoc.org/github.com/1800alex/go-utilities-fileroller?status.svg)](https://godoc.org/github.com/1800alex/go-utilities-fileroller) [![Build Status](https://travis-ci.com/1800alex/go-utilities-fileroller.png?branch=master)](https://travis-ci.com/1800alex/go-utilities-fileroller)
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



