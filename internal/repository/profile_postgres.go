package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kovalyov-valentin/profiles-service/internal/models"
	"github.com/sirupsen/logrus"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (r *UserPostgres) CreateUser(ctx context.Context, user models.Users) (int, error) {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		logrus.WithError(err).Error("Error connecting to the database")
		return 0, err
	}
	defer conn.Close()

	var id int
	query := fmt.Sprintf(""+
		"INSERT INTO %s (name, surname, patronymic, age, gender, nationality) "+
		"VALUES ($1, $2, $3, $4, $5, $6) "+
		"RETURNING id", usersTable)

	row := conn.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Surname,
		user.Patronymic,
		user.Age,
		user.Gender,
		user.Nationality)

	if err := row.Err(); err != nil {
		logrus.WithError(err).Error("Error executing query")
		return 0, err
	}

	if err := row.Scan(&id); err != nil {
		logrus.WithError(err).Error("Error scanning row")
		return 0, err
	}

	logrus.WithFields(logrus.Fields{
		"Name":        user.Name,
		"Surname":     user.Surname,
		"Patronymic":  user.Patronymic,
		"Age":         user.Age,
		"Gender":      user.Gender,
		"Nationality": user.Nationality,
	}).Info("Record created successfully. ID:", id)
	return id, nil
}

func (r *UserPostgres) GetUser(ctx context.Context, id int) (models.Users, error) {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		logrus.WithError(err).Error("Error connecting to the database")
		return models.Users{}, err
	}
	defer conn.Close()

	query := fmt.Sprintf(
		"SELECT name, surname, patronymic, age, gender, nationality "+
			"FROM %s "+
			"WHERE id =$1",
		usersTable)

	var user models.Users
	if err := conn.GetContext(
		ctx,
		&user,
		query,
		id); err != nil {
		logrus.WithError(err).Error("Failed to execute database query")
		return models.Users{}, err
	}
	logrus.WithFields(logrus.Fields{
		"id":   id,
		"name": user.Name,
		// Добавьте другие поля, которые вам нужны для логирования
	}).Info("User record retrieved successfully")
	return user, err
}

func (r *UserPostgres) GetUsers(ctx context.Context) ([]models.Users, error) {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		logrus.WithError(err).Error("Error connecting to the database")
		return nil, err
	}
	defer conn.Close()

	query := fmt.Sprintf("SELECT * FROM %s", usersTable)

	var users []models.Users
	if err := conn.SelectContext(
		ctx,
		&users,
		query); err != nil {
		logrus.WithError(err).Error("Error executing database query")
		return nil, err
	}

	result := make([]models.Users, len(users))
	for i, user := range users {
		result[i] = models.Users(user)
	}

	logrus.Info("Successfully fetched all records from the database")
	return result, nil

}

func (r *UserPostgres) UpdateUser(ctx context.Context, id int, user models.Users) error {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		logrus.WithError(err).Error("Error connecting to the database")
		return err
	}
	defer conn.Close()

	query := fmt.Sprintf(
		"UPDATE %s "+
			"SET name=$1, surname=$2, patronymic=$3 "+
			"WHERE id=$4",
		usersTable)

	if _, err := conn.ExecContext(
		ctx,
		query,
		user.Name,
		user.Surname,
		user.Patronymic,
		id); err != nil {
		logrus.WithError(err).Error("Failed to execute update query")
		return err
	}

	logrus.Infof("Record with ID %d updated successfully", id)
	return nil
}

func (r *UserPostgres) DeleteUser(ctx context.Context, id int) error {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		logrus.WithError(err).Error("Error connecting to the database")
		return err
	}
	defer conn.Close()

	query := fmt.Sprintf(
		"DELETE FROM %s "+
			"WHERE id=$1",
		usersTable)

	result, err := conn.ExecContext(ctx, query, id)
	if err != nil {
		logrus.WithError(err).Error("Error executing DELETE query")
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		logrus.WithError(err).Error("Error getting RowsAffected")
		return err
	}

	if row == 0 {
		return fmt.Errorf("Record with ID %d not found", id)

	}
	logrus.WithField("ID", id).Info("Record deleted successfully")
	return nil
}
