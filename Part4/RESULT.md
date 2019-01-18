RESULTS PART 4
--------------------

### c implementattion
> *The c routine works as spected returning 0.*

### go implementattion
> *The go routine works as spected returning th result 0. I have discovered that the GOMAXPROCS of my computer is already setted with the number of CPUs before using "runtime.GOMAXPROCS(runtime.NumCPU())"*

### python implementattion
> *The python routine is the only one that differ from the 0 result. The routine returns a large positive number that vary over the runs. This is due to the GIL. The global variable i is overwritten in an unconsistent way since the program is not runing in in parallel.*




