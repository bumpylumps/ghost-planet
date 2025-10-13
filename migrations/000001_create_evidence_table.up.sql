CREATE TABLE IF NOT EXISTS evidence (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    text_notes text[],
    audio_notes text[], 
    photos text[], 
    evps text[], 
    Visibility bool,
    version integer NOT NULL DEFAULT 1
);