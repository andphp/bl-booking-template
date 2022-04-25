export interface ApiSearchParams {
  name: string
  path: string
  method: string
  api_group: string
}

export interface ApiItem {
  id: number
  operator: []
  name: string
  path: string
  description: string
  apiGroup: string
  method: string
  createdAt: string
}

export type ApiList = ApiItem[]
