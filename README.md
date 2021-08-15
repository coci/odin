# odin
Odin is static blog engine written with golang. Odin uses github pages to deploy your blog.


### Requirements:
- support Unix (linux, MacOs , ....)


### Installation
``
go get github.com/coci/odin
``


## Usage :
1- first you need initialize odin :
```bash
odin init
```
2- you need configure odin :
```bash
odin config
```
3- you need make post:
```bash
odin new "how to work with odin"
```
this command will generate post template for you in :

``
path-to-blog-dir/content
``

4- you need to build your blog :
```bash
odin build
```

5- you have to push your blog :
```bash
odin push
```

## Consideration:

**note :** you need to run step 1 and 2 for once in lifetime.

**note :** you need to write your blog post in markdown format.

**note :** you can find post template ( file that you have to write your blog) in 'content' directory in root of blog.

**note :** everytime you change existing post or add new post you have to **build** and **push** again( step 4 and 5).

**note :** change CNAME file in root of project with your domain.
