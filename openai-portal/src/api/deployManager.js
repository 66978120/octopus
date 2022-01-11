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
export function deleteDeploy(data) {
  return request({
    url: `/v1/deploymanage/deploymodel`,
    method: 'post',
    data
  })
}
// 停止部署服务
export function stopDeploy(data) {
  return request({
    url: `/v1/deploymanage/deploymodel/${data.id}/stop`,
    method: 'post',
    data
  })
}
// 获取模型部署详情
export function deployDetail(params) {
  return request({
    url: `/v1/deploymanage/modeldeploy/{params}`,
    method: 'post',
    params
  })
}
// 获取模型部署事件列表
export function deployEvent(params) {
  return request({
    url: `/v1/deploymanage/modeldeployevent`,
    method: 'get',
    params
  })
}
// 创建部署服务
export function createDeploy(data) {
  return request({
    url: `/v1/deploymanage/deploymodel`,
    method: 'post',
    data
  })
}
// 模型服务调用
export function startDeploy(data) {
  return request({
    url: `/v1/deploymanage/deploymodel/infer`,
    method: 'get',
    data
  })
}
