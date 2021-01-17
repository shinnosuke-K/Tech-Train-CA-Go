CREATE TABLE IF NOT EXISTS tech_train_ca_go.users (
    user_id varchar(32) not null primary key,
    user_name varchar(255),
    reg_at datetime,
    update_at datetime
);