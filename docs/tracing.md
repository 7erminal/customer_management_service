# Request tracing

`LOG_LEVEL` selects debug, info (default), warn or error for application logs.
Access logs always emit at INFO, regardless of that setting. Logs are JSON on
stdout. Beego's file logger, request debug output, and ORM SQL logging are not
used. Startup/background logs use empty correlation IDs when no request exists.
The local `internal/trace` package is a small adapted subset of the gateway's
`github.com/msys/go-common/trace`, gateway commit
`a79a370ee178066e3f5a163ce69f0813cb6b249c`. Replace it with the published
module when available; its recovery/access helpers need a privacy and contract
review first (its panic value and raw path logging are not suitable here).

A BeforeRouter filter propagates validated IDs, while an outer handler also
covers static resources and requests rejected before that filter. To avoid PII
and tokens in path parameters, the access log's `path` and `route` are the
matched route template, or `<unmatched>` if none is known; neither is the raw
URL. A panic after headers/body have already been written cannot reliably change
the HTTP status or replace the body. The recovery log still records the stack,
but not the panic value. ID generation uses crypto/rand; no trace IDs are used
for authorization, routing, database keys, or idempotency.

Example:

```json
{"time":"2026-11-27T12:00:00Z","level":"INFO","msg":"access","service":"customer_management_service","session_id":"session_123","trace_id":"trace_1234","request_id":"request_123","method":"GET","path":"/v1/customers/:id","route":"/v1/customers/:id","status":200,"duration_ms":12,"bytes":80}
```

## Compatibility and remaining work

* Existing error bodies and HTTP statuses are unchanged. The new helper and
  panic envelope include `error.request_id`; legacy errors do not necessarily
  include it, but all responses carry the request and trace headers. Some
  clients parse `StatusCode`/`StatusDesc` in the body, including when HTTP is
  200. Bare string error responses also exist. Changing those shapes needs a
  separately approved client/gateway compatibility decision.
* The unsafe legacy Beego log calls are **suppressed, not migrated**. This
  avoids leaking passwords, tokens, query data and personal information, but
  means their operational events are currently unavailable. Active non-migrated
  calls remain in `controllers/users.go`, `controllers/customers.go`,
  `controllers/customer_emergency_contacts.go`, `controllers/shops.go`,
  `controllers/customer_credentials.go`, `controllers/customer_guarantors.go`,
  `controllers/functions/register_account.go`, `controllers/functions/helpers.go`,
  `controllers/roles.go`, `controllers/role_permissions.go`,
  `controllers/userExtraDetails.go`, `controllers/newsletter_customers.go`,
  `controllers/identification_types.go`, `models/customers.go`, `models/users.go`,
  `models/userinvites.go`, and `models/currencies.go`. Migrate these individually
  to static, allowlisted events with request context before re-enabling them.
  Beego framework and standard-library logs are suppressed too; startup/transport diagnostics need a
  safe structured replacement. ORM debug output is discarded.
* Context-aware ORM operations cover the main individual customer/user/account
  create, lookup, update and invitation operations and their lookup helpers.
  List/search/count and delete operations (`GetAll*`, `Get*Count`,
  `DeleteCustomer`, `DeleteUsers`, `DeleteAccounts`), invite list searches,
  and other controllers/models remain on background-context ORM methods.
  The existing compatibility wrappers also use background contexts for any
  callers not yet migrated. Extend these in follow-up work.
* SMTP invitation delivery is an in-process goroutine, not a queue. Its context
  retains the IDs after the HTTP request ends; no IDs are added to email content.
  No queue/cache publisher or consumer was found.
* Docker and the CI runner target Go 1.22; the deploy workflow builds on a
  remote host, whose Go installation must also be upgraded separately.

## Proposed schema change — not applied

After verifying the live schema and agreeing retention/ownership, add these
nullable observation-only columns to `customers`, `auth_users`, and `accounts`:

```sql
ALTER TABLE customers
  ADD COLUMN trace_id VARCHAR(64) NULL,
  ADD COLUMN request_id VARCHAR(64) NULL,
  ADD COLUMN session_id VARCHAR(64) NULL;
ALTER TABLE auth_users
  ADD COLUMN trace_id VARCHAR(64) NULL,
  ADD COLUMN request_id VARCHAR(64) NULL,
  ADD COLUMN session_id VARCHAR(64) NULL;
ALTER TABLE accounts
  ADD COLUMN trace_id VARCHAR(64) NULL,
  ADD COLUMN request_id VARCHAR(64) NULL,
  ADD COLUMN session_id VARCHAR(64) NULL;
```

This is a proposal, not a migration: the application does not write these
columns. Decide whether they record creation or last update before implementing
persistence. No primary key or idempotency change is proposed.
