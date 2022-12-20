CREATE TABLE users (
  id uuid NOT NULL,
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL,
  name character varying NOT NULL,
  email character varying NOT NULL,
  PRIMARY KEY (id)
);
