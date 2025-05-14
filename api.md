# Apps

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go">kernel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go#AppDeployResponse">AppDeployResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go">kernel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go#AppInvokeResponse">AppInvokeResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go">kernel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go#AppGetInvocationResponse">AppGetInvocationResponse</a>

Methods:

- <code title="post /apps/deploy">client.Apps.<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go#AppService.Deploy">Deploy</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go">kernel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go#AppDeployParams">AppDeployParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go">kernel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go#AppDeployResponse">AppDeployResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /apps/invoke">client.Apps.<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go#AppService.Invoke">Invoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go">kernel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go#AppInvokeParams">AppInvokeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go">kernel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go#AppInvokeResponse">AppInvokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /apps/invocations/{id}">client.Apps.<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go#AppService.GetInvocation">GetInvocation</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go">kernel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go#AppGetInvocationResponse">AppGetInvocationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Browser

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go">kernel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go#BrowserNewSessionResponse">BrowserNewSessionResponse</a>

Methods:

- <code title="post /browser">client.Browser.<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go#BrowserService.NewSession">NewSession</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go">kernel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go#BrowserNewSessionParams">BrowserNewSessionParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go">kernel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/kernel-go#BrowserNewSessionResponse">BrowserNewSessionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
