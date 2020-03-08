
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table if not exists possessions (
    posse_id varchar(20) not null primary key,
    user_id varchar(20) not null,
    chara_id varchar(20) not null,
    reg_time_jst datetime
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table possessions;
