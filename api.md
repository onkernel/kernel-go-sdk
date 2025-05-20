# Apps

Response Types:

- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppListResponse">AppListResponse</a>

Methods:

- <code title="get /apps">client.Apps.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppListParams">AppListParams</a>) ([]<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppListResponse">AppListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Deployments

Response Types:

- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppDeploymentNewResponse">AppDeploymentNewResponse</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppDeploymentFollowResponseUnion">AppDeploymentFollowResponseUnion</a>

Methods:

- <code title="post /deploy">client.Apps.Deployments.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppDeploymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppDeploymentNewParams">AppDeploymentNewParams</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppDeploymentNewResponse">AppDeploymentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /apps/{id}/events">client.Apps.Deployments.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppDeploymentService.Follow">Follow</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppDeploymentFollowResponseUnion">AppDeploymentFollowResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Invocations

Response Types:

- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppInvocationNewResponse">AppInvocationNewResponse</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppInvocationGetResponse">AppInvocationGetResponse</a>

Methods:

- <code title="post /invocations">client.Apps.Invocations.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppInvocationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppInvocationNewParams">AppInvocationNewParams</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppInvocationNewResponse">AppInvocationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /invocations/{id}">client.Apps.Invocations.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppInvocationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#AppInvocationGetResponse">AppInvocationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Browsers

Response Types:

- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserNewResponse">BrowserNewResponse</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserGetResponse">BrowserGetResponse</a>

Methods:

- <code title="post /browsers">client.Browsers.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserNewParams">BrowserNewParams</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserNewResponse">BrowserNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /browsers/{id}">client.Browsers.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserGetResponse">BrowserGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
