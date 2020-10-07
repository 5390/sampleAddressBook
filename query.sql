CREATE TABLE address_book(
id serial PRIMARY KEY,
first_name iVARCHAR(50) NOT NULL,
last_name iVARCHAR(50) NOT NULL,
address_line_1 iVARCHAR(255) NOT NULL,
address_line_1 VARCHAR(255) NOT NULL,
state VARCHAR(255) NOT NULL,
zipcode VARCHAR(255) NOT NULL
);


INSERT INTO public.address_book(
	first_name, last_name, address_line_1, address_line_2, state, zipcode)
	VALUES('Dave','Smith','123 main st.','seattle','wa','43'),
('Alice','Smith','123 Main St.','Seattle','WA','45'),
('bOb','Williams','234 2nd Ave.','Tacoma','WA','26'),
('Carol','Johnson','234 2nd Ave','Seattle','WA','67'),
('TOm','Bombadillo','1212 Maple Street','Florida','GA','520'),
('Jimbo','Jones','82 Pine Street','Atlanta','GA','2'),
('Jackie','Jones','82 Pine Street','Atlanta','GA','6'),
('Tommy','Jones','82 Pine Street','Atlanta','GA','29'),
('tammy','Jones','82 Pine Street','Atlanta','GA','27'),
('EvE','Smith','234 2nd Ave.','Tacoma','WA','25'),
('Frank','Jones','234 2nd Ave.','Tacoma','FL','23'),
('george','Brown','345 3rd Blvd., Apt. 200','Seattle','WA','19'),
('Helen','Brown','345 3rd Blvd. Apt. 200','Seattle','WA','18'),
('Ian','smith','123 main st ','Seattle','Wa','18'),
('Jane','Smith','123 Main St.','Seattle','WA','13');