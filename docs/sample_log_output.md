# Sample log output (representative)

Structured JSON via `log/slog` to stdout and
`../logs/customer_management_application.log`. Level via `LOG_LEVEL` env
or `loglevel` in `conf/app.conf` (default `info`).

Every application line carries `service`, `session_id`, `trace_id`,
`request_id`.

## One access-log line per request

Inbound: `X-Session-Id: sess_1a2b3c4d`, `X-Trace-Id: trace_9f86d081`,
`X-Request-Id: req_2c26b46b`:

```json
{"time":"2026-11-27T16:39:12.488Z","level":"INFO","msg":"http_request","service":"customer_management_service","session_id":"sess_1a2b3c4d","trace_id":"trace_9f86d081","request_id":"req_2c26b46b","method":"POST","path":"/v1/customers/add-customer","route":"/v1/customers/add-customer","status":200,"duration_ms":42.317,"bytes":512}
```

Health check / direct call with no headers (IDs generated, trace falls
back to request ID):

```json
{"time":"2026-11-27T16:39:13.102Z","level":"INFO","msg":"http_request","service":"customer_management_service","session_id":"","trace_id":"9f86d081884c7d659a2feaa0c55ad015","request_id":"9f86d081884c7d659a2feaa0c55ad015","method":"GET","path":"/","route":"","status":200,"duration_ms":0.412,"bytes":0}
```

404 (early response — still one line, response carries
`X-Request-Id`/`X-Trace-Id`):

```json
{"time":"2026-11-27T16:39:14.010Z","level":"INFO","msg":"http_request","service":"customer_management_service","session_id":"","trace_id":"2c26b46b68ffc68ff99b453c1d304134","request_id":"2c26b46b68ffc68ff99b453c1d304134","method":"GET","path":"/v1/not-a-route","route":"","status":404,"duration_ms":0.198,"bytes":3}
```

## Business log line with IDs (and business ID once the record exists)

```json
{"time":"2026-11-27T16:39:12.455Z","level":"INFO","msg":"Customer added","service":"customer_management_service","session_id":"sess_1a2b3c4d","trace_id":"trace_9f86d081","request_id":"req_2c26b46b","customer_id":1234}
```

```json
{"time":"2026-11-27T16:39:12.470Z","level":"INFO","msg":"Generated customer number","service":"customer_management_service","session_id":"sess_1a2b3c4d","trace_id":"trace_9f86d081","request_id":"req_2c26b46b","customer_id":1234,"customer_number":"CUST001234"}
```

## Panic recovery (server-side only; client gets JSON, no stack)

```json
{"time":"2026-11-27T16:39:15.332Z","level":"ERROR","msg":"panic recovered","service":"customer_management_service","session_id":"sess_1a2b3c4d","trace_id":"trace_9f86d081","request_id":"req_2c26b46b","error":"runtime error: index out of range [5] with length 3","stack":"goroutine 42 [running]:\nruntime/debug.Stack()\n\t/usr/local/go/src/runtime/debug/stack.go:24 +0x64\nmain.(*CustomersController).Post(...)\n","method":"POST","path":"/v1/customers"}
```

```json
{"time":"2026-11-27T16:39:15.333Z","level":"INFO","msg":"http_request","service":"customer_management_service","session_id":"sess_1a2b3c4d","trace_id":"trace_9f86d081","request_id":"req_2c26b46b","method":"POST","path":"/v1/customers","route":"/v1/customers/:id","status":500,"duration_ms":3.891,"bytes":106}
```

Client response body for the panic:

```json
{"error":{"code":"internal_error","message":"internal server error","request_id":"req_2c26b46b"}}
```

## Background e-mail goroutine (IDs restored from the spawning request)

```json
{"time":"2026-11-27T16:39:20.777Z","level":"DEBUG","msg":"sending e-mail failed","service":"customer_management_service","session_id":"sess_1a2b3c4d","trace_id":"trace_9f86d081","request_id":"req_2c26b46b","error":"dial tcp: smtp.gmail.com:587: connect refused"}
```

## Notes

- Query-string values, request bodies, headers, cookies, tokens and
  passwords are never logged. Query strings are stripped (access log
  uses `r.URL.Path` only).
- `orm.Debug` (SQL with parameter values) is now off by default;
  `LOG_ORM_DEBUG=true` re-enables it when needed.
- Beego's built-in dev access line is suppressed (`logs.SetLevel`), so
  exactly one access-log line is emitted per request.
