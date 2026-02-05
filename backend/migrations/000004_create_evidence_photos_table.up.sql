CREATE TABLE evidence_photos (
    id bigserial PRIMARY KEY,
    evidence_id bigint NOT NULL REFERENCES evidence(id) ON DELETE CASCADE,
    source_url text NOT NULL,
    thumbnail_url text,
    caption text,
    file_type varchar(10),
    file_size_bytes bigint,
    created_at timestamptz NOT NULL DEFAULT NOW()
)