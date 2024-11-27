CREATE TABLE groups ( --bs: low
        id INTEGER NOT NULL PRIMARY KEY,
        name TEXT NOT NULL --bs: company
);

CREATE TABLE members ( --bs: low
        id INTEGER NOT NULL PRIMARY KEY,
        username TEXT NOT NULL, --bs: username_random
        displayName TEXT NOT NULL, --bs: fullname
        password TEXT NOT NULL --bs: val; 1234
);

CREATE TABLE member_groups ( --bs: medium
        group_id INTEGER NOT NULL, --bs: rel
        member_id INTEGER NOT NULL, --bs: rel
        PRIMARY KEY (group_id, member_id),
        FOREIGN KEY (group_id) REFERENCES groups (id),
        FOREIGN KEY (member_id) REFERENCES members (id)
);
CREATE TABLE items ( --bs: medium
        id INTEGER PRIMARY KEY NOT NULL,
        timestamp INTEGER NOT NULL, --bs: timestamp_epoch
        name TEXT NOT NULL, --bs: product
        price REAL NOT NULL, --bs: num; 1to200
        author_id INTEGER NOT NULL, --bs: rel
        group_id INTEGER NOT NULL, --bs: rel
        reimbursement BOOLEAN DEFAULT 0, --bs: bool; 0.05
        FOREIGN KEY (group_id) REFERENCES groups (id),
        FOREIGN KEY (author_id) REFERENCES members (id)
);

CREATE TABLE expense_participants (
    item_id INTEGER NOT NULL, -- Foreign key referencing the items table
    member_id INTEGER NOT NULL, -- Foreign key referencing the members table
    PRIMARY KEY (item_id, member_id), -- Composite primary key (ensures unique pairs)
    FOREIGN KEY (item_id) REFERENCES items (id), -- Foreign key constraint to items table
    FOREIGN KEY (member_id) REFERENCES members (id) -- Foreign key constraint to members table
);
