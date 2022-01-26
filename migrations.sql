CREATE TABLE products(
    id          serial primary key,
    sku         varchar(255) not null,
    name        varchar(255) not null,
    category    varchar(255),
    price       int not null
);


INSERT INTO products (sku, name, category, price)
VALUES
        ('000001','BV Lean leather ankle boots','boots',89000),
        ('000002','BV Lean leather ankle boots','boots',99000),
        ('000003','Ashlington leather ankle boots','boots',71000),
        ('000004','Naima embellished suede sandals','sandals',79500),
        ('000005','Nathane leather sneakers','sneakers',59000);