-- PROPOSED migration (NOT applied): nullable observability columns for
-- gateway tracing identifiers. They are for correlation with logs only
-- and are never used for authentication or as keys.
--
-- No primary keys change, no idempotency handling is added. Apply after
-- review, e.g. via a beego migration file in database/migrations/ with
-- these statements in Up() and DROP COLUMN in Down().
--
-- ID format: 8-64 chars from [A-Za-z0-9_-] (validated by
-- tracing.SanitizeID); session_id may be empty when the gateway did not
-- send one.

ALTER TABLE `customers`
    ADD COLUMN `trace_id`    VARCHAR(64) NULL DEFAULT NULL,
    ADD COLUMN `request_id`  VARCHAR(64) NULL DEFAULT NULL,
    ADD COLUMN `session_id`  VARCHAR(64) NULL DEFAULT NULL;

ALTER TABLE `auth_users`
    ADD COLUMN `trace_id`    VARCHAR(64) NULL DEFAULT NULL,
    ADD COLUMN `request_id`  VARCHAR(64) NULL DEFAULT NULL,
    ADD COLUMN `session_id`  VARCHAR(64) NULL DEFAULT NULL;

ALTER TABLE `accounts`
    ADD COLUMN `trace_id`    VARCHAR(64) NULL DEFAULT NULL,
    ADD COLUMN `request_id`  VARCHAR(64) NULL DEFAULT NULL,
    ADD COLUMN `session_id`  VARCHAR(64) NULL DEFAULT NULL;

ALTER TABLE `shops`
    ADD COLUMN `trace_id`    VARCHAR(64) NULL DEFAULT NULL,
    ADD COLUMN `request_id`  VARCHAR(64) NULL DEFAULT NULL,
    ADD COLUMN `session_id`  VARCHAR(64) NULL DEFAULT NULL;
