const error = {
  /* 10000~11000 基础错误*/
  // 1001: '未知错误',
  // 10002: '结构体拷贝失败',
  // 10003: '参数错误',
  // 10004: '版本字符串非法',
  // 10005: '压缩失败',
  // 10006: '解压失败',
  // 10007: 'url解析失败',
  // 10008: '时间解析失败',
  // 10009: '密码加密失败',
  // 10010: '文件夹不存在',
  // 10011: '操作文件夹失败',
  // 10012: '操作文件失败',
  // http请求相关错误
  // 10020: '获取http request实体失败',
  // 10021: 'http请求失败',
  // 10022: 'http读取回包失败',
  // 10023: 'json序列化失败',
  // 10024: 'json反序列化失败',
  // 10025: 'http写入失败',
  // 10026: 'http绑定form失败',
  // minio操作相关错误
  // 10030: 'minio初始化失败',
  // 10031: '桶已存在',
  // 10032: '桶不存在',
  // 10033: '检查桶存在失败',
  // 10034: '创建桶失败',
  // 10035: '删除桶失败',
  // 10036: '生成上传临时url失败',
  // 10037: '生成下载临时url失败',
  // 10038: '查看对象列表失败',
  // 10039: '查看对象失败',
  // 10040: '对象不存在',
  // db操作相关错误
  // 10050: 'db初始化失败',
  // 10051: 'db列表查询失败',
  // 10052: 'db单条查询失败',
  // 10053: 'db查询列表条数失败',
  // 10054: 'db查询为空',
  // 10055: 'db查询条件为空',
  // 10056: 'db更新失败',
  // 10057: 'db插入失败',
  // 10058: 'db删除失败',
  // 10059: 'db操作缺少主键',
  // k8s操作相关错误
  // 10070: 'k8s创建service失败',
  // 10071: 'k8s删除service失败',
  // 10072: 'k8s创建ingress失败',
  // 10073: 'k8s删除ingress失败',
  // Harbor操作相关错误
  // 10080: 'harbor项目已存在',
  // 10081: 'harbor创建项目失败',
  // 10082: 'harbor查询项目失败',
  // redis操作相关错误
  // 10100: 'redis解析url失败',
  // 10101: 'redisHSet失败',
  // 10102: 'redisHGet失败',
  // 10103: 'redisHDel失败',

  /* 11001~12000 资源管理错误*/
  11001: '删除资源池失败',
  11002: '创建资源池失败',
  11003: '查询资源池失败',
  11004: '获取一个资源池失败',
  11005: '更新资源池信息失败',
  11006: '创建资源池失败',
  11007: '创建资源池失败',
  11008: '删除资源失败',
  11009: '更新资源失败',
  11010: '绑定节点失败',
  11011: '创建资源规格失败',
  11012: '获取资源规格列表失败',
  11013: '删除资源规格失败',
  11014: '获取资源规格失败',
  11015: '资源发现失败',
  11016: '资源名字不存在(创建资源规格)',
  11017: '资源规格不存在',
  11018: '资源名已存在(创建自定义资源)',
  11019: '获取节点列表失败',

  /* 12001~13000 算法管理错误*/
  12001: '查找最新公共算法版本失败',
  12002: '查找最新算法版本失败',
  12003: '删除算法不是我的',
  12004: '新增算法重复',
  12005: '新增算法版本重复',
  12006: ' 查找的算法不在权限范围内',
  12007: ' 算法版本文件不存在',
  12008: '算法版本文件已存在',
  12009: '公共算法版本已存在',
  12010: '旧版本算法文件状态未就绪',
  12011: '无权提交此算法版本文件',

  /* 13001~14000 镜像管理错误*/
  13001: '制作镜像状态异常',
  13002: '上传错误镜像来源类型',
  13003: '镜像已存在',
  13004: '镜像不存在',
  13005: '无权限镜像操作',

  /* 14001~15000 开发管理错误*/
  14001: 'NoteBook状态不允许操作',
  14002: 'NoteBook使用镜像不在权限范围内',
  14003: 'NoteBook使用算法不在权限范围内',
  14004: '余额不足',
  14005: '解析资源规格失败',
  14006: 'NoteBook使用镜像状态不允许操作',
  14007: 'NoteBook使用算法状态不允许操作',
  14008: 'NoteBook重复',
  14009: 'NoteBook使用数据集不在权限范围内',
  14010: 'NoteBook使用数据集状态不允许操作',

  /* 15001~16000 训练管理错误*/
  15001: '训练使用镜像不在权限范围内',
  15002: '训练使用数据集不在权限范围内',
  15003: '训练使用算法不在权限范围内',
  // 15004: '训练pipeline请求失败',
  15004: '系统错误，稍后请重试',
  15005: '删除任务失败, 只能删除终结状态下的任务',
  15006: '余额不足',
  15007: '训练任务已终止，无需停止',
  15008: '训练任务所使用镜像,状态不允许操作',
  15009: '训练任务所使用算法,状态不允许操作',
  15010: '训练任务所使用数据集,状态不允许操作',
  15011: '训练任务重名导致的联合索引冲突',
  15022: '训练分布式子任务重名',
  /* 16001~17000 用户管理错误*/
  // 用户登录
  16001: '用户未登录',
  16002: '用户状态不在权限范围内',
  16003: '创建token失败',
  16004: '解析token失败',
  16005: 'token无效',
  16006: 'sessionId为空',
  16007: 'session不存在',
  16008: 'session获取失败',
  16009: '认证失败',
  16010: '认证失败,token被重置',

  // 用户管理
  16020: '用户不存在',
  16021: '用户已存在',
  16022: '用户id错误',
  16023: '空间已存在',
  16024: '空间不存在',
  16025: '用户无空间权限',
  16026: '空间与资源池已绑定',
  /* 17001~18000 计费管理错误*/
  17001: '获取锁失败',
  17002: '状态不允许操作',

  /* 18001~19000 模型管理错误*/
  18001: '没权限操作模型',
  18002: '模型重复',
  18003: '模型版本重复',
  18004: '模型版本文件不存在',
  18005: '模型版本文件已存在',
  18006: '创建模型时归属算法不存在',

  /* 19001~20000 数据集管理错误*/
  19001: '数据集文件不存在',
  19002: '数据集已分享',
  19003: '没有权限操作',
  19004: ' 数据集重复',
  19005: '状态不允许操作',

  /* 20001-21000 第三方平台管理错误*/
  20001: '平台名称重复',
  20002: '平台存储配置名称重复',
  20003: '批量获取平台信息错误',
  20004: '配置值不正确',
  20005: '配置项不存在',
  20006: 'http请求失败',

  /* 21001-22000 云际错误*/
  21001: '云际请求失败',
  21002: '无权限访问',
  // 标签错误发
  19501: '标签被引用，不能删除',
  19502: '标签重复',
  19503: '标签不合法',
  19504: '预置标签不可更改'
}
// export function getErrorMsg(errorCode) {
//   let message = ''
//   for (var p in error) {
//     if (String(errorCode) === p) {
//       message = error[p]
//       return message
//     }
//   }
//   return '系统错误'
// }
export const mixin = {
  methods:{
    getErrorMsg(errorCode){
      let message = ''
      for (var p in error) {
        if (String(errorCode) === p) {
          message = error[p]
          return message
        }
      }
      return '系统错误'
    }
  }
}
