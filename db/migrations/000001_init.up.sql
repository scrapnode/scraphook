CREATE TABLE IF NOT EXISTS workspaces (
    user_id VARCHAR(64) NOT NULL,
    id VARCHAR(64) NOT NULL,
    name VARCHAR(256) NOT NULL,
    created_at BIGINT DEFAULT 0,
    updated_at BIGINT DEFAULT 0,
    deleted_at BIGINT DEFAULT 0,
    CONSTRAINT scraphook_workspaces_pk PRIMARY KEY (id)
);
CREATE INDEX scraphook_workspaces_idx_user_id ON workspaces (user_id);
CREATE INDEX scraphook_workspaces_idx_deleted_at ON workspaces (deleted_at);

CREATE TABLE IF NOT EXISTS webhooks (
    workspace_id VARCHAR(64) NOT NULL,
    id VARCHAR(64) NOT NULL,
    name VARCHAR(256) NOT NULL,
    created_at BIGINT DEFAULT 0,
    updated_at BIGINT DEFAULT 0,
    deleted_at BIGINT DEFAULT 0,
    CONSTRAINT scraphook_webhooks_pk PRIMARY KEY (id)
);
CREATE INDEX scraphook_webhooks_idx_user_id ON workspaces (user_id);
CREATE INDEX scraphook_webhooks_idx_deleted_at ON workspaces (deleted_at);


