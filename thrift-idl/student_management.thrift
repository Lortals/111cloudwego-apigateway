namespace go student.api

struct RegisterStudentRequest {
    1: required string name;
    2: required string age;
    3: required string email;
    4: required string address;
}

struct RegisterStudentResponse {
    1: string status;
    2: string id;
}

struct GetStudentRequest {
    1: required string id;
}

struct GetStudentResponse {
    1: string name;
    2: string age;
    3: string email;
    4: string address;
}

service StudentManagement {
    RegisterStudentResponse registerStudent(1: RegisterStudentRequest req);
    GetStudentResponse getStudent(1: GetStudentRequest req);
}
