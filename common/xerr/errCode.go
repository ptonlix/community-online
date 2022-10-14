package xerr

//成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

//全局错误码
const SERVER_COMMON_ERROR uint32 = 100001
const REUQEST_PARAM_ERROR uint32 = 100002
const DB_ERROR uint32 = 100005
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100006

//用户模块 110 代码
const USER_MESSAGE_ERROR uint32 = 110001          //短信验证码校验失败
const TOKEN_EXPIRE_ERROR uint32 = 110002          // TOKEN过期
const TOKEN_GENERATE_ERROR uint32 = 110003        // TOKEN生成失败
const USER_LOGIN_VERIFY_ERROR uint32 = 110004     // 账号或密码验证失败
const CAPTCHA_GENERATE_ERROR uint32 = 110005      // 图形校验码生成失败
const CAPTCHA_VERIFY_ERROR uint32 = 110006        // 图形校验码校验失败
const USER_NOT_EXIST_ERROR uint32 = 110007        // 用户不存在
const USER_ALREADY_REGISTER_ERROR uint32 = 110008 // 用户已注册
