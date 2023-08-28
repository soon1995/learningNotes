# MySQL Crash Course

Author: Ben Forta
Originally published: 2005

> Misuse Causes Confusion

Database Software is actually called the Database Management System (DBMS), not database.
A Database might be a file stored on a hard drive, but it might not.
You always use the DBMS and it accesses the database for you.

> Breaking Up Data

Breaking Up Data is extremely important. For example, city, state and ZIP code
should always be separate columns. By breaking these out, it becomes possible to sort
or filter data by specific columns. If city and state are combined
into one column, it would be extremely difficult to sort or filter by state.

> Always Define Primary Keys

Always define primary keys, although primary keys are not actually required. So future
data manipulation is possible and manageable.

> Primary Key Best Practices

- Don't update values in primary key columns.

- Don't reuse values in primary key columns.

- Don't use values that might change in primary key columns. E.g. when you
use a name as a primary key to identify a supplier, you would
have to change the primary key when the supplier merges and changes its name.

## Command

```bash
mysql
mysql u ben
mysql u ben p h myserver P 9999 # u username, h host, P port
mysql help
mysql> select * from a;
mysql> select * from a \g # show column line by line
mysql> help
mysql> \h # same with help
mysql> help select # to obtain help on using the SELECT statement
mysql> quit
mysql> exit

# Information
mysql> SHOW DATABASES; # SHOW command is used to display database, table, column, user, privileges and more information
mysql> SHOW TABLES;
mysql> SHOW COLUMNS FROM customers;
mysql> DESCRIBE customers; # same with SHOW COLUMNS FROM customers;
mysql> SHOW STATUS; # extensive server status information
mysql> SHOW CREATE DATABASE; # used to display the MySQL statement used to create specified databases
mysql> SHOW CREATE TABLE; # ditto but tables
mysql> SHOW GRANTS; # display security rights granted to users (all users or a specific user)
mysql> SHOW ERRORS;
mysql> SHOW WARNINGS;
mysql> HELP SHOW;

# Database
mysql> USE databaseabc;


```

## Term

### Schema

Table layout and properties.

### Clause

SQL Statements are made up of clauses, some required and some optional.
A clause usually consists of a keyword and supplied data.

`FROM` Clause
`ORDER BY`Clause
`WHERE` Clause

## SELECT

### Better Off Not Using the `*` Wildcard

Unless you really do need every column in the table.

Even though use of wildcards might save you the time and effort needed to
list the desired columns, **retrieving unnecessary columns usually slows down
the performance of your retrieval and your application.**

### DISTINCT

`SELECT DISTINCT vend_id FROM products.`

`DISTINCT` keyword applies to all columns, not just the one it precedes.
If you were to specify `SELECT DISTINCT vend_id, prod_price`, all rows
would be retrieved unless both of the specified columns were distinct.

### LIMIT

`SELECT prod_name FROM products LIMIT 5`

`SELECT prod_name FROM products LIMIT 5,5`, the first number is
where to start, the second is the number of rows to retrieve.

**The first row retrieved is row `0`, not row `1`. LIMIT 1,1 retrived
the second row, not the first one**

However, it is confusing, MYSQL 5 has an alternative syntax:
`SELECT prod_name FROM products LIMIT 4 OFFSET 3`, it
get 4 rows starting from row 3 (The fourth row)

### SORT

> **Case Sensitivity?**

It depends on how the database is set up. Default behavior in MySQL
(and indeed most DBMSs) is **not case sensitive**. However, administrators
can change this behaviour if needed. **If your database contains lots
of foreign language characters, this might become necessary**.

### WHERE

> **Case Sensitive**

By default, MySQL is not case sensitive when performing matches, so
`WHERE prod_name = 'fuses'`, `Fuses` and `fuses` matched.

> `NULL` **and Nonmatches**

When `WHERE a != 'b'`. `NULL` rows will **not** returned. The database
does not know whether they match.

> `OR` and `AND` Order

`AND` first, `OR` later.

Use parentheses to explicitly group related operators.

Don't ever rely on the default evaluation order even if it is exactly what
you want. There is no downside to use parentheses, and you are
always better off eliminating any ambiguity.

> `IN` is better than a lot of `OR`

- Syntax is far cleaner and easier to read.

- 

> **SQL vs Application Filtering**

Application Filtering is strongly discouraged.

1. Databases are optimized to perform filtering quickly and efficiently.

2. The server has to send unneeded data across network connections, resulting in a
waste of network bandwidth resources.

## To explore

- SHOW STATUS;

- SHOW GRANTS;

- SHOW ERRORS;

- SHOW WARNINGS;

- WHERE a = 'b' to see if the null returned
