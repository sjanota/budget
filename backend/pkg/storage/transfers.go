package storage

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Storage) CreateTransfer(ctx context.Context, reportID models.MonthlyReportID, in *models.TransferInput) (*models.Transfer, error) {
	if err := s.validateTransferInput(ctx, reportID, in); err != nil {
		return nil, err
	}

	toInsert := &models.Transfer{
		ID:            primitive.NewObjectID(),
		Amount:        in.Amount,
		Title:         in.Title,
		FromAccountID: in.FromAccountID,
		ToAccountID:   in.ToAccountID,
		Date:          in.Date,
	}

	find := doc{
		"_id": reportID,
	}
	update := doc{
		"$push": doc{
			"transfers": toInsert,
		},
	}
	res, err := s.monthlyReports.UpdateOne(ctx, find, update)
	if err != nil {
		return nil, err
	} else if res.MatchedCount == 0 {
		return nil, ErrNoReport
	}
	return toInsert, nil
}

func (s *Storage) UpdateTransfer(ctx context.Context, reportID models.MonthlyReportID, id primitive.ObjectID, changeSet ChangeSet) (*models.Transfer, error) {
	find := doc{"_id": reportID, "transfers._id": id}
	project := doc{
		"transfers": doc{
			"$elemMatch": doc{
				"_id": id,
			},
		},
	}
	updateFields := doc{}
	for field, value := range changeSet.Changes() {
		updateFields["transfers.$."+field] = value
	}
	update := doc{
		"$set": updateFields,
	}
	res := s.monthlyReports.FindOneAndUpdate(ctx, find, update, options.FindOneAndUpdate().SetProjection(project).SetReturnDocument(options.After))
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	result := &models.MonthlyReport{}
	err := res.Decode(result)
	return result.Transfer(id), err
}

func (s *Storage) GetTransfersTotalForAccount(ctx context.Context, reportID models.MonthlyReportID, accountID primitive.ObjectID) (models.Amount, error) {
	result, err := s.monthlyReports.Aggregate(ctx, list{
		doc{"$match": doc{"_id": reportID}},
		doc{"$unwind": "$transfers"},
		doc{"$facet": doc{
			"to": list{
				doc{"$match": doc{"transfers.toaccountid": accountID}},
				doc{"$group": doc{
					"_id": nil,
					"val": doc{
						"$sum": "$transfers.amount",
					},
				}},
			},
			"from": list{
				doc{"$match": doc{"transfers.fromaccountid": accountID}},
				doc{"$group": doc{
					"_id": nil,
					"val": doc{
						"$sum": "$transfers.amount",
					},
				}},
			},
		}},
	})
	if err != nil {
		return models.NewAmount(), err
	}
	if !result.Next(ctx) {
		return models.NewAmount(), nil
	}
	type a struct {
		Val models.Amount
	}
	sums := struct {
		From []a
		To   []a
	}{}
	err = result.Decode(&sums)
	to := models.NewAmount()
	if len(sums.To) > 0 {
		to = sums.To[0].Val
	}
	from := models.NewAmount()
	if len(sums.From) > 0 {
		from = sums.From[0].Val
	}
	sub := to.Sub(from)
	return sub, err
}

func (s *Storage) validateTransferInput(ctx context.Context, reportID models.MonthlyReportID, in *models.TransferInput) error {
	budget, err := s.GetBudget(ctx, reportID.BudgetID)
	if err != nil {
		return err
	}
	if budget.Account(in.FromAccountID) == nil {
		return ErrInvalidReference
	}
	if budget.Account(in.ToAccountID) == nil {
		return ErrInvalidReference
	}
	if !reportID.Month.Contains(in.Date) {
		return ErrWrongDate
	}
	return nil
}

func (s *Storage) validateTransferUpdate(ctx context.Context, reportID models.MonthlyReportID, in *models.TransferUpdate) error {
	budget, err := s.GetBudget(ctx, reportID.BudgetID)
	if err != nil {
		return err
	}
	if in.FromAccountID != nil && budget.Account(*in.FromAccountID) == nil {
		return ErrInvalidReference
	}
	if in.ToAccountID != nil && budget.Account(*in.ToAccountID) == nil {
		return ErrInvalidReference
	}
	if in.Date != nil && !reportID.Month.Contains(*in.Date) {
		return ErrWrongDate
	}
	return nil
}
