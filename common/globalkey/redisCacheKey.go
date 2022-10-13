package globalkey

/**
redis key except "model cache key"  in here,
but "model cache key" in model
*/

// CacheUserTokenKey /** 用户登陆的token
const CacheUserTokenKey = "user_token:%d"

// 手机短信验证码缓存
const CacheSmsPhoneKey = "sms_phone:%s"
const CacheSmsPhoneKeyExp = 5 * 60 //5分钟超时
