import { Device } from '@/models'
import { MenuChild } from '@/models/menu.interface'
import { PureSettings } from '@ant-design/pro-layout/lib/defaultSettings'
import { Role } from './login'

export type Locale = 'zh-cn' | 'en-us'

export interface CurrentUserResult {
  username: string
  role: Role
  accountName: string
  activeColor: string
  avatar: string
  baseColor: string
  departmentID: number
  email: string
  emailVerifiedAt: string
  fullName: string
  id: number
  lastIP: string
  lastLoginAt: string
  lastToken: string
  nickName:string
  phone: string
  roleID: number
  sideMode: string
  uuid: string
}

export interface User {
  username: string

  /** menu list for init tagsView */
  menuList: MenuChild[]

  /** login status */
  logged: boolean

  role: Role

  /** user's device */
  device: Device

  /** menu collapsed status */
  collapsed: boolean

  /** notification count */
  noticeCount: number

  /** user's language */
  locale: Locale

  /** Is first time to view the site ? */
  newUser: boolean

  settings: PureSettings
  avatar: string
}
