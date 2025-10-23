CREATE TABLE "channels"(
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "guild_id" INT NOT NULL,
    FOREIGN KEY (guild_id) REFERENCES guilds(id) ON DELETE CASCADE
)
