# pipeline
Simple function pipe lining package

[![Go Report Card](https://goreportcard.com/badge/github.com/julio77it/pipeline)](https://goreportcard.com/report/github.com/julio77it/pipeline)

```go
	pipe, err := pipeline.Chain(
		f1,
		f2,
		f3,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	pipe.Process(input)
```

It uses a lot of the reflect package, if functions in the pipeline haven't compatibles arguments and return values you'll get panic

> Reflection is never clear.

[![IMAGE ALT TEXT HERE](https://img.youtube.com/vi/PAAkCSZUG1c/0.jpg)](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=922s)