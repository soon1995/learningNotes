# MySQL Crash Course

Author: Ben Forta
Originally published: 2005, The content here might be outdated.

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

You can specify a port number for connections to a local server,
too. However, as indicated previously, connections to localhost on
Unix use a socket file by default, so unless you force a
TCP/IP connection as previously described, any option that specifies a port number is ignored.

For this command, the program uses a socket file on Unix and the --port option is ignored:

```bash
mysql --port=13306 --host=localhost
```

To cause the port number to be used, force a TCP/IP connection. For example, invoke the program in either of these ways:

```bash
mysql --port=13306 --host=127.0.0.1
mysql --port=13306 --protocol=TCP
```

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
`WHERE` Clause
`GROUP BY` Clause
`ORDER BY`Clause

### Foreign Key

A column in one table that contains the primary key values from
another (said B) table, thus defining the relationship between tables.

It:

- ensure B's information is never repeated, so time and space are not wasted.

- if B's information changes, just update in B

- As no data is repeated, the data used is obviously consistent,
making data reporting and manipulation much simpler.

Because of this, relational databases scale far better than
non-relational databases.

### Scale

Able to handle and increasing load without failing. A
well-designed database or application is said to *scale well*.

### Cartesian Product

The result returned by a table relationship without a join condition.

The join that return a Cartesian Product is called **cross join**.

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

- execute more quickly than lists of `OR`

- can work with subqueries

> **SQL vs Application Filtering**

Application Filtering is strongly discouraged.

1. Databases are optimized to perform filtering quickly and efficiently.

2. The server has to send unneeded data across network connections, resulting in a
waste of network bandwidth resources.

## Wildcard Filtering

```sql
% --
_ -- matches a single character
```

### Case-sensitivity

Depending on ow MySQL is configured.

### Watch for Trailing Spaces

'%hello' would not have matched 'hello '.

Solutions:

1. append another '%' at the end of wildcard

2. **trim**

### Do not overuse wildcards

if another search operator will do, use it instead.

### Not to use them at the beginning of search pattern

Unless absolutely necessary.

It is because search patterns that begin with wildcards are the slowest
to process.

## Regular Expressions

```sql
-- contains the text 1000
SELECT ... WHERE prod_name REGEXP '1000' ...
-- contains the text 1000 or 2000
SELECT ... WHERE prod_name REGEXP '1000|2000' ...

-- . matches any single character
SELECT ... WHERE prod_name REGEXP 'JetPack .000' ...
-- specify a set of characters, it is [1|2|3] Ton
SELECT ... WHERE prod_name REGEXP '[123] Ton' ...
-- negated
SELECT ... WHERE prod_name REGEXP '[^123] Ton' ...

-- matcing range
SELECT ... WHERE prod_name REGEXP '[0123456789] Ton' ...
SELECT ... WHERE prod_name REGEXP '[0-9] Ton' ...
SELECT ... WHERE prod_name REGEXP '[2-7] Ton' ...
SELECT ... WHERE prod_name REGEXP '[a-z] Ton' ...

-- matching special characters: \\
SELECT ... WHERE prod_name REGEXP 'JetPack \\.000' ...
\\f - form feed
\\n - line feed
\\r - carriage return
\\t - tab
\\\ - \

-- matching character classes
[:alnum:] - letter or digits, (sames as [a-zA-Z-0.9])
[:alpha:] - letter (same as [a-zA-Z])
[:blank:] - space or tab (same as [\\t ])
[:cntrl:] - ASCII control characters (ASCII 0 through 31 and 127)
[:digit:] - [0-9]
e.g WHERE prod_name REGXEP '[[:digit]]{4}'

-- Repetition metacharacters
. - 0 or more matches
+ - 1 or more matches (equavalent to [1,])
? - 0 or 1 match (equavalent to [0,1])
[n] - specific number of matches
[n,] - No less than a specified number of mathces
e.g. '\\([0-9] sticks?\\)' matches 'TNT (1 stick)', 'TNT (5 sticks)'

-- Anchor metacharacters
^ - start of text
$ - end of text
[[:<:]] - start of word
[[:>:]] - end of word
e.g. WHERE prod_name REGEXP '^[0-9\\.]'
```

**Tip**: **Matches are not case-sensitive**, to force case-sensitivity,
you can use `BINARY` keyword:

`WHERE prod_name REGEXP BINARY 'JetPack .000'`

**Tip**: You can use `SELECT` to test regular expressions without
using database tables. `REGEXP` checks always return 0 or 1

`SELECT 'hello' REGEXP '[0-9]'`

### \ OR \\?

Most regexp implementation use a single backslash to escape special
characters. However, MySQL requires two backslashes (MySQL itself interprets
one, and the regular expression library interprets the other).

## Creating Calculated Fields

```sql
-- Concatenating
`SELECT Concat(ven_name, ' (', vend_country, ')') FROM vendors ORDER BY vend_name`

-- Trim
RTrim
LTrim
Trim
`SELECT Concat(RTrim(vend_name), ' (', RTrim(vend_country), ')') FROM vendors ORDER BY vend_name`

-- Mathematical Operator
+
-
*
/
```

### **Client Versus Server Formatting**

It is far quicker to perform conversions and reformatting on the
database server than it is to perform them within the client because
DBMS are built to perform this type of processing quickly and efficiently.

### Most DBMS use operators + or || for concatenation

MySQL uses the Concate() function.

### How to test calculations

Use `SELECT`.

`SELECT 3 * 2;`

`SELECT RTrim('abc')`

`SELECT NOW()`

## Data Manipulation Functions

### Functions are less portable than SQL

Some DBMS supports functions that others don't.

Thus, many SQL programmers opt not to use any implementation-specific features. Although
this si a somewhat noble and idealistic view, it is not always in the best interests of application
performance. If you opt not to use these functions, you make your application code
work harder.

If you decide to use functions, make sure you comment your code well, so that at a
later date you (or your friends) will know exactly to which SQL implementation
you were writing.

### Functions

#### Text Functions

```sql
Left() -- returns characters from left of string
Length() -- returns the length of a string
Locate() -- finds a substring within a string
Lower() -- converts string to lowercase
LTrim() -- trims white space from left of string
Right() -- returns characters from right of string
RTrim() -- trims white space from right of string
Soundex() -- returns a string's soundex value, see below for explaination
SubString() -- returns characters from within a string
Upper() -- Converts string to uppercase
```

> SOUNDEX

SOUNDEX is an algorithm that converts any string of text into an alphanumeric
pattern describing the phonetic representation of that text. Enabling strings
to be compared by how they sound rather than how they have been typed.
Although SOUNDEX is not a SQL concept, MySQL (like many other DBMSs)

e.g.

```sql
SELECT cust_name, cust_contact
FROM customers
WHERE cust_contact = Soundex('Y. Lie');
|-----------|--------------|
| cust_name | cust_contact |
|-----------|--------------|
| Coyote Inc| Y Lee        |
|-----------|--------------|
```

#### Numberic Functions

```sql
Abs() -- Returns a number's absolute value
Cos() -- Returns the trigonometric cosine of a specified angle
Sin() -- Returns the trigonometric sine of a specified angle
Tan() -- Returns the trigonometric tangent of a specified angle
Exp() -- Returns the exponential value of a specific number
Mod() -- Returns the remainder of a division operation
Pi() -- Returns the value of pi
Rand() -- Returns a random number
Sqrt() -- Returns the square root of a specified number
```

#### Date and time Functions

```sql
AddData() -- Add to a date (days, weeks, and so on)
AddTime() -- Add to a time (hours, minutes, and so on)
CurDate() -- Returns the current date
CurTime() -- Returns the current time
Data() -- Returns the date portion of a date time
DateDiff() -- Calculates the difference between two dates
Date_Add() -- Highly flexible date arithmetic function
Date_Format() -- Returns a formatted date or time string
Day() -- Returns the day portion of a date
DayOfWeek() -- Returns the day of week for a date
Hour() -- Returns the hour portion of a time
Minute() -- Returns the minute portion of a time
Second() -- Returns the second portion of a time
Year() -- Returns the year portion of a date
Month() -- Returns the month portion of a date
Now() -- Returns the current date and time
Time() -- Returns the time portion of a date time
```

**Tip**: Always use four-digit year. It is far safer so MySQL does not have
to make any assumptions for you although it supports Two-digit years.

**Tip**: The date must be in the format `yyyy-mm-dd` when inserting, updating,
or filtering using `WHERE` clauses. Although other date formats might be recognized,
this format is the preferred date format because it eliminates ambiguity

##### **Date vs Datetime**

Is `WHERE order_date = '2005-09-01'` safe if its datatype is datetime?

NO, if order_date is `2005-09-01 11:30:05`, then this SQL will return nothing.

The solution: `WHERE Date(order_date) = '2005-09-01'`

##### Select orders placed in September 2005

`WHERE Date(order_date) BETWEEN '2005-09-01' AND '2005-09-30'`

or

`WHERE Year(order_date) = 2005 AND Month(order_date) = 9`

### System functions

#### Aggregate Functions

```sql
AVG() -- Returns a column's average value
COUNT() -- Returns the number of rows in a column
MAX() -- Returns a column's highest value
MIN() -- Returns a column's lowest value
SUM() -- Returns the sum of a column's value
DISTINCT() -- Returns distinct rows
```

**IMPORTANT**

- `NULL` Values Column are ignored by the `AVG()`, `MAX()`, `MIN()`, `SUM()` function

- `NULL` Values Column are ignored by the `COUNT()` function when COUNT(column)

##### `DISTINCT` With Other Aggregate function

`ALL` is default if distinct is not specified

```sql
SELECT AVG(DISTINCT prod_price) AS avg_price
FROM products
WHERE vend_id = 1003;
```

`DISTINCT` may only be used with count() if a column name is specified,
it may not be used with `COUNT(*)`, so `COUNT(DISTINCT *)` is not allowed

##### **Combining Aggregate Functions**

```sql
SELECT COUNT(*) AS num_items,
    MIN(prod_price) AS price_min,
    MAX(prod_price) AS price_max,
    AVG(prod_price) AS price_avg
FROM products
```

## Grouping Data

### Rules

1. `GROUP BY` can contain as many columns as you want.

2. Aside from the aggregate calculations statements, every column
in your `SELECT` statement should be present in the `GROUP BY` clause

3. If the grouping column contains a row with a `NULL` value, `NULL`
will be returned as a group.

4. The `GROUP BY` clause must come after any `WHERE` clause
and before any `ORDER BY` clause.

### ROLLUP

Using `ROLLUP` To obtain values at each group and at a summary level
(for each group)

`SELECT vend_id, COUNT(*) AS num_prods FROM products GROUP BY vend_id WITH ROLLUP;`

Example Result:

my:root@localhost=> SELECT user_id, count(*) FROM accounts GROUP BY user_id WITH ROLLUP;
 user_id | count(*)
---------+----------
       1 |        2
       2 |        2
         |        4

### Difference between `HAVING` and `WHERE`

`WHERE` before data is grouped.

`HAVING` after data is grouped.

## Working with Subqueries

```sql
SELECT cust_id
FROM orders
WHERE order_num IN (SELECT order_num
                    FROM orderitems
                    WHERE prod_id = 'TNT2');
```

**Tip**: Always formatting your SQL for easier read and debug

```sql
SELECT cust_name, cust_contact
FROM customers
WHERE cust_id IN (SELECT cust_id
                  FROM orders
                  WHERE order_num IN (SELECT order_num
                                      FROM orderitems
                                      WHERE prod_id = 'TNT2'));
```

**Caution**: Columns Must Match when using a subquery in a WHERE clause

**Caution**: Subqueries is not always the most efficient way to
perform this type of data retrieval, although it might be. Visit
[joining table](#joining-tables)

### Using Subqueries as calculated fields

```sql
SELECT cust_name,
       cust_state,
       (SELECT COUNT(*)
        FROM orders
        WHERE orders.cust_id = customers.cust_id) AS orders)
FROM customers
ORDER BY cust_name;
```

Above example, subquery is executed once for every customer retrieved.
The subquery is executed n times as same of rows returned.

**Note**: There are always more than one solution.

**Tip**: Build Queries with Subqueries incrementally, build and test
the innermost query first. Testing and debugging queries with subqueries
can be tricky.

## Joining Tables

**Tip**: There is often more than one way to perform any given SQL operation.
And there is rarely a definitive right or wrong way. Therefore, it is
often worth experimenting with different selection mechanisms
to find the one that works best for you.

### Referential Integrity

It is important that only valid data is inserted into relational columns.

To prevent invalid data from occurring, we can maintaining referential integrity by
specifying the primary and foreign keys as part of the table definition. [Creating
and Manipulating Tables](#creating-and-manipulating-tables)

### Use Join or Where?

Per the ANSI SQL specification, use of the `INNER JOIN` is preferable.
Using explicit join syntax ensures that you will never forget the join
condition. And it can affect performance, too (in some cases).

### Self Join instead of Subqueries

Although the end result is the same, sometimes these joins
execute far more quickly than they do subqueries. It is usually
worth experimenting with both to determine which performs better.

### Performance of Join

MySQL process joins at run-time, relating each table as specified.
This process can become very resource intensive, so be careful
not to use join tables unnecessarily. The more tables you join,
the more performance degrades.

### Natural Join

A natural join simply eliminates those multiple occurrences to only
one of each column is returned. There is no need to write an `ON`, natural
join will detect and join the column with similar name.

[here to understand more](https://www.w3resource.com/sql/joins/natural-join.php)

### Outer Join

Include rows that have no related rows.

- Right outer join

- Left outer join

### Use Joins with Aggregate Functions

```sql
SELECT customers.cust_name,
       customers.cust_id,
       COUNT(orders.order_num) AS num_ord
FROM customers INNER JOIN orders ON customers.cust_id = orders.cust_id
GROUP BY customers.cust_id;
```

## Combining Queries

`UNION`, `UNION_ALL`

Scenarios to use combined queries:

- To return similarly structured data from different tables in a single query

- To perform multiple queries against a single table returning the data as one query

**Tip**: `UNION` can be used to combine queries of different tables.

### Combining Queries vs Multiple `WHERE`

It can be accomplished by having multiple WHERE

The performance of each of the two techniques can vary based on the queries used.
Again, it is always good to experiment to determine which is preferable for specific
queries.

`WHERE` cannot accomplish `UNION ALL` which `UNION ALL` does not remove duplicate rows

### Examples

> You want to include all products made by vendor `1001` and `1002`

WHERE: `WHERE vendor_id IN (1001,1002) OR prod_price <= 5`

UNION:

```sql
SELECT .. FROM products WHERE prod_price <= 5
UNION
SELECT .. FROM products WHERE vender_id IN (1001,1002)
```

### Rules

1. Must be comprised of two or more `SELECT` statements

2. Each query in a `UNION` must contain the same columns, expressions, or aggregate functions
(although columns need not be listed in the same order). **Caution**, if the columns is not in same
order, it will compare based on column **instead of column name**.

Result:
my:root@localhost=> SELECT number, user_id FROM accounts WHERE user_id = 2 UNION select user_id, number FROM accounts WHERE balance = 0;
 number | user_id
--------+---------
 4123   | 2
 5123   | 2
 1      | ABC123
 2      | 4123
 2      | 5123

3. Column datatypes must be compatible: They need not be the exact same type, but they
must be of a type that MySQL can implicitly convert

### Union vs Union all

`UNION` removes any duplicate rows

Using `UNION ALL`, MySQL does not eliminate duplicates

### Sort With Union

Only one `ORDER BY` Clause may be used, and it must occur
after the final `SELECT` statement. **Multiple `ORDER BY` clauses
are not allowed**.

## Full-Text Searching

Not all engines support Full-Text Searching. The most commonly used
engines are MyISAM and InnoDB; the former supports full-text searching
and the latter does not.

`LIKE` and `REGEXP` has several very important limitations:

1. Performance: they requires that MySQL try and match each and every row in a table (
and table indexes are rarely of use in these searches). They can be very time-consuming
as the number of rows to be searched grows.

2. Explicit control: very difficult or not possible to explicitly control what is and what
is not matched. e.g. the first word is indeed matched.

3. No Intelligent results: No ranking. And searching for a specific word would not
find rows that did not contain that word but did contain other related words.

When using full-text searching, MySQL creates an indext of the words (in specified columns),
and search can be made against those words. Instead of look at each row individually, analyzing
and processing each word individually.

### Usage Notes

- When indexing full-text data, short words are ignored and are excluded from the
index. Short words are defined as **those having three or fewer characters** (this
number can be changed if needed).

- MySQL comes with a built-in list of *stopwords*, words that are always ignored
when indexing full-text data. This list can be overriden if needed.
`SELECT * FROM INFORMATION_SCHEMA.INNODB_FT_DEFAULT_STOPWORD;`

- Many words appear so frequently that searching on them would be useless
(too many results would be returned). As such, MySQL honors a 50% rule:
if a word appears in 50% or more rows, it is treated as a stopword and
is effectively ignored. (The 50% rule is not used for IN BOOLEAN MODE)

- Full-text searching never returns any results if there are fewer than
three rows in a table (because every word is always in at least 50% of the rows)

- Single quote characters in words are ignored. For example, `don't` is indexed as `dont`.

- Languages that don't have word delimiters (including Japanese and Chinese) will not
return full-text results properly. E.g. No rows returned.

- Full-text searching is only supported in the MyISAM database engine.

**Note**

One feature supported by many full-text search engines is **proximity searching**.
The ability to search for words that are near each other (in the same sentence,
same paragraph, or no more than a specific number of words apart, and so on)

This feature are not yet supported by MySQL full-text searching, although this
is planned for a future release.

### Columns Must be Indexed

After indexing, `SELECT` can be used with `MATCH()` and `Against()` to actually perform
the searches.

**Tip**: MySQL handles all indexing and re-indexing automatically

### Create Table

```sql
CREATE TABLE productnotes
(
    note_id         int         NOT NULL AUTO_INCREMENT,
    prod_id         char(10)    NOT NULL,
    note_date       datetime    NOT NULL,
    note_text       text        NULL,
    PRIMARY KEY(note_id),
    FULLTEXT(note_text)
) ENGINE=MyISAM;
```

Full-text searching is enabled when a table is created. Once defined,
MySQL automatically maintains the index.

**When rows are added, updated, or deleted, the index is automatically
updated accordingly.**

**Tip**: Don't use `FULLTEXT` When Importing Data as updating indexes takes
timenot a lot of time, but time nonetheless.

### Performing Full-Text Searches

```sql
SELECT note_text
FROM productnotes
WHERE Match(note_text) Against('rabbit');
```

**Note**:

The value passed to `MATCH()` Must be the
same as the one used in the `FULLTEXT()` definition. If multiple
columns are specified, all of them must be listed (and in the correct order)

```sql
CREATE TABLE multiple (
  first_text TEXT null,
  second_text TEXT null,
  FULLTEXT(first_text, second_text)
) ENGINE=MyISAM;

select * from multiple where match(first_text, second_text) against('foolish');
```

**Note**:

Searches are not case sensitive (unless `BINARY` mode is used)

### Result of Full-Text Searches

- they are ranked. Rows with a higher rank are returned first. [Demonstrate how ranking works](#demonstrate-how-ranking-works)

### Demonstrate How Ranking Works

```sql
SELECT note_text,
       Match(note_text) Against('rabbit') AS 'rank'
FROM productnotes;
```

The ranking is calculated by

- the number of words in the row

- the total number of words in the entire index

- the number of rows that contain the word

**Note**:

If multiple search terms are specified, those that contain the most matching words
will be ranked higher than those with less.

### Query Expansion

```sql
SELECT note_text
FROM productnotes
WHERE Match(note_text) Against('rabbit' WITH QUERY EXPANSION);
```

Some row has nothing to do with `rabbit`, but as it contains words that are
also in the rows with `rabbit`, it was retrieved, too.

**Tip**: The more row in your table, the better the results
return when using query expansion.

### Boolean Text Searches

In boolean mode, you may provide specifics as to

- Words to be matched

- Words to be excluded (row would not be returned even other specified were matched.)

- Ranking hints (specifying which words are more important than others)

- Expression grouping

- and more

**Tip**: Useable even without a `FULLTEXT` index. However, this would be a
very slow operation

**Tip**: In boolean mode, rows will not be returned sorted descending by ranking score.
**They are ranked, but not sorted**.

| Privilege | Description |
| --- | --- |
| + | Include, word must be present |
| - | Exclude, word must not be present |
| > | Include, and increase ranking value |
| < | Include, and decrease ranking value |
| () | Group words into subexpressions (allowing them to be include, excluded, ranked, and so fort as a group) |
| ~ | Negate a word's ranking value |
| * | Wildcard at end of word |
| "" | Defines a phrase (as opposed to a list of individual words, the entire phrase is matched for inclusion or exclusion) |

```sql
SELECT note_text
FROM productnotes
WHERE Match(note_text) Against('heavy -rope*' IN BOOLEAN MODE);
-- this match rows that have heavy but without rope
```

```sql
Against ('rabbit bait' IN BOOLEAN MODE);
# is different with
Against ('+rabbit +bait' IN BOOLEAN MODE);

# The former: `I have rabbit` row would be retrieved
# The latter: would not
```

```sql
Against ('"rabbit bait"' IN BOOLEAN MODE);
# It match `rabbit bait hello`
# It would not match `rabbit hello bait`
```

```sql
Against ('>rabbit <carrot' IN BOOLEAN MODE);
# Match both `rabbit` and `carrot`, increasing the rank of the former and decreasing
the rank of the latter
```

```sql
Against ('+rabbit +(<carrot)' IN BOOLEAN MODE);
# Match both `rabbit` and `commbination`, decreasing the rank of the latter
```

## Insert Data

`INSERT` statements usually generate no output.

**Tip**: `INSERT INTO xtable VALUES(...)` is indeed simple,
it is not at all safe and should generally be avoided at all costs. It
is highly dependent on the order in which the columns are defined in the table.

**Tip**: Always use a columns list.

### Better Performance

> **INSERT LOW_PRIORITY INTO**

`INSERT` can be time consuming, especially if there are many indexes to be
updated. If data retrieval is of utmost important (as is usually is), you can
instruct MySQL to lower the priority of your `INSERT` statement:
`INSERT LOW_PRIORITY INTO ...`. I could be applied to `UPDATE` AND `DELETE`

> Use single `INSERT` statement

Single `INSERT` statement is faster than multiple `INSERT` statements.

### INSERT SELECT

```sql
INSERT INTO customers (cust_id,
  cust_contact,
  cust_email,
  cust_name,
  cust_address,
  cust_city,
  cust_state,
  cust_zip,
  cust_country)
SELECT cust_id,
  cust_contact,
  cust_email,
  cust_name,
  cust_address,
  cust_city,
  cust_state,
  cust_zip,
  cust_country
FROM custnew;
```

**Tip**: There is no requirement that the column names match.
**In fact, the column position is used, so the first column
in the `SELECT` (regardless of its name) will be used to populate
the first specified table column, and so on.**

## Update and Delete

**Tip**: To continue processing updates, even if an error
occurs, use the `IGNORE` keyword like: `UPDATE IGNORE customers ...`

**Tip**: Use `TRUNCATE TABLE` if you want to delete all rows from a table instead
of `DELETE`. `TRUNCATE` actually drops and recreates the table instead of deleting each
row individually.

**Tip**: Before using a `WHERE` to `UPDATE` OR `DELETE`, first test it with
a `SELECT` to make sure it is filtering the right records

**Tip**: Use database enforced referential integrity, so MySQL will not
allow the deletion of rows that have data in other tables related to them.

## Creating and Manipulating Tables

```sql
CREATE TABLE customers
(
  cust_id       int         NOT NULL AUTO_INCREMENT,
  cust_name     char(50)    NOT NULL,
  cust_address  char(50)    NULL,
  PRIMARY KEY (cust_id)
) ENGINE=InnoDB;
```

**Tip**: Whitespace is ignored in MySQL statements. Always format your statement
for easier reading and editing.

**Tip**: Creating an exist table name is not allowed.

**Tip**: `CREATE TABLE IF NOT EXISTS table_name` does not check if the table exist.

## Constraint

### AUTO_INCREMENT

Only one `AUTO_INCREMENT` column is allowed per table,
and it must be indexed (e.g. make it a primary key).

Simply specify a value in the `INSERT` statement as it is unique
to overriding `AUTO_INCREMENT` column. Subsequent incrementing
will start using

- the value manually inserted if the value is larger or equal to existing auto_increment value.

- existing auto_increment value if the value manually inserted is lesser than auto_increment value.

#### Determine the `AUTO_INCREMENT_VALUE`

`SELECT last_insert_id();`

**Tip**: Different session has different result of last_insert_id();

### DEFAULT

**Caution**: In MySQL, functions as `DEFAULT` value are not allowed unlike most DBMSs.

**Tip**: Use `DEFAULT` instead of `NULL` Values, especially in columns that will be
used in calculations or data groupings.

## Engine Types

Engine is used internally to process your request. For the most part,
the engine is buried within the DBMS and you need not pay much attention to it.

**Tip**: when creating a table without `ENGINE=` statement, the default engine is used
(most likely `MyISAM`)

### Types

- `InnoDB` is a transaction-safe engine. ~~It does not support full-text searching~~(supported now).

- `Memory` is functionally equivalent to `MyISAM`, but as data is stored in memory (instead of on disk)
it is extremely fast (and ideally suited for temporary table). It does not support transactional processing.

- `MyISAM` is a very high-performance engine. It supports full-text searching, but it does not
support transactional processing.

**Caution**: Foreign Keys Can't Span Engines.

## ALTER TABLE

```sql
-- Foreign key
ALTER TABLE table_name ADD CONTRAINT fk_orderitems_orders
FOREIGN KEY (order_num) REFERENCES orders (order_num);

--
```

### Complex table structure changes

Usually require a manual move process involving these steps:

1. Create a new table with the new column layout.

2. Use the `INSERT SELECT` statement to copy the data from
the old table to the new table. Use conversion functions and calculated fields, if needed.

- Verify that the new table contains the desired data.

- Rename the old table (or delete it, if you are really brave).

- Rename the new table with the name previously used by the old table.

- Re-create any triggers, stored procedures, indexes, and foreign keys as needed.

### Rename Table

```sql
RENAME TABLE customers2 TO customers;

-- Multiple:
RENAME TABLE backup_customers TO customers,
             backup_vendors TO vendors,
             backup_products TO products;
```

## Views

Common uses:

- To reuse SQL statements

- To simplify complex SQL operations -- without having to know the details
of the underlying query itself.

- To expose parts of a table instead of complete tables -- more secure

- To change data formatting and representation.

**Caution**: Views contain no data, any retrieval needed to
execute a query must be processed every time the view is used. If you
create complex views with multiple joins and filters, or if you
nest views, **you may find that performance is dramatically degraded**.

### Rules and Restrictions

- views must be uniquely named

- No limit to the number of views that can be created.

- You must have security access to create views

- Views can be nested (one view retrieves data from another view)

- `ORDER BY` may be used in a view, but it will be overriden if `ORDER BY`
is also used in the `SELECT` that retrieves data from the view.

- Views cannot be indexed, nor can they have triggers or default values
associated with them

  - It return error if we try to set: viewName  is not BASE TABLE

- Views can be used in conjunction with tables, for example, to create a
`SELECT` statement which joins a table and a view.

### SQL

```sql
-- Create View
CREATE VIEW

-- View the statement used to create a view
SHOW CREATE VIEW viewname;

-- Remove a view
DROP VIEW viewname;

-- Update a view
-- Drop first then create
-- OR
CREATE OR REPLACE VIEW
```

#### Example Create

> Simplify Complex Joins

```sql
CREATE VIEW productcustomers AS
SELECT cust_name, cust_contact, prod_id
FROM customers, orders, orderitems
WHERE customers.cust_id = orders.cust_id
  AND orderitems.order_num = orders.order_num;
```

```sql
CREATE VIEW accountusers AS
SELECT accounts.id, accounts.number, accounts.balance, users.name, accounts.balance*2 balance2
FROM accounts, users
WHERE accounts.user_id = users.id;
```

**Note**: Both `WHERE` clause (in the view and the one passed to it) will be combined automatically.

> Reformat Retrieved Data

```sql
SELECT Concat(RTrim(vend_name),  ' (', RTrim(vend_country), ')')
  AS vend_title
FROM vendors ORDER BY vend_name
```

> Using Views with calculated fields

```sql
SELECT prod_id, quantity, item_price, quantity*item_price AS expanded_price ...
```

### Updating Views

**Caution**: In reality, views are primarily used for data retrieval. Not for updates

If you add or remove rows from a view that you are actually removing them from the underlying table.

Whether the view cannot be updated if following are used:

- Grouping (using `GROUP BY` and `HAVING`)

- Joins

- Subqueries

- Unions

- Aggregate functions (`Min()`, `Count()`, `Sum()`, and so forth)

- `DISTINCT`

- Derived (calculated) columns

## Stored Procedures

Stored procedures are very useful and should be used whenever possible.

> Benefits

- To simplify complex operations by encapsulating processes into a single easy-to-use unit.

- To ensure data integrity by not requiring that a series of steps be created over and over.
All developers and applications use the same stored procedure.

- To simplify change management. If tables, column names, or business logic changes,
only the stored procedure code needs to be updated.

- Security, restricting access to underlying data via stored procedures reduces the
chance of data corruption

- to improve performance, as stored procedures typically execute quicker than do individual

> **Downside**

- tend to be more complex to write than basic SQL statements

- you might not have the security access needed to create stored procedures. (You can still use them
although you can't write them)

### Executing

```sql
CALL productpricing(@pricelow, @pricehigh, @priceaverage);
```

### Creating

```sql
CREATE PROCEDURE productpricing()
BEGIN
  SELECT Avg(prod_price) AS priceaverage
  FROM products;
END;
```

**Note**: The default MySQL delimiter is `;`, however, the sql
command-line utility also uses `;` as a delimiter. If the command-line
utility were to interpret the `;` characters inside of the stored procedure
itself, those would not end up becoming part of the stored procedure, **that would
make the SQL in the stored procedure syntactically invalid.**

Solution is temporarily change the command-line utility delimiter:

```mysql
DELIMITER //

CREATE PROCEDURE productpricing()
BEGIN
  SELECT Avg(prod_price) AS priceaverage
  FROM products;
END //
```

```mysql
DELIMITER //

CREATE PROCEDURE productpricing()
BEGIN
  SELECT Avg(balance) AS balanceavg
  FROM accounts;
END //
```

### Drop

```sql
DROP PROCEDURE productpricing;
```

**Tip**: `DROP PROCEDURE IF EXISTS` doesn't throw an error if the named procedure does not actually exist

### With Parameters

> OUT

```sql
DELIMITER //
CREATE PROCEDURE productpricing(
  OUT pl DECIMAL(8,2),
  OUT ph DECIMAL(8,2),
  OUT pa DECIMAL(8,2)
)
BEGIN
  SELECT Min(prod_price) INTO pl FROM products;
  SELECT Max(prod_price) INTO ph FROM products;
  SELECT Avg(prod_price) INTO pa FROM products;
END //

DELIMITER ;
```

```sql
-- table:
CREATE TABLE products (
  prod_price DECIMAL(8,2) DEFAULT 0
) ENGINE=InnoDB;
INSERT INTO products (prod_price) VALUES (1),(2),(5);
```

Recordsets is not an allowed type, and so multiple rows and columns could not
be returned via a parameter. This is why three parameters are used in above example.

**NOTE**: Variable names (MySQL) must begin with `@`.

```sql
CALL productpricing(@pricelow, @pricehigh, @priceaverage);
SELECT @priceaverage; -- after call, display price average
SELECT @priceaverage, @pricehigh, @pricelow;
```

> IN

```sql
DELIMITER //
CREATE PROCEDURE ordertotal(
  IN onumber INT,
  OUT ototal DECIMAL(8,2)
)
BEGIN
  SELECT Sum(item_price*quantity)
  FROM orderitems
  WHERE order_number = onumber
  INTO ototal;
END //
DELIMITER ;

-- use it
CALL ordertotal(20005, @total);
```

```sql
-- table:
CREATE TABLE orderitems (
  order_number INT PRIMARY KEY AUTO_INCREMENT,
  item_price DECIMAL(8,2),
  quantity INT
) ENGINE=InnoDB;
INSERT INTO orderitems (item_price, quantity) VALUES
  (22.0, 2),
  (32.0, 1),
  (42.0, 3),
  (12.0, 3)
  ;
```

#### Perfect Example

```sql
-- Name: ordertotal
-- Parameters: onumber = order number
--             taxable = 0 if not taxable, 1 if taxable
--             ototal = order total variable

DELIMITER //

CREATE PROCEDURE ordertotal (
  IN onumber INT,
  IN taxable BOOLEAN,
  OUT ototal DECIMAL(8,2)
) COMMENT 'OBTAIN order total, optionally adding tax'

BEGIN
  -- Declare variable for total
  DECLARE total DECIMAL(8,2);
  -- Declare tax percentage
  DECLARE taxrate INT DEFAULT 6;

  -- Get the order total
  SELECT SUM(item_price*quantity)
  FROM orderitems
  WHERE order_number = onumber
  INTO total;

  -- Is this taxable?
  IF taxable THEN
    -- Yes, so add taxrate to the total
    SELECT (total+(total/100*taxrate)) INTO total;
  END IF;

  -- and finally, save to out variable
  SELECT total INTO ototal;
END //
DELIMITER ;
```

**NOTE**: `COMMENT` is not required, but if specified, is
displayed in `SHOW PROCEDURE STATUS` result.

Use `SHOW PROCEDURE STATUS LIKE 'ordertotal'` to limit the results.

### Inspecting Stored Procedures

`SHOW CREATE PROCEDURE ordertotal;`

## Cursors

> Cursor is the result set retrieved by the statement. Once
the cursor is stored, applications can scroll or browse up
and down through the data as needed.

**Caution**: Unlike most DBMSs, MySQL cursors may only be used
within stored procedures (and functions).

### Creating Cursors

```sql
DELIMITER //
CREATE PROCEDURE processorders()
BEGIN
  DECLARE ordernumbers CURSOR
  FOR
  SELECT ordernum FROM orders;
END //
DELIMITER ;
```

### Open and close cursor

```sql
OPEN ordernumbers;
```

When the `OPEN` statement is processed, the query is executed, and the retrieved data
is stored for subsequent browsing and scrolling.

After cursor processing is complete, the cursor should be closed using the `CLOSE`
statement. `CLOSE` frees up any internal memory and resources used by the cursor.

```sql
CLOSE ordernumbers;
```

**Caution**: Cursor cannot be reused without being opened again.
However, a cursor does not need to be declared again to be used

**Tip**: If you do not explicitly close a cursor, MySQL will close it automatically
when the `END` statement is reached.

```sql
DELIMITER //
CREATE PROCEDURE processorders()
BEGIN
  DECLARE ordernumbers CURSOR
  FOR
  SELECT ordernum FROM orders;

  -- Open the cursor
  OPEN ordernumbers;

  -- Close the cursor
  CLOSE ordernumbers;
END //
DELIMITER ;
```

### Using Cursor Data

After a cursor is opened, each row can be accessed individually using a `FETCH` statement.
`FETCH` specifies what is to be retrieved (desired columns) and where retrieved data should
be stored. It also advances the internal row pointer within the cursor (call next)

```sql
DELIMITER //
CREATE PROCEDURE processorders()
BEGIN
  -- Declare local variables
  DECLARE o INT;

  DECLARE ordernumbers CURSOR
  FOR
  SELECT ordernum FROM orders;

  -- Open the cursor
  OPEN ordernumbers;

  -- Get order number
  FETCH ordernumbers INTO o;

  -- Close the cursor
  CLOSE ordernumbers;
END //
DELIMITER ;
```

```sql
DELIMITER //
CREATE PROCEDURE processorders()
BEGIN
  -- Declare local variables
  DECLARE done BOOLEAN DEFAULT 0;
  DECLARE o INT;

  -- Declare the cursor
  DECLARE ordernumbers CURSOR
  FOR
  SELECT ordernum FROM orders;

  -- Declare continue handler
  DECLARE CONTINUE HANDLER FOR SQLSTATE '02000' SET done=1;

  -- Open the cursor
  OPEN ordernumbers;

  -- Loop through all rows
  REPEAT
    -- Get order number
    FETCH ordernumbers INTO o;

  -- END of loop
  UNTIL done END REPEAT;

  -- Close the cursor
  CLOSE ordernumbers;
END //
DELIMITER ;
```

`CONTINUE HANDLER`,  code will be executed when a condition occurs.
`SQLSTATE '02000'` is a `not found` condition and so it occurs when `REPEAT`
cannot continue because there are no more rows to loop through.

[For more error codes](https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html)

**Caution**: Local variables defined with `DECLARE` must be defined before any cursors or
handlers are defined, **and handlers must be defined after any cursor**. Failure to follow
this sequencing will generate an error message.

**Note**

`REPEAT` OR `LOOP`?  `LOOP` is manually exited using a `LEAVE` statement. In general,
the syntax of the `REPEAT` statement makes it better suited for looping through cursors.

### Example

```sql
DELIMITER //
CREATE PROCEDURE processorders()
BEGIN
  -- Declare local variables
  DECLARE done BOOLEAN DEFAULT 0;
  DECLARE o INT;
  DECLARE t DECIMAL(8,2);

  -- Declare the cursor
  DECLARE ordernumbers CURSOR
  FOR
  SELECT ordernum FROM orders;

  -- Declare continue handler
  DECLARE CONTINUE HANDLER FOR SQLSTATE '02000' SET done=1;

  CREATE TABLE IF NOT EXISTS ordertotals
    (order_num INT, total DECIMAL(8,2));

  -- Open the cursor
  OPEN ordernumbers;

  -- Get order number
  FETCH ordernumbers INTO o;

  -- Loop through all rows
  REPEAT
    -- Get the total for this order;
    CALL ordertotal(o, 1, t);

    -- Insert order and total into ordertotals
    INSERT INTO ordertotals (order_num, total) VALUES(o, t);

    -- Get order number
    FETCH ordernumbers INTO o;
  -- END of loop
  UNTIL done END REPEAT;

  -- Close the cursor
  CLOSE ordernumbers;
END //
DELIMITER ;
```

## Triggers

Example usage:

- Every time a customer is added to a database table, **check** that
the phone number is formatted correctly and the state abbreviation is in uppercase

- Every time a product is ordered, subtract the ordered quantity from the
number in stock.

- Whenever a row is deleted, save a copy in an archive table.

### Create Triggers

Specify:

- Unique trigger name

- The table to which the trigger is to be associated

- The action that the trigger should respond to (`DELETE`, `INSERT`, Or `UPDATE`)

- When the trigger should be executed (before or after processing)

**Tip**: Keep Trigger Name unique per database although MySQL allows it.
Other DBMSs restricted trigger name unique per database

```sql
CREATE TRIGGER newproduct AFTER INSERT ON products
FOR EACH ROW SELECT 'Product added';
```

**Above example doesn't work: error - not allowed to return a result set from a trigger**

```sql
DELIMITER //
CREATE TRIGGER neworderitem AFTER INSERT ON orderitems
FOR EACH ROW 
BEGIN
  INSERT INTO orders values (null,5);
END //
DELIMITER ;
```

- `FOR EACH ROW` for each inserted row

**Note**: Triggers only supported on table, **not on views**. However,
the trigger in table will be triggered if you make changes in the view that table being referred

**Tip**: When `BEFORE` trigger fail, MySQL will not perform the requested
operation. If `BEFORE` or sql statement failed, the `AFTER` trigger will not be executed.

If `AFTER` trigger failed, the sql statement will be rollback.

**Tip**: Trigger cannot be updated or overwritten. To modify a trigger,
it must be dropped and re-created.

### How many trigger can be used?

> ~~Up to six triggers are supported per table~~

~~Before and after each of `INSERT`, `UPDATE`, `DELETE`~~

*I have tried to create same INSERT AFTER, it works*

### Drop

`DROP TRIGGER name;`

### `INSERT` Triggers

- Within `INSERT` trigger code, you can refer to a virtual table named `NEW` to access the rows being inserted

- In a `BEFORE INSERT` trigger, the values in `NEW` may also be updated (allowing you to change values about to be inserted)

- For `AUTO_INCREMENT` columns, `NEW` will contain `0` before and the new automatically generated value after

**Use `BEFORE` for any data validation and cleanup**

### `UPDATE` Triggers

**Use `BEFORE` for any data validation and cleanup**

- Within `UPDATE` trigger code, you can refer to a virtual table named `OLD`
to access the previous values and the `NEW` to access the new updated values.

- In a `BEFORE UPDATE` trigger, the values in `NEW` may also be updated (allowing)
you to change values about to be used in the `UPDATE` statement

- The values in `OLD` are all read-only and cannot be updated.

Below examples ensures the state abbreviations are always in uppercase

```sql
CREATE TRIGGER updatevendor BEFORE UPDATE ON vendors
FOR EACH ROW SET NEW.vend_state = UPPER(NEW.vend_state);

CREATE TRIGGER updateorderitems BEFORE UPDATE ON orderitems
FOR EACH ROW SET NEW.abc = UPPER(NEW.abc);
```

### `DELETE` Triggers

- Within `DELETE` trigger code, you can refer to a virtual table named `OLD` to access the rows being deleted.

- The values in `OLD` are all read-only and cannot be updated.

Below example use `OLD` to save rows about to be deleted into an archive table

```yaml
DELIMITER //
CREATE TRIGGER deleteorder BEFORE DELETE ON orders
FOR EACH ROW
BEGIN
  INSERT INTO archive_orders(order_num) VALUES (OLD.ordernum);
END //
DELIMITER ;
```

The advantage of using `BEFORE DELETE` (instead of `AFTER`) is if order could not be archived,
the `DELETE` itself will be aborted. Although the `DELETE` itself will be rollback if order could not be archived in `AFTER`.

**Tip**: there is no harm of `BEGIN` and `END` not being used. With `BEGIN` and `END`, multiple SQL statements are allowed

### Other notes on trigger

- Creating triggers might require special security access

- Triggers should be used to ensure data consistency (case, formatting, and so on).

  - The advantage of performing this type of processing in a trigger is it happens transparently regardless of
  client application.

- **One interesting use for triggers is in creating an audit trail**

## Other

### Difference between `NULL` and false value

1. Some aggregate functions ignore `NULL` column

2. `NULL` is a group when using `GROUP BY`

3. When WHERE a <> 'b', null value will not return

## To explore

- Experimenting Selfjoins vs subqueries

- Appendix B to learn more INSERT

- default engine

```bash
 docker run -it --rm \
    imega/mysql-client \
    mysql --host=localhost --port=3306 --user=root --password=qwer
```

`sudo apt-get install mysql-client`
