//go:generate go run ${PWD}/../cmd/gents/main.go -input ${PWD}/events.go -output ${PWD}/../frontend/src/events.gen.ts

package backend

const (
	EventAppReady   = "app:ready"
	EventAppClose   = "app:close"
	EventDataUpdate = "data:update"
	EventUserLogin  = "user:login"
)
