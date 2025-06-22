# Rendezvous
A rendezvous point for goroutines waiting for or announcing the 
occurrence of an event.

Example:
A1 -> B2
A1 should be done before B2
B1 -> A2
B1 should be done before A

Cond Var is especially made for this. 

## Condition Variable
### Consists
* The Mutex protects the shared data that represents the condition.
* The Cond manages the waiting and waking of goroutines that depend on that condition.

### Methods:
* Wait: Suspends goroutine and allows other goroutine to progress.
    - Releases withheld lock (needed for others to progress to change the awaited condition)
    - Suspends current goroutine
    - Awakens later and reacquires lock
* Signal: Wakes 1 waiting goroutine
* Broadcast: Wakes all waiting goroutine

NOTE: Without Signal/Broadcast, wait won't complete.


