CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE category (
    id UUID       PRIMARY KEY DEFAULT uuid_generate_v4(),
    category_name VARCHAR(25) NOT NULL,
    created_at    TIMESTAMP DEFAULT NOW(),
    created_by    UUID NOT NULL,
    update_at     TIMESTAMP NULL,
    update_by     UUID NULL
);

CREATE TABLE orders (
    id UUID      PRIMARY KEY DEFAULT uuid_generate_v4() ,
    customer_id  UUID NOT NULL,
    total_price  INT NOT NULL,
    created_at   TIMESTAMP DEFAULT NOW(),
    created_by   UUID NOT NULL,
    update_at    TIMESTAMP NULL,
    update_by    UUID NULL
);

CREATE TABLE product (
    id           UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    id_category  UUID REFERENCES category (id),
    description  TEXT,
    price        INT,
    stock        INT,
    product_name VARCHAR(40),
    created_at   TIMESTAMP DEFAULT NOW(),
    created_by   UUID NOT NULL,
    update_at    TIMESTAMP NULL,
    update_by    UUID NULL
);

CREATE TABLE detail_order (
    id UUID      PRIMARY KEY DEFAULT uuid_generate_v4(),
    id_order     UUID NOT NULL REFERENCES orders (id),
    product_id   UUID NOT NULL REFERENCES product (id),
    quantity     INT NOT NULL,
    sub_total    INT NOT NULL,
    created_at   TIMESTAMP DEFAULT NOW(),
    created_by   UUID NOT NULL,
    update_at    TIMESTAMP NULL,
    update_by    UUID NULL
);

CREATE TABLE wishlist (
    id UUID      PRIMARY KEY DEFAULT uuid_generate_v4(),
    customer_id  UUID NOT NULL,
    product_id   UUID NOT NULL,
    created_at   TIMESTAMP DEFAULT NOW(),
    created_by   UUID NOT NULL,
    update_at    TIMESTAMP NULL,
    update_by    UUID NULL
);