# go

- Simpler objects, now this could be an advantage or disadvantage, but the idea is that Go is essentially object oriented although, some might disagree, but it has this concept of objects, and these objects are generally, the object orientation is a little simplified as compared to other languages. So, this is good, it makes it easier to code.

- Go is actually a simpler object-oriented implementation than you would see in other languages like C++.

- Okay, now, this translation step, it can go on in one of roughly two ways. It can be compiled, it can be a compilation or interpretation, okay?

- Now, a compiled language is a language where the translation from high-level language to machine code happens one time before you execute the code, before you deploy the code, happens one time, okay? So, like in C, C++, Go, Java partially, there's a compiler, and you compile the code.

- So, the idea behind a compiled languages, the key thing we want to bring out anyway is the fact that this translation occurs once, it doesn't occur while you're running the code, right?

- So, the other way to do this is interpreted, interpretation. In interpreted language what happens is, instructions are translated while the code is executed.

- So, in addition to actually executing the instruction, you've got to do this translation from the instruction into the equivalent machine code, so that slows you down. So, the translation occurs every time you run the Python code, say or Java, the Java byte code.

- One big difference you can see is, its compiled code is generally faster to execute, that's because you don't have to do the translation every time you run the code, so it is going to be faster. 

- So, the thing about interpreters is that, the interpreter itself, that program that is doing the translation of your code, it can help you, it can handle things that you, as a programmer, don't want to handle. For instance, in Python, I don't have to declare my variable types. I can just start using a variable and the interpreter will say, "It looks like he's using it as an integer, make it an integer." So, that's something that the programmer doesn't have to think about. 

- Another thing that interpreters commonly often have is memory management. 

- that is when you're done using an object, you want to get rid of that, de-allocate it from memory, and that happens automatically in an interpreted language, so the interpreter handles that. So, that's a good thing about interpreters. 

- It's a compiled language, but it has some of the good features of interpreted language, specifically it has garbage collection. So, garbage collection is the automatic memory management.

- So, Go has a garbage collection code included, so when it compiles your code, it also compiles garbage collection into your code automatically, and this is typically only done by interpreter. 

- So, this is a compiled language that actually has garbage collection, which is a really good feature. Now, downside is that it slows down execution a bit, but it's an efficient garbage collector, so it doesn't slow down much and you get a lot of advantage of having this automatic garbage collection.


# Go Tools

### go build - compiles the program

- arguments can be a list of packages or a list of .go files
- creates an executable for the main packages, same name as the first .go file
- .exe suffix for executable in windows


### go doc - prints documentation for a package 


### go fmt - formats source code files


### go get - downloads packages and installs them


### go list - lists all installed packages


### go run - compiles .go files and runs the executable


### go test - runs tests using files ending in _test.go




# variables

- all variables must have declarations by name and type.
like : var x int


## variable Scope

- the places in code where a variable can be accessed.






