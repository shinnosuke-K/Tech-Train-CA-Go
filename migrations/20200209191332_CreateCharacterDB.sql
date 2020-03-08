
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table if not exists characters (
    chara_id varchar(20) not null primary key,
    chara_name varchar(255),
    reg_time_jst datetime,
    rare int
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table  characters;
