Enter password: **********
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 9
Server version: 8.0.23 MySQL Community Server - GPL

Copyright (c) 2000, 2021, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.


mysql> show engines;
+--------------------+---------+----------------------------------------------------------------+--------------+------+------------+
| Engine             | Support | Comment                                                        | Transactions | XA   | Savepoints |
+--------------------+---------+----------------------------------------------------------------+--------------+------+------------+
| MEMORY             | YES     | Hash based, stored in memory, useful for temporary tables      | NO           | NO   | NO         |
| MRG_MYISAM         | YES     | Collection of identical MyISAM tables                          | NO           | NO   | NO         |
| CSV                | YES     | CSV storage engine                                             | NO           | NO   | NO         |
| FEDERATED          | NO      | Federated MySQL storage engine                                 | NULL         | NULL | NULL       |
| PERFORMANCE_SCHEMA | YES     | Performance Schema                                             | NO           | NO   | NO         |
| MyISAM             | YES     | MyISAM storage engine                                          | NO           | NO   | NO         |
| InnoDB             | DEFAULT | Supports transactions, row-level locking, and foreign keys     | YES          | YES  | YES        |
| BLACKHOLE          | YES     | /dev/null storage engine (anything you write to it disappears) | NO           | NO   | NO         |
| ARCHIVE            | YES     | Archive storage engine                                         | NO           | NO   | NO         |
+--------------------+---------+----------------------------------------------------------------+--------------+------+------------+
9 rows in set (0.21 sec)

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| authentication     |
| information_schema |
| myisam             |
| mysql              |
| performance_schema |
| sreenath           |
| sys                |
+--------------------+
7 rows in set (0.84 sec)

mysql> use myisam;
Database changed

mysql> show tables;
+------------------+
| Tables_in_myisam |
+------------------+
| employee         |
| student          |
+------------------+
2 rows in set (0.34 sec)

mysql> drop table student;
Query OK, 0 rows affected (1.68 sec)

mysql> create table student(Id int,Name varchar(20),Address varchar(20)) ENGINE=MYISAM;
Query OK, 0 rows affected (0.22 sec)


mysql> insert into student values(1,'sreenath','nellore');
Query OK, 1 row affected (0.12 sec)

mysql> select * from student;
+------+----------+---------+
| Id   | Name     | Address |
+------+----------+---------+
|    1 | sreenath | nellore |
+------+----------+---------+
1 row in set (0.05 sec)

mysql> insert into student values(2,'sree','nellore');
Query OK, 1 row affected (0.05 sec)

mysql> insert into student values(3,'praveen','podalakure');
Query OK, 1 row affected (0.03 sec)

mysql> insert into student values(4,'yash','podalakure');
Query OK, 1 row affected (0.06 sec)

mysql> insert into student values(5,'kumarr','badhvale');
Query OK, 1 row affected (0.05 sec)

mysql> insert into student values(6,'avir','bad');
Query OK, 1 row affected (0.06 sec)

mysql> insert into student values(7,'vdfvgb','ghodfv');
Query OK, 1 row affected (0.03 sec)

mysql> insert into student values(8,'rthhy','gbgrh');
Query OK, 1 row affected (0.05 sec)

mysql> insert into student values(9,'raju','home');
Query OK, 1 row affected (0.08 sec)

mysql> insert into student values(10,'raju','hoween');
Query OK, 1 row affected (0.04 sec)

mysql> insert into student values(11,'ergb','fgtefv');
Query OK, 1 row affected (0.10 sec)

mysql> insert into student values(12,'egwfrb','fwsa');
Query OK, 1 row affected (0.10 sec)

mysql> insert into student values(13,'qwdw','wfra');
Query OK, 1 row affected (0.22 sec)

mysql> insert into student values(14,'awaf','wavr');
Query OK, 1 row affected (0.07 sec)

mysql> insert into student values(15,'rajasekhar','tirupati');
Query OK, 1 row affected (0.05 sec)

mysql> select * from student;
+------+------------+------------+
| Id   | Name       | Address    |
+------+------------+------------+
|    1 | sreenath   | nellore    |
|    2 | sree       | nellore    |
|    3 | praveen    | podalakure |
|    4 | yash       | podalakure |
|    5 | kumarr     | badhvale   |
|    6 | avir       | bad        |
|    7 | vdfvgb     | ghodfv     |
|    8 | rthhy      | gbgrh      |
|    9 | raju       | home       |
|   10 | raju       | hoween     |
|   11 | ergb       | fgtefv     |
|   12 | egwfrb     | fwsa       |
|   13 | qwdw       | wfra       |
|   14 | awaf       | wavr       |
|   15 | rajasekhar | tirupati   |
+------+------------+------------+
15 rows in set (0.00 sec)

mysql> show table status\G
*************************** 1. row ***************************
           Name: employee
         Engine: InnoDB
        Version: 10
     Row_format: Dynamic
           Rows: 3
 Avg_row_length: 5461
    Data_length: 16384
Max_data_length: 0
   Index_length: 0
      Data_free: 0
 Auto_increment: NULL
    Create_time: 2021-02-24 17:42:28
    Update_time: 2021-02-24 17:44:56
     Check_time: NULL
      Collation: utf8mb4_0900_ai_ci
       Checksum: NULL
 Create_options:
        Comment:
*************************** 2. row ***************************
           Name: student
         Engine: MyISAM
        Version: 10
     Row_format: Dynamic
           Rows: 15
 Avg_row_length: 24
    Data_length: 368
Max_data_length: 281474976710655
   Index_length: 1024
      Data_free: 0
 Auto_increment: 1
    Create_time: 2021-02-25 11:19:31
    Update_time: 2021-02-25 11:26:54
     Check_time: NULL
      Collation: utf8mb4_0900_ai_ci
       Checksum: NULL
 Create_options:
        Comment:
2 rows in set (0.15 sec)
change innodb engine table into ,myiasm engine storage
------------------------------------------------------
mysql> alter table Employee engine=MYISAM;
Query OK, 3 rows affected (2.12 sec)
Records: 3  Duplicates: 0  Warnings: 0

mysql> show table status\G 
*************************** 1. row ***************************
           Name: employee
         Engine: MyISAM
        Version: 10
     Row_format: Dynamic
           Rows: 3
 Avg_row_length: 5461
    Data_length: 16384
Max_data_length: 0
   Index_length: 0
      Data_free: 0
 Auto_increment: NULL
    Create_time: 2021-02-25 11:49:02
    Update_time: 2021-02-24 17:44:56
     Check_time: NULL
      Collation: utf8mb4_0900_ai_ci
       Checksum: NULL
 Create_options:
        Comment:
*************************** 2. row ***************************
           Name: student
         Engine: MyISAM
        Version: 10
     Row_format: Dynamic
           Rows: 15
 Avg_row_length: 24
    Data_length: 368
Max_data_length: 281474976710655
   Index_length: 1024
      Data_free: 0
 Auto_increment: 1
    Create_time: 2021-02-25 11:19:31
    Update_time: 2021-02-25 11:26:54
     Check_time: NULL
      Collation: utf8mb4_0900_ai_ci
       Checksum: NULL
 Create_options:
        Comment:
2 rows in set (0.00 sec)
mysql> alter table student engine=innodb
    -> ;
Query OK, 15 rows affected (1.71 sec)
Records: 15  Duplicates: 0  Warnings: 0

mysql> show table status\G
*************************** 1. row ***************************
           Name: employee
         Engine: MyISAM
        Version: 10
     Row_format: Dynamic
           Rows: 3
 Avg_row_length: 5461
    Data_length: 16384
Max_data_length: 0
   Index_length: 0
      Data_free: 0
 Auto_increment: NULL
    Create_time: 2021-02-25 11:49:02
    Update_time: 2021-02-24 17:44:56
     Check_time: NULL
      Collation: utf8mb4_0900_ai_ci
       Checksum: NULL
 Create_options:
        Comment:
*************************** 2. row ***************************
           Name: student
         Engine: InnoDB
        Version: 10
     Row_format: Dynamic
           Rows: 15
 Avg_row_length: 24
    Data_length: 368
Max_data_length: 281474976710655
   Index_length: 1024
      Data_free: 0
 Auto_increment: 1
    Create_time: 2021-02-25 11:52:26
    Update_time: 2021-02-25 11:26:54
     Check_time: NULL
      Collation: utf8mb4_0900_ai_ci
       Checksum: NULL
 Create_options:
        Comment:
2 rows in set (0.07 sec)
mysql> select * from student;
+------+------------+------------+
| Id   | Name       | Address    |
+------+------------+------------+
|    1 | sreenath   | nellore    |
|    2 | sree       | nellore    |
|    3 | praveen    | podalakure |
|    4 | yash       | podalakure |
|    5 | kumarr     | badhvale   |
|    6 | avir       | bad        |
|    7 | vdfvgb     | ghodfv     |
|    8 | rthhy      | gbgrh      |
|    9 | raju       | home       |
|   10 | raju       | hoween     |
|   11 | ergb       | fgtefv     |
|   12 | egwfrb     | fwsa       |
|   13 | qwdw       | wfra       |
|   14 | awaf       | wavr       |
|   15 | rajasekhar | tirupati   |
+------+------------+------------+
15 rows in set (0.00 sec)
