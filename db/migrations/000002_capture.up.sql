CREATE TABLE IF NOT EXISTS messages (
    timestamps bigint DEFAULT 0,
    bucket VARCHAR(64) NOT NULL,
    workspace_id VARCHAR(64) NOT NULL,
    webhook_id VARCHAR(64) NOT NULL,
    id VARCHAR(64) NOT NULL,
    headers TEXT,
    body TEXT,
    method VARCHAR(64),
    CONSTRAINT messages_pk PRIMARY KEY (id)
);
CREATE INDEX messages_idx_webhook_id ON messages (webhook_id DESC);
