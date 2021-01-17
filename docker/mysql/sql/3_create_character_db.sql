CREATE TABLE IF NOT EXISTS tech_train_ca_go.characters (
    chara_id varchar(20) not null primary key,
    chara_name varchar(255),
    reg_at datetime,
    rare int
);