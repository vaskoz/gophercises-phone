use phones;

CREATE TABLE phone (number VARCHAR(100) NOT NULL, PRIMARY KEY (number));

INSERT INTO phone VALUES("1234567890");
INSERT INTO phone VALUES("123 456 7891");
INSERT INTO phone VALUES("(123) 456 7892");
INSERT INTO phone VALUES("(123) 456-7893");
INSERT INTO phone VALUES("123-456-7894");
INSERT INTO phone VALUES("123-456-7890");
INSERT INTO phone VALUES("1234567892");
INSERT INTO phone VALUES("(123)456-7892");

