-- PRAGMA foreign_keys = ON; -- Para habilitar as constraints fk

CREATE TABLE collection (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    icon TEXT,
    collection_id INTEGER,
    create_at INTEGER,
    update_at INTEGER,

    CONSTRAINT collection_parent_fk
    FOREIGN KEY (collection_id)
    REFERENCES collection (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE collection_seq (
    id integer
);

insert into collection_seq values (1);

insert into collection (_id, name, icon, collection_id, create_at, update_at) 
values('265d38b8-5764-4336-882f-901e77634727', 'Work', 'https://cdn-icons-png.freepik.com/512/7398/7398232.png', null, 1715508444, 1715508444);

insert into collection (_id, name, icon, collection_id, create_at, update_at) 
values('84593371-9456-4aac-85eb-acb579a0899b', 'Study', null, null, 1715508444, 1715508444);

insert into collection (_id, name, icon, collection_id, create_at, update_at) 
values('835abdbe-1395-4197-b05a-cbedd6814aa2', 'Warning', null, 15, 1715508444, 1715508444);

CREATE TABLE read_later (
    id INTEGER PRIMARY KEY,
    _id text NOT NULL,
    url TEXT not null,
    title TEXT NOT NULL,
    collection_id INTEGER,
    cover TEXT,
    create_at INTEGER,
    update_at INTEGER,

    CONSTRAINT read_later_id_uk UNIQUE (_id),

    CONSTRAINT read_later_collection_fk
    FOREIGN KEY (collection_id)
    REFERENCES collection (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE  
);