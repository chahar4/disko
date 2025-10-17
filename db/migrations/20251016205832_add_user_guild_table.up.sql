CREATE TABLE user_guild (
    user_id INT NOT NULL,
    guild_id INT NOT NULL,
    PRIMARY KEY (user_id, guild_id),
    FOREIGN KEY (user_id) REFERENCES "users"(id) ON DELETE CASCADE,
    FOREIGN KEY (guild_id) REFERENCES guilds(id) ON DELETE CASCADE
)

