// Code generated by sqlc. DO NOT EDIT.
// source: emotes.sql

package queries

import (
	"context"
)

const acquireEmotePack = `-- name: AcquireEmotePack :exec
INSERT INTO Acquired_Emote_Packs (Pack_ID, User_ID)
VALUES ($1, $2)
`

type AcquireEmotePackParams struct {
	PackID uint64 `json:"pack_id"`
	UserID uint64 `json:"user_id"`
}

func (q *Queries) AcquireEmotePack(ctx context.Context, arg AcquireEmotePackParams) error {
	_, err := q.exec(ctx, q.acquireEmotePackStmt, acquireEmotePack, arg.PackID, arg.UserID)
	return err
}

const addEmoteToPack = `-- name: AddEmoteToPack :exec
INSERT INTO Emote_Pack_Emotes (Pack_ID, Image_ID, Emote_Name)
VALUES ($1, $2, $3)
`

type AddEmoteToPackParams struct {
	PackID    uint64 `json:"pack_id"`
	ImageID   string `json:"image_id"`
	EmoteName string `json:"emote_name"`
}

func (q *Queries) AddEmoteToPack(ctx context.Context, arg AddEmoteToPackParams) error {
	_, err := q.exec(ctx, q.addEmoteToPackStmt, addEmoteToPack, arg.PackID, arg.ImageID, arg.EmoteName)
	return err
}

const createEmotePack = `-- name: CreateEmotePack :exec
INSERT INTO Emote_Packs (Pack_ID, Pack_Name, User_ID)
VALUES ($1, $2, $3)
`

type CreateEmotePackParams struct {
	PackID   uint64 `json:"pack_id"`
	PackName string `json:"pack_name"`
	UserID   uint64 `json:"user_id"`
}

func (q *Queries) CreateEmotePack(ctx context.Context, arg CreateEmotePackParams) error {
	_, err := q.exec(ctx, q.createEmotePackStmt, createEmotePack, arg.PackID, arg.PackName, arg.UserID)
	return err
}

const deleteEmoteFromPack = `-- name: DeleteEmoteFromPack :exec
DELETE FROM Emote_Pack_Emotes
WHERE Pack_ID = $1
	AND Image_ID = $2
`

type DeleteEmoteFromPackParams struct {
	PackID  uint64 `json:"pack_id"`
	ImageID string `json:"image_id"`
}

func (q *Queries) DeleteEmoteFromPack(ctx context.Context, arg DeleteEmoteFromPackParams) error {
	_, err := q.exec(ctx, q.deleteEmoteFromPackStmt, deleteEmoteFromPack, arg.PackID, arg.ImageID)
	return err
}

const deleteEmotePack = `-- name: DeleteEmotePack :exec
DELETE FROM Emote_Packs
WHERE Pack_ID = $1
	AND User_ID = $2
`

type DeleteEmotePackParams struct {
	PackID uint64 `json:"pack_id"`
	UserID uint64 `json:"user_id"`
}

func (q *Queries) DeleteEmotePack(ctx context.Context, arg DeleteEmotePackParams) error {
	_, err := q.exec(ctx, q.deleteEmotePackStmt, deleteEmotePack, arg.PackID, arg.UserID)
	return err
}

const dequipEmotePack = `-- name: DequipEmotePack :exec
DELETE FROM Acquired_Emote_Packs
WHERE Pack_ID = $1
	AND User_ID = $2
`

type DequipEmotePackParams struct {
	PackID uint64 `json:"pack_id"`
	UserID uint64 `json:"user_id"`
}

func (q *Queries) DequipEmotePack(ctx context.Context, arg DequipEmotePackParams) error {
	_, err := q.exec(ctx, q.dequipEmotePackStmt, dequipEmotePack, arg.PackID, arg.UserID)
	return err
}

const getEmotePackEmotes = `-- name: GetEmotePackEmotes :many
SELECT Image_ID,
	Emote_Name
FROM Emote_Pack_Emotes
WHERE Pack_ID = $1
`

type GetEmotePackEmotesRow struct {
	ImageID   string `json:"image_id"`
	EmoteName string `json:"emote_name"`
}

func (q *Queries) GetEmotePackEmotes(ctx context.Context, packID uint64) ([]GetEmotePackEmotesRow, error) {
	rows, err := q.query(ctx, q.getEmotePackEmotesStmt, getEmotePackEmotes, packID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetEmotePackEmotesRow
	for rows.Next() {
		var i GetEmotePackEmotesRow
		if err := rows.Scan(&i.ImageID, &i.EmoteName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEmotePacks = `-- name: GetEmotePacks :many
SELECT Emote_Packs.Pack_ID,
	Emote_Packs.User_ID,
	Emote_Packs.Pack_Name
FROM Emote_Packs
	INNER JOIN Acquired_Emote_Packs ON Acquired_Emote_Packs.Pack_ID = Emote_Packs.Pack_ID
WHERE Acquired_Emote_Packs.User_ID = $1
`

type GetEmotePacksRow struct {
	PackID   uint64 `json:"pack_id"`
	UserID   uint64 `json:"user_id"`
	PackName string `json:"pack_name"`
}

func (q *Queries) GetEmotePacks(ctx context.Context, userID uint64) ([]GetEmotePacksRow, error) {
	rows, err := q.query(ctx, q.getEmotePacksStmt, getEmotePacks, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetEmotePacksRow
	for rows.Next() {
		var i GetEmotePacksRow
		if err := rows.Scan(&i.PackID, &i.UserID, &i.PackName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPackOwner = `-- name: GetPackOwner :one
SELECT Pack_ID
FROM Emote_Packs
WHERE Pack_ID = $1
`

func (q *Queries) GetPackOwner(ctx context.Context, packID uint64) (uint64, error) {
	row := q.queryRow(ctx, q.getPackOwnerStmt, getPackOwner, packID)
	var pack_id uint64
	err := row.Scan(&pack_id)
	return pack_id, err
}
