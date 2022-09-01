CREATE TABLE IF NOT EXISTS substances (
    id serial primary key
    , name text NOT NULL
);

CREATE TABLE IF NOT EXISTS drugs (
    id serial primary key
    , name text
    , substance_id int REFERENCES substances (id)
);

INSERT INTO substances (id , name)
    VALUES (1 , 'ramipril')
    , (2 , 'tamoxifen')
ON CONFLICT
    DO NOTHING;

INSERT INTO drugs (id , name , substance_id)
    VALUES (1 , 'drug A' , 1)
    , (2 , 'drug B' , 1)
    , (3 , 'drug C' , 2)
    , (4 , 'drug D' , 2)
ON CONFLICT
    DO NOTHING;

