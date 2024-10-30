package repository

import (
	"context"
	"fmt"
	"golphin/internal/database"
	user "golphin/src/domain/user"
	"golphin/src/utils/encrypt"
	"log"
)

func insertUser(db *database.Queries, user user.User) error {
    err := db.InsertUser(context.Background(), database.InsertUserParams{
		ID: 	  *user.ID(),
        Username: user.Username(),
        Email:    user.Email(),
        Password: user.Password(),
    })
    if err != nil {
        return fmt.Errorf("failed to insert user: %w", err)
    }
    log.Println("User inserted successfully")
    return nil
}

func Auth(user user.User) error {
	dbConn, err := database.ConnectDB()
	if err != nil {
        return fmt.Errorf("failed to connect to database: %w", err)
    }
	defer dbConn.Close()

	db := database.New(dbConn)
	err = insertUser(db, user)
	if err != nil {
		return err
	}
	return nil
}

func verifyUser(db *database.Queries, usernameOrEmail, password string) (*database.User, error) {
    user, err := db.FindUser(context.Background(), database.FindUserParams{
		Username: usernameOrEmail,
        Email:    usernameOrEmail,
    })
    if err != nil {
        return nil, fmt.Errorf("failed to find user: %w", err)
    }

    if !encrypt.IsPasswordValid(user.Password, password) {
        return nil, fmt.Errorf("invalid credentials")
    }

    return &user, nil
}

func Login(desc, password string) (string, error) {
    dbConn, err := database.ConnectDB()
    if err != nil {
        return "", fmt.Errorf("failed to connect to database: %w", err)
    }
    defer dbConn.Close()

    db := database.New(dbConn)
    dbUser, dbErr := verifyUser(db, desc, password)
    if dbErr != nil {
        return "", dbErr
    }
    
    if dbUser == nil {
        return "", fmt.Errorf("invalid credentials")
    }
    return dbUser.Username, nil
}