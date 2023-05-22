SELECT 'CREATE DATABASE auth' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'auth')\gexec
\c auth;

CREATE TABLE users (
    email TEXT PRIMARY KEY,
    login TEXT NOT NULL,
    password TEXT NOT NULL,
    role VARCHAR(10) DEFAULT 'MEMBER'
);

INSERT INTO users (login, password, email, role)
VALUES (
    'Olenka',
    'Z55+a4XKM9tBClCi4/z9soOEK1/th6bWGveVqhgZTth3uoXAt+afxpy9m77Mo+y7LHJVKipxbOQL1u90V9oceaHaQATc9DH5UB8SEtYg/I6NKyrrnjQasdy7NBN++6834ZErQEsA6+9DmIr4ER3H2ecnQbXiRjBHQ5M2hzvTqc8=',
    'Olenka@mail.ru',
    'ADMIN'
);

INSERT INTO users (login, password, email, role)
VALUES (
    'Katya',
    'b7t2PnSfjgF7/7Rr+gvOd5whra5HP7q9bV6AXp5sdRfQN0R4ashgfSr6hXi8KxkWQVf3ebmOAngocSc6Wo9HOX/I6OxIACEptozQ4eOwC0PR15ZO3w5SlOWMe6+wyjaJwOdOjhcPHQ1cP5DxxkWlIY+p/7XjcqHUNMzYdYQss8I=',
    'katya@mail.ru',
    'MEMBER'
);
