create table if not exists tech_train_ca_go.rare_rate (
    id int not null auto_increment,
    rarity int,
    probability float,
    PRIMARY KEY (id)
);