package storage_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/sjanota/budget/backend/pkg/storage/mock"

	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/stretchr/testify/require"

	"github.com/sjanota/budget/backend/pkg/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	containerName = "budget-storage-tests-mongo"
)

var (
	testStorage *storage.Storage
)

func TestMain(m *testing.M) {
	var retCode int
	withDockerMongo(func() {
		retCode = m.Run()
	})
	os.Exit(retCode)
}

func before(t *testing.T) context.Context {
	drop(t)
	return context.Background()
}

func drop(t *testing.T) {
	t.Log("Drop DB")
	err := testStorage.Drop(context.Background())
	if err != nil {
		t.Fatalf("Cannot drop DB: %s", err)
	}

	t.Log("Init DB")
	err = testStorage.Init(context.Background())
	if err != nil {
		t.Fatalf("Cannot init DB: %s", err)
	}
}

func withDockerMongo(test func()) {
	defer func() {
		log.Println("Deleting mongo container")
		err := deleteMongoContainer()
		if err != nil {
			panic(err)
		}
	}()
	log.Println("Creating mongo container")
	port, err := testRunMongoContainer()
	if err != nil {
		panic(err)
	}

	fmt.Println(port)
	testStorage, err = storage.New("mongodb://localhost:" + port + "/test-db")
	if err != nil {
		panic(fmt.Errorf("cannot create testStorage: %s", err))
	}

	log.Println("Port", port)
	log.Println("Running tests")
	test()
}

func testRunMongoContainer() (string, error) {
	cmd := exec.Command("docker", "create", "--expose=27017", "-P", "--name="+containerName, "mongo:3.6")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("cannot create mongo container: %s", out)
	}
	cmd = exec.Command("docker", "start", containerName)
	out, err = cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("cannot create mongo container: %s", out)
	}
	cmd = exec.Command("docker", "inspect", "-f='{{ (index (index .NetworkSettings.Ports \"27017/tcp\") 0).HostPort }}'", containerName)
	out, err = cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("cannot get mongo container port: %s", out)
	}
	return strings.Trim(string(out), "'\n"), nil
}

func deleteMongoContainer() error {
	cmd := exec.Command("docker", "rm", "-f", containerName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("cannot delete mongo container: %s", out)
	}
	return nil
}

func whenSomeBudgetExists(t *testing.T, ctx context.Context) *models.Budget {
	budget, err := testStorage.CreateBudget(ctx, primitive.NewObjectID(), primitive.NewObjectID())
	require.NoError(t, err)
	return budget
}

func whenSomeEnvelopeExists(t *testing.T, ctx context.Context, budgetID primitive.ObjectID) *models.Envelope {
	input := &models.EnvelopeInput{Name: mock.Name(), Limit: mock.Amount()}
	envelope, err := testStorage.CreateEnvelope(ctx, budgetID, input)
	require.NoError(t, err)
	return envelope
}

func whenSomeCategoryExists(t *testing.T, ctx context.Context, budgetID, envelopeID primitive.ObjectID) *models.Category {
	input := &models.CategoryInput{Name: mock.Name(), EnvelopeID: envelopeID}
	category, err := testStorage.CreateCategory(ctx, budgetID, input)
	require.NoError(t, err)
	return category
}

func whenSomeAccountExists(t *testing.T, ctx context.Context, budgetID primitive.ObjectID) *models.Account {
	input := &models.AccountInput{Name: mock.Name()}
	account, err := testStorage.CreateAccount(ctx, budgetID, input)
	require.NoError(t, err)
	return account
}
