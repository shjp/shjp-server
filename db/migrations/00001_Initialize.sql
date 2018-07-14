-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE groups (
  id UUID PRIMARY KEY NOT NULL,
  name VARCHAR(60) NOT NULL,
  description TEXT
);

CREATE TABLE roles (
  id UUID PRIMARY KEY NOT NULL,
  group_id UUID NOT NULL REFERENCES groups (id),
  name TEXT NOT NULL,
  privilege INT NOT NULL
);

CREATE TABLE members (
  id UUID PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  baptismal_name TEXT,
  birthday TIMESTAMP,
  feast_day TIMESTAMP,
  created TIMESTAMP NOT NULL DEFAULT now(),
  last_active TIMESTAMP,
  account_type TEXT,
  account_hash TEXT
);

CREATE TABLE events (
  id UUID PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  date TIMESTAMP,
  length INT,
  creator UUID NOT NULL REFERENCES members (id),
  deadline TIMESTAMP,
  allow_maybe BOOLEAN NOT NULL,
  description TEXT,
  location POINT,
  location_description TEXT
);

CREATE TABLE announcements (
  id UUID PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  created TIMESTAMP NOT NULL DEFAULT now(),
  updated TIMESTAMP,
  author_id UUID NOT NULL REFERENCES members (id),
  content TEXT
);

CREATE TABLE comments (
  id UUID PRIMARY KEY NOT NULL,
  author_id UUID NOT NULL REFERENCES members (id),
  created TIMESTAMP DEFAULT now(),
  updated TIMESTAMP,
  parent_id UUID NOT NULL,
  parent_type VARCHAR(20) NOT NULL
    CHECK(parent_type IN ('announcement', 'event')),
  content TEXT
);

CREATE TABLE groups_members (
  member_id UUID NOT NULL REFERENCES members (id),
  group_id UUID NOT NULL REFERENCES groups (id),
  role_id UUID REFERENCES roles (id),
  status VARCHAR(20) NOT NULL
    CHECK(status IN ('accepted', 'pending'))
);

CREATE UNIQUE INDEX group_member_index ON groups_members (member_id, group_id);

CREATE TABLE groups_events (
  group_id UUID NOT NULL REFERENCES groups (id),
  event_id UUID NOT NULL REFERENCES events (id)
);

CREATE TABLE groups_announcements (
  group_id UUID NOT NULL REFERENCES groups (id),
  announcement_id UUID NOT NULL REFERENCES announcements (id)
);

CREATE TABLE members_events (
  member_id UUID NOT NULL REFERENCES members (id),
  event_id UUID NOT NULL REFERENCES events (id),
  rsvp VARCHAR(10) NOT NULL
    CHECK(rsvp IN ('yes', 'no', 'maybe', 'unanswered'))
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE members_events;
DROP TABLE groups_announcements;
DROP TABLE groups_events;
DROP TABLE groups_members;
DROP TABLE comments;
DROP TABLE events;
DROP TABLE announcements;
DROP TABLE members;
DROP TABLE roles;
DROP TABLE groups;