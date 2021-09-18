CREATE TABLE customer
(
    customer_id INTEGER PRIMARY KEY,
    name NCHAR(20) NOT NULL,
    address NCHAR VARYING(20),
    postal_code CHAR(8) REFERENCES area,
    phone_number CHAR(12),
    mail_address VARCHAR(30) UNIQUE
);

CREATE TABLE area
(
    postal_code CHAR(8) PRIMARY KEY,
    country CHAR(12)
);

CREATE VIEW japanese_customer
(
    customer_id,
    name
)
AS
    (
    SELECT customer_id, name
    FROM customer INNER JOIN area ON customer.postal_code = area.postal_code
    WHERE
        area.country = 'japan'
);

INSERT INTO
    customer
    (
    customer_id,
    name,
    address,
    postal_code,
    phone_number,
    mail_address
    )
VALUES
    (
        '0000',
        'name_00',
        'address_00',
        '00000000',
        '000-0000-000',
        'mail@foo.com'
    );

INSERT INTO
    area
    (postal_code, country)
VALUES
    ('00000000', 'japan');

SELECT
    *
FROM
    japanese_customer;