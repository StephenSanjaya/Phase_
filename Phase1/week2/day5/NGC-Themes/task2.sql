CREATE DATABASE company;
use company;

CREATE TABLE IF NOT EXISTS Departments {
    departmentID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
};

CREATE TABLE IF NOT EXISTS Employees {
    employeeID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    departmentID INT,
    FOREIGN KEY (departmentID) REFERENCES Departments(departmentID)
};

CREATE TABLE IF NOT EXISTS Projects {
    projectID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
};

CREATE TABLE IF NOT EXISTS Employees_Projects {
    employee_projectID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    employeeID INT,
    projectID INT,
    FOREIGN KEY (employeeID) REFERENCES Employees(employeeID),
    FOREIGN KEY (projectID) REFERENCES Projects(projectID)
};