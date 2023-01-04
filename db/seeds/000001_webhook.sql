-- cleanup
DELETE FROM webhooks WHERE id = 'wh_demo';
-- insert a new one
INSERT INTO webhooks (workspace_id, id, name, created_at, updated_at) VALUES ('ws_default', 'wh_demo', 'demo webhook', 1645488000000000000, 0);

-- cleanup
DELETE FROM webhook_tokens WHERE webhook_id = 'wh_demo';
-- insert a new one
INSERT INTO webhook_tokens (webhook_id, id, token, created_at) VALUES ('wh_demo', 'wht_simple','notsimpleasyouthought', 1645488000000000000);

-- cleanup
DELETE FROM endpoints WHERE webhook_id = 'wh_demo';
INSERT INTO endpoints (workspace_id, webhook_id, id, name, uri, created_at, updated_at) VALUES ('ws_default', 'wh_demo', 'ep_httpbinpost','httpbin.org POST', 'https://httpbin.org/post', 1645488000000000000, 0);
INSERT INTO endpoints (workspace_id, webhook_id, id, name, uri, created_at, updated_at) VALUES ('ws_default', 'wh_demo', 'ep_httpbinpatch','httpbin.org PATCH', 'https://httpbin.org/patch', 1645488000000000000, 0);

DELETE FROM endpoint_rules WHERE endpoint_id = 'ep_httpbinpost';
DELETE FROM endpoint_rules WHERE endpoint_id = 'ep_httpbinpatch';
INSERT INTO endpoint_rules (endpoint_id, id, rule, negative, priority, created_at, updated_at) VALUES ('ep_httpbinpost', 'epr_schedulepost', 'regex::wh_demo' , FALSE, 0, 1645488000000000000, 0);
INSERT INTO endpoint_rules (endpoint_id, id, rule, negative, priority, created_at, updated_at) VALUES ('ep_httpbinpatch', 'epr_schedulepatch', 'regex::wh_demo' , FALSE, 0, 1645488000000000000, 0);