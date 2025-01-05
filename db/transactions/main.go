package main

import (
	"context"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Entity represents a database entity
type Entity struct {
	ID   uint
	Name string
}

// Repository defines the methods for database operations and transaction handling
type Repository interface {
	Create(ctx context.Context, entity *Entity) error
	Update(ctx context.Context, entity *Entity) error
	Transaction(ctx context.Context, fn func(repo Repository) error) error
}

// repository implements the Repository interface and holds the GORM DB instance
type repository struct {
	db *gorm.DB
}

// withTx creates a new repository instance with the given transaction
func (r *repository) withTx(tx *gorm.DB) Repository {
	return &repository{
		db: tx,
	}
}

// Transaction manages the transaction lifecycle
func (r *repository) Transaction(ctx context.Context, fn func(repo Repository) error) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	repo := r.withTx(tx)
	err := fn(repo)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// Create adds a new entity to the database
func (r *repository) Create(ctx context.Context, entity *Entity) error {
	return r.db.Create(entity).Error
}

// Update modifies an existing entity in the database
func (r *repository) Update(ctx context.Context, entity *Entity) error {
	return r.db.Save(entity).Error
}

// Service handles the business logic
type Service struct {
	repo Repository
}

// PerformBusinessLogic performs multiple database operations within a transaction
func (s *Service) PerformBusinessLogic(ctx context.Context) error {
	return s.repo.Transaction(ctx, func(repo Repository) error {
		if err := repo.Create(ctx, &Entity{Name: "Example"}); err != nil {
			return err
		}
		return repo.Update(ctx, &Entity{ID: 1, Name: "Updated Example"})
	})
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Entity{})

	repo := &repository{db: db}
	service := &Service{repo: repo}

	ctx := context.Background()
	if err := service.PerformBusinessLogic(ctx); err != nil {
		fmt.Printf("Transaction failed: %v\n", err)
	} else {
		fmt.Println("Transaction succeeded")
	}
}
