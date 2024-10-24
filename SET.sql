CREATE DATABASE api_fiber;

psql -U postgres -d api_fiber
--
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
SELECT uuid_generate_v4();


--
INSERT INTO public.users (id, name, user_name, email, password, created_at, updated_at)
VALUES (
           uuid_generate_v4(),
           'jamal',
           'jamal',
           'jamal@example.com',
           '$2a$12$rzraOu3SA74DtdBSZrRfR.gcWd5yNTomvo571IvXDnhe2mAOVbpa.',
           NOW(),
           NOW()
       );
plain text 123