-- Table: english
CREATE TABLE english
(
    id       BIGSERIAL PRIMARY KEY,
    word     VARCHAR(255) NOT NULL UNIQUE, -- UNIQUE để tránh từ vựng trùng lặp
    phonetic VARCHAR(255) DEFAULT '',
    audio    TEXT         DEFAULT ''
);

-- Table: topics
CREATE TABLE topics
(
    id      BIGSERIAL PRIMARY KEY,
    name    VARCHAR(255) NOT NULL,
    user_id VARCHAR(128) NOT NULL
);

-- Table: vietnamese
CREATE TABLE vietnamese
(
    id      BIGSERIAL PRIMARY KEY,
    signify VARCHAR(255) NOT NULL UNIQUE -- Đảm bảo nghĩa tiếng Việt không bị trùng
);

-- Table: vocabularies
CREATE TABLE vocabularies
(
    id              BIGSERIAL PRIMARY KEY,
    en              BIGINT       NOT NULL, -- Quan hệ tới bảng english
    vi              BIGINT       NOT NULL, -- Quan hệ tới bảng vietnamese
    user_id         VARCHAR(128) NOT NULL, -- Tham chiếu tới User ID từ service authentication
    part_of_speech  VARCHAR(25),           -- Loại từ (danh từ, động từ, etc.)
    img             TEXT                  DEFAULT '',
    appear_en_vi    INT          NOT NULL DEFAULT 0,
    appear_vi_en    INT          NOT NULL DEFAULT 0,
    checking_en     BOOLEAN      NOT NULL DEFAULT TRUE,
    checking_vi     BOOLEAN      NOT NULL DEFAULT TRUE,
    checking_listen BOOLEAN      NOT NULL DEFAULT TRUE,
    checking_speak  BOOLEAN      NOT NULL DEFAULT TRUE,
    create_date     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    learning_day    TIMESTAMP,
    level           INT          NOT NULL DEFAULT 0,
    CONSTRAINT fk_vocabularies_english FOREIGN KEY (en) REFERENCES english (id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_vocabularies_vietnamese FOREIGN KEY (vi) REFERENCES vietnamese (id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- Table: vocabularies_topics (many-to-many)
CREATE TABLE vocabularies_topics
(
    topic      BIGINT NOT NULL,
    vocabulary BIGINT NOT NULL,
    PRIMARY KEY (vocabulary, topic),
    CONSTRAINT fk_vt_topics FOREIGN KEY (topic) REFERENCES topics (id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_vt_vocabularies FOREIGN KEY (vocabulary) REFERENCES vocabularies (id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- Indexes for performance
CREATE INDEX idx_english_word ON english (word);
CREATE INDEX idx_vietnamese_signify ON vietnamese (signify);
CREATE INDEX idx_topics_user_id ON topics (user_id);
CREATE INDEX idx_vocabularies_user_id ON vocabularies (user_id);
CREATE INDEX idx_vocabularies_en ON vocabularies (en);
CREATE INDEX idx_vocabularies_vi ON vocabularies (vi);
