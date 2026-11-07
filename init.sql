-- TABLE des participants pour le vote

CREATE TABLE IF NOT EXISTS participant (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS kelompok (
    id SERIAL PRIMARY KEY,
    kelas VARCHAR(100) NOT NULL,
    nama_kelompok VARCHAR(255) NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS kelompok_favorit (
    id SERIAL PRIMARY KEY, 
    id_participant INT NOT NULL, 
    id_kelompok INT NOT NULL,
    kelas VARCHAR(100) NOT NULL,
    nama_kelompok VARCHAR(255) NOT NULL,
    skor INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    CONSTRAINT fk_participant FOREIGN KEY (id_participant) REFERENCES participant(id) ON DELETE CASCADE,
    CONSTRAINT fk_kelompok FOREIGN KEY (id_kelompok) REFERENCES kelompok(id) ON DELETE CASCADE,
    CONSTRAINT unique_vote UNIQUE (id_participant, id_kelompok)
);
-- Table de jurys
CREATE TABLE IF NOT EXISTS referall (
    id SERIAL PRIMARY KEY,
    kode_referral VARCHAR(50) NOT NULL UNIQUE,
    nama_penilai VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table des notes
CREATE TABLE IF NOT EXISTS penilaian(
    id SERIAL PRIMARY KEY, 
    kode_referral VARCHAR(50) NOT NULL,
    id_kelompok INT NOT NULL,
    nama_penilai VARCHAR(255) NOT NULL,
    nilai int NOT NULL,
    tahun int NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    CONSTRAINT fk_kelompok_penilaian FOREIGN KEY (id_kelompok) REFERENCES kelompok(id) ON DELETE CASCADE,
    CONSTRAINT unique_penilaian UNIQUE(kode_referral, id_kelompok)

);

