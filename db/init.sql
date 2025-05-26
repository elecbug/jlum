CREATE TABLE IF NOT EXISTS affiliation (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    link VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS author (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    english_name VARCHAR(100) NOT NULL,
    email VARCHAR(100),
    orcid VARCHAR(19),
    affiliation_id INT NOT NULL,
    FOREIGN KEY (affiliation_id) REFERENCES affiliation(id)
);

CREATE TABLE IF NOT EXISTS author_role (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS journal_db (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS journal (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    link VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS paper (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    journal_id INT NOT NULL,
    volume INT,
    issue INT,
    page_s INT,
    page_e INT,
    year INT,
    month INT,
    doi VARCHAR(100),
    isbn VARCHAR(100),
    issn VARCHAR(100),
    eissn VARCHAR(100),
    abstract TEXT,
    file_data LONGBLOB,
    file_name VARCHAR(255),
    file_type VARCHAR(100),
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (journal_id) REFERENCES journal(id)
);

CREATE TABLE IF NOT EXISTS paper_author (
    id INT AUTO_INCREMENT PRIMARY KEY,
    paper_id INT NOT NULL,
    author_id INT NOT NULL,
    author_role_id INT NOT NULL,
    FOREIGN KEY (paper_id) REFERENCES paper(id),
    FOREIGN KEY (author_id) REFERENCES author(id),
    FOREIGN KEY (author_role_id) REFERENCES author_role(id)
);

CREATE TABLE IF NOT EXISTS journal_journal_db (
    id INT AUTO_INCREMENT PRIMARY KEY,
    journal_id INT NOT NULL,
    journal_db_id INT NOT NULL,
    FOREIGN KEY (journal_id) REFERENCES journal(id),
    FOREIGN KEY (journal_db_id) REFERENCES journal_db(id)
);

CREATE TABLE IF NOT EXISTS paper_reference (
    id INT AUTO_INCREMENT PRIMARY KEY,
    paper_id INT NOT NULL,
    reference_paper_id INT NOT NULL,
    FOREIGN KEY (paper_id) REFERENCES paper(id),
    FOREIGN KEY (reference_paper_id) REFERENCES paper(id)
);

CREATE TABLE IF NOT EXISTS paper_keyword (
    id INT AUTO_INCREMENT PRIMARY KEY,
    paper_id INT NOT NULL,
    keyword VARCHAR(100) NOT NULL,
    FOREIGN KEY (paper_id) REFERENCES paper(id)
);

CREATE TABLE IF NOT EXISTS paper_tag (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS paper_paper_tag (
    id INT AUTO_INCREMENT PRIMARY KEY,
    paper_id INT NOT NULL,
    paper_tag_id INT NOT NULL,
    FOREIGN KEY (paper_id) REFERENCES paper(id),
    FOREIGN KEY (paper_tag_id) REFERENCES paper_tag(id)
);

CREATE TABLE IF NOT EXISTS user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS record (
    id INT AUTO_INCREMENT PRIMARY KEY,
    paper_id INT NOT NULL,
    user_id INT NOT NULL,
    note TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (paper_id) REFERENCES paper(id),
    FOREIGN KEY (user_id) REFERENCES user(id)
);