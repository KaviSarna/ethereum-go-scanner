--
--    SPDX-License-Identifier: Apache-2.0
--

CREATE USER matic with password 'password';

CREATE DATABASE eth owner matic;
\c eth;
--

-- ----------------------------
--  Table structure for `transaction`
-- ----------------------------
--DROP TABLE IF EXISTS transaction;
CREATE TABLE transaction (
  id SERIAL PRIMARY KEY,
  txHash varchar(256) DEFAULT NULL,
  blockNo varchar(10) DEFAULT NULL,
  fromAdr varchar(256) DEFAULT NULL,
  toAdr varchar(256) DEFAULT NULL
);

ALTER table transaction owner to matic;

-- GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA PUBLIC to matic;
