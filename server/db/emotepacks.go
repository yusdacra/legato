package db

import (
	"github.com/harmony-development/legato/server/db/queries"
	"github.com/ztrue/tracerr"
)

func (db HarmonyDB) CreateEmotePack(userID, packID uint64, packName string) error {
	err := db.queries.CreateEmotePack(ctx, queries.CreateEmotePackParams{
		UserID:   userID,
		PackID:   packID,
		PackName: packName,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return err
}

func (db HarmonyDB) IsPackOwner(userID, packID uint64) (bool, error) {
	owner, err := db.queries.GetPackOwner(ctx, packID)
	err = tracerr.Wrap(err)
	if err != nil {
		return false, err
	}
	return owner == userID, nil
}

func (db HarmonyDB) AddEmoteToPack(packID uint64, imageID string, name string) error {
	err := db.queries.AddEmoteToPack(ctx, queries.AddEmoteToPackParams{
		PackID:    packID,
		ImageID:   imageID,
		EmoteName: name,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return err
}

func (db HarmonyDB) DeleteEmoteFromPack(packID uint64, imageID string) error {
	err := db.queries.DeleteEmoteFromPack(ctx, queries.DeleteEmoteFromPackParams{
		PackID:  packID,
		ImageID: imageID,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return err
}

func (db HarmonyDB) DeleteEmotePack(packID uint64) error {
	err := db.queries.DeleteEmotePack(ctx, queries.DeleteEmotePackParams{
		PackID: packID,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return err
}

func (db HarmonyDB) GetEmotePacks(userID uint64) ([]queries.GetEmotePacksRow, error) {
	emotes, err := db.queries.GetEmotePacks(ctx, userID)
	err = tracerr.Wrap(err)
	if err != nil {
		db.Logger.CheckException(err)
		return nil, err
	}
	return emotes, nil
}

func (db HarmonyDB) GetEmotePackEmotes(packID uint64) ([]queries.GetEmotePackEmotesRow, error) {
	data, err := db.queries.GetEmotePackEmotes(ctx, packID)
	err = tracerr.Wrap(err)
	return data, err
}

func (db HarmonyDB) DequipEmotePack(userID, packID uint64) error {
	return tracerr.Wrap(db.queries.DequipEmotePack(ctx, queries.DequipEmotePackParams{
		PackID: packID,
		UserID: userID,
	}))
}
