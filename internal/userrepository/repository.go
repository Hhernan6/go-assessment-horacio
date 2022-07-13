package userrepository

import (
	"database/sql"
	"fmt"
)

type Client struct {
	DB *sql.DB
}

// GetUser() will get a user from the databasa
func (c *Client) GetUser(userId string) (User, error) {
	var user User

	row := c.DB.QueryRow("SELECT * FROM badass_users WHERE id = ?", userId)

	err := row.Scan(&user)
	if err != nil {
		return User{}, fmt.Errorf("error getting user %s: %w", userId, err)
	}

	return user, nil
}

// DeleteUser() will delete a user from the database
func (c *Client) DeleteUser(userId string) error {
	_, err := c.DB.Exec("DELETE FROM badass_users WHERE id = ?", userId)
	if err != nil {
		return fmt.Errorf("error deleting user %s: %w", userId, err)
	}

	return nil
}

// UpdateUser() will update a user in the database
func (c *Client) UpdateUser(user User) error {
	_, err := c.DB.Exec("UPDATE badass_users SET first_name=?, last_name=? WHERE id = ?", user.FirstName, user.LastName, user.Id)
	if err != nil {
		return fmt.Errorf("error updating user %s: %w", user.Id, err)
	}

	return nil
}

// CreateUser() will update a new user in the database
func (c *Client) CreateUser(firstName, lastName string) error {
	_, err := c.DB.Exec("INSERT INTO badass_users(first_name, last_name) VALUES(?,?)", firstName, lastName)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}
