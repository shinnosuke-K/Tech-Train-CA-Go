create table if not exists tech_train_ca_go.possessions (
    posse_id varchar(36) not null,
    user_id varchar(36) not null,
    chara_id varchar(255) not null,
    quantity int,
    reg_at datetime,
    update_at datetime,
    PRIMARY KEY (posse_id),
    FOREIGN KEY (user_id) REFERENCES tech_train_ca_go.users(user_id)
);

