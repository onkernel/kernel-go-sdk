# Shared Response Types

- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk/shared#AppAction">AppAction</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk/shared#ErrorDetail">ErrorDetail</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk/shared#ErrorEvent">ErrorEvent</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk/shared#ErrorModel">ErrorModel</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk/shared#HeartbeatEvent">HeartbeatEvent</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk/shared#LogEvent">LogEvent</a>

# Deployments

Response Types:

- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#DeploymentStateEvent">DeploymentStateEvent</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#DeploymentNewResponse">DeploymentNewResponse</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#DeploymentGetResponse">DeploymentGetResponse</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#DeploymentListResponse">DeploymentListResponse</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#DeploymentFollowResponseUnion">DeploymentFollowResponseUnion</a>

Methods:

- <code title="post /deployments">client.Deployments.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#DeploymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#DeploymentNewParams">DeploymentNewParams</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#DeploymentNewResponse">DeploymentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /deployments/{id}">client.Deployments.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#DeploymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#DeploymentGetResponse">DeploymentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /deployments">client.Deployments.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#DeploymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#DeploymentListParams">DeploymentListParams</a>) ([]<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#DeploymentListResponse">DeploymentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /deployments/{id}/events">client.Deployments.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#DeploymentService.Follow">Follow</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#DeploymentFollowParams">DeploymentFollowParams</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#DeploymentFollowResponseUnion">DeploymentFollowResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

# Invocations

Response Types:

- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#InvocationStateEvent">InvocationStateEvent</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#InvocationNewResponse">InvocationNewResponse</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#InvocationGetResponse">InvocationGetResponse</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#InvocationUpdateResponse">InvocationUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#InvocationFollowResponseUnion">InvocationFollowResponseUnion</a>

Methods:

- <code title="post /invocations">client.Invocations.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#InvocationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#InvocationNewParams">InvocationNewParams</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#InvocationNewResponse">InvocationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /invocations/{id}">client.Invocations.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#InvocationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#InvocationGetResponse">InvocationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /invocations/{id}">client.Invocations.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#InvocationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#InvocationUpdateParams">InvocationUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#InvocationUpdateResponse">InvocationUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /invocations/{id}/browsers">client.Invocations.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#InvocationService.DeleteBrowsers">DeleteBrowsers</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /invocations/{id}/events">client.Invocations.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#InvocationService.Follow">Follow</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#InvocationFollowResponseUnion">InvocationFollowResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Browsers

Params Types:

- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserPersistenceParam">BrowserPersistenceParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserPersistence">BrowserPersistence</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserNewResponse">BrowserNewResponse</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserGetResponse">BrowserGetResponse</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserListResponse">BrowserListResponse</a>

Methods:

- <code title="post /browsers">client.Browsers.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserNewParams">BrowserNewParams</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserNewResponse">BrowserNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /browsers/{id}">client.Browsers.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserGetResponse">BrowserGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /browsers">client.Browsers.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserListResponse">BrowserListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /browsers">client.Browsers.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserDeleteParams">BrowserDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /browsers/{id}">client.Browsers.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserService.DeleteByID">DeleteByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Replays

Response Types:

- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserReplayListResponse">BrowserReplayListResponse</a>
- <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserReplayStartResponse">BrowserReplayStartResponse</a>

Methods:

- <code title="get /browsers/{id}/replays">client.Browsers.Replays.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserReplayService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserReplayListResponse">BrowserReplayListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /browsers/{id}/replays/{replay_id}">client.Browsers.Replays.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserReplayService.Download">Download</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, replayID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserReplayDownloadParams">BrowserReplayDownloadParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /browsers/{id}/replays">client.Browsers.Replays.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserReplayService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserReplayStartParams">BrowserReplayStartParams</a>) (<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserReplayStartResponse">BrowserReplayStartResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /browsers/{id}/replays/{replay_id}/stop">client.Browsers.Replays.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserReplayService.Stop">Stop</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, replayID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk">kernel</a>.<a href="https://pkg.go.dev/github.com/onkernel/kernel-go-sdk#BrowserReplayStopParams">BrowserReplayStopParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
