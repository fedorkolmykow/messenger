\connect postgres

CREATE DATABASE avitomes;

\connect avitomes


CREATE TABLE Users (
	user_id serial NOT NULL,
	username varchar(255) NOT NULL UNIQUE,
	created_at timestamp NOT NULL,
	CONSTRAINT Users_pk PRIMARY KEY (user_id)
) WITH (
  OIDS=FALSE
);



CREATE TABLE Messages (
	message_id serial NOT NULL,
	author_id integer NOT NULL,
	chat_id integer NOT NULL,
	text VARCHAR(255) NOT NULL,
	created_at timestamp NOT NULL,
	CONSTRAINT Messages_pk PRIMARY KEY (message_id)
) WITH (
  OIDS=FALSE
);



CREATE TABLE Chat_Users (
	chat_users_id serial NOT NULL,
	chat_id integer NOT NULL,
	user_id integer NOT NULL,
	CONSTRAINT Chat_Users_pk PRIMARY KEY (chat_users_id)
) WITH (
  OIDS=FALSE
);



CREATE TABLE Chats (
	chat_id serial NOT NULL,
	name varchar(255) NOT NULL UNIQUE,
	created_at timestamp NOT NULL,
	CONSTRAINT Chats_pk PRIMARY KEY (chat_id)
) WITH (
  OIDS=FALSE
);




ALTER TABLE Messages ADD CONSTRAINT Messages_fk0 FOREIGN KEY (author_id) REFERENCES Users(user_id);
ALTER TABLE Messages ADD CONSTRAINT Messages_fk1 FOREIGN KEY (chat_id) REFERENCES Chats(chat_id);

ALTER TABLE Chat_Users ADD CONSTRAINT Chat_Users_fk0 FOREIGN KEY (chat_id) REFERENCES Chats(chat_id);
ALTER TABLE Chat_Users ADD CONSTRAINT Chat_Users_fk1 FOREIGN KEY (user_id) REFERENCES Users(user_id);





INSERT INTO Users (user_id, username, created_at)
	VALUES	
	(0, 'Mark Green', current_timestamp),
	(1, 'Jack Blue', current_timestamp),
	(2, 'Mary Bloody', current_timestamp),
	(3, 'Daria Purple', current_timestamp);
	
INSERT INTO Chats (chat_id, name, created_at)
	VALUES	
	(0, 'Rats', current_timestamp),
	(1, 'Students', current_timestamp),
	(2, 'Teachers', current_timestamp),
	(3, 'World government', current_timestamp);
	
INSERT INTO Chat_Users (chat_users_id, chat_id, user_id)
	VALUES	
	(0, 0, 0),
	(1, 0, 2),
	(2, 1, 1),
	(3, 1, 2),
	(4, 2, 0),
	(5, 2, 3),
	(6, 3, 0),
	(7, 3, 2),
	(8, 3, 3);
	
INSERT INTO Messages (message_id, author_id, chat_id, text, created_at)
	VALUES	
	(0, 2, 1, 'OMG, LOOK ON THIS!', current_timestamp),
	(1, 0, 3, 'Is it a bird?', current_timestamp),
	(2, 2, 3, 'Is it a plane?', current_timestamp),
	(3, 3, 3, 'No, it is Superman!', current_timestamp);

select setval(' id_sequence', (select max( id ) + 1 from  Messages));
select setval(' id_sequence', (select max( id ) + 1 from  Chat_Users));
select setval(' id_sequence', (select max( id ) + 1 from  Chats));
select setval(' id_sequence', (select max( id ) + 1 from  Users));