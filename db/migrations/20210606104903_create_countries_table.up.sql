CREATE TABLE IF NOT EXISTS countries (
    alpha_2_code text PRIMARY KEY NOT NULL CHECK(alpha_2_code <> ''),
    english_sort_name text NOT NULL
);
