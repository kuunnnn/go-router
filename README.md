# go-router


## match

`///user///info///` === `/user/info` === `user/info`

1. `/user/phone/`
2. `/user/name/`
3. `/user/:id/`
4. `/user/*AvatarUrl/`

: 和 * 同时存在同一级则无法进入 *
