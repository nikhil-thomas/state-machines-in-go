# Minimal StateMachine
*minimal statemachine fitting addon-service use case*

## Run

```go
go run main.go
```

## Sample Output
```text
Attempt transition pending-state to ready-state
before: current state: pending-state
after: current state: pending-state

Attempt transition pending-state to installing-state
before: current state: pending-state
after: current state: installing-state

Attempt transition installing-state to ready-state
before: current state: installing-state
after: current state: ready-state
```