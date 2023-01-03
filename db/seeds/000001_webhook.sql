-- cleanup
DELETE FROM webhooks WHERE id = 'wh_demo';
-- insert a new one
INSERT INTO webhooks (workspace_id, id, name, created_at, updated_at) VALUES ('ws_default', 'wh_demo', 'demo webhook', 1645488000000000000, 0);

-- cleanup
DELETE FROM webhook_tokens WHERE id = 'wht_simple';
-- insert a new one
INSERT INTO webhook_tokens (webhook_id, id, token, created_at) VALUES ('wh_demo', 'wht_simple','notsimpleasyouthought', 1645488000000000000);
