CREATE TABLE IF NOT EXISTS Users (
    User_ID BIGSERIAL NOT NULL,
    PRIMARY KEY (User_ID)
);

CREATE TABLE IF NOT EXISTS Local_Users (
    User_ID BIGSERIAL NOT NULL,
    Email TEXT UNIQUE NOT NULL,
    Password BYTEA NOT NULL,
    FOREIGN KEY (User_ID) REFERENCES Users (User_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Foreign_Users (
    User_ID BIGSERIAL NOT NULL,
    Home_Server TEXT NOT NULL,
    Local_User_ID BIGSERIAL UNIQUE NOT NULL,
    FOREIGN KEY (Local_User_ID) REFERENCES Users (User_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Federation_Nonces (
    Nonce TEXT NOT NULL,
    User_ID BIGSERIAL NOT NULL,
    Home_Server TEXT NOT NULL,
    FOREIGN KEY (User_ID) REFERENCES Users (User_ID) ON DELETE CASCADE,
    PRIMARY KEY (Nonce)
);

CREATE TABLE IF NOT EXISTS Profiles (
    User_ID BIGSERIAL NOT NULL,
    Username TEXT UNIQUE NOT NULL,
    Avatar TEXT,
    Status SMALLINT NOT NULL,
    FOREIGN KEY (User_ID) REFERENCES Users (User_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Guild_List (
    User_ID BIGSERIAL NOT NULL,
    Guild_ID BIGSERIAL NOT NULL,
    Home_Server TEXT NOT NULL,
    Position TEXT NOT NULL,
    FOREIGN KEY (User_ID) REFERENCES Users (User_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS User_Metadata (
    User_ID BIGSERIAL NOT NULL,
    App_ID TEXT NOT NULL,
    Metadata TEXT NOT NULL,
    FOREIGN KEY (User_ID) REFERENCES Users (User_ID) ON DELETE CASCADE,
    PRIMARY KEY (App_ID)
);

CREATE TABLE IF NOT EXISTS Sessions (
    Session TEXT PRIMARY KEY NOT NULL,
    User_ID BIGSERIAL NOT NULL,
    Expiration BIGINT NOT NULL,
    FOREIGN KEY (User_ID) REFERENCES Users (User_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Guilds (
    Guild_ID BIGSERIAL PRIMARY KEY NOT NULL,
    Owner_ID BIGSERIAL NOT NULL,
    Guild_Name TEXT NOT NULL,
    Picture_URL TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Roles (
    Guild_ID BIGSERIAL NOT NULL,
    Role_ID BIGSERIAL NOT NULL,
    Name TEXT NOT NULL,
    Color INTEGER NOT NULL,
    Hoist BOOLEAN NOT NULL,
    Pingable BOOLEAN NOT NULL,
    Position TEXT NOT NULL,
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE,
    PRIMARY KEY (Role_ID)
);

CREATE TABLE IF NOT EXISTS Roles_Members (
    Guild_ID BIGSERIAL NOT NULL,
    Role_ID BIGSERIAL NOT NULL,
    Member_ID BIGSERIAL NOT NULL,
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE,
    FOREIGN KEY (Role_ID) REFERENCES Roles (Role_ID) ON DELETE CASCADE,
    FOREIGN KEY (Member_ID) REFERENCES Users (User_ID) ON DELETE CASCADE,
    PRIMARY KEY (Guild_ID, Role_ID, Member_ID)
);

--migration-only DO $$ BEGIN
CREATE TYPE PermissionsNode AS (Node TEXT, Allow BOOLEAN);

--migration-only EXCEPTION
--migration-only WHEN duplicate_object THEN null;
--migration-only END $$;
CREATE TABLE IF NOT EXISTS Permissions (
    Guild_ID BIGSERIAL NOT NULL,
    Channel_ID BIGINT,
    Role_ID BIGSERIAL NOT NULL,
    Nodes PermissionsNode [] NOT NULL,
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE,
    FOREIGN KEY (Role_ID) REFERENCES Roles (Role_ID) ON DELETE CASCADE,
    FOREIGN KEY (Channel_ID) REFERENCES Channels (Channel_ID) ON DELETE CASCADE,
    UNIQUE (Guild_ID, Channel_ID, Role_ID)
);

CREATE TABLE IF NOT EXISTS Guild_Members (
    User_ID BIGSERIAL NOT NULL,
    Guild_ID BIGSERIAL NOT NULL,
    UNIQUE (User_ID, Guild_ID),
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Invites (
    Invite_ID TEXT PRIMARY KEY UNIQUE,
    Uses INTEGER NOT NULL DEFAULT 0,
    Possible_Uses INTEGER DEFAULT -1,
    Guild_ID BIGSERIAL NOT NULL,
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Channels (
    Channel_ID BIGSERIAL PRIMARY KEY UNIQUE,
    Guild_ID BIGSERIAL,
    Channel_Name TEXT NOT NULL,
    Position TEXT NOT NULL,
    Category BOOLEAN NOT NULL,
    IsVoice BOOLEAN NOT NULL,
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Messages (
    Message_ID BIGSERIAL PRIMARY KEY,
    Guild_ID BIGSERIAL NOT NULL,
    Channel_ID BIGSERIAL NOT NULL,
    User_ID BIGSERIAL NOT NULL,
    Created_At TIMESTAMP NOT NULL,
    Edited_At TIMESTAMP,
    Content TEXT NOT NULL,
    Embeds jsonb,
    Actions jsonb,
    Overrides bytea,
    Reply_To_ID BIGINT DEFAULT 0,
    Attachments text [],
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE,
    FOREIGN KEY (Channel_ID) REFERENCES Channels (Channel_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Files (
    Hash BYTEA PRIMARY KEY NOT NULL,
    File_ID TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Rate_Limit_Whitelist_IP (IP TEXT NOT NULL PRIMARY KEY);

CREATE TABLE IF NOT EXISTS Rate_Limit_Whitelist_User (
    User_ID BIGSERIAL NOT NULL,
    FOREIGN KEY (User_ID) REFERENCES Users (User_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Emote_Packs (
    Pack_ID BIGSERIAL NOT NULL,
    Pack_Name TEXT NOT NULL,
    User_ID BIGSERIAL NOT NULL,
    PRIMARY KEY (Pack_ID),
    FOREIGN KEY (User_ID) REFERENCES Users (User_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Emote_Pack_Emotes (
    Pack_ID BIGSERIAL NOT NULL,
    Image_ID TEXT NOT NULL,
    Emote_Name TEXT NOT NULL,
    FOREIGN KEY (Pack_ID) REFERENCES Emote_Packs (Pack_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Acquired_Emote_Packs (
    Pack_ID BIGSERIAL NOT NULL,
    User_ID BIGSERIAL NOT NULL,
    FOREIGN KEY (Pack_ID) REFERENCES Emote_Packs (Pack_ID) ON DELETE CASCADE,
    FOREIGN KEY (User_ID) REFERENCES Users (User_ID) ON DELETE CASCADE
);