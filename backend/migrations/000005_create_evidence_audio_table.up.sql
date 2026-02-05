CREATE TABLE IF NOT EXISTS evidence_audio (
    id bigserial PRIMARY key,
    evidence_id bigint NOT NULL REFERENCES evidence(id) ON DELETE CASCADE,
    title text,
    source_url text NOT NULL,
    duration interval,
    file_size_bytes bigint,
    is_evp boolean DEFAULT false,
    created_at timestamptz NOT NULL DEFAULT NOW()
)