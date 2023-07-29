namespace go addition.management

struct AdditionRequest {
    1: required string name;
    2: required i32 Age;
    3: required string email;
    4: required string address;
}

struct AdditionResponse {
    1: string status;
}

service AdditionManagement {
    AdditionResponse addStudent(1: AdditionRequest req);
}

