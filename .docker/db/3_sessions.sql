
CREATE TABLE sess (
	sess_id VARCHAR(36) PRIMARY KEY DEFAULT uuid_generate_v4(),
	sess_expiration TIMESTAMP NOT NULL DEFAULT NOW() + INTERVAL '10 minutes',
	sess_access_token TEXT NOT NULL,
	sess_refresh_token TEXT NOT NULL,

	sess_student_id VARCHAR(36) UNIQUE REFERENCES students(student_id) ON DELETE CASCADE
);

CREATE INDEX sess_idx ON sess (sess_id);
CREATE INDEX sess_studentx ON sess (sess_student_id);
