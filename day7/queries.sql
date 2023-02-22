-- create tictactoe table named ttt
CREATE DATABASE ttt;

-- create students table
CREATE TABLE IF NOT EXISTS ttt.students (
	id INTEGER AUTO_INCREMENT,
	name VARCHAR(50),
	student_id VARCHAR(10),
	gpa FLOAT,
	PRIMARY KEY (id)
);

-- retrive all students from database with all the attributes
SELECT * FROM ttt.students;

-- retrive a single student with id with all the attributes
SELECT * FROM ttt.students WHERE id = 1;

-- retrive a single student with id with only name and gpa attributes
SELECT (name, gpa) gpa FROM ttt.students WHERE id = 1;

-- update the student with id=1 and set the name to Mohammad
UPDATE ttt.students SET name='Mohammad' WHERE id=1;

-- delete the student with id=2
DELETE FROM ttt.students WHERE id=2;

-- delete all of the students
DELETE FROM ttt.students;