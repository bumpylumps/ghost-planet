CREATE TABLE IF NOT EXISTS evidence (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    created_by_user_id bigint,
    text_notes jsonb NOT NULL DEFAULT '[]',
    audio_notes jsonb NOT NULL DEFAULT '[]', 
    photos jsonb NOT NULL DEFAULT '[]', 
    evps jsonb NOT NULL DEFAULT '[]', 
    visibility boolean NOT NULL,
    version integer NOT NULL DEFAULT 1
);