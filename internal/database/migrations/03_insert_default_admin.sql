-- +migrate Up

INSERT INTO admins (id, username, pass, created_at, is_active) VALUES ('8e52f943-5106-4622-b0c2-569f2f2b51a4', 'admin', '1fG5oBZCP8Xp', '2025-04-20T07:20:50.52Z', true);

-- +migrate Down

DELETE * FROM admins WHERE id = "8e52f943-5106-4622-b0c2-569f2f2b51a4";
