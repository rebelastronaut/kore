/**
 * Copyright 2020 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package persistence

import (
	"context"
	"errors"

	"github.com/appvia/kore/pkg/persistence/model"

	"github.com/jinzhu/gorm"
	"github.com/prometheus/client_golang/prometheus"
)

// Members is the team members interface
type Members interface {
	// AddUser is responsible for adding a user to a team
	AddUser(context.Context, string, string, []string) error
	// Delete is responsible for removing a member from the team
	Delete(context.Context, *model.Member) error
	// DeleteBy removes based on a filter
	DeleteBy(context.Context, ...ListFunc) error
	// ListMembers returns a list of members in a team
	List(context.Context, ...ListFunc) ([]*model.Member, error)
	// Preload adds to the query preload
	Preload(...string) Members
	// Add is responsible for adding a member to a team
	Update(context.Context, *model.Member) error
}

// membersImpl implements the above interface
type membersImpl struct {
	Interface
	// load is for preloading
	load []string
	// conn is the db connection
	conn *gorm.DB
}

// AddUser is responsible for adding a user to a team
func (m *membersImpl) AddUser(ctx context.Context, user, team string, roles []string) error {
	timed := prometheus.NewTimer(setLatency)
	defer timed.ObserveDuration()

	u, err := m.Users().Get(ctx, user)
	if err != nil {
		return err
	}
	t, err := m.Teams().Get(ctx, team)
	if err != nil {
		return err
	}

	return m.Members().Update(ctx, &model.Member{
		UserID: u.ID,
		TeamID: t.ID,
		Roles:  roles,
	})
}

// List returns a list of teams for a specific user
func (m *membersImpl) List(ctx context.Context, opts ...ListFunc) ([]*model.Member, error) {
	timed := prometheus.NewTimer(listLatency)
	defer timed.ObserveDuration()

	terms := ApplyListOptions(opts...)

	var list []*model.Member

	q := Preload(m.load, m.conn).
		Select("m.*").
		Table("members m").
		Joins("LEFT JOIN teams t ON t.id = m.team_id").
		Joins("LEFT JOIN users u ON u.id = m.user_id")

	if terms.HasTeam() {
		q = q.Where("t.name = ?", terms.GetTeam())
	}
	if terms.HasUser() {
		q = q.Where("u.username = ?", terms.GetUser())
	}
	if err := q.Preload("Team").Find(&list).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return nil, err
		}

		return []*model.Member{}, nil
	}

	return list, nil
}

// DeleteBy is responsible for deleting by filter
func (m *membersImpl) DeleteBy(ctx context.Context, filters ...ListFunc) error {
	timed := prometheus.NewTimer(deleteLatency)
	defer timed.ObserveDuration()

	if len(filters) <= 0 {
		return errors.New("no filters for delete by on users")
	}

	terms := ApplyListOptions(filters...)

	q := m.conn.
		Model(&model.Member{}).
		Select("m.*").
		Table("members m").
		Joins("JOIN teams t ON t.id = m.team_id").
		Joins("JOIN users u ON u.id = m.user_id")

	if terms.HasUser() {
		q = q.Where("u.username = ?", terms.GetUser())
	}
	if terms.HasTeam() {
		q = q.Where("t.name = ?", terms.GetTeam())
	}

	list := []*model.Member{}
	if err := q.Find(&list).Error; err != nil {
		return err
	}

	for _, x := range list {
		if err := m.conn.Model(&model.Member{}).Delete(x).Error; err != nil {
			return err
		}
	}

	return nil
}

// Delete is responsible for removing a member from the team
func (m *membersImpl) Delete(ctx context.Context, member *model.Member) error {
	timed := prometheus.NewTimer(deleteLatency)
	defer timed.ObserveDuration()

	if member.UserID == 0 {
		return errors.New("no user id defined")
	}
	if member.TeamID == 0 {
		return errors.New("no team id defined")
	}

	return m.conn.
		Where("user_id = ?", member.UserID).
		Where("team_id = ?", member.TeamID).
		Delete(member).
		Error
}

// Add is responsible for adding a member to a team
func (m membersImpl) Update(ctx context.Context, member *model.Member) error {
	timed := prometheus.NewTimer(setLatency)
	defer timed.ObserveDuration()

	return m.conn.FirstOrCreate(member, member).Error
}

// Preload adds proloading to the queries
func (m *membersImpl) Preload(v ...string) Members {
	m.load = append(m.load, v...)

	return m
}
