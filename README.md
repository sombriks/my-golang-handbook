# My Golang Study Handbook

Personal golang exercises from various sources

Each exercise sample a unique behavior and might or might not rely on
techniques already shown before.

## How to run each sample code

Enter inside each individual exercise folder and
`go run <filename>.go [arguments maybe]`.

A dedicated README.md file might exist sometimes, explaining how to run the
sample otherwise.

## What I am studying and what I am not

I am looking at general operations that can be sampled on small snippets.

I am following a meta-structure for multipurpose programming languages creating
sample cases for specific features that such kind of language are supposed to
provide in order to solve common problems that computers should resolve.

In a very simplistic way, i am exploring all possible ways to grab data from the
disk, put it into memory, do some calculations and then put it back into storage.

I am not structuring big brains project configs, there will be other projects
just to do that.

~~I am not into syntax abuse party~~ [I may do a little syntax abuse 
sometimes](exercises/0005-todo-list-file/todo-list-v1.go), yet all snippets
should be clean and objective, easy to understand.

I am not evangelizing the lang, the lang should evangelize by itself.

I am not here to say that one IDE/Editor is the best for the job, pick whatever
you want to punch code against the system.

## A roadmap-ish path

Besides following [a few](https://cursos.alura.com.br/formacao-go) 
[online](https://www.udemy.com/course/go-programming-language/learn/lecture/5985510?start=0#overview)
[courses](https://www.udemy.com/course/multithreading-in-go-lang/learn/lecture/18225828?start=0#overview),
there is no defined path here.

However, I expect a few key techniques to be sampled over those exercises:

- Hello world
- Handle low level (console) user input (interactive, parameters, env vars)
- Comparison
- Loops
- Basic I/O (files)
- Unit testing
- Concurrency I (classic mutexes and wait groups)
- Concurrency II (channels)
- Reflection
- Intermediate I/O (TCP/UDP)
- Database connection
- REST services
- Service Streams (kafka, nats, mqtt)
- Kubernetes Controllers and Operators

After that, mostly everything else ends up in reading docs from some library

## Study sources

Several links used to understand at least one specific thing resent in the
snippets.

- <https://discord.gg/golang>
- <https://gophers.slack.com/>
- <https://go.dev/tour/flowcontrol/9>
- <https://gobyexample.com/command-line-arguments>
- <https://gobyexample.com/structs>
- <https://gobyexample.com/maps>
- <http://golangtutorials.blogspot.com/2011/06/methods-on-structs.html>
- <https://gosamples.dev/string-to-bool/>
- <https://golangdocs.com/>
- <https://pkg.go.dev/fmt>
- <https://pkg.go.dev/os#OpenFile>
- <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world>
- <https://www.sohamkamani.com/golang/enums/>
- <https://stackoverflow.com/a/69785897/420096>
- <https://stackoverflow.com/a/51957548/420096>
- <https://stackoverflow.com/a/67307008/420096>
- <https://tutorialedge.net/golang/go-constructors-tutorial/>
- <https://www.freecodecamp.org/news/go-beginners-handbook/>
- <https://www.golangprograms.com/>
- <https://www.groundlabs.com/blog/anatomy-of-a-credit-card/>
- <https://zetcode.com/golang/rune/>
- <https://udemy.com/course/multithreading-in-go-lang/>
- <https://go.dev/blog/laws-of-reflection>
- <https://go.dev/doc/tutorial/database-access>
- <https://github.com/gostor/awesome-go-storage>
- <https://go.dev/doc/modules/managing-dependencies>
- <https://www.digitalocean.com/community/tutorials/how-to-use-dates-and-times-in-go>
- <https://go.dev/blog/using-go-modules>
- <https://gobyexample.com/json>
- <https://www.calhoun.io/6-tips-for-using-strings-in-go/>
- <https://go.dev/tour/methods/15>
- <https://pkg.go.dev/encoding/json#Unmarshal>
