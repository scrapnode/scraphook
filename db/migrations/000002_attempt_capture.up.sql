CREATE TABLE IF NOT EXISTS messages
(
    timestamps   bigint DEFAULT 0,
    bucket       VARCHAR(64) NOT NULL,
    workspace_id VARCHAR(64) NOT NULL,
    webhook_id   VARCHAR(64) NOT NULL,
    id           VARCHAR(64) NOT NULL,
    headers      TEXT,
    body         TEXT,
    method       VARCHAR(64),
    CONSTRAINT messages_pk PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS requests
(
    timestamps   bigint DEFAULT 0,
    bucket       VARCHAR(64) NOT NULL,
    workspace_id VARCHAR(64) NOT NULL,
    webhook_id   VARCHAR(64) NOT NULL,
    endpoint_id  VARCHAR(64) NOT NULL,
    message_id   VARCHAR(64) NOT NULL,
    id           VARCHAR(64) NOT NULL,
    uri          TEXT,
    status       INT,
    headers      TEXT,
    body         TEXT,
    method       VARCHAR(64),
    CONSTRAINT requests_pk PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS responses
(
    timestamps   bigint DEFAULT 0,
    bucket       VARCHAR(64) NOT NULL,
    workspace_id VARCHAR(64) NOT NULL,
    webhook_id   VARCHAR(64) NOT NULL,
    endpoint_id  VARCHAR(64) NOT NULL,
    message_id   VARCHAR(64) NOT NULL,
    request_id   VARCHAR(64) NOT NULL,
    id           VARCHAR(64) NOT NULL,
    uri          TEXT,
    status       INT,
    headers      TEXT,
    body         TEXT,
    CONSTRAINT responses_pk PRIMARY KEY (id)
);
