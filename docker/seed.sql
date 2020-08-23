CREATE TABLE budgets(
   id serial PRIMARY KEY,
   amount INTEGER,
   created_at timestamp default NULL,
   updated_at timestamp default NULL
);

-- INSERT INTO budgets (amount) VALUES (100);
-- INSERT INTO budgets (amount) VALUES (200);
-- INSERT INTO budgets (amount) VALUES (300);