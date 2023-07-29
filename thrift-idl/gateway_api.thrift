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

struct College {
    1: required string name(go.tag = 'json:"name"'),
    2: string address(go.tag = 'json:"address"'),
}

struct Student {
    1: required i32 id(api.body='id'),
    2: required string name(api.body='name'),
    3: required string sex(api.body='sex'),
    4: required i32 age(api.body='age'),
    5: required College college(api.body='college'),
    6: optional list<string> email(api.body='email'),
}

struct RegisterResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

struct QueryReq {
    1: required i32 id(api.query='id')
}

//----------------------service-------------------
service StudentService {
    RegisterResp Register(1: Student student)(api.post = '/add-student-info')
    Student Query(1: QueryReq req)(api.get = '/query')
}

service Gateway {
   AdditionResponse addNumbers(1: AdditionRequest req) (api.post="/add");
   MultiplicationResponse multiplyNumbers(1: MultiplicationRequest req) (api.post="/multiply");
   DivisionResponse divideNumbers(1: DivisionRequest req) (api.post="/divide");
}