CREATE TABLE IF NOT EXISTS tech_train_ca_go.users (
    id varchar(36) not null,
    name varchar(255),
    reg_at datetime,
    update_at datetime,
    PRIMARY KEY (id)
);