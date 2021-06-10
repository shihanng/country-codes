ALTER TABLE countries
    DROP COLUMN short_name;

ALTER TABLE countries
    DROP COLUMN short_name_lower_case;

ALTER TABLE countries
    DROP COLUMN full_name;

ALTER TABLE countries
    DROP COLUMN alpha_3_code;

ALTER TABLE countries
    DROP COLUMN numeric_code;

ALTER TABLE countries
    DROP COLUMN remarks;

ALTER TABLE countries
    DROP COLUMN independent;

ALTER TABLE countries
    DROP COLUMN territory_name;

ALTER TABLE countries
    DROP COLUMN status;

ALTER TABLE countries RENAME COLUMN english_short_name TO english_sort_name;
