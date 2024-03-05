INSERT INTO students (name, email, major, year_of_study)
VALUES 
    ("Budi", "Budi@gmail.com", "MajorA", "2022"),
    ("Yudi", "Yudi@gmail.com", "MajorC", "2022"),
    ("Rudi", "Rudi@gmail.com", "MajorA", "2022"),
    ("Siti", "Siti@gmail.com", "MajorA", "2022"),
    ("Fandi", "Fandi@gmail.com", "MajorB", "2022");

INSERT INTO courses (title, instructor, schedule, credit_hours)
VALUES 
    ("TitleA", "Mr. X", "Monday", 2),
    ("TitleB", "Mr. S", "Tuesday", 1),
    ("TitleC", "Mr. V", "Wednesday", 2),
    ("TitleD", "Mr. Z", "Thursday", 2),
    ("TitleE", "Mr. Y", "Friday", 4);

INSERT INTO student_courses (student_id, course_id, preferred_schedule)
VALUES 
    (1, 2, "Monday"),
    (2, 1, "Tuesday"),
    (1, 3, "Wednesday"),
    (4, 4, "Thursday"),
    (5, 1, "Friday");


