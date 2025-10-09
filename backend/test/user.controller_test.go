package test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"main/internal/controllers"
	"main/internal/models"
)

var testClient *mongo.Client

const testDBName = "gdp-nexus"
const testCollectionName = "users"

// --- Funciones de Utilidad (Setup y Teardown) ---

// setupTestDB inicializa la conexión y limpia la colección de usuarios.
func setupTestDB(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	require.NoError(t, err, "Failed to connect to MongoDB. Is it running?")

	err = client.Ping(ctx, nil)
	require.NoError(t, err, "Failed to ping MongoDB")

	testClient = client

	// Limpieza: Borrar todos los documentos antes de la suite de pruebas.
	// _, err = testClient.Database(testDBName).Collection(testCollectionName).DeleteMany(context.Background(), bson.M{})
	// require.NoError(t, err, "Failed to clear the users collection")
}

// --- Suite de Pruebas Principal (CRUD) ---

func TestUsersCRUD(t *testing.T) {
	setupTestDB(t)

	testUser := models.User{
		Email:    "main.test.user@example.com",
		Password: "SecurePassword123",
	}

	defer func() {
		collection := testClient.Database(testDBName).Collection(testCollectionName)
		filter := bson.M{"email": testUser.Email}

		_, err := collection.DeleteOne(context.Background(), filter)
		if err != nil && err != mongo.ErrNoDocuments {
			log.Printf("Error during final teardown: could not delete user %s: %v", testUser.Email, err)
		}
	}()

	// ----------------------------------------------------
	// A. CREACIÓN EXITOSA (Setup de datos para las siguientes pruebas)
	// ----------------------------------------------------
	t.Run("A_CreateUser_Success", func(t *testing.T) {
		err := controllers.CreateUser(testClient, testUser)
		require.NoError(t, err, "Should create user successfully")

		t.Logf("Created user: %s", testUser.Email)
	})

	// ----------------------------------------------------
	// B. OBTENER USUARIO (Verifica la creación del paso A)
	// ----------------------------------------------------
	t.Run("B_GetUserByEmail_Found", func(t *testing.T) {
		foundUser, err := controllers.GetUserByEmail(testClient, testUser.Email)

		require.NoError(t, err, "Should find the user created in step A")
		require.Equal(t, testUser.Email, foundUser.Email, "Emails must match")

		t.Logf("Found user: %s", foundUser.Email)
	})

	// ----------------------------------------------------
	// C. CREACIÓN FALLIDA (Conflicto de email duplicado)
	// ----------------------------------------------------
	t.Run("C_CreateUser_Conflict", func(t *testing.T) {
		err := controllers.CreateUser(testClient, testUser)

		require.Error(t, err, "Should return an error for duplicate email")
		require.Contains(t, err.Error(), "Email is already registered", "Error message must indicate conflict")

		t.Logf("Conflict error: %s", err.Error())
	})

	// ----------------------------------------------------
	// D. OBTENER USUARIO FALLIDO (No encontrado)
	// ----------------------------------------------------
	t.Run("D_GetUserByEmail_NotFound", func(t *testing.T) {
		_, err := controllers.GetUserByEmail(testClient, "nonexistent.user@email.com")

		require.Error(t, err, "Should return an error when user is not found")
		require.Equal(t, mongo.ErrNoDocuments, err, "Error must be 'no documents found'")

		t.Logf("Not found error: %s", err.Error())
	})
}
