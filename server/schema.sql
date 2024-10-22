CREATE TABLE groups (
        id INTEGER NOT NULL PRIMARY KEY,
        name TEXT NOT NULL 
);

CREATE TABLE members (
        id INTEGER NOT NULL PRIMARY KEY,
        username TEXT NOT NULL,
        displayName TEXT NOT NULL,
        password TEXT NOT NULL 
);

CREATE TABLE member_groups (
        group_id INTEGER NOT NULL ,
        member_id INTEGER NOT NULL ,
        PRIMARY KEY (group_id, member_id),
        FOREIGN KEY (group_id) REFERENCES groups (id),
        FOREIGN KEY (member_id) REFERENCES members (id)
);
CREATE TABLE items (
        id INTEGER PRIMARY KEY NOT NULL,
        timestamp INTEGER NOT NULL,
        name TEXT NOT NULL ,
        price REAL NOT NULL ,
        author_id INTEGER NOT NULL ,
        group_id INTEGER NOT NULL ,
        FOREIGN KEY (group_id) REFERENCES groups (id),
        FOREIGN KEY (author_id) REFERENCES members (id)
);
