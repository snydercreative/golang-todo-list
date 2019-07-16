DO $$
DECLARE 
    lastvalue BIGINT;

BEGIN
    TRUNCATE todolist CASCADE;

    INSERT INTO todoList (name, created) VALUES ('Test list 1', NOW()) RETURNING id INTO lastvalue;

    INSERT INTO todo (listid, name, created) VALUES (lastvalue, 'Test item A1', NOW());   
    INSERT INTO todo (listid, name, created) VALUES (lastvalue, 'Test item A2', NOW());
    INSERT INTO todo (listid, name, created) VALUES (lastvalue, 'Test item A3', NOW());
    INSERT INTO todo (listid, name, created) VALUES (lastvalue, 'Test item A4', NOW());

    INSERT INTO todoList (name, created) VALUES ('Test list 2', NOW()) RETURNING id INTO lastvalue;

    INSERT INTO todo (listid, name, created) VALUES (lastvalue, 'Test item B1', NOW());
    INSERT INTO todo (listid, name, created) VALUES (lastvalue, 'Test item B2', NOW());

    INSERT INTO todoList (name, created) VALUES ('Test list 3', NOW()) RETURNING id INTO lastvalue;

    INSERT INTO todo (listid, name, created) VALUES (lastvalue, 'Test item C1', NOW());
    INSERT INTO todo (listid, name, created) VALUES (lastvalue, 'Test item C2', NOW());
    INSERT INTO todo (listid, name, created) VALUES (lastvalue, 'Test item C3', NOW());
END $$;