# SQL 

## Playground

* [paiza.io](https://paiza.io/en/languages/mysql)

## SQL samples

### Create Table and View

```sql
CREATE TABLE area (
  postal_code CHAR(8) PRIMARY KEY,
  country CHAR(12)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

CREATE TABLE customer (
  customer_id INTEGER PRIMARY KEY,
  name NCHAR(20) NOT NULL,
  address NCHAR VARYING(20),
  postal_code CHAR(8),
  phone_number CHAR(12),
  mail_address VARCHAR(30) UNIQUE,
  age INT,
    FOREIGN KEY (postal_code)
    REFERENCES area (postal_code)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=INNODB DEFAULT CHARSET=utf8;

CREATE VIEW japanese_customer (customer_id, name) AS (
  SELECT customer_id, name
  FROM customer
    INNER JOIN area
      ON customer.postal_code=area.postal_code
  WHERE area.country='japan'
);

/**
 * When you want to insert some sqls, BULK INSERT is used generally.
 * But exam considers basic sql only, so I write simple query.
 */
INSERT INTO area (postal_code, country) VALUES ('postal_0', 'japan');
INSERT INTO area (postal_code, country) VALUES ('postal_1', 'japan');
INSERT INTO area (postal_code, country) VALUES ('postal_2', 'japan');
INSERT INTO area (postal_code, country) VALUES ('postal_3', 'usa');
INSERT INTO area (postal_code, country) VALUES ('postal_4', 'usa');
INSERT INTO customer (
  customer_id, name, address, postal_code, phone_number, mail_address, age
) VALUES (
  '0000', 'name_00', 'address_00', 'postal_0', '000-0000-000', '00@foo.com', 0
);
INSERT INTO customer (
  customer_id, name, address, postal_code, phone_number, mail_address, age
) VALUES (
  '0001', 'name_01', 'address_01', 'postal_1', '000-0000-001', '01@foo.com', 1
);
INSERT INTO customer (
  customer_id, name, address, postal_code, phone_number, mail_address, age
) VALUES (
  '0002', 'name_02', 'address_02', 'postal_2', '000-0000-002', '02@foo.com', 2
);
INSERT INTO customer (
  customer_id, name, address, postal_code, phone_number, mail_address, age
) VALUES (
  '0003', 'name_03', 'address_03', 'postal_3', '000-0000-003', '03@foo.com', 3
);
INSERT INTO customer (
  customer_id, name, address, postal_code, phone_number, mail_address, age
) VALUES (
  '0004', 'name_04', 'address_04', 'postal_4', '000-0000-004', '04@foo.com', 4
);

/**
 * check Table
 */
SELECT '==============================' AS 'print customer table';
SELECT * FROM customer;
SELECT '==============================' AS 'print area table';
SELECT * FROM area;

/**
 * check VIEW
 */
SELECT '==============================' AS 'print japanese_customer view';
SELECT * FROM japanese_customer;

/**
 * check ON UPDATE CASCADE
 */
UPDATE area SET postal_code='postal_9' WHERE postal_code='postal_0';
SELECT '==============================' AS 'print updated customer table';
SELECT * FROM customer;
SELECT '==============================' AS 'print updated area table';
SELECT * FROM area;

/**
 * check ON DELETE CASCADE
 */
DELETE FROM area WHERE postal_code='postal_9';
SELECT '==============================' AS 'print deleted customer table';
SELECT * FROM customer;
SELECT '==============================' AS 'print deleted area table';
SELECT * FROM area;

/**
 * check INDEX
 */
CREATE INDEX name ON customer (name);
SELECT '==============================' AS 'use index, and print customer table';
SELECT * FROM customer;
SELECT '==============================' AS 'does not use index, and print area table';
SELECT * FROM area;

/**
 * check transaction
 */
SET TRANSACTION ISOLATION LEVEL READ COMMITTED;
START TRANSACTION READ WRITE;
UPDATE area SET country='canada' WHERE country='usa';
COMMIT;
SELECT '==============================' AS 'after commit, print area table';
SELECT * FROM area;

SET TRANSACTION ISOLATION LEVEL READ COMMITTED;
START TRANSACTION READ WRITE;
UPDATE area SET country='usa' WHERE country='canada';
ROLLBACK;
SELECT '==============================' AS 'after rollback, print area table';
SELECT * FROM area;

/**
 * check SELECT
 */
SELECT '==============================' AS 'print country and sum of age in japan';
SELECT customer_id, country, SUM(age) FROM customer
INNER JOIN area ON customer.postal_code=area.postal_code
WHERE area.country='japan' OR area.country='china'
GROUP BY customer.customer_id, area.country
  HAVING customer.customer_id='0001' OR customer.customer_id='0002'
  ORDER BY customer.customer_id DESC
```
