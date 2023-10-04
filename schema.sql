CREATE TABLE user (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(128) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);
CREATE TABLE gist (
    id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(128) NOT NULL,
    content VARCHAR(500) NOT NULL,
    writer INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`writer`) REFERENCES `user` (`id`) ON UPDATE CASCADE
);
-- INSERT INTO gist (title, content, writer)
-- VALUES (
--         "Hello Mysql speaking.",
--         "Go is sooo amazing",
--         1,
--     ),
--     (
--         "I dont know whether i learn java or not..",
--         "java sucks",
--         1,
--     ),
--     (" hi ", " i want to find study mates.", 2),
--     (" wtf ", " whattt ? ", 2);
-- INSERT INTO user (name)
-- VALUES (" Milky "),
--     (" Stephen ");