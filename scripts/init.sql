DROP TABLE IF EXISTS Parking;
DROP TABLE IF EXISTS Block;
DROP TABLE IF EXISTS Slot;
DROP TABLE IF EXISTS Vehicle;

CREATE TABLE Parking (
    id SERIAL PRIMARY KEY,
    name VARCHAR UNIQUE,
    region VARCHAR,
    max INTEGER
);

CREATE TABLE Block (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    parking_id INTEGER REFERENCES Parking(id),
    UNIQUE (name, parking_id)
);

CREATE TABLE Slot (
    id SERIAL PRIMARY KEY,
    number INTEGER,
    parking_id INTEGER REFERENCES Parking(id),
    block_id INTEGER REFERENCES Block(id),
    is_empty BOOLEAN DEFAULT TRUE,
    UNIQUE (number, block_id),
    CHECK (number >= 0 AND number <= 49)
);

CREATE TABLE Vehicle (
    id SERIAL PRIMARY KEY,
    wincode VARCHAR UNIQUE,
    parking_id INTEGER REFERENCES Parking(id) NOT NULL,
    block_id INTEGER REFERENCES Block(id) NOT NULL,
    slot_id INTEGER REFERENCES Slot(id) NOT NULL
);

CREATE OR REPLACE FUNCTION NumToChar(num INT)
RETURNS VARCHAR AS $$
DECLARE
    result VARCHAR;
BEGIN
    SELECT chr(num + 65) INTO result;
    RETURN result;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION AddParking(n text, region text, max INT)
RETURNS INT AS $$
DECLARE
    pid             INT;
    block_count     INT;
    remaining_slots INT;
    bid             INT;
BEGIN
    INSERT INTO parking(name, region, max) VALUES (n, region, max);
    SELECT id INTO pid FROM parking WHERE name = n;

    block_count = max / 50 - 1;
    remaining_slots = max % 50;
    
    FOR i IN 0..block_count LOOP
        INSERT INTO block(name, parking_id) VALUES (NumToChar(i), pid);
        SELECT id INTO bid FROM block WHERE name = NumToChar(i) AND parking_id = pid;

        FOR j IN 0..49 LOOP
            INSERT INTO slot(number, parking_id, block_id) VALUES (j, pid, bid);
        END LOOP;
    END LOOP;

    IF remaining_slots > 0 THEN
        INSERT INTO block(name, parking_id) VALUES (NumToChar(block_count + 1), pid);
        SELECT id INTO bid FROM block WHERE name = NumToChar(block_count + 1) AND parking_id = pid;

        FOR j IN 0..remaining_slots - 1 LOOP
            INSERT INTO slot(number, parking_id, block_id) VALUES (j, pid, bid);
        END LOOP;
    END IF;

    RETURN pid;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION FindFreeSlot(parking_name VARCHAR)
RETURNS TABLE (block_name VARCHAR, slot_number INT) AS $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM parking WHERE name = parking_name) THEN
        RAISE EXCEPTION 'Parking with name % does not exist', parking_name;
    END IF;

    RETURN QUERY
    SELECT b.name, s.number
    FROM parking p
    JOIN block b ON p.id = b.parking_id
    JOIN slot s ON b.id = s.block_id
    WHERE p.name = parking_name AND s.is_empty = TRUE
    ORDER BY b.name, s.number
    LIMIT 1;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION AddVehicle(wincode VARCHAR, parking_name VARCHAR)
RETURNS TABLE (block_name VARCHAR, slot_number INT) AS $$
DECLARE
    pid INT;
    bid INT;
    sid INT;
BEGIN
    IF NOT EXISTS (SELECT 1 FROM parking WHERE name = parking_name) THEN
        RAISE EXCEPTION 'Parking with name % does not exist', parking_name;
    END IF;

    SELECT id INTO pid FROM parking WHERE name = parking_name;

    SELECT b.id INTO bid
    FROM block b
    JOIN slot s ON b.id = s.block_id
    WHERE b.parking_id = pid AND s.is_empty = TRUE
    ORDER BY b.name ASC, s.number ASC
    LIMIT 1;
    
    IF bid IS NULL THEN
        RAISE EXCEPTION 'No free slot available in parking %', parking_name;
    END IF;
    
    SELECT s.id INTO sid
    FROM slot s
    WHERE s.block_id = bid AND s.is_empty = TRUE
    ORDER BY s.number ASC
    LIMIT 1;
    
    IF NOT FOUND THEN
        RAISE EXCEPTION 'No free slot available in parking %', parking_name;
    END IF;

    INSERT INTO vehicle (wincode, parking_id, block_id, slot_id) VALUES (wincode, pid, bid, sid);
    UPDATE slot SET is_empty = FALSE WHERE id = sid;
    
    RETURN QUERY
    SELECT b.name, s.number
    FROM block b
    JOIN slot s ON b.id = s.block_id
    WHERE s.id = sid;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION FindVehicle(wc VARCHAR)
RETURNS TABLE (parking_name VARCHAR, block_name VARCHAR, slot_number INT) AS $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM vehicle WHERE wincode = wc) THEN
        RAISE EXCEPTION 'Vehicle with wincode % does not exist', wc;
    END IF;

    RETURN QUERY
    SELECT p.name, b.name, s.number
    FROM vehicle v
    JOIN parking p ON v.parking_id = p.id
    JOIN block b ON v.block_id = b.id
    JOIN slot s ON v.slot_id = s.id
    WHERE v.wincode = wc;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION RemoveVehicle(wc VARCHAR)
RETURNS TABLE (parking_name VARCHAR, block_name VARCHAR, slot_number INT) AS $$
DECLARE
    pid INT;
    bid INT;
    sid INT;
BEGIN
    SELECT v.parking_id, v.block_id, v.slot_id INTO pid, bid, sid
    FROM vehicle v
    WHERE v.wincode = wc;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'Vehicle with wincode % does not exist', wc;
    END IF;

    DELETE FROM vehicle WHERE wincode = wc;
    UPDATE slot SET is_empty = TRUE WHERE id = sid;

    RETURN QUERY
    SELECT p.name, b.name, s.number
    FROM parking p
    JOIN block b ON p.id = b.parking_id
    JOIN slot s ON b.id = s.block_id
    WHERE s.id = sid;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION FindAllFreeSlots(parking_name VARCHAR)
RETURNS TABLE (block_name VARCHAR, slot_number INT) AS $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM parking WHERE name = parking_name) THEN
        RAISE EXCEPTION 'Parking with name % does not exist', parking_name;
    END IF;

    RETURN QUERY
    SELECT b.name, s.number
    FROM parking p
    JOIN block b ON p.id = b.parking_id
    JOIN slot s ON b.id = s.block_id
    WHERE p.name = parking_name AND s.is_empty = TRUE
    ORDER BY b.name, s.number;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION FindAllFullSlots(parking_name VARCHAR)
RETURNS TABLE (block_name VARCHAR, slot_number INT, wincode VARCHAR) AS $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM parking WHERE name = parking_name) THEN
        RAISE EXCEPTION 'Parking with name % does not exist', parking_name;
    END IF;

    RETURN QUERY
    SELECT b.name, s.number, v.wincode
    FROM parking p
    JOIN block b ON p.id = b.parking_id
    JOIN slot s ON b.id = s.block_id
    JOIN vehicle v ON s.id = v.slot_id
    WHERE p.name = parking_name AND NOT s.is_empty
    ORDER BY b.name, s.number, v.wincode;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION CheckAllSlots(parking_name VARCHAR)
RETURNS TABLE (block_name VARCHAR, slot_number VARCHAR, is_empty BOOLEAN) AS $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM parking WHERE name = parking_name) THEN
        RAISE EXCEPTION 'Parking with name % does not exist', parking_name;
    END IF;

    RETURN QUERY
    SELECT b.name, s.number, s.is_empty
    FROM parking p
    JOIN block b ON p.id = b.parking_id
    JOIN slot s ON b.id = s.block_id
    WHERE p.name = parking_name
    ORDER BY b.name, s.number;
END;
$$ LANGUAGE plpgsql;

