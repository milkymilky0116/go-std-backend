INSERT INTO snippets (title, content, created, expires)
VALUES (
        ?,
        ?,
        UTC_TIMESTAMP(),
        DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY)
    )