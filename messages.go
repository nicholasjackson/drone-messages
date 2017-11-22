package messages

const (
	// MessageFlight is the name of a flight message for the drone
	MessageFlight = "drone.flight"
	// MessageFaceDetection is the name of a message when a new face has been detected
	MessageFaceDetection = "image.facedetection"
)

const (
	// CommandTakeOff instructs a drone taking off
	CommandTakeOff = "takeoff"
	// CommandLand instructs a drone to land
	CommandLand = "land"
)

// Flight defines a flight instruction message
type Flight struct {
	Command string
	Value   int
}

// FaceDetected defines a face detection message
type FaceDetected struct {
	Faces  []Rectangle
	Bounds Rectangle
}

// Rectangle defines a rectangular shape
type Rectangle struct {
	X      int
	Y      int
	Height int
	Width  int
}
