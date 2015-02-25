package runner

import "github.com/mgutz/dat"

// Session represents a business unit of execution for some connection
type Session struct {
	*Tx
}

// NewSession instantiates a Session for the Connection
func (conn *Connection) NewSession() (*Session, error) {
	tx, err := conn.Begin()
	if err != nil {
		return nil, err
	}
	return &Session{tx}, nil
}

// Close closes the session.
func (sess *Session) Close() error {
	err := sess.Tx.AutoCommit()
	if err != nil {
		dat.Events.EventErr("session.close", err)
	}
	return err
}