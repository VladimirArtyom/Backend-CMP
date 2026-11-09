-- TABLE des participants pour le vote
DROP TABLE IF EXISTS participant;
DROP TABLE IF EXISTS t_group;
DROP TABLE IF EXISTS fav_group;
DROP TABLE IF EXISTS referral;
DROP TABLE IF EXISTS grades;

CREATE TABLE IF NOT EXISTS participant (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS t_group(
    id SERIAL PRIMARY KEY,
    class VARCHAR(100) NOT NULL,
    group_name VARCHAR(255) NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS votes(
    id SERIAL PRIMARY KEY, 
    id_participant INT NOT NULL, 
    id_group INT NOT NULL,
    class VARCHAR(100) NOT NULL,
    group_name VARCHAR(255) NOT NULL,
    score INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    CONSTRAINT fk_participant FOREIGN KEY (id_participant) REFERENCES participant(id) ON DELETE CASCADE,
    CONSTRAINT fk_group FOREIGN KEY (id_group) REFERENCES t_group(id) ON DELETE CASCADE,
    CONSTRAINT unique_vote UNIQUE (id_participant, id_group)
);
-- Table de jurys
CREATE TABLE IF NOT EXISTS referall (
    id SERIAL PRIMARY KEY,
    referral_code VARCHAR(50) NOT NULL UNIQUE,
    grader_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table des notes
CREATE TABLE IF NOT EXISTS grades(
    id SERIAL PRIMARY KEY, 
    referral_code VARCHAR(50) NOT NULL,
    group_id INT NOT NULL,
    group_name VARCHAR(255) NOT NULL,
    grade int NOT NULL,
    year int NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    --CONSTRAINT fk_group_grade FOREIGN KEY (group_id) REFERENCES t_group(id) ON DELETE CASCADE,
    CONSTRAINT unique_grade UNIQUE(referral_code, group_id)

);

