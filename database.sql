create database IT_Ticketing_Platform;

create type ticket_status as enum('Open', 'Closed');

create table Ticket(
    Ticket_ID uuid primary key,
    Status ticket_status not null,
    Category_ID uuid,
    Subcategory_ID uuid,
    Title varchar(100) not null,
    Description varchar(5000),
    Customer_Contact varchar(100) not null,
    Opened_Data timestamp default(current_timestamp) not null,
    Updated_At timestamp default(current_timestamp) not null,
    Closed_Date timestamp,
    Assigned_Team uuid,
    Assigned_Analyst uuid,
    Opened_By uuid,
    Closed_By uuid
);

create table Category(
    Category_ID uuid primary key,
    Category_Name varchar(100) not null unique
);

create table Subcategory(
    Subcategory_ID uuid primary key,
    Category_ID uuid not null references Category(Category_ID) on delete cascade,
    Subcategory_Name varchar(100) not null
);

alter table Ticket
add constraint fk_ticket_category foreign key (Category_ID) references Category(Category_ID) on delete set null,
add constraint fk_ticket_subcategory foreign key (Subcategory_ID) references Subcategory(Subcategory_ID) on delete set null;

create table Analyst(
    Analyst_ID uuid primary key,
    First_Name varchar(50) not null,
    Last_Name varchar(50) not null,
    Email varchar(100) not null unique,
    Password varchar not null,
    Phone_Number varchar(20) not null unique,
    Team_ID uuid,
    Number_of_Open_Tickets int not null default(0),
    Number_of_Opened_Tickets int not null default(0),
    Number_of_Closed_Tickets int not null default(0)
);

create table Manager(
    Manager_ID uuid primary key references Analyst(Analyst_ID) on delete cascade
);

create table Administrator(
    Administrator_ID uuid primary key references Analyst(Analyst_ID) on delete cascade
);

create table Team(
    Team_ID uuid primary key,
    Team_Name varchar(100) not null unique,
    Manager_ID uuid references Manager(Manager_ID) on delete set null
);

alter table Ticket
add constraint fk_ticket_assigned_analyst foreign key (Assigned_Analyst) references Analyst(Analyst_ID) on delete set null,
add constraint fk_ticket_assigned_team foreign key (Assigned_Team) references Team(Team_ID) on delete set null,
add constraint fk_ticket_opened_by foreign key (Opened_By) references Analyst(Analyst_ID) on delete set null,
add constraint fk_ticket_closed_by foreign key (Closed_By) references Analyst(Analyst_ID) on delete set null;

alter table Analyst
add constraint fk_analyst_team foreign key (Team_ID) references Team(Team_ID) on delete set null;

create table Ticket_Reopen(
    Reopen_ID uuid primary key,
    Ticket_ID uuid not null references Ticket(Ticket_ID) on delete cascade,
    Reopened_By uuid references Analyst(Analyst_ID) on delete set null,
    Reopened_Date timestamp not null default(current_timestamp)
);

create table Assignment_History(
    Assignment_ID uuid primary key,
    Ticket_ID uuid not null references Ticket(Ticket_ID) on delete cascade,
    Assigned_From uuid references Analyst(Analyst_ID) on delete set null,
    Assigned_to_Analyst uuid references Analyst(Analyst_ID) on delete set null,
    Assigned_to_Team uuid references Team(Team_ID) on delete set null,
    Assignment_Message varchar(500),
    Assignment_Date timestamp not null default(current_timestamp)
);

insert into Team values(gen_random_uuid(), 'IT Support', null);
insert into Team values(gen_random_uuid(), 'Hardware', null);
insert into Team values(gen_random_uuid(), 'Software', null);
insert into Team values(gen_random_uuid(), 'Utilities', null);

--select (password = crypt('vladan123', password)) as password from analyst; RETURNS TRUE IF password is correct, OTHERWISE FALSE
--Adminstrator
CREATE EXTENSION IF NOT EXISTS pgcrypto;
insert into Analyst(Analyst_ID, First_Name, Last_Name, Email, Password, Phone_Number, Team_ID) values(gen_random_uuid(), 'Vladan', 'Tesic', 'vladan@gmail.com', crypt('vladan123', gen_salt('md5')), '123456789', (select Team_ID from Team where Team_Name = 'Administrators'));
insert into Administrator values((select Analyst_ID from Analyst where Email = 'vladan@gmail.com'));

--IT Support Manager
insert into Analyst(Analyst_ID, First_Name, Last_Name, Email, Password, Phone_Number, Team_ID) values(gen_random_uuid(), 'Marko', 'Markovic', 'marko@gmail.com', crypt('marko123', gen_salt('md5')), '987654321', (select Team_ID from Team where Team_Name = 'IT Support'));
insert into Manager values((select Analyst_ID from Analyst where Email = 'marko@gmail.com'));
update Team set Manager_ID = (select Analyst_ID from Analyst where Email = 'marko@gmail.com') where Team_Name = 'IT Support';

--Hardware Manager
insert into Analyst(Analyst_ID, First_Name, Last_Name, Email, Password, Phone_Number, Team_ID) values(gen_random_uuid(), 'Sara', 'Saric', 'sara@gmail.com', crypt('sara123', gen_salt('md5')), '123654789', (select Team_ID from Team where Team_Name = 'Hardware'));
insert into Manager values((select Analyst_ID from Analyst where Email = 'sara@gmail.com'));
update Team set Manager_ID = (select Analyst_ID from Analyst where Email = 'sara@gmail.com') where Team_Name = 'Hardware';

--Software Manager
insert into Analyst(Analyst_ID, First_Name, Last_Name, Email, Password, Phone_Number, Team_ID) values(gen_random_uuid(), 'Ana', 'Anic', 'ana@gmail.com', crypt('ana123', gen_salt('md5')), '123456987', (select Team_ID from Team where Team_Name = 'Software'));
insert into Manager values((select Analyst_ID from Analyst where Email = 'ana@gmail.com'));
update Team set Manager_ID = (select Analyst_ID from Analyst where Email = 'ana@gmail.com') where Team_Name = 'Software';

--Utilities Manager
insert into Analyst(Analyst_ID, First_Name, Last_Name, Email, Password, Phone_Number, Team_ID) values(gen_random_uuid(), 'Petar', 'Petrovic', 'petar@gmail.com', crypt('petar123', gen_salt('md5')), '321456789', (select Team_ID from Team where Team_Name = 'Utilities'));
insert into Manager values((select Analyst_ID from Analyst where Email = 'petar@gmail.com'));
update Team set Manager_ID = (select Analyst_ID from Analyst where Email = 'petar@gmail.com') where Team_Name = 'Utilities';

--IT Support Analysts
insert into Analyst(Analyst_ID, First_Name, Last_Name, Email, Password, Phone_Number, Team_ID) values(gen_random_uuid(), 'Janko', 'Jankovic', 'janko@gmail.com', crypt('janko123', gen_salt('md5')), '321654789', (select Team_ID from Team where Team_Name = 'IT Support'));
insert into Analyst(Analyst_ID, First_Name, Last_Name, Email, Password, Phone_Number, Team_ID) values(gen_random_uuid(), 'Katarina', 'Katic', 'katarina@gmail.com', crypt('katarina123', gen_salt('md5')), '321654987', (select Team_ID from Team where Team_Name = 'IT Support'));

--Hardware Analysts
insert into Analyst(Analyst_ID, First_Name, Last_Name, Email, Password, Phone_Number, Team_ID) values(gen_random_uuid(), 'Nikola', 'Nikolic', 'nikola@gmail.com', crypt('nikola123', gen_salt('md5')), '987456123', (select Team_ID from Team where Team_Name = 'Hardware'));
insert into Analyst(Analyst_ID, First_Name, Last_Name, Email, Password, Phone_Number, Team_ID) values(gen_random_uuid(), 'Branko', 'Brankovic', 'branko@gmail.com', crypt('branko123', gen_salt('md5')), '987654123', (select Team_ID from Team where Team_Name = 'Hardware'));

--Software Analysts
insert into Analyst(Analyst_ID, First_Name, Last_Name, Email, Password, Phone_Number, Team_ID) values(gen_random_uuid(), 'Jana', 'Jankovic', 'jana@gmail.com', crypt('jana123', gen_salt('md5')), '789654123', (select Team_ID from Team where Team_Name = 'Software'));
insert into Analyst(Analyst_ID, First_Name, Last_Name, Email, Password, Phone_Number, Team_ID) values(gen_random_uuid(), 'Boris', 'Borisic', 'boris@gmail.com', crypt('boris123', gen_salt('md5')), '789456321', (select Team_ID from Team where Team_Name = 'Software'));

--Utilites Analysts
insert into Analyst(Analyst_ID, First_Name, Last_Name, Email, Password, Phone_Number, Team_ID) values(gen_random_uuid(), 'Julia', 'Julic', 'julia@gmail.com', crypt('julia123', gen_salt('md5')), '123789456', (select Team_ID from Team where Team_Name = 'Utilites'));
insert into Analyst(Analyst_ID, First_Name, Last_Name, Email, Password, Phone_Number, Team_ID) values(gen_random_uuid(), 'Aleksandra', 'Aleksic', 'aleksandra@gmail.com', crypt('aleksandra123', gen_salt('md5')), '321987654', (select Team_ID from Team where Team_Name = 'Utilites'));

insert into Category values(gen_random_uuid(), 'Account issues');
insert into Category values(gen_random_uuid(), 'Software install');
insert into Category values(gen_random_uuid(), 'Software removal');
insert into Category values(gen_random_uuid(), 'Software issues');
insert into Category values(gen_random_uuid(), 'Laptop');
insert into Category values(gen_random_uuid(), 'Desktop');
insert into Category values(gen_random_uuid(), 'Hardware request');
insert into Category values(gen_random_uuid(), 'Utilities');
insert into Category values(gen_random_uuid(), 'Other');

insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Account issues'), 'Forgotten email');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Account issues'), 'Email change');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Account issues'), 'Forgotten password');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Account issues'), 'Password change');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Account issues'), 'Phone number change');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Account issues'), 'Other');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Software install'), 'Microsoft Office');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Software install'), 'Adobe');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Software install'), 'IDE');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Software install'), 'Internet browser');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Software install'), 'Other');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Software removal'), 'Microsoft Office');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Software removal'), 'Adobe');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Software removal'), 'IDE');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Software removal'), 'Internet browser');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Software removal'), 'Other');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Software issues'), 'Microsoft Office');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Software issues'), 'Adobe');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Software issues'), 'IDE');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Software issues'), 'Internet browser');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Software issues'), 'Other');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Laptop'), 'Monitor');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Laptop'), 'Power on/off');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Laptop'), 'Speed issues');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Laptop'), 'Memory full');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Laptop'), 'Memory extension');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Laptop'), 'Other');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Desktop'), 'Monitor');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Desktop'), 'Power on/off');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Desktop'), 'Speed issues');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Desktop'), 'Memory full');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Desktop'), 'Memory extension');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Desktop'), 'Other');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Hardware request'), 'Monitor');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Hardware request'), 'Keyboard');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Hardware request'), 'Mouse');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Hardware request'), 'Headphones');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Hardware request'), 'Laptop');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Hardware request'), 'Desktop');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Hardware request'), 'Other');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Utilities'), 'Vending machine');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Utilities'), 'Water dispenser');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Utilities'), 'Printer');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Utilities'), 'Fridge');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Utilities'), 'Toilet');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Utilities'), 'Sink');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Utilities'), 'Other');
insert into Subcategory values(gen_random_uuid(), (select Category_ID from Category where Category_Name = 'Other'), 'Other');