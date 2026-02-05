CREATE TABLE IF NOT EXISTS evidence_textnotes (
    id bigserial PRIMARY KEY,
    evidence_id bigint NOT NULL REFERENCES evidence(id) ON DELETE CASCADE,
    subject text,
    body text NOT NULL,
    location_id bigint,
    created_at timestamptz NOT NULL DEFAULT NOW() 
)