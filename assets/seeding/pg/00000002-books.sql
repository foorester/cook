--SEED
INSERT INTO books (id, name, description, owner_id, created_at, updated_at)
VALUES ('7c679e0e-888c-4e50-9f5b-80e75b15f998', 'Book 1', 'First book owned by user 1', '0792b97b-4f88-42a8-a035-1d0aad0ae7f8', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

--SEED
INSERT INTO books (id, name, description, owner_id, created_at, updated_at)
VALUES ('2a8b0e19-63eb-4ebc-828a-7111e88fc54f', 'Book 2', 'Second book owned by user 1', '0792b97b-4f88-42a8-a035-1d0aad0ae7f8', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

--SEED
INSERT INTO books (id, name, description, owner_id, created_at, updated_at)
VALUES ('af1365d5-63c3-4465-908f-5b04a94d733b', 'Empty Book 1', 'An empty book owned by user 2', 'b1c20e60-ec1c-4fae-97b9-b4d0578b0123', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

--SEED
INSERT INTO books (id, name, description, owner_id, created_at, updated_at)
VALUES ('9c48153b-c9bb-4b8f-92bf-11bca18c7ad1', 'Empty Book 2', 'An empty book owned by user 3', '7d399e9e-9df0-4dcb-a733-3d4a8be80123', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
