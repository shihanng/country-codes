ALTER TABLE countries
    ADD COLUMN short_name text;

ALTER TABLE countries
    ADD COLUMN short_name_lower_case text;

ALTER TABLE countries
    ADD COLUMN full_name text;

ALTER TABLE countries
    ADD COLUMN alpha_3_code text;

ALTER TABLE countries
    ADD COLUMN numeric_code text;

ALTER TABLE countries
    ADD COLUMN remarks text;

ALTER TABLE countries
    ADD COLUMN independent text;

ALTER TABLE countries
    ADD COLUMN territory_name text;

ALTER TABLE countries
    ADD COLUMN status text;

ALTER TABLE countries RENAME COLUMN english_sort_name TO english_short_name;
