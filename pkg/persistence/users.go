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

// Users defines the users interface to the store
type Users interface {
	// Delete removes a user from the store
	Delete(context.Context, *model.User) (*model.User, error)
	// Exists check if the user exists
	Exists(context.Context, string) (bool, error)
	// Get returns a user from the store
	Get(context.Context, string) (*model.User, error)
	// List returns a list of users from the store
	List(context.Context, ...ListFunc) ([]*model.User, error)
	// Size returns the number of users
	Size(context.Context) (int64, error)
	// Update updates a user in the store
	Update(context.Context, *model.User) error
	// transaction set the transaction
	transaction(*gorm.DB) Users
}

// userImpl handles access to the users model
type userImpl struct {
	Interface

	conn *gorm.DB
}

// Delete removes a user from the store
func (u userImpl) Delete(ctx context.Context, user *model.User) (*model.User, error) {
	timed := prometheus.NewTimer(deleteLatency)
	defer timed.ObserveDuration()

	if user.ID == 0 && user.Username == "" {
		return nil, errors.New("invalid user for deletion: must have id or username")
	}

	q := u.conn
	if user.ID != 0 {
		q = q.Where("id = ?", user.ID)
	} else if user.Username != "" {
		q = q.Where("username = ?", user.Username)
	}

	return user, q.Delete(&model.User{}).Error
}

// Exists check if the user exists
func (u userImpl) Exists(ctx context.Context, name string) (bool, error) {
	if _, err := u.Get(ctx, name); err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return false, err
		}

		return false, nil
	}

	return true, nil
}

// Get returns a user from the store
func (u userImpl) Get(ctx context.Context, name string) (*model.User, error) {
	timed := prometheus.NewTimer(getLatency)
	defer timed.ObserveDuration()

	user := &model.User{}

	err := u.conn.Where("username = ?", name).Find(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// List returns a list of users from the store
func (u userImpl) List(ctx context.Context, opts ...ListFunc) ([]*model.User, error) {
	timed := prometheus.NewTimer(listLatency)
	defer timed.ObserveDuration()

	terms := ApplyListOptions(opts...)

	q := u.conn
	if terms.HasDisabled() {
		q = q.Where("disabled = ?", terms.GetDisabled())
	}
	if terms.HasUser() {
		q = q.Where("username = ?", terms.GetUser())
	}
	if terms.HasID() {
		q = q.Where("id = ?", terms.GetID())
	}

	var list []*model.User

	return list, q.Find(&list).Error
}

func (u userImpl) Size(ctx context.Context) (int64, error) {
	timed := prometheus.NewTimer(listLatency)
	defer timed.ObserveDuration()

	var count int64

	return count, u.conn.Model(&model.User{}).Count(&count).Error
}

// Update updates a user in the store
func (u userImpl) Update(ctx context.Context, user *model.User) error {
	timed := prometheus.NewTimer(setLatency)
	defer timed.ObserveDuration()

	return u.conn.Model(&model.Team{}).
		Where("username = ?", user.Username).
		Assign(&model.User{
			Disabled: user.Disabled,
			Email:    user.Email,
			Username: user.Username,
		}).
		FirstOrCreate(user).
		Error
}

func (u *userImpl) transaction(db *gorm.DB) Users {
	u.conn = db

	return u
}
