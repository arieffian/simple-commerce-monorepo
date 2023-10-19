CREATE UNIQUE INDEX IF NOT EXISTS products_slug_unique_idx ON products ( slug );
CREATE UNIQUE INDEX IF NOT EXISTS products_sku_unique_idx ON products ( sku );
CREATE INDEX IF NOT EXISTS products_status_idx ON products ( status );