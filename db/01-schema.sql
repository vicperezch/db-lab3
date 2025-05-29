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
    FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
);

CREATE OR REPLACE VIEW enrollments_info AS
SELECT 
	e.id AS id,
	c.name AS course_name,
	s.first_name AS first_name,
	s.last_name AS last_name,
	s.level AS student_level,
	c.mode AS course_mode,
	e.enrollment_date AS enrollment_date
FROM enrollments AS e
LEFT JOIN students AS s ON s.id = e.student_id
LEFT JOIN courses AS c ON c.id = e.course_id;
