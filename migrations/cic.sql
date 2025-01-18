
-- Table for english
CREATE TABLE english
(
    id       BIGSERIAL PRIMARY KEY,
    word     VARCHAR(255) NOT NULL,
    phonetic VARCHAR(255) DEFAULT '',
    audio    TEXT         DEFAULT ''
);

-- Table for reuse_account_ids
CREATE TABLE reuse_account_ids
(
    id BIGSERIAL PRIMARY KEY
);

-- Table for reuse_english_ids
CREATE TABLE reuse_english_ids
(
    id BIGSERIAL PRIMARY KEY
);

-- Table for reuse_topics_ids
CREATE TABLE reuse_topics_ids
(
    id BIGSERIAL PRIMARY KEY
);

-- Table for reuse_user_ids
CREATE TABLE reuse_user_ids
(
    id INT PRIMARY KEY
);

-- Table for reuse_vietnamese_ids
CREATE TABLE reuse_vietnamese_ids
(
    id BIGSERIAL PRIMARY KEY
);

-- Table for reuse_vocabularies_ids
CREATE TABLE reuse_vocabularies_ids
(
    id BIGSERIAL PRIMARY KEY
);

-- Table for topics
CREATE TABLE topics
(
    id      BIGSERIAL PRIMARY KEY,
    name    VARCHAR(255) NOT NULL,
    of_user INT          NOT NULL DEFAULT 2
);

-- Table for users
CREATE TABLE users
(
    id          SERIAL PRIMARY KEY,
    email       VARCHAR(250) NOT NULL,
    img         VARCHAR(255) NOT NULL,
    theme       INT          NOT NULL DEFAULT 0,
    user_name   VARCHAR(100),
    password    VARCHAR(64),
    pin         VARCHAR(64)  NOT NULL,
    pin_wrong   INT          NOT NULL DEFAULT 0,
    token       VARCHAR(128),
    update_time TIMESTAMP    NOT NULL
);

-- Table for vietnamese
CREATE TABLE vietnamese
(
    id      BIGSERIAL PRIMARY KEY,
    signify VARCHAR(255) NOT NULL
);

-- Table for vocabularies
CREATE TABLE vocabularies
(
    id              BIGSERIAL PRIMARY KEY,
    en              BIGINT    NOT NULL,
    part_of_speech  VARCHAR(25),
    img             TEXT               DEFAULT '',
    vi              BIGINT    NOT NULL,
    user_own        INT       NOT NULL DEFAULT 2,
    appear_en_vi    INT       NOT NULL DEFAULT 0,
    appear_vi_en    INT       NOT NULL DEFAULT 0,
    checking_en     BOOLEAN   NOT NULL DEFAULT TRUE,
    checking_vi     BOOLEAN   NOT NULL DEFAULT TRUE,
    checking_listen BOOLEAN   NOT NULL DEFAULT TRUE,
    checking_speak  BOOLEAN   NOT NULL DEFAULT TRUE,
    create_date     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    learning_day    TIMESTAMP,
    level           INT       NOT NULL DEFAULT 0
);

-- Table for vocabularies_topics
CREATE TABLE vocabularies_topics
(
    topic      BIGINT NOT NULL,
    vocabulary BIGINT NOT NULL,
    PRIMARY KEY (vocabulary, topic)
);

-- Constraints for the tables

-- Foreign key for topics
ALTER TABLE topics
    ADD CONSTRAINT fk_topics_users FOREIGN KEY (of_user) REFERENCES users (id) ON UPDATE CASCADE;

-- Foreign keys for vocabularies
ALTER TABLE vocabularies
    ADD CONSTRAINT fk_vocabularies_english FOREIGN KEY (en) REFERENCES english (id) ON UPDATE CASCADE,
    ADD CONSTRAINT fk_vocabularies_users FOREIGN KEY (user_own) REFERENCES users (id) ON UPDATE CASCADE,
                                                                                                                                                                                                  ADD CONSTRAINT fk_vocabularies_vietnamese FOREIGN KEY (vi) REFERENCES vietnamese (id) ON UPDATE CASCADE;

-- Foreign keys for vocabularies_topics
ALTER TABLE vocabularies_topics
    ADD CONSTRAINT fk_vt_topics FOREIGN KEY (topic) REFERENCES topics (id) ON UPDATE CASCADE,
    ADD CONSTRAINT fk_vt_vocabularies FOREIGN KEY (vocabulary) REFERENCES vocabularies (id) ON UPDATE CASCADE;

INSERT INTO users (id, email, img, theme, user_name, password, pin, pin_wrong, token, update_time)
VALUES (2, 'nguyentandat.mail@gmail.com', '/views/assets/img/users/2.png', 1, 'Nguyễn Đạt',
        'd5f60fe3a2d39b5140ed55697062b7b13089d76c768a6a2e50d818eefba3b1e6',
        'a6adec41a464d024020ce23f745fae62c4050499b0176a9fbc888c25346a1826', 0,
        'p889I8i2TDc4OAzbmKWMHEW4JgaH5Z0WYqxSUUZcjSCwnevKD2vgkFq8bOfKzRkczoyE4SU7VvoCBmNHLHbQT4Z56c6kZx3cY9OONzMXscRiA0FIK4H42m9VaXJO5Ant',
        '2024-12-12 01:44:36');
INSERT INTO accounts (id, user_name, password, allocate, of_user)
VALUES (4, 'nguyentandat16052000', 'xyUJZwbM8L46MbjDfz1py5xFp7QxIC6ZFUONvo1u8iw=', 'Root email', 2),
       (7, 'nguyentandat16052000.cd', 'i9ElnftTtKX35XRQxjnCbPzneQQiqwF5hEImb/7Y8ok=', 'Branch of Root mail', 2),
       (8, 'tandatnguyen16052000', 'Y6RUIejFxoFQTAwzpijnhmxxvAqO8Rf/ahsX+xLGhqc=', 'Branch of Root mail', 2),
       (9, 'vi.nguyentandat', 'f2F1ZqaHtfP7vdAUAcOEHlTHnDZ0tewiS8zTjj7bbV8=', 'Branch of Root mail', 2),
       (10, 'ntandat.vi', 'KXLzPhHDJxbPPNvhedIVgTFcVKPCKEtpygPvTTAh1TU=', 'Branch of Root mail', 2),
       (11, 'trendingshopsaigon', 'Bi35MAMnpL9xUESc/3W3kCqbZZO+NLD3qZa/NzKoNQw=', 'Branch of Root mail', 2),
       (17, 'nguyentandat.use', '1+x36gJTNdgCAOMZL72wE1CFyYOWGTEdMTK1j1ojIuI=', 'Branch of Root mail use', 2),
       (18, 'nguyentandat.on', '+TDOmPF29eJHmFA+HOVMx6cIa4p87gkLASxXVAPUoJs=', 'Branch of Root mail use', 2),
       (19, 'nguyentandat.mail', 'KzpSZ/jfN550excby7QbzrOA1b6bqc1KKqRw9BxCu5A=', 'Root mail use', 2),
       (20, 'nguyentandat.dv', 'IW0M1YVDfGmv5AURw8+5CYZ12+1s2K4/IjhroeSgUu7D4kBEutrmggc/MNy4GY2P', 'Root mail dv', 2),
       (21, 'nguyentandat.dev.vn', 'lH/mymwLctcqbYL41kzvuzbhZLL9rEkudekoHqnZlgJQZm/XS5unxFrvxm9wOvld',
        'Branch of Root mail dv', 2),
       (22, 'dat.dev.vn', 'dTVb9UNnUelFHqnejmH9TbPbq7GEUtgSFsSlTBBOd74IHGFa9Nu9O/SXm2WUHkyG', 'Branch of Root mail dv',
        2),
       (23, 'ntdat.dv', 'JjU6pVfVAvqDaX8mO9+61mvHb1ZeVIHQ0wQfN838SY4czdnofioy4tllxyrFOOWp', 'Branch of Root mail dv',
        2),
       (24, 'nguyentandat.email.dev', 'esxELoBXLnMORodbVwL+f9ECQQp7t8sz9AKjRTljLUA=', 'Branch of Root mail dv', 2),
       (26, 'cic.dat.dev@gmail.com', 'o2TutiLT0yYoayFpgK6k9/vVL6YKSxQEOLXVE/15ZCk=', 'Branch of Root mail', 2);
