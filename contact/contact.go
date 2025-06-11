package contact

type Contact struct {
	Name  string
	Phone string
	Email string
}

type Manager struct {
	contacts map[string]Contact
}

func NewManager() *Manager {
	return &Manager{contacts: make(map[string]Contact)}
}

func (m *Manager) AddContact(contact Contact) {
	m.contacts[contact.Email] = contact
}

func (m *Manager) GetContactByEmail(email string) (Contact, bool) {
	contact, ok := m.contacts[email]
	return contact, ok
}

func (m *Manager) RemoveContactByEmail(email string) bool {
	_, ok := m.contacts[email]
	if ok {
		delete(m.contacts, email)
		return true
	}
	return false
}

func (m *Manager) ContactList() []Contact {
	contacts := make([]Contact, 0, len(m.contacts))

	for _, contact := range m.contacts {
		contacts = append(contacts, contact)
	}
	return contacts
}
