-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `users` (
    user_id varchar(32) not null primary key,
    token varchar(255),
    user_name varchar(255),
    reg_time_jst datetime,
    update_time_jst datetime
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table users;
