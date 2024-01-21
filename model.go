// Copyright (c) Gabriel de Quadros Ligneul
// SPDX-License-Identifier: Apache-2.0 (see LICENSE)

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"time"
)

const DB_DIR = ".local/share/clockin"
const DB_FILE = "clockin.json"

const WORKING_HOURS = 8 * time.Hour

type WorkHours struct {
	ClockIn  time.Time `json:"clockIn"`
	ClockOut time.Time `json:"clockOut"`
}

type Model struct {
	workHours map[string]WorkHours
}

func NewModel() (*Model, error) {
	model := &Model{
		workHours: make(map[string]WorkHours),
	}
	err := model.load()
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *Model) Show(when time.Time) (WorkHours, bool) {
	workHours, ok := m.workHours[getKey(when)]
	return workHours, ok
}

func (m *Model) ClockIn(when time.Time) (WorkHours, error) {
	_, ok := m.workHours[getKey(when)]
	if ok {
		return WorkHours{}, errors.New("already clocked-in for the day")
	}
	workHours := WorkHours{
		ClockIn:  when,
		ClockOut: when.Add(WORKING_HOURS),
	}
	m.workHours[getKey(when)] = workHours
	err := m.store()
	if err != nil {
		return WorkHours{}, err
	}
	return workHours, nil
}

func (m *Model) load() error {
	// Get DB path
	dbDir, err := dbDir()
	if err != nil {
		return err
	}

	// Create DB dir if it doesn't exist
	if _, err := os.Stat(dbDir); err != nil {
		const fileMode = 0700
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dbDir, fileMode); err != nil {
				return fmt.Errorf("create DB dir: %w", err)
			}
		} else {
			return fmt.Errorf("read DB dir: %w", err)
		}
	}

	// Read DB file
	dbFile := path.Join(dbDir, DB_FILE)
	data, err := os.ReadFile(dbFile)
	if err != nil {
		if os.IsNotExist(err) {
			// It is fine if the file doesn't exist yet
			return nil
		} else {
			return fmt.Errorf("read DB file: %w", err)
		}
	}

	// Parse JSON
	if err := json.Unmarshal(data, &m.workHours); err != nil {
		return fmt.Errorf("parse DB file: %w", err)
	}
	return nil
}

func (m *Model) store() error {
	// Encode JSON
	data, err := json.MarshalIndent(m.workHours, "", "  ")
	if err != nil {
		panic(err)
	}

	// Get DB path
	dbDir, err := dbDir()
	if err != nil {
		return err
	}
	dbFile := path.Join(dbDir, DB_FILE)

	// Write DB file
	const fileMode = 0600
	if err := os.WriteFile(dbFile, data, fileMode); err != nil {
		return fmt.Errorf("write DB file: %w", err)
	}
	return nil
}

func dbDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("get home dir: %w", err)
	}
	return path.Join(homeDir, DB_DIR), nil
}

func getKey(when time.Time) string {
	return when.Format("2006-01-02")
}
