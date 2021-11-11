import request from '@/utils/request'

export function judgeParam(params) {
  let conditions = []
  conditions.push(`pageSize=` + params.pageSize);
  conditions.push(`pageIndex=` + params.pageIndex);
  params.orderBy?conditions.push(`orderBy=` + params.orderBy):null;
  params.sortBy?conditions.push(`sortBy=` + params.sortBy):null;
  params.searchKey?conditions.push(`searchKey=` + params.searchKey):null;
  params.createdAtGte?conditions.push(`createdAtGte=` + params.createdAtGte):null;
  params.createdAtLt?conditions.push(`createdAtLt=` + params.createdAtLt):null;
  params.status?conditions.push(`status=` + params.status):null;
  params.fileStatus?conditions.push(`fileStatus=` + params.fileStatus):null;
  params.algorithmVersion?conditions.push(`algorithmVersion=` + params.algorithmVersion):null;
  return conditions
}

export async function getNotebookList(payload){
  let conditions = judgeParam(payload)
  const res = await request({
    url: "/v1/developmanage/notebook?" + conditions.join("&"),
    method: "get",  
  })
  return res
}

export async function stopNotebook(id){
  const res = await request({
    url: `/v1/developmanage/notebook/${id}/stop`,
    method: "post",  
  })
  return res
}

export async function getUserAlgorithmList(payload){
  let conditions = judgeParam(payload)
  const res = await request({
    url: `/v1/algorithmmanage/allalgorithm?` + conditions.join("&"),
    method: "get",  
  })
  return res
}

export async function getPresetAlgorithmList(payload){
  let conditions = judgeParam(payload)
  const res = await request({
    url: "/v1/algorithmmanage/prealgorithm?" + conditions.join("&"),
    method: "get",  
  })
  return res
}

export async function getAlgorithmVersionList(payload){
  let conditions = judgeParam(payload)
  const res = await request({
    url: `/v1/algorithmmanage/algorithm/${payload.algorithmId}?` + conditions.join("&"),
    method: "get",  
  })
  return res
}

export async function queryAlgorithmVersion(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithm/${payload.algorithmId}/version/${payload.version}`,
    method: 'get',
  })
  return res
}

export async function addPreAlgorithm(payload){
  const res = await request({
    url: `/v1/algorithmmanage/prealgorithm`,
    method: "post",  
    data: payload 
  })
  return res
}

export async function addPreAlgorithmVersion(payload){
  const res = await request({
    url: `/v1/algorithmmanage/prealgorithm/${payload.algorithmId}`,
    method: "post",  
    data: {
      oriVersion : payload.oriVersion,
      algorithmDescript : payload.algorithmDescript
    } 
  })
  return res
}

export async function uploadPreAlgorithm(payload){
  const res = await request({
    url: `/v1/algorithmmanage/prealgorithm/${payload.algorithmId}/version/${payload.version}/upload`,
    method: "post",  
    data: payload 
  })
  return res
}

export async function preAlgorithmFinishUpload(payload){
  const res = await request({
    url: `/v1/algorithmmanage/prealgorithm/${payload.algorithmId}/version/${payload.version}/uploadconfirm`,
    method: "put",
    data: payload
  })
  return res
}

export async function compressAlgorithm(payload){
  const res = await request({
    url: `/v1/algorithmmanage/algorithm/${payload.algorithmId}/version/${payload.version}/downloadcompress`,
    method: "post",
  })
  return res
}

export async function downloadAlgorithmVersion(payload){
  const res = await request({
    url: `/v1/algorithmmanage/algorithm/${payload.algorithmId}/version/${payload.version}/download?domain=${payload.domain}&compressAt=${payload.compressAt}`,
    method: "get",
  })
  return res
}

export async function deletePreAlgorithmVersion(payload){
  const res = await request({
    url: `/v1/algorithmmanage/prealgorithm/${payload.algorithmId}/version/${payload.version}`,
    method: "delete",
  })
  return res
}

export async function deletePreAlgorithm(algorithmId){
  const res = await request({
    url: `/v1/algorithmmanage/prealgorithm/${algorithmId}`,
    method: "delete",
  })
  return res
}