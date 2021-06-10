CREATE TABLE IF NOT EXISTS languages (
    alpha_2_code text PRIMARY KEY NOT NULL CHECK (alpha_2_code <> '')
    , alpha_3_code text NOT NULL CHECK (alpha_3_code <> '')
);
