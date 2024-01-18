CREATE DATABASE student_project_management;

USE student_project_management;

CREATE TABLE
    calendar_account (
        id INT NOT NULL auto_increment,
        login VARCHAR(50) NOT NULL,
        email VARCHAR(100) NOT NULL,
        PRIMARY KEY(id)
    );

CREATE TABLE
    repository_account (
        id INT NOT NULL auto_increment,
        login VARCHAR(50) NOT NULL,
        email VARCHAR(100) NOT NULL,
        PRIMARY KEY(id)
    );

CREATE TABLE
    professor (
        id INT NOT NULL auto_increment,
        name VARCHAR(50) NOT NULL,
        surname VARCHAR(50) NOT NULL,
        middle_name VARCHAR(50) NOT NULL,
        position VARCHAR(100) NOT NULL,
        calendar_account_id INT,
        calendar_id VARCHAR(100),
        repo_account_id INT,
        PRIMARY KEY(id),
        FOREIGN KEY (calendar_account_id) REFERENCES calendar_account(id) ON DELETE CASCADE,
        FOREIGN KEY (repo_account_id) REFERENCES repository_account(id) ON DELETE CASCADE
    );

CREATE TABLE
    student (
        id INT NOT NULL auto_increment,
        name VARCHAR(50) NOT NULL,
        surname VARCHAR(50) NOT NULL,
        middle_name VARCHAR(50) NOT NULL,
        uni_group VARCHAR(20) NOT NULL,
        PRIMARY KEY(id)
    );

CREATE TABLE
    project_status (
        id INT NOT NULL auto_increment,
        name VARCHAR(50) NOT NULL,
        PRIMARY KEY(id)
    );

CREATE TABLE
    repository (
        id INT NOT NULL auto_increment,
        name VARCHAR(100) NOT NULL,
        owner VARCHAR(100) NOT NULL,
        PRIMARY KEY(id)
    );

CREATE TABLE
    project (
        id INT NOT NULL auto_increment,
        theme VARCHAR(100) NOT NULL,
        year INT NOT NULL,
        supervisor_id INT NOT NULL,
        student_id INT NOT NULL,
        status_id INT NOT NULL,
        repoId INT,
        cloud_drive_folder_id VARCHAR(100) NOT NULL,
        PRIMARY KEY(id),
        FOREIGN KEY (supervisor_id) REFERENCES professor(id) ON DELETE CASCADE,
        FOREIGN KEY (student_id) REFERENCES student(id) ON DELETE CASCADE,
        FOREIGN KEY (status_id) REFERENCES project_status(id) ON DELETE CASCADE,
        FOREIGN KEY (repoId) REFERENCES repository(id) ON DELETE CASCADE
    );

CREATE TABLE
    project_stage (
        id INT NOT NULL auto_increment,
        name VARCHAR(50) NOT NULL,
        PRIMARY KEY(id)
    );

CREATE TABLE
    task_status (
        id INT NOT NULL auto_increment,
        name VARCHAR(50) NOT NULL,
        PRIMARY KEY(id)
    );
CREATE TABLE
    meeting_status (
        id INT NOT NULL auto_increment,
        name VARCHAR(50) NOT NULL,
        PRIMARY KEY(id)
    );

CREATE TABLE
    task (
        id INT NOT NULL auto_increment,
        name VARCHAR(50) NOT NULL,
        description VARCHAR(300) NOT NULL,
        deadline DATETIME NOT NULL,
        status_id INT NOT NULL,
        project_id INT NOT NULL,
        project_stage_id INT NOT NULL,
        cloud_drive_folder_id VARCHAR(100) NOT NULL,
        PRIMARY KEY(id),
        FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE,
        FOREIGN KEY (project_stage_id) REFERENCES project_stage(id) ON DELETE CASCADE
    );

CREATE TABLE
    Meeting (
        id INT NOT NULL auto_increment,
        organizer_id INT NOT NULL,
        participant_id INT NOT NULL,
        time DATETIME NOT NULL,
        is_online BOOLEAN NOT NULL,
        status_id INT NOT NULL,
        description VARCHAR(200),
        PRIMARY KEY(id),
        FOREIGN KEY (organizer_id) REFERENCES professor(id) ON DELETE CASCADE,
        FOREIGN KEY (participant_id) REFERENCES student(id) ON DELETE CASCADE,
        FOREIGN KEY (status_id) REFERENCES meeting_status(id) ON DELETE CASCADE
    );