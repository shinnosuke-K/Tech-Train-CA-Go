create table if not exists tech_train_ca_go.gachas (
    id int not null auto_increment,
    rarity int,
    weights int,
    PRIMARY KEY (id)
);

insert into tech_train_ca_go.gachas(rarity, weights)
values (5, 10),
       (4, 20),
       (3, 50),
       (2, 75),
       (1, 100)