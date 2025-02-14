-- CreateEnum
CREATE TYPE "learn_type" AS ENUM ('reading', 'study', 'review', 'exercise');

-- CreateTable
CREATE TABLE "users" (
    "id" SERIAL NOT NULL,
    "name" VARCHAR(70) NOT NULL,
    "email" VARCHAR(70) NOT NULL,
    "password" VARCHAR(140) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "avarages" (
    "months" INTEGER NOT NULL,
    "weeks" INTEGER NOT NULL,
    "days" INTEGER NOT NULL,
    "total" INTEGER NOT NULL,
    "userId" INTEGER NOT NULL
);

-- CreateTable
CREATE TABLE "subjects" (
    "id" SERIAL NOT NULL,
    "name" VARCHAR(150) NOT NULL,
    "total_time" INTEGER NOT NULL,
    "avarage" INTEGER NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "subjects_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "contexts" (
    "id" SERIAL NOT NULL,
    "total_time" INTEGER NOT NULL,
    "name" VARCHAR(150) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "contexts_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "subject_context" (
    "subject_id" INTEGER NOT NULL,
    "context_id" INTEGER NOT NULL,

    CONSTRAINT "subject_context_pkey" PRIMARY KEY ("subject_id","context_id")
);

-- CreateTable
CREATE TABLE "reviews" (
    "id" SERIAL NOT NULL,
    "date" TIMESTAMPTZ NOT NULL,
    "done" BOOLEAN NOT NULL,
    "position" SMALLINT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "reviews_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "subject_review" (
    "subject_id" INTEGER NOT NULL,
    "review_id" INTEGER NOT NULL,

    CONSTRAINT "subject_review_pkey" PRIMARY KEY ("subject_id","review_id")
);

-- CreateTable
CREATE TABLE "learns" (
    "id" SERIAL NOT NULL,
    "content" VARCHAR(150) NOT NULL,
    "start" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "end" TIMESTAMPTZ NOT NULL,
    "type" "learn_type" NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "subjectId" INTEGER,

    CONSTRAINT "learns_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "days" (
    "id" SERIAL NOT NULL,
    "date" TIMESTAMPTZ NOT NULL,
    "total_time" INTEGER NOT NULL,
    "userId" INTEGER,

    CONSTRAINT "days_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "weeks" (
    "id" SERIAL NOT NULL,
    "sunday" DATE NOT NULL,
    "average_all_days" INTEGER NOT NULL,
    "average_days_studied" INTEGER NOT NULL,
    "total_time" INTEGER NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "userId" INTEGER,

    CONSTRAINT "weeks_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "months" (
    "id" SERIAL NOT NULL,
    "firstDay" DATE NOT NULL,
    "average_all_days" INTEGER NOT NULL,
    "average_days_studied" INTEGER NOT NULL,
    "total_time" INTEGER NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "userId" INTEGER,

    CONSTRAINT "months_pkey" PRIMARY KEY ("id")
);
