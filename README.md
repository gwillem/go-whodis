# Who Dis?

Generate a friendly Git(hub) repository URL for the current file and line number in your Go code. Useful for logs, templates, cron jobs and such, so people can quickly find out where to edit the application. Uses `runtime.Caller()` and the main module name to guesstimate a repo URL. Currently supports Github but support for other public git services can be easily added.

# Usage
```go
import "github.com/gwillem/go-whodis"

// ...

fmt.Println("Generated by", whodis.URL())
// Generated by https://github.com/gwillem/someapp/blob/HEAD/pkg/pkg.go#L11
```
