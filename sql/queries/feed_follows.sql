-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows(id, created_at,updated_at,user_id,feed_id)
    VALUES ($1,$2,$3,$4,$5)
    RETURNING *
)
SELECT inserted_feed_follow.*,u.name AS user_name,f.name AS feed_name
FROM inserted_feed_follow
INNER JOIN users u ON inserted_feed_follow.user_id = u.id
INNER JOIN feeds f ON inserted_feed_follow.feed_id = f.id;

-- name: DeleteFeedFollow :one
DELETE FROM feed_follows
WHERE feed_follows.user_id = $1 AND feed_follows.feed_id = $2
RETURNING *;

-- name: GetFeedFollowsForUser :many
SELECT users.name AS user_name,feeds.name AS feed_name,feed_follows.*
FROM users
INNER JOIN feed_follows on users.id = feed_follows.user_id
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE users.name = $1;