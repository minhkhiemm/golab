# Go scheduler goroutine internals

The Go scheduler has three key structures that handle the workload of goroutines: 
M struct, P struct, and G struct.

## The M struct

The M struct is labeled M for **machine**.

The M struct is a representation of an OS thread.

It contains a pointer that points to the runnable goroutine global queue
(defined by the P struct). M retrieves its work from the P struct.
M contains the free and waiting goroutines that are ready to be executed.
Some notable M struct parameters are the following:
	- A goroutine that contains a scheduling stack(go)
	- **Thread local storage(tls)**
	- A P struct for executing Go code(p)

## The P struct

The P struct is labeled P for **processor**.

The P struct represents a logical processor. This is set by GOMAXPROCS.
Which should be equivalent to the number of cores available after go 1.5.

P maintains a queue of all of the goroutines(defined by the G struct).
When you invoke a new goroutine using the Go excutor, this new goroutine gets
inserted into P's queue. If P doesn't have an associated M struct, it will allocate
a new M.

Some notable P struct parameters are the following:
	- The P struct ID(id)
	- A back link to a associated M struct if applicable(m)
	- A pool of available defer structs(deferpool)
	- The queue of runnable goroutines(runq)
	- A struct of available Gs(gFree)

## The G struct

The G struct is labeled G for **goroutine**.

The G struct represents the stack parameters of a single goroutine.

It includes information on a couple of different parameters that are important
for a goroutine.

G structs get created for every new goroutine, as well as goroutines for the runtime.

Some notable G struct parameters are the following:
	- The current value of the stack pointers(stack.lo and stack.hi)
	- The current value of Go and C stack growth prologues (stackguard0 and stackguard1)
	- The current value of the M struct(m)

