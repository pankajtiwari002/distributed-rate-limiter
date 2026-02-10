local key = KEYS[1]
local capacity = tonumber(ARGV[1])
local refill_rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])

local data = redis.call("HMGET", key, "tokens", "last_time")

local tokens = tonumber(data[1]) or capacity
local last_time = tonumber(data[2]) or now

local delta = math.max(0, now - last_time)
local refill = delta * refill_rate
tokens = math.min(capacity, tokens + refill)

local allowed = 0
if tokens >= 1 then
	tokens = tokens - 1
	allowed = 1
end

redis.call("HMSET", key, "tokens", tokens, "last_time", now)
redis.call("EXPIRE", key, 120)

return allowed
