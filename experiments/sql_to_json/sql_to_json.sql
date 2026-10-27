DROP TABLE IF EXISTS customer;

CREATE TABLE customer (
    id INTEGER PRIMARY KEY,
    name TEXT,
    address TEXT,
    city TEXT,
    zip TEXT
);

INSERT INTO customer (name, address, city, zip) values ('Zed', '1010 Blank Lane', 'San Francisco, CA', '94102');
INSERT INTO customer (name, address, city, zip) values ('Joe', '345 Blank Lane', 'San Francisco, CA', '94102');
INSERT INTO customer (name, address, city, zip) values ('Mary', '9823 Blank Lane', 'New York, NY', '07020');
INSERT INTO customer (name, address, city, zip) values ('Alfred', '2222 Blank Lane', 'New York, NY', '07020');
INSERT INTO customer (name, address, city, zip) values ('Gonzo', '874 Blank Lane', 'Salt Lake City, UT', '84044');
INSERT INTO customer (name, address, city, zip) values ('Blaine', '3488 Blank Lane', 'Salt Lake City, UT', '84044');
INSERT INTO customer (name, address, city, zip) values ('Justine', '010102 Blank Lane', 'Chicago, IL', '60007');
INSERT INTO customer (name, address, city, zip) values ('Alberta', '65 Blank Lane', 'Chicago, IL', '60007');
