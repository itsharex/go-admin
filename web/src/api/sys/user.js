import request from '@/utils/request'
import md5 from 'md5'

// 查询用户列表
export function listUser(query) {
  return request({
    url: '/admin-api/v1/sys-user',
    method: 'get',
    params: query
  })
}

// 查询用户详细
export function getUser(userId) {
  return request({
    url: '/admin-api/v1/sys-user/' + userId,
    method: 'get'
  })
}

export function getUserInit() {
  return request({
    url: '/admin-api/v1/sys-user/',
    method: 'get'
  })
}

// 新增用户
export function addUser(data) {
  return request({
    url: '/admin-api/v1/sys-user',
    method: 'post',
    data: data
  })
}

// 修改用户
export function updateUser(data) {
  return request({
    url: '/admin-api/v1/sys-user/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除用户
export function delUser(data) {
  return request({
    url: '/admin-api/v1/sys-user',
    method: 'delete',
    data: data
  })
}

// 导出用户
export function exportUser(query) {
  return request({
    url: '/admin-api/v1/sys-user/export',
    method: 'get',
    params: query
  })
}

// 用户密码重置
export function resetUserPwd(userId, password) {
  const data = {
    userId,
    password: md5(password)
  }
  return request({
    url: '/admin-api/v1/sys-user/pwd/reset',
    method: 'put',
    data: data
  })
}

// 用户状态修改
export function changeUserStatus(e) {
  const data = {
    userId: e.id,
    status: e.status
  }
  return request({
    url: '/admin-api/v1/sys-user/status',
    method: 'put',
    data: data
  })
}

export function updateProfile(data) {
  return request({
    url: '/admin-api/v1/sys-user/profile',
    method: 'put',
    data: data
  })
}

// 修改用户个人信息
export function updateUserProfile(data) {
  return request({
    url: '/admin-api/v1/sys-user/profile',
    method: 'put',
    data: data
  })
}

// 下载用户导入模板
export function importTemplate() {
  return request({
    url: '/admin-api/v1/sys-user/importTemplate',
    method: 'get'
  })
}

// 用户密码重置
export function updateProfilePwd(oldPassword, newPassword) {
  const data = {
    oldPassword: md5(oldPassword),
    newPassword: md5(newPassword)
  }
  return request({
    url: '/admin-api/v1/sys-user/profile/pwd',
    method: 'put',
    data: data
  })
}

// 用户头像上传
export function updateProfileAvatar(data) {
  return request({
    url: '/admin-api/v1/sys-user/profile/avatar',
    method: 'post',
    data: data
  })
}

// 查询用户个人信息
export function getUserProfile() {
  return request({
    url: '/admin-api/v1/sys-user/profile',
    method: 'get'
  })
}

// 获取验证码
export function getCodeImg() {
  return request({
    url: '/admin-api/v1/captcha',
    method: 'get'
  })
}

// 查询 获取所有配置信息
export function getSetting() {
  return request({
    url: '/admin-api/v1/config/list',
    method: 'get'
  })
}

// login 登陆
export function login(data) {
  data.password = md5(data.password)
  return request({
    url: '/admin-api/v1/login',
    method: 'post',
    data
  })
}

// logout 退出
export function logout() {
  return request({
    url: '/admin-api/v1/sys-user/logout',
    method: 'get'
  })
}

// refreshtoken 刷新token
export function refreshtoken(data) {
  return request({
    url: '/refreshtoken',
    method: 'post',
    data
  })
}
