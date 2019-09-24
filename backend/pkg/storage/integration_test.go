package storage_test

import (
	"context"
	"fmt"
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
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

func beforeWithBudget(t *testing.T) (context.Context, *models.Budget, func()) {
	ctx, cancel := context.WithCancel(before(t))
	name := primitive.NewObjectID().Hex()

	budget, err := testStorage.Budgets().Create(ctx, name)
	require.NoError(t, err)

	return ctx, budget, func() {
		_, _ = testStorage.Budgets().Delete(ctx, budget.ID)
		cancel()
	}
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
	cmd := exec.Command("docker", "create", "--expose=27017", "-P", "--name="+containerName, "mongo:4.1")
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

func strPtr(s string) *string {
	return &s
}

func idPtr(id primitive.ObjectID) *primitive.ObjectID {
	return &id
}
