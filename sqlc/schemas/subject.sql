CREATE TABLE "subjects" (
    "id" SERIAL NOT NULL,
    "name" VARCHAR(150) NOT NULL,
    "total_time" INTEGER NOT NULL,
    "avarage" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "subjects_pkey" PRIMARY KEY ("id"),
    CONSTRAINT fk_subjects_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);