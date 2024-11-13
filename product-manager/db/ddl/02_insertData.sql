/*
 -- Query: SELECT * FROM yorushika_product.products
 LIMIT 0, 1000
 
 -- Date: 2024-11-06 20:58
 */
INSERT INTO
  `products` (
    `id`,
    `created_at`,
    `updated_at`,
    `deleted_at`,
    `text`,
    `name`,
    `price`,
    `img`
  )
VALUES
  (
    1,
    '2024-07-28 18:04:57.576',
    '2024-07-28 18:04:57.576',
    NULL,
    'テスト',
    'サンプル',
    1000,
    '画像'
  );