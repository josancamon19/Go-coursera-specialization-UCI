# Week 2

**Processes**

- An instance of a running program
- Things unique to a process
    1. Memory 
    - Virtual address space
    - Code, stack, heap, shared libraries
    2. Registers
    - Program counter, data regs, stack ptr,

**Operating Systems**

- Allows many processes to execute concurrently
- Processes are switched quickly **20ms**
- User has the impression of parallelism
- Operating system must give processes fair access to resources

Scheduling Processes

- Operating system schedules processes for execution
- Gives the illusion of parallel execution

    ![Week%202%200adcb6ce75c54adeb3e530fe6020c01c/Screen_Shot_2020-11-15_at_9.39.47_PM.png](Week%202%200adcb6ce75c54adeb3e530fe6020c01c/Screen_Shot_2020-11-15_at_9.39.47_PM.png)

- OS gives fair access to CPU, memory, etc.

**Context Switch**

- Control flow changes from one process to another
- Process "context" must be swapped

    ![Week%202%200adcb6ce75c54adeb3e530fe6020c01c/Screen_Shot_2020-11-15_at_9.40.38_PM.png](Week%202%200adcb6ce75c54adeb3e530fe6020c01c/Screen_Shot_2020-11-15_at_9.40.38_PM.png)

### **Threads vs. Processes**

- Threads share some context
- Many threads can exist in one process
- OS Schedules threads rather than processes

**Goroutines**

- Like a thread in Go
- Many Goroutines execute within a single OS thread

    ![Week%202%200adcb6ce75c54adeb3e530fe6020c01c/Screen_Shot_2020-11-15_at_9.43.10_PM.png](Week%202%200adcb6ce75c54adeb3e530fe6020c01c/Screen_Shot_2020-11-15_at_9.43.10_PM.png)

**Go Runtime Scheduler**

- Schedules goroutines inside an OS thread
- Like a little OS inside a single OS thread
- Logical processor is mapped to a thread

    ![Week%202%200adcb6ce75c54adeb3e530fe6020c01c/Screen_Shot_2020-11-15_at_9.44.08_PM.png](Week%202%200adcb6ce75c54adeb3e530fe6020c01c/Screen_Shot_2020-11-15_at_9.44.08_PM.png)

**Interleavings**

- Each 20 ms change within threads →

    Happens at machine code, is not deterministic. 
    E.g. i += 1 → `reads i, increases i, writes i`

- Order of execution within a task is known
- Order of execution between concurrent tasks is unknown

**Race Conditions**

- Outcome depends on non-deterministic ordering
- Races occur due to **communication**

    ![Week%202%200adcb6ce75c54adeb3e530fe6020c01c/Screen_Shot_2020-11-15_at_9.49.08_PM.png](Week%202%200adcb6ce75c54adeb3e530fe6020c01c/Screen_Shot_2020-11-15_at_9.49.08_PM.png)