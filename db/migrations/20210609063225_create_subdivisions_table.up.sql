CREATE TABLE subdivisions (
    country_code text NOT NULL
    , code_31662 text NOT NULL
    , name text
    , local_variant text
    , language_code text
    , romanization_system text
    , parent_subdivision text
    , PRIMARY KEY (country_code , code_31662)
    , FOREIGN KEY (country_code) REFERENCES countries (alpha_2_code) ON DELETE CASCADE
    , FOREIGN KEY (language_code) REFERENCES languages (alpha_2_code) ON DELETE CASCADE
);
