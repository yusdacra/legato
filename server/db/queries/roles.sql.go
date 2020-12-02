// Code generated by sqlc. DO NOT EDIT.
// source: roles.sql

package queries

import (
	"context"
	"database/sql"
	"encoding/json"
)

const addUserToRole = `-- name: AddUserToRole :exec
INSERT INTO Roles_Members (Guild_ID, Role_ID, Member_ID)
VALUES ($1, $2, $3) ON CONFLICT DO NOTHING
`

type AddUserToRoleParams struct {
	GuildID  uint64 `json:"guild_id"`
	RoleID   uint64 `json:"role_id"`
	MemberID uint64 `json:"member_id"`
}

func (q *Queries) AddUserToRole(ctx context.Context, arg AddUserToRoleParams) error {
	_, err := q.exec(ctx, q.addUserToRoleStmt, addUserToRole, arg.GuildID, arg.RoleID, arg.MemberID)
	return err
}

const createRole = `-- name: CreateRole :one
INSERT INTO Roles (
        Guild_ID,
        Role_ID,
        Name,
        Color,
        Hoist,
        Pingable,
        Position
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7
    ) RETURNING guild_id, role_id, name, color, hoist, pingable, position
`

type CreateRoleParams struct {
	GuildID  uint64 `json:"guild_id"`
	RoleID   uint64 `json:"role_id"`
	Name     string `json:"name"`
	Color    int32  `json:"color"`
	Hoist    bool   `json:"hoist"`
	Pingable bool   `json:"pingable"`
	Position string `json:"position"`
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error) {
	row := q.queryRow(ctx, q.createRoleStmt, createRole,
		arg.GuildID,
		arg.RoleID,
		arg.Name,
		arg.Color,
		arg.Hoist,
		arg.Pingable,
		arg.Position,
	)
	var i Role
	err := row.Scan(
		&i.GuildID,
		&i.RoleID,
		&i.Name,
		&i.Color,
		&i.Hoist,
		&i.Pingable,
		&i.Position,
	)
	return i, err
}

const deleteRole = `-- name: DeleteRole :exec
DELETE FROM Roles
WHERE Guild_ID = $1
    AND Role_ID = $2
`

type DeleteRoleParams struct {
	GuildID uint64 `json:"guild_id"`
	RoleID  uint64 `json:"role_id"`
}

func (q *Queries) DeleteRole(ctx context.Context, arg DeleteRoleParams) error {
	_, err := q.exec(ctx, q.deleteRoleStmt, deleteRole, arg.GuildID, arg.RoleID)
	return err
}

const getPermissions = `-- name: GetPermissions :one
SELECT Nodes
FROM Permissions
WHERE Guild_ID = $1
    AND Channel_ID = $2
    AND Role_ID = $3
`

type GetPermissionsParams struct {
	GuildID   uint64        `json:"guild_id"`
	ChannelID sql.NullInt64 `json:"channel_id"`
	RoleID    uint64        `json:"role_id"`
}

func (q *Queries) GetPermissions(ctx context.Context, arg GetPermissionsParams) (json.RawMessage, error) {
	row := q.queryRow(ctx, q.getPermissionsStmt, getPermissions, arg.GuildID, arg.ChannelID, arg.RoleID)
	var nodes json.RawMessage
	err := row.Scan(&nodes)
	return nodes, err
}

const getPermissionsWithoutChannel = `-- name: GetPermissionsWithoutChannel :one
SELECT Nodes
FROM Permissions
WHERE Guild_ID = $1
    AND Role_ID = $2
`

type GetPermissionsWithoutChannelParams struct {
	GuildID uint64 `json:"guild_id"`
	RoleID  uint64 `json:"role_id"`
}

func (q *Queries) GetPermissionsWithoutChannel(ctx context.Context, arg GetPermissionsWithoutChannelParams) (json.RawMessage, error) {
	row := q.queryRow(ctx, q.getPermissionsWithoutChannelStmt, getPermissionsWithoutChannel, arg.GuildID, arg.RoleID)
	var nodes json.RawMessage
	err := row.Scan(&nodes)
	return nodes, err
}

const getRolePosition = `-- name: GetRolePosition :one
SELECT Position
FROM Roles
WHERE Role_ID = $1
    AND Guild_ID = $2
`

type GetRolePositionParams struct {
	RoleID  uint64 `json:"role_id"`
	GuildID uint64 `json:"guild_id"`
}

func (q *Queries) GetRolePosition(ctx context.Context, arg GetRolePositionParams) (string, error) {
	row := q.queryRow(ctx, q.getRolePositionStmt, getRolePosition, arg.RoleID, arg.GuildID)
	var position string
	err := row.Scan(&position)
	return position, err
}

const getRolesForGuild = `-- name: GetRolesForGuild :many
SELECT guild_id, role_id, name, color, hoist, pingable, position
FROM Roles
WHERE Guild_ID = $1
ORDER BY Position
`

func (q *Queries) GetRolesForGuild(ctx context.Context, guildID uint64) ([]Role, error) {
	rows, err := q.query(ctx, q.getRolesForGuildStmt, getRolesForGuild, guildID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.GuildID,
			&i.RoleID,
			&i.Name,
			&i.Color,
			&i.Hoist,
			&i.Pingable,
			&i.Position,
		); err != nil {
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

const moveRole = `-- name: MoveRole :exec
UPDATE Roles
SET Position = $1
WHERE Role_ID = $2
    AND Guild_ID = $3
`

type MoveRoleParams struct {
	Position string `json:"position"`
	RoleID   uint64 `json:"role_id"`
	GuildID  uint64 `json:"guild_id"`
}

func (q *Queries) MoveRole(ctx context.Context, arg MoveRoleParams) error {
	_, err := q.exec(ctx, q.moveRoleStmt, moveRole, arg.Position, arg.RoleID, arg.GuildID)
	return err
}

const removeUserFromRole = `-- name: RemoveUserFromRole :exec
DELETE FROM Roles_Members
WHERE Guild_ID = $1
    AND Role_ID = $2
    AND Member_ID = $3
`

type RemoveUserFromRoleParams struct {
	GuildID  uint64 `json:"guild_id"`
	RoleID   uint64 `json:"role_id"`
	MemberID uint64 `json:"member_id"`
}

func (q *Queries) RemoveUserFromRole(ctx context.Context, arg RemoveUserFromRoleParams) error {
	_, err := q.exec(ctx, q.removeUserFromRoleStmt, removeUserFromRole, arg.GuildID, arg.RoleID, arg.MemberID)
	return err
}

const rolesForUser = `-- name: RolesForUser :many
SELECT Role_ID
FROM Roles_Members
WHERE Guild_ID = $1
    AND Member_ID = $2
`

type RolesForUserParams struct {
	GuildID  uint64 `json:"guild_id"`
	MemberID uint64 `json:"member_id"`
}

func (q *Queries) RolesForUser(ctx context.Context, arg RolesForUserParams) ([]uint64, error) {
	rows, err := q.query(ctx, q.rolesForUserStmt, rolesForUser, arg.GuildID, arg.MemberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []uint64
	for rows.Next() {
		var role_id uint64
		if err := rows.Scan(&role_id); err != nil {
			return nil, err
		}
		items = append(items, role_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setPermissions = `-- name: SetPermissions :exec
INSERT INTO Permissions (Guild_ID, Channel_ID, Role_ID, Nodes)
VALUES ($1, $2, $3, $4) ON CONFLICT (Guild_ID, Channel_ID, Role_ID) DO
UPDATE
SET Nodes = EXCLUDED.Nodes
`

type SetPermissionsParams struct {
	GuildID   uint64          `json:"guild_id"`
	ChannelID sql.NullInt64   `json:"channel_id"`
	RoleID    uint64          `json:"role_id"`
	Nodes     json.RawMessage `json:"nodes"`
}

func (q *Queries) SetPermissions(ctx context.Context, arg SetPermissionsParams) error {
	_, err := q.exec(ctx, q.setPermissionsStmt, setPermissions,
		arg.GuildID,
		arg.ChannelID,
		arg.RoleID,
		arg.Nodes,
	)
	return err
}

const setRoleColor = `-- name: SetRoleColor :exec
UPDATE Roles
    SET Color = $1
    WHERE Guild_ID = $2
      AND Role_ID = $3
`

type SetRoleColorParams struct {
	Color   int32  `json:"color"`
	GuildID uint64 `json:"guild_id"`
	RoleID  uint64 `json:"role_id"`
}

func (q *Queries) SetRoleColor(ctx context.Context, arg SetRoleColorParams) error {
	_, err := q.exec(ctx, q.setRoleColorStmt, setRoleColor, arg.Color, arg.GuildID, arg.RoleID)
	return err
}

const setRoleHoist = `-- name: SetRoleHoist :exec
UPDATE Roles
    SET Hoist = $1
    WHERE Guild_ID = $2
      AND Role_ID = $3
`

type SetRoleHoistParams struct {
	Hoist   bool   `json:"hoist"`
	GuildID uint64 `json:"guild_id"`
	RoleID  uint64 `json:"role_id"`
}

func (q *Queries) SetRoleHoist(ctx context.Context, arg SetRoleHoistParams) error {
	_, err := q.exec(ctx, q.setRoleHoistStmt, setRoleHoist, arg.Hoist, arg.GuildID, arg.RoleID)
	return err
}

const setRoleName = `-- name: SetRoleName :exec
UPDATE Roles
    SET Name = $1
    WHERE Guild_ID = $2
      AND Role_ID = $3
`

type SetRoleNameParams struct {
	Name    string `json:"name"`
	GuildID uint64 `json:"guild_id"`
	RoleID  uint64 `json:"role_id"`
}

func (q *Queries) SetRoleName(ctx context.Context, arg SetRoleNameParams) error {
	_, err := q.exec(ctx, q.setRoleNameStmt, setRoleName, arg.Name, arg.GuildID, arg.RoleID)
	return err
}

const setRolePingable = `-- name: SetRolePingable :exec
UPDATE Roles
    SET Pingable = $1
    WHERE Guild_ID = $2
      AND Role_ID = $3
`

type SetRolePingableParams struct {
	Pingable bool   `json:"pingable"`
	GuildID  uint64 `json:"guild_id"`
	RoleID   uint64 `json:"role_id"`
}

func (q *Queries) SetRolePingable(ctx context.Context, arg SetRolePingableParams) error {
	_, err := q.exec(ctx, q.setRolePingableStmt, setRolePingable, arg.Pingable, arg.GuildID, arg.RoleID)
	return err
}
