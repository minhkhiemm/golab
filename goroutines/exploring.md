# Exploring goroutines

Go is a language designed with concurrency in mind.
Concurrency is the ability to execute independent processes.
Goroutines are a construct in Go that can help with concurrency.
They are often referred to as *lightweight threads*.
In other languages, threads are handled by the OS.
This in turn, uses a larger-sized call stack and usually handles less concurrency with a given memory stack size.
Goroutines are functions or methods that run within the Go runtime concurrently and don't connect to the underlying OS.
The scheduler within the Go language manages goroutines' life cycles. 
The system's scheduler has a lot of overhead as well, so limiting the number of threads being utilized can help to improve performance.

## The Go scheduler

There are a couple of different pieces involved in the management of goroutine life cycles by the Go runtime scheduler.
The Go scheduler was changed in its second iteration, which was derived from a design document written by Dmitry Vyukov,
released in Go 1.1. In this design doc, Vyukov discusses te initial Go scheduler and how to implement a work-sharing and work-stealing scheduler, as originally prescribed by Dr Robert D. Blumofe and Dr. Charles E. Leiserson in MIT paper entitled, *Scheduling Multithreaded Computations by Work Stealing*.
The fundamental concept behind this paper is to ensure dynamic, multithreaded computation in order to ensure that processors are utilized efficiently while maintaining memory requirements.

Goroutines only have a stack size of 2 KB on inception. This is one of the reasons why goroutines
are preferred for a lot of concurrent programming-because it is much easier to have tens or hundreds of 
thousands of goroutines in one program. Threads in other languages can take up megabytes of space,
making them a lot less flexible. If more memory is needed, Go's functions have the ability to allocate more
memory in another place in available memory space to help the goroutine space grow. By default, the runtime gives the new
stack twice the amount of memory.

Goroutines block a running thread onl on system calls. when this occures, the runtime 
takes another thread from the scheduler struct. These are used for other goroutines
that are waiting to be executed.

Work sharing is a process in which a scheduler migrates new threads to other processors for
work distribution. Work stealing performs a similar action, but in which the underutilized processors
steal threads from other processors.  Following the work-stealing pattern much more efficient and, in turn,
gives higher throughput to the goroutines that run on top of the kernel's scheduler. Lastly, GO's scheduler
implements spinning threads. Spinning threads will utilize extra CPU cycles over preempting a thread.
Threads spin in three different ways:
    - When a thread is not attached to a processor.
    - When making a goroutine ready will unblock an OS thread onto an idle processor.
    - When a thread is running but no goroutines are attached to it. This idle thread will continue to search
for runnable goroutines to execute.