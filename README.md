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
1. Create a blog > POST "/api/blogs/create" > Enter three key values in Body (x-www-form-urlencoded) `title, content, author`
2. Get all blogs > GET "/api/blogs" > This get an array of all blogs.
3. Get a specific blog > GET "/api/blog/{_id}" > This gets the object of the specified blog id.
4. Delete a specific blog > DELETE "/api/blog/{_id}/delete" > This deletes the blog identified by the specified id. 
5. Update a specific blog > PUT "/api/blog/{_id}/update" > This updates the specified blog. This also requires three key values in Body (x-www-form-urlencoded) `title, content, author`
