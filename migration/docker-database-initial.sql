create table users(
    id serial primary key,
    name varchar(50) not null,
    email varchar(60) not null unique,
    password varchar(70) not null
);

INSERT INTO users(name, email, password) VALUES
('Luan', 'lluantenorio7@gmail.com', '$2a$12$7LVHCrakFO6wFtM23uvxQOj6pVsEm7bPPAwL2kn2rFuvgvgVDFSzG'),
('Cleiton', 'cleiton@gmail.com', '$2a$12$DRjW1RnB0z4Rt5lF2oJPFeAkgxrw9mZ5YuctvMYyLGE5/YcLVUvsC');
