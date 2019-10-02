# blog-api-with-golang
This is a basic blog api designed using golang.

# Getting started
1. Install Golang on your machine.
  1.1. Make sure you have GOPATH set in your environment variables.
  1.2. Ensure it using `echo %GOPATH%`
2. Get this project by this command: `go get -u https://github.com/anirudhganwal06/blog-api-with-golang.git`
3. This will take some time because it downloads this project and downloads all the imported dependencies.
4. Now, `cd blog-api-with-golang`
5. Now, run a mogodb server on your local machine, which by default runs on port :27017.
6. Run `go build` to build the go project in a executable file.
7. Run the executable by just typing `blog-api-with-golang`
8. If the above command prompt Server running @5000, then you are good to go.

# Usage

>>>> Testing can be done using POSTMAN

## Endpoints
### POST "/api/blogs/create"
Enter three key values in Body (x-www-form-urlencoded) `title, content, author`
