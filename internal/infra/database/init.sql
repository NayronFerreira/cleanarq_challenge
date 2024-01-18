CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(255) PRIMARY KEY,
    price INT NOT NULL,
    tax INT NOT NULL,
    final_price INT NOT NULL
);

INSERT INTO orders (id, price, tax, final_price) VALUES
('0afd4e8e-b7da-483d-b4db-b26395245417', 2000, 2000, 4000),
('182a4ace-d4d2-4098-8481-b692a713a06b', 4000, 2000, 6000),
('bbc63cad-9f9d-403b-a5e0-703b3467b6f9', 4000, 3000, 7000),
('bde5d595-9671-48e8-9edf-31a8f5cabbde', 4000, 3000, 7000);