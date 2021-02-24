create table if not exists tech_train_ca_go.possessions (
    id varchar(36) not null,
    user_id varchar(36) not null,
    chara_id varchar(255) not null,
    reg_at datetime,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES tech_train_ca_go.users(id),
    INDEX user_index(user_id)
);

