CREATE TABLE products(
    id          serial primary key,
    sku         varchar(255) not null,
    name        varchar(255) not null,
    category    varchar(255),
    price       int not null
);


INSERT INTO products (sku, name, category, price)
VALUES
        ('1','test1','cat1',100),
        ('2','test2','cat2',200);