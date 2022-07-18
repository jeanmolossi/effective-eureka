CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE students (
	student_id VARCHAR(36) PRIMARY KEY DEFAULT uuid_generate_v4(),
	student_email VARCHAR(255) NOT NULL UNIQUE,
	student_password VARCHAR(255) NOT NULL,

	student_created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	student_updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX student_idx ON students (student_id);
CREATE INDEX student_emailx ON students (student_email);
