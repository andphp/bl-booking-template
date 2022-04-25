import { MenuList } from '@/models/menu.interface'
import { LoginParams, LoginResult } from '@/models/login'
import { CurrentUserResult } from '@/models/user'
import { useBatch, useCreate, useGetList, useGetOne, useUpdate } from '../utils/request'

const projectResource = '/projects'

export const useLogin = () => {
  return useCreate<LoginParams, LoginResult>('/auth/sign/in')
}

export const useGetCurrentUser = () => {
  return useGetOne<CurrentUserResult>('CurrentUser', '/auth/user/info')
}

export const useGetCurrentMenus = () => {
  return useGetList<MenuList>('CurrentMenuList', '/auth/menu')
}
export const useGetProjects = (pagination: any, filters: any) => {
  return useGetList<API.ProjectPagination>('Projects', projectResource, pagination, filters)
}
export const useAddProject = () => {
  return useCreate<API.Project, API.Project>(projectResource)
}

export const useUpdateProject = () => {
  return useUpdate<API.Project>(projectResource)
}

export const useBatchDeleteProject = () => {
  return useBatch(projectResource + ':batchDelete')
}
