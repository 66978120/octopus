import request from '@/utils/request'
// 获取模型部署列表
export function getDeployList(params) {
  return request({
    url: `/v1/deploymanage/modeldeploy`,
    method: 'get',
    params
  })
}
// 删除部署服务
export function deleteDeploy(params) {
  return request({
    url: `/v1/deploymanage/modeldeploy`,
    method: 'delete',
    params
  })
}
// 停止部署服务
export function stopDeploy(data) {
  return request({
    url: `/v1/deploymanage/modeldeploy/${data}/stop`,
    method: 'post'
  })
}
// 获取模型部署详情
export function deployDetail(id) {
  return request({
    url: `/v1/deploymanage/modeldeploy/${id}`,
    method: 'get'
  })
}
// 获取模型部署运行信息
export function deployMessage(params) {
  return request({
    url: `/v1/deploymanage/modeldeployevent`,
    method: 'get',
    params
  })
}
// 创建部署服务
export function createDeploy(data) {
  return request({
    url: `/v1/deploymanage/modeldeploy`,
    method: 'post',
    data
  })
}
// 模型服务调用
export function startDeploy(data) {
  return request({
    url: `/v1/deploymanage/modeldeploy/infer`,
    method: 'post',
    data
  })
}
