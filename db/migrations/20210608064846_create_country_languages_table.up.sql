CREATE TABLE country_languages (
    country_code text NOT NULL
    , language_code text NOT NULL
    , local_short_name text
    , PRIMARY KEY (country_code , language_code)
    , FOREIGN KEY (country_code) REFERENCES countries (alpha_2_code) ON DELETE CASCADE
    , FOREIGN KEY (language_code) REFERENCES languages (alpha_2_code) ON DELETE CASCADE
);
