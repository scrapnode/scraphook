-- cleanup
DELETE FROM workspaces WHERE id = 'ws_default';
-- insert a new one
INSERT INTO workspaces (user_id, id, name, created_at, updated_at) VALUES ('ski_root', 'ws_default', 'default workspace', 1645488000000, 0);

-- cleanup
DELETE FROM webhooks WHERE id = 'wh_demo';
-- insert a new one
INSERT INTO webhooks (workspace_id, id, name, created_at, updated_at) VALUES ('ws_default', 'wh_demo', 'demo webhook', 1645488000000, 0);

-- cleanup
DELETE FROM webhook_tokens WHERE webhook_id = 'wh_demo';
-- insert a new one
INSERT INTO webhook_tokens (webhook_id, id, name, token, created_at) VALUES ('wh_demo', 'wht_simple', 'simple demo webhook token','notsimpleasyouthought', 1645488000000);

-- cleanup
DELETE FROM endpoints WHERE webhook_id = 'wh_demo';
INSERT INTO endpoints (workspace_id, webhook_id, id, name, uri, created_at, updated_at) VALUES ('ws_default', 'wh_demo', 'ep_httpbinpost','httpbin.org POST', 'https://httpbin.org/post', 1645488000000, 0);
INSERT INTO endpoints (workspace_id, webhook_id, id, name, uri, created_at, updated_at) VALUES ('ws_default', 'wh_demo', 'ep_httpbinput','httpbin.org PUT', 'https://httpbin.org/put', 1645488000000, 0);

DELETE FROM endpoint_rules WHERE endpoint_id = 'ep_httpbinpost';
DELETE FROM endpoint_rules WHERE endpoint_id = 'ep_httpbinput';
INSERT INTO endpoint_rules (endpoint_id, id, rule, negative, priority, created_at, updated_at) VALUES ('ep_httpbinpost', 'epr_schedulepost', 'regex::wh_demo' , FALSE, 0, 1645488000000, 0);
INSERT INTO endpoint_rules (endpoint_id, id, rule, negative, priority, created_at, updated_at) VALUES ('ep_httpbinput', 'epr_scheduleput', 'regex::wh_demo' , FALSE, 0, 1645488000000, 0);