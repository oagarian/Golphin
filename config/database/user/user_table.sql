CREATE EXTENSION IF NOT EXISTS  "uuid-ossp";
CREATE TABLE IF NOT EXISTS users {
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(16) NOT NULL UNIQUE,
    email VARCHAR(128) UNIQUE NOT NULL,
    password VARCHAR(128) NOT NULL
};