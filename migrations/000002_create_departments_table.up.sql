-- Create a `departments` table.
CREATE TABLE IF NOT EXISTS departments (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    manager VARCHAR(100) NOT NULL
);
