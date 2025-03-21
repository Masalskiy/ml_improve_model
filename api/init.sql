CREATE TABLE IF NOT EXISTS products (
    name VARCHAR(255),
    scu VARCHAR(50),
    link VARCHAR(255),
    image_link VARCHAR(255),
    description TEXT,
    id_product BIGINT PRIMARY KEY
);

-- Добавим тестовые данные
INSERT INTO products (name, scu, link, image_link, description, id_product) VALUES
    ('Тестовый продукт 1', 'SKU001', 'http://example.com/1', 'http://example.com/image1.jpg', 'Описание продукта 1', 1),
    ('Тестовый продукт 2', 'SKU002', 'http://example.com/2', 'http://example.com/image2.jpg', 'Описание продукта 2', 2); 