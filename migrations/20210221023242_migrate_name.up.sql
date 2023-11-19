CREATE TABLE "posts" (
    "post_id" bigint ,
    "word" varchar not null ,
    "count"   integer not null,
    PRIMARY KEY ("post_id","word")
);