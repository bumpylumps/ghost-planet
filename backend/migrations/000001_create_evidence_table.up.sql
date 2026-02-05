CREATE TABLE IF NOT EXISTS evidence (
    id bigserial PRIMARY KEY,
    investigation_id bigint NOT NULL, 
    location_id bigint NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    created_by_user_id bigint,
    visibility boolean NOT NULL,
    version integer NOT NULL DEFAULT 1
);