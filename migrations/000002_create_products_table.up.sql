-- Create a `products` table.
CREATE TABLE IF NOT EXISTS products (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    supplier VARCHAR(100) NOT NULL,
    department_id INTEGER NULL,
    price DECIMAL(10,2) NULL,
    FOREIGN KEY (department_id)
        REFERENCES departments(id)
        ON DELETE SET NULL
);
