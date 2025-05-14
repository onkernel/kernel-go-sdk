# Apps

Response Types:

- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppDeployResponse">AppDeployResponse</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppInvokeResponse">AppInvokeResponse</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppGetInvocationResponse">AppGetInvocationResponse</a>

Methods:

- <code title="post /apps/deploy">client.Apps.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppService.Deploy">Deploy</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppDeployParams">AppDeployParams</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppDeployResponse">AppDeployResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /apps/invoke">client.Apps.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppService.Invoke">Invoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppInvokeParams">AppInvokeParams</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppInvokeResponse">AppInvokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /apps/invocations/{id}">client.Apps.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppService.GetInvocation">GetInvocation</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppGetInvocationResponse">AppGetInvocationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Browser

Response Types:

- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserNewSessionResponse">BrowserNewSessionResponse</a>

Methods:

- <code title="post /browser">client.Browser.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserService.NewSession">NewSession</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserNewSessionParams">BrowserNewSessionParams</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserNewSessionResponse">BrowserNewSessionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
