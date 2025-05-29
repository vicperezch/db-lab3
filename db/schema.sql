CREATE TYPE LEVEL AS ENUM('BACHELOR', 'MASTER', 'DOCTORATE');
CREATE TYPE GENDER AS ENUM('MALE', 'FEMALE', 'OTHER', 'UNSPECIFIED');
CREATE TYPE MODE AS ENUM('IN_PERSON', 'ONLINE');

CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL,
    level LEVEL NOT NULL,
    gender GENDER NOT NULL
);

CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    description TEXT,
    mode MODE NOT NULL DEFAULT 'IN_PERSON'
);

CREATE TABLE enrollments (
    id SERIAL PRIMARY KEY,
    student_id INTEGER NOT NULL,
    course_id INTEGER NOT NULL,
    enrollment_date DATE NOT NULL CHECK (enrollment_date <= CURRENT_DATE),
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
);
