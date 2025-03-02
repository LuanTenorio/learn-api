-- Write your migrate up statements here

ALTER TABLE subjects ADD COLUMN user_id INTEGER NOT NULL;

ALTER TABLE subjects 
ADD CONSTRAINT fk_subjects_users FOREIGN KEY (user_id) 
REFERENCES users(id) ON DELETE CASCADE;

---- create above / drop below ----

ALTER TABLE subjects DROP CONSTRAINT fk_subjects_users;
ALTER TABLE subjects DROP COLUMN user_id;


-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
