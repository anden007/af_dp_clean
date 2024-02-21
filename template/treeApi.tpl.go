// Code generated by hero.
// source: D:\GoProjects\dp_clean_core\template\treeApi.tpl
// DO NOT EDIT!
package template

import "bytes"

func TreeApi(apiName string, apiPath string, vueName string, buffer *bytes.Buffer) {
	buffer.WriteString(`
// 统一请求路径前缀在libs/axios.js中修改
// import { getRequest, postBodyRequest, postRequest, deleteRequest } from '@/libs/axios';

// `)
	buffer.WriteString(apiName)
	buffer.WriteString(`获取一级数据
export const init`)
	buffer.WriteString(apiName)
	buffer.WriteString(` = (params) => {
    return getRequest('`)
	buffer.WriteString(apiPath)
	buffer.WriteString(`/getByParentId/0', params)
}
// `)
	buffer.WriteString(apiName)
	buffer.WriteString(`加载子级数据
export const load`)
	buffer.WriteString(apiName)
	buffer.WriteString(` = (id, params) => {
    return getRequest('`)
	buffer.WriteString(apiPath)
	buffer.WriteString(`/getByParentId/' + id, params)
}
// `)
	buffer.WriteString(apiName)
	buffer.WriteString(`添加
export const add`)
	buffer.WriteString(apiName)
	buffer.WriteString(` = (params) => {
    return postBodyRequest('`)
	buffer.WriteString(apiPath)
	buffer.WriteString(`/add', params)
}
// `)
	buffer.WriteString(apiName)
	buffer.WriteString(`编辑
export const edit`)
	buffer.WriteString(apiName)
	buffer.WriteString(` = (params) => {
    return postBodyRequest('`)
	buffer.WriteString(apiPath)
	buffer.WriteString(`/edit', params)
}
// `)
	buffer.WriteString(apiName)
	buffer.WriteString(`删除
export const delete`)
	buffer.WriteString(apiName)
	buffer.WriteString(` = (params) => {
    return postRequest('`)
	buffer.WriteString(apiPath)
	buffer.WriteString(`/delByIds/', params)
}
// `)
	buffer.WriteString(apiName)
	buffer.WriteString(`搜索
export const search`)
	buffer.WriteString(apiName)
	buffer.WriteString(` = (params) => {
    return getRequest('`)
	buffer.WriteString(apiPath)
	buffer.WriteString(`/search', params)
}`)

}