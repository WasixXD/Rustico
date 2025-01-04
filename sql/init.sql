DROP TABLE IF EXISTS PROJECTS;


CREATE TABLE projects (
    id INTEGER PRIMARY KEY,
    project_name TEXT,
    description TEXT,
    created_at DATETIME DEFAULT (DATETIME('now', 'localtime')),
    github_url TEXT
);