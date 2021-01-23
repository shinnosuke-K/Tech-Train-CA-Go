CREATE TABLE IF NOT EXISTS tech_train_ca_go.characters (
    chara_id int not null auto_increment,
    chara_name varchar(255),
    reg_at datetime,
    rarity int,
    PRIMARY KEY (chara_id)
);