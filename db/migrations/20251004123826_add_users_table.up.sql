CREATE TABLE "users"(
    "id" BIGSERIAL PRIMARY KEY,
    "username" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL
)