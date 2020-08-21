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
	(1, 'Mark Green', current_timestamp),
	(2, 'Jack Blue', current_timestamp),
	(3, 'Mary Bloody', current_timestamp),
	(4, 'Daria Purple', current_timestamp);
	
INSERT INTO Chats (chat_id, name, created_at)
	VALUES	
	(1, 'Rats', current_timestamp),
	(2, 'Students', current_timestamp),
	(3, 'Teachers', current_timestamp),
	(4, 'World government', current_timestamp);
	
INSERT INTO Chat_Users (chat_users_id, chat_id, user_id)
	VALUES	
	(1, 1, 1),
	(2, 1, 3),
	(3, 2, 2),
	(4, 2, 3),
	(5, 3, 1),
	(6, 3, 4),
	(7, 4, 1),
	(8, 4, 3),
	(9, 4, 4);
	
INSERT INTO Messages (message_id, author_id, chat_id, text, created_at)
	VALUES	
	(1, 3, 2, 'OMG, LOOK ON THIS!', current_timestamp),
	(2, 1, 4, 'Is it a bird?', current_timestamp),
	(3, 3, 4, 'Is it a plane?', current_timestamp),
	(4, 4, 4, 'No, it is Superman!', current_timestamp);

 select setval('messages_message_id_seq', (select max(message_id)  from Messages));
 select setval('chats_chat_id_seq', (select max(chat_id) from Chats));
 select setval('chat_users_chat_users_id_seq', (select max(chat_users_id) from Chat_Users));
 select setval('users_user_id_seq', (select max(user_id) from Users));

