-- a. List of students enrolled in a specific course
SELECT *
FROM Student s
JOIN Enrollment e ON s.student_id = e.student_id
WHERE e.course_id = 1

-- b. List of courses taught by a specific professor   
SELECT *
FROM Course c
JOIN TeachingAssignment ta ON c.course_id = ta.course_id
WHERE ta.professor_id = 1

-- c. List of professor teaching a specific course
SELECT *
FROM Professor p
JOIN TeachingAssignment ta ON p.professor_id = ta.professor_id
WHERE ta.course_id = 1