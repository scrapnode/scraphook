CREATE TABLE IF NOT EXISTS workspaces (
    user_id VARCHAR(64) NOT NULL,
    id VARCHAR(64) NOT NULL,
    name VARCHAR(256) NOT NULL,
    created_at BIGINT DEFAULT 0,
    updated_at BIGINT DEFAULT 0,
    CONSTRAINT workspaces_pk PRIMARY KEY (id)
);
CREATE INDEX workspaces_idx_user_id ON workspaces (user_id DESC);

CREATE TABLE IF NOT EXISTS webhooks (
    workspace_id VARCHAR(64) NOT NULL,
    id VARCHAR(64) NOT NULL,
    name VARCHAR(256) NOT NULL,
    created_at BIGINT DEFAULT 0,
    updated_at BIGINT DEFAULT 0,
    CONSTRAINT webhooks_pk PRIMARY KEY (id)
);
CREATE INDEX webhooks_idx_workspace_id ON webhooks (workspace_id DESC);

CREATE TABLE IF NOT EXISTS webhook_tokens (
    webhook_id VARCHAR(64) NOT NULL,
    id VARCHAR(64) NOT NULL,
    token VARCHAR(256) NOT NULL,
    created_at BIGINT DEFAULT 0,
    CONSTRAINT webhook_tokens_pk PRIMARY KEY (id)
);
CREATE INDEX webhooks_idx_webhook_id ON webhook_tokens (webhook_id DESC);

CREATE TABLE IF NOT EXISTS endpoints (
    workspace_id VARCHAR(64) NOT NULL,
    webhook_id VARCHAR(64) NOT NULL,
    id VARCHAR(64) NOT NULL,
    name VARCHAR(256) NOT NULL,
    uri VARCHAR(1024) NOT NULL,
    created_at BIGINT DEFAULT 0,
    updated_at BIGINT DEFAULT 0,
    CONSTRAINT endpoints_pk PRIMARY KEY (id)
);
CREATE INDEX endpoints_idx_workspace_webhook_ids ON endpoints (workspace_id DESC, webhook_id DESC);

CREATE TABLE IF NOT EXISTS endpoint_rules (
    endpoint_id VARCHAR(64) NOT NULL,
    id VARCHAR(64) NOT NULL,
    rule VARCHAR(2048) NOT NULL,
    negative BOOLEAN DEFAULT FALSE,
    priority INT DEFAULT 0,
    created_at BIGINT DEFAULT 0,
    updated_at BIGINT DEFAULT 0,
    CONSTRAINT endpoint_rules_pk PRIMARY KEY (id)
);
CREATE INDEX endpoint_rules_idx_endpoint_id ON endpoint_rules (endpoint_id DESC);
