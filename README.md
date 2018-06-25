# gorm-bug-1731

https://github.com/jinzhu/gorm/issues/1731

## Reproduction steps

- `CREATE DATABASE test_gorm_db`;
- Apply the SQL from `sql.sql` (partition is created for 06.2018);
- `go run main.go`.

No insertion occurs:

```
felian@felian-VirtualBox:~/go_code/src/github.com/mtfelian/gorm-bug-1731$ go run main.go

(/home/felian/go_code/src/github.com/mtfelian/gorm-bug-1731/main.go:36)
[2018-06-25 19:03:47]  sql: no rows in result set

(/home/felian/go_code/src/github.com/mtfelian/gorm-bug-1731/main.go:36)
[2018-06-25 19:03:47]  [4.35ms]  INSERT INTO "model" ("id","rec_time") VALUES ('34f620bf-b456-4397-a434-cc157b2b8633','2018-06-25 19:03:47') RETURNING "model"."id"
[0 rows affected or returned ]
>>>>> CREATE: sql: no rows in result set

(/home/felian/go_code/src/github.com/mtfelian/gorm-bug-1731/main.go:41)
[2018-06-25 19:03:47]  [0.32ms]  SELECT count(*) FROM "model"
[0 rows affected or returned ]
0 rows found
That's all.
```

Manual requests:

```
felian@felian-VirtualBox:~/go_code/src/github.com/mtfelian/gorm-bug-1731$ psql --user=postgres --dbname=test_gorm_db
psql (10.4 (Ubuntu 10.4-2.pgdg14.04+1), server 9.6.9)
Type "help" for help.

test_gorm_db=# INSERT INTO "model" ("id","rec_time") VALUES ('34f620bf-b456-4397-a434-cc157b2b8633','2018-06-25 19:03:47') RETURNING "model"."id";
 id
----
(0 rows)

INSERT 0 0

test_gorm_db=# SELECT count(*) FROM "model";
 count
-------
     1
(1 row)

test_gorm_db=# SELECT * FROM "model";
                  id                  |      rec_time
--------------------------------------+---------------------
 34f620bf-b456-4397-a434-cc157b2b8633 | 2018-06-25 19:03:47
(1 row)


```