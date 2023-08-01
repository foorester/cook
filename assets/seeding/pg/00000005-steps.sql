--SEED
INSERT INTO steps (id, name, description, recipe_id, duration, created_at, updated_at)
VALUES ('c2ad0f71-7cf0-4464-b15c-2e18d295d953', 'Boil the Pasta', 'Bring a large pot of salted water to a boil. Add the spaghetti and cook until al dente.', '99e8cf36-af30-453e-8605-1c33c950d121', '12 minutes', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO steps (id, name, description, recipe_id, duration, created_at, updated_at)
VALUES ('a05a3a13-01c6-487b-8531-432a9c7c27bb', 'Make the Bolognese Sauce', 'In a large skillet, brown the ground beef. Add chopped tomatoes, tomato paste, onions, garlic, and herbs. Simmer until the sauce thickens.', '99e8cf36-af30-453e-8605-1c33c950d121', '30 minutes', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO steps (id, name, description, recipe_id, duration, created_at, updated_at)
VALUES ('a1c6c1a3-6c2b-40c8-a739-2e1d6d7b6fbb', 'Combine Pasta and Sauce', 'Drain the pasta and add it to the skillet with the Bolognese sauce. Toss until the pasta is coated with the sauce.', '99e8cf36-af30-453e-8605-1c33c950d121', '2 minutes', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

--SEED
INSERT INTO steps (id, name, description, recipe_id, duration, created_at, updated_at)
VALUES ('9f6e80c9-8119-4610-af38-78216a516c27', 'Prepare the Chicken', 'Cut the chicken thighs into bite-sized pieces and marinate them with yogurt, ginger, garlic, and spices.', 'f1ee3b1f-8df3-4e1f-9e89-4f0c7cbb8a23', '15 minutes', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO steps (id, name, description, recipe_id, duration, created_at, updated_at)
VALUES ('6bc4570e-197a-4d9c-9db9-bc6da5dd9b75', 'Sauté the Onion', 'Heat oil in a pan and sauté the chopped onion until golden brown.', 'f1ee3b1f-8df3-4e1f-9e89-4f0c7cbb8a23', '5 minutes', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO steps (id, name, description, recipe_id, duration, created_at, updated_at)
VALUES ('5405c0a1-b84e-4f62-86e9-978e582c015f', 'Cook the Chicken', 'Add marinated chicken to the pan and cook until it turns tender and cooked through.', 'f1ee3b1f-8df3-4e1f-9e89-4f0c7cbb8a23', '20 minutes', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

--SEED
INSERT INTO steps (id, name, description, recipe_id, duration, created_at, updated_at)
VALUES ('777af69b-8c33-441c-bc1b-7c6011f5f0a6', 'Marinate the Pork', 'Mix the spices, pineapple juice, and vinegar. Marinate the sliced pork in the mixture for at least 2 hours or overnight.', '021d9561-5a3a-4ebd-9850-7e1720f6e02b', '2 hours', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO steps (id, name, description, recipe_id, duration, created_at, updated_at)
VALUES ('94b6f2db-88af-467f-ba7b-7b6deac80354', 'Prepare the Taco Filling', 'Sauté the marinated pork with diced pineapple until fully cooked and caramelized.', '021d9561-5a3a-4ebd-9850-7e1720f6e02b', '15 minutes', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO steps (id, name, description, recipe_id, duration, created_at, updated_at)
VALUES ('9de28181-2c40-465f-8de0-22e27ebe28c7', 'Assemble the Tacos', 'Warm the tortillas, and fill them with the cooked pork, diced onions, cilantro, and salsa.', '021d9561-5a3a-4ebd-9850-7e1720f6e02b', '10 minutes', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
