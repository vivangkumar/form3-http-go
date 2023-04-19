package fixtures

var BadRequestError = `{
  "error_message": "Message parsing failed: Unexpected character (';' (code 34)): was expecting comma to separate Object entries ",
  "error_code": "d0a17902-63ed-4cb6-a8e8-fac5ca31b0b7"
}`

var ForbiddenError = `{
  "error": "invalid_grant",
  "error_description": "Wrong email or password."
}`

var ConflictError = `{
  "error_message": "Duplicate id f72c5098-bf0f-4526-a215-54e5c1e2e687",
  "error_code": "4bc0fa5d-231e-43f3-af79-8fc371d95a31"
}`
