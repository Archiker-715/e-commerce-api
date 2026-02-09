CREATE TABLE users_cart (
    user_id UUID NOT NULL,
    product_id BIGINT NOT NULL,
    name VARCHAR(500) NOT NULL,
    count BIGINT NOT NULL,
    unit_price BIGINT NOT NULL,
    total_price BIGINT NOT NULL,
    CONSTRAINT fk_product
        FOREIGN KEY (product_id)
        REFERENCES product (product_id)
        ON DELETE CASCADE
);

CREATE FUNCTION check_product_quantity()
RETURN TRIGGER AS $$
BEGIN
    IF NEW.count > (
        SELECT count FROM products WHERE product_id = NEW.product_id
    ) THEN 
        RAISE EXCEPTION 'Нельзя добавить больше товаров, чем есть в наличии';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_check_quantity
BEFORE INSERT OR UPDATE ON users_cart
FOR EACH ROW
EXECUTE FUNCTION check_product_quantity();