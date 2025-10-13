ALTER TABLE evidence
    ADD CONSTRAINT evidence_text_notes_check CHECK (array_length(text_notes, 1) <= 100),
    ADD CONSTRAINT evidence_audio_notes_check CHECK (array_length(audio_notes, 1) <= 100),
    ADD CONSTRAINT evidence_photos_check CHECK (array_length(photos, 1) <= 100),
    ADD CONSTRAINT evidence_evps_check CHECK (array_length(evps, 1) <= 100),
    ADD CONSTRAINT evidence_visibility_check CHECK (Visibility IN (TRUE, FALSE));
