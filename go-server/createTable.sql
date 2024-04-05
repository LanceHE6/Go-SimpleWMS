CREATE TABLE IF NOT EXISTS user (
                                    uid VARCHAR(20) primary key not null ,
                                    account VARCHAR(255) unique not null,
                                    password VARCHAR(255),
                                    nickname VARCHAR(255),
                                    permission INT default 1,
                                    register_time VARCHAR(255),
                                    token VARCHAR(255) default '',
                                    phone VARCHAR(255) default ''
);
CREATE TABLE IF NOT EXISTS goods_type (
                                          gtid VARCHAR(20) primary key not null,
                                          name VARCHAR(255),
                                          type_code VARCHAR(255) default '',
                                          add_time VARCHAR(255)
);
CREATE TABLE IF NOT EXISTS department (
                                          did VARCHAR(20) primary key not null,
                                          name VARCHAR(255),
                                          add_time VARCHAR(255)
);
CREATE TABLE IF NOT EXISTS staff (
                                     sid VARCHAR(20) primary key not null,
                                     name VARCHAR(255),
                                     phone VARCHAR(255) default '',
                                     department VARCHAR(255) default '',
                                     add_time VARCHAR(255),
                                     foreign key (department) references department(did) on delete set null
);
CREATE TABLE IF NOT EXISTS warehouse (
                                         wid VARCHAR(20) primary key not null,
                                         name VARCHAR(255),
                                         add_time VARCHAR(255),
                                         comment VARCHAR(255) default '',
                                         manager VARCHAR(255),
                                         status INT default 1,
                                         foreign key (manager) references staff(sid) on delete set null
);
CREATE TABLE IF NOT EXISTS inventory_type (
                                         itid VARCHAR(20) primary key not null,
                                         name VARCHAR(255),
                                         type_code VARCHAR(255) default '',
                                         add_time VARCHAR(255)
);
CREATE TABLE IF NOT EXISTS unit (
                                              unid varchar(20) primary key not null,
                                              name VARCHAR(255),
                                              add_time VARCHAR(255)
);