CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE courses (
	course_id VARCHAR(36) PRIMARY KEY DEFAULT uuid_generate_v4(),
	course_title VARCHAR(255) NOT NULL,
	course_description TEXT NOT NULL DEFAULT 'Sem descrição',
	course_thumb TEXT NULL,
	course_published BOOLEAN NOT NULL DEFAULT FALSE,
	course_created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	course_updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX course_idx ON courses (course_id);
CREATE INDEX course_publishedx ON courses (course_published);

CREATE TABLE modules (
	module_id VARCHAR(36) PRIMARY KEY DEFAULT uuid_generate_v4(),
	module_title VARCHAR(255) NOT NULL,
	module_description TEXT NOT NULL DEFAULT 'Sem descrição',
	module_thumb TEXT NULL,
	module_published BOOLEAN NOT NULL DEFAULT FALSE,
	module_created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	module_updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

	course_id VARCHAR(36) REFERENCES courses(course_id) ON DELETE CASCADE
);

CREATE INDEX module_idx ON modules (module_id);
CREATE INDEX module_publishedx ON modules (module_published);
CREATE INDEX module_course_idx ON modules (course_id);

CREATE TABLE sections (
	section_id VARCHAR(36) PRIMARY KEY DEFAULT uuid_generate_v4(),
	section_index INTEGER NOT NULL DEFAULT 999,
	section_title VARCHAR(255) NOT NULL,
	section_published BOOLEAN NOT NULL DEFAULT FALSE,
	section_created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	section_updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

	module_id VARCHAR(36) REFERENCES modules(module_id) ON DELETE CASCADE,
	course_id VARCHAR(36) REFERENCES courses(course_id)
);

CREATE INDEX section_idx ON sections (section_id);
CREATE INDEX section_publishedx ON sections (section_published);
CREATE INDEX section_module_idx ON sections (module_id);
CREATE INDEX section_course_idx ON sections (course_id);

CREATE TABLE lessons (
	lesson_id VARCHAR(36) PRIMARY KEY DEFAULT uuid_generate_v4(),
	lesson_index INTEGER NOT NULL DEFAULT 999,
	lesson_title VARCHAR(255) NOT NULL,
	lesson_description TEXT NULL,
	lesson_thumb TEXT NULL,
	lesson_video TEXT NULL,
	lesson_video_preview TEXT NULL,
	lesson_published BOOLEAN NOT NULL DEFAULT FALSE,
	lesson_created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	lesson_updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

	section_id VARCHAR(36) REFERENCES sections(section_id) ON DELETE CASCADE
);

CREATE INDEX lesson_idx ON lessons (lesson_id);
CREATE INDEX lesson_publishedx ON lessons (lesson_published);
CREATE INDEX lesson_section_idx ON lessons (section_id);

CREATE TABLE attachments (
	attachment_id VARCHAR(36) PRIMARY KEY DEFAULT uuid_generate_v4(),
	attachment_name VARCHAR(255) NOT NULL,
	attachment_type VARCHAR(255) NOT NULL,
	attachment_published BOOLEAN NOT NULL DEFAULT FALSE,
	attachment_created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	attachment_updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

	lesson_id VARCHAR(36) REFERENCES lessons(lesson_id) ON DELETE CASCADE
);

CREATE INDEX attachment_idx ON attachments (attachment_id);
CREATE INDEX attachment_publishedx ON attachments (attachment_published);
CREATE INDEX attachment_lesson_idx ON attachments (lesson_id);
