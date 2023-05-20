deps:
	npm i

test: export DATABASE_URL=file:test.db
test: deps
	npm run db:push
	npm test
