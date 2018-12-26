CREATE TABLE IF NOT EXISTS users (
  user_id	          VARCHAR(36) PRIMARY KEY,
  access_token      VARCHAR(512) DEFAULT '',
  full_name         VARCHAR(512),
  avatar_url        VARCHAR(1024),
  telegraph_token   VARCHAR(64),
  created_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  expire_at         TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW() 
);

CREATE INDEX ON users (user_id);

CREATE TABLE IF NOT EXISTS posts (
  post_id         VARCHAR(36) PRIMARY KEY,
  title	          VARCHAR(256) NOT NULL,
  path	          VARCHAR(256),
  telegraph_url	  VARCHAR(1024),
  description	    VARCHAR(1024),
  ipfs_id         VARCHAR(64),
  trace_id          VARCHAR(36),
  content         TEXT NOT NULL,
  user_id	        VARCHAR(36) NOT NULL,
  created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
  updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX ON posts (post_id, user_id);

CREATE TABLE IF NOT EXISTS subscribers (
  user_id	        VARCHAR(36) NOT NULL,
  subscriber_id     VARCHAR(36) NOT NULL,
  created_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (user_id, subscriber_id)
);

CREATE TABLE IF NOT EXISTS supports (
  support_id            	VARCHAR(36) PRIMARY KEY,
  supporter_id	        	VARCHAR(36) NOT NULL,
  author_id					VARCHAR(36) NOT NULL,
  post_id					VARCHAR(36) NOT NULL,
  asset_id         			VARCHAR(36) NOT NULL,
  trace_id         			VARCHAR(36) NOT NULL,
  amount        			VARCHAR(256),
  created_at        	TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
