package storage_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/pkg/errors"

	"github.com/sjanota/budget/backend/pkg/models"
	mock_models "github.com/sjanota/budget/backend/pkg/models/mocks"
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

func before() context.Context {
	return context.Background()
}

func initDB() error {
	log.Println("Init DB")
	err := testStorage.Init(context.Background())
	if err != nil {
		return errors.Wrap(err, "while initializing DB")
	}
	return nil
}

func withDockerMongo(test func()) {
	defer func() {
		log.Println("Deleting mongo container")
		err := deleteMongoContainer()
		if err != nil {
			log.Println(err)
		}

		err = pruneVolumes()
		if err != nil {
			log.Println(err)
		}
	}()
	log.Println("Creating mongo container")
	port, err := testRunMongoContainer()
	if err != nil {
		panic(err)
	}

	testStorage, err = storage.New("mongodb://localhost:" + port + "/test-db")
	if err != nil {
		panic(errors.Wrap(err, "while creating testStorage"))
	}

	err = initDB()
	if err != nil {
		panic(err)
	}

	log.Println("DB port", port)
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

func pruneVolumes() error {
	cmd := exec.Command("docker", "volume", "prune")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("cannot prune volumes: %s", out)
	}
	return nil
}

func whenSomeBudgetExists(t *testing.T, ctx context.Context) *models.Budget {
	budget, err := testStorage.CreateBudget(ctx, "", mock_models.Month())
	require.NoError(t, err)
	return budget
}

func whenSomeEnvelopeExists(t *testing.T, ctx context.Context, budgetID primitive.ObjectID) *models.Envelope {
	input := &models.EnvelopeInput{Name: *mock_models.Name(), Limit: mock_models.Amount()}
	envelope, err := testStorage.CreateEnvelope(ctx, budgetID, input)
	require.NoError(t, err)
	return envelope
}

func whenSomeCategoryExists(t *testing.T, ctx context.Context, budgetID, envelopeID primitive.ObjectID) *models.Category {
	input := &models.CategoryInput{Name: *mock_models.Name(), EnvelopeID: envelopeID}
	category, err := testStorage.CreateCategory(ctx, budgetID, input)
	require.NoError(t, err)
	return category
}

func whenSomeAccountExists(t *testing.T, ctx context.Context, budgetID primitive.ObjectID) *models.Account {
	input := &models.AccountInput{Name: *mock_models.Name()}
	account, err := testStorage.CreateAccount(ctx, budgetID, input)
	require.NoError(t, err)
	return account
}

func whenSomeMonthlyReportExists(t *testing.T, ctx context.Context, budgetID primitive.ObjectID) *models.MonthlyReport {
	month := mock_models.Month()
	report, err := testStorage.CreateMonthlyReport(ctx, budgetID, month, make([]*models.PlanInput, 0))
	require.NoError(t, err)
	return report
}

func whenSomeExpenseExists(t *testing.T, ctx context.Context, accountID, categoryID1, categoryID2 primitive.ObjectID, report *models.MonthlyReport) *models.Expense {
	input := mock_models.ExpenseInput().
		WithDate(mock_models.DateInReport(report)).
		WithAccount(accountID).
		WithCategories(
			mock_models.ExpenseCategoryInput().WithCategory(categoryID1),
			mock_models.ExpenseCategoryInput().WithCategory(categoryID2),
		)
	expense, err := testStorage.CreateExpense(ctx, report.ID, input)
	require.NoError(t, err)
	return expense
}

func whenSomePlanExists(t *testing.T, ctx context.Context, envelopeID1, envelopeID2 primitive.ObjectID, report *models.MonthlyReport) *models.Plan {
	input := mock_models.PlanInput().
		WithTo(envelopeID1).
		WithFrom(&envelopeID2)
	plan, err := testStorage.CreatePlan(ctx, report.ID, input)
	require.NoError(t, err)
	return plan
}

func whenSomeTransferExists(t *testing.T, ctx context.Context, accountID1, accountID2 primitive.ObjectID, report *models.MonthlyReport) *models.Transfer {
	input := mock_models.TransferInput().
		WithTo(accountID1).
		WithFrom(&accountID2).
		WithDate(mock_models.DateInReport(report))
	transfer, err := testStorage.CreateTransfer(ctx, report.ID, input)
	require.NoError(t, err)
	return transfer
}