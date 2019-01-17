# Reasons for concurrency and parallelism


To complete this exercise you will have to use git. Create one or several commits that adds answers to the following questions and push it to your groups repository to complete the task.

When answering the questions, remember to use all the resources at your disposal. Asking the internet isn't a form of "cheating", it's a way of learning.

 ### What is concurrency? What is parallelism? What's the difference?
 > *Concurrency means multiple tasks which start, run, and complete in overlapping time periods, in no specific order.If something is running concurrently is not necessarily running at the same time but in such a way that is perciebed to be parallel. Parallelism literally physically run parts of tasks or multiple tasks. Parallelism requires hardware with multiple processing units, essentially. In single core CPU, we may get concurrency but not parallelism.*
 
 ### Why have machines become increasingly multicore in the past decade?
 > *The reason can be derived from the last question. Since our world works in a parallel way it is logical tha we want to solve problems using paralellism. Other reason is that by moving to extra cores on the single processor chip, manufacturers avoided problems with the clock speeds by effectively multiplying the amount of data that could be handled by the CPU.*

 ### What kinds of problems motivates the need for concurrent execution?
 (Or phrased differently: What problems do concurrency help in solving?)
 > *Some of the driving forces for concurrency are external. That is, they are imposed by the demands of the environment. In real-world systems many things are happening simultaneously and must be addressed “in real-time” by software. To do so, many real-time software systems must be “reactive.” They must respond to externally generated events which may occur at somewhat random times, in some-what random order, or both.*
 
 ### Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
 (Come back to this after you have worked on part 4 of this exercise)
 > *It depends on the problem we are solving.For some problems it is essential to use concurrency,even if it complicates the work it makes the problem solvable.*
 
 ### What are the differences between processes, threads, green threads, and coroutines?
 > *-Both processes and threads are independent sequences of execution. The typical difference is that threads (of the same process) run in a shared memory space, while processes run in separate memory spaces.*
 > * *
 > *-Green threads are "user-level threads". They are scheduled by an "ordinary" user-level process, not by the kernel. So they can be used to simulate multi-threading on platforms that don't provide that capability.*
 > * *
 > *-Coroutines are a form of sequential processing: only one is executing at any given time.*
 
 ### Which one of these do `pthread_create()` (C/POSIX), `threading.Thread()` (Python), `go` (Go) create?
 > *-The `pthread_create()` function starts a new thread in the calling process.*
 > * *
 > *-`threading.Thread()` is the constructor of the class Thread.*
 > * *
 > *-The `go` is placed beforea function call to call it like a gorutine. A goroutine is a lightweight thread managed by the Go runtime.*
 
 ### How does pythons Global Interpreter Lock (GIL) influence the way a python Thread behaves?
 > *The Python Global Interpreter Lock or GIL, in simple words, is a mutex (or a lock) that allows only one thread to hold the control of the Python interpreter.
 This means that only one thread can be in a state of execution at any point in time. The impact of the GIL isn’t visible to developers who execute single-threaded programs, but it can be a performance bottleneck in CPU-bound and multi-threaded code.Since the GIL allows only one thread to execute at a time even in a multi-threaded architecture with more than one CPU core, the GIL has gained a reputation as an “infamous” feature of Python.*
 
 ### With this in mind: What is the workaround for the GIL (Hint: it's another module)?
 > *The most popular way is to use a multi-processing approach where you use multiple processes instead of threads. Each Python process gets its own Python interpreter and memory space so the GIL won’t be a problem. Python has a multiprocessing module which lets us create processes easily.*
 
 ### What does `func GOMAXPROCS(n int) int` change? 
 > *GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns the previous setting. If n < 1, it does not change the current setting. The number of logical CPUs on the local machine can be queried with NumCPU.*
