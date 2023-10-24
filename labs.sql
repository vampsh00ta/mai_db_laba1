/*проишлозло дтп*/
begin;
    insert into dtp (date,area,street,coords,category,metro)  values(now(),'string','string','string, string','string','string') returning id;
    select c.id  from crew as c join  gai on gai.id = c.gai_id
                 where c.duty = true
                   and c.id not  in (select crew_id from crew_dtp)
                   and  gai.metro = ? limit 1;
    insert into  dtp_description(text, time,dtp_id) values('были вызваны сотрудники дтп',now(),?);
    insert into crew_dtp (dtp_id,crew_id) values(?,?);
commit;

/*какие та дополнения по дтп*/
begin;
    insert into  dtp_description(text, time,dtp_id) values('сотрудники приехали',now(),?);
commit;

/*найти номер машины по pts*/
begin;
    select * from vehicle where pts = ?;
commit;

/*проверить принадлежит ли машина человеку по номеру*/
begin;
    select v.id from (select * from vehicle where pts = '1488' ) as v
    join person_vehicle on person_vehicle.vehicle_id = v.id
    join (select  *  from person where name  = 'ya' and surname  = 'ya' and  patronymic = 'ya' and passport = 11111) as p on  p.id = person_vehicle.person_id
commit;

/*машины задействованные в дтп*/
begin;
    select vehicle.model ,vehicle.pts from (select * from dtp where id = 1 ) as d
    join participant_of_dtp on  participant_of_dtp.dtp_id = d.id
    join person on person.id = participant_of_dtp.person_id
    join person_vehicle on person_vehicle.person_id = person.id
    join  vehicle on person_vehicle.vehicle_id = vehicle.id;
commit;
/*добавить машину*/
begin;
insert into vehicle ( pts, model, category)  values('asd','asd','asd') returning id;
commit;
/*добавить пользователя*/
begin;
insert into person (name, surname, patronymic, birthday, passport, citizenship)
values ('asd','asd','asd',now(),1111,'asd')
returning id;
commit;

/*добавить сотрудника дпс*/
begin;
insert into police_officer (rank, person_id)
values ('asd',1)
    returning id;
commit;
/*добавить взвод*/
begin;
    insert into crew (p_officer_id_1,p_officer_id_2,gai_id,time,duty)
    values (1,2,1,now(),false)
commit;
/*добавить гаи*/
begin;
insert into gai (area, metro)
values ('a','a')
returning id;
    commit;
/*добавить участника дтп*/
begin;
    select id from violation  where law_number = ?;
    select  id  from person where name  = ? and surname  = ? and  patronymic = ? and passport = ?;
    insert into participant_of_dtp (violation_id,vehicle_id , person_id , dtp_id , role)
    values (?,?,?,?);
commit;

/*
-------------------------
ТАБЛИЦЫ
-------------------------
*/

begin;


CREATE TABLE person (
    id serial primary key,
     name varchar(255),
     surname varchar(255),
     patronymic varchar(255)  NULL,
     birthday timestamp ,
     passport integer,
     citizenship varchar(255)
    );

CREATE TABLE vehicle (
         id serial primary key,
         pts varchar(255),
         model varchar(255),
         category varchar(255)
     );


CREATE TABLE person_vehicle (
     vehicle_id integer    REFERENCES vehicle (id),
     person_id integer    REFERENCES person (id)

            );



CREATE TABLE violation (
    id serial primary key,
    law varchar(255),
    law_number varchar(255)
   );



CREATE TABLE police_officer (
     id serial primary key,
     rank varchar,
     person_id integer  REFERENCES person (id)
);

create table gai
(
    id    serial primary key,
    area  varchar(255),
    metro varchar(255)

);
CREATE TABLE crew (
      id serial primary key,
      p_officer_id_1 integer  REFERENCES police_officer (id),
      p_officer_id_2 integer  REFERENCES police_officer (id),
      gai_id   integer  REFERENCES gai (id),
      time timestamp  DEFAULT CURRENT_TIMESTAMP,
      duty boolean
                  );


CREATE TABLE dtp
(
    id       serial primary key,
    date     timestamp DEFAULT CURRENT_TIMESTAMP,
    area     varchar(255),
    street   varchar(255),
    metro    varchar(255),
    coords   varchar(255),
    category varchar(255)
);

CREATE TABLE dtp_description (
    id serial primary key,
     text varchar(255),
     time timestamp DEFAULT CURRENT_TIMESTAMP,
     dtp_id integer  REFERENCES dtp (id)
                             );


CREATE TABLE crew_dtp (
    dtp_id integer  REFERENCES dtp (id),
    crew_id integer  REFERENCES crew (id)
                  );

CREATE TABLE crew_po (
    po_id integer  REFERENCES police_officer (id),
    crew_id integer  REFERENCES crew (id)
);
CREATE TABLE participant_of_dtp (
    id serial primary key ,
    violation_id integer REFERENCES violation (id),
    vehicle_id integer  REFERENCES vehicle (id),
    person_id integer  REFERENCES person (id),
    dtp_id integer    REFERENCES dtp (id),
    role varchar(255)
    );



commit;


