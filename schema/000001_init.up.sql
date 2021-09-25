CREATE TABLE mocks
(
    id serial not null unique,
    name varchar(255) not null,
    description text,
    status int not null,
    content_type varchar(255) not null,
    request_method varchar(255) not null,
    route_path varchar(255) not null,
    body_type varchar(255),
    body_content text,
    script_type varchar(255),
    script text,
    active boolean not null,
    mock_order int not null
);

CREATE TABLE headers
(
    id serial not null unique,
    name varchar(255) not null,
    value varchar(255) not null,
    mock_id int references mocks (id) on delete cascade not null
);
