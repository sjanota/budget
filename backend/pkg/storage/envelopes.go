package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type envelopesRepository struct {
	*repository
	storage *Storage
}

func newEnvelopesRepository(s *Storage) *envelopesRepository {
	return &envelopesRepository{
		repository: &repository{
			storage:    s,
			collection: s.db.Collection("envelopes"),
		},
	}
}

type Envelopes struct {
	*envelopesRepository
	budgetID primitive.ObjectID
}

func (r envelopesRepository) session(budgetID primitive.ObjectID) *Envelopes {
	return &Envelopes{
		envelopesRepository: &r,
		budgetID:            budgetID,
	}
}

func (r *Envelopes) FindAll(ctx context.Context) ([]*models.Envelope, error) {
	result := make([]*models.Envelope, 0)
	err := r.find(ctx, doc{budgetID: r.budgetID}, func(d decodeFunc) error {
		e := &models.Envelope{}
		err := d(e)
		if err != nil {
			return err
		}
		result = append(result, e)
		return nil
	})
	return result, err
}

func (r *Envelopes) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Envelope, error) {
	result := &models.Envelope{}
	err := r.findByID(ctx, id, result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return result, err
}

func (r *Envelopes) ReplaceByID(ctx context.Context, id primitive.ObjectID, input models.EnvelopeInput) (*models.Envelope, error) {
	result := &models.Envelope{}
	replacement := input.ToModel(r.budgetID)
	err := r.replaceOne(ctx, doc{budgetID: r.budgetID}, replacement, result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return result, err
}

func (r *Envelopes) Insert(ctx context.Context, input models.EnvelopeInput) (*models.Envelope, error) {
	if err := r.expectBudget(ctx, r.budgetID); err != nil {
		return nil, err
	}

	envelope := input.ToModel(r.budgetID)
	result, err := r.collection.InsertOne(ctx, envelope)
	if err != nil {
		return nil, err
	}
	return envelope.WithID(result.InsertedID.(primitive.ObjectID)), nil
}
