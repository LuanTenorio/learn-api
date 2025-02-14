-- Write your migrate up statements here

-- CreateIndex
CREATE UNIQUE INDEX "users_email_key" ON "users"("email");

-- CreateIndex
CREATE UNIQUE INDEX "avarages_userId_key" ON "avarages"("userId");

-- CreateIndex
CREATE INDEX "subjects_name_idx" ON "subjects"("name");

-- CreateIndex
CREATE INDEX "contexts_name_idx" ON "contexts"("name");

-- CreateIndex
CREATE INDEX "reviews_date_done_idx" ON "reviews"("date", "done");

-- CreateIndex
CREATE INDEX "learns_content_type_idx" ON "learns"("content", "type");

-- CreateIndex
CREATE INDEX "days_date_idx" ON "days"("date");

-- CreateIndex
CREATE INDEX "weeks_sunday_idx" ON "weeks"("sunday");

-- CreateIndex
CREATE INDEX "months_firstDay_idx" ON "months"("firstDay");
