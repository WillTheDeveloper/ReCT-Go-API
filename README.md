# ReCT-Go-API
New and improved API for the ReCT community of developers.

## Usage

1. Clone the repository
2. Open in root directory of repository
3. Run ``` go tidy ``` to get packages
4. Run ``` go run . ``` to start the server
5. Start making requests. Available requests are below.

## Available requests

| Request | Returns | Method |
|---------|---------|--------|
| /users | List of users | GET |
| /projects | List of projects | GET |
| /projects/{id} | Returns a project with matching ID | GET |
| /users | Create a new user | POST |
| /users/{id} | Returns a user with matching ID | GET |
| /projects | Create a new project | POST |