package sarif

// Notification represents different Notifications in the SARIF spec
type Notification struct {
	PropertyBag
	Level      *string     `json:"level,omitempty"`
	Message    Message     `json:"message"`
	Locations  []*Location `json:"locations,omitempty"`
	Properties Properties  `json:"properties,omitempty"`
}

// NewNotification ...
func NewNotification() *Notification {
	return &Notification{}
}

func newLevelNotification(level string) *Notification {
	return NewNotification().WithLevel(level)
}

// WithLevel ...
func (n *Notification) WithLevel(level string) *Notification {
	n.Level = &level
	return n
}

// WithMessage ...
func (n *Notification) WithMessage(message *Message) *Notification {
	n.Message = *message
	return n
}

// WithLocation ...
func (n *Notification) WithLocation(location *Location) *Notification {
	n.Locations = append(n.Locations, location)
	return n
}

// WithProperties specifies properties for a rule and returns the updated rule
func (n *Notification) WithProperties(properties Properties) *Notification {
	n.Properties = properties
	return n
}

// AttachPropertyBag adds a property bag to a rule
func (n *Notification) AttachPropertyBag(pb *PropertyBag) {
	n.Properties = pb.Properties
}
