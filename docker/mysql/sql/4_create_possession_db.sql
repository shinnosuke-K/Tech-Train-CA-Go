create table if not exists tech_train_ca_go.possessions (
    posse_id varchar(20) not null primary key,
    user_id varchar(20) not null,
    chara_id varchar(20) not null,
    reg_at datetime
);

