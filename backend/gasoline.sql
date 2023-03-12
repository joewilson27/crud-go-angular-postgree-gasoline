-- Table: buku

-- DROP TABLE buku;

CREATE TABLE gasoline
(
 id serial NOT NULL,
 jumlah_liter INT NOT NULL,
 premium INT NOT NULL,
 pertalite INT NOT NULL,
 CONSTRAINT pk_gasoline PRIMARY KEY (id )
)
WITH (
 OIDS=FALSE
);
ALTER TABLE gasoline
 OWNER TO postgres;

-- INSERT DATA INTO TABLE 
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (1,6450,7650);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (2,12900,15300);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (3,19350,22950);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (4,25800,30600);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (5,32250,38250);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (6,38700,45900);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (7,45150,53550);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (8,51600,61200);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (9,58050,68850);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (10,64500,76500);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (11,70950,84150);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (12,77400,91800);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (13,83850,99450);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (14,90300,107100);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (15,96750,107100);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (16,103200,122400);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (17,109650,130050);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (18,116100,137700);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (19,122550,145350);
INSERT INTO gasoline(jumlah_liter, premium, pertalite) VALUES (20,129000,153000);