# Move from auto-migration to versioned migrations

## Goal

The goal of this exercise is to move from auto-migration to versioned migrations.
The existing code is using auto-migration to create the tables `User` and `Company`.
The goal is to create the table `Address` using a migration file and use it in the code.

## Steps

1. Create a new migration file:
```bash
make db-migrate-create name="delete_automigrate"
```

2. Dump current database schema:
```bash
make db-export-schema
```

3. Force `schema_migrations` table creation with the current schema:
```bash
migrate -verbose -database $DB_DSN -path migrations force 1
```

4. Create a new migration file:
```bash
make db-migrate-create name="delete_automigrate"
```

5. Edit the new migration file.

6. Delete from the code `db.AutoMigrate`

7. Create in the code a new Table Address

8. Use the new table in the code
