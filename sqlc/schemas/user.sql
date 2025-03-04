CREATE TABLE "users" (
    "id" SERIAL NOT NULL,
    "name" VARCHAR(70) NOT NULL,
    "email" VARCHAR(70) NOT NULL,
    "password" VARCHAR(140) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);