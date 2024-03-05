-- a.Retrieve the list of all students enrolled in a specific course
SELECT s.name
FROM student_courses sc
JOIN students s ON sc.student_id = s.student_id
WHERE sc.course_id = 1

-- b. Find all the courses a particular student has registered for.    
SELECT c.title
FROM student_courses sc
JOIN courses c ON sc.course_id = c.course_id
WHERE student_id = 1

-- c. Get the preferred schedule of a student for a specific course.
SELECT preferred_schedule
FROM student_courses 
where course_id = 1