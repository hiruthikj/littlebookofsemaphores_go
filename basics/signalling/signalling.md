# Signalling

Signalling of one goroutine to another goroutine indicating something is done.
Could be implemented in channels (idiomatic way), condition variables (best for multiple goroutines) and waitGroup (not preferred)

## Definitions/Analogies:
* WaitGroup - a checklist, waits for all to be done
* channel - messaging between goroutines (and syncing  too due to blocking nature of adding and removing)
* Condition Variable - waits for a condition (rendezvous point) to happen, only then all goroutines proceeds

Ref: 
https://hackernoon.com/understanding-synccond-in-go-a-guide-for-beginners