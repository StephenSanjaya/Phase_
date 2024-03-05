INSERT INTO Student (student_name)
VALUES
    ("Budi"),
    ("Siti"),
    ("Rudi");

INSERT INTO Professor (professor_name)
VALUES
    ("Mr. X"),
    ("Mr. Y"),
    ("Mr. V"),
    ("Mr. Z"),
    ("Mrs. U");

INSERT INTO Course (course_title, max_capacity)
VALUES
    ("CourseA", 200),
    ("CourseB", 100),
    ("CourseC", 300),
    ("CourseD", 200),
    ("CourseE", 150);

INSERT INTO Enrollment (student_id, course_id, enrollment_date)
VALUES
    (1, 2, "2022-02-22"),
    (2, 2, "2022-02-22");

INSERT INTO TeachingAssignment (professor_id, course_id, start_date)
VALUES
    (1, 2, "2022-02-22"),
    (2, 1, "2022-02-22");


