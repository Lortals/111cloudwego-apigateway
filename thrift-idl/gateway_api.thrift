namespace go api

struct AdditionRequest {
    1: required string FirstNum (api.body="FirstNum");
    2: required string SecondNum (api.body="SecondNum")
}

struct AdditionResponse {
    1: string Sum;
}

struct MultiplicationRequest {
    1: required string FirstNum (api.body="FirstNum");
    2: required string SecondNum (api.body="SecondNum")
}

struct MultiplicationResponse {
    1: string Product;
}

struct DivisionRequest {
    1: required string FirstNum (api.body="FirstNum");
    2: required string SecondNum (api.body="SecondNum")
}

struct DivisionResponse {
    1: string Quotient;
}

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

service Gateway {
   AdditionResponse addNumbers(1: AdditionRequest req) (api.post="/add");
   MultiplicationResponse multiplyNumbers(1: MultiplicationRequest req) (api.post="/multiply");
   DivisionResponse divideNumbers(1: DivisionRequest req) (api.post="/divide");
   RegisterStudentResponse registerStudent(1: RegisterStudentRequest req) (api.post="/student/register");
   GetStudentResponse getStudent(1: GetStudentRequest req) (api.post="/student/get");
}