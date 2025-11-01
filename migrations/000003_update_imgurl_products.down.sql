-- Reverse the rename (for rollback)
ALTER TABLE products RENAME COLUMN img_url TO imgurl;