create table if not exists tech_train_ca_go.rare_rate (
    id int not null auto_increment,
    rarity int,
    probability float,
    PRIMARY KEY (id)
);

insert into tech_train_ca_go.rare_rate(rarity, probability)
values (5, 2.5),
       (4, 7.5),
       (3, 10),
       (2, 20),
       (1, 60)