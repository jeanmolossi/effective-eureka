INSERT INTO courses (
	course_id,
	course_title,
	course_description,
	course_thumb,
	course_published,
	course_created_at,
	course_updated_at
) VALUES (
	DEFAULT,
	'Curso de Node.js',
	'Curso de Node.js',
	'https://www.opus-software.com.br/wp-content/uploads/2018/09/nodejs.jpg',
	TRUE,
	DEFAULT,
	DEFAULT
), (
	DEFAULT,
	'Curso de React.js',
	'Curso de React.js',
	'https://miro.medium.com/max/1838/0*1V_xALlt1BCKvFBW.jpeg',
	TRUE,
	DEFAULT,
	DEFAULT
);

INSERT INTO modules (
	module_id,
	module_title,
	module_description,
	module_thumb,
	module_published,
	module_created_at,
	module_updated_at,
	course_id
) VALUES (
	DEFAULT,
	'Introdução',
	'Introdução',
	'https://www.opus-software.com.br/wp-content/uploads/2018/09/nodejs.jpg',
	TRUE,
	DEFAULT,
	DEFAULT,
	(SELECT course_id FROM courses WHERE course_title = 'Curso de Node.js')
), (
	DEFAULT,
	'Introdução',
	'Introdução',
	'https://miro.medium.com/max/1838/0*1V_xALlt1BCKvFBW.jpeg',
	TRUE,
	DEFAULT,
	DEFAULT,
	(SELECT course_id FROM courses WHERE course_title = 'Curso de React.js')
);

INSERT INTO sections (
	section_id,
	section_index,
	section_title,
	section_published,
	section_created_at,
	section_updated_at,
	module_id,
	course_id
) VALUES (
	DEFAULT,
	1,
	'Introdução',
	TRUE,
	DEFAULT,
	DEFAULT,
	(SELECT module_id FROM modules WHERE module_title = 'Introdução' LIMIT 1),
	(SELECT course_id FROM courses WHERE course_title = 'Curso de Node.js')
), (
	DEFAULT,
	1,
	'Introdução',
	TRUE,
	DEFAULT,
	DEFAULT,
	(SELECT module_id FROM modules WHERE module_title = 'Introdução' LIMIT 1 OFFSET 1),
	(SELECT course_id FROM courses WHERE course_title = 'Curso de React.js')
);

INSERT INTO lessons (
	lesson_id,
	lesson_index,
	lesson_title,
	lesson_description,
	lesson_thumb,
	lesson_published,
	lesson_created_at,
	lesson_updated_at,
	section_id
) VALUES (
	DEFAULT,
	1,
	'Introdução',
	'Introdução',
	'https://www.opus-software.com.br/wp-content/uploads/2018/09/nodejs.jpg',
	TRUE,
	DEFAULT,
	DEFAULT,
	(SELECT section_id FROM sections WHERE section_title = 'Introdução' LIMIT 1)
), (
	DEFAULT,
	1,
	'Introdução',
	'Introdução',
	'https://miro.medium.com/max/1838/0*1V_xALlt1BCKvFBW.jpeg',
	TRUE,
	DEFAULT,
	DEFAULT,
	(SELECT section_id FROM sections WHERE section_title = 'Introdução' LIMIT 1 OFFSET 1)
);
