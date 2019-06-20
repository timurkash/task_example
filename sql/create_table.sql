CREATE TABLE task_example.tasks (
	guid uuid NOT NULL,
	status varchar(10) NOT NULL,
	cdate timestamp NOT NULL DEFAULT now(),
	udate timestamp NOT NULL DEFAULT now(),
	CONSTRAINT tasks_pk PRIMARY KEY (guid)
);
