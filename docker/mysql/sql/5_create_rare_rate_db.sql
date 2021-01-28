create table if not exists tech_train_ca_go.gachas (
    id int not null auto_increment,
    rarity int,
    probability float,
    PRIMARY KEY (id)
);

insert into tech_train_ca_go.gachas(rarity, probability)
values (5, 2.5),
       (4, 7.5),
       (3, 10),
       (2, 20),
       (1, 60)