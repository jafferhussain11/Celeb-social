These two steps are doing different levels of database connection setup:
Step 1: GORM Connection
godb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
What it does:

Creates a GORM wrapper around the database connection
Sets up GORM's ORM layer (models, migrations, query builder, etc.)
Returns: *gorm.DB - the ORM interface

What you get:

GORM methods like .Find(), .Create(), .Where()
Automatic SQL generation
Model mapping, hooks, etc.

Step 2: Raw SQL Connection
gosqlDB, err := db.DB()
What it does:

Extracts the underlying *sql.DB from GORM's wrapper
Gives you direct access to the raw database driver
Returns: *sql.DB - the standard library database connection

What you get:

Direct SQL execution with .Query(), .Exec()
Connection pool configuration
Database-level settings

Why Do Both?
You typically get the *sql.DB to configure connection settings:
go// Step 1: GORM setup
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// Step 2: Get raw connection for configuration
sqlDB, err := db.DB()

// Configure connection pool (this needs *sql.DB)
sqlDB.SetMaxIdleConns(10)
sqlDB.SetMaxOpenConns(100)
sqlDB.SetConnMaxLifetime(time.Hour)

// Now use GORM for actual queries
db.Find(&users)
db.Create(&newUser)
Side Note - Bug in Your Code:
gofmt.Sprintf("Unable to connect to database") // This does nothing!
Should be:
gofmt.Println("Unable to connect to database") // Actually prints
// or
log.Fatal("Unable to connect to database")   // Logs and exits
Summary:

gorm.Open() → Sets up the ORM layer
db.DB() → Gets the raw connection for configuration
Use db (GORM) for queries, use sqlDB for connection settings