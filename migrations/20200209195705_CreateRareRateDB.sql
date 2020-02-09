
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table if not exists rare_rate (
    rare int,
    rate int
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table rare_rate;
