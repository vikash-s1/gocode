package components

// RegularUser represents a standard chat user
type RegularUser struct {
	*BaseUser
}

// NewRegularUser creates a new RegularUser
func NewRegularUser(name string) *RegularUser {
	return &RegularUser{
		BaseUser: NewBaseUser(name),
	}
}

// HandleNotification handles notifications specific to regular users
func (ru *RegularUser) HandleNotification(event string, data interface{}) {
	// Regular users handle all standard notifications
	ru.BaseUser.HandleNotification(event, data)
}