package storage

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Storage) CreatePlan(ctx context.Context, reportID models.MonthlyReportID, in *models.PlanInput) (*models.Plan, error) {
	if err := s.validatePlanInput(ctx, reportID, in); err != nil {
		return nil, err
	}

	toInsert := &models.Plan{
		ID:              primitive.NewObjectID(),
		CurrentAmount:   in.CurrentAmount,
		Title:           in.Title,
		FromEnvelopeID:  in.FromEnvelopeID,
		ToEnvelopeID:    in.ToEnvelopeID,
		RecurringAmount: in.RecurringAmount,
	}

	find := doc{
		"_id": reportID,
	}
	update := doc{
		"$push": doc{
			"plans": toInsert,
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

func (s *Storage) UpdatePlan(ctx context.Context, reportID models.MonthlyReportID, id primitive.ObjectID, update models.PlanUpdate) (*models.Plan, error) {
	if err := s.validatePlanUpdate(ctx, reportID, update); err != nil {
		return nil, err
	}

	find := doc{"_id": reportID, "plans._id": id}
	project := doc{
		"plans": doc{
			"$elemMatch": doc{
				"_id": id,
			},
		},
	}
	updateFields := doc{}
	for field, value := range update {
		updateFields["plans.$."+field] = value
	}
	updateDoc := doc{
		"$set": updateFields,
	}
	res := s.monthlyReports.FindOneAndUpdate(ctx, find, updateDoc, options.FindOneAndUpdate().SetProjection(project).SetReturnDocument(options.After))
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	result := &models.MonthlyReport{}
	err := res.Decode(result)
	return result.Plan(id), err
}

func (s *Storage) DeletePlan(ctx context.Context, reportID models.MonthlyReportID, id primitive.ObjectID) (*models.Plan, error) {
	find := doc{"_id": reportID}
	project := doc{
		"plans": doc{
			"$elemMatch": doc{
				"_id": id,
			},
		},
	}
	updateDoc := doc{
		"$pull": doc{"plans": doc{"_id": id}},
	}
	res := s.monthlyReports.FindOneAndUpdate(ctx, find, updateDoc, options.FindOneAndUpdate().SetProjection(project).SetReturnDocument(options.Before))
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	result := &models.MonthlyReport{}
	err := res.Decode(result)
	return result.Plan(id), err
}

func (s *Storage) GetPlansTotalForEnvelope(ctx context.Context, reportID models.MonthlyReportID, accountID primitive.ObjectID) (models.Amount, error) {
	result, err := s.monthlyReports.Aggregate(ctx, list{
		doc{"$match": doc{"_id": reportID}},
		doc{"$unwind": "$plans"},
		doc{"$facet": doc{
			"to": list{
				doc{"$match": doc{"plans.toenvelopeid": accountID}},
				doc{"$group": doc{
					"_id": nil,
					"val": doc{
						"$sum": "$plans.amount",
					},
				}},
			},
			"from": list{
				doc{"$match": doc{"plans.fromenvelopeid": accountID}},
				doc{"$group": doc{
					"_id": nil,
					"val": doc{
						"$sum": "$plans.amount",
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

func (s *Storage) validatePlanInput(ctx context.Context, reportID models.MonthlyReportID, in *models.PlanInput) error {
	budget, err := s.GetBudget(ctx, reportID.BudgetID)
	if err != nil {
		return err
	}
	if in.FromEnvelopeID != nil && budget.Envelope(*in.FromEnvelopeID) == nil {
		return ErrInvalidReference
	}
	if budget.Envelope(in.ToEnvelopeID) == nil {
		return ErrInvalidReference
	}
	return nil
}

func (s *Storage) validatePlanUpdate(ctx context.Context, reportID models.MonthlyReportID, in models.PlanUpdate) error {
	budget, err := s.GetBudget(ctx, reportID.BudgetID)
	if err != nil {
		return err
	}
	if id, has := in.ToEnvelopeID(); has && budget.Envelope(id) == nil {
		return ErrInvalidReference
	}
	if id, has := in.FromEnvelopeID(); has && id != nil && budget.Envelope(*id) == nil {
		return ErrInvalidReference
	}
	return nil
}
