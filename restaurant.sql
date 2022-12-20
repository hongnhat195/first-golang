CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE go_delivery.restaurants (
  id  BIGSERIAL NOT NULL ,
  owner_id int NULL,
  name varchar(50) NOT NULL,
  addr varchar(255) NOT NULL,
  city_id int DEFAULT NULL,
  lat decimal DEFAULT NULL,
  lng decimal DEFAULT NULL,
  cover json NULL,
  logo json NULL,
  shipping_fee_per_km decimal DEFAULT '0',
  status int NOT NULL DEFAULT '1',
  created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
  PRIMARY KEY (id)
) ;

CREATE INDEX owner_id ON go_delivery.restaurants USING btree (owner_id);
create  index city_id on go_delivery.restaurants using btree (city_id);

